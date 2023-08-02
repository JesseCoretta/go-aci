package aci

/*
parse.go acts as a bridge between the go-aci package and the ANTLR4
parser/lexer subsystems for various ACIv3 and LDAP components.
*/

import (
	//"bytes"
	//"github.com/JesseCoretta/go-aci/internal/ldapparser"
	"github.com/JesseCoretta/go-aci/internal/aciparser"
)

/*
parseTR reads and processes tokens into a sequence of zero (0) or more
Target Rule instances, which are returned alongside an error.

The chop return value informs the calling method which index integer
from which to continue, thus avoiding tokens already processed.

Target Rules are not nested; if a non-zero targetRules stack is returned,
it shall contain a like-number of Condition instances, each bearing a
target keyword (e.g.: `targetattr`), an equality (=) OR negated (!=)
comparison operator and a doubled-quoted expression value. Note that Target
Rule conditions are **ALWAYS** parenthetical.
*/
func parseTR(tokens []string) (chop int, targetRules Rule, err error) {

	// vals stores a sequence of value (and value RELATED)
	// tokens when detected. This var is purged of its
	// contents whenever the end of the value sequence is
	// reached.
	var vals []string = make([]string, 0)

	// Initialize our target rule stack; max capacity is nine (9)
	targetRules = T()

	var kw string   // target rule keyword
	var cop string  // target rule comparison operator
	var last string // previous token
	var next string // upcoming token
	var cready bool // condition ready for assembly

	var done bool
	for index, token := range tokens {
		if index > 0 {
			last = tokens[index-1]
		}

		if index+1 < len(tokens) {
			// look-ahead to the next iteration's token
			next = tokens[index+1]
		}

		_, istok := matchOp(token)

		switch {

		// token is a target rule keyword
		case matchTKW(lc(token)) != TargetKeyword(0x0):
			kw = lc(token)

		// token is a target rule comparison operator
		case istok:
			cop = token
			if cready = len(kw) > 0 && len(cop) > 0; !cready {
				err = errorf("Target Rule comparison operator and/or keyword not detected near '%s'; cannot declare target rule condition readiness", token)
				return
			}

		// token is anchor, meaning there are no more
		// target rules to process.
		case token == `version 3.0; acl`:
			if chop = index; chop > 0 {
				done = true
				break
			}

		// generalized token fallback will capture quoted values as well as
		// logical symbolic operators (||, &&) between quoted values (or as
		// part of a quoted value).
		default:
			if cready {
				// condition is ready for assembly AND a non-zero
				// double-quoted value is the current token. We
				// will also accept symbolic operators if we're
				// dealing with a multi-valued expression.
				if !isQuoted(token) && (token != `||` && token != `&&`) {
					err = errorf("Bogus Target Rule condition expression between '%s' [%d] and '%s' [%d]; value must be a non-zero string enclosed within double quotes, or a symbolic list (||,&&) of same",
						last, index-1, next, index+1)
					return
				}
				// strip quotes, as go-stackage provides encaps
				// without the need for literal storage of such
				// characters.

				// increment chop index by one (1)
				chop++

				// Save this value; we don't yet know if this
				// value is merely one (1) of multiple values
				// as opposed to a single value alone.
				vals = append(vals, token)

				// Look ahead to see what is coming next. If
				// another quoted value or symbolic operator
				// are detected, we know we're not done yet.
				// In that case, break out of this case to
				// continue at the next for-loop iteration.
				if next == `||` || next == `&&` || isQuoted(next) {
					break
				}

				// Keep track of when individual values, not
				// an entire value, are quoted in a list.
				var vEncapsulated bool

				// assert the comparison operator
				oper, mo := matchOp(cop)
				if !mo {
					err = errorf("Unidentified or misaligned target rule comparison operator '%s'; aborting",
						cop)
					return
				}

				// This is the last (or only!) value component. We can
				// now analyze the keyword and the value(s) to ascertain
				// the appropriate instance type for condition storage
				// (and to perform other context-specific sanity checks).
				//
				// Begin with an assertion switch upon the target keyword
				// (which we already vetted as sane) ...
				switch key := matchTKW(kw); key {

				case TargetScope:
					if len(vals) != 1 {
						err = errorf("Unexpected number of %s search scopes; want %d, got %d",
							key, 1, len(vals))
						return
					}

					scn := unquote(vals[0])
					sc := strToScope(scn)

					// base is a fallback for a bogus scope, so
					// if the user did not originally request
					// base, we know they requested something
					// totally unsupported.
					if sc.String() == `base` && !eq(scn, `base`) {
						err = errorf("Bogus %s value: '%s'", key, scn)
						return
					}

					targetRules.Push(sc.Eq().Encap(`"`).Paren(true))

				case TargetFilter:
					if len(vals) != 1 {
						err = errorf("Unexpected number of %s search filters; want %d, got %d",
							key, 1, len(vals))
						return
					}

					f := unquote(vals[0])
					targetRules.Push(TFilter().Push(f).Eq())

				case TargetAttr:
					var tat Rule = TAttrs().Encap()

					// target rule is either or both of the following:
					// A: one (1) double-quoted AT
					// B: one (1) double-quoted LIST of unquoted ATs in symbolic OR context
					for x := 0; x < len(vals); x++ {
						var value string = vals[x]

						if contains(value, `||`) {

							// Type-B confirmed
							for ix, O := range split(unquote(value), `||`) {
								if len(O) == 0 {
									continue
								}

								if x == 0 && ix == 0 {
									if !isQuoted(vals[x]) && isQuoted(O) {
										vEncapsulated = true
										tat.Encap()
									} else if !isQuoted(O) {
										tat.Encap(`"`)
									}
								}

								tat.Push(ATName(trimS(unquote(O))))
							}

						} else {
							// Type-A confirmed
							if x == 0 {
								if isQuoted(value) {
									vEncapsulated = true
									tat.Encap(`"`)
								}
							}
							tat.Push(ATName(trimS(unquote(value))))
						}
					}

					var c Condition
					if oper == Ne {
						if !vEncapsulated {
							c = tat.Ne().Encap(`"`)
						} else {
							c = tat.Ne().Encap()
						}
					} else {
						if !vEncapsulated {
							c = tat.Eq().Encap(`"`)
						} else {
							c = tat.Eq().Encap()
						}
					}

					targetRules.Push(c)

				case TargetCtrl, TargetExtOp:
					var toid Rule
					if key == TargetExtOp {
						toid = ExtOps()
					} else {
						toid = Ctrls()
					}

					// target rule is either or both of the following:
					// A: one (1) double-quoted OID
					// B: one (1) double-quoted LIST of unquoted OIDs in symbolic OR context
					for x := 0; x < len(vals); x++ {
						var value string = vals[x]

						if contains(value, `||`) {

							// Type-B confirmed
							for ix, O := range split(unquote(value), `||`) {
								if len(O) == 0 {
									continue
								}

								if x == 0 && ix == 0 {
									if !isQuoted(vals[x]) && isQuoted(O) {
										vEncapsulated = true
										toid.Encap()
									} else if !isQuoted(O) {
										toid.Encap(`"`)
									}
								}

								value = trimS(unquote(O))
								o, _ := newObjectID(key, value)
								toid.Push(ObjectIdentifier{o})
							}
						} else {
							if x == 0 {
								if isQuoted(value) {
									vEncapsulated = true
									toid.Encap(`"`)
								}
							}

							// Type-A confirmed
							value = trimS(unquote(value))
							o, _ := newObjectID(key, value)
							toid.Push(ObjectIdentifier{o})
						}
					}

					var c Condition
					if oper == Ne {
						if !vEncapsulated {
							c = toid.Ne().Encap(`"`)
						} else {
							c = toid.Ne().Encap()
						}
					} else {
						if !vEncapsulated {
							c = toid.Eq().Encap(`"`)
						} else {
							c = toid.Eq().Encap()
						}
					}

					targetRules.Push(c)

				case Target, TargetTo, TargetFrom:
					var tdnr Rule = TDNs().setCategory(key.String())

					// target rule is either or both of the following:
					// A: one (1) double-quoted DN
					// B: one (1) double-quoted LIST of unquoted DNs in symbolic OR context
					for x := 0; x < len(vals); x++ {
						var value string = vals[x]
						if contains(value, `||`) {

							// Type-B confirmed
							for ix, O := range split(unquote(value), `||`) {
								if len(O) == 0 {
									continue
								}

								if x == 0 && ix == 0 {
									if !isQuoted(vals[x]) && isQuoted(O) {
										vEncapsulated = true
										tdnr.Encap()
									} else if !isQuoted(O) {
										tdnr.Encap(`"`)
									}
								}

								D := trimS(unquote(O))
								if !hasPfx(D, LocalScheme) {
									err = errorf("Illegal %s distinguishedName slice: [index:%d;value:%s] missing LDAP local scheme (%s)",
										key, x, D, LocalScheme)
									return
								}

								tdnr.Push(DistinguishedName{newDistinguishedName(D[len(LocalScheme):], key)})
							}

						} else {

							// Type-A confirmed
							if x == 0 {
								if isQuoted(value) {
									vEncapsulated = true
									tdnr.Encap(`"`)
								}
							}

							D := unquote(value)
							if !hasPfx(D, LocalScheme) {
								err = errorf("Illegal %s distinguishedName: [index:%d;value:%s] missing LDAP local scheme (%s)",
									key, x, D, LocalScheme)
								return
							}

							tdnr.Push(DistinguishedName{newDistinguishedName(D[len(LocalScheme):], key)})
						}
					}

					var c Condition
					if oper == Ne {
						if !vEncapsulated {
							c = tdnr.Ne().Encap(`"`)
						} else {
							c = tdnr.Ne().Encap()
						}
					} else {
						if !vEncapsulated {
							c = tdnr.Eq().Encap(`"`)
						} else {
							c = tdnr.Eq().Encap()
						}
					}

					targetRules.Push(c)

				case TargetAttrFilters:
					// TODO
					//if len(vals) != 1 {
					//	err = errorf("Target Rule keyword '%s' supports single values only, but %d values were found: %v",
					//		kw,len(vals),vals)
					//	return
					//}
				default:
					err = errorf("Unhandled target rule type '%s'", key)
					return
				}

				// Reset for next target rule condition, if any
				kw = ``
				oper = nil
				cready = false
				vals = make([]string, 0)
			}
		}

		if done {
			break
		}
	}

	return
}

/*
parseBR processes the input slice values (tokens) into a Bind Rule stack (outer).
It, alongside an error and chop index, are returned when processing stops or completes.
*/
func parseBR(tokens []string, pspan int) (chop int, outer Rule, err error) {
	// Don't bother processing tokens that could
	// never possibly represent a valid Bind Rule
	// expression (UNLESS we're recursing within
	// a parenthetical bind rule expression).
	if len(tokens) < 4 && pspan == 0 {
		err = errorf("Empty bind rule input value, or value is below minimum possible length for validity: %v [%d<4]", tokens, len(tokens))
		return
	}

	// oparen remembers whether the entirely of the
	// bind rule statement, whether nested or not,
	// is parenthetical.
	var oparen bool
	if len(tokens) > 4 {
		 oparen = (tokens[0] == `(` && tokens[4] != `)` && tokens[len(tokens)-3] == `)` && pspan <2)
	}

	// Create temporary storage vars for some of
	// the Condition components that will need
	// to be preserved across loops.
	var kw string                         // Bind Rule Condition keyword
	var cop string                        // Bind Rule Condition comparison operator
	var vals []string = make([]string, 0) // Bind Rule Condition expression(s)

	// convenient true/false bind rule keyword
	// recognizer func.
	isKW := func(o string) bool {
		return matchBKW(lc(o)) != BindKeyword(0x0)
	}

	// convenient comparison operator token
	// recognizer func.
	isTokOp := func(o string) (ok bool) {
		_, ok = matchOp(o)
		return
	}

	// Find (and remember) the first (1st)
	// Boolean WORD operator encountered.
	var bopf bool
	for i := 0; i < len(tokens); i++ {
		if isWordOp(tokens[i]) {
			// a known Boolean WORD operator has
			// been found; create the outer stack
			// accordingly.
			outer = ruleByLoP(tokens[i])
			bopf = true
			break
		}
	}

	// If no Boolean WORD operator was
	// encountered, just fallback to AND
	// for convenience.
	if !bopf {
		outer = ruleByLoP(`AND`)
	}

	var cparen bool // parenthetical condition marker
	var iparen bool // parenthetical inner stack marker

	// Create a temporary map to store
	// Condition and Rule instances.
	slices := make(map[int]any, 0)

	// Whenever a valid Condition is ready to be assembled
	// and pushed into the outer stack, this func is called.
	pushBR := func(k, id string, t any, o any) (err error) {

		// Assemble Condition instance c
		// using keyword k, token value t
		// and comparison operator o.
		c := Cond(k, t, o).
			setID(id).
			Paren(cparen).
			NoPadding(!ConditionPadding)

		// be double certain the condition
		// is truly valid, else we do NOT
		// want to add it to the stack.
		if err = c.Valid(); err != nil {
			return
		}

		// Assignh Condition c into temporary
		// component map
		slices[len(slices)] = c
		return
	}

	// make it known when we're ready to push a new
	// (complete) condition into the outer stack.
	var cready bool

	// certain methods executed recursively (such as
	// for nested elements) shall return a "chop index"
	// which defines the integer index to which we
	// should "jump ahead" once the recursion(s)
	// finish.
	var skipTo int

	var next string
	var last string
	var seen []string

	// iterate each of the tokens.
	var ct int = -1
	for _, token := range tokens {
		// actual iteration counter. We're unable to
		// rely on range index because the iterable
		// var (tokens) is continuously truncated
		// along the way.
		ct++

		if len(tokens) > 1 {
			// look-ahead to the next iteration's token
			next = tokens[1]
		}

		switch {

		// token is a parenthetical opener. This case switch
		// shall launch a new parseBR (recursive) thread in
		// which nested (superior) parentheticals are found.
		case token == `(`:
			chop++
			pspan++

			// Before beginning this loop, check to see if the fifth
			// (5th) token is a Boolean WORD operator. Also check to
			// see if there is a parenthetical Condition by itself.
			if len(tokens) >= 5 {

				if next == `(` {
					// Launch a new inner recursion of this
					// same function.
					var inner Rule
					if skipTo, inner, err = parseBR(tokens[1:], pspan); err != nil {
						return
					}

					// Done processing!
					if skipTo-1 == len(tokens) {
						chop = skipTo - 2
						tokens = tokens[len(tokens)-2:]
						if inner.Len() > 0 {
							slices[len(slices)] = inner // save stack
						}
						break
					}

					// If the inner stack has at least one
					// (1) element, preserve it for the end
					// stack element, else take no action.
					if inner.Len() > 0 {
						inner.Paren(true)
						tokens = tokens[skipTo:]    // truncate tokens already processed through recursion
						chop += skipTo              // sum our "skip to" index with our return chop index
						slices[len(slices)] = inner // save stack
					}

					break
				}

				// evaluate inner parentheticals
				iparen = isWordOp(tokens[4]) && tokens[len(tokens)-3] != `)`

				// evaluate condition parentheticals
				cparen = isKW(tokens[1]) && tokens[4] == `)`
			}

		// token is a parenthetical closer. Match only if we're
		// NOT in the middle of Condition assembly.
		case token == `)` && !cready:
			chop++

			// perform some checks to ensure the number
			// of parentheticals ( '(' and ')' ) are equal
			// (i.e.: one closer (R) for every opener (L)).
			//
			// The ONE AND ONLY exception is the check that
			// discounts the semicolon terminator as the last
			// (previous) token, as that indicates the closure
			// of the anchor (pre-bind rule) and the end of
			// the ACI as a whole. While this function doesn't
			// see those tokens, we already knew to expect it,
			// thus we can handle that condition here.
			if pspan < 0 {
				err = errorf("Unbalanced parenthetical expression detected near or within '%s%s%s' (token:%d, negative paren)",
					last, token, next, ct)
				return
			} else if pspan == 0 && last != `;` {
				err = errorf("Unbalanced parenthetical expression detected near or within '%s%s%s' (token:%d, zero paren)",
					last, token, next, ct)
				return
			}

			// pspan is known to be a positive, non-zero
			// integer, so it is safe to decrement. But
			pspan--

			// If decrementing pspan by one (1) results
			// in a zero value, then we have exited the
			// parenthetical (stack) expression.
			if pspan == 0 {
				break
			}

			// evaluate condition parentheticals inside
			// (superior) closing parenthetical.
			if len(seen) >= 3 {

				cparen = seen[len(seen)-4] == `(` &&
					isKW(seen[len(seen)-3]) &&
					isQuoted(seen[len(seen)-1]) &&
					pspan > 0
			}

			// If the NEXT token is a logical Boolean WORD
			// operator, we take special steps because we
			// are currently within a parenthetical bind
			// rule expression component.
			if isWordOp(next) {
					chop++
					var ttoken string = next
					if eq(next, `and not`) {
						// go-stackage's negation stacks use the
						// category of 'NOT', as opposed to ACIv3's
						// 'AND NOT' operator equivalent. Take the
						// 'NOT' portion of the value, using its
						// original case-folding scheme, and save
						// it for stack tagging later.
						ttoken = ttoken[4:]
					}

					// If the category (word operator) is not
					// the same as the token, this means a new
					// distinct (inner) stack is beginning (and
					// not a continuation of outer).
					if !eq(ttoken, outer.Category()) {
						// We need to offset the truncation factor
						// of our token slices when the 'AND NOT'
						// logical Boolean WORD operator is used,
						// as it will erroneously be interpreted
						// as two (2) distinct tokens.
						var offset int
						if eq(ttoken, `not`) {
							offset+=2 // +2 because we're in "look ahead" mode (and 'not' is 'and not').
						}

						// Launch a new inner recursion of this
						// same function.
						var inner Rule
						if skipTo, inner, err = parseBR(tokens[offset:], pspan); err != nil {
							return
						}

						// If the inner stack has at least one
						// (1) element, preserve it for the end
						// stack element, else take no action.
						if inner.Len() > 0 {
							inner.setCategory(ttoken)   // mark the inner stack's logical Boolean WORD operator
							tokens = tokens[skipTo:]    // truncate tokens already processed through recursion
							chop += skipTo              // sum our "skip to" index with our return chop index
							slices[len(slices)] = inner // save stack
						}
					} else if tokens[1] == `(` && tokens[5] != `)` {
		                                // Launch a new inner recursion of this
		                                // same function.
		                                var inner Rule
		                                if skipTo, inner, err = parseBR(tokens[1:], pspan); err != nil {
		                                        return
		                                }

		                                // If the inner stack has at least one
		                                // (1) element, preserve it for the end
		                                // stack element, else take no action.
		                                if inner.Len() > 0 {
							inner.Paren(true)
		                                        inner.setCategory(ttoken)   // mark the inner stack's logical Boolean WORD operator
		                                        tokens = tokens[skipTo:]    // truncate tokens already processed through recursion
		                                        chop += skipTo              // sum our "skip to" index with our return chop index
		                                        slices[len(slices)] = inner // save stack
		                                }
		                        }
			}

		// token is a keyword
		case isKW(token):
			chop++
			kw = lc(token)

		// token is a comparison operator
		case isTokOp(token):
			chop++
			cop = token
			cready = len(kw) > 0 && len(cop) > 0

		// token is a boolean word operator
		case isWordOp(token) && !cready:
			chop++

			var ttoken string = token

			if eq(token, `and not`) {
				// go-stackage's negation stacks use the
				// category of 'NOT', as opposed to ACIv3's
				// 'AND NOT' operator equivalent. Take the
				// 'NOT' portion of the value, using its
				// original case-folding scheme, and save
				// it for stack tagging later.
				ttoken = ttoken[4:]
			}

			// If the category (word operator) is not
			// the same as the token, this means a new
			// distinct (inner) stack is beginning (and
			// not a continuation of outer).
			if !eq(ttoken, outer.Category()) {

				// We need to offset the truncation factor
				// of our token slices when the 'AND NOT'
				// logical Boolean WORD operator is used,
				// as it will erroneously be interpreted
				// as two (2) distinct tokens.
				var offset int
				if eq(ttoken, `not`) {
					offset++
				}

				// Launch a new inner recursion of this
				// same function.
				var inner Rule
				if skipTo, inner, err = parseBR(tokens[offset:], pspan); err != nil {
					return
				}

				// If the inner stack has at least one
				// (1) element, preserve it for the end
				// stack element, else take no action.
				if inner.Len() > 0 {
					inner.setCategory(ttoken)   // mark the inner stack's logical Boolean WORD operator
					tokens = tokens[skipTo:]    // truncate tokens already processed through recursion
					chop += skipTo              // sum our "skip to" index with our return chop index
					slices[len(slices)] = inner // save stack
				}

			} else if tokens[1] == `(` && tokens[5] != `)` {
                                // Launch a new inner recursion of this
                                // same function.
                                var inner Rule
                                if skipTo, inner, err = parseBR(tokens[1:], pspan); err != nil {
                                        return
                                }

                                // If the inner stack has at least one
                                // (1) element, preserve it for the end
                                // stack element, else take no action.
                                if inner.Len() > 0 {
					inner.Paren(true)
                                        inner.setCategory(ttoken)   // mark the inner stack's logical Boolean WORD operator
                                        tokens = tokens[skipTo:]    // truncate tokens already processed through recursion
                                        chop += skipTo              // sum our "skip to" index with our return chop index
                                        slices[len(slices)] = inner // save stack
                                }
			}

		// token is a semicolon, which means the end of a PermissionBindRule
		// instance. We don't need to do anything, and we don't need to keep
		// this token, so we match it separately. DO NOT match if we are in
		// the middle of Condition assembly.
		case token == `;` && !cready:

		// generalized token fallback will capture quoted values as well as
		// logical symbolic operators (||, &&) between quoted values (or as
		// part of a quoted value).
		default:
			if cready {
				// condition is ready for assembly AND a non-zero
				// double-quoted value is the current token. We
				// will also accept symbolic operators if we're
				// dealing with a multi-valued expression.
				if !isQuoted(token) && (token != `||` && token != `&&`) {
					err = errorf("Bogus Bind Rule condition expression between '%s' [%d] and '%s' [%d]; value must be a non-zero string enclosed within double quotes, or a symbolic list (||,&&) of same",
						last, ct-1, next, ct+1)
					return
				}

				// strip quotes, as go-stackage provides encaps
				// without the need for literal storage of such
				// characters.

				// increment chop index by one (1)
				chop++

				// Look ahead to see what is coming next. If
				// another symbolic operator is detected, we
				// know we're not done yet. In that case, we
				// break out of this case to continue at the
				// next for-loop iteration.
				if next == `||` || next == `&&` {
					// keep the delim for now, we'll need
					// it for quote analysis!
					vals = append(vals, token)
					break
				} else if !isQuoted(next) && next != `;` && next != `)` {
					err = errorf("Misaligned value expression [%s -> %s]", token, next)
					return
				}

				// Save this value; we don't yet know if this
				// value is merely one (1) of multiple values
				// as opposed to a single value alone.
				vals = append(vals, token)

				// assert the comparison operator
				oper, mo := matchOp(cop)
				if !mo {
					err = errorf("Unidentified or misaligned target rule comparison operator '%s'; aborting",
						cop)
					return
				}

				// ascertain our parenthetical status, as well
				// as our circumscribing (enclosing) rule, if
				// applicable ...
				var id string = `bind`
				if (oparen || iparen) && cparen {
					id = `enveloped_parenthetical_bind`
				} else if (oparen || iparen) && !cparen {
					id = `enveloped_bind`
				} else if cparen {
					id = `parenthetical_bind`
				}

				// Keep track of when individual values, not
				// an entire value, are quoted in a list.
				var vEncapsulated bool

				// This is the last (or only!) value component. We can
				// now analyze the keyword and the value(s) to ascertain
				// the appropriate instance type for condition storage
				// (and to perform other context-specific sanity checks).
				//
				// Begin with an assertion switch upon the target keyword
				// (which we already vetted as sane) ...
				switch key := matchBKW(kw); key {

				case BindUDN, BindRDN, BindGDN:

					// prepare a stack for our DN value(s)
					var bdn Rule
					if key == BindRDN {
						bdn = RDNs()
					} else if key == BindGDN {
						bdn = GDNs()
					} else {
						bdn = UDNs()
					}

					// bind rule is either or both of the following:
					// A: one (1) double-quoted DN
					// B: one (1) double-quoted LIST of unquoted DNs in symbolic OR context
					for x := 0; x < len(vals); x++ {
						var value string = vals[x]
						if contains(value, `||`) {

							// Type-B confirmed
							for ix, O := range split(unquote(value), `||`) {
								if len(O) == 0 {
									continue
								}

								if x == 0 && ix == 0 {
									if !isQuoted(vals[x]) && isQuoted(O) {
										vEncapsulated = true
										bdn.Encap()
									} else if !isQuoted(O) {
										bdn.Encap(`"`)
									}
								}

								D := trimS(unquote(O))
								if !hasPfx(D, LocalScheme) {
									err = errorf("Illegal %s distinguishedName slice: [index:%d;value:%s] missing LDAP local scheme (%s)",
										key, x, D, LocalScheme)
									return
								}

								bdn.Push(DistinguishedName{newDistinguishedName(D[len(LocalScheme):], key)})
							}

						} else {

							// Type-A confirmed
							if x == 0 {
								if isQuoted(value) {
									vEncapsulated = true
									bdn.Encap(`"`)
								}
							}

							D := unquote(value)
							if !hasPfx(D, LocalScheme) {
								err = errorf("Illegal %s distinguishedName: [index:%d;value:%s] missing LDAP local scheme (%s)",
									key, x, D, LocalScheme)
								return
							}

							bdn.Push(DistinguishedName{newDistinguishedName(D[len(LocalScheme):], key)})
						}
					}

					var c Condition
					if oper == Ne {
						if !vEncapsulated {
							c = bdn.Ne().Encap(`"`)
						} else {
							c = bdn.Ne().Encap()
						}
					} else {
						if !vEncapsulated {
							c = bdn.Eq().Encap(`"`)
						} else {
							c = bdn.Eq().Encap()
						}
					}

					slices[len(slices)] = c.Paren(cparen).
						setID(id).
						setCategory(key.String())

				case BindUAT, BindGAT:
					// TEMPORARY
					// assemble and store new bind rule condition
					// components for eventual migration to stack
					if err = pushBR(key.String(), id, join(vals, ``), oper); err != nil {
						return
					}

				case BindToD:
					if len(vals) != 1 {
						err = errorf("Unexpected number of %s values; want %d, got %d",
							key, 1, len(vals))
						return
					}

					// extract clocktime from raw value, remove
					// quotes and any L/T WHSP
					raw := trimS(unquote(vals[0]))
					thyme := ToD(raw)
					if thyme.String() != raw {
						err = errorf("Unexpected %s clock time parsing result; want '%s', got '%s' (hint: use military time; 0000 through 2400)",
							key, raw, thyme)
						return
					}

					slices[len(slices)] = Cond(key, thyme, oper).
						Paren(cparen).
						Encap(`"`).
						setID(id).
						NoPadding(!ConditionPadding).
						setCategory(key.String())

				case BindDoW:
					if len(vals) != 1 {
						err = errorf("Unexpected number of %s values; want %d, got %d",
							key, 1, len(vals))
						return
					}

					// extract auth method from raw value, remove
					// quotes and any L/T WHSP and analyze
					raw := trimS(unquote(vals[0]))
					var dw DayOfWeek
					if dw, err = parseDoW(raw); err != nil {
						return
					}

					slices[len(slices)] = Cond(key, dw, oper).
						Paren(cparen).
						Encap(`"`).
						setID(id).
						NoPadding(!ConditionPadding).
						setCategory(key.String())

				case BindAM:
					if len(vals) != 1 {
						err = errorf("Unexpected number of %s values; want %d, got %d",
							key, 1, len(vals))
						return
					}

					// extract auth method from raw value, remove
					// quotes and any L/T WHSP and analyze
					raw := trimS(unquote(vals[0]))
					am := matchAuthMethod(raw)
					if !eq(am.String(), raw) {
						err = errorf("Unexpected %s auth method parsing result; want '%s', got '%s' (hint: choose one (1) of `none`, `simple`, `ssl` and `sasl`)",
							key, raw, am)
						return
					}

					slices[len(slices)] = Cond(key, am, oper).
						Paren(cparen).
						Encap(`"`).
						setID(id).
						NoPadding(!ConditionPadding).
						setCategory(key.String())

				case BindSSF:
					if len(vals) != 1 {
						err = errorf("Unexpected number of %s values; want %d, got %d",
							key, 1, len(vals))
						return
					}

					// extract factor from raw value, remove
					// quotes and any L/T WHSP
					raw := trimS(unquote(vals[0]))
					fac := SSF(raw)
					if fac.String() != raw {
						err = errorf("Unexpected security strength factor parsing result; want '%s', got '%s' (hint: valid range is 0-256)",
							raw, fac)
						return
					}

					slices[len(slices)] = Cond(key, fac, oper).
						Paren(cparen).
						Encap(`"`).
						setID(id).
						NoPadding(!ConditionPadding).
						setCategory(key.String())

				case BindIP, BindDNS:
					if len(vals) != 1 {
						err = errorf("Unexpected number of %s values; want %d, got %d",
							key, 1, len(vals))
						return
					}

					if key == BindIP {
						// extract IP Address(es) from raw value,
						// remove quotes and any L/T WHSP and then
						// split for iteration.
						raw := split(trimS(unquote(vals[0])), `,`)
						var addr IPAddr
						for ipa := 0; ipa < len(raw); ipa++ {
							addr.Set(raw[ipa])
						}

						if len(raw) != addr.Len() {
							err = errorf("Unexpected number of IP addresses parsed; want '%d', got '%d'",
								len(raw), addr.Len())
							return
						}

						slices[len(slices)] = Cond(key, addr, oper).
							Paren(cparen).
							Encap(`"`).
							setID(id).
							NoPadding(!ConditionPadding).
							setCategory(key.String())

					} else {
						// extract FQDN from raw value, remove
						// quotes and any L/T WHSP.
						raw := trimS(unquote(vals[0]))
						fq := DNS(raw)

						if fq.String() != raw {
							err = errorf("Unexpected FQDN parsing result; want '%s', got '%s'", raw, fq)
							return
						}

						slices[len(slices)] = Cond(key, fq, oper).
							Paren(cparen).
							Encap(`"`).
							setID(id).
							NoPadding(!ConditionPadding).
							setCategory(key.String())
					}
				}

				// ###################################################################

				// We're done; reset for any subsequent
				// conditions that might be present.
				cready = false           // falsify bind rule condition readiness flag
				cparen = false           // falsify bind rule condition parenthetical flag
				kw = ``                  // clear bind rule condition keyword
				vals = make([]string, 0) // clear bind rule expression(s)
			}
		}

		if chop == -1 {
			break
		}

		// If we have more than one (1) token remaining
		// to process in the next loop(s), make a note
		// of the upcoming token, and then truncate the
		// token slices to remove the current token as
		// we're done with it.
		if len(tokens) > 1 {
			// make a note of the current token for the
			// next iteration's "look-behind" and "seen"
			// capabilities before exiting this loop.
			last = token
			seen = append(seen, last)

			// Truncate the token slices remaining to
			// remove the token we've already handled
			// (and archived as 'seen').
			tokens = tokens[1:]
		} else {
			// No tokens left to process. Break out
			// of this for-loop so we don't spin our
			// wheels forever.
			break
		}
	}

	// With our orderable sequence of processed
	// Condition and Rule (stack) instances, we
	// will now cycle through each map value by
	// the corresponding enum key, and will add
	// (push) said values into the appropriate
	// hierarchical structure in order to keep
	// the sequence the same as the manner in
	// which the user originally entered it.
	//
	// Individual conditions that are contiguous
	// between the same Boolean WORD operator(s)
	// are all funneled into a single stack that
	// is pushed into the return stack.
	//
	// Individual stacks (Rules) are enveloped in
	// an outer stack, which is then added to the
	// return stack.
	if len(slices) > 0 {

		// Initialize our transfer map
		R := make(map[int]Rule, 0)

		// Set outer parenthesis, if applicable
		outer.Paren(oparen)

		var prev string
		for i := 0; i < len(slices); i++ {
			// If we've progressed at least one (1)
			// slice, we'll retain knowledge of the
			// previous slice type, expressed using
			// a single capital letter.
			if i > 0 {
				if _, ok := slices[i-1].(Condition); ok {
					prev = `C` // Previous slice was a Condition
				} else if _, ok = slices[i-1].(Rule); ok {
					prev = `R` // Previous slice was a Rule (stack)
				}
			}

			switch tv := slices[i].(type) {

			// Slice value is a single Condition
			case Condition:
				// Bail out if at any point a Condition
				// instance is bogus, or not well-formed.
				if err = tv.Valid(); err != nil {
					return
				}

				// We need to initialize a new stack IF
				// we've just started, OR if the previous
				// slice was a Rule (and not a Condition)
				if len(R) == 0 || prev == `R` {
					R[len(R)] = ruleByLoP(outer.Category()).
						Paren(hasSfx(tv.ID(), `_bind`) && tv.ID() != `parenthetical_bind`)
				}

				// Push the current condition instance
				// into the most recent stack found
				// within our temporary map.
				R[len(R)-1].Push(tv)
				//if tv.ID() == `parenthetical_bind` {
				//	printf("NNOW[%s] %s\n", tv.ID(), R[len(R)-1])
				//	R[len(R)-1].Paren(true)
				//}

				// Set "C" (for Condition) as the last-seen
				// marker value.
				prev = `C`

				// Slice value is a single stack (Rule)
			case Rule:
				// Bail out if at any point a stack (Rule)
				// instance is encountered that contains no
				// elements (is empty). This would SEEM to
				// indicate an empty pair of parentheticals
				// resides within the bind rule expression.
				if tv.Len() == 0 {
					err = errorf("Empty parenthetical expression found '()'; aborting")
					return
				}

				// Envelope the current stack within an
				// appropriate Boolean WORD-based stack
				// and add to our transfer map.
				R[len(R)] = ruleByLoP(tv.Category()).Push(tv).setID(tv.ID())

				// Set "R" (for Rule) as the last-seen
				// marker value.
				prev = `R`
			}
		}

		if len(R) == 0 {
			// Empty bind rule is a fatal error.
			err = errorf("Empty bind rule expression found; aborting")
			return
		} else if len(R) == 1 {
			if R[0].Len() == 1 {
				// If there is only one (1) element
				// in the Rule map, use the return
				// as the variable, as opposed to
				// pushing into it.
				Z, _ := R[0].Index(0)
				if assert, aok := Z.(Rule); aok && assert.Len() > 0 {
					outer = assert
				} else if assert2, bok := Z.(Condition); bok {
					outer.Push(assert2)
				}
			} else if R[0].Len() != 0 {
				outer = R[0]
			} else {
				err = errorf("Empty parenthetical expression found '()'; aborting")
				return
			}
			return
		}

		// With multiple bind rule expression elements
		// assembled and enveloped as needed, push each
		// piece (in the original order) into our return
		// stack.
		for i := 0; i < len(R); i++ {
			xxx := R[i].ID()
			printf("[%s] %s %s\n", outer.ID(), xxx, R[i])
			outer.Push(R[i].
				setID(outer.ID()).
				setCategory(outer.Category()))
		}
	}

	return
}

/*
parsePerm reads and processes token slices into an instance of Permission, which is
the first (1st) component of a PermissionBindRule instance. It, alongside an error
and chop index, are returned when processing stops or completes.
*/
func parsePerm(tokens []string) (chop int, perm Permission, err error) {
	var disp string
	var privs []any
	var done bool

	for _, token := range tokens {
		// closing paren during perm
		// mode means perm has ended
		// and at least one (1) bind
		// rule is beginning.
		switch lc(token) {
		case `allow`, `deny`:
			disp = lc(token)
		case `;`, `(`, `,`:
			// do nothing
		case `)`:
			done = true
		default:
			privs = append(privs, lc(token))
		}

		chop++
		if done {
			break
		}
	}

	// assemble permission
	if disp == `allow` {
		if len(privs) == 0 {
			perm = Allow(`none`)
		} else {
			perm = Allow(privs...)
		}
	} else {
		if len(privs) == 0 {
			perm = Deny(`all`, `proxy`)
		} else {
			perm = Deny(privs...)
		}
	}

	return
}

/*
parsePBR reads and processes a sequence of tokens into one (1) Permission and one (1)
bind rule. An error and a chop index is returned alongside these components.
*/
func parsePBR(tokens []string) (chop int, pbr []PermissionBindRule, err error) {
	var mode string = `permission` // starting mode is always permission

	pbr = make([]PermissionBindRule, 0)

	for _, token := range tokens {
		if len(tokens) <= 1 {
			return
		}

		switch token {

		case `allow`, `deny`:
			switch mode {

			case `permission`:
				var skipTo int
				var br Rule
				var perm Permission
				if skipTo, perm, err = parsePerm(tokens); err != nil {
					return
				}

				tokens = tokens[skipTo:]
				chop = skipTo
				if skipTo, br, err = parseBR(tokens, 0); err != nil {
					return
				}
				pbr = append(pbr, PB(perm, br))

				// Done processing!
				if skipTo-1 == len(tokens) {
					chop = -1
					return
				}

				tokens = tokens[skipTo:]
				chop += skipTo
				mode = `bind`
			}

		case `;`:
			chop++
			if mode == `bind` {
				mode = `permission`
				continue
			}

		case `,`:
			chop++
			if mode == `permission` {
				// don't preserve permission
				// delimiters.
				continue
			}
		}
	}

	if len(pbr) == 0 {
		err = errorf("No Permission Bind Rule(s) found; aborting")
	}

	return
}

/*
parseInstruction returns a populated instance of Instruction, alongside
an error instance. This is the top-level parser function, which handles
all lower-levels of value recognition.

The input argument, expr, is the string-based ACIv3 expression in its
complete form.
*/
func parseInstruction(expr string) (a Instruction, err error) {
	// if expr is zero-length, absolutely nothing
	// can be done and is considered a user error.
	if len(expr) == 0 {
		err = errorf("Cannot process zero-length instruction")
		return
	}

	// Initialize our new Instruction instance. This is the
	// return value defined in the signature, and is that
	// which the user expects in exchange for their text
	// expression.
	a = ACI()

	// Remove unneeded contiguous WHSP (tab/space) as
	// well as newlines. Also remove any leading or
	// trailing WHSP, contiguous or not.
	expr = condenseWHSP(expr)

	// Tokenize our aci string input
	tokens := aciparser.InstructionToTokens(expr)

	// Always start with target mode, as
	// targetRules will be the first items
	// encountered (if defined).
	var mode string = `target`

	// Keep track of where we'll be next.
	var next string

	// Keep track of the so-called "chop index",
	// which is used following recursion-based
	// processing phases to avoid superfluous
	// handling of already-seen tokens.
	var skipTo int

	// Iterate our tokenized ANTLR char stream
	// and handle each token accordingly.
	for index, token := range tokens {

		// If recursion was performed, we MAY need to skip
		// ahead to avoid processing tokens already handled.
		// This only happens IF the chop index is non-zero
		// AND is higher than the current index AND does not
		// exceed the current length of the token slices.
		if skipTo != 0 && skipTo > index && skipTo < len(tokens) {
			continue
		}

		// If we have a ways to go, store the next (upcoming)
		// token so we can "look ahead" if needed.
		if len(tokens) > 2 {
			// before we handle the current token, make
			// a note of the next token, as we may need
			// to evaluate it along the way.
			next = tokens[1]
		} else {
			// Remaining tokens are too few, so we bail.
			break
		}

		// Perform a value switch to analyze the current token
		// and see if it conforms to the components we expect
		// to find.
		switch {

		// If we found the anchor, that means the expression is
		// "targetless", meaning no target rules were specified.
		case token == `version 3.0; acl`:
			mode = `acl`

		// If we found a semicolon, we know the current mode is
		// ending (OR the current value of the current mode has
		// been handled).
		case token == `;`:

			// Perform a mode switch so we can take appropriate
			// action based on the current stage of processing.
			switch mode {

			// A semicolon while acl mode is in effect
			// means the acl is about to end, and the
			// PermissionBindRule phase is next. Thus
			// the tokens slices is FIFO trimmed.
			case `acl`:

				// Recurse into PermissionBindRule processing
				// phase. Note that an ACI will always have one
				// (1) OR MORE of these.
				var pbr []PermissionBindRule
				if skipTo, pbr, err = parsePBR(tokens[1:]); err != nil {
					return
				}

				// Add the resultant permission + bind rule(s)
				// instance to our return Instruction.
				for p := 0; p < len(pbr); p++ {
					a.Set(pbr[p])
				}

				// If skipTo is minus one, this indicates we
				// finished processing the PBR section of the
				// ACI.
				if skipTo == -1 {
					return
				}

				// Truncate the token slices to begin where we
				// just left off, thus avoiding superfluous
				// processing of already-seen tokens.
				tokens = tokens[skipTo:]
			}

		// We we found an opening parenthesis, we know that we're either
		// in the target rule processing phase, OR just finishing said
		// phase and are about to move onto the ACI "anchor".
		case token == `(`:
			if next == `version 3.0; acl` {
				mode = `acl`
				continue
			}

			// perform a mode switch so we can take appropriate
			// action based on the current stage of processing.
			switch mode {

			// an opening parenthesis while target mode is in
			// effect means that we're about to receive one (1)
			// or more target rule conditions.
			case `target`:

				// prepare our targetRules Rule stack, into
				// which one (1) or more target rule condition
				// instances shall be pushed.
				var targetRules Rule

				// recurse into parseTR, extract the target
				// rule expressions and obtain our chop index.
				if skipTo, targetRules, err = parseTR(tokens[1:]); err != nil {
					return
				}

				// We are done processing target rules. We
				// know for certain the next mode is 'acl',
				// so set it now.
				mode = `acl`

				// Truncate the token slices to begin where
				// we just left off.
				tokens = tokens[skipTo:]

				// Set our new targetRules Rule instance (IF
				// non-zero) to the ACI instance (a).
				if targetRules.Len() > 0 {
					a.Set(targetRules)
				}
			}

		// default is a catch-all for any token not explicitly handled
		// in the above case statements.
		default:

			// If the mode is acl, we're expecting a double-quoted
			// access control label (hence the acronym). Strip the
			// quotes off, as go-stackage handles encapsulation
			// without the need for preserving characters literally,
			// and add the naked string value.
			if mode == `acl` {
				a.Set(unquote(token))
			}
		}
	}

	// We're done, return the Instruction along with
	// (what is likely) a nil error.
	return
}

/*
assertParsedConditionExpression will read the provided keyword and expression
value (which should be a double-quoted non-zero value) and determine the correct
destination type instance for storage.
*/
//func assertParsedConditionExpression(kw, expr string) (any, error) {
//}

/*
func parseFilter(filter string) (r Rule, err error) {
	if len(filter) < 3 {
		err = errFilterTooSmall
		return
	}

	// tokenize the filter string value
	tokens := ldapparser.FilterToTokens(filter)
	r = Rule(stackageList())
	_, err = marshalFilter(tokens,r)
	return
}
*/

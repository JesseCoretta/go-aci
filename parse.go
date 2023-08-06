package aci

/*
parse.go acts as a bridge between the go-aci package and the ANTLR4
parser/lexer subsystems for various ACIv3 and LDAP components.
*/

import (
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

	// Initialize our target rule stack; max capacity is nine (9)
	targetRules = T()

	var (
		kw, // target rule keyword
		cop, // target rule comparison operator
		last, // previous token
		next string // upcoming token

		// vals stores a sequence of value (and value RELATED)
		// tokens when detected. This var is purged of its
		// contents whenever the end of the value sequence is
		// reached.
		vals []string = make([]string, 0)

		done,
		cready bool // condition ready for assembly
	)

	// iterate tokens, looking for target rule elements.
	for index, token := range tokens {
		// get previous and upcoming tokens, if possible.
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
			cready = len(kw) > 0

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
				if err = targetExprValueInvalid(token, last, next, index); err != nil {
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
				if targetRuleNotDone(next) {
					break
				}

				var (
					// Prepare our condition for target rule creation
					c Condition
				)

				// This is the last (or only!) value component. We can
				// now analyze the keyword and the value(s) to ascertain
				// the appropriate instance type for condition storage
				// (and to perform other context-specific sanity checks).

				if c, err = assertTargetRule(vals, kw, cop); err != nil {
					return
				}
				targetRules.Push(c)

				// Reset for next target rule condition, if any
				kw = ``
				cop = ``
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

func targetExprValueInvalid(token, last, next string, index int) (err error) {
	if !isQuoted(token) && (token != `||` && token != `&&`) {
		err = errorf("Bogus Target Rule condition expression between '%s' [%d] and '%s' [%d]; value must be a non-zero string enclosed within double quotes, or a symbolic list (||,&&) of same",
			last, index-1, next, index+1)
	}
	return
}

func targetExprReady(kw, op string) bool {
	return len(kw) > 0 && len(op) > 0
}

func targetRuleNotDone(next string) bool {
	return next == `||` || next == `&&` || isQuoted(next)
}

func assertTargetRule(vals []string, kw, op string) (c Condition, err error) {
	// Begin with an assertion switch upon the target keyword
	// (which we already vetted as sane) ...
	switch key := matchTKW(kw); key {

	case TargetScope, TargetFilter:
		if len(vals) != 1 {
			err = errorf("Unexpected number of %s values; want %d, got %d",
				key, 1, len(vals))
			return
		}

		if key == TargetScope {
			c, err = assertTargetScope(unquote(vals[0]), op)
		} else {
			f := unquote(vals[0])
			c = TFilter().Push(f).Eq()
		}

	case TargetAttr:
		c, err = assertTargetAttributes(vals, op)

	case TargetCtrl, TargetExtOp:
		c, err = assertTargetOID(vals, op, key)

	case Target, TargetTo, TargetFrom:
		c, err = assertTargetDN(vals, op, key)

	case TargetAttrFilters:
		// TODO
		//if len(vals) != 1 {
		//      err = errorf("Target Rule keyword '%s' supports single values only, but %d values were found: %v",
		//              kw,len(vals),vals)
		//      return
		//}

	default:
		err = errorf("Unhandled target rule type '%s'", key)
	}

	return
}

func assertTargetScope(value string, op string) (c Condition, err error) {
	if len(value) == 0 {
		err = errorf("Zero-length LDAP Search Scope detected; aborting")
		return
	}

	scn := unquote(value)
	sc := strToScope(scn)

	// base is a fallback for a bogus scope, so
	// if the user did not originally request
	// base, we know they requested something
	// totally unsupported.
	if sc == noScope {
		err = errorf("Bogus %s value: '%s'", TargetScope, scn)
		return
	}

	c, err = conditionByOperator(op, sc)

	return
}

func assertTargetOID(vals []string, op string, key TargetKeyword) (c Condition, err error) {
	var vencap bool
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
						vencap = true
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
					vencap = true
					toid.Encap(`"`)
				}
			}

			// Type-A confirmed
			value = trimS(unquote(value))
			o, _ := newObjectID(key, value)
			toid.Push(ObjectIdentifier{o})
		}
	}

	c, err = conditionByOperator(op, toid)
	if !vencap {
		c.Encap(`"`)
		return
	}
	c.Encap()

	return
}

func assertTargetDN(vals []string, op string, key TargetKeyword) (c Condition, err error) {
	var vencap bool
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
						vencap = true
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
					vencap = true
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

	c, err = conditionByOperator(op, tdnr)
	if !vencap {
		c.Encap(`"`)
		return
	}
	c.Encap()

	return
}

func assertTargetAttributes(vals []string, op string) (c Condition, err error) {
	var vencap bool
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
						vencap = true
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
					vencap = true
					tat.Encap(`"`)
				}
			}
			tat.Push(ATName(trimS(unquote(value))))
		}
	}

	c, err = conditionByOperator(op, tat)
	if !vencap {
		c.Encap(`"`)
		return
	}
	c.Encap()

	return
}

/*
parseBR processes the input slice values (tokens) into a Bind Rule stack (outer).
It, alongside an error and chop index, are returned when processing stops or completes.
*/
func parseBR(tokens []string, depth, pspan int) (chop int, outer Rule, err error) {
	// Don't bother processing tokens that could
	// never possibly represent a valid Bind Rule
	// expression (UNLESS we're recursing within
	// a parenthetical bind rule expression).
	if len(tokens) < 4 && pspan == 0 {
		err = errorf("Empty bind rule input value, or value is below minimum possible length for validity: %v [%d<4]", tokens, len(tokens))
		return
	}

	// keep track of how many times we recurse
	// into this function.
	depth++

	// Create temporary storage vars for some of
	// the Condition components that will need
	// to be preserved across loops.
	var (
		kw, // Bind Rule Condition keyword
		cop, // Bind Rule Condition comparison operator
		last, // previous token
		next string // upcoming token

		vals []string = make([]string, 0) // Bind Rule Condition expression(s)
		seen []string                     // tokens already processed

		cready, // marker for condition assembly readiness
		cparen bool // parenthetical marker for condition instances

		ct     int = -1 // running total
		skipTo int      // skip-ahead per recursion return

		slices map[int]any = make(map[int]any, 0)

		// convenient true/false bind rule keyword
		// recognizer func.
		isKW func(string) bool = func(o string) bool {
			return matchBKW(lc(o)) != BindKeyword(0x0)
		}

		// convenient comparison operator token
		// recognizer func.
		isTokOp func(string) bool = func(o string) (ok bool) {
			_, ok = matchOp(o)
			return
		}
	)

	_, isb := isBS(tokens)
	_, opai, _, _ := parenState(join(tokens, ``))
	opw, _ := hasOp(tokens)
	//if !opf {
	//	singc = true
	//}
	outer = ruleByLoP(opw)

	// iterate each of the tokens.
	for _, token := range tokens {
		// actual iteration counter. We're unable to
		// rely on range index because the iterable
		// var (tokens) is continuously truncated
		// along the way.
		ct++

		if len(tokens) > 2 {
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

				// evaluate condition parentheticals
				cparen = isKW(tokens[1]) && tokens[4] == `)`

				if next == `(` {
					// Launch a new inner recursion of this
					// same function.
					var inner Rule
					_, _, oip, _ := parenState(join(tokens, ``))
					//printf("QQQQ Recursing[D:%d;T:%d;P:%d;O:%t;B:%t] %v\n", depth, tat, pai, oip, bal, tokens)
					if skipTo, inner, err = parseBR(tokens[1:], depth, pspan); err != nil {
						return
					}

					// Done processing!
					if skipTo-1 == len(tokens) {
						chop = skipTo - 2
						tokens = tokens[len(tokens)-2:]
					} else {
						tokens = tokens[skipTo:] // truncate tokens already processed through recursion
					}

					if inner.Len() > 0 {
						_, _, oip, _ = parenState(join(tokens, ``))
						inner.Paren(isb || oip) // this influences rule #20
						//printf("QQQQ Recursion returned [L:%d;D:%d;T:%d;P:%d;O:%t;B:%t] %s\n", inner.Len(), depth, tat, pai, oip, bal, inner)
						slices[len(slices)] = inner // save stack
					}
				}
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
						offset += 2 // +2 because we're in "look ahead" mode (and 'not' is 'and not').
					}

					// if next is another parenthetical opener,
					// make a note of it.
					iparen := tokens[2] == `(`
					printf("IPAREN:%t [tokens:%v]\n", iparen, tokens[2:])

					// Launch a new inner recursion of this
					// same function.
					var inner Rule
					if skipTo, inner, err = parseBR(tokens[offset:], depth, pspan); err != nil {
						return
					}

					tokens = tokens[skipTo:] // truncate tokens already processed through recursion
					chop += skipTo           // sum our "skip to" index with our return chop index

					// If the inner stack has at least one
					// (1) element, preserve it for the end
					// stack element, else take no action.
					if inner.Len() > 0 {
						//inner.Paren(oip)	    // not needed?
						inner.Paren(iparen)         // not needed?
						inner.setCategory(ttoken)   // mark the inner stack's logical Boolean WORD operator
						slices[len(slices)] = inner // save stack
						printf("INNER IS: %s\n", inner)
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
				_, _, oip, _ := parenState(join(tokens[offset:], ``))
				if skipTo, inner, err = parseBR(tokens[offset:], depth, pspan); err != nil {
					return
				}

				tokens = tokens[skipTo:] // truncate tokens already processed through recursion
				chop += skipTo           // sum our "skip to" index with our return chop index

				// If the inner stack has at least one
				// (1) element, preserve it for the end
				// stack element, else take no action.
				if inner.Len() > 0 {
					inner.Paren(oip)
					inner.setCategory(ttoken)   // mark the inner stack's logical Boolean WORD operator
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
				// ensure the sequence of values is well-formed
				// and contains a pattern we're expecting ...
				var brake bool
				if brake, err = checkBindRuleConditionValueStream(last, token, next, ct); err != nil {
					return
				}

				// increment chop index by one (1)
				chop++

				// Save this value; we don't yet know if this
				// value is merely one (1) of multiple values
				// as opposed to a single value alone.
				vals = append(vals, token)

				// If we encountered a DELIMITER, break out of
				// this case to continue at next token.
				if brake {
					break
				}

				// This is the last (or only!) value component. We can
				// now analyze the keyword and the value(s) to ascertain
				// the appropriate instance type for condition storage
				// (and to perform other context-specific sanity checks).
				var c Condition
				if c, err = parseBindRuleCondition(vals, kw, cop); err != nil {
					return
				}

				// Save the condition for handling later ...
				slices[len(slices)] = c.Paren(cparen).
					NoPadding(!ConditionPadding).
					setID(bindRuleID)

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
					R[len(R)] = ruleByLoP(outer.Category())
				}

				// Push the current condition instance
				// into the most recent stack found
				// within our temporary map.
				R[len(R)-1].Push(tv)

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

				R[len(R)] = ruleByLoP(tv.Category()).
					Paren(tv.isParen() && pspan == 0)
				R[len(R)-1].Push(tv)
				printf("RECEIVED: %s\n", tv)

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
					printf("Assigning (%T) %s to outer\n", assert, assert)
					outer = assert

				} else if assert2, bok := Z.(Condition); bok && !assert2.IsZero() {
					printf("Assigning (%T) %s to outer\n", assert2, assert2)
					outer.Push(assert2)
				}

			} else if R[0].Len() == 0 {
				err = errorf("Empty parenthetical expression found '()'; aborting")

			} else {
				outer = R[0]
			}

			return
		}

		if (isBPC(tokens) && opai == 0) || (isb && opai > 0) {
			//outer.Paren(depth==0 && pspan==0)	// influences #41
			printf("tokens: %#v\n", R)
			outer.Paren(depth == 0 && pspan == 0)
		}

		// With multiple bind rule expression elements
		// assembled and enveloped as needed, push each
		// piece (in the original order) into our return
		// stack.
		for i := 0; i < len(R); i++ {
			ident := R[i].ID()
			tot, pairs, oip, bal := parenState(R[i].String())

			printf("[%s|%d] [D:%d;T:%d;P:%d;O:%t;B:%t] %s\n",
				ident,
				R[i].Len(),
				depth,
				tot,
				pairs,
				oip,
				bal,
				R[i])

			outer.Push(R[i].Paren(depth > 0 && isb))
			//setCategory(outer.Category()))
		}
	}

	return
}

func assertBindRuleUGRDN(vals []string, key BindKeyword, op string) (c Condition, err error) {
	if len(vals) == 0 {
		err = errorf("Empty bind rule value")
		return
	}

	var vencap bool
	var value string = vals[0]
	if hasPfx(value, LocalScheme) && contains(vals[0], `?`) {
		var uri LDAPURI

		if uri, err = parseLDAPURI(value, key); err != nil {
			return
		}

		c, err = conditionByOperator(op, uri)
		return
	}

	// prepare a stack for our DN value(s)
	bdn := ruleByDNKeyword(key)

	// bind rule is either or both of the following:
	// A: one (1) double-quoted DN
	// B: one (1) double-quoted LIST of unquoted DNs in symbolic OR context
	for x := 0; x < len(vals); x++ {
		value = vals[x]
		if contains(value, `||`) {

			// Type-B confirmed
			for ix, O := range split(unquote(value), `||`) {
				if len(O) == 0 {
					continue
				}

				if x == 0 && ix == 0 {
					if !isQuoted(vals[x]) && isQuoted(O) {
						vencap = true
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
					vencap = true
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

	c, err = conditionByOperator(op, bdn)
	if !vencap {
		c.Encap(`"`)
		return
	}
	c.Encap()

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
	perm = assemblePermissionByDisposition(disp, privs)

	return
}

/*
assemblePermissionByDisposition shifts the return value to include all privilege
abstraction found within privs, and will initialize said return value based on
the disposition (allow or deny) selected by the user.
*/
func assemblePermissionByDisposition(disp string, privs []any) (perm Permission) {
	if disp == `allow` {
		if len(privs) == 0 {
			perm = Allow(`none`)
		} else {
			perm = Allow(privs...)
		}
		return
	}

	if len(privs) == 0 {
		perm = Deny(`all`, `proxy`)
	} else {
		perm = Deny(privs...)
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
				if skipTo, br, err = parseBR(tokens, -1, 0); err != nil {
					return
				}
				pbr = append(pbr, PB(perm, br))
				printf("Appended %s to PBR\n", br)

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
		if len(tokens) <= 2 {
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

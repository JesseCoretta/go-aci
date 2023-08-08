package aci

func parseBindRule(tokens []string, depth, pspan int, word ...string) (outer Rule, skip int, err error) {
        // Don't bother processing tokens that could
        // never possibly represent a valid Bind Rule
        // expression (UNLESS we're recursing within
        // a parenthetical bind rule expression).
        if len(tokens) < 4 && depth==0 {
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
                kw,		// Bind Rule Condition keyword
                cop string	// Bind Rule Condition comparison operator

                vals []string // Bind Rule Condition expression(s)

		cc Rule = Rule(stackageBasic())
                cready, // marker for condition assembly readiness
		iparen,
		oparen,
                cparen bool // parenthetical marker for condition instances

                ct     int = -1 // running total

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

	if len(word) > 0 {
		outer = ruleByLoP(word[0])
	} else {
		opw, _ := hasOp(tokens)
		outer = ruleByLoP(opw)
	}

	oparen = ( tokens[0] == `(` &&
		( tokens[len(tokens)-1] == `)` &&
		tokens[len(tokens)-2] != `;` ) )

	var cct string = outer.Category()

        // iterate each of the tokens.
        for _, token := range tokens {
                // actual iteration counter. We're unable to
                // rely on range index because the iterable
                // var (tokens) is continuously truncated
                // along the way.
                ct++

		var inner Rule
		var done bool

		switch {
		case cready:
			var stop int
			stop, vals = readQuotedValues(tokens)
                        var c Condition
                        if c, err = parseBindRuleCondition(vals, kw, cop); err != nil {
                                return
                        }
			tokens = tokens[stop+1:]
			c.Paren(cparen)
			printf("Pushing %s (cct:%s)\n", c, cct)
			//outer.Push(c)
			cc.Push(c)

			vals = []string{}
			cparen = false
			cready = false
			kw = ``
			cop = ``

		case token == `;`:
			if len(tokens) <= 2 {
				// nothing else to process
				skip = -1
			} else {
				// another perm/bind rule pair may be next
				skip++
			}
			done = true

		case isKW(token):
			kw = token
			tokens = tokens[1:]
			continue

		case isTokOp(token):
			cop = token
			tokens = tokens[1:]
			cready = len(kw) >0 && len(cop) > 0
			continue

		case isQuoted(token):
			var stop int
			if cready, stop, vals, err = getQuotedValues(kw,cop,tokens); err != nil {
				return
			}

			tokens = tokens[stop:]
			continue

		// Boolean word operator ("AND", "OR", "AND NOT")
		case isWordOp(token):
			// this boolean operator merely continues the
			// expression, and does not signify a switch
			// to another operator, e.g.: AND -> OR, which
			// would result in a new recursion.
			if eq(token, inner.Category()) {
				tokens = tokens[1:]
				continue
			}

			// boolean operator differs from the current
			// (outer) operator. Begin new recursion, and
			// pass the desired word to this same function
			// which will result in an alloc for a new stack.
			//var innot Rule
			if eq(token, `and not`) {
				// negated condition/stack
				var innot Rule
				innot, skip, err = parseBindRule(tokens[1:], depth, pspan) //; innot.Len() > 0 {
				innot.setCategory(`or`)
				inner = ruleByLoP(token[4:]).Push(innot)
				//inner = ruleByLoP(token[4:]).setCategory(`or`).Push(unwrapRule(innot))
				//inner.Push(unwrapRule(innot).setCategory(`or`)).setCategory(token[4:])

			} else {
				// AND or OR condition/stack
				inner, skip, err = parseBindRule(tokens[1:], depth, pspan, token) //; innot.Len() > 0 {
				inner.setCategory(lc(token)) //outer.Category())
			}

		// Found a closing parenthetical
		case token == `)`:

			tokens = tokens[1:]
			pspan--
			if pspan < 0 {
				err = errorf("Unbalanced parenthetical; want 0, got %d (hint: missing an opener?)",pspan)
			} else if len(tokens) <= 2 {
				skip = -1
				iparen = false
				cparen = false
			} else if pspan == 0 {
				iparen = false
				cparen = false
			}

		// Found an opening parenthetical
		case token == `(`:
			//_, _, oup, _ := parenState(join(tokens, ``))
			tokens = tokens[1:]
			cparen = isBC(tokens[0:3]) && tokens[3] == `)`
			iparen = isBC(tokens[0:3]) && isWordOp(tokens[3])
			cct, _ = hasOp(tokens)
			pspan++

		// fallback == error condition
		default:
			err = errorf("[%d] Unhandled token '%s'\n", ct, token)

		}

		if err != nil {
			return
		}

                if cc.Len() > 0 {
                        grp := ruleByLoP(outer.Category()).Paren(iparen && cc.Len() > 1)
                        for j := 0; j < cc.Len(); j++ {
                                jidx, _ := cc.Index(j)
                                grp.Push(jidx)
                        }
                        outer.Push(grp)
                        cc.reset()
                        cct = ``
                }

		outer = transferToOuterStack(iparen, oparen, inner, outer)

		if done {
			break
		}

		switch skip {
		case 0:
			continue
		default:
			if skip == -1 || !( 0 <= skip && skip <= len(tokens)-1 ) {
				tokens = []string{}
				return
			}
			tokens = tokens[skip:]
			skip = 0
		}
	}

	return
}

func getQuotedValues(kw,op string, t []string) (cready bool, stop int, v []string, err error) {
	if stop, v = readQuotedValues(t); len(v) == 0 {
	        err = errorf("No values parsed from token stream '%v'", t)
	        return
	}

	t = t[stop:]
	cready = len(kw) > 0 && len(op) > 0

	return
}

func transferToOuterStack(iparen, oparen bool, inner, outer Rule) Rule {
	//printf("inner[%d]:%s, outer[%d]:%s, [oparen:%t; iparen:%t]\n",
	//	inner.Len(), inner, outer.Len(), outer, oparen, iparen)

	if inner.Len() == 0 {
		return outer
	}

	//var last string
        for i := 0; i < inner.Len(); i++ {
                slice, _ := inner.Index(i)
                prev, _ := outer.Index(outer.Len()-1)

		// Switch on inner slice type
                switch tv := slice.(type) {

                // Current inner slice is a Condition
                case Condition:

                        // Last-added outer slice was a Rule
                        switch uv := prev.(type) {
                        case Rule:
				//last = `R`
                                uv.Push(tv)
                                //printf("[SLICE:%d] PUSH innerR[%d] into %T: '%T:%s:%s; last:%s'\n",
                                //        i, uv.Len(), tv, uv, uv.Category() , tv, uv)

                        // Last-added outer slice was a Condition
                        case Condition:
				//last = `C`
                                var envl Rule = ruleByLoP(outer.Category()).Push(tv)

                                //printf("[SLICE:%d] PUSH tv into envl [%d]: '%T:%s:%s'\n",
                                //        i, envl.Len(), envl, envl.Category(), envl)

                                outer.Push(envl)
                                //printf("[SLICE:%d] PUSH inner? [%s] into outer [%d]: '%T:%s:%s'\n",
                                //        i, envl, outer.Len(), outer, outer.Category(), outer)
                        }

                // Current inner slice is a Rule
                case Rule:

			// Inner Rule's Boolean WORD operator does not
			// match Inner slice [i] WORD operator. This
			// would indicate a new stack (Rule) is coming
			// up ...

			//printf("INNERC:%s; TVC:%s; OUTERC:%s\n",
			//	inner.Category(),
			//	tv.Category(),
			//	outer.Category())

                        if inner.Category() != tv.Category() {
                                var envl Rule = ruleByLoP(inner.Category()).Push(tv)
                                outer.Push(envl)
                                //printf("[SLICE:%d] PUSH cat [cat:%s;envl:%d;%s] into outer [%d]: '%T::%s[%s]'\n",
                                //        i, tv.Category(), envl.Len(), envl, outer.Len(), outer, outer.Category(), outer)
				break
                        }

			// Push Inner slice [i] into Outer Rule.
                        outer.Push(tv)
                        //printf("[SLICE:%d] ##PUSH## tv: '%T::%s'\n", i, tv, tv)

                }
		//printf("Last was %s\n", last)
        }

	outer.Paren(oparen)
	printOuter(outer)

	return outer
}

func printOuter(outer Rule, tabs ...string) {
	var tab string
	if len(tabs) > 0 {
		for t := 0 ; t < len(tabs)-1; t++ {
			tab += string(rune(9))
		}
	}

	if len(tab) > 0 {
		tab += string(rune(9))
	}

	//printf("\n%s :: OUTER :: [len:%d]: %T[%s] >>> \n",
	//	tab, outer.Len(), outer, outer.Category())

        for i := 0; i < outer.Len(); i++ {
                sl, _ := outer.Index(i)

                switch tv := sl.(type) {
                case Rule:
			printf("%s [%d]: %T[%s]\n", tab, i, tv, tv.Category())
			printOuter(tv, tab, string(rune(9)))

			//case Condition:
			//tab += string(rune(9))
                        //printf("%s [%d]: %T[ocat:%s]: '%s'\n",
                        //        tab, i, tv, outer.Category(), tv)
                }
        }
}

/*
parseBR processes the input slice values (tokens) into a Bind Rule stack (outer).
It, alongside an error and chop index, are returned when processing stops or completes.
*/
func parseBR(tokens []string, depth, pspan int) (chop int, outer Rule, err error) {
	return
}

//	// Don't bother processing tokens that could
//	// never possibly represent a valid Bind Rule
//	// expression (UNLESS we're recursing within
//	// a parenthetical bind rule expression).
//	if len(tokens) < 4 && pspan == 0 {
//		err = errorf("Empty bind rule input value, or value is below minimum possible length for validity: %v [%d<4]", tokens, len(tokens))
//		return
//	}
//
//	// keep track of how many times we recurse
//	// into this function.
//	depth++
//
//	// Create temporary storage vars for some of
//	// the Condition components that will need
//	// to be preserved across loops.
//	var (
//		kw, // Bind Rule Condition keyword
//		cop, // Bind Rule Condition comparison operator
//		last, // previous token
//		next string // upcoming token
//
//		vals []string = make([]string, 0) // Bind Rule Condition expression(s)
//		seen []string                     // tokens already processed
//
//		cready, // marker for condition assembly readiness
//		cparen bool // parenthetical marker for condition instances
//
//		ct     int = -1 // running total
//		skipTo int      // skip-ahead per recursion return
//
//		slices map[int]any = make(map[int]any, 0)
//
//		// convenient true/false bind rule keyword
//		// recognizer func.
//		isKW func(string) bool = func(o string) bool {
//			return matchBKW(lc(o)) != BindKeyword(0x0)
//		}
//
//		// convenient comparison operator token
//		// recognizer func.
//		isTokOp func(string) bool = func(o string) (ok bool) {
//			_, ok = matchOp(o)
//			return
//		}
//	)
//
//	_, isb := isBS(tokens)
//	_, opai, _, _ := parenState(join(tokens, ``))
//	opw, _ := hasOp(tokens)
//	//if !opf {
//	//	singc = true
//	//}
//	outer = ruleByLoP(opw)
//
//	// iterate each of the tokens.
//	for _, token := range tokens {
//		// actual iteration counter. We're unable to
//		// rely on range index because the iterable
//		// var (tokens) is continuously truncated
//		// along the way.
//		ct++
//
//		if len(tokens) > 2 {
//			// look-ahead to the next iteration's token
//			next = tokens[1]
//		}
//
//		switch {
//
//		// token is a parenthetical opener. This case switch
//		// shall launch a new parseBR (recursive) thread in
//		// which nested (superior) parentheticals are found.
//		case token == `(`:
//			chop++
//			pspan++
//
//			// Before beginning this loop, check to see if the fifth
//			// (5th) token is a Boolean WORD operator. Also check to
//			// see if there is a parenthetical Condition by itself.
//			if len(tokens) >= 5 {
//
//				// evaluate condition parentheticals
//				cparen = isKW(tokens[1]) && tokens[4] == `)`
//
//				if next == `(` {
//					// Launch a new inner recursion of this
//					// same function.
//					var inner Rule
//					_, _, oip, _ := parenState(join(tokens, ``))
//					//printf("QQQQ Recursing[D:%d;T:%d;P:%d;O:%t;B:%t] %v\n", depth, tat, pai, oip, bal, tokens)
//					if skipTo, inner, err = parseBR(tokens[1:], depth, pspan); err != nil {
//						return
//					}
//
//					// Done processing!
//					if skipTo-1 == len(tokens) {
//						chop = skipTo - 2
//						tokens = tokens[len(tokens)-2:]
//					} else {
//						tokens = tokens[skipTo:] // truncate tokens already processed through recursion
//					}
//
//					if inner.Len() > 0 {
//						_, _, oip, _ = parenState(join(tokens, ``))
//						inner.Paren(isb || oip) // this influences rule #20
//						//printf("QQQQ Recursion returned [L:%d;D:%d;T:%d;P:%d;O:%t;B:%t] %s\n", inner.Len(), depth, tat, pai, oip, bal, inner)
//						slices[len(slices)] = inner // save stack
//					}
//				}
//			}
//
//		// token is a parenthetical closer. Match only if we're
//		// NOT in the middle of Condition assembly.
//		case token == `)` && !cready:
//			chop++
//
//			// perform some checks to ensure the number
//			// of parentheticals ( '(' and ')' ) are equal
//			// (i.e.: one closer (R) for every opener (L)).
//			//
//			// The ONE AND ONLY exception is the check that
//			// discounts the semicolon terminator as the last
//			// (previous) token, as that indicates the closure
//			// of the anchor (pre-bind rule) and the end of
//			// the ACI as a whole. While this function doesn't
//			// see those tokens, we already knew to expect it,
//			// thus we can handle that condition here.
//			if pspan < 0 {
//				err = errorf("Unbalanced parenthetical expression detected near or within '%s%s%s' (token:%d, negative paren)",
//					last, token, next, ct)
//				return
//			} else if pspan == 0 && last != `;` {
//				err = errorf("Unbalanced parenthetical expression detected near or within '%s%s%s' (token:%d, zero paren)",
//					last, token, next, ct)
//				return
//			}
//
//			// pspan is known to be a positive, non-zero
//			// integer, so it is safe to decrement. But
//			pspan--
//
//			// If decrementing pspan by one (1) results
//			// in a zero value, then we have exited the
//			// parenthetical (stack) expression.
//			if pspan == 0 {
//				break
//			}
//
//			// If the NEXT token is a logical Boolean WORD
//			// operator, we take special steps because we
//			// are currently within a parenthetical bind
//			// rule expression component.
//			if isWordOp(next) {
//				chop++
//				var ttoken string = next
//				if eq(next, `and not`) {
//					// go-stackage's negation stacks use the
//					// category of 'NOT', as opposed to ACIv3's
//					// 'AND NOT' operator equivalent. Take the
//					// 'NOT' portion of the value, using its
//					// original case-folding scheme, and save
//					// it for stack tagging later.
//					ttoken = ttoken[4:]
//				}
//
//				// If the category (word operator) is not
//				// the same as the token, this means a new
//				// distinct (inner) stack is beginning (and
//				// not a continuation of outer).
//				if !eq(ttoken, outer.Category()) {
//					// We need to offset the truncation factor
//					// of our token slices when the 'AND NOT'
//					// logical Boolean WORD operator is used,
//					// as it will erroneously be interpreted
//					// as two (2) distinct tokens.
//					var offset int
//					if eq(ttoken, `not`) {
//						offset += 2 // +2 because we're in "look ahead" mode (and 'not' is 'and not').
//					}
//
//					// if next is another parenthetical opener,
//					// make a note of it.
//					iparen := tokens[2] == `(`
//					printf("IPAREN:%t [tokens:%v]\n", iparen, tokens[2:])
//
//					// Launch a new inner recursion of this
//					// same function.
//					var inner Rule
//					if skipTo, inner, err = parseBR(tokens[offset:], depth, pspan); err != nil {
//						return
//					}
//
//					tokens = tokens[skipTo:] // truncate tokens already processed through recursion
//					chop += skipTo           // sum our "skip to" index with our return chop index
//
//					// If the inner stack has at least one
//					// (1) element, preserve it for the end
//					// stack element, else take no action.
//					if inner.Len() > 0 {
//						//inner.Paren(oip)	    // not needed?
//						inner.Paren(iparen)         // not needed?
//						inner.setCategory(ttoken)   // mark the inner stack's logical Boolean WORD operator
//						slices[len(slices)] = inner // save stack
//						printf("INNER IS: %s\n", inner)
//					}
//				}
//			}
//
//		// token is a keyword
//		case isKW(token):
//			chop++
//			kw = lc(token)
//
//		// token is a comparison operator
//		case isTokOp(token):
//			chop++
//			cop = token
//			cready = len(kw) > 0 && len(cop) > 0
//
//		// token is a boolean word operator
//		case isWordOp(token) && !cready:
//			chop++
//
//			var ttoken string = token
//
//			if eq(token, `and not`) {
//				// go-stackage's negation stacks use the
//				// category of 'NOT', as opposed to ACIv3's
//				// 'AND NOT' operator equivalent. Take the
//				// 'NOT' portion of the value, using its
//				// original case-folding scheme, and save
//				// it for stack tagging later.
//				ttoken = ttoken[4:]
//			}
//
//			// If the category (word operator) is not
//			// the same as the token, this means a new
//			// distinct (inner) stack is beginning (and
//			// not a continuation of outer).
//			if !eq(ttoken, outer.Category()) {
//
//				// We need to offset the truncation factor
//				// of our token slices when the 'AND NOT'
//				// logical Boolean WORD operator is used,
//				// as it will erroneously be interpreted
//				// as two (2) distinct tokens.
//				var offset int
//				if eq(ttoken, `not`) {
//					offset++
//				}
//
//				// Launch a new inner recursion of this
//				// same function.
//				var inner Rule
//				_, _, oip, _ := parenState(join(tokens[offset:], ``))
//				if skipTo, inner, err = parseBR(tokens[offset:], depth, pspan); err != nil {
//					return
//				}
//
//				tokens = tokens[skipTo:] // truncate tokens already processed through recursion
//				chop += skipTo           // sum our "skip to" index with our return chop index
//
//				// If the inner stack has at least one
//				// (1) element, preserve it for the end
//				// stack element, else take no action.
//				if inner.Len() > 0 {
//					inner.Paren(oip)
//					inner.setCategory(ttoken)   // mark the inner stack's logical Boolean WORD operator
//					slices[len(slices)] = inner // save stack
//				}
//
//			}
//
//		// token is a semicolon, which means the end of a PermissionBindRule
//		// instance. We don't need to do anything, and we don't need to keep
//		// this token, so we match it separately. DO NOT match if we are in
//		// the middle of Condition assembly.
//		case token == `;` && !cready:
//
//		// generalized token fallback will capture quoted values as well as
//		// logical symbolic operators (||, &&) between quoted values (or as
//		// part of a quoted value).
//		default:
//			if cready {
//				// ensure the sequence of values is well-formed
//				// and contains a pattern we're expecting ...
//				var brake bool
//				if brake, err = checkBindRuleConditionValueStream(last, token, next, ct); err != nil {
//					return
//				}
//
//				// increment chop index by one (1)
//				chop++
//
//				// Save this value; we don't yet know if this
//				// value is merely one (1) of multiple values
//				// as opposed to a single value alone.
//				vals = append(vals, token)
//
//				// If we encountered a DELIMITER, break out of
//				// this case to continue at next token.
//				if brake {
//					break
//				}
//
//				// This is the last (or only!) value component. We can
//				// now analyze the keyword and the value(s) to ascertain
//				// the appropriate instance type for condition storage
//				// (and to perform other context-specific sanity checks).
//				var c Condition
//				if c, err = parseBindRuleCondition(vals, kw, cop); err != nil {
//					return
//				}
//
//				// Save the condition for handling later ...
//				slices[len(slices)] = c.Paren(cparen).
//					NoPadding(!ConditionPadding).
//					setID(bindRuleID)
//
//				// ###################################################################
//
//				// We're done; reset for any subsequent
//				// conditions that might be present.
//				cready = false           // falsify bind rule condition readiness flag
//				cparen = false           // falsify bind rule condition parenthetical flag
//				kw = ``                  // clear bind rule condition keyword
//				vals = make([]string, 0) // clear bind rule expression(s)
//			}
//		}
//
//		if chop == -1 {
//			break
//		}
//
//		// If we have more than one (1) token remaining
//		// to process in the next loop(s), make a note
//		// of the upcoming token, and then truncate the
//		// token slices to remove the current token as
//		// we're done with it.
//		if len(tokens) > 1 {
//			// make a note of the current token for the
//			// next iteration's "look-behind" and "seen"
//			// capabilities before exiting this loop.
//			last = token
//			seen = append(seen, last)
//
//			// Truncate the token slices remaining to
//			// remove the token we've already handled
//			// (and archived as 'seen').
//			tokens = tokens[1:]
//		} else {
//			// No tokens left to process. Break out
//			// of this for-loop so we don't spin our
//			// wheels forever.
//			break
//		}
//	}
//
//	// With our orderable sequence of processed
//	// Condition and Rule (stack) instances, we
//	// will now cycle through each map value by
//	// the corresponding enum key, and will add
//	// (push) said values into the appropriate
//	// hierarchical structure in order to keep
//	// the sequence the same as the manner in
//	// which the user originally entered it.
//	//
//	// Individual conditions that are contiguous
//	// between the same Boolean WORD operator(s)
//	// are all funneled into a single stack that
//	// is pushed into the return stack.
//	//
//	// Individual stacks (Rules) are enveloped in
//	// an outer stack, which is then added to the
//	// return stack.
//	if len(slices) > 0 {
//
//		// Initialize our transfer map
//		R := make(map[int]Rule, 0)
//
//		var prev string
//		for i := 0; i < len(slices); i++ {
//			// If we've progressed at least one (1)
//			// slice, we'll retain knowledge of the
//			// previous slice type, expressed using
//			// a single capital letter.
//			if i > 0 {
//				if _, ok := slices[i-1].(Condition); ok {
//					prev = `C` // Previous slice was a Condition
//				} else if _, ok = slices[i-1].(Rule); ok {
//					prev = `R` // Previous slice was a Rule (stack)
//				}
//			}
//
//			switch tv := slices[i].(type) {
//
//			// Slice value is a single Condition
//			case Condition:
//				// Bail out if at any point a Condition
//				// instance is bogus, or not well-formed.
//				if err = tv.Valid(); err != nil {
//					return
//				}
//
//				// We need to initialize a new stack IF
//				// we've just started, OR if the previous
//				// slice was a Rule (and not a Condition)
//				if len(R) == 0 || prev == `R` {
//					R[len(R)] = ruleByLoP(outer.Category())
//				}
//
//				// Push the current condition instance
//				// into the most recent stack found
//				// within our temporary map.
//				R[len(R)-1].Push(tv)
//
//				// Set "C" (for Condition) as the last-seen
//				// marker value.
//				prev = `C`
//
//			// Slice value is a single stack (Rule)
//			case Rule:
//				// Bail out if at any point a stack (Rule)
//				// instance is encountered that contains no
//				// elements (is empty). This would SEEM to
//				// indicate an empty pair of parentheticals
//				// resides within the bind rule expression.
//				if tv.Len() == 0 {
//					err = errorf("Empty parenthetical expression found '()'; aborting")
//					return
//				}
//
//				// Envelope the current stack within an
//				// appropriate Boolean WORD-based stack
//				// and add to our transfer map.
//
//				R[len(R)] = ruleByLoP(tv.Category()).
//					Paren(tv.isParen() && pspan == 0)
//				R[len(R)-1].Push(tv)
//				printf("RECEIVED: %s\n", tv)
//
//				// Set "R" (for Rule) as the last-seen
//				// marker value.
//				prev = `R`
//			}
//		}
//
//		if len(R) == 0 {
//			// Empty bind rule is a fatal error.
//			err = errorf("Empty bind rule expression found; aborting")
//			return
//
//		} else if len(R) == 1 {
//			if R[0].Len() == 1 {
//				// If there is only one (1) element
//				// in the Rule map, use the return
//				// as the variable, as opposed to
//				// pushing into it.
//				Z, _ := R[0].Index(0)
//				if assert, aok := Z.(Rule); aok && assert.Len() > 0 {
//					printf("Assigning (%T) %s to outer\n", assert, assert)
//					outer = assert
//
//				} else if assert2, bok := Z.(Condition); bok && !assert2.IsZero() {
//					printf("Assigning (%T) %s to outer\n", assert2, assert2)
//					outer.Push(assert2)
//				}
//
//			} else if R[0].Len() == 0 {
//				err = errorf("Empty parenthetical expression found '()'; aborting")
//
//			} else {
//				outer = R[0]
//			}
//
//			return
//		}
//
//		if (isBPC(tokens) && opai == 0) || (isb && opai > 0) {
//			//outer.Paren(depth==0 && pspan==0)	// influences #41
//			printf("tokens: %#v\n", R)
//			outer.Paren(depth == 0 && pspan == 0)
//		}
//
//		// With multiple bind rule expression elements
//		// assembled and enveloped as needed, push each
//		// piece (in the original order) into our return
//		// stack.
//		for i := 0; i < len(R); i++ {
//			ident := R[i].ID()
//			tot, pairs, oip, bal := parenState(R[i].String())
//
//			printf("[%s|%d] [D:%d;T:%d;P:%d;O:%t;B:%t] %s\n",
//				ident,
//				R[i].Len(),
//				depth,
//				tot,
//				pairs,
//				oip,
//				bal,
//				R[i])
//
//			outer.Push(R[i].Paren(depth > 0 && isb))
//			//setCategory(outer.Category()))
//		}
//	}
//
//	return
//}
//

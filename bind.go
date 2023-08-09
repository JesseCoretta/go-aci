package aci

/*
bind.go contains bind rule parsing functions and methods.
*/

/*
parseBindRule is the top-level parser for a sequence of bind rule expressions with,
or without, nesting and/or Boolean WORD operators. A populated outer Rule instance,
along with a token stream trimming integer and an error are returned.
*/
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
                kw,             // Bind Rule Condition keyword
                cop string      // Bind Rule Condition comparison operator

                // Bind Rule expression values are single
                // quoted values, or a sequence of values
                // (with varying quotation schemes) that
                // are delimited with a double pipe (||).
                vals []string

                // temporary contiguous condition stack,
                // mainly to preserve the parenthetical
                // nature of more than one (1) contiguous
                // condition within a single boolean WORD
                // operator's purview.
                cc Rule = Rule(stackageBasic())

                cready,         // marker for condition assembly readiness
                iparen,         // parenthetical marker for inner expressions
                oparen,         // parenthetical marker for outer expressions
                cparen bool     // parenthetical marker for condition instances

                ct     int = -1 // running total of tokens processed

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

        // During recursive calls of this function, a
        // boolean word operator token may be passed
        // as a (variadic) parameter. If and when this
        // happens, use this "word" (e.g.: 'AND', OR'
        // or 'AND NOT') as the initializer value for
        // a new (nested) outer stack.
        if len(word) > 0 {
                outer = ruleByLoP(word[0])
        } else {
                // No "word" was found, so just scan
                // ahead for the first (1st) word we
                // encounter and use that. If no word
                // was found at all, default to AND.
                opw, _ := hasOp(tokens)
                outer = ruleByLoP(opw)
        }

        // An outer parenthetical statement is qualified
        // by an opening (left) parenthesis character
        // (ASCII #40) as token zero (0), a closing (right)
        // parenthesis character (ASCII #41) as the final
        // character, and the second-to-last character NOT
        // being a PBR-termination semicolon. Note this is
        // not reliable by itself, and is sometimes used
        // in (negated) conjunction with another Boolean
        // parenthetical marker.
        oparen = ( tokens[0] == `(` &&
                ( tokens[len(tokens)-1] == `)` &&
                tokens[len(tokens)-2] != `;` ) )

        // iterate each of the tokens.
        for _, token := range tokens {
                // actual iteration counter. We're unable to
                // rely on range index because the iterable
                // var (tokens) is continuously truncated
                // along the way.
                ct++

                // inner is a loop-scoped stack used to
                // act as a temporary container for one
                // or more stacks parsed as a result of
                // a recursive (self-executing) call of
                // this same function.
                var inner Rule

                // done is a marker for when a processing
                // job should finish before all tokens are
                // processed.
                var done bool

                switch {

                // cready indicates that we're ready to
                // parse a given condition's value(s) and
                // assemble a condition. If this case is
                // matched, it means we've already acquired
                // the bind keyword and comparison operator.
                case cready:

                        var (
                                stop int
                                c Condition
                        )

                        // scan for quoted values, and stop at the
                        // first thing that is neither a quoted
                        // value nor a double pipe (||) delimiter.
                        stop, vals = readQuotedValues(tokens)

                        // assemble the condition with all needed
                        // vars now in-hand.
                        if c, err = parseBindRuleCondition(vals, kw, cop); err != nil {
                                return
                        }

                        // wherever the quoted value scanner left
                        // off, truncate our tokens to resume at
                        // that point (and not to re-process any
                        // tokens already seen).
                        tokens = tokens[stop+1:]

                        // If the condition is parenthetical itself,
                        // tell go-stackage to reflect this trait.
                        c.Paren(cparen)

                        // Save the new condition in our temporary
                        // contiguous condition stack.
                        cc.Push(c)

                        // If the sequence of contiguous condition
                        // instances is parenthetical as a whole,
                        // tell go-stackage to reflect this.
                        cc.Paren(oparen)

                        // Reset all pertinent variables for any
                        // condition tokens not yet processed.
                        vals = []string{}
                        cparen = false
                        cready = false
                        kw = ``
                        cop = ``

                // token is a terminator of a single permission/bind rule
                // "pair"; if we see this, we finish processing regardless
                // of whether any tokens remain. This function is usually
                // called by the parsePBR function, which keeps an eye on
                // the tokens from an external PoV.
                case token == `;`:
                        if len(tokens) <= 2 {
                                // nothing else to process
                                skip = -1
                        } else {
                                // another perm/bind rule pair may be next
                                skip++
                        }
                        done = true

                // token is a known bind rule keyword (e.g.: 'userdn'). If
                // this has matched, it is guaranteed that a condition has
                // just begun to undergo processing.
                case isKW(token):
                        kw = token
                        tokens = tokens[1:]
                        continue

                // token is a known comparison operator (e.g.: '>=', '=').
                // This would only appear immediately after a known Bind
                // Rule keyword was encountered and recorded.
                case isTokOp(token):
                        cop = token
                        tokens = tokens[1:]
                        cready = len(kw) >0 && len(cop) > 0
                        continue

                // token is a single quoted value. This would only appear
                // as the third (3rd) and final component in a Condition.
                case isQuoted(token):
                        var stop int
                        if cready, stop, vals, err = getQuotedValues(kw,cop,tokens); err != nil {
                                return
                        }

                        tokens = tokens[stop:]
                        continue

                // token is a Boolean WORD operator ("AND", "OR", "AND NOT").
                // This indicates multiple condition instances, which may or
                // may not be nested within other stacks.
                case isWordOp(token):

                        // this boolean operator merely continues the
                        // expression, and does not signify a switch
                        // to another operator, e.g.: AND -> OR, which
                        // would result in a new recursion.
                        if eq(token, outer.Category()) {
                                tokens = tokens[1:]
                                continue
                        }

                        // boolean operator differs from the current
                        // (outer) operator. Begin new recursion, and
                        // pass the desired word to this same function
                        // which will result in an alloc for a new stack.
                        if eq(token, `and not`) {

                                // negated condition/stack
                                var innot Rule
                                innot, skip, err = parseBindRule(tokens[1:], depth, pspan)
                                inner = ruleByLoP(token[4:])
                                if innot.Len() == 1 {
                                        innar, _ := innot.Index(0)
                                        switch tv := innar.(type) {
                                        case Rule:
                                                if tv.Len() == 1 {
                                                        inner.Push(tv.setCategory(`or`)).setCategory(token[4:])
                                                }
                                        case Condition:
                                                inner.Push(tv).setCategory(token[4:])
                                        }
                                } else {
                                        inner = ruleByLoP(token[4:])
                                        inner.Push(innot.setCategory(`or`))
                                }

                        } else {

                                // AND or OR condition/stack
                                inner, skip, err = parseBindRule(tokens[1:], depth, pspan, token)
                                inner.setCategory(lc(token))
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

                        tokens = tokens[1:]
                        cparen = isBC(tokens[0:3]) && tokens[3] == `)`
                        iparen = isBC(tokens[0:3]) && isWordOp(tokens[3])
                        pspan++

                // fallback == error condition
                default:
                        err = errorf("[%d] Unhandled token '%s'\n", ct, token)

                }

                // Go no further if any errors were encountered.
                if err != nil {
                        return
                }

                // Transfer our temporary contiguous condition
                // stack's contents into a new stack, which
                // is pushed singularly into the outer stack.
                if cc.Len() > 0 {
                        grp := ruleByLoP(outer.Category())
                        for j := 0; j < cc.Len(); j++ {
                                jidx, _ := cc.Index(j)
                                /*
                                printf("DEF:%s [%s;%s;oparen:%t;iparen:%t;ccparen:%t] [%T]\n",
                                        objectString(jidx),
                                        objectCategory(jidx),
                                        objectIdent(jidx),
                                        oparen,
                                        iparen,
                                        cc.isParen(),
                                        jidx)
                                */
                                grp.Push(jidx)
                        }

                        outer.Push(grp).Paren(oparen || iparen)
                        cc.reset()
                }

                // Traverse the inner stack, which contains nested
                // stacks and/or conditions that were acquired as
                // a result of a recursive (self-executing) call
                // of this same function. Migrate the inner stack's
                // contents into outer prior to a return.
                outer = transferToOuterBindRule(iparen, oparen, inner, outer)

                // Break out of the for-loop if we've been ordered
                // to do so ...
                if done {
                        break
                }

                // If we've made it here, we still have tokens left
                // to process. If any recursive calls have been made,
                // it might be necessary to skip ahead, should the
                // skip integer value be non-zero.
                switch skip {
                case 0:
                        continue
                default:
                        // If skip falls outside of the expected boundaries
                        // trash whatever tokens were remaining and return.
                        if skip == -1 || !( 0 <= skip && skip <= len(tokens)-1 ) {
                                tokens = []string{}
                                return
                        }

                        // Truncate our token stream, and reset the skip
                        // integer marker.
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

func transferToOuterBindRule(iparen, oparen bool, inner, outer Rule) Rule {

        if inner.Len() == 0 {
                return outer
        }

        r := ruleByLoP(outer.Category())
        r.Push(outer.Paren(oparen)).Paren(oparen)

        //printf("[len:%d]: %T[%s]: %s\n", r.Len(), r, r.Category(), r)

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

                        // Last-added outer slice was a Condition
                        case Condition:
                                //last = `C`
                                r.Push(ruleByLoP(outer.Category()).Push(tv))
                        }

                // Current inner slice is a Rule
                case Rule:

                        // Inner Rule's Boolean WORD operator does not
                        // match Inner slice [i] WORD operator. This
                        // would indicate a new stack (Rule) is coming
                        // up ...

                        tv.Paren(!oparen)
                        //printf("RULE :: 1/2 [len:%d;isparen:%t;iparen:%t;oparen:%t]: %T[%s]: %s\n", inner.Len(), inner.isParen(), iparen, outer.isParen(), inner, inner.Category(), inner)
                        //printf("RULE :: 2/2 [len:%d;isparen:%t;iparen:%t;oparen:%t]: %T[%s]: %s\n", tv.Len(), tv.isParen(), iparen, tv.isParen(), tv, tv.Category(), tv)
                        if inner.Category() != tv.Category() {
                                r.Push(ruleByLoP(inner.Category()).Push(tv))
                                break
                        }

                        // Push Inner slice [i] into Outer Rule.
                        r.Push(tv)

                }
        }

        return r
}

/*
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

        for i := 0; i < outer.Len(); i++ {
                sl, _ := outer.Index(i)

                switch tv := sl.(type) {
                case Rule:
                        printf("%s [%d; len:%d]: %T[%s]\n", tab, i, tv.Len(), tv, tv.Category())
                        if tv.Len() > 0 {
                                sub, _ := tv.Index(0)
                                printf("%T: %s\n", sub, objectString(sub))
                        }
                        printOuter(tv, tab, string(rune(9)))
                }
        }
}
*/

func bindRuleAssertWordOperator(t []string, token, icat string, depth, pspan int) (r Rule, skip int, ok bool, err error) {
        // this boolean operator merely continues the
        // expression, and does not signify a switch
        // to another operator, e.g.: AND -> OR, which
        // would result in a new recursion.
        if eq(token, icat) {
                return
        }

        // boolean operator differs from the current
        // (outer) operator. Begin new recursion, and
        // pass the desired word to this same function
        // which will result in an alloc for a new stack.
        //var innot Rule
        if eq(token, `and not`) {
                // negated condition/stack
                var innot Rule
                if innot, skip, err = parseBindRule(t[1:], depth, pspan); innot.Len() == 1 {
                        innar, _ := innot.Index(0)
                        r = ruleByLoP(token[4:])
                        switch tv := innar.(type) {
                        case Rule:
                                if tv.Len() == 1 {
                                        r.Push(tv.setCategory(`or`)).setCategory(token[4:])
                                }
                        case Condition:
                                r.Push(tv).setCategory(token[4:])
                        }
                } else {
                        r = ruleByLoP(token[4:])
                        r.Push(innot.setCategory(`or`))
                }

        } else {
                // AND or OR condition/stack
                r, skip, err = parseBindRule(t[1:], depth, pspan, token)
                r.setCategory(lc(token))
        }

        return
}

/*
parseBindRuleCondition returns a new Condition and an error instance following an attempt
to parse the stream of values (vals) along with the Bind Rule keyword (kw) and comparison
operator (cop).

This function is executed by parseBindRule during the Instruction parsing process.
*/
func parseBindRuleCondition(vals []string, kw, cop string) (c Condition, err error) {
	if len(vals) == 0 {
		err = errorf("Empty bind rule condition values; aborting")
		return
	}

	// assert the comparison operator
	if _, mo := matchOp(cop); !mo {
		err = badComparisonOperatorErr(cop)
		return
	}

	// This is the last (or only!) value component. We can
	// now analyze the keyword and the value(s) to ascertain
	// the appropriate instance type for condition storage
	// (and to perform other context-specific sanity checks).
	//
	// Begin with an assertion switch upon the target keyword
	// (which we already vetted as sane) ...
	switch key := matchBKW(kw); key {

	case BindUDN, BindRDN, BindGDN:
		c, err = assertBindRuleUGRDN(vals, key, cop)

	case BindUAT, BindGAT:
		c, err = assertBindRuleUGAttr(vals, key, cop)

	case BindToD:
		c, err = assertBindRuleTimeOfDay(vals, cop)

	case BindDoW:
		c, err = assertBindRuleDayOfWeek(vals, cop)

	case BindAM:
		c, err = assertBindRuleAuthMethod(vals, cop)

	case BindSSF:
		c, err = assertBindRuleSecurityStrengthFactor(vals, cop)

	case BindIP, BindDNS:
		c, err = assertBindRuleNet(vals, key, cop)
	}

	return
}

func checkBindRuleConditionValueStream(last, token, next string, ct int) (skip bool, err error) {
	// condition is ready for assembly AND a non-zero
	// double-quoted value is the current token. We
	// will also accept symbolic operators if we're
	// dealing with a multi-valued expression.
	if !isQuoted(token) && (token != `||` && token != `&&`) {
		err = errorf("Bogus Bind Rule condition expression between '%s' [%d] and '%s' [%d]; value must be a non-zero string enclosed within double quotes, or a symbolic list (||,&&) of same",
			last, ct-1, next, ct+1)
		return
	}

	switch {
	case next == `||` || next == `&&`:
		skip = true
	case !isQuoted(next) && next != `;` && next != `)`:
		err = errorf("Misaligned bind rule value expression ['%s' -> '%s']", token, next)
	}

	return
}

func assertBindRuleUGAttr(vals []string, key BindKeyword, op string) (c Condition, err error) {
	if err = unexpectedBindRuleConditionValueErr(key, 1, len(vals)); err != nil {
		return
	}

	// Prepare the expression value for our
	// Condition.
	ugat := trimS(unquote(vals[0]))
	if hasPfx(ugat, LocalScheme) {
		// value is an LDAP URI
		var uri LDAPURI
		if uri, err = parseLDAPURI(ugat, key); err != nil {
			return
		}

		c, err = conditionByOperator(op, uri)
		return
	}

	if hasPfx(ugat, `parent[`) {

		// value is an inheritance attributeBindTypeOrValue
		var inh Inheritance
		if inh, err = parseInheritance(ugat); err != nil {
			return
		}

		c, err = conditionByOperator(op, inh)
		return
	}

	// value is a standard attributeBindTypeOrValue
	var atb AttributeBindTypeOrValue
	if atb, err = parseATBTV(ugat, key); err != nil {
		return
	}
	c, err = conditionByOperator(op, atb)

	return
}

func assertBindRuleTimeOfDay(vals []string, op string) (c Condition, err error) {
	if err = unexpectedBindRuleConditionValueErr(BindToD, 1, len(vals)); err != nil {
		return
	}

	// extract clocktime from raw value, remove
	// quotes and any L/T WHSP
	raw := trimS(unquote(vals[0]))
	thyme := ToD(raw)
	if err = badClockTimeErr(raw, thyme.String()); err != nil {
		return
	}

	c, err = conditionByOperator(op, thyme)
	return
}

func assertBindRuleDayOfWeek(vals []string, op string) (c Condition, err error) {
	if err = unexpectedBindRuleConditionValueErr(BindDoW, 1, len(vals)); err != nil {
		return
	}

	// extract auth method from raw value, remove
	// quotes and any L/T WHSP and analyze
	raw := trimS(unquote(vals[0]))
	var dw DayOfWeek
	if dw, err = parseDoW(raw); err != nil {
		return
	}

	c, err = conditionByOperator(op, dw)
	return
}

func assertBindRuleAuthMethod(vals []string, op string) (c Condition, err error) {
	if err = unexpectedBindRuleConditionValueErr(BindAM, 1, len(vals)); err != nil {
		return
	}

	// extract auth method from raw value, remove
	// quotes and any L/T WHSP and analyze
	raw := trimS(unquote(vals[0]))
	am := matchAuthMethod(raw)
	if err = badAMErr(raw, am.String()); err != nil {
		return
	}

	c, err = conditionByOperator(op, am)
	return
}

func assertBindRuleSecurityStrengthFactor(vals []string, op string) (c Condition, err error) {
	if err = unexpectedBindRuleConditionValueErr(BindSSF, 1, len(vals)); err != nil {
		return
	}

	// extract factor from raw value, remove
	// quotes and any L/T WHSP
	raw := trimS(unquote(vals[0]))
	fac := SSF(raw)
	if err = badSecurityStrengthFactorErr(raw, fac.String()); err != nil {
		return
	}

	c, err = conditionByOperator(op, fac)
	return
}

func assertBindRuleNet(vals []string, key BindKeyword, op string) (c Condition, err error) {
	if err = unexpectedBindRuleConditionValueErr(key, 1, len(vals)); err != nil {
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

		if err = badIPErr(len(raw), addr.Len()); err != nil {
			return
		}

		c, err = conditionByOperator(op, addr)
		return
	}

	// extract FQDN from raw value, remove
	// quotes and any L/T WHSP.
	raw := trimS(unquote(vals[0]))
	fq := DNS(raw)
	if err = badDNSErr(raw, fq.String()); err != nil {
		return
	}

	c, err = conditionByOperator(op, fq)
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

func badClockTimeErr(raw, thyme string) (err error) {
	if thyme != raw {
		err = errorf("Unexpected %s clock time parsing result; want '%s', got '%s' (hint: use military time; 0000 through 2400)",
			BindToD, raw, thyme)
	}
	return
}

func badSecurityStrengthFactorErr(raw, fac string) (err error) {
	if fac != raw {
		err = errorf("Unexpected security strength factor parsing result; want '%s', got '%s' (hint: valid range is 0-256)",
			raw, fac)
	}
	return
}

func unexpectedBindRuleConditionValueErr(key BindKeyword, want, got int) (err error) {
	if want != got {
		err = errorf("Unexpected number of %s values; want %d, got %d", key, want, got)
	}
	return
}

func badIPErr(want, got int) (err error) {
	if want != got {
		err = errorf("Unexpected '%s' values parsed; want '%d', got '%d'", BindIP, want, got)
	}
	return
}

func badDNSErr(raw, fqdn string) (err error) {
	if raw != fqdn {
		err = errorf("Unexpected '%s' values parsed; want '%s', got '%s'", BindDNS, raw, fqdn)
	}
	return
}

func badAMErr(raw, am string) (err error) {
	if !eq(am, raw) {
		err = errorf("Unexpected %s auth method parsing result; want '%s', got '%s'", BindAM, raw, am)
	}
	return
}

func badComparisonOperatorErr(cop string) (err error) {
	err = errorf("Unidentified or misaligned bind rule comparison operator '%s'; aborting", cop)
	return
}

func badBindRuleKeyword(cop string) (err error) {
	err = errorf("Unknown bind rule keyword '%s'", cop)
	return
}

const bindRuleID = `bind`

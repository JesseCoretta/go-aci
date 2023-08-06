package aci

/*
parseBindRuleCondition returns a new Condition and an error instance following an attempt
to parse the stream of values (vals) along with the Bind Rule keyword (kw) and comparison
operator (cop).

This function is executed by parseBR during the Instruction parsing process.
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

func bindRuleCaseWordOperator(tokens []string, ocat string, chop, depth, pspan int) (inner Rule, skipTo int, err error) {
	if len(tokens) < 5 {
		err = errorf("Cannot process token stream for word operator; too few tokens (%d): %v",
			len(tokens), tokens)
		return
	}

	var token string = tokens[0]
	var next string = tokens[1]
	var ttoken string
	var offset int = 1

	// go-stackage's negation stacks use the
	// category of 'NOT', as opposed to ACIv3's
	// 'AND NOT' operator equivalent. Take the
	// 'NOT' portion of the value, using its
	// original case-folding scheme, and save
	// it for stack tagging later.
	if eq(token, `and not`) {
		ttoken = token[4:]
		printf("TOCK: %s [OFF:%d]\n", ttoken)
	} else if eq(next, `and not`) {
		ttoken = next[4:]
		offset++
		printf("TICK: %s [OFF:%d]\n", ttoken)
	}
	offset = 5

	// If the category (word operator) is not
	// the same as the token, this means a new
	// distinct (inner) stack is beginning (and
	// not a continuation of outer).
	if !eq(ttoken, ocat) {

		// We need to offset the truncation factor
		// of our token slices when the 'AND NOT'
		// logical Boolean WORD operator is used,
		// as it will erroneously be interpreted
		// as two (2) distinct tokens.
		if eq(ttoken, `not`) {
			offset++
		}

		// look ahead for an opening parenthetical ...
		iparen := tokens[2] == `(`

		// Launch a new inner recursion of this
		// same function.
		_, _, oip, _ := parenState(join(tokens[offset:], ``))
		if skipTo, inner, err = parseBR(tokens[offset:], depth, pspan); err != nil {
			return
		}

		printf("B SKIPTO:%d\n", skipTo)
		skipTo += chop - offset + 1
		printf("A SKIPTO:%d\n", skipTo)

		// If the inner stack has at least one
		// (1) element, preserve it for the end
		// stack element, else take no action.
		if inner.Len() > 0 {
			inner.Paren(oip || iparen)
			inner.setCategory(ttoken) // mark the inner stack's logical Boolean WORD operator
		}
	}

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

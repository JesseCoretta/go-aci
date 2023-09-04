package aci

import (
	"errors"
)

/*
errorf wraps errors.New and returns a non-nil instance of error
based upon a non-nil/non-zero msg input value with optional args.
*/
func errorf(msg any, x ...any) error {
	switch tv := msg.(type) {
	case string:
		if len(tv) > 0 {
			return errors.New(sprintf(tv, x...))
		}
	case error:
		if tv != nil {
			return errors.New(sprintf(tv.Error(), x...))
		}
	}

	return nil
}

func nilInstanceErr(x any) error {
	return errorf("%T instance is nil", x)
}

/*
badInhErr returns an error describing the appropriate syntax and displaying the offending value.
*/
func badInhErr(bad string) error {
	return errorf("Bad Inheritance value '%s'; must conform to 'parent[0-9+].<at>#<bt_or_av>'", bad)
}

func badCopErr(instance any) error {
	return errorf("%T contains an bogus underlying value", instance)
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

func unexpectedBindConditionValueErr(key BindKeyword, want, got int) (err error) {
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

func badPTBRuleKeywordErr(candidate any, typ string, want, got any) error {
	var kw, kg string
	switch tv := want.(type) {
	case string:
		kw = tv
	case Keyword:
		kw = tv.String()
	}

	switch tv := got.(type) {
	case string:
		kg = tv
	case Keyword:
		kg = tv.String()
	}

	emsg := "Unknown or unresolvable %s rule keyword or category for %T: want '%s', got '%s'"
	return errorf(emsg, typ, candidate, kw, kg)
}

func noTBRuleExpressionValues(candidate any, typ string, key Keyword) error {
	emsg := "Found no %s %s rule expression value(s) during %T assertion"
	return errorf(emsg, key, typ, candidate)
}

func badAttributeBindTypeOrValueErr(x string) error {
	return errorf("Invalid AttributeBindTyoeOrValue instance: must conform to '<at>#<bt_or_av>', got '%s'", x)
}

func badObjectIdentifierErr(x string) error {
	return errorf("Invalid ObjectIdentifier instance: must conform to 'N[.N]+', got '%s'", x)
}

func badObjectIdentifierKeywordErr(key TargetKeyword) error {
	emsg := "Invalid %s and/or %T[%s] value(s)"
	return errorf(emsg, `ObjectIdentifier`, key, key)
}

func unexpectedKindErr(receiver any, want, got string) error {
	return errorf("Unexpected %T.Kind result: should be '%s', got '%s'", receiver, want, got)
}

func illegalSliceTypeErr(receiver, want, got any) error {
	return errorf("Illegal slice type within %T stack: should be '%T', got '%T'",
		receiver, want, got)
}

func illegalSyntaxPerTypeErr(candidate any, key Keyword, er ...error) error {
	var err error
	var kw string = `<unspecified_keyword>`
	if len(er) > 0 {
		err = er[0]
	}
	if key != nil {
		kw = key.String()
	}

	emsg := "%T [%s] syntax invalid"

	if err != nil {
		emsg += `: %v`
		return errorf(emsg, candidate, kw, err)
	}
	return errorf(emsg, candidate, kw)
}

func afosNonIdempSplitErr(d, l, o int, c rune) error {
	emsg := "Inappropriate delimiter id [%d]; non-idempotent result following '%c' split: len(vals)=%d, opct=%d"
	return errorf(emsg, d, c, l, o)
}

func afoMissingPrefixErr() error {
	emsg := "%T instance is missing required %T prefix: needs either add= or delete="
	return errorf(emsg, AttributeFilterOperation{}, AttributeOperation(0))
}

func aoBadPrefixErr() error {
	emsg := "Invalid %T value prefix; must be add= or delete="
	return errorf(emsg, AttributeOperation(0))
}

func uriBadPrefixErr() error {
	emsg := "Invalid %T value prefix; must be %s"
	return errorf(emsg, LDAPURI{}, LocalScheme)
}

func afMissingDelimiterErr(af AttributeFilter) error {
	emsg := "No attr:filter delimiter (%c) found in %T"
	return errorf(emsg, ':', af)
}

func instructionNoLabelErr() error {
	emsg := "%T has no name (ACL); set a string name value using %T.Set"
	return errorf(emsg, Instruction{}, Instruction{})
}

func levelsNotFoundErr() error {
	emsg := "No level identifiers parsed; aborting"
	return errorf(emsg)
}

func rightNotfound(x string) error {
	emsg := "Right '%s' unknown"
	return errorf(emsg, x)
}

func dowBadTimeErr() error {
	emsg := "%T instance describes invalid timeofday"
	return errorf(emsg, TimeOfDay{})
}

func dowBadDayErr(x any) error {
	emsg := "%T instance describes invalid dayofweek: %v"
	return errorf(emsg, DayOfWeek{}, x)
}

func noPermissionDispErr() error {
	emsg := "%T has no disposition (allow/deny)"
	return errorf(emsg, Permission{})
}

func fqdnInvalidLabelErr(l domainLabel) error {
	emsg := "%T has a bad label %v"
	return errorf(emsg, FQDN{}, l)
}

func fqdnInvalidLenErr(l int) error {
	emsg := "%T size requirements not met (%d v. %d), or len < 2"
	return errorf(emsg, FQDN{}, l, fqdnMax)
}

func parseBindRuleInvalidExprTypeErr(receiver, want, got any) error {
	emsg := "Unexpected %T within %T; wanted %T"
	return errorf(emsg, got, receiver, want)
}

func parseBindRulesHierErr(stk any, b BindRules) error {
	emsg := "Unable to cast %T hierarchy into %T; aborting"
	return errorf(emsg, stk, b)
}

func unexpectedStringResult(typ, want, got string) error {
	emsg := "Unexpected string %s result; want %s, got %s"
	return errorf(emsg, typ, want, got)
}

func unexpectedValueCountErr(typ string, want, got int) error {
	emsg := "Unexpected number of %s values; want %d, got %d"
	return errorf(emsg, typ, want, got)
}

func bogusValueErr(typ, bogus string) error {
	emsg := "Bogus %s value: '%s'"
	return errorf(emsg, typ, bogus)
}

func generalErr(typ string, err error) error {
	emsg := "General %s error: %v"
	return errorf(emsg, typ, err)
}

func noValueErr(candidate any, typ string) error {
	emsg := "Found no %s expression during processing of %T instance"
	return errorf(emsg, typ, candidate)
}

func pushError(receiver, candidate any, key Keyword, emsg string, er ...error) error {
	var err error
	var kw string = `<unspecified_keyword>`
	if len(er) > 0 {
		err = er[0]
	}
	if key != nil {
		kw = key.String()
	}

	if err != nil {
		return errorf(emsg, candidate, receiver, kw, err)
	}
	return errorf(emsg, candidate, receiver, kw)
}

func pushErrorNotUnique(receiver, candidate any, key Keyword, er ...error) error {
	emsg := "Cannot push non-unique or ineligible %T into %T [%s]"
	return pushError(receiver, candidate, key, emsg, er...)
}

func pushErrorNilOrZero(receiver, candidate any, key Keyword, er ...error) error {
	var emsg string = "Cannot push zero-length or nil %T into %T [%s]: %v"
	return pushError(receiver, candidate, key, emsg, er...)
}

func pushErrorBadType(receiver, candidate any, key Keyword, er ...error) error {
	var emsg string = "Push request of %T type violates %T [%s] PushPolicy"
	return pushError(receiver, candidate, key, emsg, er...)
}

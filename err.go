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

func missingLocalSchemeDNErr(key Keyword, x int, D string) (err error) {
	if !hasPfx(D, LocalScheme) {
		err = errorf("Illegal %s distinguishedName slice: [index:%d;value:%s] missing LDAP local scheme (%s)",
			key, x, D, LocalScheme)
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

func badComparisonOperatorErr(cop string) (err error) {
	err = errorf("Unidentified or misaligned bind rule comparison operator '%s'; aborting", cop)
	return
}

func badBindRuleKeyword(cop string) (err error) {
	err = errorf("Unknown bind rule keyword '%s'", cop)
	return
}

func badAttributeBindTypeOrValueErr(x string) error {
	return errorf("Invalid AttributeBindTyoeOrValue instance: must conform to '<at>#<bt_or_av>', got '%s'", x)
}

func badObjectIdentifierErr(x string) error {
	return errorf("Invalid ObjectIdentifier instance: must conform to 'N[.N]+', got '%s'", x)
}

func unexpectedKindErr(r any, want, got string) error {
	return errorf("Unexpected %T.Kind result: should be '%s', got '%s'", r, want, got)
}

func illegalSliceType(r, want, got any) error {
	return errorf("Illegal slice type within %T stack: should be '%T', got '%T'",
		r, want, got)
}

package aci

import (
	"testing"
)

func matchBadDNString(x string) (string, bool) {
	return x, (eq(x, badBDN) || eq(x, badTDN))
}

func TestBindDistinguishedName_codecov(t *testing.T) {
	for kw, fn := range map[BindKeyword]func(string) BindDistinguishedName{
		BindUDN: UDN,
		BindRDN: RDN,
		BindGDN: GDN,
	} {
		for _, dn := range []string{
			`uid=jesse,ou=People,dc=example,dc=com`,
			`cn=Courtney Tolana,ou=Contractors,ou=People,dc=example,dc=com`,
		} {
			var O BindDistinguishedName

			if err := O.Valid(); err == nil {
				t.Errorf("%s failed: invalid %T returned no validity error",
					t.Name(), O)
			}

			if !O.IsZero() {
				t.Errorf("%s failed: non-zero %T", t.Name(), O)
			}

			if msg, bad := matchBadDNString(O.String()); !bad {
				t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
					t.Name(), msg, O)
			}

			// process OID
			O = fn(dn)

			// OIDs qualify for equality and negated equality
			// comparison operators.
			cops := map[ComparisonOperator]func() BindRule{
				Eq: O.Eq,
				Ne: O.Ne,
			}

			// try every comparison operator supported in
			// this context ...
			for c := 1; c < len(cops)+1; c++ {
				cop := ComparisonOperator(c)
				wcop := sprintf("%s %s %q", O.Keyword(), cop, O.String())

				// create targetrule B using comparison
				// operator (cop).
				if B := cops[cop](); B.String() != wcop {
					err := unexpectedStringResult(kw.String(), wcop, B.String())
					t.Errorf("%s failed [%s rule]: %v", t.Name(), kw.String(), err)
				}

			}
		}
	}
}

func TestTargetDistinguishedName_codecov(t *testing.T) {
	for kw, fn := range map[TargetKeyword]func(string) TargetDistinguishedName{
		Target:     TDN,
		TargetTo:   TTDN,
		TargetFrom: TFDN,
	} {
		for _, dn := range []string{
			`uid=jesse,ou=People,dc=example,dc=com`,
			`cn=Courtney Tolana,ou=Contractors,ou=People,dc=example,dc=com`,
		} {
			var O TargetDistinguishedName

			if err := O.Valid(); err == nil {
				t.Errorf("%s failed: invalid %T returned no validity error",
					t.Name(), O)
			}

			if !O.IsZero() {
				t.Errorf("%s failed: non-zero %T", t.Name(), O)
			}

			if msg, bad := matchBadDNString(O.String()); !bad {
				t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
					t.Name(), msg, O)
			}

			// process OID
			O = fn(dn)

			// OIDs qualify for equality and negated equality
			// comparison operators.
			cops := map[ComparisonOperator]func() TargetRule{
				Eq: O.Eq,
				Ne: O.Ne,
			}

			// try every comparison operator supported in
			// this context ...
			for c := 1; c < len(cops)+1; c++ {
				cop := ComparisonOperator(c)
				wcop := sprintf("( %s %s %q )", O.Keyword(), cop, O.String())

				// create targetrule T using comparison
				// operator (cop).
				if T := cops[cop](); T.String() != wcop {
					err := unexpectedStringResult(kw.String(), wcop, T.String())
					t.Errorf("%s failed [%s rule]: %v", t.Name(), kw.String(), err)
				}

			}
		}
	}
}

func TestBindDistinguishedNames_codecov(t *testing.T) {
	var Os BindDistinguishedNames

	for kw, fn := range map[BindKeyword]func(...any) BindDistinguishedNames{
		BindUDN: UDNs,
		BindRDN: RDNs,
		BindGDN: GDNs,
	} {
		Os = fn()
		Os.Push(kw)

		for _, dn := range []string{
			`uid=jesse,ou=People,dc=example,dc=com`,
			`cn=Courtney Tolana,ou=Contractors,ou=People,dc=example,dc=com`,
		} {
			var O BindDistinguishedName
			var Ol int = Os.Len()

			if err := O.Valid(); err == nil {
				t.Errorf("%s multival failed: invalid %T returned no validity error",
					t.Name(), O)
			}

			if msg, bad := matchBadDNString(O.String()); !bad {
				t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
					t.Name(), msg, O)
			}

			if Os.Push(O); Os.Len() > Ol {
				t.Errorf("%s multival failed: invalid %T instance pushed into %T without error",
					t.Name(), O, Os)
			}

			// process OID
			O = Os.F()(dn)

			Ol = Os.Len()
			Os.Push(O)
			if Os.Len() != Ol+1 {
				t.Errorf("%s multival failed: valid %T[%s] instance (%s) not pushed into %T[%s; len:%d]",
					t.Name(), O, O.Keyword(), O, Os, Os.Keyword(), Ol)
			}

			// OIDs qualify for equality and negated equality
			// comparison operators.
			cops := map[ComparisonOperator]func() BindRule{
				Eq: Os.Eq,
				Ne: Os.Ne,
			}

			// try every comparison operator supported in
			// this context ...
			for c := 1; c < len(cops)+1; c++ {
				cop := ComparisonOperator(c)
				wcop := sprintf("( %s %s %q )", O.Keyword(), cop, Os.String())

				// create targetrule T using comparison
				// operator (cop).
				if B := cops[cop](); B.Paren().String() != wcop {
					err := unexpectedStringResult(kw.String(), wcop, B.String())
					t.Errorf("%s multival failed [%s rule]: %v", t.Name(), kw.String(), err)
				}

			}
		}
	}
}

/*
func TestDistinguishedName_codecov(t *testing.T) {
	var (
		bd BindDistinguishedName
		td TargetDistinguishedName
	)

	user := `cn=Jesse Coretta,ou=People,dc=example,dc=com`
	if td.Set(user, Target); td.IsZero() {
		t.Errorf("%s failed [%T.Set]: instance is nil",
			t.Name(), td)
	}
	if bd.Set(user, BindUDN); bd.IsZero() {
		t.Errorf("%s failed [%T.Set]: instance is nil",
			t.Name(), bd)
	}

	want := LocalScheme + user
	if got := td.String(); want != got {
		t.Errorf("%s failed [%T.String compare]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), td, want, got)
	}

	if got := bd.String(); want != got {
		t.Errorf("%s failed [%T.String compare]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), bd, want, got)
	}
}
*/

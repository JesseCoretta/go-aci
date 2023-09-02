package aci

import (
	"fmt"
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
			O.isDistinguishedNameContext()
			_ = O.Eq()
			_ = O.Ne()
			_ = O.Len()

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

			// process DN
			O = fn(dn)
			_ = O.Len()
			dnBefore := O.String()

			O.Set(dn) // set to same DN using Set method
			dnAfter := O.String()
			if dnBefore != dnAfter {
				err := unexpectedStringResult(kw.String(), dnBefore, dnAfter)
				t.Errorf("%s failed: %v",
					t.Name(), err)
			}

			O.Set(`159`)
			O.Set(``)
			O.Set(`uid=jesse,ou=People,dc=example,dc=com`, kw)

			// try every comparison operator supported in
			// this context ...
			brf := O.BRF()
			for i := 0; i < brf.Len(); i++ {
				cop, meth := brf.Index(i + 1)
				wcop := sprintf("( %s %s \"ldap:///uid=jesse,ou=People,dc=example,dc=com\" )", O.Keyword(), cop)
				if T := meth(); T.Paren().String() != wcop {
					err := unexpectedStringResult(O.String(), wcop, T.String())
					t.Errorf("%s [%s] multival failed [%s rule]: %v",
						t.Name(), O.Keyword(), kw, err)
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
			O.isDistinguishedNameContext()
			_ = O.Eq()
			_ = O.Ne()
			_ = O.Len()

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

			// process DN
			O = fn(dn)
			_ = O.Len()
			dnBefore := O.String()

			O.Set(dn) // set to same DN using Set method
			dnAfter := O.String()
			if dnBefore != dnAfter {
				err := unexpectedStringResult(kw.String(), dnBefore, dnAfter)
				t.Errorf("%s failed: %v",
					t.Name(), err)
			}

			O.Set(`159`)
			O.Set(``)
			O.Set(`#barf`, kw)

			// DNs qualify for equality and negated equality
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
	var Id string = Os.ID()
	var Kw Keyword = Os.Keyword()
	_ = sprintf("%v", Kw)
	_ = Os.Eq()
	_ = Os.Ne()
	Os.reset()
	_ = Os.setQuoteStyle(0)
	_ = Os.setQuoteStyle(1)

	for kw, fn := range map[BindKeyword]func(...any) BindDistinguishedNames{
		BindUDN: UDNs,
		BindGDN: GDNs,
		BindRDN: RDNs,
	} {
		Os = fn()
		Os.Push(kw) // reset the keyword for this iteration

		for _, dn := range []string{
			`uid=jesse,ou=People,dc=example,dc=com`,
			`cn=Courtney Tolana,ou=Contractors,ou=People,dc=example,dc=com`,
		} {
			var O BindDistinguishedName
			var Ol int = Os.Len()

			if err := O.Valid(); err == nil {
				t.Errorf("%s [%s] multival failed: invalid %T returned no validity error",
					t.Name(), Id, O)
			}

			if msg, bad := matchBadDNString(O.String()); !bad {
				t.Errorf("%s [%s] failed: unexpected string result; want '%s', got '%s'",
					t.Name(), Id, msg, O)
			}

			if Os.Push(O); Os.Len() > Ol {
				t.Errorf("%s [%s]multival failed: invalid %T instance pushed into %T without error",
					t.Name(), Id, O, Os)
			}

			// process DN
			O = Os.F()(dn)

			Ol = Os.Len()
			Os.Push(O)
			Os.Push(Os.Pop())
			Os.Push(O) // try to introduce duplicate
			Id = Os.ID()
			Kw = Os.Keyword()

			if Os.Len() != Ol+1 {
				t.Errorf("%s [%s] multival failed: valid %T[%s] instance (%s) not pushed into %T[%s; len:%d]",
					t.Name(), Id, O, O.Keyword(), O, Os, Kw, Ol)
			}

			_ = Os.setQuoteStyle(0)
			_ = Os.setQuoteStyle(1)

			// DNs qualify for equality and negated equality
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
					t.Errorf("%s [%s] multival failed [%s rule]: %v", t.Name(), Id, kw.String(), err)
				}

			}
		}
	}
}

func TestTargetDistinguishedNames_codecov(t *testing.T) {
	var Os TargetDistinguishedNames
	var Id string = Os.ID()
	var Kw Keyword = Os.Keyword()
	_ = sprintf("%v", Kw)
	_ = Os.Eq()
	_ = Os.Ne()
	Os.reset()
	_ = Os.setQuoteStyle(0)
	_ = Os.setQuoteStyle(1)

	for kw, fn := range map[TargetKeyword]func(...any) TargetDistinguishedNames{
		Target:     TDNs,
		TargetTo:   TTDNs,
		TargetFrom: TFDNs,
	} {
		Os = fn()
		Os.Push(kw) // reset the keyword for this iteration

		for _, dn := range []string{
			`uid=jesse,ou=People,dc=example,dc=com`,
			`cn=Courtney Tolana,ou=Contractors,ou=People,dc=example,dc=com`,
		} {
			var O TargetDistinguishedName
			var Ol int = Os.Len()

			if err := O.Valid(); err == nil {
				t.Errorf("%s [%s] multival failed: invalid %T returned no validity error",
					t.Name(), Id, O)
			}

			if msg, bad := matchBadDNString(O.String()); !bad {
				t.Errorf("%s [%s] failed: unexpected string result; want '%s', got '%s'",
					t.Name(), Id, msg, O)
			}

			if Os.Push(O); Os.Len() > Ol {
				t.Errorf("%s [%s] multival failed: invalid %T instance pushed into %T without error",
					t.Name(), Id, O, Os)
			}

			// process DN
			O = Os.F()(dn)

			Ol = Os.Len()
			Os.Push(O)
			Os.Push(Os.Pop())
			Os.Push(O) // try to introduce duplicate
			Id = Os.ID()
			Kw = Os.Keyword()

			if Os.Len() != Ol+1 {
				t.Errorf("%s [%s] multival failed: valid %T[%s] instance (%s) not pushed into %T[%s; len:%d]",
					t.Name(), Id, O, O.Keyword(), O, Os, Kw, Ol)
			}

			_ = Os.setQuoteStyle(0)
			_ = Os.setQuoteStyle(1)

			// DNs qualify for equality and negated equality
			// comparison operators.
			cops := map[ComparisonOperator]func() TargetRule{
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
				if B := cops[cop](); B.String() != wcop {
					err := unexpectedStringResult(kw.String(), wcop, B.String())
					t.Errorf("%s [%s] multival failed [%s rule]: %v", t.Name(), Id, kw.String(), err)
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

/*
This example demonstrates the SHA-1 hash comparison between two (2) instances of
BindDistinguishedName using the Compare method.

These seemingly different distinguished names actually evaluate as equal because
the local LDAP prefix (ldap:///) is stripped off when the UDN (or similar) function
is run. When the String method is executed for a distinguished name at a later point,
said prefix is artificially imposed. Therefore in this scenario, the two values are
in fact equal.
*/
func ExampleBindDistinguishedName_Compare() {
	dn1 := UDN(`uid=jesse,ou=People,dc=example,dc=com`)
	dn2 := UDN(`ldap:///uid=jesse,ou=People,dc=example,dc=com`)

	fmt.Printf("Hashes are equal: %t", dn1.Compare(dn2))
	// Output: Hashes are equal: true
}

/*
This example demonstrates the SHA-1 hash comparison between two (2) instances of
BindDistinguishedNames using the Compare method.

The comparison evaluates as false because, while the contents match, their ordering
does not.
*/
func ExampleBindDistinguishedNames_Compare() {
	adns := UDNs(`uid=jesse,ou=People,dc=example,dc=com`, `uid=courtney,ou=People,dc=example,dc=com`)
	odns := UDNs(`uid=courtney,ou=People,dc=example,dc=com`, `uid=jesse,ou=People,dc=example,dc=com`)

	fmt.Printf("Hashes are equal: %t", odns.Compare(adns))
	// Output: Hashes are equal: false
}

/*
This example demonstrates the SHA-1 hash comparison between two (2) instances of
BindDistinguishedNames using the Compare method.

The comparison evaluates as false because, while the contents match, their ordering
does not.
*/
func ExampleTargetDistinguishedNames_Compare() {
	adns := TFDNs(`uid=jesse,ou=People,dc=example,dc=com`, `uid=courtney,ou=People,dc=example,dc=com`)
	odns := TFDNs(`uid=courtney,ou=People,dc=example,dc=com`, `uid=jesse,ou=People,dc=example,dc=com`)

	fmt.Printf("Hashes are equal: %t", odns.Compare(adns))
	// Output: Hashes are equal: false
}

/*
This example demonstrates the SHA-1 hash comparison between two (2) instances of
TargetDistinguishedName using the Compare method.

The comparison evaluates as false because the case folding schemes do not match
for the `T` in Tolana.
*/
func ExampleTargetDistinguishedName_Compare() {
	dn1 := TDN(`cn=Courtney tolana,ou=People,dc=example,dc=com`)
	dn2 := TDN(`cn=Courtney Tolana,ou=People,dc=example,dc=com`)

	fmt.Printf("Hashes are equal: %t", dn1.Compare(dn2))
	// Output: Hashes are equal: false
}

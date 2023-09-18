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
			_ = O.Keyword()
			_ = O.Valid()
			_ = O.IsZero()
			_ = O.Len()
			_ = isDNAlias(AllDN.String())
			_ = isDNAlias(``)

			if err := O.Valid(); err == nil {
				t.Errorf("%s failed: invalid %T returned no validity error",
					t.Name(), O)
				return
			}

			if !O.IsZero() {
				t.Errorf("%s failed: non-zero %T", t.Name(), O)
				return
			}

			if msg, bad := matchBadDNString(O.String()); !bad {
				t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
					t.Name(), msg, O)
				return
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
				return
			}

			O.Set(`159`)
			O.Set(``)
			_ = O.Valid()
			O.Set(`uid=jesse,ou=People,dc=example,dc=com`, kw)

			// try every comparison operator supported in
			// this context ...
			brm := O.BRM()
			for i := 0; i < brm.Len(); i++ {
				cop, meth := brm.Index(i + 1)
				wcop := sprintf("( %s %s \"ldap:///uid=jesse,ou=People,dc=example,dc=com\" )", O.Keyword(), cop)
				if T := meth(); T.Paren().String() != wcop {
					err := unexpectedStringResult(O.String(), wcop, T.String())
					t.Errorf("%s [%s] multival failed [%s rule]: %v",
						t.Name(), O.Keyword(), kw, err)
					return
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
			_ = O.Keyword()
			_ = O.Valid()
			_ = O.IsZero()
			_ = O.Len()

			if err := O.Valid(); err == nil {
				t.Errorf("%s failed: invalid %T returned no validity error",
					t.Name(), O)
				return
			}

			if !O.IsZero() {
				t.Errorf("%s failed: non-zero %T", t.Name(), O)
				return
			}

			if msg, bad := matchBadDNString(O.String()); !bad {
				t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
					t.Name(), msg, O)
				return
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
				return
			}

			O.Set(`159`)
			O.Set(``)
			_ = O.Valid()
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
					return
				}

			}
		}
	}
}

func TestBindDistinguishedNames_codecov(t *testing.T) {
	var ctx DistinguishedNameContext = UDN(`uid=jesse,ou=People,dc=example,dc=com`)
	var Os BindDistinguishedNames
	var Id string = Os.ID()
	var Kw Keyword = Os.Keyword()
	_ = sprintf("%v", Kw)
	_ = Os.Eq()
	_ = Os.Ne()
	_ = Os.Push()
	_ = Os.Push(``)
	_ = Os.Push(`fhksjthg4`)
	_ = Os.Push(`_1`)
	_ = Os.Push(ctx)
	_ = Os.Push(URI(`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName?one?(&(objectClass=contractor)(status=active))`))

	Os.reset()
	_ = Os.setQuoteStyle(0)
	_ = Os.setQuoteStyle(1)

	Os.setExpressionValues(Target, []string{}...)
	Os.setExpressionValues(BindUDN, []string{}...)
	Os.setExpressionValues(BindSSF, []string{`1`}...)
	Os.setExpressionValues(BindSSF, []string{`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName?one?(&(objectClass=contractor)(status=active))`}...)
	Os.setExpressionValues(BindUDN, []string{`325Ga_`}...)

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
			`ldap:///ou=People,dc=example,dc=com?cn,sn,objectClass?one?(objectClass=employee)`,
		} {
			var O BindDistinguishedName
			var Ol int = Os.Len()

			if err := O.Valid(); err == nil {
				t.Errorf("%s [%s] multival failed: invalid %T returned no validity error",
					t.Name(), Id, O)
				return
			}

			if msg, bad := matchBadDNString(O.String()); !bad {
				t.Errorf("%s [%s] failed: unexpected string result; want '%s', got '%s'",
					t.Name(), Id, msg, O)
				return
			}

			if Os.Push(O); Os.Len() > Ol {
				t.Errorf("%s [%s]multival failed: invalid %T instance pushed into %T without error",
					t.Name(), Id, O, Os)
				return
			}

			// process DN
			O = Os.F()(dn)

			Ol = Os.Len()
			Os.Push(O)
			Os.Push(``)
			Os.Push(nil)
			Os.Push('a')
			Os.Contains(dn)
			Os.Push(Os.Pop())
			Os.Push(O) // try to introduce duplicate
			Id = Os.ID()
			Kw = Os.Keyword()

			if Os.Len() != Ol+1 {
				t.Errorf("%s [%s] multival failed: valid %T[%s] instance (%s) not pushed into %T[%s; len:%d]",
					t.Name(), Id, O, O.Keyword(), O, Os, Kw, Ol)
				return
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
					return
				}

			}
		}
	}
}

func TestTargetDistinguishedNames_codecov(t *testing.T) {
	var ctx DistinguishedNameContext = UDN(`uid=jesse,ou=People,dc=example,dc=com`)
	var Os TargetDistinguishedNames
	var Id string = Os.ID()
	var Kw Keyword = Os.Keyword()
	_ = sprintf("%v", Kw)
	_ = Os.Eq()
	_ = Os.Ne()
	_ = Os.Push()
	_ = Os.Push(``)
	_ = Os.Push(`fhksjthg4`)
	_ = Os.Push(`_1`)
	_ = Os.Push(ctx)

	Os.reset()
	_ = Os.setQuoteStyle(0)
	_ = Os.setQuoteStyle(1)

	Os.setExpressionValues(BindGDN, []string{}...)
	Os.setExpressionValues(Target, []string{}...)
	Os.setExpressionValues(BindSSF, []string{`1`}...)
	Os.setExpressionValues(BindSSF, []string{`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName?one?(&(objectClass=contractor)(status=active))`}...)
	Os.setExpressionValues(TargetTo, []string{`325Ga_`}...)

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
				return
			}

			if msg, bad := matchBadDNString(O.String()); !bad {
				t.Errorf("%s [%s] failed: unexpected string result; want '%s', got '%s'",
					t.Name(), Id, msg, O)
				return
			}

			if Os.Push(O); Os.Len() > Ol {
				t.Errorf("%s [%s] multival failed: invalid %T instance pushed into %T without error",
					t.Name(), Id, O, Os)
				return
			}

			// process DN
			O = Os.F()(dn)

			Ol = Os.Len()
			Os.Push(O)
			Os.Push(``)
			Os.Push(nil)
			Os.Push('a')
			Os.Contains(dn)
			Os.Push(Os.Pop())
			Os.Push(O) // try to introduce duplicate
			Id = Os.ID()
			Kw = Os.Keyword()

			if Os.Len() != Ol+1 {
				t.Errorf("%s [%s] multival failed: valid %T[%s] instance (%s) not pushed into %T[%s; len:%d]",
					t.Name(), Id, O, O.Keyword(), O, Os, Kw, Ol)
				return
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
					return
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

func ExampleBindDistinguishedNames_Contains() {
	dn := GDNs(
		`cn=Executives,ou=Groups,dc=example,dc=com`,
		`cn=Executive Assistants,ou=Groups,dc=example,dc=com`,
	)

	execs := `cn=Executives,ou=Groups,dc=example,dc=com`

	fmt.Printf("Found: %t", dn.Contains(execs))
	// Output: Found: true
}

func ExampleTargetDistinguishedNames_Contains() {
	dn := TFDNs(
		`cn=*,ou=Profiles,dc=example,dc=com`,
		`cn=*,ou=People,dc=example,dc=com`,
	)

	groups := `cn=*,ou=Groups,dc=example,dc=com`

	fmt.Printf("Found: %t", dn.Contains(groups))
	// Output: Found: false
}

func ExampleBindDistinguishedNames_Push() {
	var odns BindDistinguishedNames = UDNs()
	odns.Push(
		`uid=jesse,ou=People,dc=example,dc=com`,
		`uid=jimmy,ou=People,dc=example,dc=com`,
	)

	fmt.Printf("Length: %d", odns.Len())
	// Output: Length: 2
}

func ExampleTargetDistinguishedNames_Push() {
	var odns TargetDistinguishedNames = TDNs()
	odns.Push(
		`uid=jesse,ou=People,dc=example,dc=com`,
		`uid=jimmy,ou=People,dc=example,dc=com`,
	)

	fmt.Printf("Length: %d", odns.Len())
	// Output: Length: 2
}

func ExampleBindDistinguishedNames_Pop() {
	var odns BindDistinguishedNames = UDNs()
	odns.Push(
		`uid=jesse,ou=People,dc=example,dc=com`,
		`uid=jimmy,ou=People,dc=example,dc=com`,
	)

	popped := odns.Pop()

	fmt.Printf("%s", popped)
	// Output: ldap:///uid=jimmy,ou=People,dc=example,dc=com
}

func ExampleTargetDistinguishedNames_Pop() {
	var odns TargetDistinguishedNames = TDNs()
	odns.Push(
		`uid=jesse,ou=People,dc=example,dc=com`,
		`uid=jimmy,ou=People,dc=example,dc=com`,
	)

	popped := odns.Pop()

	fmt.Printf("%s", popped)
	// Output: ldap:///uid=jimmy,ou=People,dc=example,dc=com
}

func ExampleBindDistinguishedNames_String() {
	var odns BindDistinguishedNames = UDNs()
	odns.Push(
		`uid=jesse,ou=People,dc=example,dc=com`,
		`uid=jimmy,ou=People,dc=example,dc=com`,
	)

	fmt.Printf("%s", odns)
	// Output: ldap:///uid=jesse,ou=People,dc=example,dc=com || ldap:///uid=jimmy,ou=People,dc=example,dc=com
}

func ExampleTargetDistinguishedNames_String() {
	var odns TargetDistinguishedNames = TDNs()
	odns.Push(
		`uid=jesse,ou=People,dc=example,dc=com`,
		`uid=jimmy,ou=People,dc=example,dc=com`,
	)

	fmt.Printf("%s", odns)
	// Output: ldap:///uid=jesse,ou=People,dc=example,dc=com || ldap:///uid=jimmy,ou=People,dc=example,dc=com
}

func ExampleBindDistinguishedNames_F() {
	var odns BindDistinguishedNames = UDNs()
	funk := odns.F()
	odns.Push(funk(`uid=jesse,ou=People,dc=example,dc=com`))

	fmt.Printf("%s", odns)
	// Output: ldap:///uid=jesse,ou=People,dc=example,dc=com
}

func ExampleTargetDistinguishedNames_F() {
	var odns TargetDistinguishedNames = TDNs()
	funk := odns.F()
	odns.Push(funk(`uid=jesse,ou=People,dc=example,dc=com`))

	fmt.Printf("%s", odns)
	// Output: ldap:///uid=jesse,ou=People,dc=example,dc=com
}

func ExampleBindDistinguishedNames_Len() {
	var odns BindDistinguishedNames = UDNs()
	funk := odns.F()
	odns.Push(funk(`uid=jesse,ou=People,dc=example,dc=com`))

	fmt.Printf("Length: %d", odns.Len())
	// Output: Length: 1
}

func ExampleTargetDistinguishedNames_Len() {
	var odns TargetDistinguishedNames = TDNs()
	funk := odns.F()
	odns.Push(funk(`uid=jesse,ou=People,dc=example,dc=com`))

	fmt.Printf("Length: %d", odns.Len())
	// Output: Length: 1
}

func ExampleBindDistinguishedNames_Eq() {
	var dn BindDistinguishedNames = UDNs(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("%s", dn.Eq())
	// Output: userdn = "ldap:///uid=jesse,ou=People,dc=example,dc=com"
}

func ExampleBindDistinguishedNames_Ne() {
	var dn BindDistinguishedNames = UDNs(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("%s", dn.Ne())
	// Output: userdn != "ldap:///uid=jesse,ou=People,dc=example,dc=com"
}

func ExampleTargetDistinguishedNames_Eq() {
	var dn TargetDistinguishedNames = TDNs(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("%s", dn.Eq())
	// Output: ( target = "ldap:///uid=jesse,ou=People,dc=example,dc=com" )
}

func ExampleTargetDistinguishedNames_Ne() {
	var dn TargetDistinguishedNames = TDNs(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("%s", dn.Ne())
	// Output: ( target != "ldap:///uid=jesse,ou=People,dc=example,dc=com" )
}

func ExampleBindDistinguishedNames_Keyword() {
	var dn BindDistinguishedNames = UDNs(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("Keyword: %s", dn.Keyword())
	// Output: Keyword: userdn
}

func ExampleTargetDistinguishedNames_Keyword() {
	var dn TargetDistinguishedNames = TDNs(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("Keyword: %s", dn.Keyword())
	// Output: Keyword: target
}

func ExampleTargetDistinguishedNames_IsZero() {
	var odns TargetDistinguishedNames
	fmt.Printf("Zero: %t", odns.IsZero())
	// Output: Zero: true
}

func ExampleBindDistinguishedNames_IsZero() {
	var odns BindDistinguishedNames
	fmt.Printf("Zero: %t", odns.IsZero())
	// Output: Zero: true
}

func ExampleBindDistinguishedNames_Valid() {
	var odns BindDistinguishedNames
	fmt.Printf("Valid: %t", odns.Valid() == nil)
	// Output: Valid: false
}

func ExampleTargetDistinguishedNames_Valid() {
	var odns TargetDistinguishedNames
	fmt.Printf("Valid: %t", odns.Valid() == nil)
	// Output: Valid: false
}

func ExampleUDNs() {
	udns := UDNs(
		`uid=jesse,ou=People,dc=example,dc=com`,
		`uid=courtney,ou=People,dc=example,dc=com`,
	)
	fmt.Printf("%s", udns)
	// Output: ldap:///uid=jesse,ou=People,dc=example,dc=com || ldap:///uid=courtney,ou=People,dc=example,dc=com
}

func ExampleRDNs() {
	rdns := RDNs(
		`cn=Default,ou=Profiles,dc=example,dc=com`,
		`cn=Administrative,ou=Profiles,dc=example,dc=com`,
	)
	fmt.Printf("%d", rdns.Len())
	// Output: 2
}

func ExampleGDNs() {
	gdns := GDNs(
		`uid=Executives,ou=Groups,dc=example,dc=com`,
		`uid=Executive Assistants,ou=Groups,dc=example,dc=com`,
	)
	fmt.Printf("%s", gdns)
	// Output: ldap:///uid=Executives,ou=Groups,dc=example,dc=com || ldap:///uid=Executive Assistants,ou=Groups,dc=example,dc=com
}

/*
This example demonstrates the call of a select slice
member from the receiver using the Index method.
*/
func ExampleBindDistinguishedNames_Index() {
	odns := UDNs(`uid=jesse,ou=People,dc=example,dc=com`, `uid=courtney,ou=People,dc=example,dc=com`)

	fmt.Printf("%s", odns.Index(0))
	// Output: ldap:///uid=jesse,ou=People,dc=example,dc=com
}

/*
This example demonstrates the call of a select slice
member from the receiver using the Index method.
*/
func ExampleTargetDistinguishedNames_Index() {
	odns := TFDNs(`uid=jesse,ou=People,dc=example,dc=com`, `uid=courtney,ou=People,dc=example,dc=com`)

	fmt.Printf("%s", odns.Index(0))
	// Output: ldap:///uid=jesse,ou=People,dc=example,dc=com
}

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

func ExampleTargetDistinguishedName_Valid() {
	var dn TargetDistinguishedName
	fmt.Printf("Valid: %t", dn.Valid() == nil)
	// Output: Valid: false
}

func ExampleBindDistinguishedName_Valid() {
	var dn BindDistinguishedName
	fmt.Printf("Valid: %t", dn.Valid() == nil)
	// Output: Valid: false
}

func ExampleBindDistinguishedName_IsZero() {
	var dn BindDistinguishedName
	fmt.Printf("Zero: %t", dn.IsZero())
	// Output: Zero: true
}

func ExampleTargetDistinguishedName_IsZero() {
	var dn TargetDistinguishedName
	fmt.Printf("Zero: %t", dn.IsZero())
	// Output: Zero: true
}

func ExampleBindDistinguishedName_Keyword() {
	var dn BindDistinguishedName = UDN(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("Keyword: %s", dn.Keyword())
	// Output: Keyword: userdn
}

func ExampleBindDistinguishedName_Kind() {
	var dn BindDistinguishedName = UDN(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("Kind: %s", dn.Kind())
	// Output: Kind: bind
}

func ExampleTargetDistinguishedName_Kind() {
	var dn TargetDistinguishedName = TFDN(`uid=*,ou=People,dc=example,dc=com`)
	fmt.Printf("Kind: %s", dn.Kind())
	// Output: Kind: target
}

func ExampleBindDistinguishedName_Len() {
	var dn BindDistinguishedName = UDN(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("Length %d", dn.Len())
	// Output: Length 1
}

func ExampleTargetDistinguishedName_Len() {
	var dn TargetDistinguishedName = TFDN(`uid=*,ou=People,dc=example,dc=com`)
	fmt.Printf("Length %d", dn.Len())
	// Output: Length 1
}

func ExampleBindDistinguishedName_String() {
	var dn BindDistinguishedName = UDN(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("%s", dn)
	// Output: ldap:///uid=jesse,ou=People,dc=example,dc=com
}

func ExampleTargetDistinguishedName_String() {
	var dn TargetDistinguishedName = TDN(`cn=Executives,ou=Groups,dc=example,dc=com`)
	fmt.Printf("%s", dn)
	// Output: ldap:///cn=Executives,ou=Groups,dc=example,dc=com
}

func ExampleBindDistinguishedName_Eq() {
	var dn BindDistinguishedName = UDN(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("%s", dn.Eq())
	// Output: userdn = "ldap:///uid=jesse,ou=People,dc=example,dc=com"
}

func ExampleBindDistinguishedName_Ne() {
	var dn BindDistinguishedName = UDN(`uid=courtney,ou=People,dc=example,dc=com`)
	fmt.Printf("%s", dn.Ne())
	// Output: userdn != "ldap:///uid=courtney,ou=People,dc=example,dc=com"
}

func ExampleTargetDistinguishedName_Eq() {
	var dn TargetDistinguishedName = TTDN(`cn=Distribution List,ou=Groups,dc=example,dc=com`)
	fmt.Printf("%s", dn.Eq())
	// Output: ( target_to = "ldap:///cn=Distribution List,ou=Groups,dc=example,dc=com" )
}

func ExampleTargetDistinguishedName_Ne() {
	var dn TargetDistinguishedName = TFDN(`uid=courtney,ou=People,dc=example,dc=com`)
	fmt.Printf("%s", dn.Ne())
	// Output: ( target_from != "ldap:///uid=courtney,ou=People,dc=example,dc=com" )
}

func ExampleTargetDistinguishedName_Keyword() {
	var dn TargetDistinguishedName = TFDN(`uid=*,ou=People,dc=example,dc=com`)
	fmt.Printf("Keyword: %s", dn.Keyword())
	// Output: Keyword: target_from
}

func ExampleBindDistinguishedName_BRM() {
	var dn BindDistinguishedName
	fmt.Printf("%d available comparison operator methods", dn.BRM().Len())
	// Output: 2 available comparison operator methods
}

func ExampleBindDistinguishedNames_BRM() {
	var dn BindDistinguishedNames
	fmt.Printf("%d available comparison operator methods", dn.BRM().Len())
	// Output: 2 available comparison operator methods
}

func ExampleTargetDistinguishedName_TRM() {
	var dn TargetDistinguishedName
	fmt.Printf("%d available comparison operator methods", dn.TRM().Len())
	// Output: 2 available comparison operator methods
}

func ExampleTargetDistinguishedNames_TRM() {
	var dn TargetDistinguishedNames
	fmt.Printf("%d available comparison operator methods", dn.TRM().Len())
	// Output: 2 available comparison operator methods
}

func ExampleTDN() {
	dn := TDN(`cn=Executives,ou=Groups,dc=example,dc=com`)
	fmt.Printf("%s", dn)
	// Output: ldap:///cn=Executives,ou=Groups,dc=example,dc=com
}

func ExampleTTDN() {
	dn := TTDN(`cn=Executives,ou=Groups,dc=example,dc=com`)
	fmt.Printf("%s", dn)
	// Output: ldap:///cn=Executives,ou=Groups,dc=example,dc=com
}

func ExampleTFDN() {
	dn := TFDN(`cn=Executives,ou=Groups,dc=example,dc=com`)
	fmt.Printf("%s", dn)
	// Output: ldap:///cn=Executives,ou=Groups,dc=example,dc=com
}

func ExampleUDN() {
	dn := UDN(`cn=Jesse Coretta,ou=People,dc=example,dc=com`)
	fmt.Printf("%s", dn)
	// Output: ldap:///cn=Jesse Coretta,ou=People,dc=example,dc=com
}

func ExampleGDN() {
	dn := GDN(`cn=Executives,ou=Groups,dc=example,dc=com`)
	fmt.Printf("%s", dn)
	// Output: ldap:///cn=Executives,ou=Groups,dc=example,dc=com
}

func ExampleRDN() {
	dn := RDN(`cn=Role Profile,ou=Roles,dc=example,dc=com`)
	fmt.Printf("%s", dn)
	// Output: ldap:///cn=Role Profile,ou=Roles,dc=example,dc=com
}

func ExampleBindDistinguishedName_Set() {
	var dn BindDistinguishedName
	dn.Set(`cn=Role Profile,ou=Roles,dc=example,dc=com`, BindRDN)
	fmt.Printf("%s", dn.Eq())
	// Output: roledn = "ldap:///cn=Role Profile,ou=Roles,dc=example,dc=com"
}

func ExampleTargetDistinguishedName_Set() {
	var dn TargetDistinguishedName
	dn.Set(`cn=*,($attr.ou),dc=example,dc=com`, Target)
	fmt.Printf("%s", dn.Ne())
	// Output: ( target != "ldap:///cn=*,($attr.ou),dc=example,dc=com" )
}

func ExampleBindDistinguishedNames_ID() {
	var dn BindDistinguishedNames
	fmt.Printf("%s", dn.ID())
	// Output: bind
}

func ExampleTargetDistinguishedNames_ID() {
	var dn TargetDistinguishedNames
	fmt.Printf("%s", dn.ID())
	// Output: target
}

func ExampleTDNs() {
	tdns := TDNs(
		`uid=jesse,ou=People,dc=example,dc=com`,
		`uid=courtney,ou=People,dc=example,dc=com`,
	)
	fmt.Printf("%s contains %d DNs", tdns.Keyword(), tdns.Len())
	// Output: target contains 2 DNs
}

func ExampleTTDNs() {
	tdns := TTDNs(
		`uid=jesse,ou=People,dc=example,dc=com`,
		`uid=courtney,ou=People,dc=example,dc=com`,
	)
	fmt.Printf("%s contains %d DNs", tdns.Keyword(), tdns.Len())
	// Output: target_to contains 2 DNs
}

func ExampleTFDNs() {
	tdns := TFDNs(
		`uid=jesse,ou=People,dc=example,dc=com`,
		`uid=courtney,ou=People,dc=example,dc=com`,
	)
	fmt.Printf("%s contains %d DNs", tdns.Keyword(), tdns.Len())
	// Output: target_from contains 2 DNs
}

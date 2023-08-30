package aci

import (
	"fmt"
	"testing"
)

func TestObjectIdentifier_codecov(t *testing.T) {
	for kw, fn := range map[TargetKeyword]func(...any) ObjectIdentifier{
		TargetCtrl:  Ctrl,
		TargetExtOp: ExtOp,
	} {
		for _, oid := range []string{
			`1.3.6.1.4.1.56521.999.83`,
			`1.3.6.1.4.1.56521.999.84`,
			`1.3.6.1.4.1.56521.999.85`,
			`1.3.6.1.4.1.56521.999.86`,
			`1.3.6.1.4.1.56521.999.87`,
			`1.3.6.1.4.1.56521.999.88`,
			`1.3.6.1.4.1.56521.999.89`,
			`1.3.6.1.4.1.56521.999.90`,
			`1.3.6.1.4.1.56521.999.91`,
		} {
			var O ObjectIdentifier
			if err := O.Valid(); err == nil {
				t.Errorf("%s failed: invalid %T returned no validity error",
					t.Name(), O)
			}

			if !O.IsZero() {
				t.Errorf("%s failed: unexpectted non-nil %T instance",
					t.Name(), O)
			}

			if O.String() != badDotNot {
				t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
					t.Name(), badDotNot, O)
			}

			// process OID
			O = fn(oid)

			if err := O.Valid(); err != nil {
				t.Errorf("%s failed: %v", t.Name(), err)
			}

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

func TestObjectIdentifier(t *testing.T) {
	got := Ctrl(`1.3.6.1.4.1.56521.999.5`)
	want := `1.3.6.1.4.1.56521.999.5`

	if want != got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

/*
This example demonstrates the creation of a single LDAP Extended Operation
Object Identifier, intended for use in the creation of extop TargetRule
expressions.
*/
func ExampleExtOp() {
	o := ExtOp(`1.3.6.1.4.1.56521.999.5`)
	fmt.Printf("%s", o)
	// Output: 1.3.6.1.4.1.56521.999.5
}

/*
This example demonstrates the creation of a single LDAP Control Object Identifier,
intended for use in the creation of targetcontrol TargetRule expressions.
*/
func ExampleCtrl() {
	o := Ctrl(`1.3.6.1.4.1.56521.999.5`)
	fmt.Printf("%s", o)
	// Output: 1.3.6.1.4.1.56521.999.5
}

/*
This example demonstrates the creation of a multi-valued targetcontrol (LDAP Control)
Target Rule expression.
*/
func ExampleCtrls() {
	// note: these are phony OIDs
	o1 := Ctrl(`1.3.6.1.4.1.56521.999.5`)
	o2 := Ctrl(`1.3.6.1.4.1.56521.999.7`)

	// Initialize the stack (Ctrls) and
	// immediately push o1 and o2.
	ctrls := Ctrls().Push(o1, o2)

	fmt.Printf("%s", ctrls.Eq())
	// Output: ( targetcontrol = "1.3.6.1.4.1.56521.999.5 || 1.3.6.1.4.1.56521.999.7" )
}

/*
This example demonstrates the creation of a multi-valued extop (LDAP Extended
Operation) Target Rule expression.
*/
func ExampleExtOps() {
	// note: these are phony OIDs
	o1 := ExtOp(`1.3.6.1.4.1.56521.999.5`)
	o2 := ExtOp(`1.3.6.1.4.1.56521.999.7`)

	// Initialize the stack (Ctrls) and
	// immediately push o1 and o2.
	exop := ExtOps().Push(o1, o2)

	fmt.Printf("%s", exop.Eq())
	// Output: ( extop = "1.3.6.1.4.1.56521.999.5 || 1.3.6.1.4.1.56521.999.7" )
}

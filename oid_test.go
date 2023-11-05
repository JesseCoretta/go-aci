package aci

import (
	"fmt"
	"testing"
)

func TestObjectIdentifiers_codecov(t *testing.T) {

	_ = isDotNot(``)
	_ = isDotNot(`this is not dot not`)

	for keyword, Oidsfn := range map[TargetKeyword]func(...any) ObjectIdentifiers{
		TargetCtrl:  Ctrls,
		TargetExtOp: ExtOps,
	} {
		var Oids ObjectIdentifiers
		_ = Oids.Ne()
		_ = Oids.Eq()

		_ = objectIdentifiersPushPolicy(Oids, ``, keyword)
		_ = objectIdentifiersPushPolicy(Oids, nil, keyword)
		_ = objectIdentifiersPushPolicy(Oids, float64(1), keyword)
		_ = objectIdentifiersPushPolicy(Oids, BindUDN, keyword)
		_ = objectIdentifiersPushPolicy(Oids, keyword, keyword)
		_ = objectIdentifiersPushPolicy(Oids, `hello`, keyword)

		_ = Oids.Len()
		Oids.reset()
		Oids.resetKeyword(keyword)
		Oids.resetKeyword(keyword.String())
		Oids.Push(keyword)
		_ = Oids.Keyword()
		_ = Oids.Kind()
		_ = Oids.Push()
		_ = Oids.Push(nil)
		_ = Oids.Push(``)
		_ = Oids.Push('a')
		_ = Oids.Ne()
		_ = Oids.Eq()
		_ = Oids.Valid()
		_ = Oids.setQuoteStyle(0)
		_ = Oids.setQuoteStyle(1)
		Oids.isObjectIdentifierContext()

		Oids = Oidsfn() // init
		_ = Oids.Len()
		_ = objectIdentifiersPushPolicy(Oids, ``, keyword)
		_ = objectIdentifiersPushPolicy(Oids, nil, keyword)
		_ = objectIdentifiersPushPolicy(Oids, BindUDN, keyword)
		_ = objectIdentifiersPushPolicy(Oids, keyword, keyword)
		_ = objectIdentifiersPushPolicy(Oids, float64(1), keyword)
		_ = objectIdentifiersPushPolicy(Oids, `hello`, keyword)

		Oids.reset()
		Oids.resetKeyword(keyword)
		Oids.resetKeyword(keyword.String())
		Oids.Push(keyword)
		_ = Oids.Keyword()
		_ = Oids.Kind()
		_ = Oids.Push()
		_ = Oids.Push(nil)
		_ = Oids.Push(``)
		_ = Oids.Push('a')
		_ = Oids.Ne()
		_ = Oids.Eq()
		_ = Oids.Valid()
		_ = Oids.setQuoteStyle(0)
		_ = Oids.setQuoteStyle(1)
		Oids.isObjectIdentifierContext()

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
			var (
				Ol   int                           = Oids.Len()
				ofn  func(...any) ObjectIdentifier = Oids.F()
				Oid  ObjectIdentifier
				octx ObjectIdentifierContext
			)

			Oid.isObjectIdentifierContext()
			_ = Oid.Len()
			_ = Oid.Kind()
			_ = Oid.Valid()
			_ = Oid.Keyword()

			if err := testEmptyOidContext(t, keyword, Oid, Oids, Ol); err != nil {
				t.Errorf(err.Error())
				return
			}

			// process OID
			Oid = ofn(oid)
			Oid = ofn(Oid)
			badOid := ofn(5465734.3, '𝝅', '🤮', '🤮', '🤮', '🤮', '🤮', '🤮', '🤮', '☜', 'Ҩ', 'ↂ', '⼼', 'ↂ', 'Ҩ', '☞')
			_ = badOid.String()
			_ = badOid.Valid()
			_ = badOid.Eq()
			Ol = Oids.Len()
			_ = Oid.Len()
			_ = Oid.Kind()
			_ = Oid.Keyword()
			_ = Oid.Ne()
			_ = Oid.Eq()
			_ = Oid.Valid()

			if Oids.Push(Oid); !Oids.Contains(oid) {
				t.Errorf("%s [%s] multival failed: valid %T[%s] instance (%s) not pushed into %T[%s; len:%d]",
					t.Name(), keyword, Oid, Oid.Keyword(), Oid, Oids, Oids.Keyword(), Ol)
				return
			}
			_ = Oids.Contains('𝝅')
			_ = Oids.Contains(43785)
			_ = Oids.Contains(nil)
			_ = Oids.Contains(``)

			popped := Oids.Pop()
			Oids.Push(popped)
			Oids.Push(popped.String())
			Oids.Push(3444444.445)
			Oids.Push()
			Oids.Push(``) // crashes go-objectid
			Oids.Push('𝝅')
			_ = Oids.Keyword()
			_ = Oids.Kind()

			_ = Oids.setQuoteStyle(0)
			_ = Oids.setQuoteStyle(1)

			for sop, trmn := range []func() TargetRuleMethods{
				Oid.TRM,
				Oids.TRM,
			} {
				octx = testMakeOidContext(sop, Oid, Oids)
				trm := trmn()
				for i := 0; i < trm.Len(); i++ {
					cop, meth := trm.Index(i + 1)
					if meth == nil {
						t.Errorf("%s [%s] multival failed: expected %s method (%T), got nil",
							t.Name(), keyword, cop.Context(), meth)
						return
					}

					wcop := sprintf("( %s %s %q )", octx.Keyword(), cop, octx.String())
					if T := meth(); T.String() != wcop {
						err := unexpectedStringResult(octx.String(), wcop, T.String())
						t.Errorf("%s [%s] multival failed [%s rule]: %v",
							t.Name(), keyword, octx.Keyword(), err)
						return
					}
				}
			}
		}
	}
}

func testMakeOidContext(phase int, oid ObjectIdentifier, oids ObjectIdentifiers) (octx ObjectIdentifierContext) {
	if phase == 0 {
		octx = oid
		return
	}

	octx = oids
	return
}

func testEmptyOidContext(t *testing.T, kw Keyword, oid ObjectIdentifier, oids ObjectIdentifiers, ol int) (err error) {
	_ = oid.Keyword()
	_ = oid.Kind()
	_ = oid.Valid()
	_ = oids.Keyword()
	_ = oids.Kind()
	_ = oids.Valid()

	err = oid.Valid()
	if err != nil {
		if err.Error() != `aci.ObjectIdentifier instance is nil` {
			err = errorf("%s [%s] multival failed: invalid %T returned no validity error",
				t.Name(), kw, oid)
		} else {
			err = nil
		}
	} else {
		err = errorf("%s [%s] multival failed: invalid %T returned no validity error",
			t.Name(), kw, oid)
	}

	if oid.String() != badDotNot {
		err = errorf("%s [%s] multival failed: unexpected string result; want '%s', got '%s'",
			t.Name(), kw, badDotNot, oid)
	}

	if oids.Push(oid); oids.Len() > ol {
		err = errorf("%s [%s] multival failed (len): invalid %T (%s) pushed into %T without error",
			t.Name(), kw, oid, oid, oids)
	}

	if oids.Contains(oid) {
		err = errorf("%s [%s] multival failed (contains): invalid %T instance pushed into %T without error",
			t.Name(), kw, oid, oids)
	}

	return
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

func ExampleObjectIdentifier_Compare() {
	o1 := Ctrl(`1.3.6.1.4.1.56521.999.5`)
	o2 := ExtOp(`1.3.6.1.4.1.56521.999.5`)
	fmt.Printf("%t", o1.Compare(o2))
	// Output: true
}

func ExampleObjectIdentifiers_Compare() {
	o1 := Ctrls(`1.3.6.1.4.1.56521.999.5`, `1.3.6.1.4.1.56521.999.6`)
	o2 := Ctrls(`1.3.6.1.4.1.56521.999.7`, `1.3.6.1.4.1.56521.999.6`)
	fmt.Printf("%t", o1.Compare(o2))
	// Output: false
}

/*
This example demonstrates the creation of a multi-valued targetcontrol (LDAP Control)
[TargetRule] expression.
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
Operation) [TargetRule] expression.
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

/*
This example demonstrates the manual creation of an ObjectIdentifier instance using the
string representation of an ASN.1 Object Identifier in dot notation, and a valid keyword
context for a TargetRule.
*/
func ExampleOID() {
	o1 := `1.3.6.1.4.1.56521.999.5` // note: phony OID

	o := OID(o1, TargetExtOp)
	//o := OID(o1, TargetCtrl)	// alternative

	fmt.Printf("OID:%s, Type:%s", o, o.Keyword())
	// Output: OID:1.3.6.1.4.1.56521.999.5, Type:extop
}

/*
This example demonstrates the creation of a multi-valued extop (LDAP Extended
Operation) equality TargetRule expression. Push is used to submit the instances
of ObjectIdentifier to the stack.
*/
func ExampleObjectIdentifiers_Push() {
	// note: these are phony OIDs
	o1 := ExtOp(`1.3.6.1.4.1.56521.999.5`)
	o2 := ExtOp(`1.3.6.1.4.1.56521.999.7`)

	// Initialize the stack (Ctrls) and
	// immediately push o1 and o2.
	exop := ExtOps().Push(o1, o2)

	fmt.Printf("%s", exop.Eq())
	// Output: ( extop = "1.3.6.1.4.1.56521.999.5 || 1.3.6.1.4.1.56521.999.7" )
}

/*
This is an example of a LIFO stack slice removal using the Pop method.
*/
func ExampleObjectIdentifiers_Pop() {
	// note: these are phony OIDs
	o1 := ExtOp(`1.3.6.1.4.1.56521.999.5`)
	o2 := ExtOp(`1.3.6.1.4.1.56521.999.7`)
	o3 := ExtOp(`1.3.6.1.4.1.56521.999.9`)

	// Initialize the stack (Ctrls) and
	// immediately push o1 and o2.
	exop := ExtOps().Push(o1, o2, o3)
	popped := exop.Pop()

	fmt.Printf("Removed %T (%s), stack length now %d", popped, popped, exop.Len())
	// Output: Removed aci.ObjectIdentifier (1.3.6.1.4.1.56521.999.9), stack length now 2
}

/*
This example demonstrates the creation of a multi-valued extop (LDAP Extended
Operation) equality TargetRule expression.
*/
func ExampleObjectIdentifiers_Eq() {
	// note: these are phony OIDs
	o1 := ExtOp(`1.3.6.1.4.1.56521.999.5`)
	o2 := ExtOp(`1.3.6.1.4.1.56521.999.7`)

	// Initialize the stack (Ctrls) and
	// immediately push o1 and o2.
	exop := ExtOps().Push(o1, o2)

	fmt.Printf("%s", exop.Eq())
	// Output: ( extop = "1.3.6.1.4.1.56521.999.5 || 1.3.6.1.4.1.56521.999.7" )
}

/*
This example demonstrates the creation of a multi-valued extop (LDAP Extended
Operation) equality TargetRule expression.
*/
func ExampleObjectIdentifier_Eq() {
	fmt.Printf("%s", ExtOp(`1.3.6.1.4.1.56521.999.5`).Eq())
	// Output: ( extop = "1.3.6.1.4.1.56521.999.5" )
}

/*
This example demonstrates the creation of a multi-valued targetcontrol (LDAP Controls)
negated equality TargetRule expression.
*/
func ExampleObjectIdentifiers_Ne() {
	// note: these are phony OIDs
	o1 := Ctrl(`1.3.6.1.4.1.56521.999.5`)
	o2 := Ctrl(`1.3.6.1.4.1.56521.999.7`)

	// Initialize the stack (Ctrls) and
	// immediately push o1 and o2.
	ctrls := Ctrls().Push(o1, o2)

	fmt.Printf("%s", ctrls.Ne())
	// Output: ( targetcontrol != "1.3.6.1.4.1.56521.999.5 || 1.3.6.1.4.1.56521.999.7" )
}

/*
This example demonstrates the creation of a multi-valued extop (LDAP Extended
Operation) negated equality TargetRule expression.
*/
func ExampleObjectIdentifier_Ne() {
	fmt.Printf("%s", ExtOp(`1.3.6.1.4.1.56521.999.5`).Ne())
	// Output: ( extop != "1.3.6.1.4.1.56521.999.5" )
}

/*
This example demonstrates the string representation of the receiver.
*/
func ExampleObjectIdentifier_String() {
	fmt.Printf("%s", ExtOp(`1.3.6.1.4.1.56521.999.5`))
	// Output: 1.3.6.1.4.1.56521.999.5
}

/*
This example demonstrates the (mostly) useless nature of the Len method,
which only exists to satisfy Go's interface signature requirements.
*/
func ExampleObjectIdentifier_Len() {
	fmt.Printf("%d", ExtOp(`1.3.6.1.4.1.56521.999.5`).Len())
	// Output: 1
}

/*
This example demonstrates use of the Len method to return the number
of slices present within the receiver as an integer.
*/
func ExampleObjectIdentifiers_Len() {
	ctrls := Ctrls().Push(
		// note: these are phony OIDs
		Ctrl(`1.3.6.1.4.1.56521.999.5`),
		Ctrl(`1.3.6.1.4.1.56521.999.7`),
	)
	fmt.Printf("%d", ctrls.Len())
	// Output: 2
}

/*
This example demonstrates use of the Keyword method to obtain the
current Keyword context from the receiver.
*/
func ExampleObjectIdentifiers_Keyword() {
	ctrls := Ctrls().Push(
		// note: these are phony OIDs
		Ctrl(`1.3.6.1.4.1.56521.999.5`),
		Ctrl(`1.3.6.1.4.1.56521.999.7`),
	)
	fmt.Printf("%s", ctrls.Keyword())
	// Output: targetcontrol
}

/*
This example demonstrates use of the Kind method to obtain the
string form of the current Keyword context from the receiver.
*/
func ExampleObjectIdentifiers_Kind() {
	ctrls := Ctrls().Push(
		// note: these are phony OIDs
		Ctrl(`1.3.6.1.4.1.56521.999.5`),
		Ctrl(`1.3.6.1.4.1.56521.999.7`),
	)
	fmt.Printf("%s", ctrls.Kind())
	// Output: targetcontrol
}

/*
This example demonstrates use of the Keyword method to obtain the
current Keyword context from the receiver.
*/
func ExampleObjectIdentifier_Keyword() {
	fmt.Printf("%s", ExtOp(`1.3.6.1.4.1.56521.999.5`).Keyword())
	// Output: extop
}

/*
This example demonstrates use of the Kind method to obtain the
string form of the current Keyword context from the receiver.
*/
func ExampleObjectIdentifier_Kind() {
	fmt.Printf("%s", ExtOp(`1.3.6.1.4.1.56521.999.5`).Kind())
	// Output: extop
}

/*
This example demonstrates the string representation of the receiver.
*/
func ExampleObjectIdentifiers_String() {
	// Initialize the stack (Ctrls) and
	// immediately push into it.
	ctrls := Ctrls().Push(
		// note: these are phony OIDs
		Ctrl(`1.3.6.1.4.1.56521.999.5`),
		Ctrl(`1.3.6.1.4.1.56521.999.7`),
	)

	fmt.Printf("%s", ctrls)
	// Output: 1.3.6.1.4.1.56521.999.5 || 1.3.6.1.4.1.56521.999.7
}

/*
This example demonstrates the use of the IsZero method upon a nil receiver.
*/
func ExampleObjectIdentifier_IsZero() {
	var oid ObjectIdentifier
	fmt.Printf("Zero: %t", oid.IsZero())
	// Output: Zero: true
}

/*
This example demonstrates the use of the IsZero method upon a nil receiver.
*/
func ExampleObjectIdentifiers_IsZero() {
	var oids ObjectIdentifiers
	fmt.Printf("Zero: %t", oids.IsZero())
	// Output: Zero: true
}

/*
This example demonstrates the use of the ID method.
*/
func ExampleObjectIdentifiers_ID() {
	var oids ObjectIdentifiers = ExtOps() // must be initialized, as there are two (2) types of OIDs here
	fmt.Printf("ID: %s", oids.ID())
	// Output: ID: target
}

/*
This example demonstrates the use of the Valid method upon a nil receiver.
*/
func ExampleObjectIdentifier_Valid() {
	var oid ObjectIdentifier
	fmt.Printf("Valid: %t", oid.Valid() == nil)
	// Output: Valid: false
}

/*
This example demonstrates the use of the Valid method upon a nil receiver.
*/
func ExampleObjectIdentifiers_Valid() {
	var oids ObjectIdentifiers
	fmt.Printf("Valid: %t", oids.Valid() == nil)
	// Output: Valid: false
}

/*
This example demonstrates the population of an object identifier stack
and a subsequent presence check of one of its members.
*/
func ExampleObjectIdentifiers_Contains() {
	ctrls := Ctrls(
		`1.3.6.1.4.1.56521.999.5`,
		`1.3.6.1.4.1.56521.999.6`,
		`1.3.6.1.4.1.56521.999.7`,
	)

	fmt.Printf("Contains OID: %t", ctrls.Contains(`1.3.6.1.4.1.56521.999.5`))
	// Output: Contains OID: true
}

/*
This example demonstrates the use of the Index method to obtain a single
slice and print its Keyword value.
*/
func ExampleObjectIdentifiers_Index() {
	ctrls := Ctrls(
		`1.3.6.1.4.1.56521.999.5`,
		`1.3.6.1.4.1.56521.999.6`,
		`1.3.6.1.4.1.56521.999.7`,
	)

	fmt.Printf("Slice keyword: %s", ctrls.Index(1).Keyword())
	// Output: Slice keyword: targetcontrol
}

/*
This example demonstrates the use of the TRM method to obtain a list of available
comparison operator identifiers and methods, and a subsequent call of Contains
to determine whether Greater Than (Gt) is among them.
*/
func ExampleObjectIdentifiers_TRM() {
	var oids ObjectIdentifiers
	fmt.Printf("Allows greater-than: %t", oids.TRM().Contains(Gt))
	// Output: Allows greater-than: false
}

/*
This example demonstrates the use of the F method to obtain the appropriate
package level function to be used to craft additional slices for push into
the receiver.
*/
func ExampleObjectIdentifiers_F() {
	var oids ObjectIdentifiers = Ctrls()
	funk := oids.F()
	newval := funk(`1.3.6.1.4.1.56521.999.5`)
	fmt.Printf("New %T: %s", newval, newval)
	// Output: New aci.ObjectIdentifier: 1.3.6.1.4.1.56521.999.5
}

/*
This example demonstrates the use of the TRM method to obtain a list of available
comparison operator identifiers and methods, and a subsequent call of Contains
to determine whether Greater Than (Gt) is among them.
*/
func ExampleObjectIdentifier_TRM() {
	var oid ObjectIdentifier
	fmt.Printf("Allows greater-than: %t", oid.TRM().Contains(Gt))
	// Output: Allows greater-than: false
}

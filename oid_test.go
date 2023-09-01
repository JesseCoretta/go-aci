package aci

import (
	"fmt"
	"testing"
)

func TestObjectIdentifiers_codecov(t *testing.T) {

	for keyword, Oidsfn := range map[TargetKeyword]func(...any) ObjectIdentifiers{
		TargetCtrl:  Ctrls,
		TargetExtOp: ExtOps,
	} {
		var Oids ObjectIdentifiers = Oidsfn()

		_ = Oids.Len()
		Oids.reset()
		Oids.resetKeyword(keyword)
		Oids.resetKeyword(keyword.String())
		Oids.Push(keyword)
		_ = Oids.Keyword()
		_ = Oids.Kind()
		_ = Oids.Ne()
		_ = Oids.Eq()
		_ = Oids.Valid()
		_ = Oids.setQuoteStyle(1)
		_ = Oids.setQuoteStyle(0)
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

			_ = Oids.setQuoteStyle(1)
			_ = Oids.setQuoteStyle(0)

			for sop, trfn := range []func() TargetRuleMethods{
				Oid.TRF,
				Oids.TRF,
			} {
				octx = testMakeOidContext(sop, Oid, Oids)
				trf := trfn()
				for i := 0; i < trf.Len(); i++ {
					cop, meth := trf.Index(i + 1)
					if meth == nil {
						t.Errorf("%s [%s] multival failed: expected %s method (%T), got nil",
							t.Name(), keyword, cop.Context(), meth)
					}

					wcop := sprintf("( %s %s %q )", octx.Keyword(), cop, octx.String())
					if T := meth(); T.String() != wcop {
						err := unexpectedStringResult(octx.String(), wcop, T.String())
						t.Errorf("%s [%s] multival failed [%s rule]: %v",
							t.Name(), keyword, octx.Keyword(), err)
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
Operation) TargetRule expression.
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

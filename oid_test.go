package aci

import (
	"fmt"
	"testing"
)

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

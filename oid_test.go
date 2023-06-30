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

func ExampleExtOp() {
	o := ExtOp(`1.3.6.1.4.1.56521.999.5`)
	fmt.Printf("%s", o)
	// Output: 1.3.6.1.4.1.56521.999.5
}

func ExampleCtrl() {
	o := Ctrl(`1.3.6.1.4.1.56521.999.5`)
	fmt.Printf("%s", o)
	// Output: 1.3.6.1.4.1.56521.999.5
}


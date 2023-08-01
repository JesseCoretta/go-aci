package aci

import (
	"fmt"
	"testing"
)

func TestSecurityStrengthFactor_Set_stringTermMaximum(t *testing.T) {
	var s SecurityStrengthFactor
	s.Set(`full`)
	want := `256`
	got := s.String()
	if want != got {
		t.Errorf("%s failed; want '%s', got '%s'", t.Name(), want, got)
	}
}

func TestSecurityStrengthFactor_Set_stringTermNone(t *testing.T) {
	var s SecurityStrengthFactor
	s.Set(`noNe`) // test case insignificance
	want := `0`
	if want != s.String() {
		t.Errorf("%s failed; want '%s', got '%s'", t.Name(), want, s)
	}
}

func TestSecurityStrengthFactor_noSet(t *testing.T) {
	// same test as stringTermNone, except don'
	// actually set any value ...
	var x SecurityStrengthFactor
	want := `0`
	if want != x.String() {
		t.Errorf("%s failed; want '%s', got '%s'", t.Name(), want, x)
	}
}

func TestSecurityStrengthFactor_Set_noInit(t *testing.T) {
	var s SecurityStrengthFactor
	s.Set(171)
	want := `171`
	if want != s.String() {
		t.Errorf("%s failed; want '%s', got '%s'", t.Name(), want, s)
	}
}

func TestSecurityStrengthFactor_Set_fromInit(t *testing.T) {
	var s SecurityStrengthFactor
	s.Set(71)
	want := `71`
	if want != s.String() {
		t.Errorf("%s failed; want '%s', got '%s'", t.Name(), want, s)
	}
}

func ExampleSecurityStrengthFactor_Set_byWordNoFactor() {
	var s SecurityStrengthFactor
	s.Set(`noNe`) // case is not significant
	fmt.Printf("%s", s)
	// Output: 0
}

func ExampleSecurityStrengthFactor_Set_byWordMaxFactor() {
	var s SecurityStrengthFactor
	s.Set(`FULL`) // case is not significant
	//s.Set(`max`) // alternative term
	fmt.Printf("%s", s)
	// Output: 256
}

func ExampleSecurityStrengthFactor_Set_byNumber() {
	var s SecurityStrengthFactor
	s.Set(128)
	fmt.Printf("%s\n", s)
	// Output: 128
}

func ExampleSecurityStrengthFactor_Eq() {
	var s SecurityStrengthFactor
	fmt.Printf("%s", s.Set(128).Eq().Paren())
	// Output: (ssf = "128")
}

func ExampleSSF() {
	// convenient alternative to "var X SecurityStrengthFactor, X.Set(...) ..."
	fmt.Printf("%s", SSF(128))
	// Output: 128
}

func ExampleSSF_setLater() {
	s := SSF() // this is functionally the same ...
	// var s SecurityStrengthFactor // ... as this.

	// ... later in your code ...

	fmt.Printf("%s", s.Set(127))
	// Output: 127
}

func TestAnonymous_eqne(t *testing.T) {
	want := `authmethod != "none"`
	got := Anonymous.Ne()
	if want != got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}

	// reset
	want = `authmethod = "none"`
	got = Anonymous.Eq()
	if want != got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

func ExampleAuthMethod_Ne() {
	fmt.Printf("%s", Anonymous.Ne())
	// Output: authmethod != "none"
}

func ExampleAuthMethod_Eq() {
	fmt.Printf("%s", SASL.Eq())
	// Output: authmethod = "SASL"
}

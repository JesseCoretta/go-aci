package aci

import (
	"fmt"
	"testing"
)

func TestSecurityStrengthFactor(t *testing.T) {
	var (
		factor SecurityStrengthFactor
		typ    string = BindSSF.String()
		err    error
	)

	for i := 0; i < 257; i++ {
		want := itoa(i) // what we expect (string representation)

		got := factor.Set(i) // set by iterated integer
		if want != got.String() {
			err = unexpectedStringResult(typ, want, got.String())
			t.Errorf("%s failed [%s int]: %v",
				t.Name(), typ, err)
			return
		}

		// reset using string representation of iterated integer
		if got = factor.Set(want); want != got.String() {
			err = unexpectedStringResult(typ, want, got.String())
			t.Errorf("%s failed [%s str]: %v",
				t.Name(), typ, err)
			return
		}

		brm := factor.BRM()
		for c := 0; c < brm.Len(); c++ {
			cop, meth := brm.Index(c + 1)
			wcop := sprintf("%s %s %q", got.Keyword(), cop, got)

			// create bindrule B using comparison
			// operator (cop).
			if B := meth(); B.String() != wcop {
				err = unexpectedStringResult(typ, wcop, B.String())
			}

			if err != nil {
				t.Errorf("%s failed [%s rule]: %v", t.Name(), typ, err)
				return
			}
		}
		factor.clear() // codecov

	}

	// try to set our factor using special keywords
	// this package understands ...
	for word, value := range map[string]string{
		`mAx`:  `256`,
		`full`: `256`,
		`nOnE`: `0`,
		`OFF`:  `0`,
		`fart`: `0`,
	} {
		if got := factor.Set(word); got.String() != value {
			err = unexpectedStringResult(typ, value, got.String())
			t.Errorf("%s failed [factor word '%s']: %v", t.Name(), word, err)
			return
		}
	}
}

func TestAuthenticationMethod(t *testing.T) {
	// codecov
	_ = noAuth.Eq()
	_ = noAuth.Ne()

	for idx, auth := range authMap {
		if matchAuthenticationMethod(idx) == noAuth {
			t.Errorf("%s failed: unable to match auth method by index (%d)",
				t.Name(), idx)
			return
		} else if matchAuthenticationMethod(auth.String()) == noAuth {
			t.Errorf("%s failed: unable to match auth method by string (%s)",
				t.Name(), auth.String())
			return
		}
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
	// Output: ( ssf = "128" )
}

func ExampleSecurityStrengthFactor_Ne() {
	var s SecurityStrengthFactor
	fmt.Printf("%s", s.Set(128).Ne().Paren())
	// Output: ( ssf != "128" )
}

func ExampleSecurityStrengthFactor_Lt() {
	var s SecurityStrengthFactor
	fmt.Printf("%s", s.Set(128).Lt())
	// Output: ssf < "128"
}

func ExampleSecurityStrengthFactor_Le() {
	var s SecurityStrengthFactor
	fmt.Printf("%s", s.Set(128).Le().Paren())
	// Output: ( ssf <= "128" )
}

func ExampleSecurityStrengthFactor_Gt() {
	var s SecurityStrengthFactor
	fmt.Printf("%s", s.Set(128).Gt().Paren())
	// Output: ( ssf > "128" )
}

func ExampleSecurityStrengthFactor_Ge() {
	var s SecurityStrengthFactor
	fmt.Printf("%s", s.Set(128).Ge().Paren())
	// Output: ( ssf >= "128" )
}

func ExampleSecurityStrengthFactor_String() {
	var s SecurityStrengthFactor = SSF(128)
	fmt.Printf("%s", s)
	// Output: 128
}

func ExampleSecurityStrengthFactor_Valid() {
	var s SecurityStrengthFactor
	fmt.Printf("Valid: %t", s.Valid() == nil) // zero IS valid, technically speaking!
	// Output: Valid: true
}

func ExampleSecurityStrengthFactor_IsZero() {
	var s SecurityStrengthFactor
	fmt.Printf("Zero: %t", s.IsZero())
	// Output: Zero: true
}

func ExampleSecurityStrengthFactor_Keyword() {
	var s SecurityStrengthFactor
	fmt.Printf("Keyword: %s", s.Keyword())
	// Output: Keyword: ssf
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

func ExampleAuthenticationMethod_BRM() {
	meths := Anonymous.BRM()
	fmt.Printf("%d available aci.BindRuleMethod instances", meths.Len())
	// Output: 2 available aci.BindRuleMethod instances
}

func ExampleAuthenticationMethod_Ne() {
	fmt.Printf("%s", Anonymous.Ne())
	// Output: authmethod != "NONE"
}

func ExampleAuthenticationMethod_Eq() {
	fmt.Printf("%s", SASL.Eq())
	// Output: authmethod = "SASL"
}

func ExampleSecurityStrengthFactor_BRM() {
	var factor SecurityStrengthFactor = SSF(128)
	meths := factor.BRM()
	fmt.Printf("%d available aci.BindRuleMethod instances", meths.Len())
	// Output: 6 available aci.BindRuleMethod instances
}

func ExampleAuthenticationMethod_String() {
	fmt.Printf("%s", EXTERNAL)
	// Output: SASL EXTERNAL
}

/*
This example demonstrates the SHA-1 hash comparison between two (2)
AuthenticationMethod instances using the Compare method.
*/
func ExampleAuthenticationMethod_Compare() {
	fmt.Printf("%t", Anonymous.Compare(EXTERNAL))
	// Output: false
}

/*
This example demonstrates the SHA-1 hash comparison between two (2)
SecurityStrengthFactor instances using the Compare method.
*/
func ExampleSecurityStrengthFactor_Compare() {
	ssf1 := SSF(`101`)
	ssf2 := SSF(101)

	fmt.Printf("Hashes are equal: %t", ssf1.Compare(ssf2))
	// Output: Hashes are equal: true
}

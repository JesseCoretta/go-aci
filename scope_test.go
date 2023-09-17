package aci

import (
	"fmt"
	"testing"
)

func TestScope(t *testing.T) {
	for idx, raw := range []string{
		`baSe`,
		`oNe`,
		`sub`,
	} {
		if sc := Scope(raw); sc == noScope {
			t.Errorf("%s failed; failed to parse scope name '%s'",
				t.Name(), raw)
			return
		}
		if sc := Scope(idx); sc == noScope {
			t.Errorf("%s failed; failed to parse scope index '%d'",
				t.Name(), idx)
			return
		}
	}
}

func TestScope_targetRules(t *testing.T) {
	for idx, raw := range []string{
		`baSe`,
		`oNelevEL`,
		`subTREe`,
		`subORdinate`,
	} {

		oper := `=`
		want := sprintf("( targetscope %s %q )", oper, lc(raw))
		tscope := Scope(raw)

		sEq := tscope.Eq()
		if sEq.IsZero() {
			t.Errorf("%s failed; failed to create TargetRule with scope '%s'",
				t.Name(), tscope)
			return
		}

		if sc := Scope(idx); sc.Target() != lc(raw) {
			t.Errorf("%s failed; failed to parse scope index '%d'\nwant: '%s'\n got: '%s'",
				t.Name(), idx, lc(raw), sc.Target())
			return
		}

		if sEq.String() != want {
			t.Errorf("%s failed; failed to create equality TargetRule: want '%s', got '%s'",
				t.Name(), want, sEq)
			return
		}

		// negated targetscope rule is illegal per ACIv3 syntax.
		// Make sure we can't generate that particular scope, but
		// are able to generate all others.
		oper = `!` + oper
		want = sprintf("( targetscope %s %q )", oper, lc(raw))
		if sNe := tscope.Ne(); sNe != badTargetRule {
			t.Errorf("%s failed; created illegal targetscope rule '%s' (%s)",
				t.Name(), sNe, want)
			return
		}
	}
}

func ExampleSearchScope_Compare() {
	fmt.Printf("%s == %s: %t", SingleLevel, BaseObject, SingleLevel.Compare(BaseObject))
	// Output: onelevel == base: false
}

func ExampleSearchScope_Keyword() {
	fmt.Printf("%s", SingleLevel.Keyword())
	// Output: targetscope
}

func ExampleSearchScope_String() {
	fmt.Printf("%s", SingleLevel)
	// Output: onelevel
}

func ExampleSearchScope_Target() {
	fmt.Printf("%s", Subordinate) // only valid for target rule scenarios, never URIs!
	// Output: subordinate
}

func ExampleSearchScope_TRM() {
	fmt.Printf("Allows Ne: %t", SingleLevel.TRM().Contains(Ne))
	// Output: Allows Ne: false
}

func ExampleSearchScope_Ne() {
	fmt.Printf("%s", SingleLevel.Ne()) // ILLEGAL!!!!
	// Output:
}

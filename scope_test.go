package aci

import (
	"testing"
)

func TestScope(t *testing.T) {
	for idx, raw := range []string{
		`baSe`,
		`oNe`,
		`sub`,
	} {
		if sc := Scope(raw); sc.String() != lc(raw) {
			t.Errorf("%s failed; failed to parse scope name '%s'",
				t.Name(), raw)
		}
		if sc := Scope(idx); sc.String() != lc(raw) {
			t.Errorf("%s failed; failed to parse scope index '%d'",
				t.Name(), idx)
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
		}

		if sc := Scope(idx); sc.Target() != lc(raw) {
			t.Errorf("%s failed; failed to parse scope index '%d'\nwant: '%s'\n got: '%s'",
				t.Name(), idx, lc(raw), sc.Target())
		}

		if sEq.String() != want {
			t.Errorf("%s failed; failed to create equality TargetRule: want '%s', got '%s'",
				t.Name(), want, sEq)
		}

		// negated targetscope rule is illegal per ACIv3 syntax.
		// Make sure we can't generate that particular scope, but
		// are able to generate all others.
		oper = `!` + oper
		want = sprintf("( targetscope %s %q )", oper, lc(raw))
		if sNe := tscope.Ne(); !sNe.IsZero() {
			t.Errorf("%s failed; created illegal targetscope rule '%s' (%s)",
				t.Name(), sNe, want)
		}
	}
}

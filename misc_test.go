package aci

import (
	"testing"
)

var copMap map[int]string = map[int]string{
	int(Eq): Eq.String(),
	int(Ne): Ne.String(),
	int(Lt): Lt.String(),
	int(Le): Le.String(),
	int(Gt): Gt.String(),
	int(Ge): Ge.String(),
}

var copKeywordTable map[string]map[Keyword][]bool = map[string]map[Keyword][]bool{
	bindRuleID: {
		BindUDN: {true, true, false, false, false, false}, // userdn
		BindRDN: {true, true, false, false, false, false}, // roledn
		BindGDN: {true, true, false, false, false, false}, // groupdn
		BindIP:  {true, true, false, false, false, false}, // ip
		BindAM:  {true, true, false, false, false, false}, // authmethod
		BindDNS: {true, true, false, false, false, false}, // dns
		BindUAT: {true, true, false, false, false, false}, // userattr
		BindGAT: {true, true, false, false, false, false}, // groupattr
		BindDoW: {true, true, false, false, false, false}, // dayofweek
		BindSSF: {true, true, true, true, true, true},     // ssf
		BindToD: {true, true, true, true, true, true},     // timeofday
	},

	targetRuleID: {
		Target:            {true, true, false, false, false, false},  // target
		TargetTo:          {true, true, false, false, false, false},  // target_to
		TargetFrom:        {true, true, false, false, false, false},  // target_from
		TargetCtrl:        {true, true, false, false, false, false},  // targetcontrol
		TargetAttr:        {true, true, false, false, false, false},  // targetattr
		TargetExtOp:       {true, true, false, false, false, false},  // extop
		TargetScope:       {true, false, false, false, false, false}, // targetscope
		TargetFilter:      {true, true, false, false, false, false},  // targetfilter
		TargetAttrFilters: {true, false, false, false, false, false}, // targattrfilters
	},
}

var strInSliceMap map[int]map[int][]bool = map[int]map[int][]bool{
	// case match
	0: {
		0: {true, true, true, true, true},
		1: {true, true, true, true, true},
	},

	// case fold
	1: {
		0: {true, true, true, true, true},
		1: {true, true, true, true, true},
	},
}

func TestStrInSlice(t *testing.T) {
	for idx, fn := range []func(string, []string) bool{
		strInSlice,
		strInSliceFold,
	} {
		for i, values := range [][]string{
			{`cAndidate1`, `blarGetty`, `CANndidate7`, `squatcobbler`, `<censored>`},
			{`Ó-aîï4Åø´øH«w%);<wÃ¯`, `piles`, `4378295fmitty`, string(rune(0)), `broccolI`},
		} {
			for j, val := range values {
				result_expected := strInSliceMap[idx][i][j]

				// warp the candidate value such that
				// it no longer matches the slice from
				// whence it originates. j² is used as
				// its quicker and less stupid than
				// adding a rand generator.
				if isPowerOfTwo(j) {
					var R []rune = []rune(val)
					for g, h := 0, len(R)-1; g < h; g, h = g+1, h-1 {
						R[g], R[h] = R[h], R[g]
					}
					val = string(R)
					result_expected = !result_expected // invert
				}

				result_received := fn(val, values)
				if result_expected != result_received {
					t.Errorf("%s[%d->%d] failed; []byte(%v) in %v: %t (wanted %t)",
						t.Name(), i, j, []byte(val), values, result_received, result_expected)
				}
			}
		}
	}
}

/*
TestOperator_codecov shall test every possible permutation of B/T keywords and
ComparisonOperator. Each permutation result will be compared with the expected
Boolean value.

Also perform various simple checks to satisfy codecov
*/
func TestOperator_codecov(t *testing.T) {
	if keywordAllowsComparisonOperator(Target, float64(3.14592)) {
		t.Errorf("%s failed; resolution error: illegal type permitted", t.Name())
	}

	for i := 0; i < len(copMap); i++ {

		// attempt to resolve the operator
		oper := ComparisonOperator(i + 1)
		if cop := matchCOP(copMap[i+1]); cop != oper {
			t.Errorf("%s failed; resolution error: want '%s', got '%s'",
				t.Name(), oper, cop)
		}

		// traverse the above b/t operator tables
		for _, typ := range []string{
			bindRuleID,
			targetRuleID,
		} {
			// iterate each b/t table's contents,
			// verifying the index (when cast to
			// a keyword index) reveals the same
			// (expected) Boolean results as those
			// reported by the keywordAllows.. func.
			for k, got := range copKeywordTable[typ] {
				want := keywordAllowsComparisonOperator(k, oper)
				if want != got[i] {
					t.Errorf("%s failed; illegal %s operator+keyword [%s + %s @ cop[%d]]: want '%t', got '%t'",
						t.Name(), typ, k, oper, i, want, got[i])
				}

				// retry, using cop STRING instead of actual instance value
				want = keywordAllowsComparisonOperator(k, oper.String())
				if want != got[i] {
					t.Errorf("%s failed; illegal %s operator+keyword [%s + %s]: want '%t', got '%t'",
						t.Name(), typ, k, oper, want, got[i])
				}

				// retry, using keyword STRING instead of actual instance value
				want = keywordAllowsComparisonOperator(k.String(), oper)
				if want != got[i] {
					t.Errorf("%s failed; illegal %s operator+keyword [%s + %s]: want '%t', got '%t'",
						t.Name(), typ, k, oper, want, got[i])
				}
			}
		}
	}
}

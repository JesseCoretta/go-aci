package aci

import (
	"fmt"
	"testing"
)

func ExampleComparisonOperator_stringers() {
	for _, cop := range []ComparisonOperator{
		Eq, Ne, Lt, Gt, Le, Ge,
	} {
		fmt.Printf("[%d] %s (%s)[%s]\n",
			int(cop),
			cop.Description(),
			cop.Context(),
			cop)
	}

	// Output:
	// [1] Equal To (Eq)[=]
	// [2] Not Equal To (Ne)[!=]
	// [3] Less Than (Lt)[<]
	// [4] Greater Than (Gt)[>]
	// [5] Less Than Or Equal (Le)[<=]
	// [6] Greater Than Or Equal (Ge)[>=]
}

func TestComparisonOperator_codecov(t *testing.T) {
	var lousyCop ComparisonOperator = ComparisonOperator(7)
	_ = lousyCop.String()
	_ = lousyCop.Context()
	_ = lousyCop.Description()
	_ = lousyCop.Valid()
	_ = lousyCop.Compare(Eq)
	_ = lousyCop.Compare(`=`)
	_ = lousyCop.Compare(`eq`)
	_ = lousyCop.Compare(`equal to`)
	_ = lousyCop.Compare(3.14567)

	// test permutations of keywords and cops

	permutations := map[string]map[Keyword][]any{
		`valid`: {
			// target keywords
			Target:            {`eq`, `ne`, Eq, Ne, 1, 2},
			TargetTo:          {`eq`, `ne`},
			TargetFrom:        {`eq`, `ne`},
			TargetCtrl:        {`eq`, `ne`},
			TargetAttr:        {`eq`, `ne`},
			TargetFilter:      {`eq`, `ne`, Ne, 2},
			TargetExtOp:       {`eq`, `ne`},
			TargetScope:       {`eq`},
			TargetAttrFilters: {`eq`, 1, Eq},

			// bind keywords
			BindUDN: {`eq`, `ne`, Eq, "equal to", `EQ`},
			BindGDN: {`eq`, Ne, `not equal to`, `NE`, `ne`},
			BindRDN: {`eq`, `ne`},
			BindDNS: {`eq`, `ne`},
			BindUAT: {`eq`, `ne`},
			BindGAT: {`eq`, `ne`},
			BindDoW: {`eq`, `ne`},
			BindIP:  {`eq`, `ne`},
			BindAM:  {`eq`, `ne`},
			BindToD: {`eq`, 4, `ne`, Le, `LE`, 6, `le`, Lt, `LT`, `lt`, 3, Ge, `GE`, `ge`, Gt, `GT`, `gt`},
			BindSSF: {`eq`, 1, `ne`, Le, `LE`, 5, `le`, Lt, `LT`, 2, `lt`, Ge, `GE`, `ge`, Gt, `GT`, `gt`},
		},
	}

	for typ, kwmap := range permutations {
		for kw, values := range kwmap {
			for i := 0; i < len(values); i++ {
				op := values[i]
				if !keywordAllowsComparisonOperator(kw, op) {
					t.Errorf("%s [%s] failed: %s %T [%v] denied or not resolved",
						t.Name(), kw, typ, badCop, op)
					return
				}
			}
		}
	}
}

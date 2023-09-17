package aci

import (
	"fmt"
	"testing"
)

func ExampleComparisonOperator_Compare() {
	fmt.Printf("%Ts are identical: %t", Ne, Ne.Compare(Eq))
	// Output: aci.ComparisonOperators are identical: false
}

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

func ExampleComparisonOperator_Valid() {
	var unknown ComparisonOperator = ComparisonOperator(12)
	fmt.Printf("Is a known %T: %t", unknown, unknown.Valid() == nil)
	// Output: Is a known aci.ComparisonOperator: false
}

/*
This example demonstrates the string representation for all
known ComparisonOperator constants.
*/
func ExampleComparisonOperator_String() {
	for _, cop := range []ComparisonOperator{
		Eq, Ne, Lt, Gt, Le, Ge,
	} {
		fmt.Printf("%s\n", cop)
	}
	// Output:
	// =
	// !=
	// <
	// >
	// <=
	// >=
}

/*
This example demonstrates the use of the Context method to show
all the context name for all ComparisonOperator constants.
*/
func ExampleComparisonOperator_Context() {
	for _, cop := range []ComparisonOperator{
		Eq, Ne, Lt, Gt, Le, Ge,
	} {
		fmt.Printf("%s\n", cop.Context())
	}
	// Output:
	// Eq
	// Ne
	// Lt
	// Gt
	// Le
	// Ge
}

/*
This example demonstrates the use of the Description method to
show all descriptive text for all ComparisonOperator constants.
*/
func ExampleComparisonOperator_Description() {
	for _, cop := range []ComparisonOperator{
		Eq, Ne, Lt, Gt, Le, Ge,
	} {
		fmt.Printf("%s\n", cop.Description())
	}
	// Output:
	// Equal To
	// Not Equal To
	// Less Than
	// Greater Than
	// Less Than Or Equal
	// Greater Than Or Equal
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

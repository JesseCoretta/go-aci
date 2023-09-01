package aci

import (
	"fmt"
)

func ExampleComparisonOperator_stringers() {
	for _, cop := range []ComparisonOperator{
		Eq, Ne, Lt, Le, Gt, Ge,
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
	// [5] Less Than Or Equal (Le)[<=]
	// [4] Greater Than (Gt)[>]
	// [6] Greater Than Or Equal (Ge)[>=]
}

package aci

/*
cop.go contains comparison operator types and methods.
*/

var (
	comparisonOperatorMap              map[string]ComparisonOperator
	permittedTargetComparisonOperators map[TargetKeyword][]ComparisonOperator
	permittedBindComparisonOperators   map[BindKeyword][]ComparisonOperator
)

/*
String wraps go-stackage's ComparisonOperator.String method.
*/
func (r ComparisonOperator) String() string {
	return castAsCop(r).String()
}

/*
Context wraps go-stackage's ComparisonOperator.Context method.
*/
func (r ComparisonOperator) Context() string {
	return castAsCop(r).Context()
}

/*
matchCOP reads the *string representation* of a
ComparisonOperator instance and returns
the appropriate ComparisonOperator const.

A bogus ComparisonOperator (badCop, 0x0)
shall be returned if a match was not made.
*/
func matchCOP(op string) ComparisonOperator {
	for k, v := range comparisonOperatorMap {
		if op == k {
			return v
		}
	}

	return badCop
}

/*
keywordAllowsComparisonOperator returns a boolean value indicative of
whether Keyword input value kw allows ComparisonOperator op
for use in T/B rule instances.

Certain keywords, such as TargetScope, allow only certain operators,
while others, such as BindSSF, allow the use of ALL operators.
*/
func keywordAllowsComparisonOperator(kw, op any) bool {
	// identify the comparison operator,
	// save as cop var.
	var cop ComparisonOperator
	switch tv := op.(type) {
	case string:
		cop = matchCOP(tv)
	case ComparisonOperator:
		cop = tv
	default:
		return false
	}

	// identify the keyword, and
	// pass it onto the appropriate
	// map search function.
	switch tv := kw.(type) {
	case string:
		if bkw := matchBKW(tv); bkw != BindKeyword(0x0) {
			return bindKeywordAllowsComparisonOperator(bkw, cop)

		} else if tkw := matchTKW(tv); tkw != TargetKeyword(0x0) {
			return targetKeywordAllowsComparisonOperator(tkw, cop)
		}
	case BindKeyword:
		return bindKeywordAllowsComparisonOperator(tv, cop)
	case TargetKeyword:
		return targetKeywordAllowsComparisonOperator(tv, cop)
	}

	return false
}

/*
bindKeywordAllowsComparisonOperator is a private function called by keywordAllowsCompariso9nOperator.
*/
func bindKeywordAllowsComparisonOperator(key BindKeyword, cop ComparisonOperator) bool {
	// look-up the keyword within the permitted cop
	// map; if found, obtain slices of cops allowed
	// by said keyword.
	cops, found := permittedBindComparisonOperators[key]
	if !found {
		return false
	}

	// iterate the cops slice, attempting to perform
	// a match of the input cop candidate value and
	// the current cops slice [i].
	for i := 0; i < len(cops); i++ {
		if cop == cops[i] {
			return true
		}
	}

	return false
}

/*
targetKeywordAllowsComparisonOperator is a private function called by keywordAllowsCompariso9nOperator.
*/
func targetKeywordAllowsComparisonOperator(key TargetKeyword, cop ComparisonOperator) bool {
	// look-up the keyword within the permitted cop
	// map; if found, obtain slices of cops allowed
	// by said keyword.
	cops, found := permittedTargetComparisonOperators[key]
	if !found {
		return false
	}

	// iterate the cops slice, attempting to perform
	// a match of the input cop candidate value and
	// the current cops slice [i].
	for i := 0; i < len(cops); i++ {
		if cop == cops[i] {
			return true
		}
	}

	return false
}

func init() {
	comparisonOperatorMap = map[string]ComparisonOperator{
		Eq.String(): Eq,
		Ne.String(): Ne,
		Lt.String(): Lt,
		Le.String(): Le,
		Gt.String(): Gt,
		Ge.String(): Ge,
	}

	// populate the allowed comparison operator map per each
	// possible TargetRule keyword
	permittedTargetComparisonOperators = map[TargetKeyword][]ComparisonOperator{
		Target:            {Eq, Ne},
		TargetTo:          {Eq, Ne},
		TargetFrom:        {Eq, Ne},
		TargetCtrl:        {Eq, Ne},
		TargetAttr:        {Eq, Ne},
		TargetExtOp:       {Eq, Ne},
		TargetScope:       {Eq},
		TargetFilter:      {Eq, Ne},
		TargetAttrFilters: {Eq},
	}

	// populate the allowed comparison operator map per each
	// possible BindRule keyword
	permittedBindComparisonOperators = map[BindKeyword][]ComparisonOperator{
		BindUDN: {Eq, Ne},
		BindRDN: {Eq, Ne},
		BindGDN: {Eq, Ne},
		BindIP:  {Eq, Ne},
		BindAM:  {Eq, Ne},
		BindDNS: {Eq, Ne},
		BindUAT: {Eq, Ne},
		BindGAT: {Eq, Ne},
		BindDoW: {Eq, Ne},
		BindSSF: {Eq, Ne, Lt, Le, Gt, Ge},
		BindToD: {Eq, Ne, Lt, Le, Gt, Ge},
	}
}

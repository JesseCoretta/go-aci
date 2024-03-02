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
String wraps [stackage.ComparisonOperator.String] method. This will return the comparison operator character(s) required by the ACIv3 syntax for a particular expressive statement.

For example, if the receiver is the Eq [ComparisonOperator] constant, the returned string value shall be `=`.

If the receiver is bogus, or describes an unknown [ComparisonOperator] value, the default [stackage] tag "<invalid_operator>" is returned.

See the [ComparisonOperator] constant definitions for details.
*/
func (r ComparisonOperator) String() string {
	return r.cast().String()
}

/*
Context returns the "name" of the [ComparisonOperator] constant.

For example, if the receiver represents the Eq [ComparisonOperator] constant, the returned string value shall be `Eq`.

If the receiver is bogus, or describes an unknown [ComparisonOperator] value, the default [stackage] tag "<invalid_operator>" is returned.

See the [ComparisonOperator] constant definitions for details.
*/
func (r ComparisonOperator) Context() string {
	switch r {
	case Eq:
		return `Eq`
	case Ne:
		return `Ne`
	case Lt:
		return `Lt`
	case Le:
		return `Le`
	case Gt:
		return `Gt`
	case Ge:
		return `Ge`
	}

	return r.String() // [stackage] "<invalid_operator>"
}

/*
Description returns a short description of the receiver instance's context.

For instance, if the receiver is the Eq [ComparisonOperator] constant, the returned string value shall be `Equal To`.

If the receiver is bogus, or describes an unknown [ComparisonOperator] value, the default [stackage] tag "<invalid_operator>" is returned.

This method is largely for convenience, and many individuals may feel it only has any practical applications in the areas of documentation, diagram creation or some other similar activity.

However, a prudent cybersecurity expert may argue that this method can be used to aid in the (critical) area of proofreading newly-devised or modified access control statements. A person could very easily mistake >= and <=, certainly if they're not paying attention. One such mistake could spell disaster.

Additionally, use of this method as a means to auto-generate [Instruction] comments (for LDIF configurations, or similar) can greatly help an admin more easily READ and UNDERSTAND the statements in question.

See the [ComparisonOperator] const definitions for details.
*/
func (r ComparisonOperator) Description() string {
	switch r {
	case Eq:
		return `Equal To`
	case Ne:
		return `Not Equal To`
	case Lt:
		return `Less Than`
	case Le:
		return `Less Than Or Equal`
	case Gt:
		return `Greater Than`
	case Ge:
		return `Greater Than Or Equal`
	}

	return r.String() // [stackage] "<invalid_operator>"
}

/*
Valid returns an error instance following the process of verifying the receiver to be a known [ComparisonOperator] instance.  This does NOT, however, imply feasibility of use with any particular type in the creation of [BindRule] or [TargetRule] instances.
*/
func (r ComparisonOperator) Valid() (err error) {
	if !isValidCopNumeral(int(r)) {
		err = badCopErr(r)
	}

	return
}

/*
Compare shall resolve the input [ComparisonOperator] candidate (cop) and, if successful, shall perform an equality assertion between it and the receiver instance. The assertion result is returned as a bool instance.

In the case of the string representation of a given candidate input value, case-folding is not a significant factor.
*/
func (r ComparisonOperator) Compare(cop any) bool {
	switch tv := cop.(type) {
	case ComparisonOperator:
		return tv == r
	case int:
		return int(tv) == int(r)
	case string:
		return strInSliceFold(tv, []string{
			r.Description(),
			r.Context(),
			r.String(),
		})
	}

	return false
}

/*
isValidCopNumeral merely returns the Boolean evaluation result of a check to see whether integer x falls within a numerical range of one (1) through six (6).

This range represents the absolute minimum and maximum numerical values for any valid instance of the stackage.ComparisonOperator type (and, by necessity, the go-aci [ComparisonOperator] alias type as well).
*/
func isValidCopNumeral(x int) bool {
	return (1 <= x && x <= 6)
}

/*
keywordAllowsComparisonOperator returns a Boolean value indicative of whether Keyword input value kw allows [ComparisonOperator] op for use in T/B rule instances.

Certain keywords, such as [TargetScope], allow only certain operators, while others, such as [BindSSF], allow the use of ALL operators.
*/
func keywordAllowsComparisonOperator(kw, op any) (allowed bool) {
	// identify the comparison operator,
	// save as cop var.
	var cop ComparisonOperator
	switch tv := op.(type) {
	case string:
		cop = matchCOP(tv)
	case ComparisonOperator:
		cop = tv
	case int:
		cop = ComparisonOperator(tv)
	default:
		return
	}

	// identify the keyword, and
	// pass it onto the appropriate
	// map search function.
	switch tv := kw.(type) {
	case string:
		if bkw := matchBKW(tv); bkw != BindKeyword(0x0) {
			allowed = bindKeywordAllowsComparisonOperator(bkw, cop)

		} else if tkw := matchTKW(tv); tkw != TargetKeyword(0x0) {
			allowed = targetKeywordAllowsComparisonOperator(tkw, cop)
		}
	case BindKeyword:
		allowed = bindKeywordAllowsComparisonOperator(tv, cop)
	case TargetKeyword:
		allowed = targetKeywordAllowsComparisonOperator(tv, cop)
	}

	return
}

/*
matchCOP reads the *string representation* of a ComparisonOperator instance and returns the appropriate ComparisonOperator constant.

A bogus ComparisonOperator (badCop, 0x0) shall be returned if a match was not made.
*/
func matchCOP(op string) ComparisonOperator {
	for _, v := range comparisonOperatorMap {
		if strInSliceFold(op, []string{
			v.String(),
			v.Context(),
			v.Description(),
		}) {
			return v
		}
	}

	return badCop
}

/*
bindKeywordAllowsComparisonOperator is a private function called by keywordAllowsComparisonOperator.
*/
func bindKeywordAllowsComparisonOperator(key BindKeyword, cop ComparisonOperator) (allowed bool) {
	// look-up the keyword within the permitted cop
	// map; if found, obtain slices of cops allowed
	// by said keyword.
	if cops, found := permittedBindComparisonOperators[key]; found {
		// iterate the cops slice, attempting to perform
		// a match of the input cop candidate value and
		// the current cops slice [i].
		for i := 0; i < len(cops); i++ {
			if cop == cops[i] {
				allowed = true
				break
			}
		}
	}

	return
}

/*
targetKeywordAllowsComparisonOperator is a private function called by keywordAllowsComparisonOperator.
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

package aci

/*
target.go contains target rule(s) types, functions and methods.
*/

import (
	parser "github.com/JesseCoretta/go-antlraci"
	"github.com/JesseCoretta/go-stackage"
)

var (
	badTargetRule  TargetRule
	badTargetRules TargetRules
)

/*
TargetRule is a stackage.Condition type alias intended to represent
a single Target Rule; that is, one (1) Target Rule keyword, one (1)
comparison operator and one (1) or more string values (called an
'expression').

For example:

	( targetscope = "subordinate" )

Instances of this type may be assembled manually by users, or may be
created logically as a result of textual parsing. Users may also want
to use convenient Eq and Ne methods extended through various types
(as permitted) for simplicity.

Instances of this type shall appear within TargetRules instances.

TargetRule instances are always parenthetical. No parenthetical control
methods exist for instances of this type.
*/
type TargetRule stackage.Condition

/*
TR returns a populated instance of TargetRule. Note there are more
convenient ways of crafting this type instance and in general this
package-level is not needed unless the user wishes to craft the
TargetRule instance in one shot as opposed to incrementally.
*/
func TR(kw, op, ex any) TargetRule {
	return newTargetRule(kw, op, ex)
}

/*
newTargetRule is a private function called by the TR function.
*/
func newTargetRule(kw, op, ex any) (t TargetRule) {
	t.SetKeyword(kw)
	t.SetOperator(op)
	t.SetExpression(ex)

	castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!RulePadding)

	return
}

func (r TargetRule) isConditionContextQualifier() bool {
	return true
}

func (r TargetRule) isRuleContextQualifier() bool {
	return true
}

/*
Valid wraps go-stackage's Condition.Valid method.
*/
func (r TargetRule) Valid() (err error) {
	_t := castAsCondition(r)

	if !keywordAllowsComparisonOperator(_t.Keyword(), _t.Operator()) {
		return
	}
	err = _t.Valid()
	return
}

/*
Kind returns the string literal `condition` to identify the receiver
as a stackage.Condition type alias.
*/
func (r TargetRule) Kind() string {
	return `condition`
}

/*
setID wraps go-stackage's Stack.SetID method.
*/
func (r TargetRule) setID(id string) {
	castAsCondition(r).SetID(id)
}

/*
setCategory wraps go-stackage's Stack.SetCategory method.
*/
func (r TargetRule) setCategory(cat string) {
	castAsCondition(r).SetCategory(cat)
}

/*
Category wraps go-stackage's Condition.Category method.
*/
func (r TargetRule) Category() string {
	return castAsCondition(r).Category()
}

/*
ID wraps go-stackage's Condition.ID method.
*/
func (r TargetRule) ID() string {
	return castAsCondition(r).ID()
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Condition.String method.
*/
func (r TargetRule) String() string {
	return castAsCondition(r).String()
}

/*
SetQuoteStyle allows the election of a particular multivalued
quotation style offered by the various adopters of the ACIv3
syntax.

See the const definitions for MultivalOuterQuotes (default)
and MultivalSliceQuotes for details.
*/
func (r TargetRule) SetQuoteStyle(style int) TargetRule {
	_r := castAsCondition(r)

	switch key := r.Keyword(); key {
	case Target, TargetTo, TargetFrom,
		TargetExtOp, TargetCtrl, TargetAttr:
		switch tv := _r.Expression().(type) {
		case stackage.Stack:
			if key == TargetAttr {
				AttributeTypes(tv).setQuoteStyle(style)
			} else if key == TargetExtOp || key == TargetCtrl {
				ObjectIdentifiers(tv).setQuoteStyle(style)
			} else {
				TargetDistinguishedNames(tv).setQuoteStyle(style)
			}

			// Toggle the individual value quotation scheme
			// to the INVERSE of the Stack quotation scheme
			// set above
			if style == MultivalSliceQuotes {
				_r.Encap()
			} else {
				_r.Encap(`"`)
			}
		}
	}
	//r = TargetRule(*_r)

	return r

}

/*
SetKeyword wraps go-stackage's Condition.SetKeyword method.
*/
func (r *TargetRule) SetKeyword(kw any) {
	x := castAsCondition(*r)
	x.SetKeyword(kw)
	*r = TargetRule(*x)
}

/*
SetOperator wraps go-stackage's Condition.SetOperator method.
Valid input types are stackage.ComparisonOperator (using the
local const aliases) or their string equivalents (e.g.: `>=`
for Ge).
*/
func (r *TargetRule) SetOperator(op any) {
	if !keywordAllowsComparisonOperator(r.Keyword(), op) {
		return
	}

	var cop stackage.ComparisonOperator
	switch tv := op.(type) {
	case string:
		cop = matchCOP(tv)
	case stackage.Operator:
		cop = tv.(stackage.ComparisonOperator)
	default:
		// bogus operator type
		return
	}

	// operator not known? bail out
	if cop == stackage.ComparisonOperator(0) {
		return
	}

	x := castAsCondition(*r)
	x.SetOperator(cop)
	*r = TargetRule(*x)
}

/*
SetExpression wraps go-stackage's Condition.SetExpression method.
*/
func (r *TargetRule) SetExpression(expr any) {
	x := castAsCondition(*r)
	x.SetExpression(expr)
	*r = TargetRule(*x)
}

/*
Keyword wraps go-stackage's Condition.Keyword method and
resolves the raw value into a TargetKeyword. Failure to do
so will return a bogus Keyword.
*/
func (r TargetRule) Keyword() Keyword {
	k := castAsCondition(r).Keyword()
	var kw any = matchTKW(k)
	return kw.(TargetKeyword)
}

/*
Operator wraps go-stackage's Condition.Operator method.
*/
func (r TargetRule) Operator() stackage.Operator {
	return castAsCondition(r).Operator()
}

/*
Expression wraps go-stackage's Condition.Expression method.
*/
func (r TargetRule) Expression() any {
	return castAsCondition(r).Expression()
}

/*
setQuoteStyle shall set the receiver instance to the quotation
scheme defined by integer i.

DECOM
*/
/*
func (r TargetRule) setQuoteStyle(i int) {
	if r.IsZero() {
		return
	}

	_t := castAsCondition(r).Encap()
	if i == 1 {
		_t.Encap(`"`)
	} else {
		_t.Encap()
	}
}*/

/*
IsZero wraps go-stackage's Condition.IsZero method.
*/
func (r TargetRule) IsZero() bool {
	return castAsCondition(r).IsZero()
}

/*
TargetRules is a stackage.Stack type alias intended to store and express
one (1) or more Target Rule statements.

For example:

	( targetscope = "subordinate" )( targetattr = "cn || sn || givenName || objectClass" )

Instances of this type may be assembled manually by users, or may be
created logically as a result of textual parsing. See the T function
for easily initializing and returning instances of this type.

Instances of this type will not allow nesting (i.e.: the addition of any
stackage.Stack type alias instances). Only individual TargetRule instances
may be pushed into instances of this type.
*/
type TargetRules stackage.Stack

/*
func (r TargetRules) isStackContextQualifier() bool {
	return true
}
*/

/*
func (r TargetRules) isRuleContextQualifier() bool {
	return true
}
*/

/*
Kind returns the string literal `stack` to identify the receiver as
a stackage.Stack type alias.
*/
func (r TargetRules) Kind() string {
	return `stack`
}

/*
TRs creates and returns a new instance of TargetRules with an initialized
embedded stack configured to function as a Target Rule store that is
meant to contain one (1) or more Condition instances, each of which bear
one (1) of the following Target Rule keyword constants:

• Target

• TargetTo

• TargetFrom

• TargetAttr

• TargetCtrl

• TargetScope

• TargetFilter

• TargetAttrFilters

• TargetExtOp

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
TargetRule type (or its string equivalent) during initialization. This is merely a
more convenient alternative to separate initialization and push procedures.

Please note that instances of this design are set with a maximum capacity
of nine (9) for both the following reasons:

• There are only said number of Target Rule keywords supported within the
ACI syntax specification honored by this package, and ...

• Individual Target Rule keywords can only be used once per ACI; in other
words, one cannot specify multiple `target` conditions within the same
TargetRules instance.

Instances of this design generally are assigned to top-level instances of
Instruction, and never allow nesting elements (e.g.: other stackage.Stack
derived type aliases).
*/
func TRs(x ...any) (t TargetRules) {
	// create a native stackage.Stack
	// and configure before typecast.
	_t := stackList(9).
		NoNesting(true).
		SetDelimiter(``).
		SetID(targetRuleID).
		NoPadding(!StackPadding).
		SetCategory(targetRuleID)

	// cast _t as a proper TargetRules instance
	// (t). We do it this way to gain access to
	// the method for the *specific instance*
	// being created (t), thus allowing a custom
	// push policy to be set.
	t = TargetRules(_t)

	// Set custom push policy per go-stackage
	// signatures.
	_t.SetPushPolicy(t.pushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_t.Push(x...)

	return
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Stack.String method.
*/
func (r TargetRules) String() string {
	_t, _ := castAsStack(r)
	return _t.String()
}

/*
setCategory wraps go-stackage's Stack.SetCategory method.
*/
func (r TargetRules) setCategory(cat string) {
	if r.IsZero() {
		return
	}

	_t, _ := castAsStack(r)
	_t.SetCategory(cat)
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r TargetRules) IsZero() bool {
	_t, _ := castAsStack(r)
	return _t.IsZero()
}

/*
reset wraps go-stackage's Stack.Reset method. This is a private
method in go-aci.
*/
func (r TargetRules) reset() {
	if r.IsZero() {
		return
	}

	_t, _ := castAsStack(r)
	_t.Reset()
}

/*
ID wraps go-stackage's Stack.ID method.
*/
func (r TargetRules) ID() string {
	if r.IsZero() {
		return ``
	}

	_t, _ := castAsStack(r)
	return _t.ID()
}

/*
Category wraps go-stackage's Stack.Category method.
*/
func (r TargetRules) Category() string {
	if r.IsZero() {
		return ``
	}

	_t, _ := castAsStack(r)
	return _t.Category()
}

/*
setID wraps go-stackage's Stack.SetID method.
*/
func (r TargetRules) setID(id string) {
	if r.IsZero() {
		return
	}

	_t, _ := castAsStack(r)
	_t.SetID(id)
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r TargetRules) Len() int {
	_t, _ := castAsStack(r)
	return _t.Len()
}

/*
Push wraps go-stackage's Stack.Push method.
*/
func (r TargetRules) Push(x ...TargetRule) TargetRules {
	_r, _ := castAsStack(r)
	// iterate variadic input arguments
	for i := 0; i < len(x); i++ {
		_r.Push(x[i])
	}

	return r
}

/*
Pop wraps go-stackage's Stack.Pop method. An instance of
TargetRule is returned following a call of this method.

Within the context of the receiver type, a RuleContext, if
non-nil, can only represent a TargetRule instance.
*/
func (r TargetRules) Pop() TargetRule {
	_r, _ := castAsStack(r)
	x, _ := _r.Pop()

	assert, _ := x.(TargetRule)
	return assert
}

/*
remove wraps go-stackage's Stack.Remove method.
*/
func (r TargetRules) remove(idx int) bool {
	_t, _ := castAsStack(r)
	_, ok := _t.Remove(idx)
	return ok
}

/*
Index wraps go-stackage's Stack.Index method.
*/
func (r TargetRules) Index(idx int) TargetRule {
	_r, _ := castAsStack(r)
	y, _ := _r.Index(idx)

	assert, _ := y.(TargetRule)
	return assert
}

/*
ReadOnly wraps go-stackage's Stack.ReadOnly method.
*/
func (r TargetRules) ReadOnly(state ...bool) {
	_t, _ := castAsStack(r)
	_t.ReadOnly(state...)
}

/*
Insert wraps go-stackage's Stack.Insert method.
*/
func (r TargetRules) insert(x any, left int) bool {
	_t, _ := castAsStack(r)

	assert, ok := x.(TargetRule)
	if !ok {
		return false
	}

	ok = _t.Insert(assert, left)
	//r = TargetRules(_t)
	return ok
}

/*
replace wraps go-stackage's Stack.Replace method.
*/
func (r TargetRules) replace(x any, idx int) bool {
	_t, _ := castAsStack(r)

	_, ok := x.(TargetRule)
	if !ok {
		return false
	}

	return _t.Replace(x, idx)
}

/*
NoPadding wraps go-stackage's Stack.NoPadding method.
*/
func (r TargetRules) NoPadding(state ...bool) {
	_t, _ := castAsStack(r)
	_t.NoPadding(state...)
}

/*
Valid wraps go-stackage's Stack.Valid method.
*/
func (r TargetRules) Valid() (err error) {
	_t, _ := castAsStack(r)
	err = _t.Valid()
	return
}

/*
setQuoteStyle shall set the receiver instance to the quotation
scheme defined by integer i.  This is a private method called
during parsing of AttributeTypes instances, typically following
the ANTLR phase.
*/
func (r TargetRules) setQuoteStyle(i int) {
	if r.IsZero() {
		return
	}

	_t, _ := castAsStack(r)
	if i == 1 {
		_t.Encap()
	} else {
		_t.Encap(`"`)
	}
}

/*
targetRulesPushPolicy conforms to the PushPolicy signature
defined within go-stackage.  This function will be called
privately whenever an instance is pushed into a particular
stackage.Stack (or alias) type instance.

Only TargetRule instances are to be cleared for push executions.
*/
func (r TargetRules) pushPolicy(x any) (err error) {
	switch tv := x.(type) {
	case TargetRule:
		if tv.IsZero() {
			err = errorf("Push request of %T into %T [%s] failed: instance is nil",
				tv, r, tv.Keyword())
		}
		if r.contains(tv.Keyword()) {
			err = errorf("Cannot push non-unique or invalid %T into %T [%s]",
				x, r, tv.Keyword())
		}
	default:
		err = errorf("Push request of %T type violates %T PushPolicy", x, r)
	}
	return
}

/*
Contains returns a boolean value indicative of whether value x,
if a string or TargetKeyword instance, already resides within
the receiver instance.

Case is not significant in the matching process.
*/
func (r TargetRules) Contains(x any) bool {
	return r.contains(x)
}

/*
contains is a private method called by AttributeTypes.Contains.
*/
func (r TargetRules) contains(x any) bool {
	if r.Len() == 0 {
		return false
	}

	var candidate string

	switch tv := x.(type) {
	case string:
		candidate = tv

	case Keyword:
		if kw := matchTKW(tv.String()); kw == TargetKeyword(0x0) {
			return false
		}
		candidate = tv.String()
	default:
		return false
	}

	if len(candidate) == 0 {
		return false
	}

	for i := 0; i < r.Len(); i++ {
		tr := r.Index(i).Keyword().String()
		// case is not significant here.
		if eq(tr, candidate) {
			return true
		}
	}

	return false
}

/*
ParseTargetRule processes the raw input string value,
which should represent a single Target Rule expressive
statement, into an instance of TargetRule. This, along
with an error instance, are returned upom completion
of processing.
*/
func ParseTargetRule(raw string) (TargetRule, error) {
	return parseTargetRule(raw)
}

/*
parseTargetRule is a private function which converts the
stock stackage.Condition instance assembled by go-antlraci
and casts as a go-aci TargetRule instance, which will be
returned alongside an error upon completion of processing.
*/
func parseTargetRule(raw string) (TargetRule, error) {
	_t, err := parser.ParseTargetRule(raw)
	return TargetRule(*_t), err
}

/*
ParseTargetRules processes the raw input string value,
which should represent one (1) or more valid Target Rule
expressive statements, into an instance of TargetRules.
This, alongside an error instance, are returned at the
completion of processing.
*/
func ParseTargetRules(raw string) (TargetRules, error) {
	return parseTargetRules(raw)
}

/*
parseTargetRules is a private function which converts the
stock stackage.Stack instance assembled by go-antlraci and
coaxes the raw string values into proper value-appropriate
type instances made available by go-aci.
*/
func parseTargetRules(raw string) (TargetRules, error) {
	// Call our go-antlraci (parser) package's
	// ParseTargetRules function, and get the
	// results (or bail if error).
	var _t stackage.Stack
	var err error
	if _t, err = parser.ParseTargetRules(raw); err != nil {
		return badTargetRules, err
	}

	// transfer (copy) Target Rule references from _t into _z.
	_z, _ := castAsStack(_t)
	t := TargetRules(_z)

	// iterate our target rule slice members, identifying
	// each by integer index i.
	for i := 0; i < t.Len(); i++ {
		tr := t.Index(0)

		// Extract individual expression value
		// from TargetRule (ntv), and recreate it
		// using the proper type, replacing the
		// original. For example, a `target_to`
		// (DN) Target Rule with a RuleExpression
		// value of:
		//
		//   []string{<dn1>,<dn2>,<dn3>}
		//
		// ... shall be replaced with:
		//
		//   <stack alias type>-idx#------val-
		//   DistinguishedNames[<N1>] -> <dn1>
		//                     [<N2>] -> <dn2>
		//                     [<N3>] -> <dn3>
		if err = tr.assertExpressionValue(); err != nil {
			return badTargetRules, err
		}

		// overwrite old (tv @ index i) with new (ntv)
		//t.replace(ntv, i)
	}

	return t, err
}

/*
assertExpressionValue will update the underlying go-antlraci temporary expression type
with a proper value-appropriate type defined within the go-aci package.

An error is returned upon processing completion.
*/
func (r TargetRule) assertExpressionValue() (err error) {
	// grab the raw value from the receiver. If it is
	// NOT an instance of parser.RuleExpression, then
	// bail out.
	expr, ok := r.Expression().(parser.RuleExpression)
	if !ok {
		err = errorf("Unexpected %T within %T; wanted %T", expr, r, parser.RuleExpression{})
		return
	}

	// our proper type-converted expression
	// value(s) shall reside as an any, as
	// stackage.Condition allows this and
	// will make things simpler.
	var ex any

	// perform a target keyword switch upon
	// a resolution attempt of the value.
	switch key := matchTKW(r.Keyword().String()); key {

	case TargetScope, TargetFilter, TargetAttrFilters:
		// value is a targetscope, targetfilter or a
		// targattrfilters expressive statement. We
		// handle them here because they're strictly
		// single-valued.

		if key == TargetScope {
			// value is a target scope
			ex, err = assertTargetScope(expr)

		} else if key == TargetAttrFilters {
			// value is a targattrfilters
			ex, err = assertTargetAttrFilters(expr)

		} else {
			// value (seems to be) an LDAP Search Filter
			// TODO - assertion func
			//ex, err = assertTargetFilter(expr)
			ex = Filter(expr.Values[0])
		}

	case TargetAttr:
		// value is a targetattr expressive statement,
		// possibly multi-valued.
		ex = assertTargetAttributes(expr)

	case TargetCtrl, TargetExtOp:
		// value is a targetcontrol or extop expressive
		// statement, possibly multi-valued.
		ex, err = assertTargetOID(expr, key)

	case Target, TargetTo, TargetFrom:
		// value is a target, target_to or target_from
		// expressive statement, possibly multi-valued
		ex, err = assertTargetTFDN(expr, key)

	default:
		// value is ... bogus
		err = errorf("Unhandled target rule type '%s'", key)
	}

	if err != nil {
		return
	}

	r.SetExpression(ex)

	return
}

/*
assertTargetScope processes the raw expression value (expr) provided by go-antlraci
into a proper instance of SearchScope (ex), which is returned alongside an instance of
error (err).
*/
func assertTargetScope(expr parser.RuleExpression) (ex SearchScope, err error) {
	if expr.Len() != 1 {
		err = errorf("Unexpected number of %s values; want %d, got %d",
			TargetScope, 1, expr.Len())
		return
	}

	// base is a fallback for a bogus scope, so
	// if the user did not originally request
	// base, we know they requested something
	// totally unsupported.
	if ex = strToScope(expr.Values[0]); ex == noScope {
		err = errorf("Bogus %s value: '%s'", TargetScope, expr.Values[0])
	}

	return
}

/*
assertTargetOID is handler for all possible OID values used within Target Rule expressive
statements. In particular, this handles `targetcontrol` and `extop`.

An ObjectIdentifiers instance is returned in the event that the raw value(s) represent one
(1) or more legal ASN.1 Object Identifiers in "dot notation".

Quotation schemes are supported seamlessly and either scheme shall be honored per the ANTLR4
parsed content.
*/
func assertTargetOID(expr parser.RuleExpression, key TargetKeyword) (ex ObjectIdentifiers, err error) {
	// Don't waste time if expression values
	// are nonexistent.
	if expr.Len() == 0 {
		err = errorf("Found no %s target rule expression value(s) during DN assertion", key)
		return
	}

	// create an appropriate container based on the
	// Target Rule keyword.
	ex = stackByOIDKeyword(key)

	// Honor the established quotation scheme that
	// was observed during ANTLR4 processing.
	ex.setQuoteStyle(expr.Style)

	// Assign the raw (DN) values to the
	// return value. If nothing was found,
	// bail out now.
	if ex.setExpressionValues(key, expr.Values...); ex.Len() == 0 {
		err = errorf("Found none of targetcontrol or extop %T instances",
			ObjectIdentifier{})
		return
	}

	return
}

/*
assertTargetTFDN is handler for all possible DN values used within Target Rule expressive
statements. In particular, this handles `target`, `target_to` and `target_from` keyword
contexts.

A DistinguishedNames instance is returned in the event that the raw value(s) represent one
(1) or more legal LDAP Distinguished Name value.

Quotation schemes are supported seamlessly and either scheme shall be honored per the ANTLR4
parsed content.
*/
func assertTargetTFDN(expr parser.RuleExpression, key TargetKeyword) (ex TargetDistinguishedNames, err error) {
	// Don't waste time if expression values
	// are nonexistent.
	if expr.Len() == 0 {
		err = errorf("Found no %s target rule expression value(s) during DN assertion",
			key)
		return
	}

	// create an appropriate container based on the
	// Target Rule keyword.
	tdn := stackByTDNKeyword(key)

	// Honor the established quotation scheme that
	// was observed during ANTLR4 processing.
	tdn.setQuoteStyle(expr.Style)

	// Assign the raw (DN) values to the
	// return value. If nothing was found,
	// bail out now.
	if tdn.setExpressionValues(key, expr.Values...); tdn.Len() == 0 {
		err = errorf("Found none of userdn, groupdn or roledn %T instances", badTargetDN)
		return
	}

	// Envelope our DN stack within an
	// 'any' instance, which is returned.
	ex = tdn

	return
}

/*
assertTargetAttributes is a private functions called during the processing of a TargetRule
expressive statement bearing the `targetattr` keyword context. An instance of AttributeTypes
is returned.
*/
func assertTargetAttributes(expr parser.RuleExpression) (ex AttributeTypes) {
	ex = TAs()
	ex.setQuoteStyle(expr.Style)

	for i := 0; i < expr.Len(); i++ {
		ex.Push(AT(expr.Values[i]))
	}

	return
}

/*
assertTargetAttrFilters is a private function called during the processing of a TargetRule
expressive statement bearing the `targattrfilters` keyword context. An instance of the
AttributeFilterOperations type is returned, alongside an error instance, when processing is
complete.
*/
func assertTargetAttrFilters(expr parser.RuleExpression) (ex AttributeFilterOperations, err error) {
	if expr.Len() != 1 {
		err = errorf("Unexpected number of %s values; want %d, got %d",
			TargetAttrFilters, 1, expr.Len())
		return
	}

	if idx := idxr(expr.Values[0], ','); idx != -1 {
		// First, try to split on a comma rune (ASCII #44).
		// This is the default, and is the most common char
		// for use in delimiting values of this form.
		ex, err = parseAttributeFilterOperations(expr.Values[0], 0)

	} else if idx = idxr(expr.Values[0], ';'); idx != -1 {
		// If no comma was found, try semicolon (ASCII #59).
		ex, err = parseAttributeFilterOperations(expr.Values[0], 1)

	} else if hasAttributeFilterOperationPrefix(expr.Values[0]) {
		// Still nothing? Try AttributeFilterOperation (whether
		// multivalued or not).
		var afo AttributeFilterOperation
		if afo, err = parseAttributeFilterOperation(expr.Values[0]); err != nil {
			return
		}
		ex = AFOs(afo)

	} else {
		// The only other thing it could be is a bare AttributeFilter.
		var af AttributeFilter
		if af, err = parseAttributeFilter(expr.Values[0]); err != nil {
			return
		}

		ex = AFOs(AddOp.AFO(af)) // we have to choose one, 'add' seems safer than 'delete'
	}

	return
}

const targetRuleID = `target`

package aci

/*
filter.go contains types, functions, methods and constants that pertain to basic
LDAP Search Filter concepts, as well as TargetRule filter-related abstracts.
*/

import (
	"github.com/JesseCoretta/go-stackage"
)

/*
SearchFilter is a struct type that embeds an LDAP search filter. Instances of this type
may be used in a variety of areas, from LDAPURI composition to targetfilter rules.
*/
type SearchFilter struct {
	*searchFilter
}

/*
searchFilter is a private (pointer!) type embedded within instances of SearchFilter.
*/
type searchFilter struct {
	string
}

/*
IsZero returns a boolean value indicative of whether the receiver is nil, or unset.
*/
func (r SearchFilter) IsZero() bool {
	return r.searchFilter == nil
}

/*
Filter initializes (and optionally sets) a new instance of SearchFilter.
Instances of this kind are used in LDAPURIs, as well as certain target
rules.
*/
func Filter(x ...any) (r SearchFilter) {
	r = SearchFilter{new(searchFilter)}
	r.searchFilter.set(x...)
	return
}

/*
String is a stringer method that returns the string representation of
an LDAP Search Filter.
*/
func (r SearchFilter) String() string {
	if r.searchFilter == nil {
		return ``
	}

	return r.searchFilter.string
}

/*
Set assigns the provided value as the LDAP Search Filter instance within the
receiver. Note that this should only be done once, as filters cannot easily
built "incrementally" by the user.
*/
func (r SearchFilter) Set(x ...any) SearchFilter {
	r.searchFilter.set(x...)
	return r
}

/*
set is a private method executed by SearchFilter.Set.
*/
func (r *searchFilter) set(x ...any) {
	if len(x) == 0 {
		return
	}

	if r == nil {
		r = new(searchFilter)
	}

	switch tv := x[0].(type) {
	case string:
		if len(tv) > 0 {
			r.string = tv
		}
	}
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To a `targetfilter` Target Keyword
context.
*/
func (r SearchFilter) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(TargetFilter)
	t.SetOperator(Eq)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!ConditionPadding).
		SetCategory(TargetFilter.String())

	t = TargetRule(*_t)
	return t
}

/*
Ne initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Not-Equal-To a `targetfilter` Target Keyword
context.
*/
func (r SearchFilter) Ne() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(TargetFilter)
	t.SetOperator(Ne)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!ConditionPadding).
		SetCategory(TargetFilter.String())

	t = TargetRule(*_t)
	return t
}

/*
AttributeFilter is a struct type that embeds an AttributeTyp and filter-style Rule.
Instances of this type are a component in the creation of Target Rule definitions
based upon the targattrfilters keyword.
*/
type AttributeFilter struct {
	*atf
}

/*
atf is the embedded type (as a pointer!) within instances of AttributeFilter.
*/
type atf struct {
	AttributeType // single LDAP AttributeType
	SearchFilter  // single LDAP Search Filter
}

/*
AttributeOperation defines either an Add Operation or a Delete Operation.

Constants of this type are used in AttributeFilters instances.
*/
type AttributeOperation uint8

/*
AttributeOperation constants are used to initialize and return AttributeFilters
instances based on one (1) of the possible two (2) constants defined below.
*/
const (
	noAOp AttributeOperation = iota
	AddOp                    // add=
	DelOp                    // delete=
)

/*
AF initializes, optionally sets and returns a new instance of AttributeFilter,
which is a critical component of the `targattrfilters` Target Rule.

Input values must be either a SearchFilter or an AttributeType.
*/
func AF(x ...any) AttributeFilter {
	return AttributeFilter{newAttrFilter(x...)}
}

func newAttrFilter(x ...any) *atf {
	a := new(atf)
	if len(x) > 0 {
		a.set(x...)
	}
	return a
}

/*
Set assigns the provided address component to the receiver and
returns the receiver instance in fluent-form.

Multiple values can be provided in variadic form, or piecemeal.
*/
func (r *AttributeFilter) Set(x ...any) *AttributeFilter {
	if r.atf == nil {
		r.atf = new(atf)
	}

	r.atf.set(x...)
	return r
}

/*
set is a private method called by AttributeFilter.Set.
*/
func (r *atf) set(x ...any) {
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case string:
			if isIdentifier(tv) {
				r.AttributeType = AT(tv)
			} else {
				r.SearchFilter = Filter(tv)
			}
		case AttributeType:
			r.AttributeType = tv
		case SearchFilter:
			r.SearchFilter = tv
		}
	}
}

/*
String is a stringer method that returns the string representation of the
receiver instance.
*/
func (r AttributeFilter) String() string {
	if err := r.Valid(); err != nil {
		return ``
	}

	return sprintf("%s:%s", r.atf.AttributeType, r.atf.SearchFilter)
}

/*
Valid returns an error indicative of whether the receiver is in an aberrant
state.
*/
func (r AttributeFilter) Valid() (err error) {
	if r.IsZero() {
		err = errorf("%T instance is nil", r)
	}
	return
}

/*
IsZero returns a boolean value indicative of whether the receiver is nil,
or unset.
*/
func (r AttributeFilter) IsZero() bool {
	if r.atf == nil {
		return true
	}
	return r.atf.SearchFilter.IsZero() &&
		r.atf.AttributeType.IsZero()
}

/*
Mode returns one (1) of the following string values, indicating the
operational disposition of the receiver:

• `add`

• `delete`

See the AttributeOperation constants for details.
*/
func (r AttributeOperation) Mode() string {
	return r.String()
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r AttributeOperation) String() string {
	if r == DelOp {
		return `delete`
	}
	return `add`
}

type AttributeFilterOperations stackage.Stack

/*
AFOs returns a freshly initialized instance of AttributeFilterOperations, configured
to store one (1) or more AttributeFilterOperation instances for the purpose of crafting
TargetRule instances which bear the `targattrfilters` keyword context.

Instances of this design are not generally needed elsewhere.

Values are automatically joined using stackage.List() with JoinDelim for comma
delimitation.
*/
func AFOs() AttributeFilterOperations {
	return AttributeFilterOperations(stackList().
		JoinDelim(`,`).
		SetID(targetRuleID).
		NoPadding(!RulePadding).
		SetCategory(TargetAttrFilters.String()).
		SetPushPolicy(attrFilterOpsPushPolicy))
}

/*
SetDelimMode controls the delimitation scheme employed by the receiver.

Some vendors use semicolon (ASCII #59) for delimitation. This can be achieved
using an integer value of one (1) to execute this method.

Other vendors use a comma (ASCII #44) for delimitation. This is the default,
but can be set using any integer other than one (1).
*/
func (r AttributeFilterOperations) SetDelimMode(i int) AttributeFilterOperations {
	nx, conv := castAsStack(r) // cast to stackage.Stack to set category
	if !conv {
		return r
	}

	switch i {
	case 1:
		nx.JoinDelim(`;`)
	default:
		nx.JoinDelim(`,`)
	}

	r = AttributeFilterOperations(nx)
	return AttributeFilterOperations(r)
}

/*
Push wraps go-stackage's Stack.Push method.
*/
func (r AttributeFilterOperations) Push(x any) AttributeFilterOperations {
	_r, _ := castAsStack(r)
	_r.Push(x)
	r = AttributeFilterOperations(_r)

	return r
}

/*
Pop wraps go-stackage's Stack.Pop method.
*/
func (r AttributeFilterOperations) Pop() (slice any) {
	_r, _ := castAsStack(r)
	slice, _ = _r.Pop()

	return slice
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r AttributeFilterOperations) Len() int {
	_r, _ := castAsStack(r)
	return _r.Len()
}

/*
Index wraps go-stackage's Stack.Index method. Note that the
Boolean OK value returned by go-stackage by default will be
shadowed and not obtainable by the caller.
*/
func (r AttributeFilterOperations) Index(idx int) (x AttributeFilterOperation) {
	_r, _ := castAsStack(r)
	slice, _ := _r.Index(idx)
	if assert, ok := slice.(AttributeFilterOperation); ok {
		x = assert
	}

	return
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r AttributeFilterOperations) IsZero() bool {
	_r, _ := castAsStack(r)
	return _r.IsZero()
}

/*
String is a stringer method that returns the string representation of
the receiver instance.
*/
func (r AttributeFilterOperations) String() string {
	_r, _ := castAsStack(r)
	return _r.String()
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To a `targattrfilters` keyword context.
*/
func (r AttributeFilterOperations) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(TargetAttrFilters)
	t.SetOperator(Eq)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!ConditionPadding).
		SetCategory(TargetAttrFilters.String())

	t = TargetRule(*_t)
	return t
}

/*
Ne performs no useful task, as negated equality comparison does not apply
to TargetRule instances that bear the `targattrfilters` keyword context.

This method exists solely to convey this message and conform to Go's interface
qualifying signature. When executed, this method will return a bogus TargetRule.
*/
func (r AttributeFilterOperations) Ne() TargetRule { return badTargetRule }

func attrFilterOpsPushPolicy(x any) (err error) {
	switch tv := x.(type) {
	case AttributeFilterOperation:
		if tv.IsZero() {
			err = errorf("%T denied per PushPolicy method; nil %T", tv)
		}
	}

	return
}

func attrFilterOpPushPolicy(x any) (err error) {
	switch tv := x.(type) {
	case AttributeFilter:
		if tv.IsZero() {
			err = errorf("%T denied per PushPolicy method; nil %T", tv)
		}
	}

	return
}

/*
AttributeFilterOperation is a stackage.Stack type alias used to store Target Attribute
Filters expressions, specifically those used within TargetRule instances bearing the
`targattrfilters` TargetRule keyword context.

See also the AttributeFilterOperations type and its methods.
*/
type AttributeFilterOperation stackage.Stack

/*
Push wraps go-stackage's Stack.Push method.
*/
func (r AttributeFilterOperation) Push(x any) AttributeFilterOperation {
	_r, _ := castAsStack(r)
	_r.Push(x)

	return r
}

/*
Pop wraps go-stackage's Stack.Pop method.
*/
func (r AttributeFilterOperation) Pop() (af AttributeFilter) {
	_r, _ := castAsStack(r)
	slice, _ := _r.Pop()

	if assert, ok := slice.(AttributeFilter); ok {
		af = assert
	}

	return
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r AttributeFilterOperation) Len() int {
	_r, _ := castAsStack(r)
	return _r.Len()
}

/*
Index wraps go-stackage's Stack.Index method. Note that the
Boolean OK value returned by go-stackage by default will be
shadowed and not obtainable by the caller.
*/
func (r AttributeFilterOperation) Index(idx int) (slice any) {
	_r, _ := castAsStack(r)
	slice, _ = _r.Index(idx)
	return
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r AttributeFilterOperation) IsZero() bool {
	_r, _ := castAsStack(r)
	return _r.IsZero()
}

/*
String is a stringer method that returns the string representation of
the receiver instance.
*/
func (r AttributeFilterOperation) String() string {
	if r.IsZero() {
		return ``
	}

	aop := r.Operation()
	afs := r.String()

	return sprintf("%s=%s", aop, afs)
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To a `targattrfilters` keyword context.
*/
func (r AttributeFilterOperation) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(TargetAttrFilters)
	t.SetOperator(Eq)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren().
		SetID(targetRuleID).
		NoPadding(!ConditionPadding).
		SetCategory(TargetAttrFilters.String())

	t = TargetRule(*_t)
	return t
}

/*
Ne performs no useful task, as negated equality comparison does not apply
to TargetRule instances that bear the `targattrfilters` keyword context.

This method exists solely to convey this message and conform to Go's interface
qualifying signature. When executed, this method will return a bogus TargetRule.
*/
func (r AttributeFilterOperation) Ne() TargetRule { return badTargetRule }

/*
AFO returns a freshly initialized instance of AttributeFilterOperation, configured
to store one (1) or more AttributeFilter instances for the purpose of crafting
TargetRule instances which bear the `targattrfilters` keyword context.

Instances of this design are not generally needed elsewhere.

Values are automatically ANDed using stackage.And() in symbol (&&) mode.
*/
func AFO() AttributeFilterOperation {
	return AttributeFilterOperation(stackAnd().
		Symbol(`&&`).
		SetID(targetRuleID).
		NoPadding(!RulePadding).
		SetCategory(TargetAttrFilters.String()).
		SetPushPolicy(attrFilterOpPushPolicy))
}

/*
AFs returns an instance of AttributeFilters based upon the input AttributeFilter
instances.

The instance of AttributeFilters contains an ANDed Rule instance using symbols (`&&`)
and bears the categorical string label of `attrfilters`.
*/
func (r AttributeOperation) AFO(x ...AttributeFilter) (afo AttributeFilterOperation) {
	afo = AFO()
	cat := sprintf("%s_%s", TargetAttrFilters, r)
	afo.setCategory(cat)
	for i := 0; i < len(x); i++ {
		afo.Push(x[i]) // append new slices
	}

	return
}

func (r AttributeFilterOperation) Category() string {
	nx, conv := castAsStack(r) // cast to stackage.Stack to set category
	if !conv {
		return ``
	}
	return nx.Category()
}

func (r AttributeFilterOperation) setCategory(cat string) {
	nx, conv := castAsStack(r) // cast to stackage.Stack to set category
	if !conv {
		return
	}

	nx.SetCategory(cat)
	//r = AttributeFilterOperation(nx)
}

/*
Operation returns AddOp or DelOp as extracted from the receiver's categorical
label. If invalid, an invalid AttributeOperation value is returned.
*/
func (r AttributeFilterOperation) Operation() AttributeOperation {
	switch x := trimPfx(r.Category(), TargetAttrFilters.String()+`_`); lc(x) {
	case `add`:
		return AddOp
	case `delete`:
		return DelOp
	}

	return noAOp
}

func hasAttributeFilterOperationPrefix(raw string) bool {
	switch {
	case hasPfx(raw, `add=`):
		return true

	case hasPfx(raw, `delete=`):
		return true
	}

	return false
}

func parseAttributeFilterOperations(raw string, delim int) (afos AttributeFilterOperations, err error) {
	var vals []string
	switch delim {
	case 1:
		vals = split(raw, `;`)
	default:
		vals = split(raw, `,`)
	}

	afos = AFOs()
	for i := 0; i < len(vals); i++ {
		var afo AttributeFilterOperation
		if afo, err = parseAttributeFilterOperation(vals[i]); err != nil {
			return
		}

		afos.Push(afo)
	}

	return
}

func parseAttributeFilterOperation(raw string) (afo AttributeFilterOperation, err error) {
	var (
		val string
		aop AttributeOperation
		seq []string
	)

	if aop, val, err = parseAttrFilterOperPreamble(raw); err != nil {
		return
	}

	afo = aop.AFO()
	cat := sprintf("%s_%s", TargetAttrFilters, aop)
	afo.setCategory(cat)
	seq = split(trimS(val), `&&`)

	for j := 0; j < len(seq); j++ {
		var af AttributeFilter
		if af, err = parseAttributeFilter(trimS(seq[j])); err != nil {
			return
		}
		afo.Push(af)
	}

	return
}

func parseAttributeFilter(raw string) (r AttributeFilter, err error) {
	idx := idxr(raw, ':')
	if idx == -1 {
		err = errorf("No AttributeFilter delim (:) found in %T", r)
		return
	}

	at := AT(raw[:idx])
	f := Filter(raw[idx+1:])
	r.Set(at, f)

	return
}

func parseAttrFilterOperPreamble(raw string) (aop AttributeOperation, value string, err error) {
	switch {

	case hasPfx(raw, `add=`):
		aop = AddOp
		value = raw[4:]

	case hasPfx(raw, `delete=`):
		aop = DelOp
		value = raw[7:]

	default:
		err = errorf("Invalid %T value prefix; must be add= or delete=", aop)
	}

	return
}

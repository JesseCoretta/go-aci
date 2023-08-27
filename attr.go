package aci

/*
attr.go contains LDAP AttributeType types and methods.
*/

import (
	"github.com/JesseCoretta/go-stackage"
)

/*
invalid value constants used as stringer method returns when
something goes wrong :/
*/
const (
	badAT = `<invalid_attribute_type>`
	badAV = `<invalid_attribute_value>`
)

/*
AttributeBindTypeOrValue contains a statement of the following syntax:

	<AttributeName>#<BindType -OR- AttributeValue>

Instances of this type are used in certain Bind Rules, particularly those that
involve user-attribute or group-attribute keywords.
*/
type AttributeBindTypeOrValue struct {
	BindKeyword // BindUAT or BindGAT keywords only!
	*atbtv      // PTR
}

/*
atbtv is the embedded (BY POINTER!) type found within instances of AttributeBindTypeOrValue.

Slices are as follows:
  - 0: <atname> (AttributeType)
  - 1: <atv> (BindType Keyword -OR- AttributeValue)
*/
type atbtv [2]any

/*
IsZero returns a boolean value indicative of whether the receiver is nil,
or unset.
*/
func (r AttributeBindTypeOrValue) IsZero() bool {
	if r.atbtv == nil {
		return true
	}

	return r.BindKeyword == 0x0
}

/*
UAT (User-Attribute Name + Bind Type -OR- Attribute Value) returns an initialized
instance of AttributeBindTypeOrValue configured for rules that leverage the `userattr`
Bind Rule keyword.
*/
func UAT(x ...any) AttributeBindTypeOrValue {
	return userOrGroupAttr(BindUAT, x...)
}

/*
GAT (Group-Attribute Name + Bind Type -OR- Attribute Value) returns an initialized
instance of AttributeBindTypeOrValue configured for rules that leverage the `groupattr`
Bind Rule keyword.
*/
func GAT(x ...any) AttributeBindTypeOrValue {
	return userOrGroupAttr(BindGAT, x...)
}

/*
userOrGroupAttr is a private package level function called by either the GroupAttr or
UserAttr function. This function is the base initializer for the AttributeBindTypeOrValue
instance returned by said functions.
*/
func userOrGroupAttr(t BindKeyword, x ...any) (A AttributeBindTypeOrValue) {
	A = AttributeBindTypeOrValue{
		t, new(atbtv),
	}

	if len(x) != 0 {
		A.atbtv.set(x...)
	}

	return
}

/*
Set assigns value(s) x to the receiver. The value(s) must be AttributeType
and/or AttributeValue instances, created via the package-level AT and AV
functions respectively.
*/
func (r *AttributeBindTypeOrValue) Set(x ...any) *AttributeBindTypeOrValue {
	r.atbtv.set(x...)
	return r
}

/*
Eq initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Equal-To a `userattr` or `groupattr` Bind
keyword context.
*/
func (r AttributeBindTypeOrValue) Eq() BindRule {
	if r.atbtv.isZero() {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(r.BindKeyword)
	b.SetOperator(Eq)
	b.SetExpression(r)

	_b := castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!ConditionPadding).
		SetCategory(r.BindKeyword.String())

	b = BindRule(*_b)
	return b
}

/*
Ne initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Not-Equal-To a `userattr`, `groupattr`
Bind keyword context.

Negated equality BindRule instances should be used with caution.
*/
func (r AttributeBindTypeOrValue) Ne() BindRule {
	if r.atbtv.isZero() {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(r.BindKeyword)
	b.SetOperator(Ne)
	b.SetExpression(r)

	_b := castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!ConditionPadding).
		SetCategory(r.BindKeyword.String())

	b = BindRule(*_b)
	return b
}

/*
isZero returns a boolean value indicative of whether the receiver is nil, or
unset.
*/
func (r *atbtv) isZero() bool {
	if r == nil {
		return true
	}
	return (r[0] == nil && r[1] == nil)
}

/*
String is a stringer method that returns the string representation of the
receiver.
*/
func (r atbtv) String() string {
	// AttributeType will always
	// be used.
	var at AttributeType

	// Only one (1) of the following
	// vars will be used.
	var bt BindType
	var av AttributeValue

	// Assert the attributeType value
	// or bail out.
	at, assert := r[0].(AttributeType)
	if !assert {
		return ``
	}

	// First see if the value is a BindType
	// keyword, as those are few and easily
	// identified.
	if bt, assert = r[1].(BindType); !assert || bt == BindType(0x0) {
		// If not a BindType kw, see if it
		// appears to be an AttributeValue.
		if av, assert = r[1].(AttributeValue); !assert || len(*av.string) == 0 {
			// value is neither an AttributeValue
			// nor BindType kw; bail out.
			return ``
		}

		// AttributeValue wins
		return sprintf("%s#%s", at, av)
	}

	// BindType wins
	return sprintf("%s#%s", at, bt)
}

/*
set assigns one (1) or more values (x) to the receiver. Only
AttributeType, AttributeValue and BindType instances shall be
assigned.

Note that if a string value is detected, it will be cast as
the appropriate type and assigned to the appropriate slice in
the receiver, but ONLY if said slice is nil.
*/
func (r *atbtv) set(x ...any) {
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case AttributeType:
			if r[0] == nil {
				r[0] = tv
			}
		case AttributeValue, BindType:
			if r[1] == nil {
				r[1] = tv
			}
		case string:
			if r[0] == nil {
				r[0] = AT(tv)
			} else {
				if bt := matchBT(tv); bt != BindType(0x0) {
					r[1] = bt
				} else {
					r[1] = AV(tv)
				}
			}
		}
	}
}

/*
String is a stringer method that returns the string representation of the
receiver.
*/
func (r AttributeBindTypeOrValue) String() (s string) {
	s = badAV
	if r.atbtv != nil {
		s = r.atbtv.String()
	}
	return
}

/*
parseATBTV parses the input string (x) in an attempt to marshal its contents
into an instance of AttributeBindTypeOrValue (A), which is returned alongside
an error (err).

The optional BindKeyword argument (bkw) allows the BindGAT (groupattr) Bind
Rule keyword to be set, else the default of BindUAT (userattr) will take
precedence.
*/
func parseATBTV(x string, bkw ...any) (A AttributeBindTypeOrValue, err error) {
	// Obtain the index number for ASCII #35 (NUMBER SIGN).
	// If minus one (-1), input value x is totally bogus.
	idx := idxr(x, '#')
	if idx == -1 {
		err = badAttributeBindTypeOrValueErr(x)
		return
	}

	// Set the groupattr keyword if requested, else
	// use the default of userattr.
	kw := assertATBTVBindKeyword(bkw...)

	// If the remaining portion of the value is, in
	// fact, a known BIND TYPE keyword, pack it up
	// and ship it out.
	if bt := matchBT(x[idx+1:]); bt != BindType(0x0) {
		A = userOrGroupAttr(kw, AT(x[:idx]), bt)
		return
	}

	// Remaining portion of the value would appear
	// to be an attribute value, so pack it up and
	// send it off.
	A = userOrGroupAttr(kw, AT(x[:idx]), AV(x[idx+1:]))
	return
}

/*
AttributeType embeds a pointer value that reflects a single AttributeType name
such as `manager` or `cn`.
*/
type AttributeType struct {
	*string
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To a `targetattr` keyword context.
*/
func (r AttributeType) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(TargetAttr)
	t.SetOperator(Eq)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!ConditionPadding).
		SetCategory(TargetAttr.String())

	t = TargetRule(*_t)
	return t
}

/*
Ne initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Not-Equal-To a `targetattr` keyword context.

Negated equality TargetRule instances should be used with caution.
*/
func (r AttributeType) Ne() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(TargetAttr)
	t.SetOperator(Ne)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!ConditionPadding).
		SetCategory(TargetAttr.String())

	t = TargetRule(*_t)
	return t
}

/*
AT initializes, sets and returns an AT instance in one shot. The
input value x shall be a string attributeType name (e.g.: `manager`).
*/
func AT(x string) (A AttributeType) {
	if len(x) != 0 {
		A = AttributeType{&x}
	}

	return
}

/*
String returns the string representation of the underlying
value within the receiver. The return value shall reflect
an attributeType name, such as `manager` or `cn`.
*/
func (r AttributeType) String() (s string) {
	s = badAT
	if r.string != nil {
		return (*r.string)
	}

	return
}

/*
IsZero returns a boolean value indicative of whether the receiver is nil,
or unset.
*/
func (r AttributeType) IsZero() bool {
	if r.string == nil {
		return true
	}
	return len(*r.string) == 0
}

/*
AttributeValue embeds a pointer value that reflects an attributeType
assertion value.
*/
type AttributeValue struct {
	*string
}

/*
AV initializes, sets and returns an AttributeValue instance in one shot. The
input value x shall be a known BindType constant, such as USERDN, OR a raw
string attributeType value, such as `uid=bob,ou=People,dc=example,dc=com`.
*/
func AV(x string) (A AttributeValue) {
	if len(x) > 0 {
		A = AttributeValue{&x}
	}
	return
}

/*
String returns the string representation of the underlying value within the receiver.
The return value should be either an attributeType assertion value, or one (1) of the
five (5) possible BindType identifiers (e.g.: USERDN).
*/
func (r AttributeValue) String() (s string) {
	s = badAV
	if r.string != nil {
		s = (*r.string)
	}

	return
}

/*
AttributeTypes is an alias type for stackage.Stack, and is intended
to house one (1) or more AttributeType instances for the purpose of
expression within a BindRule or TargetRule instance.
*/
type AttributeTypes stackage.Stack

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To a `targetattr` keyword context.
*/
func (r AttributeTypes) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(TargetAttr)
	t.SetOperator(Eq)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!ConditionPadding).
		SetCategory(TargetAttr.String())

	t = TargetRule(*_t)
	return t
}

/*
Ne initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Not-Equal-To a `targetattr` keyword context.

Negated equality TargetRule instances should be used with caution.
*/
func (r AttributeTypes) Ne() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(TargetAttr)
	t.SetOperator(Ne)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!ConditionPadding).
		SetCategory(TargetAttr.String())

	t = TargetRule(*_t)
	return t
}

/*
setQuoteStyle shall set the receiver instance to the quotation
scheme defined by integer i.
*/
func (r AttributeTypes) setQuoteStyle(style int) AttributeTypes {
	_r, _ := castAsStack(r)
	if style == MultivalSliceQuotes {
		_r.Encap(`"`)
	} else {
		_r.Encap()
	}
	r = AttributeTypes(_r)
	return r
}

/*
setCategory wraps go-stackage's Stack.SetCategory method privately.

DECOM
*/
func (r AttributeTypes) setCategory(cat string) AttributeTypes {
	_r, _ := castAsStack(r)
	_r.SetCategory(cat)
	r = AttributeTypes(_r)
	return r
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r AttributeTypes) IsZero() bool {
	_r, _ := castAsStack(r)
	return _r.IsZero()
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r AttributeTypes) Len() int {
	_r, _ := castAsStack(r)
	return _r.Len()
}

/*
Index wraps go-stackage's Stack.Index method. Note that the
Boolean OK value returned by go-stackage by default will be
shadowed and not obtainable by the caller.
*/
func (r AttributeTypes) Index(idx int) (x AttributeType) {
	_r, _ := castAsStack(r)
	z, _ := _r.Index(idx)

	if assert, asserted := z.(AttributeType); asserted {
		x = assert
	}

	return
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Stack.String method.
*/
func (r AttributeTypes) String() string {
	_r, _ := castAsStack(r)
	return _r.String()
}

/*
Kind wraps go-stackage's Stack.Category method for the
purpose of identifying the context of the receiver instance.
*/
func (r AttributeTypes) Kind() string {
	_r, _ := castAsStack(r)
	return _r.Category()
}

/*
Pop wraps go-stackage's Stack.Pop method.
*/
func (r AttributeTypes) Pop() (x AttributeType) {
	_r, _ := castAsStack(r)

	z, _ := _r.Pop()
	if assert, asserted := z.(AttributeType); asserted {
		x = assert
	}

	return
}

/*
Push wraps go-stackage's Stack.Push method. Valid input types
are string and AttributeType. In the case of a string value,
it is automatically cast as an instance of AttributeType, so
long as the raw string is of a non-zero length.
*/
func (r AttributeTypes) Push(x ...any) AttributeTypes {
	_r, _ := castAsStack(r)

	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case string:
			if len(tv) != 0 {
				_r.Push(AT(tv))
			}
		case AttributeType:
			if !tv.IsZero() {
				_r.Push(tv)
			}
		}
	}

	r = AttributeTypes(_r)
	return r
}

/*
pushPolicy conforms to the PushPolicy interface signature defined within
go-stackage. This private function is called during Push attempts to a
AttributeTypes stack instance.
*/
func (r AttributeTypes) pushPolicy(x any) (err error) {
	// verify uniqueness; bail out if Boolean
	// false is return value.
	if r.contains(x) {
		err = errorf("Cannot push non-unique or invalid %T into %T [%s]",
			x, r, r.Kind())
		return
	}

	// perform type switch upon input value
	// x to determine suitability for push.
	switch tv := x.(type) {

	case string:
		// case matches a string-based LDAP AttributeType
		if len(tv) == 0 {
			err = errorf("Cannot push zero %T into %T [%s]",
				tv, r, r.Kind())
		}

	case AttributeType:
		// case matches an AttributeType instance
		if tv.IsZero() {
			err = errorf("Cannot push nil %T into %T [%s]",
				tv, r, r.Kind())
		}

	default:
		// unsuitable candidate per type
		err = errorf("%T type violates %T [%s] PushPolicy",
			tv, r, r.Kind())
	}

	return
}

/*
Contains returns a boolean value indicative of whether value x,
if a string or AttributeType instance, already resides within
the receiver instance.

Case is not significant in the matching process.
*/
func (r AttributeTypes) Contains(x any) bool {
	return r.contains(x)
}

/*
contains is a private method called by AttributeTypes.Contains.
*/
func (r AttributeTypes) contains(x any) bool {
	if r.Len() == 0 {
		return false
	}

	var candidate string

	switch tv := x.(type) {
	case string:
		candidate = tv
	case AttributeType:
		candidate = tv.String()
	default:
		return false
	}

	if len(candidate) == 0 {
		return false
	}

	for i := 0; i < r.Len(); i++ {
		// case is not significant here.
		if eq(r.Index(i).String(), candidate) {
			return true
		}
	}

	return false
}

func badAttributeBindTypeOrValueErr(x string) error {
	return errorf("Invalid AttributeBindTyoeOrValue instance: must conform to '<at>#<bt_or_av>', got '%s'", x)
}

/*
TAs returns a freshly initialized instance of AttributeTypes, configured to
store one (1) or more AttributeType instances for the purpose of TargetRule
expression when using the `targetattr` keyword context.

Values are automatically delimited using stackage.Stack.Symbol(`||`) in an
ORed Boolean stack.
*/
func TAs() (a AttributeTypes) {
	// create a native stackage.Stack
	// and configure before typecast.
	_a := stackOr().
		Symbol(`||`).
		NoNesting(true).
		SetID(targetRuleID).
		NoPadding(!RulePadding).
		SetCategory(TargetAttr.String())

	// cast _a as a proper AttributeTypes
	// instance (a). We do it this way to
	// gain access to the method for the
	// *specific instance* being created (a),
	// thus allowing things like uniqueness
	// checks, etc., to occur during push
	// attempts, providing more helpful
	// and non-generalized feedback.
	a = AttributeTypes(_a)
	_a.SetPushPolicy(a.pushPolicy)

	return
}

/*
UAs returns a freshly initialized instance of AttributeTypes, configured
to store one (1) or more AttributeType instances for the purpose of LDAP
Search URI specification of desired AttributeType names. Instances of
this design are not generally needed elsewhere.

Values are automatically comma-delimited using stackage.Stack.JoinDelim
in List mode.
*/
func UAs() (a AttributeTypes) {
	// create a native stackage.Stack
	// and configure before typecast.
	_a := stackList().
		JoinDelim(`,`).
		NoNesting(true).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(`uri_search_attributes`)

	// cast _a as a proper AttributeTypes
	// instance (a). We do it this way to
	// gain access to the method for the
	// *specific instance* being created (a),
	// thus allowing things like uniqueness
	// checks, etc., to occur during push
	// attempts, providing more helpful
	// and non-generalized feedback.
	a = AttributeTypes(_a)
	_a.SetPushPolicy(a.pushPolicy)

	return
}

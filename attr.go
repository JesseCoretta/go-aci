package aci

/*
attr.go contains LDAP AttributeType types and methods.
*/

/*
invalid value constants used as stringer method returns when
something goes wrong :/
*/
const (
	badAT = `<invalid_attribute_type>`
	badAV = `<invalid_attribute_value>`
)

var (
	badAttributeType  AttributeType  // for failed calls that return an AttributeType only
	badAttributeValue AttributeValue // for failed calls that return an AttributeValue only
)

/*
AttributeTypeContext is a convenient interface type that is qualified by the following types:

  - [AttributeType]
  - [AttributeTypes]

The qualifying methods shown below are intended to make the generalized handling of attributeTypes
slightly easier without an absolute need for type assertion at every step.

These methods are inherently read-only in nature.

To alter the underlying value, or to gain access to all of a given type's methods, type assertion
of qualifying instances shall be necessary.
*/
type AttributeTypeContext interface {
	Len() int
	String() string
	Kind() string
	Keyword() Keyword
	IsZero() bool
	Valid() error

	isAttributeTypeContext()
}

/*
AttributeBindTypeOrValue contains a statement of the following syntax:

	<AttributeName>#<BindType -OR- AttributeValue>

Instances of this type are used in certain [BindRules], particularly those that
involve user-attribute or group-attribute [BindKeyword] instances.
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
IsZero returns a Boolean value indicative of whether the receiver is nil,
or unset.
*/
func (r AttributeBindTypeOrValue) IsZero() bool {
	if r.atbtv == nil {
		return true
	}

	return r.BindKeyword == 0x0
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r AttributeBindTypeOrValue) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
ABTV will return a new instance of [AttributeBindTypeOrValue]. The required
[BindKeyword] must be either [BindUAT] or [BindGAT]. The optional input values
(x), if provided, will be used to set the instance.
*/
func ABTV(kw BindKeyword, x ...any) (a AttributeBindTypeOrValue) {
	switch kw {
	case BindUAT, BindGAT:
		a = userOrGroupAttr(kw, x...)
	}

	return
}

/*
UAT (User-Attribute Name + Bind Type -OR- Attribute Value) returns an initialized
instance of [AttributeBindTypeOrValue] configured for rules that leverage the [BindUAT]
[BindKeyword] context.
*/
func UAT(x ...any) AttributeBindTypeOrValue {
	return userOrGroupAttr(BindUAT, x...)
}

/*
GAT (Group-Attribute Name + Bind Type -OR- Attribute Value) returns an initialized
instance of [AttributeBindTypeOrValue] configured for rules that leverage the [BindGAT]
[BindKeyword] context.
*/
func GAT(x ...any) AttributeBindTypeOrValue {
	return userOrGroupAttr(BindGAT, x...)
}

/*
userOrGroupAttr is a private package level function called by either the GroupAttr or UserAttr function. This function is the base initializer for the [AttributeBindTypeOrValue] instance returned by said functions.
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
Set assigns value(s) x to the receiver. The value(s) must be [AttributeType] and/or [AttributeValue] instances, created via the package-level [AT] and [AV] functions respectively.
*/
func (r *AttributeBindTypeOrValue) Set(x ...any) *AttributeBindTypeOrValue {
	if r.IsZero() {
		r.atbtv = new(atbtv)
	}
	r.atbtv.set(x...)
	return r
}

/*
Eq initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Equal-To a [BindUAT] or [BindGAT] [BindKeyword] context.
*/
func (r AttributeBindTypeOrValue) Eq() (b BindRule) {
	if !r.atbtv.isZero() {
		b = BR(r.BindKeyword, Eq, r)
	}
	return
}

/*
Ne initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Not-Equal-To a [BindUAT] or [BindGAT] [BindKeyword] context.

Negated equality [BindRule] instances should be used with caution.
*/
func (r AttributeBindTypeOrValue) Ne() (b BindRule) {
	if !r.atbtv.isZero() {
		b = BR(r.BindKeyword, Ne, r)
	}
	return
}

/*
BRM returns an instance of [BindRuleMethods].

Each of the return instance's key values represent a single instance of the [ComparisonOperator] type that is allowed for use in the creation of [BindRule] instances which bear the receiver instance as an expression value. The value for each key is the actual [BindRuleMethod] instance for OPTIONAL use in the creation of a [BindRule] instance.

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus [BindRule] instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly populated and prepared for such activity.
*/
func (r AttributeBindTypeOrValue) BRM() BindRuleMethods {
	return newBindRuleMethods(bindRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
Keyword returns the [BindKeyword] associated with the receiver instance, enveloped as a [Keyword]. In the context of this type instance, the [BindKeyword] returned will be either [BindUAT] or [BindGAT].
*/
func (r AttributeBindTypeOrValue) Keyword() Keyword {
	var kw Keyword = r.BindKeyword
	switch kw {
	case BindGAT:
		return BindGAT
	}

	return BindUAT
}

/*
isZero returns a Boolean value indicative of whether the receiver is nil, or unset.
*/
func (r *atbtv) isZero() bool {
	if r == nil {
		return true
	}
	return (r[0] == nil && r[1] == nil)
}

/*
String is a stringer method that returns the string representation of the receiver.
*/
func (r atbtv) String() (s string) {
	// Only one (1) of the following
	// vars will be used.
	var bt BindType
	var av AttributeValue

	// Assert the attributeType value
	// or bail out.
	if at, assert := r[0].(AttributeType); assert {
		// First see if the value is a BindType
		// keyword, as those are few and easily
		// identified.
		if bt, assert = r[1].(BindType); !assert || bt == BindType(0x0) {
			// If not a BindType kw, see if it
			// appears to be an AttributeValue.
			if av, assert = r[1].(AttributeValue); !assert || len(*av.string) == 0 {
				// value is neither an AttributeValue
				// nor BindType kw; bail out.
				return
			}

			// AttributeValue wins
			s = sprintf("%s#%s", at, av)
			return
		}

		// BindType wins
		s = sprintf("%s#%s", at, bt)
	}

	return
}

/*
set assigns one (1) or more values (x) to the receiver. Only [AttributeType], [AttributeValue] and [BindType] instances shall be assigned.

Note that if a string value is detected, it will be cast as the appropriate type and assigned to the appropriate slice in the receiver, but ONLY if said slice is nil.
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
Parse reads the input string (raw) in an attempt to marshal its contents into the receiver instance (r). An error is returned at the end of the process.

If no suitable [BindKeyword] is provided (bkw), the default is [BindUAT]. Valid options are [BindUAT] and [BindGAT].
*/
func (r *AttributeBindTypeOrValue) Parse(raw string, bkw ...any) (err error) {
	var _r AttributeBindTypeOrValue
	if _r, err = parseATBTV(raw, bkw); err != nil {
		return
	}
	*r = _r

	return
}

/*
Valid returns an error indicative of whether the receiver is in an aberrant state.
*/
func (r AttributeBindTypeOrValue) Valid() (err error) {
	err = nilInstanceErr(r)
	if !r.IsZero() {
		err = nil
	}

	return
}

/*
parseATBTV parses the input string (x) in an attempt to marshal its contents
into an instance of [AttributeBindTypeOrValue] (A), which is returned alongside
an error (err).

The optional BindKeyword argument (bkw) allows the [BindGAT] (groupattr) Bind
Rule keyword to be set, else the default of [BindUAT] (userattr) will take
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

	at := AT(x[:idx])
	av := AV(x[idx+1:])

	if at.String() == badAT {
		err = badAttributeBindTypeOrValueErr(x)
		return
	}

	// If the remaining portion of the value is, in
	// fact, a known BIND TYPE keyword, pack it up
	// and ship it out.
	if bt := matchBT(x[idx+1:]); bt != BindType(0x0) {
		A = userOrGroupAttr(kw, at, bt)
		return
	}

	A = userOrGroupAttr(kw, at, av)
	return
}

/*
AttributeType embeds a pointer value that reflects a single [AttributeType] descriptor such as `manager` or `cn`. The descriptor should conform to RFC 4512 Section 2.5.
*/
type AttributeType struct {
	*string
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r AttributeType) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Eq initializes and returns a new [TargetRule] instance configured to express the evaluation of the receiver value as Equal-To a [TargetAttr] [TargetKeyword] context.
*/
func (r AttributeType) Eq() (t TargetRule) {
	if !r.IsZero() {
		t = TR(TargetAttr, Eq, r)
	}
	return
}

/*
Ne initializes and returns a new [TargetRule] instance configured to express the evaluation of the receiver value as Not-Equal-To a [TargetAttr] [TargetKeyword] context.

Negated equality [TargetRule] instances should be used with caution.
*/
func (r AttributeType) Ne() (t TargetRule) {
	if !r.IsZero() {
		t = TR(TargetAttr, Ne, r)
	}
	return
}

/*
Kind performs no useful task, as the receiver instance has no concept of a keyword, which is the typical value source for Kind calls. This method exists solely to satisfy Go's interface signature requirements and will return a zero string if executed.
*/
func (r AttributeType) Kind() string { return `` }

/*
Keyword performs no useful task, as the receiver instance has no concept of a keyword. This method exists solely to satisfy Go's interface signature requirements and will return nil if executed.
*/
func (r AttributeType) Keyword() Keyword { return nil }

func (r AttributeType) isAttributeTypeContext() {}

/*
TRM returns an instance of [TargetRuleMethods].

Each of the return instance's key values represent a single instance of the [ComparisonOperator] type that is allowed for use in the creation of [TargetRule] instances which bear the receiver instance as an expression value. The value for each key is the actual [TargetRuleMethod] instance for OPTIONAL use in the creation of a [TargetRule] instance.

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus [TargetRule] instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly initialized, populated and prepared for such activity.
*/
func (r AttributeType) TRM() TargetRuleMethods {
	return newTargetRuleMethods(targetRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
AT initializes, sets and returns an [AttributeType] instance in one shot. The input value x shall be an RFC 4512 Section 2.5 compliant descriptor (e.g.: `manager`).
*/
func AT(x string) (A AttributeType) {
	if isIdentifier(x) || x == `*` {
		A = AttributeType{&x}
	}

	return
}

/*
String returns the string representation of the underlying value within the receiver. The return value shall reflect an LDAP descriptor, such as `manager` or `cn`.
*/
func (r AttributeType) String() (s string) {
	s = badAT
	if r.string != nil {
		return (*r.string)
	}

	return
}

/*
Len returns 0 or 1 to describe an abstract length of the receiver. This method exists only to satisfy Go's interface signature requirements and need not be used.
*/
func (r AttributeType) Len() int {
	if err := r.Valid(); err != nil {
		return 0
	}
	return 1
}

/*
Valid returns an instance of error describing the aberrant state of the receiver, if applicable. At the moment, this method merely verifies nilness, as the [AttributeType] type defined within this package is strictly one dimensional, and lacks any significant mechanics for extended scrutiny.
*/
func (r AttributeType) Valid() error {
	if r.IsZero() {
		return nilInstanceErr(r)
	}

	return nil
}

/*
IsZero returns a Boolean value indicative of whether the receiver is nil, or unset.
*/
func (r AttributeType) IsZero() bool {
	if r.string == nil {
		return true
	}
	return len(*r.string) == 0
}

/*
AttributeValue embeds a pointer value that reflects an attribute value.
*/
type AttributeValue struct {
	*string
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r AttributeValue) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
AV initializes, sets and returns an [AttributeValue] instance in one shot. The input value x shall be a known [BindType] constant, such as [USERDN], OR a raw string value.
*/
func AV(x string) (A AttributeValue) {
	if len(x) > 0 {
		A = AttributeValue{&x}
	}
	return
}

/*
String returns the string representation of the underlying value within the receiver. The return value should be either an attributeType assertion value, or one (1) of the five (5) possible [BindType] identifiers (e.g.: [USERDN]).
*/
func (r AttributeValue) String() (s string) {
	s = badAV
	if r.string != nil {
		s = (*r.string)
	}

	return
}

/*
F returns the appropriate instance creator function for crafting individual [AttributeType] instances for submission to the receiver. This is merely a convenient alternative to maintaining knowledge as to which function applies to the current receiver instance.

As there is only one possibility for instances of this design, the [AT] function is returned.
*/
func (r AttributeTypes) F() func(string) AttributeType {
	return AT
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r AttributeTypes) Compare(x any) bool {
	return compareHashInstance(r, x)
}

func (r AttributeTypes) reset() {
	r.cast().Reset()
}

func (r AttributeTypes) resetKeyword(x any) {
	if r.Len() > 0 {
		return
	}

	switch tv := x.(type) {
	case Keyword:
		r.resetKeyword(tv.String())

	case string:
		_r := r.cast()
		switch lc(tv) {
		case TargetAttr.String():
			_r.SetCategory(lc(tv)).
				SetPushPolicy(r.pushPolicy)

		case BindUAT.String(), BindGAT.String():
			_r.SetCategory(tv).
				SetPushPolicy(r.pushPolicy)

		}
	}
}

/*
Eq initializes and returns a new [TargetRule] instance configured to express the evaluation of the receiver value as Equal-To a [TargetAttr] [TargetKeyword] context.
*/
func (r AttributeTypes) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}
	return TR(TargetAttr, Eq, r)
}

/*
Ne initializes and returns a new [TargetRule] instance configured to express the evaluation of the receiver value as Not-Equal-To a [TargetAttr] [TargetKeyword] context.

Negated equality [TargetRule] instances should be used with caution.
*/
func (r AttributeTypes) Ne() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	return TR(TargetAttr, Ne, r)
}

func (r AttributeTypes) isAttributeTypeContext() {}

/*
TRM returns an instance of [TargetRuleMethods].

Each of the return instance's key values represent a single instance of the [ComparisonOperator] type that is allowed for use in the creation of [TargetRule] instances which bear the receiver instance as an expression value. The value for each key is the actual [TargetRuleMethod] instance for OPTIONAL use in the creation of a [TargetRule] instance.

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus [TargetRule] instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly populated and prepared for such activity.
*/
func (r AttributeTypes) TRM() TargetRuleMethods {
	if r.Keyword() == TargetAttr {
		return newTargetRuleMethods(targetRuleFuncMap{
			Eq: r.Eq,
			Ne: r.Ne,
		})
	}

	return TargetRuleMethods{nil}
}

/*
setQuoteStyle shall set the receiver instance to the quotation
scheme defined by integer i.
*/
func (r AttributeTypes) setQuoteStyle(style int) AttributeTypes {
	_r := r.cast()
	if _r.Len() < 2 {
		_r.Encap() // not multivalued, force default
		return r
	}

	if style == MultivalSliceQuotes {
		_r.Encap(`"`)
	} else {
		_r.Encap()
	}

	return r
}

/*
IsZero wraps the [stackage.Stack.IsZero] method.
*/
func (r AttributeTypes) IsZero() bool {
	return r.cast().IsZero()
}

/*
Len wraps the [stackage.Stack.Len] method.
*/
func (r AttributeTypes) Len() int {
	return r.cast().Len()
}

/*
Index wraps the [stackage.Stack.Index] method. Note that the Boolean OK value returned by [stackage] by default will be shadowed and not obtainable by the caller.
*/
func (r AttributeTypes) Index(idx int) (x AttributeType) {
	z, _ := r.cast().Index(idx)
	if assert, asserted := z.(AttributeType); asserted {
		x = assert
	}

	return
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps the [stackage.Stack.String] method.
*/
func (r AttributeTypes) String() string {
	return r.cast().String()
}

/*
Kind wraps the [stackage.Stack.Category] method for the
purpose of identifying the context of the receiver instance.
*/
func (r AttributeTypes) Kind() string {
	if r.IsZero() {
		return `<uninitialized>`
	}
	return r.cast().Category()
}

/*
Valid returns an instance of error in the event the receiver is in
an aberrant state.
*/
func (r AttributeTypes) Valid() (err error) {
	if r.Kind() == `<uninitialized>` {
		err = nilInstanceErr(r)
	}

	return
}

/*
Keyword returns the [Keyword] associated with the receiver instance. In the context of this type instance, the [TargetKeyword] returned shall be either [TargetAttr] or [TargetFilter].
*/
func (r AttributeTypes) Keyword() Keyword {
	if r.Kind() == `<uri_search_attributes>` {
		return TargetFilter
	}

	return matchTKW(r.Kind())
}

/*
transfer will "copy" all slice references from the receiver instance into dest instance. PushPolicy controls may apply.
*/
func (r AttributeTypes) transfer(dest AttributeTypes) {
	_r := r.cast()
	_d := dest.cast()
	_r.Transfer(_d)
}

/*
Pop wraps the [stackage.Stack.Pop] method.
*/
func (r AttributeTypes) Pop() (x AttributeType) {
	z, _ := r.cast().Pop()
	if assert, asserted := z.(AttributeType); asserted {
		x = assert
	}

	return
}

/*
Push wraps the [stackage.Stack.Push] method. Valid input types are string and [AttributeType]. In the case of a string value, it is automatically cast as an instance of [AttributeType], so long as the raw string is of a non-zero length.
*/
func (r AttributeTypes) Push(x ...any) AttributeTypes {
	_r := r.cast()
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case string:
			_r.Push(AT(tv))
		default:
			_r.Push(tv)
		}
	}
	return r
}

/*
pushPolicy conforms to the PushPolicy interface signature defined within [stackage]. This private function is called during Push attempts to a [AttributeTypes] stack instance.
*/
func (r AttributeTypes) pushPolicy(x ...any) (err error) {
	// verify uniqueness; bail out if Boolean
	// false is return value.
	if r.contains(x[0]) {
		err = pushErrorNotUnique(r, x[0], matchTKW(r.Kind()))
		return
	}

	// perform type switch upon input value
	// x to determine suitability for push.
	switch tv := x[0].(type) {
	case AttributeType:
		// case matches an AttributeType instance
		if tv.IsZero() {
			err = pushErrorNilOrZero(r, tv, matchTKW(r.Kind()))
		}

	default:
		// unsuitable candidate per type
		err = pushErrorBadType(r, tv, matchTKW(r.Kind()))
	}

	return
}

/*
Contains returns a Boolean value indicative of whether value x, if a string or [AttributeType] instance, already resides within the receiver instance.

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
	}

	for i := 0; i < r.Len(); i++ {
		// case is not significant here.
		if eq(r.Index(i).String(), candidate) {
			return true
		}
	}

	return false
}

/*
TAs returns a freshly initialized instance of [AttributeTypes], configured to store one (1) or more [AttributeType] instances for the purpose of [TargetRule] expression when using the [TargetAttr] [TargetKeyword] context.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the [AttributeType] type (or its string equivalent) during initialization. This is merely a more convenient alternative to separate initialization and push procedures.

Values are automatically delimited using the [stackage.Stack.Symbol] method using the symbolic OR operator (`||`).
*/
func TAs(x ...any) (a AttributeTypes) {
	_a := stackOr().
		Symbol(`||`).
		NoNesting(true).
		SetID(targetRuleID).
		NoPadding(!StackPadding).
		SetCategory(TargetAttr.String()).
		SetPushPolicy(a.pushPolicy)

	a = AttributeTypes(_a).Push(x...)
	return
}

/*
UAs returns a freshly initialized instance of [AttributeTypes], configured to store one (1) or more [AttributeType] instances for the purpose of LDAP Search URI specification of desired [AttributeType] names. Instances of this design are not generally needed elsewhere.

Optionally, the caller may choose to push one (1) or more (valid) instances of the [AttributeType] type (or its string equivalent) during initialization. This is merely a more convenient alternative to separate initialization and push procedures.

Values are automatically comma-delimited (ASCII #44) using the [stackage.Stack.SetDelimiter] method in List mode.
*/
func UAs(x ...any) (a AttributeTypes) {
	_a := stackList().
		NoNesting(true).
		NoPadding(true).
		SetID(targetRuleID).
		SetDelimiter(rune(44)).
		SetCategory(`<uri_search_attributes>`). // URIs qualify for a few different KWs.
		SetPushPolicy(a.pushPolicy)

	a = AttributeTypes(_a).Push(x...)
	return
}

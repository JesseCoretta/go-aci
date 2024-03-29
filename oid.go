package aci

import (
	"github.com/JesseCoretta/go-objectid"
)

const badDotNot = `<invalid_object_identifier>`

var badOID ObjectIdentifier

/*
ObjectIdentifierContext is a convenient interface type that is qualified by the following types:

  - [ObjectIdentifier]
  - [ObjectIdentifiers]

The qualifying methods shown below are intended to make the generalized handling of ASN.1 object identifiers slightly easier without an absolute need for type assertion at every step. These methods are inherently read-only in nature.

To alter the underlying value, or to gain access to all of a given type's methods, type assertion of qualifying instances shall be necessary.
*/
type ObjectIdentifierContext interface {
	Len() int
	String() string
	Kind() string
	Keyword() Keyword
	IsZero() bool
	Valid() error

	isObjectIdentifierContext()
}

/*
ObjectIdentifier embeds an instance of go-objectid's DotNotation type.

Within the context of this package, instances of this type are used mainly for [TargetRule] definitions that bear the [TargetCtrl] or [TargetExtOp] [TargetKeyword] contexts.
*/
type ObjectIdentifier struct {
	*objectIdentifier
}

/*
objectIdentifier is the private embedded instance (BY POINTER) for storage within an ObjectIdentifier. It houses the actual [objectid.DotNotation] (pointer!) instance, as well as either the [TargetExtOp] or [TargetCtrl] [TargetKeyword] context.
*/
type objectIdentifier struct {
	TargetKeyword
	*objectid.DotNotation
}

/*
String wraps [objectid.DotNotation.String] method.
*/
func (r ObjectIdentifier) String() string {
	if err := r.Valid(); err != nil {
		return badDotNot
	}
	return r.objectIdentifier.DotNotation.String()
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r ObjectIdentifier) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Len returns 0 or 1 to describe an abstract length of the receiver. This method exists only to satisfy Go's interface signature requirements and need not be used.
*/
func (r ObjectIdentifier) Len() int {
	if err := r.Valid(); err != nil {
		return 0
	}
	return 1
}

func (r ObjectIdentifier) isObjectIdentifierContext() {}

/*
Keyword returns the [TargetKeyword] assigned to the receiver instance enveloped as a [Keyword] instance. This shall be the [TargetKeyword] that appears in a [TargetRule] containing the receiver instance as the expression value.
*/
func (r ObjectIdentifier) Keyword() Keyword {
	if r.IsZero() {
		return nil
	}
	return r.objectIdentifier.TargetKeyword
}

/*
Kind returns the string value for the kind of [ObjectIdentifier] described by the receiver. It will return either [TargetCtrl] or [TargetExtOp].
*/
func (r ObjectIdentifier) Kind() string {
	if r.IsZero() {
		return `<uninitialized>`
	}
	return r.objectIdentifier.TargetKeyword.String()
}

/*
TRM returns an instance of [TargetRuleMethods].

Each of the return instance's key values represent a single [ComparisonOperator] that is allowed for use in the creation of [TargetRule] instances which bear the receiver instance as an expression value. The value for each key is the actual instance method to -- optionally -- use for the creation of the [TargetRule].

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus [TargetRule] instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly populated and prepared for such activity.
*/
func (r ObjectIdentifier) TRM() TargetRuleMethods {
	return newTargetRuleMethods(targetRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
IsZero wraps [objectid.DotNotation.IsZero] method.
*/
func (r ObjectIdentifier) IsZero() bool {
	if r.objectIdentifier == nil {
		return true
	} else if r.objectIdentifier.DotNotation == nil {
		return true
	}

	return r.objectIdentifier.TargetKeyword == TargetKeyword(0x0)
}

/*
Eq initializes and returns a new [TargetRule] instance configured to express the evaluation of the receiver value as Equal-To a [TargetExtOp] or [TargetCtrl] [TargetKeyword] context.
*/
func (r ObjectIdentifier) Eq() (t TargetRule) {
	t = badTargetRule
	if err := r.Valid(); err == nil {
		t = TR(r.objectIdentifier.TargetKeyword, Eq, r)
	}
	return
}

/*
Ne initializes and returns a new [TargetRule] instance configured to express the evaluation of the receiver value as Not-Equal-To a [TargetExtOp] or [TargetCtrl] [TargetKeyword] context.

Negated equality [TargetRule] instances should be used with caution.
*/
func (r ObjectIdentifier) Ne() (t TargetRule) {
	t = badTargetRule
	if err := r.Valid(); err == nil {
		t = TR(r.objectIdentifier.TargetKeyword, Ne, r)
	}
	return
}

/*
Valid returns an instance of error in the event the receiver is in an aberrant state.
*/
func (r ObjectIdentifier) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
		return
	}

	/*
		raw := r.objectIdentifier.DotNotation.String()
		if !isDotNot(raw) {
			err = badObjectIdentifierErr(raw)
			return
		}

		if !(r.objectIdentifier.DotNotation.Len() > 0 &&
			r.objectIdentifier.TargetKeyword != TargetKeyword(0x0)) {
			err = badObjectIdentifierKeywordErr(r.objectIdentifier.TargetKeyword)
		}
	*/

	return
}

/*
OID returns a new instance of [ObjectIdentifier]. Instances of this type are used in the following scenarios:

• For appending (by Push) to an instance of [ObjectIdentifiers] when crafting a multi-valued [TargetRule] instance containing a sequence of [ObjectIdentifier] instances, OR ...
• For use directly in creating a unary [TargetRule] instance

In either case, valid [TargetKeyword] contexts are [TargetExtOp] and [TargetCtrl].
*/
func OID(x any, kw Keyword) ObjectIdentifier {
	o, _ := newObjectID(kw.(TargetKeyword), x)
	return ObjectIdentifier{o}
}

/*
Ctrl initializes a new instance of [ObjectIdentifier], which embeds an instance of go-objectid's DotNotation type.

Instances of this design are used in the creation of [TargetRule] instances that bear the [TargetCtrl] [TargetKeyword] context. OIDs produced as a result of this function are generally expected to be LDAP Control Object Identifiers.
*/
func Ctrl(x ...any) ObjectIdentifier {
	o, _ := newObjectID(TargetCtrl, x...)
	return ObjectIdentifier{o}
}

/*
Extop initializes a new instance of [ObjectIdentifier], which embeds an instance of go-objectid's DotNotation type.

Instances of this design are used in the creation of [TargetRule] instances that bear the [TargetExtOp] [TargetKeyword] context. OIDs produced as a result of this function are generally expected to be LDAP Extended Operation Object Identifiers.
*/
func ExtOp(x ...any) ObjectIdentifier {
	o, _ := newObjectID(TargetExtOp, x...)
	return ObjectIdentifier{o}
}

/*
set is a private method executed by Set.
*/
func (r *objectIdentifier) set(x ...any) (err error) {
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case ObjectIdentifier:
			if err = tv.Valid(); err == nil {
				r.DotNotation, err = objectid.NewDotNotation(tv.String())
			}
		case string:
			r.DotNotation, err = objectid.NewDotNotation(tv)
		}
	}
	return
}

/*
newObjectID is a private function called by ExtOp and Ctrl package level functions.
*/
func newObjectID(kw TargetKeyword, x ...any) (o *objectIdentifier, err error) {
	o = new(objectIdentifier)
	if err = o.set(x...); err != nil {
		return
	}
	o.TargetKeyword = kw
	return
}

func isDotNot(x string) bool {
	o, err := objectid.NewDotNotation(x)
	return err == nil && o != nil
}

/*
TRM returns an instance of [TargetRuleMethods].

Each of the return instance's key values represent a single [ComparisonOperator] that is allowed for use in the creation of [TargetRule] instances which bear the receiver instance as an expression value. The value for each key is the actual instance method to -- optionally -- use for the creation of the [TargetRule].

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus [TargetRule] instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly populated and prepared for such activity.
*/
func (r ObjectIdentifiers) TRM() TargetRuleMethods {
	return newTargetRuleMethods(targetRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
IsZero wraps the [stackage.Stack.IsZero] method.
*/
func (r ObjectIdentifiers) IsZero() bool {
	return r.cast().IsZero()
}

/*
Valid returns an instance of error in the event the receiver is in an aberrant state.
*/
func (r ObjectIdentifiers) Valid() (err error) {
	if r.Kind() == `<uninitialized>` {
		err = nilInstanceErr(r)
	}

	return
}

/*
Len wraps the [stackage.Stack.Len] method.
*/
func (r ObjectIdentifiers) Len() int {
	return r.cast().Len()
}

/*
Index wraps the [stackage.Stack.Index] method. Note that the Boolean OK value returned by [stackage] by default will be shadowed and not obtainable by the caller.
*/
func (r ObjectIdentifiers) Index(idx int) (x ObjectIdentifier) {
	y, _ := r.cast().Index(idx)
	if assert, ok := y.(ObjectIdentifier); ok {
		x = assert
	}
	return
}

func (r ObjectIdentifiers) isObjectIdentifierContext() {}

/*
String is a stringer method that returns the string representation of the receiver instance.

This method wraps the [stackage.Stack.String] method.
*/
func (r ObjectIdentifiers) String() string {
	return r.cast().String()
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r ObjectIdentifiers) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Push wraps the [stackage.Stack.Push] method.
*/
func (r ObjectIdentifiers) Push(x ...any) ObjectIdentifiers {
	_r := r.cast()

	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case string:
			_r.Push(r.F()(tv))
		default:
			_r.Push(tv)
		}
	}

	return r
}

/*
Keyword returns the [TargetKeyword] associated with the receiver instance enveloped as a [Keyword]. In the context of this type instance, the [TargetKeyword] returned is always [TargetExtOp] or [TargetCtrl].
*/
func (r ObjectIdentifiers) Keyword() Keyword {
	if r.IsZero() {
		return TargetKeyword(0x0)
	}

	kw := matchTKW(r.Kind())
	return kw
}

/*
Contains returns a Boolean value indicative of whether value x, if a string or [ObjectIdentifier] instance, already resides within the receiver instance.

Case is not significant in the matching process.
*/
func (r ObjectIdentifiers) Contains(x any) bool {
	return r.contains(x)
}

/*
contains is a private method called by ObjectIdentifiers.Contains.
*/
func (r ObjectIdentifiers) contains(x any) bool {
	if r.Len() == 0 {
		return false
	}

	var candidate string

	switch tv := x.(type) {
	case string:
		candidate = tv
	case ObjectIdentifier:
		candidate = tv.String()
	default:
		return false
	}

	if len(candidate) == 0 || candidate == badDotNot {
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

/*
Pop wraps the [stackage.Stack.Pop] method.
*/
func (r ObjectIdentifiers) Pop() (x ObjectIdentifier) {
	y, _ := r.cast().Pop()
	if assert, asserted := y.(ObjectIdentifier); asserted {
		x = assert
	}

	return
}

/*
setQuoteStyle shall set the receiver instance to the quotation scheme defined by integer i.
*/
func (r ObjectIdentifiers) setQuoteStyle(style int) ObjectIdentifiers {
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
setExpressionValues is a private method called by assertTargetTFDN for DN-based Target Rules and assertBindUGRDN for DN-based Bind Rules.
*/
func (r ObjectIdentifiers) setExpressionValues(key TargetKeyword, values ...string) (err error) {
	// iterate each string-based distinguishedName
	// in the values sequence ...
	for i := 0; i < len(values); i++ {

		// Identify this objectIdentifier value
		// as O, as referenced by index integer i.
		O := values[i]

		// Attempt to parse the raw object Identifier
		// (O) dot notation value using go-objectid.
		// Bail if ANY errors are encountered. Honor
		// the keyword in use.
		var o *objectIdentifier
		if o, err = newObjectID(key, O); err != nil {
			return
		}

		// accept the new OID value (o), pushing it
		// into the receiver instance and embedding
		// into an ObjectIdentifier struct envelope.
		r.Push(ObjectIdentifier{o})
	}

	return
}

/*
extOpsPushPolicy conforms to the PushPolicy signature defined within [stackage].  This function will be called privately whenever an instance is pushed into a particular stackage.Stack (or alias) type instance when the [TargetKeyword] `extop` is involved.

Only ObjectIdentifier instances are to be cleared for push executions, assuming they are keyword context-aligned with the destination stack.
*/
func (r ObjectIdentifiers) extOpsPushPolicy(x ...any) error {
	if r.contains(x[0]) {
		return pushErrorNotUnique(r, x[0], r.Keyword())
	}
	return objectIdentifiersPushPolicy(r, x[0], TargetExtOp)
}

/*
ctrlsPushPolicy conforms to the PushPolicy signature defined within [stackage].  This function will be called privately whenever an instance is pushed into a particular stackage.Stack (or alias) type instance when the [TargetKeyword] `targetcontrol` is involved.

Only ObjectIdentifier instances are to be cleared for push executions, assuming they are keyword context-aligned with the destination stack.
*/
func (r ObjectIdentifiers) ctrlsPushPolicy(x ...any) error {
	if r.contains(x[0]) {
		return pushErrorNotUnique(r, x[0], r.Keyword())
	}
	return objectIdentifiersPushPolicy(r, x[0], TargetCtrl)
}

/*
F returns the appropriate instance creator function for crafting individual [ObjectIdentifier] instances for submission to the receiver. This is merely a convenient alternative to maintaining knowledge as to which function applies to the current receiver instance.

The default is [Ctrl], and will be returned if the receiver is uninitialized, or if the [TargetKeyword] associated with the receiver is invalid somehow. Otherwise, [ExtOp] is returned.
*/
func (r ObjectIdentifiers) F() func(...any) ObjectIdentifier {
	switch r.Keyword() {
	case TargetExtOp:
		return ExtOp
	}

	return Ctrl
}

func (r ObjectIdentifiers) reset() {
	r.cast().Reset()
}

func (r ObjectIdentifiers) resetKeyword(x any) {
	if r.Len() == 0 {
		switch tv := x.(type) {
		case TargetKeyword:
			r.resetKeyword(tv.String())

		case string:
			_r := r.cast()

			switch lc(tv) {
			case TargetExtOp.String():
				_r.SetCategory(TargetExtOp.String()).
					SetPushPolicy(r.extOpsPushPolicy)

			case TargetCtrl.String():
				_r.SetCategory(TargetCtrl.String()).
					SetPushPolicy(r.ctrlsPushPolicy)
			}
		}
	}
}

/*
objectIdentifiersPushPolicy is called by one of the PushPolicy conformant interface signature functions -- either extOpsPushPolicy or ctrlsPushPolicy -- and is used to determine whether or not an element being pushed (into a stack, somewhere) should be accepted based on the keyword context.
*/
func objectIdentifiersPushPolicy(r, x any, kw TargetKeyword) (err error) {
	switch tv := x.(type) {

	case string:

		// case match is a string-based objectIdentifier
		if !isDotNot(tv) {
			err = badObjectIdentifierErr(tv)
		}

	case TargetKeyword:
		switch tv {
		case TargetExtOp, TargetCtrl:
			if R := r.(ObjectIdentifiers); R.Len() == 0 {
				R.resetKeyword(tv)
			}
		}

	case ObjectIdentifier:

		// case match is a proper instance of ObjectIdentifier
		if err = tv.Valid(); err != nil {
			break
		}
		err = badObjectIdentifierKeywordErr(tv.objectIdentifier.TargetKeyword)
		if tv.Keyword() == kw {
			err = nil
		}

	default:
		err = pushErrorBadType(r, x, kw)
	}

	return
}

/*
ExtOps returns a freshly initialized instance of [ObjectIdentifiers], configured to store one (1) or more [ObjectIdentifier] instances for the purpose of [TargetRule] expression when using the [TargetExtOp] [TargetKeyword].

Optionally, the caller may choose to submit one (1) or more (valid) instances of the [ObjectIdentifier] type. This is merely a more convenient alternative to separate init and push procedures.
*/
func ExtOps(x ...any) (o ObjectIdentifiers) {
	// create a native stackage.Stack
	// and configure before typecast.
	_o := stackOr().
		Symbol(`||`).
		NoNesting(true).
		SetID(targetRuleID).
		NoPadding(!StackPadding).
		SetCategory(TargetExtOp.String())

	// cast _o as a proper ObjectIdentifiers
	// instance (o). We do it this way to gain
	// access to the method for the *specific
	// instance* being created (o), thus allowing
	// things like uniqueness checks, etc., to
	// occur during push attempts, providing more
	// helpful and non-generalized feedback.
	o = ObjectIdentifiers(_o)
	_o.SetPushPolicy(o.extOpsPushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	o.Push(x...)

	return
}

/*
ExtOps returns a freshly initialized instance of [ObjectIdentifiers], configured to store one (1) or more [ObjectIdentifier] instances for the purpose of [TargetRule] expression when using the [TargetCtrl] [TargetKeyword].

Optionally, the caller may choose to submit one (1) or more (valid) instances of the [ObjectIdentifier] type. This is merely a more convenient alternative to separate init and push procedures.
*/
func Ctrls(x ...any) (o ObjectIdentifiers) {
	// create a native stackage.Stack
	// and configure before typecast.
	_o := stackOr().
		Symbol(`||`).
		NoNesting(true).
		SetID(targetRuleID).
		NoPadding(!StackPadding).
		SetCategory(TargetCtrl.String())

	// cast _o as a proper ObjectIdentifiers
	// instance (o). We do it this way to gain
	// access to the method for the *specific
	// instance* being created (o), thus allowing
	// things like uniqueness checks, etc., to
	// occur during push attempts, providing more
	// helpful and non-generalized feedback.
	o = ObjectIdentifiers(_o)
	_o.SetPushPolicy(o.ctrlsPushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	o.Push(x...)

	return
}

/*
Eq initializes and returns a new [TargetRule] instance configured to express the evaluation of the receiver value as Equal-To a [TargetExtOp] or [TargetCtrl] [TargetKeyword] context.
*/
func (r ObjectIdentifiers) Eq() TargetRule {
	if err := r.Valid(); err != nil {
		return badTargetRule
	}
	return TR(r.Kind(), Eq, r)
}

/*
Ne initializes and returns a new [TargetRule] instance configured to express the evaluation of the receiver value as Not-Equal-To a [TargetExtOp] or [TargetCtrl] [TargetKeyword] context.

Negated equality [TargetRule] instances should be used with caution.
*/
func (r ObjectIdentifiers) Ne() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}
	return TR(r.Kind(), Ne, r)
}

/*
ID wraps the [stackage.Stack.ID] method.
*/
func (r ObjectIdentifiers) ID() (s string) {
	if !r.IsZero() {
		s = r.cast().ID()
	}
	return
}

/*
Category wraps the [stackage.Stack.Category] method.
*/
func (r ObjectIdentifiers) Kind() (k string) {
	k = `<uninitialized>`
	if r.IsZero() {
		return
	}

	switch _k := lc(r.cast().Category()); _k {
	case TargetExtOp.String(),
		TargetCtrl.String():
		k = _k
	}

	return
}

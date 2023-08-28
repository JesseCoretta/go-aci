package aci

import (
	"github.com/JesseCoretta/go-objectid"
	"github.com/JesseCoretta/go-stackage"
)

const badDotNot = `<invalid_object_identifier>`

var badOID ObjectIdentifier

/*
ObjectIdentifier embeds an instance of go-objectid's
DotNotation type.

Within the context of this package, instances of this
type are used mainly for Target Rule definitions that
bear the targetcontrol or extop keywords.
*/
type ObjectIdentifier struct {
	*objectIdentifier
}

/*
objectIdentifier is the private embedded instance (BY
POINTER) for storage within an ObjectIdentifier. It
houses the actual *objectid.DotNotation instance, as
well as either the TargetExtOp keyword, or its mate
keyword, TargetControl.
*/
type objectIdentifier struct {
	TargetKeyword
	*objectid.DotNotation
}

/*
String wraps go-objectid's DotNotation.String method.
*/
func (r ObjectIdentifier) String() string {
	if err := r.Valid(); err != nil {
		return badDotNot
	}
	return r.objectIdentifier.DotNotation.String()
}

/*
IsZero wraps go-objectid's DotNotation.IsZero method.
*/
func (r ObjectIdentifier) IsZero() bool {
	if r.DotNotation == nil {
		return true
	}

	return r.objectIdentifier.DotNotation.IsZero() &&
		r.objectIdentifier.TargetKeyword == TargetKeyword(0x0)
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To an `extop` or `targetcontrol`
keyword context.
*/
func (r ObjectIdentifier) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(r.objectIdentifier.TargetKeyword)
	t.SetOperator(Eq)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!RulePadding).
		SetCategory(r.objectIdentifier.TargetKeyword.String())

	t = TargetRule(*_t)
	return t
}

/*
Ne initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Not-Equal-To an `extop` or `targetcontrol`
keyword context.

Negated equality TargetRule instances should be used with caution.
*/
func (r ObjectIdentifier) Ne() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(r.objectIdentifier.TargetKeyword)
	t.SetOperator(Ne)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!RulePadding).
		SetCategory(r.objectIdentifier.TargetKeyword.String())

	t = TargetRule(*_t)
	return t
}

/*
Valid returns an instance of error in the event the receiver is in
an aberrant state.
*/
func (r ObjectIdentifier) Valid() (err error) {
	if r.IsZero() {
		err = errorf("%T instance is nil", r)
		return
	}

	if !(r.objectIdentifier.DotNotation.Len() > 0 &&
		r.objectIdentifier.TargetKeyword != TargetKeyword(0x0)) {
		err = errorf("Invalid %T and/or %T value(s)",
			r.objectIdentifier.DotNotation,
			r.objectIdentifier.TargetKeyword)
	}

	return
}

/*
OID returns a new instance of ObjectIdentifier. Instances of
this type are used in the following scenarios:

• For appending (by Push) to an instance of ObjectIdentifers
when crafting a multi-valued TargetRule instance containing
a sequence of ObjectIdentifier instances, OR ...

• For use directly in creating a unary TargetRule instance

In either case, valid Target keyword contexts are `extop` and
`targetcontrol`.
*/
func OID(x any, kw Keyword) ObjectIdentifier {
	o, _ := newObjectID(kw.(TargetKeyword), x)
	return ObjectIdentifier{o}
}

/*
Ctrl initializes a new instance of ObjectIdentifier, which
embeds an instance of go-objectid's DotNotation type.

Instances of this design are used in the creation of Target
Rule Conditions that bear the `targetcontrol` keyword. OIDs
produced as a result of this function are generally expected
to be LDAP Control Object Identifiers.
*/
func Ctrl(x ...any) ObjectIdentifier {
	o, _ := newObjectID(TargetCtrl, x...)
	return ObjectIdentifier{o}
}

/*
ExtOp initializes a new instance of ObjectIdentifier, which
embeds an instance of go-objectid's DotNotation type.

Instances of this design are used in the creation of Target
Rule Conditions that bear the `extop` keyword. OIDs produced
as a result of this function are generally expected to be LDAP
Extended Operation Object Identifiers.
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
		case string:
			r.DotNotation, err = objectid.NewDotNotation(tv)
		}
	}
	return
}

/*
newObjectID is a private function called by ExtOp and Ctrl package
level functions.
*/
func newObjectID(kw TargetKeyword, x ...any) (o *objectIdentifier, err error) {
	o = new(objectIdentifier)
	o.set(x...)
	o.TargetKeyword = kw
	return
}

func isDotNot(x string) bool {
	o, err := objectid.NewDotNotation(x)
	return err == nil && o.Valid()
}

/*
ObjectIdentifiers is an alias type for stackage.Stack, and is intended
to house one (1) or more ObjectIdentifier instances for the purpose of
expression within a TargetRule instance.
*/
type ObjectIdentifiers stackage.Stack

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r ObjectIdentifiers) IsZero() bool {
	_r, _ := castAsStack(r)
	return _r.IsZero()
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r ObjectIdentifiers) Len() int {
	_r, _ := castAsStack(r)
	return _r.Len()
}

/*
Index wraps go-stackage's Stack.Index method. Note that the
Boolean OK value returned by go-stackage by default will be
shadowed and not obtainable by the caller.
*/
func (r ObjectIdentifiers) Index(idx int) (x ObjectIdentifier) {
	_r, _ := castAsStack(r)
	y, _ := _r.Index(idx)

	if assert, ok := y.(ObjectIdentifier); ok {
		x = assert
	}
	return
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Stack.String method.
*/
func (r ObjectIdentifiers) String() string {
	_r, _ := castAsStack(r)
	return _r.String()
}

/*
Push wraps go-stackage's Stack.Push method.
*/
func (r ObjectIdentifiers) Push(x ...any) ObjectIdentifiers {
	_r, _ := castAsStack(r)
	_r.Push(x...)
	return ObjectIdentifiers(_r)
}

func (r ObjectIdentifiers) Keyword() Keyword {
	if r.IsZero() {
		return TargetKeyword(0x0)
	}

	kw := matchTKW(r.Category())
	return kw
}

/*
Contains returns a boolean value indicative of whether value x,
if a string or ObjectIdentifier instance, already resides
within the receiver instance.

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

/*
Pop wraps go-stackage's Stack.Pop method.
*/
func (r ObjectIdentifiers) Pop() (x ObjectIdentifier) {
	_r, _ := castAsStack(r)
	y, _ := _r.Pop()

	if assert, asserted := y.(ObjectIdentifier); asserted {
		x = assert
	}

	return
}

/*
setQuoteStyle shall set the receiver instance to the quotation
scheme defined by integer i.
*/
func (r ObjectIdentifiers) setQuoteStyle(style int) {
	_r, _ := castAsStack(r)
	if _r.Len() < 2 {
		_r.Encap() // not multivalued, force default
		return
	}

	if style == MultivalSliceQuotes {
		_r.Encap(`"`)
	} else {
		_r.Encap()
	}

	return
}

/*
setExpressionValues is a private method called by assertTargetTFDN for
DN-based Target Rules and assertBindUGRDN for DN-based Bind Rules.
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
extOpsPushPolicy conforms to the PushPolicy signature
defined within go-stackage.  This function will be called
privately whenever an instance is pushed into a particular
stackage.Stack (or alias) type instance when the Target
keyword `extop` is involved.

Only ObjectIdentifier instances are to be cleared for push
executions, assuming they are keyword context-aligned with
the destination stack.
*/
func (r ObjectIdentifiers) extOpsPushPolicy(x any) error {
	if r.contains(x) {
		err := errorf("Cannot push non-unique or invalid %T into %T [%s]",
			x, r, r.Keyword())
		return err
	}
	return objectIdentifiersPushPolicy(r, x, TargetExtOp)
}

/*
ctrlsPushPolicy conforms to the PushPolicy signature
defined within go-stackage.  This function will be called
privately whenever an instance is pushed into a particular
stackage.Stack (or alias) type instance when the Target
keyword `targetcontrol` is involved.

Only ObjectIdentifier instances are to be cleared for push
executions, assuming they are keyword context-aligned with
the destination stack.
*/
func (r ObjectIdentifiers) ctrlsPushPolicy(x any) error {
	if r.contains(x) {
		err := errorf("Cannot push non-unique or invalid %T into %T [%s]",
			x, r, r.Keyword())
		return err
	}
	return objectIdentifiersPushPolicy(r, x, TargetCtrl)
}

/*
objectIdentifiersPushPolicy is called by one of the PushPolicy
conformant interface signature functions -- either extOpsPushPolicy
or ctrlsPushPolicy -- and is used to determine whether or not an
element being pushed (into a stack, somewhere) should be accepted
based on the keyword context.
*/
func objectIdentifiersPushPolicy(r, x any, kw TargetKeyword) (err error) {
	switch tv := x.(type) {

	case string:

		// case match is a string-based objectIdentifier
		if len(tv) == 0 {
			err = errorf("Cannot push zero string %T into %T [%s]", tv, r, kw)
		}

	case ObjectIdentifier:

		// case match is a proper instance of ObjectIdentifier
		if tv.IsZero() {
			err = errorf("Cannot push nil %T into %T [%s]", tv, r, kw)

		} else if tv.objectIdentifier.TargetKeyword != kw {
			err = errorf("%T push into %T failed during keyword verification (not a %s-based %T)",
				tv, r, kw, tv)
		}

	default:
		err = errorf("Push request of %T type violates %T [%s] PushPolicy",
			tv, r, kw)
	}

	return
}

func badObjectIdentifierErr(x string) error {
	return errorf("Invalid ObjectIdentifier instance: must conform to 'N[.N]+', got '%s'", x)
}

/*
ExtOps returns a freshly initialized instance of ObjectIdentifiers, configured
to store one (1) or more ObjectIdentifier instances for the purpose of Target
Rule expression when using the extop keyword.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
ObjectIdentifier type. This is merely a more convenient alternative to separate init
and push procedures.
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
	_o.Push(x...)

	return
}

/*
Ctrls returns a freshly initialized instance of ObjectIdentifiers, configured
to store one (1) or more ObjectIdentifier instances for the purpose of Target
Rule expression when using the targetcontrol keyword.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
ObjectIdentifier type. This is merely a more convenient alternative to separate init
and push procedures.
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
	_o.Push(x...)

	return
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To an `extop` or `targetcontrol`
keyword context.
*/
func (r ObjectIdentifiers) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(r.Category())
	t.SetOperator(Eq)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!RulePadding).
		SetCategory(t.Keyword().String())

	t = TargetRule(*_t)
	return t
}

/*
Ne initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Not-Equal-To an `extop` or `targetcontrol`
keyword context.

Negated equality TargetRule instances should be used with caution.
*/
func (r ObjectIdentifiers) Ne() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}

	var t TargetRule
	t.SetKeyword(r.Category())
	t.SetOperator(Ne)
	t.SetExpression(r)

	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!RulePadding).
		SetCategory(t.Keyword().String())

	t = TargetRule(*_t)
	return t
}

/*
ID wraps go-stackage's Stack.ID method.
*/
func (r ObjectIdentifiers) ID() string {
	if r.IsZero() {
		return ``
	}

	_t, _ := castAsStack(r)
	return _t.ID()
}

/*
Category wraps go-stackage's Stack.Category method.
*/
func (r ObjectIdentifiers) Category() string {
	if r.IsZero() {
		return ``
	}

	_t, _ := castAsStack(r)
	return _t.Category()
}

/*
setID wraps go-stackage's Stack.SetID method.
*/
func (r ObjectIdentifiers) setID(id string) {
	_t, _ := castAsStack(r)
	_t.SetID(id)
}

/*
setCategory wraps go-stackage's Stack.SetCategory method.
*/
func (r ObjectIdentifiers) setCategory(cat string) {
	_t, _ := castAsStack(r)
	_t.SetCategory(cat)
}

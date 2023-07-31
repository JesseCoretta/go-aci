package aci

import (
	"github.com/JesseCoretta/go-objectid"
)

const badDotNot = `<invalid_object_identifier>`

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
Eq initializes and returns a new Condition instance configured to express
the evaluation of the receiver value as Equal-To a `targetcontrol` or `extop`
Target Rule, depending on the receiver's configuration.
*/
func (r ObjectIdentifier) Eq() Condition {
	if r.IsZero() {
		return Condition{}
	}
	return Cond(r.objectIdentifier.TargetKeyword, r, Eq).
		Encap(`"`).
		Paren().
		setID(`target`).
		setCategory(r.objectIdentifier.TargetKeyword.String())
}

/*
Ne initializes and returns a new Condition instance configured to express
the evaluation of the receiver value as Not-Equal-To a `targetcontrol` or
`extop` Target Rule, depending on the receiver's configuration.
*/
func (r ObjectIdentifier) Ne() Condition {
        if r.IsZero() {
                return Condition{}
        }
        return Cond(r.objectIdentifier.TargetKeyword, r, Ne).
		Encap(`"`).
		Paren().
		setID(`target`).
		setCategory(r.objectIdentifier.TargetKeyword.String())
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

	if !( r.objectIdentifier.DotNotation.Len() > 0 &&
		r.objectIdentifier.TargetKeyword != TargetKeyword(0x0) ) {
		err = errorf("Invalid %T and/or %T value(s)",
			r.objectIdentifier.DotNotation,
			r.objectIdentifier.TargetKeyword)
	}

	return
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
	o, _ := newObjectID(TargetCtrl,x...)
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
        o, _ := newObjectID(TargetExtOp,x...)
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

// future parsers go here...

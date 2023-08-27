package aci

/*
pb.go contains PermissionBindRules types and methods. PermissionBindRules combine
Permissions with BindRules, and are a key component in the formation of a complete
ACI.
*/

import (
	"github.com/JesseCoretta/go-stackage"
)

/*
invalid value constants used as stringer method returns when
something goes wrong :/
*/
const (
	// badPB is supplied during PermissionBindRule string representation
	// attempts that fail
	badPB = `<invalid_pbrule>`
)

/*
PB contains one (1) Permission instance and one (1) BindRules instance. Instances
of this type are used to create top-level ACI Permission+Bind Rule pairs, of which
at least one (1) is required in any given access control instructor definition.

Users may compose instances of this type manually, or using the PB package
level function, which automatically invokes value checks.
*/
type PermissionBindRule struct {
	P Permission
	B BindRules
}

/*
PBR returns an instance of PermissionBindRule, bearing the Permission P and
the Bind Rule B. The values P and B shall undergo validity checks per the
conditions of the PermissionBindRule.Valid method automatically.

Instances of this kind are intended for submission (via Push) into instances
of PermissionBindRules.

Generally, an ACI only has a single PermissionBindRule, though multiple
instances of this type are allowed per the syntax specification honored
by this package.
*/
func PBR(P Permission, B BindRules) (r PermissionBindRule) {
	_r := PermissionBindRule{P, B}
	if err := _r.Valid(); err == nil {
		r = _r
	}

	return
}

/*
Valid returns an error instance should any of the following conditions
evaluate as true:

• Permission.Valid returns an error for P

• Rule.Valid returns an error for B

• Rule.Len returns zero (0) for B
*/
func (r PermissionBindRule) Valid() (err error) {
	return r.valid()
}

func (r PermissionBindRule) IsZero() bool {
	return r.P.IsZero() && r.B.IsZero()
}

func (r PermissionBindRule) Kind() string {
	return pbrRuleID
}

/*
valid is a private method invoked by PermissionBindRule.Valid.
*/
func (r PermissionBindRule) valid() (err error) {
	if err = r.P.Valid(); err != nil {
		return
	} else if err = r.B.Valid(); err != nil {
		return
	}

	if r.B.Len() == 0 {
		err = errorf("%T is zero length", r.B)
	} else if r.P.IsZero() || r.B.ID() != `bind` {
		err = errorf("%T is not a permission+bind rule (%s)", r.B, r.B.ID())
	}

	return
}

/*
ID wraps go-stackage's Stack.ID method.
*/
func (r PermissionBindRule) ID() string {
	return pbrRuleID
}

/*
Category wraps go-stackage's Stack.Category method.
*/
func (r PermissionBindRule) Category() string {
	return pbrRuleID
}

/*
String is a stringer method that returns the string representation
of the receiver.
*/
func (r PermissionBindRule) String() string {
	return r.string()
}

/*
string is a private method called by PermissionBindRule.String.
*/
func (r PermissionBindRule) string() (s string) {
	s = badPB
	if err := r.valid(); err == nil {
		s = sprintf("%s %s", r.P, r.B)
	}

	return
}

/*
PermissionBindRules is a stackage.Stack type alias used to store one (1)
or more instances of PermissionBindRule. Instances of this kind are used
in top-level Instruction (ACI) assembly.
*/
type PermissionBindRules stackage.Stack

/*
Valid wraps go-stackage's Stack.Valid method.
*/
func (r PermissionBindRules) Valid() (err error) {
	_t, _ := castAsStack(r)
	err = _t.Valid()
	return
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r PermissionBindRules) IsZero() bool {
	_r, _ := castAsStack(r)
	return _r.IsZero()
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r PermissionBindRules) Len() int {
	_r, _ := castAsStack(r)
	return _r.Len()
}

/*
Category wraps go-stackage's Stack.Category method.
*/
func (r PermissionBindRules) Category() string {
	return pbrsRuleID
}

/*
Category wraps go-stackage's Stack.ID method.
*/
func (r PermissionBindRules) ID() string {
	return pbrsRuleID
}

/*
Index wraps go-stackage's Stack.Index method and performs type
assertion in order to return an instance of PermissionBindRule.
*/
func (r PermissionBindRules) Index(idx int) (pbr PermissionBindRule) {
	_r, _ := castAsStack(r)
	x, _ := _r.Index(idx)

	if assert, asserted := x.(PermissionBindRule); asserted {
		pbr = assert
	}

	return
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Stack.String method.
*/
func (r PermissionBindRules) String() string {
	_r, _ := castAsStack(r)
	return _r.String()
}

/*
Push wraps go-stackage's Stack.Push method.
*/
func (r PermissionBindRules) Push(x ...any) PermissionBindRules {
	_r, _ := castAsStack(r)

	// iterate variadic input arguments
	for i := 0; i < len(x); i++ {

		if assert, ok := x[i].(PermissionBindRule); ok {
			// Push it!
			_r.Push(assert)
		}
	}

	r = PermissionBindRules(_r)
	return r
}

/*
Pop wraps go-stackage's Stack.Pop method. An instance of
PermissionBindRule, which may or may not be nil, is returned
following a call of this method.

Within the context of the receiver type, a PermissionBindRule,
if non-nil, can only represent a PermissionBindRule instance.
*/
func (r PermissionBindRules) Pop() (pbr PermissionBindRule) {
	_r, _ := castAsStack(r)
	x, _ := _r.Pop()

	if assert, ok := x.(PermissionBindRule); ok {
		pbr = assert
	}

	return
}

/*
permissionBindRulesPushPolicy conforms to the PushPolicy interface signature
defined within go-stackage. This private function is called during Push
attempts to a PermissionBindRules instance.
*/
func permissionBindRulesPushPolicy(x any) (err error) {
	switch tv := x.(type) {

	case PermissionBindRule:
		if tv.IsZero() {
			err = errorf("Cannot push nil %T into %T",
				tv, PermissionBindRules{})
		} else {
			if err = tv.Valid(); err != nil {
				err = errorf("Cannot push aberrant %T into %T: %v",
					tv, PermissionBindRules{}, err)
			}
		}

	default:
		err = errorf("Push request of %T type violates %T PushPolicy",
			tv, PermissionBindRules{})
	}

	return
}

/*
PBRs returns a freshly initialized instance of PermissionBindRules, configured to
store one (1) or more instances of PermissionBindRule.

Instances of this kind are used as a component in top-level Instruction (ACI)
assembly.
*/
func PBRs() PermissionBindRules {
	return PermissionBindRules(stackList().
		JoinDelim(`;`).
		SetID(pbrsRuleID).
		SetCategory(pbrsRuleID).
		NoPadding(!RulePadding).
		SetPushPolicy(permissionBindRulesPushPolicy))
}

const pbrRuleID = `pbr`
const pbrsRuleID = `pbrs`
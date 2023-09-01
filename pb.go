package aci

/*
pb.go contains PermissionBindRules types and methods. PermissionBindRules combine
Permissions with BindRules, and are a key component in the formation of a complete
ACI.
*/

/*
invalid value constants used as stringer method returns when
something goes wrong :/
*/
const (
	// badPB is supplied during PermissionBindRule string representation
	// attempts that fail
	badPB = `<invalid_pbrule>`
)

var (
	badPermissionBindRule  PermissionBindRule  // for failed calls that return a PermissionBindRule only
	badPermissionBindRules PermissionBindRules // for failed calls that return a PermissionBindRules only
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
	B BindContext // BindRule -or- BindRules are allowed
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
func PBR(P Permission, B BindContext) (r PermissionBindRule) {
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
func (r *PermissionBindRule) Valid() (err error) {
	return r.valid()
}

func (r *PermissionBindRule) IsZero() bool {
	if r == nil {
		return true
	}

	return r.P.IsZero() && r.B == nil
}

func (r PermissionBindRule) Kind() string {
	return pbrRuleID
}

/*
valid is a private method invoked by PermissionBindRule.Valid.
*/
func (r *PermissionBindRule) valid() (err error) {
	if r == nil {
		err = nilInstanceErr(r)
		return
	}

	if err = r.P.Valid(); err != nil {
		return

	} else if err = r.B.Valid(); err != nil {
		return
	}

	if r.B.Len() == 0 {
		err = nilInstanceErr(r.B)
	} else if r.P.IsZero() {
		err = nilInstanceErr(r.P)
	} else if r.B.ID() != bindRuleID {
		err = badPTBRuleKeywordErr(r.B, bindRuleID, bindRuleID, r.B.ID())
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
		s = sprintf("%s %s;", r.P, r.B)
	}

	return
}

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
defined within the go-stackage package. This private function is called during
Push attempts to a PermissionBindRules instance.
*/
func permissionBindRulesPushPolicy(x any) (err error) {
	switch tv := x.(type) {

	case PermissionBindRule:
		if tv.IsZero() {
			err = pushErrorNilOrZero(PermissionBindRules{}, tv, nil)
		} else {
			if err = tv.Valid(); err != nil {
				err = pushErrorNilOrZero(PermissionBindRules{}, tv, nil, err)
			}
		}

	default:
		err = pushErrorBadType(PermissionBindRules{}, tv, nil)
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
		SetID(pbrsRuleID).
		SetDelimiter(rune(59)).
		SetCategory(pbrsRuleID).
		NoPadding(!StackPadding).
		SetPushPolicy(permissionBindRulesPushPolicy))
}

const pbrRuleID = `pbr`
const pbrsRuleID = `pbrs`

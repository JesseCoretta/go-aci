package aci

/*
aci.go contains the top-level access control instructor methods and types.
*/

/*
Version defines the official ACI syntax version number implemented and honored by this package.
*/
const Version float32 = 3.0

/*
Instruction conforms to the ACI syntax specification associated with
the Version constant value of this package.

Instances of this type, when represented in their string form, are intended for
submission to an X.500/LDAP DSA for assignment (via the 'aci' LDAP Attribute Type)
to the relevant directory entry.
*/
type Instruction struct {
	*instruction
}

/*
ACIs initializes, optionally sets and returns a new instance of Instructions
configured to store valid Instruction instances.

Slice values are delimited using the newline rune (ASCII #10).
*/
func ACIs(x ...any) (i Instructions) {
	_i := stackList().
		NoNesting(true).
		SetID(`instructions`).
		SetDelimiter(rune(10)).
		NoPadding(!StackPadding).
		SetCategory(`instructions`).
		SetPushPolicy(instructionsPushPolicy)

	_i.Push(x...)
	i = Instructions(_i)

	return
}

/*
instruction is the embedded (pointer!) type found within initialized
instances of the Instruction type. The fields are as follows:

• N contains the string name (or "ACL") of a particular Instruction; note
that this field cannot be reset for security reasons

• T contains one (1) TargetRules instance, which is a stackage.Stack type
alias containing a sequence of zero (0) or more TargetRule instances

• PB contains one (1) PermissionBindRules instance, which is a stackage.Stack alias
type containing a sequence of one (1) or more PermissionBindRule instances
*/
type instruction struct {
	ACL  string
	TRs  TargetRules
	PBRs PermissionBindRules
}

/*
canned invalidity tag constants for when ACI-related things go awry ...
*/
const (
	// badACI is supplied during Instruction string representation
	badACI = `<invalid_aci>`
)

func instructionsPushPolicy(x any) (err error) {
	switch tv := x.(type) {
	case Instruction:
		if tv.IsZero() {
			err = nilInstanceErr(tv)
		}
	default:
		err = pushErrorBadType(Instructions{}, tv, nil)
	}

	return
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r Instructions) Len() int {
	_r, _ := castAsStack(r)
	return _r.Len()
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r Instructions) IsZero() bool {
	_r, _ := castAsStack(r)
	return _r.IsZero()
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Stack.String method.
*/
func (r Instructions) String() string {
	_r, _ := castAsStack(r)
	return _r.String()
}

/*
String is a stringer method that returns the string representation of
the receiver instance.
*/
func (r Instruction) String() string {
	if err := r.Valid(); err != nil {
		return badACI
	}

	return sprintf("%s(%s; acl \"%s\"; %s)",
		r.instruction.TRs,
		version(), // sprints Version const.
		r.instruction.ACL,
		r.instruction.PBRs)
}

/*
Push wraps go-stackage's Stack.Push method. Only Instruction
instances are permitted for push.

In the case of a string value, it is automatically cast as an
instance of BindDistinguishedName using the appropriate keyword,
so long as the raw string is of a non-zero length.
*/
func (r Instructions) Push(x ...any) Instructions {
	_r, _ := castAsStack(r)

	// iterate variadic input arguments
	for i := 0; i < len(x); i++ {
		// Push it!
		_r.Push(x[i])
	}

	return Instructions(_r)
}

/*
Pop wraps go-stackage's Stack.Pop method.
*/
func (r Instructions) Pop() (x Instruction) {
	_r, _ := castAsStack(r)
	y, _ := _r.Pop()

	if assert, asserted := y.(Instruction); asserted {
		x = assert
	}

	return
}

/*
F returns the appropriate instance creator function for crafting individual Instruction
instances for submission to the receiver. This is merely a convenient alternative to
maintaining knowledge as to which function applies to the current receiver instance.

As there is only one possibility for instances of this design, the ACI function is returned.
*/
func (r Instructions) F() func(...any) Instruction {
	return ACI
}

/*
Valid wraps go-stackage's Stack.Valid method.
*/
func (r Instructions) Valid() (err error) {
	_b, _ := castAsStack(r)
	err = _b.Valid()
	return
}

/*
Index wraps go-stackage's Stack.Index method. Note that the
Boolean OK value returned by go-stackage by default will be
shadowed and not obtainable by the caller.
*/
func (r Instructions) Index(idx int) (x Instruction) {
	_r, _ := castAsStack(r)
	y, _ := _r.Index(idx)

	if assert, ok := y.(Instruction); ok {
		x = assert
	}

	return
}

/*
T returns the Target Rule found within the underlying receiver instance.
Note that a bogus Rule is returned if the receiver is nil, or unset.
*/
func (r Instruction) TRs() (trs TargetRules) {
	if !r.IsZero() {
		trs = r.instruction.TRs
	}

	return
}

/*
PBRs returns the PermissionBindRules found within the underlying receiver instance.
Note that a bogus Rule is returned if the receiver is nil, or unset.
*/
func (r Instruction) PBRs() (pbrs PermissionBindRules) {
	if !r.IsZero() {
		pbrs = r.instruction.PBRs
	}
	return
}

/*
ACL returns the access control label of the receiver, else a zero string if unset.
*/
func (r Instruction) ACL() (acl string) {
	if !r.IsZero() {
		acl = r.instruction.ACL
	}
	return
}

/*
Valid returns an instance of error that reflects any perceived errors
or deficiencies within the receiver instance.
*/
func (r Instruction) Valid() (err error) {
	return r.instruction.valid()
}

/*
valid is a private method called by instruction.Valid.
*/
func (r instruction) valid() (err error) {
	if r.isZero() {
		err = nilInstanceErr(r)
		return
	}

	if len(r.ACL) == 0 {
		err = instructionNoLabelErr()
		return
	}

	for _, valid := range []func() error{
		r.TRs.Valid,
		r.PBRs.Valid,
	} {
		if err = valid(); err != nil {
			break
		}
	}

	return
}

/*
IsZero returns a boolean value indicative of whether the receiver
is nil, or unset.
*/
func (r Instruction) IsZero() bool {
	return r.instruction.isZero()
}

func (r *instruction) isZero() bool {
	return r == nil
}

/*
ACI initializes, (optionally) sets and returns a new instance of the
Instruction type.

Input values must conform to the following specifications per the
intended field within the return instance:

• A non-zero string value shall be used for the effective Name, or "ACL"

• One (1) or more non-zero and unique PermissionBindRule instances

• One (1) or more non-zero and unique Condition instances bearing the
`target` categorical label

Please note the following constraints for the name of the receiver:

• Value cannot be reset (i.e.: renamed)

• Value should not contain the "version <float>" statement, as that is
imposed automatically during string representation procedures
*/
func ACI(x ...any) Instruction {
	return Instruction{newACI(x...)}
}

/*
newACI is a private function invoked by the package level
ACI function for the purpose of allocating memory for a new
*instruction instance, to be embedded within
an instance of Instruction.

If any arguments are provided, they shall (possibly) be set
within the return instance.
*/
func newACI(x ...any) (a *instruction) {
	a = new(instruction)
	a.TRs = TRs()
	a.PBRs = PBRs()
	if len(x) > 0 {
		a.set(x...)
	}

	return
}

/*
Set assigns one (1) or more values to the receiver. The input value(s)
must conform to the following conditions:

• If the value is a string, it shall become the immutable name (or "ACL")
of a given Instruction instance; this value cannot be changed once set

• If the value is a TargetRule instance, it shall be appended to the
receiver's TargetRules instance

• If the value is a TargetRules instance, it shall have all stack slice
members appended to the receiver's TargetRules instance

• If the value is a PermissionBindRule, and if it is valid (i.e.: contains
exactly one (1) valid Permission statement and exactly one (1) BindRules),
it shall be appended to the receiver's PermissionBindRules stack
*/
func (r *Instruction) Set(x ...any) *Instruction {
	if r.instruction == nil {
		r.instruction = newACI()
	}
	r.instruction.set(x...)
	return r
}

/*
set is a private method invoked by newACI and Instruction.Set
to handle the addition of new ACI components through type assertion and
validity checks where applicable.
*/
func (r *instruction) set(x ...any) {
	for i := 0; i < len(x); i++ {
		r.assertInstruction(x[i])
	}
}

func (r *instruction) assertInstruction(x any) {
	switch tv := x.(type) {
	case string:
		r.setLabel(tv)
	case TargetRules:
		r.instructionTargetPush(tv)
	case TargetRule:
		// TODO :: uniqueness check
		if K := matchTKW(tv.Category()); K != TargetKeyword(0x0) {
			r.TRs.Push(tv)
		}
	case PermissionBindRule:
		r.PBRs.Push(tv)
	}
}

func (r *instruction) setLabel(x string) {
	// Only set if non-zero and if field IS zero
	// (i.e.: don't allow renaming).
	if len(x) > 0 && len(r.ACL) == 0 {
		r.ACL = x
	}
}

func (r *instruction) instructionTargetPush(x TargetRules) {
	for i := 0; i < x.Len(); i++ {
		tgt := x.Index(i)
		// TODO :: uniqueness check
		if K := matchTKW(tgt.Keyword().String()); K != TargetKeyword(0x0) {
			r.TRs.Push(tgt)
		}
	}
}

/*
version returns the string version label for the ACI syntax.
*/
func version() string {
	return sprintf("version %.1f", Version)
}

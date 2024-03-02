package aci

/*
aci.go contains the top-level access control instructor methods and types.
*/

/*
Version defines the official ACI syntax version number implemented and honored by this package.
*/
const Version float32 = 3.0

/*
Instruction conforms to the ACI syntax specification associated with the [Version] constant value of this package.

Instances of this type, when represented in their string form, are intended for submission to an X.500/LDAP DSA for assignment (via the 'aci' LDAP Attribute Type) to the relevant directory entry.
*/
type Instruction struct {
	*instruction
}

/*
ACIs initializes, optionally sets and returns a new instance of [Instructions] configured to store valid [Instruction] instances.

Slice values are delimited using the newline rune (ASCII #10).
*/
func ACIs(x ...any) (i Instructions) {
	_i := stackList().
		NoNesting(true).
		SetID(`instructions`).
		SetDelimiter(rune(10)).
		NoPadding(true).
		SetCategory(`instructions`)

		// cast _i as a proper Instructions instance
		// (i). We do it this way to gain access to
	// the method for the *specific instance*
	// being created (o), thus allowing things
	// like uniqueness checks, etc., to occur
	// during push attempts, providing more
	// helpful and non-generalized feedback.
	i = Instructions(_i)
	_i.SetPushPolicy(i.pushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	i.Push(x...)

	return
}

/*
instruction is the embedded (pointer!) type found within initialized
instances of the Instruction type. The fields are as follows:

• N contains the string name (or "ACL") of a particular Instruction; note
that this field cannot be reset for security reasons

• T contains one (1) TargetRules instance, which is a [stackage.Stack] type
alias containing a sequence of zero (0) or more [TargetRule] instances

• PB contains one (1) PermissionBindRules instance, which is a [stackage.Stack] alias
type containing a sequence of one (1) or more [PermissionBindRule] instances
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

func (r Instructions) pushPolicy(x ...any) (err error) {
	if r.contains(x[0]) {
		err = pushErrorNotUnique(r, x[0], nil)
		return
	}

	err = pushErrorBadType(Instructions{}, x[0], nil)
	switch tv := x[0].(type) {
	case Instruction:
		err = tv.Valid()
	}

	return
}

/*
Len wraps the [stackage.Stack.Len] method.
*/
func (r Instructions) Len() int {
	return r.cast().Len()
}

/*
Contains returns a Boolean value indicative of whether value x, if a string or [Instruction] instance, already resides within the receiver instance.

Case is not significant in the matching process.
*/
func (r Instructions) Contains(x any) bool {
	return r.contains(x)
}

/*
contains is a private method called by Instructions.Contains.
*/
func (r Instructions) contains(x any) bool {
	if r.Len() == 0 {
		return false
	}

	var candidate string

	switch tv := x.(type) {
	case string:
		candidate = tv
	case Instruction:
		candidate = tv.String()
	}

	candidate = condenseWHSP(candidate)
	for i := 0; i < r.Len(); i++ {
		// case is not significant here.
		if eq(r.Index(i).String(), candidate) {
			return true
		}
	}

	return false
}

/*
IsZero wraps the [stackage.Stack.IsZero] method.
*/
func (r Instructions) IsZero() bool {
	return r.cast().IsZero()
}

/*
String is a stringer method that returns the string representation of the receiver instance.

This method wraps the [stackage.Stack.String] method.
*/
func (r Instructions) String() string {
	return r.cast().String()
}

/*
String is a stringer method that returns the string representation of the receiver instance.
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
Push wraps the [stackage.Stack.Push] method. Only [Instruction] instances are permitted for push.

In the case of a string value, it is automatically cast as an instance of [BindDistinguishedName] using the appropriate [BindKeyword], so long as the raw string is of a non-zero length.
*/
func (r Instructions) Push(x ...any) Instructions {
	_r := r.cast()

	// iterate variadic input arguments
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case string:
			var ins Instruction
			if err := ins.Parse(tv); err == nil {
				_r.Push(ins)
			}
		default:
			_r.Push(tv)
		}
	}

	return r
}

/*
Pop wraps the [stackage.Stack.Pop] method.
*/
func (r Instructions) Pop() (x Instruction) {
	y, _ := r.cast().Pop()
	if assert, asserted := y.(Instruction); asserted {
		x = assert
	}

	return
}

/*
F returns the appropriate instance creator function for crafting individual [Instruction] instances for submission to the receiver. This is merely a convenient alternative to maintaining knowledge as to which function applies to the current receiver instance.

As there is only one possibility for instances of this design, the package-level [ACI] function is returned.
*/
func (r Instructions) F() func(...any) Instruction {
	return ACI
}

/*
Valid wraps the [stackage.Stack.Valid] method.
*/
func (r Instructions) Valid() (err error) {
	err = r.cast().Valid()
	return
}

/*
Index wraps the [stackage.Stack.Index] method. Note that the Boolean OK value returned by [stackage] by default will be shadowed and not obtainable by the caller.
*/
func (r Instructions) Index(idx int) (x Instruction) {
	y, _ := r.cast().Index(idx)
	if assert, ok := y.(Instruction); ok {
		x = assert
	}

	return
}

/*
T returns the [TargetRules] instance found within the underlying receiver instance. Note that a bogus [TargetRules] instance is returned if the receiver is nil, or unset.
*/
func (r Instruction) TRs() (trs TargetRules) {
	if !r.IsZero() {
		trs = r.instruction.TRs
	}

	return
}

/*
PBRs returns the [PermissionBindRules] instance found within the underlying receiver instance. Note that a bogus [PermissionBindRules] instance is returned if the receiver is nil, or unset.
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
Valid returns an instance of error that reflects any perceived errors or deficiencies within the receiver instance.
*/
func (r Instruction) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
	}
	return
}

/*
IsZero returns a Boolean value indicative of whether the receiver is nil, or unset.
*/
func (r Instruction) IsZero() bool {
	return r.instruction.isZero()
}

func (r *instruction) isZero() bool {
	return r == nil
}

/*
ACI initializes, (optionally) sets and returns a new instance of the [Instruction] type.

Input values must conform to the following specifications per the intended field within the return instance:

  - A non-zero string value shall be used for the effective Name, or "ACL"
  - One (1) [PermissionBindRules] instance
  - One (1) [TargetRules] instance

Please note the following constraints for the name of the receiver:

  - Value cannot be reset (i.e.: renamed)
  - Value should not contain the "version <float>" statement, as that is imposed automatically during string representation procedures
*/
func ACI(x ...any) Instruction {
	return Instruction{newACI(x...)}
}

/*
newACI is a private function invoked by the package level ACI function for the purpose of allocating memory for a new *instruction instance, to be embedded within an instance of Instruction.

If any arguments are provided, they shall (possibly) be set within the return instance.
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
Set assigns one (1) or more values to the receiver. The input value(s) must conform to the following conditions:

  - If the value is a string, it shall become the immutable name (or "ACL") of a given [Instruction] instance; this value cannot be changed once set
  - If the value is a [TargetRule] instance, it shall be appended to the receiver's [TargetRules] instance
  - If the value is a [TargetRules] instance, it shall have all stack slice members appended to the receiver's [TargetRules] instance
  - If the value is a [PermissionBindRule], and if it is valid (i.e.: contains exactly one (1) valid [Permission] statement and exactly one (1) [BindRules] instance), it shall be appended to the receiver's [PermissionBindRules] stack
*/
func (r *Instruction) Set(x ...any) *Instruction {
	if r.instruction == nil {
		r.instruction = newACI()
	}
	r.instruction.set(x...)
	return r
}

/*
set is a private method invoked by newACI and Instruction.Set to handle the addition of new ACI components through type assertion and validity checks where applicable.
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
		r.targetPush(tv)
	case TargetRule:
		r.TRs.Push(tv)
	case PermissionBindRule:
		r.PBRs.Push(tv)
	case PermissionBindRules:
		r.permissionBindRulesPush(tv)
	}
}

func (r *instruction) setLabel(x string) {
	// Only set if non-zero and if field IS zero
	// (i.e.: don't allow renaming).
	if len(x) > 0 && len(r.ACL) == 0 {
		r.ACL = x
	}
}

func (r *instruction) targetPush(x TargetRules) {
	for i := 0; i < x.Len(); i++ {
		tgt := x.Index(i)
		if K := matchTKW(tgt.Keyword().String()); K != TargetKeyword(0x0) {
			r.TRs.Push(tgt)
		}
	}
}

func (r *instruction) permissionBindRulesPush(x PermissionBindRules) {
	for i := 0; i < x.Len(); i++ {
		r.PBRs.Push(x.Index(i))
	}
}

/*
version returns the string version label for the ACI syntax.
*/
func version() string {
	return sprintf("version %.1f", Version)
}

package aci

/*
aci.go contains the top-level access control instructor methods and types.
*/

import (
	"github.com/JesseCoretta/go-aci/internal/aciparser"
)

/*
Instruction conforms to the ACI syntax specification associated with
the Version constant value of this package.
*/
type Instruction struct {
	*instruction
}

/*
instruction is the embedded (pointer!) type found within instances
of Instruction. The fields are as follows:

• N contains the string name (or "ACL") of a particular Instruction; note
that this field cannot be reset for security reasons

• T contains one (1) or more Condition instances that bear the `target` categorical
string label

• PB contains one (1) or more PermissionBindRule instances, each containing exactly
one (1) Permission (grant or deny) and one (1) Rule instance bearing the `bind`
categorical string label, and containing one (1) or more Condition instances that
each bear the said label.
*/
type instruction struct {
	N  string `aci:"name"`
	T  Rule   `aci:"target"`
	PB Rule   `aci:"permission_bind_rules"`
}

/*
canned invalidity tag constants for when ACI-related things go awry ...
*/
const (
	// badACI is supplied during Instruction string representation
	badACI = `<invalid_aci>`

	// badPB is supplied during PermissionBindRule string representation
	// attempts that fail
	badPB = `<invalid_pbrule>`
)

/*
ACIs initializes, optionally sets and returns a new instance of Rule configured
to store valid Instruction instances.
*/
func ACIs() Rule {
	return Rule(stackageList().JoinDelim("\n")).
		setPushPolicy().
		NoPadding(!RulePadding).
		setID(`instructions`)
}

/*
String is a stringer method that returns the string representation of
the receiver instance.
*/
func (r Instruction) String() string {
	if err := r.Valid(); err != nil {
		return badACI
	}

	return sprintf("%s(%s; acl \"%s\"; %s;)",
		r.instruction.T,
		version(), // sprints Version const.
		r.instruction.N,
		r.instruction.PB)
}

/*
T returns the Target Rule found within the underlying receiver instance.
Note that a bogus Rule is returned if the receiver is nil, or unset.
*/
func (r Instruction) T() Rule {
	var t Rule
	if r.IsZero() {
		return t
	}
	return r.instruction.T
}

/*
PB returns the Permission/Bind Rule found within the underlying receiver instance.
Note that a bogus Rule is returned if the receiver is nil, or unset.
*/
func (r Instruction) PB() Rule {
	var p Rule
	if r.IsZero() {
		return p
	}
	return r.instruction.PB
}

/*
N returns the name (or "ACL") of the receiver, else a zero string if unset.
*/
func (r Instruction) N() string {
	if r.IsZero() {
		return ``
	}
	return r.instruction.N
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
		err = errorf("%T instance is nil", r)
		return
	}

	if len(r.N) == 0 {
		err = errorf("%T has no name (ACL); set a string name value using %T.Set", r, r)
		return
	}

	for _, valid := range []func() error{
		r.T.Valid,
		r.PB.Valid,
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
PB contains one (1) Permission and one (1) Bind Rule. Instances of this type
are used to create top-level ACI Permission+Bind Rule pairs, of which at least
one (1) is required in any given access control instructor definition.

Users may compose instances of this type manually, or using the PB package
level function, which automatically invokes value checks.
*/
type PermissionBindRule struct {
	P Permission `aci:"permission_rule"`
	B Rule       `aci:"bind_rule"`
}

/*
PB returns an instance of PermissionBindRule, bearing the Permission P and
the Bind Rule B. The values P and B shall undergo validity checks per the
conditions of the PermissionBindRule.Valid method automatically.

Generally, an ACI only has a single PermissionBindRule, though multiple
instances of this type are allowed per the syntax specification honored
by this package.
*/
func PB(P Permission, B Rule) PermissionBindRule {
	r := PermissionBindRule{P, B}
	if err := r.Valid(); err != nil {
		return PermissionBindRule{}
	}
	return r
}

/*
Valid returns an error instance should any of the following conditions
evaluate as true:

• Permission.Valid returns an error for P

• Rule.Valid returns an error for B

• Rule.Len returns zero (0) for B

• Rule.Category returns a categorical label other than `bind` for B
*/
func (r PermissionBindRule) Valid() (err error) {
	return r.valid()
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
	} else if r.B.ID() != `pb` && r.B.ID() != `bind` {
		err = errorf("%T is not a permission+bind rule (%s)", r.B, r.B.ID())
	}

	return
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
func (r PermissionBindRule) string() string {
	if err := r.valid(); err != nil {
		return badPB
	}
	return sprintf("%s %s", r.P, r.B)
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
	a.T = T()
	a.PB = pbrule()
	if len(x) > 0 {
		a.set(x...)
	}

	return
}

/*
Set assigns one (1) or more values to the receiver. The input value(s)
must conform to the following conditions:

• If the value is a string, it shall become the immutable name (or "ACL")
of a given Instruction instance; this value cannot be changed
once set.

• If the value is a Condition, and if it bears the `target` categorical
label, it shall be appended to the stack of Target Rule instances

• If the value is a Rule, and if it bears the `target` categorical label,
it shall have all stack slice members appended to the receiver's Target
Rule stack.

• If the value is a PermissionBindRule, and if it is valid (i.e.: contains
exactly one (1) valid Permission statement and exactly one (1) Bind Rule
instance containing one (1) or more Bind Rule conditions), it shall be
appended to the receiver's PermissionBind Rule stack.
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
		// Only set if non-zero and if field IS zero
		// (i.e.: don't allow renaming).
		if len(tv) > 0 && len(r.N) == 0 {
			r.N = tv
		}
	case Condition:
		if K := matchTKW(tv.Category()); K != TargetKeyword(0x0) {
			r.T.Push(tv)
		}
	case Rule:
		r.assertPushRule(tv)
	case PermissionBindRule:
		r.PB.Push(tv)
	}
}

/*
assertPushRule is a private method invoked by instruction.set
when type assertion reveals a Rule instance containing Target Rule slices
meant for append within the receiver.
*/
func (r *instruction) assertPushRule(x Rule) {
	if x.Len() > 0 && matchBKW(x.Category()) != BindKeyword(0x0) || matchTKW(x.Category()) != TargetKeyword(0x0) {
		for t := 0; t < x.Len(); t++ {
			tgt, ok := x.Index(t)
			if !ok {
				continue
			}
			r.T.Push(tgt)
		}
	} else if x.ID() == `pb` && x.Len() > 0 {
		r.PB.Push(x)
	}
}

/*
parsePBR reads and processes a sequence of tokens into one (1) Permission and one (1)
bind rule. An error and a chop index is returned alongside these components.
*/
func parsePBR(tokens []string) (chop int, pbr []PermissionBindRule, err error) {
        var mode string = `permission` // starting mode is always permission

        pbr = make([]PermissionBindRule, 0)

        for _, token := range tokens {
                if len(tokens) <= 1 {
                        return
                }

                switch token {

                case `allow`, `deny`:
                        switch mode {

                        case `permission`:
                                var skipTo int
                                var br Rule
                                var perm Permission
                                if skipTo, perm, err = parsePerm(tokens); err != nil {
                                        return
                                }

                                tokens = tokens[skipTo:]
                                chop = skipTo
                                if br, skipTo, err = parseBindRule(tokens, -1, 0); err != nil {
                                        return
                                }
                                pbr = append(pbr, PB(perm, br))

                                // Done processing!
                                if ( skipTo == -1 || ( skipTo-1 == len(tokens) ) ) {
                                        chop = -1
                                        return
                                }

                                tokens = tokens[skipTo:]
                                chop += skipTo
                                mode = `bind`
                        }

                case `;`:
                        chop++
                        if mode == `bind` {
                                mode = `permission`
                                continue
                        }

                case `,`:
                        chop++
                        if mode == `permission` {
                                // don't preserve permission
                                // delimiters.
                                continue
                        }
                }
        }

        if len(pbr) == 0 {
                err = errorf("No Permission Bind Rule(s) found; aborting")
        }

        return
}

/*
parseInstruction returns a populated instance of Instruction, alongside
an error instance. This is the top-level parser function, which handles
all lower-levels of value recognition.

The input argument, expr, is the string-based ACIv3 expression in its
complete form.
*/
func parseInstruction(expr string) (a Instruction, err error) {
        // if expr is zero-length, absolutely nothing
        // can be done and is considered a user error.
        if len(expr) == 0 {
                err = errorf("Cannot process zero-length instruction")
                return
        }

        // Initialize our new Instruction instance. This is the
        // return value defined in the signature, and is that
        // which the user expects in exchange for their text
        // expression.
        a = ACI()

        // Remove unneeded contiguous WHSP (tab/space) as
        // well as newlines. Also remove any leading or
        // trailing WHSP, contiguous or not.
        expr = condenseWHSP(expr)

        // Tokenize our aci string input
        tokens := aciparser.InstructionToTokens(expr)

        // Always start with target mode, as
        // targetRules will be the first items
        // encountered (if defined).
        var mode string = `target`

        // Keep track of where we'll be next.
        var next string

        // Keep track of the so-called "chop index",
        // which is used following recursion-based
        // processing phases to avoid superfluous
        // handling of already-seen tokens.
        var skipTo int

        // Iterate our tokenized ANTLR char stream
        // and handle each token accordingly.
        for index, token := range tokens {

                // If recursion was performed, we MAY need to skip
                // ahead to avoid processing tokens already handled.
                // This only happens IF the chop index is non-zero
                // AND is higher than the current index AND does not
                // exceed the current length of the token slices.
                if skipTo != 0 && skipTo > index && skipTo < len(tokens) {
                        continue
                }

                // If we have a ways to go, store the next (upcoming)
                // token so we can "look ahead" if needed.
                if len(tokens) <= 2 {
                        // Remaining tokens are too few, so we bail.
                        break
                }

                // Perform a value switch to analyze the current token
                // and see if it conforms to the components we expect
                // to find.
                switch {

                // If we found the anchor, that means the expression is
                // "targetless", meaning no target rules were specified.
                case token == `version 3.0; acl`:
                        mode = `acl`

                // If we found a semicolon, we know the current mode is
                // ending (OR the current value of the current mode has
                // been handled).
                case token == `;`:

                        // Perform a mode switch so we can take appropriate
                        // action based on the current stage of processing.
                        switch mode {

                        // A semicolon while acl mode is in effect
                        // means the acl is about to end, and the
                        // PermissionBindRule phase is next. Thus
                        // the tokens slices is FIFO trimmed.
                        case `acl`:

                                // Recurse into PermissionBindRule processing
                                // phase. Note that an ACI will always have one
                                // (1) OR MORE of these.
                                var pbr []PermissionBindRule
                                if skipTo, pbr, err = parsePBR(tokens[1:]); err != nil {
                                        return
                                }

                                // Add the resultant permission + bind rule(s)
                                // instance to our return Instruction.
                                for p := 0; p < len(pbr); p++ {
                                        a.Set(pbr[p])
                                }

                                // If skipTo is minus one, this indicates we
                                // finished processing the PBR section of the
                                // ACI.
                                if skipTo == -1 {
                                        return
                                }

                                // Truncate the token slices to begin where we
                                // just left off, thus avoiding superfluous
                                // processing of already-seen tokens.
                                tokens = tokens[skipTo:]
                        }

                // We we found an opening parenthesis, we know that we're either
                // in the target rule processing phase, OR just finishing said
                // phase and are about to move onto the ACI "anchor".
                case token == `(`:
                        if next == `version 3.0; acl` {
                                mode = `acl`
                                continue
                        }

                        // perform a mode switch so we can take appropriate
                        // action based on the current stage of processing.
                        switch mode {

                        // an opening parenthesis while target mode is in
                        // effect means that we're about to receive one (1)
                        // or more target rule conditions.
                        case `target`:

                                // prepare our targetRules Rule stack, into
                                // which one (1) or more target rule condition
                                // instances shall be pushed.
                                var targetRules Rule

                                // recurse into parseTR, extract the target
                                // rule expressions and obtain our chop index.
                                if skipTo, targetRules, err = parseTR(tokens[1:]); err != nil {
                                        return
                                }

                                // We are done processing target rules. We
                                // know for certain the next mode is 'acl',
                                // so set it now.
                                mode = `acl`

                                // Truncate the token slices to begin where
                                // we just left off.
                                tokens = tokens[skipTo:]

                                // Set our new targetRules Rule instance (IF
                                // non-zero) to the ACI instance (a).
                                if targetRules.Len() > 0 {
                                        a.Set(targetRules)
                                }
                        }

                // default is a catch-all for any token not explicitly handled
                // in the above case statements.
                default:

                        // If the mode is acl, we're expecting a double-quoted
                        // access control label (hence the acronym). Strip the
                        // quotes off, as go-stackage handles encapsulation
                        // without the need for preserving characters literally,
                        // and add the naked string value.
                        if mode == `acl` {
                                a.Set(unquote(token))
                        }
                }
        }

        // We're done, return the Instruction along with
        // (what is likely) a nil error.
        return
}

/*
pbrule is a private function that returns a new Permission/Bind Rule stack,
suitable for storing individual PermissionBindRule instances.

This function invokes delimitation using `; ` (that is, a semi-colon char
followed by a single WHSP char) for stringification involving Rule values
with a length greater than one (1).
*/
func pbrule() Rule {
	return Rule(stackageList().JoinDelim(`; `)).
		setID(`pb`).
		setCategory(`bind`).
		NoPadding(!RulePadding).
		setPushPolicy()
}

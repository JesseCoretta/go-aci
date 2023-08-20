package aci

/*
rules.go contains the go-stackage Stack alias "Rule" and its extended
methods.
*/

import "github.com/JesseCoretta/go-stackage"

/*
T creates and returns a new instance of Rule with an initialized
embedded stack configured to function as a Target Rule store that
is meant to contain one (1) or more Condition instances, each of
which bear one (1) of the following Target Rule keyword constants:

• Target

• TargetTo

• TargetFrom

• TargetAttr

• TargetCtrl

• TargetScope

• TargetFilter

• TargetAttrFilters

• TargetExtOp

Please note that Rule instances of this design are set with a maximum
capacity of nine (9) for both the following reasons:

• There are only said number of Target Rule keywords supported within the
ACI syntax specification honored by this package, and ...

• Individual Target Rule keywords can only be used once per ACI; in other
words, one cannot specify have multiple `target` conditions within the same
ACI.

Instances of this design generally are assigned to top-level instances of
Instruction.
*/
func T() Rule {
	return Rule(stackageList(9)).
		setID(`target`).
		setCategory(`target`).
		NoPadding(!RulePadding).
		setPushPolicy()
}

/*
And returns an instance of Rule configured to express Boolean AND logical
operations. Rule instances of this design generally contain Bind Rule
Condition instances, or other nested Boolean-related Rule instances.

Rule instances of this design may also be used in TargetAttrFilters expressions
that involve multiple attr:filter combinations ANDed together by symbols (&&).

The embedded type within the return is stackage.Stack via the go-stackage
package's And function.
*/
func And() Rule {
	return Rule(stackageAnd()).
		setID(`bind`).
		setCategory(`and`).
		NoPadding(!RulePadding).
		setPushPolicy()
}

/*
Or returns an instance of Rule configured to express Boolean OR logical
operations. Rule instances of this design generally contain Bind Rule
Condition instances, or other nested Boolean-related Rule instances.

Rule instances of this design may also be used in TargetAttr expressions
that involve a sequence of attributeType names ORed together by symbols (||).

The embedded type within the return is stackage.Stack via the go-stackage
package's Or function.
*/
func Or() Rule {
	return Rule(stackageOr()).
		setID(`bind`).
		setCategory(`or`).
		NoPadding(!RulePadding).
		setPushPolicy()
}

/*
Not returns an instance of Rule configured to express Boolean NOT logical
operations. Rule instances of this design generally contain Bind Rule
Condition instances, or other nested Boolean-related Rule instances.

The embedded type within the return is stackage.Stack via the go-stackage
package's Not function.
*/
func Not() Rule {
	return Rule(stackageNot()).
		setID(`bind`).
		setCategory(`not`).
		NoPadding(!RulePadding).
		setPushPolicy()
}

/*
Rule contains an ordered sequence of conditions, which can also nest other
Condition instances. Rules can be extended horizontally and vertically without
any fixed limit.

An instance of Rule can describe a Target or Bind rule, an LDAP filter or an
arbitrary list of attributeTypes or object identifiers.

Rule is a type alias of go-stackage's Stack type and is used heavily throughout
this package.
*/
type Rule stackage.Stack

/*
String wraps go-stackage's Stack.String method.
*/
func (r Rule) String() string {
	return stackage.Stack(r).String()
}

/*
setCategory wraps go-stackage's Stack.SetCategory method.
*/
func (r Rule) setCategory(cat string) Rule {
	stackage.Stack(r).SetCategory(cat)
	return r
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r Rule) IsZero() bool {
	return stackage.Stack(r).IsZero()
}

/*
reset wraps go-stackage's Stack.Reset method. This is a private
method in go-aci.
*/
func (r Rule) reset() {
	stackage.Stack(r).Reset()
}

/*
transfer wraps go-stackage's Stack.Transfer method. This is a
private method in go-aci.
*/
func (r Rule) transfer(x Rule) bool {
	return stackage.Stack(r).Transfer(stackage.Stack(x))
}

/*
ID wraps go-stackage's Stack.ID method.
*/
func (r Rule) ID() string {
	if r.IsZero() {
		return ``
	}
	return stackage.Stack(r).ID()
}

/*
Category wraps go-stackage's Stack.Category method.
*/
func (r Rule) Category() string {
	if r.IsZero() {
		return ``
	}
	return stackage.Stack(r).Category()
}

/*
setID wraps go-stackage's Stack.SetID method.
*/
func (r Rule) setID(id string) Rule {
	stackage.Stack(r).SetID(id)
	return r
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r Rule) Len() int {
	if r.IsZero() {
		return 0
	}
	return stackage.Stack(r).Len()
}

/*
Push wraps go-stackage's Stack.Push method.
*/
func (r Rule) Push(x ...any) Rule {
	stackage.Stack(r).Push(x...)
	return r
}

/*
Pop wraps go-stackage's Stack.Pop method.
*/
func (r Rule) Pop() (any, bool) {
	return stackage.Stack(r).Pop()
}

/*
Remove wraps go-stackage's Stack.Remove method.
*/
func (r Rule) Remove(i int) (any, bool) {
	return stackage.Stack(r).Remove(i)
}

/*
Index wraps go-stackage's Stack.Index method.
*/
func (r Rule) Index(i int) (slice any) {
	slice, _ = stackage.Stack(r).Index(i)
	return
}

/*
Insert wraps go-stackage's Stack.Insert method.
*/
func (r Rule) Insert(x any, left int) bool {
	return stackage.Stack(r).Insert(x, left)
}

/*
Paren wraps go-stackage's Stack.Paren method.
*/
func (r Rule) Paren(state ...bool) Rule {
	stackage.Stack(r).Paren(state...)
	return r
}

/*
isParen wraps go-stackage's Stack.IsParen method.
*/
func (r Rule) isParen() bool {
	return stackage.Stack(r).IsParen()
}

/*
insert wraps go-stackage's Stack.Insert method.
*/
func (r Rule) insert(x any, left int) (ok bool) {
	ok = stackage.Stack(r).Insert(x,left)
	return
}

/*
Fold wraps go-stackage's Stack.Fold method.
*/
func (r Rule) Fold(state ...bool) Rule {
	stackage.Stack(r).Fold(state...)
	return r
}

/*
leadOnce wraps go-stackage's Stack.LeadOnce method.
*/
func (r Rule) leadOnce(state ...bool) Rule {
	stackage.Stack(r).LeadOnce(state...)
	return r
}

/*
NoPadding wraps go-stackage's Stack.NoPadding method.
*/
func (r Rule) NoPadding(state ...bool) Rule {
	stackage.Stack(r).NoPadding(state...)
	return r
}

/*
Encap wraps go-stackage's Stack.Encap method.
*/
func (r Rule) Encap(x ...any) Rule {
	stackage.Stack(r).Encap(x...)
	return r
}

/*
JoinDelim wraps go-stackage's Stack.JoinDelim method.
*/
func (r Rule) JoinDelim(x string) Rule {
	stackage.Stack(r).JoinDelim(x)
	return r
}

/*
Symbol wraps go-stackage's Stack.Symbol method.
*/
func (r Rule) Symbol(x ...any) Rule {
	stackage.Stack(r).Symbol(x...)
	return r
}

/*
Traverse wraps go-stackage's Stack.Traverse method.
*/
func (r Rule) Traverse(indices ...int) (any, bool) {
	return stackage.Stack(r).Traverse(indices...)
}

/*
Kind returns the string value `bind`, `pb` or `target` depending
on the configuration of the receiver.
*/
func (r Rule) Kind() (kind string) {
	return stackage.Stack(r).ID()
}

/*
Valid wraps go-stackage's Stack.Valid method.
*/
func (r Rule) Valid() (err error) {
	if err = stackage.Stack(r).Valid(); err != nil {
		return
	}

	// TODO - per Rule-type checks
	return
}

/*
Eq is a convenience method that crafts a particular equality-based
Condition instance based upon the categorical string label value
assigned to the receiver. The return instance may be used within
Target and Bind Rule expressions, and will automatically bear the
appropriate keyword and label.

For example, if a receiver instance were created with the Ctrls()
package level function, its categorical label `controls` would
result in the creation of a `targetcontrol` Target Rule.
*/
func (r Rule) Eq() (c Condition) {
	if r.IsZero() {
		return Condition{}
	}

	if k := matchTKW(r.Category()); k != TargetKeyword(0x0) {
		// Rule is Target Rule related
		c = Cond(k, r.Paren(false), Eq).
			NoPadding(!ConditionPadding).
			Encap(`"`).Paren().
			setCategory(k.String())

	} else if j := matchBKW(r.Category()); j != BindKeyword(0x0) {
		// Rule is Bind Rule related
		c = Cond(j, r.Paren(false), Eq).
			NoPadding(!ConditionPadding).
			Encap(`"`).Paren().
			setCategory(j.String())
	}

	return
}

/*
Ne is a convenience method that crafts a particular negated-equality
Condition instance based upon the categorical string label value
assigned to the receiver. The return instance may be used within
Target and Bind Rule expressions, and will automatically bear the
appropriate keyword and label.

For example, if a receiver instance were created with the Ctrls()
package level function, its categorical label `controls` would
result in the creation of a `targetcontrol` Target Rule.

Negated equality matching operators should be used with EXTREME CAUTION.
*/
func (r Rule) Ne() (c Condition) {
	if r.IsZero() {
		return Condition{}
	}

	if k := matchTKW(r.Category()); k != TargetKeyword(0x0) {
		// Rule is Target Rule related
		c = Cond(k, r.Paren(false), Ne).
			NoPadding(!ConditionPadding).
			Encap(`"`).Paren().
			setCategory(k.String())

	} else if j := matchBKW(r.Category()); j != BindKeyword(0x0) {
		// Rule is Bind Rule related
		c = Cond(j, r.Paren(false), Ne).
			NoPadding(!ConditionPadding).
			Encap(`"`).Paren().
			setCategory(j.String())
	}

	return
}

/*
canPush is a go-stackage PushPolicy-compliant method
assigned to various Rule instances for the purpose
of vetting all push attempts upon a given Rule instance.
It is executed for ALL push attempts upon Rule instances
within which this method has been implemented.

As an example, this method allows uniqueness enforcement
for Rule instances containing attributeType values.
*/
func (r Rule) canPush(x any) (err error) {
	var ok bool
	switch r.ID() {

	// list of Instruction instances
	case `instructions`:
		ok = aciCanPush(x)

	// bind rules in varied parenthetical
	// states
	case `parenthetical_bind`,
		`enveloped_parenthetical_bind`,
		`enveloped_bind`,
		`bind`, `and`, `or`, `not`:
		ok = bindRuleCanPush(x)

	// target rules
	case `target`:
		// only allow unique Condition instances which
		// bear the categorical string label `target`
		// to be considered for pushes. Uniqueness is
		// true if the given target keyword does not
		// already reside within a Condition present
		// within the receiver (e.g.: `targetfilter`).
		ok = targetRuleCanPush(r, x)

	// target attribute / URI search attribute lists
	case `attributes`, TargetAttr.String():
		ok = attrRuleCanPush(r, x)

	// permission bind rule
	case `pb`:
		ok = pbRuleCanPush(x)

	// arbitrary list
	case `list`:
		ok = listRuleCanPush(x)
	}

	// Return an error if the PushPolicy denies a request
	// for any reason.
	if !ok {
		err = pushPolicyViolationErr(r, x)
	}

	return
}

/*
setPushPolicy assigns the default Rule.canPush push policy
method to the receiver instance, which shall be executed
for all push attempts for verification reasons.
*/
func (r Rule) setPushPolicy() Rule {
	stackage.Stack(r).SetPushPolicy(r.canPush)
	return r
}

/*
bindRuleCanPush analyzes the input argument (x) to determine
eligibility for push into a Bind Rule.
*/
func bindRuleCanPush(x any) (ok bool) {
	switch tv := x.(type) {
	case string:
		if len(tv) > 0 {
			ok = true
		}
	case Condition:
		if matchBKW(tv.Keyword()) != BindKeyword(0x0) {
			ok = true
		}
	case PermissionBindRule:
		ok = true
	case Rule:
		ok = true
	case ObjectIdentifier, DistinguishedName, AttributeBindTypeOrValue:
		ok = true
	}

	return
}

/*
attrRuleCanPush is a handler function for Rule instances bearing
the `attributes` categorical string label value. This function
shall ensure uniqueness is maintained amongst slice members, and
that all pushed attributeType values are well-formed and valid.

A boolean return value indicative of "push authorization" will be
returned.
*/
func attrRuleCanPush(r Rule, x any) (ok bool) {
	// don't process any Rule instance that is
	// ill-suited for attributeType storage.
	var cat string = r.Category()
	if !eq(cat, `attributes`) && !eq(cat, TargetAttr.String()) {
		return
	}

	// Assert to string, or bail out
	at, assert := x.(string)
	if !assert {
		return
	}

	// Ensure AttributeType is well
	// formed and conformant, or is
	// an asterisk (ASCII #42).
	if !isIdentifier(at) && x != string(rune(42)) {
		return
	}

	// Iterate all indices present within
	// Rule r at this time. If the input
	// string matches one of these, push
	// shall be denied.
	for i := 0; i < r.Len(); i++ {

		// Call index number
		a := r.Index(i)

		// Assert slice index to AttributeType
		// or continue at next iteration
		atyp, assert := a.(AttributeType)
		if !assert {
			continue
		}

		// Does input value x match current
		// iteration attributeType value?
		if eq(at, atyp.String()) {
			return
		}
	}

	// seems fine ...
	ok = true
	return
}

func aciCanPush(x any) (ok bool) {
	switch tv := x.(type) {
	case string:
		ok = len(tv) > 0 // parser will do the real legwork
	case Instruction:
		ok = tv.Valid() == nil
	}

	return
}

func targetRuleCanPush(r Rule, x any) (ok bool) {
	switch tv := x.(type) {

	// scan Rule r for all Condition instances,
	// and compare their keyword with that of
	// input argument x, also an instance of
	// Condition.

	// Matched Target Rule Condition
	case Condition:
		if matchTKW(tv.Keyword()) == TargetKeyword(0x0) {
			return
		}

		// Iterate all slice members of r.
		for i := 0; i < r.Len(); i++ {
			// call index i
			sl := r.Index(i)

			// If i is a Condition, perform
			// keyword comparison.
			switch uv := sl.(type) {
			case Condition:
				if eq(uv.Keyword(), tv.Keyword()) {
					return
				}
			}
		}
		ok = true

	// matched Object Identifier
	case AttributeType:

		// Iterate all slice members of r.
		for i := 0; i < r.Len(); i++ {
			// call index i
			sl := r.Index(i)

			// If i is an OID or string, perform
			// string literal comparison
			switch uv := sl.(type) {
			case AttributeType:
				if tv.String() == uv.String() {
					return
				}
			case string:
				if tv.String() == uv {
					return
				}
			}
		}
		ok = true

	// matched Object Identifier
	case ObjectIdentifier:

		// Iterate all slice members of r.
		for i := 0; i < r.Len(); i++ {
			// call index i
			sl := r.Index(i)

			// If i is an OID or string, perform
			// string literal comparison
			switch uv := sl.(type) {
			case ObjectIdentifier:
				if tv.String() == uv.String() {
					return
				}
			case string:
				if tv.String() == uv {
					return
				}
			}
		}
		ok = true

	// matched Object Identifier
	case DistinguishedName:

		// Iterate all slice members of r.
		for i := 0; i < r.Len(); i++ {
			// call index i
			sl := r.Index(i)

			// If i is a DN or string, perform
			// string literal comparison
			switch uv := sl.(type) {
			case DistinguishedName:
				if tv.String() == uv.String() {
					return
				}
			case string:
				if tv.String() == uv {
					return
				}
			}
		}
		ok = true

	case Rule:
		if matchTKW(tv.Category()) == TargetKeyword(0x0) && tv.Len() > 0 {
			return
		}
		ok = true
	default:
		// bad type
		return
	}

	return
}

/*
TODO: improve this.
*/
func pbRuleCanPush(x any) (ok bool) {
	switch x.(type) {
	case Rule:
		ok = true
	case PermissionBindRule:
		ok = true
	}

	return
}

func listRuleCanPush(x any) (ok bool) {
	switch tv := x.(type) {
	case PermissionBindRule:
		ok = true
	case string:
		if len(tv) > 0 {
			ok = true
		}
	case Rule:
		if tv.Len() > 0 {
			ok = true
		}
	case ObjectIdentifier:
		if !tv.IsZero() {
			ok = true
		}
	}
	return
}

func stackageAnd() stackage.Stack                 { return stackage.And() }
func stackageOr() stackage.Stack                  { return stackage.Or() }
func stackageNot() stackage.Stack                 { return stackage.Not() }
func stackageList(capacity ...int) stackage.Stack { return stackage.List(capacity...) }
func stackageBasic() stackage.Stack		  { return stackage.Basic() }

/*
ruleByLoP returns the requested Rule type, or an "And()" Rule as a fallback. The input
value op should be one (1) value from one (1) of the following lines:

• `NOT`, `!(`

• `OR`, `|(`, `||`

• `AND`, `&(`, `&&`

For LDAP filter symbol operators `!(`, `|(` and `&(`, the padding bit is disabled, and
the lead-once/symbol bits are enabled in the returned Rule instance.

For double-boolean symbol operators `&&` and `||`, the symbol bit is enabled in the
return Rule instance.

For word operators `NOT`, `OR` and `AND`, case-folding is not significant in the matching
process. No additional configuration / presentation options are enabled in the return Rule
instance, unlike the above symbol scenarios.

An unrecognized operator word will return a fallback And() Rule.
*/
func ruleByLoP(op string) Rule {

	if eq(op, `NOT`) || eq(op, `AND NOT`) {
		// NOT (word)
		return Not()
	} else if eq(op, `!`) {
		// NOT (filter)
		return Rule(stackageNot().LeadOnce().Paren().NoPadding().Symbol(`!`)).setCategory(`filter`).setID(`not`)
	} else if eq(op, `OR`) {
		// OR (word)
		return Or()
	} else if eq(op, `|`) {
		// OR (filter)
		return Rule(stackageOr().LeadOnce().Paren().NoPadding().Symbol(`|`)).setCategory(`filter`).setID(`or`)
	} else if eq(op, `||`) {
		// OR (dsymbol)
		return Rule(stackageOr().Symbol(`||`)).setCategory(`or`)
	} else if eq(op, `&`) {
		// AND (filter)
		return Rule(stackageAnd().LeadOnce().Paren().NoPadding().Symbol(`&`)).setCategory(`filter`).setID(`and`)
	} else if eq(op, `&&`) {
		// AND (dsymbol)
		return Rule(stackageAnd().Symbol(`&&`)).setCategory(`and`)
	}

	// Fallback AND (word)
	return And()
}

/*
ruleByDNKeyword returns an instance of Rule based on the following:

• BindGDN (groupdn) keyword returns the Rule instance created by GDNs()

• BindRDN (roledn) keyword returns the Rule instance created by RDNs()

• BindUDN (userdn) keyword returns the Rule instance created by UDNs()

Note that the latter case, BindUDN, is a fallback for any BindKeyword
not matching any of the above.

This function is private and is used only during parsing of bind rules
which permit a list of DNs as a single logical value. It exists mainly
to keep cyclomatics low.
*/
func ruleByDNKeyword(key BindKeyword) Rule {
	// prepare a stack for our DN value(s)
	// based on the input keyword (key)
	if key == BindRDN {
		return RDNs()
	}

	if key == BindGDN {
		return GDNs()
	}

	return UDNs()
}

/*
pushPolicyViolationErr returns an instance of error to be thrown when
a push attempt of instance x is rejected by Rule stack r.
*/
func pushPolicyViolationErr(r Rule, x any, err ...error) error {
	id := r.ID()
	if len(id) == 0 {
		id = `<null>`
	}

	if len(err) > 0 {
		// caller passed a supplemental error to use
		if err[0] != nil {
			return errorf("PushPolicy violation [%T:%s]: %v", r, id, err[0])
		}
	}

	return errorf("PushPolicy violation: Stack %T [id:%s] rejected %T instance", r, id, x)
}

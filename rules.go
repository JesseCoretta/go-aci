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
	return Rule(stackageList(9).
		NoPadding()).setID(`target`).setCategory(`target`).setPushPolicy()
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
	return Rule(stackageAnd()).setID(`bind`).setCategory(`and`).setPushPolicy()
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
	return Rule(stackageOr()).setID(`bind`).setCategory(`or`).setPushPolicy()
}

/*
Not returns an instance of Rule configured to express Boolean NOT logical
operations. Rule instances of this design generally contain Bind Rule
Condition instances, or other nested Boolean-related Rule instances.

The embedded type within the return is stackage.Stack via the go-stackage
package's Not function.
*/
func Not() Rule {
	return Rule(stackageNot()).setID(`bind`).setCategory(`not`).setPushPolicy()
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
func (r Rule) Index(i int) (any, bool) {
	return stackage.Stack(r).Index(i)
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
assigned to the receiver.

For example, if a receiver instance were created with the Ctrls()
package level function, its categorical label `controls` would
result in the creation of a `targetcontrol` Target Rule.
*/
func (r Rule) Eq() (c Condition) {
	if r.IsZero() {
		return Condition{}
	}

	if k := matchTKW(r.Category()); k != TargetKeyword(0x0) {
		c = Cond(k, r.Paren(false), Eq).
			Encap(`"`).Paren().
			setCategory(k.String())
	} else if j := matchBKW(r.Category()); j != BindKeyword(0x0) {
		c = Cond(j, r.Paren(false), Eq).
			Encap(`"`).Paren().
			setCategory(j.String())
	}

	return
}

/*
Ne is a convenience method that crafts a particular negated-equality
Condition instance based upon the categorical string label value
assigned to the receiver.

For example, if a receiver instance were created with the Ctrls()
package level function, its categorical label `controls` would
result in the creation of a `targetcontrol` Target Rule.

Negated equality matching operators should be used with caution.
*/
func (r Rule) Ne() (c Condition) {
	if r.IsZero() {
		return Condition{}
	}

	for k, v := range tkwMap {
		if r.Category() == v {
			c = Cond(k, r.Paren(false), Ne).
				Encap(`"`).Paren().
				setCategory(k.String())
			break
		}
	}

	return
}

func (r Rule) canPush(x any) (err error) {
	var ok bool
	switch r.ID() {
	case `instructions`:
		ok = aciCanPush(x)
	case `parenthetical_bind`,
		`enveloped_parenthetical_bind`,
		`enveloped_bind`,
		`bind`, `and`, `or`, `not`:
		ok = bindRuleCanPush(x)
	case `target`:
		// only allow unique Condition instances which
		// bear the categorical string label `target`
		// to be considered for pushes. Uniqueness is
		// true if the given target keyword does not
		// already reside within a Condition present
		// within the receiver (e.g.: `targetfilter`).
		if ok = targetRuleCanPush(x); ok {
			if found, isCond := targetInRule(r, x); !found && isCond {
				ok = targetRuleCanPush(x)
			} else {
				printf("FAIL1\n")
			}
		} else {
			printf("FAIL0\n")
		}
	case `pb`:
		ok = pbRuleCanPush(x)
	case `list`:
		ok = listRuleCanPush(x)
	}

	if !ok {
		err = errorf("PushPolicy violation: %T (%s) does not allow appends of %T instances", r, r.ID(), x)
		printf("%v\n", err)
	}

	return
}

func (r Rule) setPushPolicy() Rule {
	stackage.Stack(r).SetPushPolicy(r.canPush)
	return r
}

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

func aciCanPush(x any) (ok bool) {
	switch tv := x.(type) {
	case string:
		// TODO - invoke high level ACI parser here.
	case Instruction:
		ok = tv.Valid() == nil
	}

	return
}

func targetInRule(r Rule, t any) (found, isCond bool) {
	var kw string
	switch tv := t.(type) {
	case Condition:
		kw = tv.Keyword()
		// if nothing resides within the
		// Rule, just give it the go-ahead.
		if r.Len() == 0 {
			isCond = true
			return
		}
	default:
		// only Condition instances are
		// valid target rule slices.
		return
	}

	for i := 0; i < r.Len(); i++ {
		var sl any
		if sl, found = r.Index(i); !found {
			continue
		}

		switch tv := sl.(type) {
		case Condition:
			isCond = true
			if found = eq(tv.Keyword(), kw); found {
				return
			}
		case Rule:
			isCond = true
		}
	}

	return
}

func targetRuleCanPush(x any) (ok bool) {
	switch tv := x.(type) {
	case Condition:
		if matchTKW(tv.Keyword()) != TargetKeyword(0x0) {
			ok = true
		}
	case Rule:
		if matchTKW(tv.Category()) != TargetKeyword(0x0) {
			ok = true
		}
	}

	return
}

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

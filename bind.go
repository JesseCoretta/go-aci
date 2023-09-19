package aci

/*
bind.go contains bind rule(s) types, functions and methods.
*/

var (
	badBindRule  BindRule
	badBindRules BindRules
)

/*
BindRuleMethods contains one (1) or more instances of BindRuleMethod,
representing a particular BindRule "builder" method for execution by
the caller.

See the Operators method extended through all eligible types for further
details.
*/
type BindRuleMethods struct {
	*bindRuleFuncMap
}

/*
newBindRuleMethods populates an instance of *bindRuleFuncMap, which
is embedded within the return instance of BindRuleMethods.
*/
func newBindRuleMethods(m bindRuleFuncMap) BindRuleMethods {
	M := make(bindRuleFuncMap, len(m))
	for k, v := range m {
		M[k] = v
	}

	return BindRuleMethods{&M}
}

/*
Index calls the input index (idx) within the internal structure of the
receiver instance. If found, an instance of ComparisonOperator and its
accompanying BindRuleMethod instance are returned.

Valid input index types are integer (int), ComparisonOperator constant
or string identifier. In the case of a string identifier, valid values
are as follows:

• For Eq (1): `=`, `Eq`, `Equal To`

• For Ne (2): `=`, `Ne`, `Not Equal To`

• For Lt (3): `=`, `Lt`, `Less Than`

• For Le (4): `=`, `Le`, `Less Than Or Equal`

• For Gt (5): `=`, `Gt`, `Greater Than`

• For Ge (6): `=`, `Ge`, `Greater Than Or Equal`

Case is not significant in the string matching process.

Please note that use of this method by way of integer or ComparisonOperator
values utilizes fewer resources than a string lookup.

See the ComparisonOperator type's Context, String and Description methods
for accessing the above string values easily.

If the index was not matched, an invalid ComparisonOperator is returned
alongside a nil BindRuleMethod. This will also apply to situations in
which the type instance which crafted the receiver is uninitialized, or
is in an otherwise aberrant state.
*/
func (r BindRuleMethods) Index(idx any) (ComparisonOperator, BindRuleMethod) {
	return r.index(idx)
}

/*
index is a private method called by BindRuleMethods.Index.
*/
func (r BindRuleMethods) index(idx any) (cop ComparisonOperator, meth BindRuleMethod) {
	if r.IsZero() {
		return
	}
	cop = badCop

	// perform a type switch upon the input
	// index type
	switch tv := idx.(type) {

	case ComparisonOperator:
		// cast cop as an int, and make recursive
		// call to this function.
		return r.Index(int(tv))

	case int:
		// there are only six (6) valid
		// operators, numbered one (1)
		// through six (6).
		//
		// this is an unnecessary cyclomatic
		// complexity factor.
		//if !isValidCopNumeral(tv) {
		//	return
		//}

		var found bool
		if meth, found = (*r.bindRuleFuncMap)[ComparisonOperator(tv)]; found {
			cop = ComparisonOperator(tv)
			return
		}

	case string:
		cop, meth = rangeBindRuleFuncMap(tv, r.bindRuleFuncMap)
	}

	return
}

func rangeBindRuleFuncMap(candidate string, fm *bindRuleFuncMap) (cop ComparisonOperator, meth BindRuleMethod) {
	// iterate all map entries, and see if
	// input string value matches the value
	// returned by these three (3) methods:
	for k, v := range *fm {
		if strInSliceFold(candidate, []string{
			k.String(),      // e.g.: "="
			k.Context(),     // e.g.: "Eq"
			k.Description(), // e.g.: "Equal To"
		}) {
			cop = k
			meth = v
			break
		}
	}

	return
}

/*
BR wraps go-stackage's Cond package-level function. In this context,
it is wrapped here to assemble and return a BindRule instance using
the so-called "one-shot" procedure. This is an option only when ALL
information necessary for the process is in-hand and ready for user
input: the keyword, comparison operator and the appropriate value(s)
expression.

Use of this function shall not require a subsequent call of BindRule's
Init method, which is needed only for so-called "piecemeal" BindRule
assembly.

Use of this function is totally optional. Users may, instead, opt to
populate the specific value instance(s) needed and execute the type's
own Eq, Ne, Ge, Gt, Le and Lt methods (when applicable) to produce an
identical return instance. Generally speaking, those methods may prove
to be more convenient -- and far safer -- than use of this function.
*/
func BR(kw, op, ex any) BindRule {
	return newBindRule(kw, op, ex)
}

/*
newBindRule is a private package-level function called by the BR
package-level function.
*/
func newBindRule(kw, op, ex any) (b BindRule) {
	b.Init()
	b.SetKeyword(kw).
		SetOperator(op).
		SetExpression(ex)

	b.cast().
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding)

	return
}

/*
Init wraps stackage.Condition's Init method. This is a required method
for situations involving the piecemeal (step-by-step) assembly of an
instance of BindRule as opposed to a one-shot creation using the BR
package-level function. It is also an ideal means for the creation of
a BindRule instance when one does not immediately possess all of the
needed pieces of information (i.e.: uncertain which keyword to use, or
when an expression value has not yet been determined, etc).

Call this method after a variable declaration but before your first
change, e.g.:

	var br BindRule
	... do other things ...
	... we're ready to set something now ...
	br.Init()
	br.SetKeyword("blarg")
	br.SetSomethingElse(...)
	...

Init need only be executed once within the lifespan of a BindRule
instance. Its execution shall result in a completely new embedded
pointer instance supplanting the previous one.

One may choose, however, to re-execute this method IF this instance
shall be reused (perhaps in a repetative or looped manner), and if
it would be desirable to 'wipe the slate clean' for some reason.
*/
func (r *BindRule) Init() BindRule {
	_r := r.cast()
	if _r.IsZero() || !_r.IsInit() {
		_r.Init()
	}
	*r = BindRule(_r)
	return *r
}

/*
Contains returns a Boolean value indicative of whether the specified ComparisonOperator,
which may be expressed as a string, int or native ComparisonOperator, is allowed for use
by the type instance that created the receiver instance. This method offers a convenient
alternative to the use of the Index method combined with an assertion value (such as Eq,
Ne, "=", "Greater Than", et al).

In other words, if one uses the FQDN.BRM method to create an instance of BindRuleMethods,
feeding Gt (Greater Than) to this method shall return false, as mathematical comparison
does not apply to instances of the FQDN type.
*/
func (r BindRuleMethods) Contains(cop any) bool {
	c, _ := r.index(cop)
	return c.Valid() == nil
}

/*
IsZero returns a Boolean value indicative of whether the receiver is
nil, or unset.
*/
func (r BindRuleMethods) IsZero() bool {
	return r.bindRuleFuncMap == nil
}

/*
Valid returns the first encountered error returned as a result of
execution of the first available BindRuleMethod instance. This is
useful in cases where a user wants to see if the desired instance(s)
of BindRuleMethod will produce a usable result.
*/
func (r BindRuleMethods) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
		return
	}

	// Eq is always available for all eligible
	// types, so let's use that unconditionally.
	// If any one method works, then all of them
	// will work.
	_, meth := r.Index(Eq)
	err = meth().Valid()
	return
}

/*
Len returns the integer length of the receiver. Note that the return
value will NEVER be less than zero (0) nor greater than six (6).
*/
func (r BindRuleMethods) Len() int {
	if r.IsZero() {
		return 0
	}

	return len((*r.bindRuleFuncMap))
}

/*
BindRuleMethod is the closure signature for methods used to build
new instances of BindRule.

The signature is qualified by the following methods extended through
all eligible types defined in this package:

• Eq

• Ne

• Lt

• Le

• Gt

• Ge

Note that certain types only support a subset of the above list. Very
few types support all of the above.
*/
type BindRuleMethod func() BindRule

/*
bindRuleFuncMap is a private type intended to be used within
instances of BindRuleMethods.
*/
type bindRuleFuncMap map[ComparisonOperator]BindRuleMethod

func (r BindRule) isBindContextQualifier() {}

/*
Traverse returns the receiver instance. This method only exists to satisfy
Go's interface signature requirements for the BindContext interface. See
BindRules.Traverse instead.
*/
func (r BindRule) Traverse(indices ...int) BindContext {
	return r
}

/*
Index returns the receiver instance. This method only exists to satisfy
Go's interface signature requirements for the BindContext interface. See
BindRules.Index instead.
*/
func (r BindRule) Index(_ int) BindContext {
	return r
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r BindRule) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Kind returns the string literal `condition` to identify the receiver
as a stackage.Condition type alias.
*/
func (r BindRule) Kind() string {
	return `condition`
}

/*
Len does not perform any useful task, and exists only to satisfy Go's
interface signature requirements and to convey this message.

A value of zero (0) is returned if the receiver instance is nil. A
value of one (1) otherwise.
*/
func (r BindRule) Len() int {
	if r.IsZero() {
		return 0
	}
	return 1
}

/*
IsNesting does not perform any useful task, and exists only to satisfy
Go's interface signature requirements and to convey this message.

A Boolean value of false is returned in any scenario.
*/
func (r BindRule) IsNesting() bool {
	return false
}

/*
Paren wraps go-stackage's Condition.Paren method.
*/
func (r BindRule) Paren(state ...bool) BindRule {
	r.cast().Paren(state...)
	return r
}

/*
IsParen wraps go-stackage's Condition.IsParen method.
*/
func (r BindRule) IsParen() bool {
	return r.cast().IsParen()
}

/*
Valid wraps go-stackage's Condition.Valid method.
*/
func (r BindRule) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
		return
	}

	err = r.cast().Valid()
	return
}

/*
ID wraps go-stackage's Stack.ID method.
*/
func (r BindRule) ID() string {
	if r.IsZero() {
		return bindRuleID
	}
	return r.cast().ID()
}

/*
Category wraps go-stackage's Stack.Category method.
*/
func (r BindRule) Category() string {
	if r.IsZero() {
		return ``
	}
	return r.Keyword().String()
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Condition.String method.
*/
func (r BindRule) String() string {
	if r.IsZero() {
		return ``
	}
	return r.cast().String()
}

/*
NoPadding wraps go-stackage's Condition.NoPadding method.
*/
func (r BindRule) NoPadding(state ...bool) BindRule {
	if r.IsZero() {
		return r
	}

	r.cast().NoPadding(state...)
	return r
}

/*
Keyword wraps go-stackage's Condition.Keyword method and
resolves the raw value into a BindKeyword. Failure to do
so will return a bogus Keyword.
*/
func (r BindRule) Keyword() Keyword {
	k := r.cast().Keyword()
	var kw any = matchBKW(k)
	return kw.(BindKeyword)
}

/*
Operator wraps go-stackage's Condition.Operator method
and casts the stackage.ComparisonOperator to the local
aci.ComparisonOperator.
*/
func (r BindRule) Operator() ComparisonOperator {
	return castCop(r.cast().Operator())
}

/*
Expression wraps go-stackage's Condition.Expression method.
*/
func (r BindRule) Expression() any {
	return r.cast().Expression()
}

/*
IsZero wraps go-stackage's Condition.IsZero method.
*/
func (r BindRule) IsZero() bool {
	return r.cast().IsZero()
}

func (r BindRules) isBindContextQualifier() {}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r BindRules) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Kind returns the string literal `stack` to identify the receiver as
a stackage.Stack type alias.
*/
func (r BindRules) Kind() string {
	return `stack`
}

/*
And returns an instance of Rule configured to express Boolean AND logical operations.
Instances of this design contain BindContext instances, which are qualified through
instances of the following types:

• BindRule

• BindRules

Optionally, the caller may choose to submit one (1) or more (valid) instances of these
types during initialization. This is merely a more convenient alternative to separate
initialization and push procedures.

The embedded type within the return is stackage.Stack via the go-stackage
package's And function.
*/
func And(x ...any) (b BindRules) {
	// create a native stackage.Stack
	// and configure before typecast.
	_b := stackAnd().
		SetID(bindRuleID).
		SetCategory(`and`).
		NoPadding(!StackPadding)

	// cast _a as a proper BindRules instance
	// (b). We do it this way to gain access
	// to the method for the *specific instance*
	// being created (b), thus allowing things
	// like uniqueness checks, etc., to occur
	// during push attempts, providing helpful
	// and non-generalized feedback.
	b = BindRules(_b)
	_b.SetPushPolicy(b.pushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_b.Push(x...)

	return
}

/*
Or returns an instance of Rule configured to express Boolean OR logical operations.
Instances of this design contain BindContext instances, which are qualified through
instances of the following types:

• BindRule

• BindRules

Optionally, the caller may choose to submit one (1) or more (valid) instances of these
types during initialization. This is merely a more convenient alternative to separate
initialization and push procedures.

The embedded type within the return is stackage.Stack via the go-stackage
package's Or function.
*/
func Or(x ...any) (b BindRules) {
	// create a native stackage.Stack
	// and configure before typecast.
	_b := stackOr().
		SetID(bindRuleID).
		SetCategory(`or`).
		NoPadding(!StackPadding)

	// cast _a as a proper BindRules instance
	// (b). We do it this way to gain access
	// to the method for the *specific instance*
	// being created (b), thus allowing things
	// like uniqueness checks, etc., to occur
	// during push attempts, providing helpful
	// and non-generalized feedback.
	b = BindRules(_b)
	_b.SetPushPolicy(b.pushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_b.Push(x...)

	return
}

/*
Not returns an instance of Rule configured to express Boolean NOT logical operations.
Instances of this design contain BindContext instances, which are qualified through
instances of the following types:

• BindRule

• BindRules

Optionally, the caller may choose to submit one (1) or more (valid) instances of these
types during initialization. This is merely a more convenient alternative to separate
initialization and push procedures.

The embedded type within the return is stackage.Stack via the go-stackage
package's Not function.
*/
func Not(x ...any) (b BindRules) {
	// create a native stackage.Stack
	// and configure before typecast.
	_b := stackNot().
		SetID(bindRuleID).
		SetCategory(`not`).
		NoPadding(!StackPadding)

	// cast _a as a proper BindRules instance
	// (b). We do it this way to gain access
	// to the method for the *specific instance*
	// being created (b), thus allowing things
	// like uniqueness checks, etc., to occur
	// during push attempts, providing helpful
	// and non-generalized feedback.
	b = BindRules(_b)
	_b.SetPushPolicy(b.pushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_b.Push(x...)

	return
}

/*
convertBindRulesHierarchy processes the orig input instance and casts
its contents in the following manner:

	stackage.Stack		-->	BindRules
	       \			 	 \
	        +- ...                           + ...
	        |				  |
	        +- stackage.Condition	 -->	  +- BindRule

The hierarchy is traversed thoroughly and will handle nested contexts
seamlessly.

This function is called following an apparently successful BindRules
parsing request through go-antlraci.
*/
func convertBindRulesHierarchy(stack any) (BindContext, bool) {
	orig, _ := castAsStack(stack)
	if orig.Len() == 0 {
		return badBindRules, false
	}

	var clean BindRules

	// Obtain the kind string from the
	// original stack.
	clean, ok := wordToStack(orig.Kind())

	// Iterate the newly-populated clean
	// instance, performing type-casting
	// as needed, possibly in recursion.
	for i := 0; i < orig.Len() && ok; i++ {
		slice, _ := orig.Index(i)

		// perform a type switch upon the
		// slice member @ index i. There
		// are two (2) valid types we may
		// encounter ...
		switch {

		// slice is a stackage.Condition.
		// We want to cast to a BindRule
		// instance, and update the string
		// value(s) to be housed within a
		// value-appropriate type defined
		// by go-aci.
		case isStackageCondition(slice):
			deref := derefC(slice)
			ntv := BindRule(deref).
				Paren(deref.IsParen())

			// Extract individual expression value
			// from BindRule (ntv), and recreate it
			// using the proper type, replacing the
			// original. For example, a User DN Bind
			// Rule with a RuleExpression value of:
			//
			//   []string{<dn1>,<dn2>,<dn3>}
			//
			// ... shall be replaced with:
			//
			//   <stack alias type>-idx#------val-
			//   DistinguishedNames[<N1>] -> <dn1>
			//                     [<N2>] -> <dn2>
			//                     [<N3>] -> <dn3>
			if err := ntv.assertExpressionValue(); err == nil {
				clean.Push(ntv)
				continue
			}
			break

		// slice is a stackage.Stack instance.
		// We want to cast to a BindRules type
		// instance, but in order to do that,
		// we'll recurse into this same function
		// using this slice as the subordinate
		// 'orig' input value.
		case isStackageStack(slice):
			stk, _ := castAsStack(slice)
			paren := stk.IsParen()
			if sub, subok := convertBindRulesHierarchy(slice); subok {
				if _, ok := sub.(BindRules); ok {
					sub.(BindRules).Paren(paren)
					uncloakBindRules(sub.(BindRules))
				}
				clean.Push(sub)
				continue
			}
			return badBindRules, false
		}
	}

	// A cheap and easy means of ensuring
	// the content really did transfer and
	// [re]cast properly, and that nothing
	// was missed.
	//ok = orig.String() == clean.String()
	ok = len(clean.String()) > 0

	// uncloak any hidden stacks
	uncloakBindRules(clean)

	return clean, ok
}

func uncloakBindRules(ctx BindRules) {
	for i := 0; i < ctx.Len(); i++ {
		if slice := ctx.Index(i); isCloaked(slice) {
			stack := uncloak(slice.(BindRules))
			ctx.Replace(stack, i)
		}
	}
}

/*
BUG: testing cornercase for go-antlraci. Temporary, do not remove yet.
*/
func uncloak(ctx BindRules) BindRules {
	stack := ctx.Index(0)
	if isCloaked(stack) {
		stack = uncloak(stack.(BindRules))
	}

	return stack.(BindRules)
}

/*
BUG: testing cornercase for go-antlraci. Temporary, do not remove yet.
*/
func isCloaked(x BindContext) bool {
	switch tv := x.(type) {
	case BindRules:
		if hasSfx(tv.Category(), `not`) || tv.Len() != 1 {
			break
		}

		if tv.Index(0).Kind() == `stack` {
			return !tv.IsParen()
		}
	}
	return false
}

func wordToStack(k string) (BindRules, bool) {
	// Perform an anonymous switch, allowing
	// the evaluation of the Boolean logical
	// "disposition" of the (outer) Bind Rules
	// instance "kind".
	switch {

	// Negated (NOT, AND NOT) operator
	case hasSfx(uc(k), `NOT`):
		return Not(), true

	// ANDed operator
	case eq(k, `AND`):
		return And(), true

	// ORed operator
	case eq(k, `OR`):
		return Or(), true
	}

	// unsupported operator
	return badBindRules, false
}

/*
SetKeyword wraps go-stackage's Condition.SetKeyword method.
*/
func (r BindRule) SetKeyword(kw any) BindRule {
	cac := r.cast()
	if !cac.IsInit() {
		r.Init()
	}
	cac.SetKeyword(kw)
	return r
}

/*
SetOperator wraps go-stackage's Condition.SetOperator method.
*/
func (r BindRule) SetOperator(op any) BindRule {
	var cop ComparisonOperator
	switch tv := op.(type) {
	case string:
		cop = matchCOP(tv)
	case ComparisonOperator:
		cop = tv
	}

	// ALL Target and Bind rules accept Eq,
	// so only scrutinize the operator if
	// it is something *other than* that.
	if cop != Eq {
		if !keywordAllowsComparisonOperator(r.Keyword(), op) {
			return r
		}
	}

	// operator not known? bail out
	if cop == ComparisonOperator(0) {
		return r
	}

	// not initialized? bail out
	if !r.cast().IsInit() {
		return r
	}

	// cast to stackage.Condition and
	// set operator value.
	r.cast().SetOperator(cop)

	return r
}

/*
SetExpression wraps go-stackage's Condition.SetExpression method.
*/
func (r BindRule) SetExpression(expr any) BindRule {
	cac := r.cast()
	if !cac.IsInit() {
		cac.Init()
	}
	cac.SetExpression(expr)
	return r
}

/*
SetQuoteStyle allows the election of a particular multivalued
quotation style offered by the various adopters of the ACIv3
syntax. In the context of a BindRule, this will only have a
meaningful impact if the keyword for the receiver is one (1)
of the following:

  - BindUDN (userdn)
  - BindRDN (roledn)
  - BindGDN (groupdn)

Additionally, the underlying type set as the expression value
within the receiver MUST be a BindDistinguishedNames instance
with two (2) or more distinguished names within.

See the const definitions for MultivalOuterQuotes (default)
and MultivalSliceQuotes for details.
*/
func (r BindRule) SetQuoteStyle(style int) BindRule {
	key := r.Keyword()

	switch tv := r.Expression().(type) {
	case BindDistinguishedNames:
		switch key {
		case BindUDN, BindGDN, BindRDN:
			tv.setQuoteStyle(style)
		}
	default:
		r.cast().Encap(`"`)
		return r
	}

	// Toggle the individual value quotation scheme
	// to the INVERSE of the Stack quotation scheme
	// set above.
	//
	// If MultivalSliceQuotes equals the style set
	// by the user, this implies that that no outer
	// encapsulation shall be used, thus _r.Encap()
	// is called for the receiver.
	//
	// But the above type instances (TDNs, OIDs, ATs)
	// will have the opposite setting imposed, which
	// enables quotation for the individual values.
	if style == MultivalSliceQuotes {
		r.cast().Encap()
	} else {
		r.cast().Encap(`"`)
	}

	return r

}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Stack.String method.
*/
func (r BindRules) String() string {
	return r.cast().String()
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r BindRules) IsZero() bool {
	return r.cast().IsZero()
}

/*
reset wraps go-stackage's Stack.Reset method. This is a private
method in go-aci.
*/
func (r BindRules) reset() {
	r.cast().Reset()
}

/*
ID wraps go-stackage's Stack.ID method.
*/
func (r BindRules) ID() string {
	return bindRuleID
}

/*
Category wraps go-stackage's Stack.Category method.
*/
func (r BindRules) Category() string {
	if r.IsZero() {
		return `<uninitialized_bindrules>`
	}
	return r.cast().Category()
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r BindRules) Len() int {
	if r.IsZero() {
		return 0
	}
	return r.cast().Len()
}

/*
IsNesting wraps go-stackage's Stack.IsNesting method.
*/
func (r BindRules) IsNesting() bool {
	if r.IsZero() {
		return false
	}
	return r.cast().IsNesting()
}

/*
Keyword wraps go-stackage's Stack.Category method and
resolves the raw value into a BindKeyword. Failure to
do so will return a bogus Keyword.
*/
func (r BindRules) Keyword() Keyword {
	var kw any = matchBKW(r.cast().Category())
	return kw.(BindKeyword)
}

/*
Push wraps go-stackage's Stack.Push method.
*/
func (r BindRules) Push(x ...any) BindRules {
	r.cast().Push(x...)
	return r
}

/*
Pop wraps go-stackage's Stack.Pop method. An instance of
BindContext, which may or may not be nil, is returned
following a call of this method.

Within the context of the receiver type, a BindContext, if
non-nil, can represent any of the following instance types:

• BindRule

• BindRules
*/
func (r BindRules) Pop() BindContext {
	return r.pop()
}

func (r BindRules) pop() (popped BindContext) {
	if r.IsZero() {
		return nil
	}

	_popped, _ := r.cast().Pop()
	switch tv := _popped.(type) {
	case BindRule:
		popped = tv
	case BindRules:
		popped = tv
	}

	return
}

/*
remove wraps go-stackage's Stack.Remove method.
*/
func (r BindRules) remove(idx int) (ok bool) {
	_, ok = r.cast().Remove(idx)
	return
}

/*
Replace wraps go-stackage's Stack.Replace method.
*/
func (r BindRules) Replace(x any, idx int) BindRules {
	return r.replace(x, idx)
}

/*
replace is a private method called by BindRules.Replace
as well as certain ANTLR->ACI parsing procedures.
*/
func (r BindRules) replace(x any, idx int) BindRules {
	if !r.IsZero() {
		r.cast().Replace(x, idx)
	}

	return r
}

/*
Index wraps go-stackage's Stack.Index method.
*/
func (r BindRules) Index(idx int) (ctx BindContext) {
	y, _ := r.cast().Index(idx)

	switch tv := y.(type) {
	case BindRule:
		ctx = tv
	case BindRules:
		ctx = tv
	}

	return
}

/*
ReadOnly wraps go-stackage's Stack.ReadOnly method.
*/
func (r BindRules) ReadOnly(state ...bool) BindRules {
	r.cast().ReadOnly(state...)
	return r
}

/*
Paren wraps go-stackage's Stack.Paren method.
*/
func (r BindRules) Paren(state ...bool) BindRules {
	r.cast().Paren(state...)
	return r
}

/*
IsParen wraps go-stackage's Stack.IsParen method.
*/
func (r BindRules) IsParen() bool {
	return r.cast().IsParen()
}

/*
Fold wraps go-stackage's Stack.Fold method to allow the case
folding of logical Boolean 'AND', 'OR' and 'AND NOT' WORD
operators to 'and', 'or' and 'and not' respectively, or vice
versa.
*/
func (r BindRules) Fold(state ...bool) BindRules {
	r.cast().Fold(state...)
	return r
}

/*
insert wraps go-stackage's Stack.Insert method.
*/
func (r BindRules) insert(x any, left int) (ok bool) {
	switch tv := x.(type) {
	case BindRule, BindRules:
		ok = r.cast().Insert(tv, left)
	}

	return
}

/*
NoPadding wraps go-stackage's Stack.NoPadding method.
*/
func (r BindRules) NoPadding(state ...bool) BindRules {
	r.cast().NoPadding(state...)
	return r
}

/*
Traverse wraps go-stackage's Stack.Traverse method.
*/
func (r BindRules) Traverse(indices ...int) (B BindContext) {
	br, ok := r.cast().Traverse(indices...)
	if ok {
		switch tv := br.(type) {
		case BindContext:
			B = tv
		default:
			B = castAsBindRule(tv)
		}
	}

	if B == nil {
		B = badBindRules
	}

	return
}

/*
Valid wraps go-stackage's Stack.Valid method.
*/
func (r BindRules) Valid() (err error) {
	err = r.cast().Valid()
	return
}

/*
pushPolicy conforms to the PushPolicy signature defined
within go-stackage. This function will be called privately
whenever an instance is pushed into a particular Stack (or
Stack alias) type instance.

Only BindContext qualifiers are to be cleared for push.
*/
func (r BindRules) pushPolicy(x ...any) (err error) {
	// perform type switch upon input value
	// x to determine suitability for push.
	switch tv := x[0].(type) {
	case BindRules:
		if err = tv.Valid(); err != nil {
			err = pushErrorNilOrZero(r, tv, matchBKW(r.Category()), err)
		}
	case BindRule:
		if err = tv.Valid(); err != nil {
			err = pushErrorNilOrZero(r, tv, matchBKW(r.Category()), err)
		}

		if tv.Keyword() != nil {
			err = badPTBRuleKeywordErr(tv, `bind`, `bindkeyword`, tv.Keyword())
			if matchBKW(tv.Keyword().String()) != BindKeyword(0x0) {
				err = nil
			}
		}

	default:
		// unsuitable candidate per type
		err = pushErrorBadType(r, tv, matchBKW(r.Category()))
	}

	return
}

/*
BindContext is a convenient interface type that is qualified by
the following types:

• BindRule

• BindRules

The qualifying methods shown below are intended to make the
handling of a structure of (likely nested) BindRules instances
slightly easier without an absolute need for type assertion at
every step. These methods are inherently read-only in nature
and represent only a subset of the available methods exported
by the underlying qualifier types.

To alter the underlying value, or to gain access to all of a
given type's methods, type assertion of qualifying instances
shall be necessary.
*/
type BindContext interface {
	// String returns the string representation of the
	// receiver instance.
	String() string

	// Keyword returns the BindKeyword, enveloped as a
	// Keyword interface value. If the receiver is an
	// instance of BindRule, the value is derived from
	// the Keyword method. If the receiver is an instance
	// of BindRules, the value is derived (and resolved)
	// using the Category method.
	Keyword() Keyword

	// IsZero returns a Boolean value indicative of the
	// receiver instance being nil, or unset.
	IsZero() bool

	// IsParen returns a Boolean value indicative of
	// whether the receiver is parenthetical.
	IsParen() bool

	// Len returns the integer length of the receiver.
	// Only meaningful when run on BindRules instances.
	Len() int

	// IsNesting returns a Boolean value indicative of
	// whether the receiver contains a stack as a value.
	// Only meaningful when run on BindRules instances.
	IsNesting() bool

	// Index returns the Nth slice from the receiver.
	// This method is only meaningful if the underlying
	// instance is BindRules.
	Index(int) BindContext

	// Traverse will walk a structure of BindContext
	// instances using a sequence of index integers.
	// An instance of BindContext is returned, or nil.
	Traverse(...int) BindContext

	// Valid returns an error instance that indicates
	// whether the receiver is in an aberrant state.
	Valid() error

	// ID will report `bind` in all scenarios.
	ID() string

	// Category will report the logical state of a BindRules
	// instance. This will read `and`, `or`, `not`. For an
	// instance of BindRule, however, it will return the string
	// form of the Keyword, if set, else a zero string.
	Category() string

	// Kind will report `stack` for a BindRules instance, or
	// `condition` for a BindRule instance
	Kind() string

	// isBindContextQualifier ensures no greedy interface
	// matching outside of the realm of bind rules. It need
	// not be accessed by users, nor is it run at any time
	// outside of unit tests to satisfy code coverage ...
	isBindContextQualifier()
}

const bindRuleID = `bind`

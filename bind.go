package aci

/*
bind.go contains bind rule(s) types, functions and methods.
*/

import (
	parser "github.com/JesseCoretta/go-antlraci"
)

var (
	badBindRule  BindRule
	badBindRules BindRules
)

func (r BindRule) isBindContextQualifier() bool {
	return true
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

An integer value of one (1) is returned in any scenario.
*/
func (r BindRule) Len() int {
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
setID wraps go-stackage's Condition.SetID method.
*/
func (r BindRule) setID(id string) {
	castAsCondition(r).SetID(id)
}

/*
setCategory wraps go-stackage's Condition.SetCategory method.
*/
func (r BindRule) setCategory(cat string) {
	castAsCondition(r).SetCategory(cat)
}

/*
Paren wraps go-stackage's Condition.Paren method.
*/
func (r BindRule) Paren(state ...bool) BindRule {
	castAsCondition(r).Paren(state...)
	return r
}

/*
Valid wraps go-stackage's Condition.Valid method.
*/
func (r BindRule) Valid() (err error) {
	_t := castAsCondition(r)
	err = _t.Valid()
	return
}

/*
ID wraps go-stackage's Stack.ID method.
*/
func (r BindRule) ID() string {
	if r.IsZero() {
		return ``
	}
	return castAsCondition(r).ID()
}

/*
Category wraps go-stackage's Stack.Category method.
*/
func (r BindRule) Category() string {
	if r.IsZero() {
		return ``
	}
	return castAsCondition(r).Category()
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Condition.String method.
*/
func (r BindRule) String() string {
	return castAsCondition(r).String()
}

/*
Keyword wraps go-stackage's Condition.Keyword method and
resolves the raw value into a BindKeyword. Failure to do
so will return a bogus Keyword.
*/
func (r BindRule) Keyword() Keyword {
	k := castAsCondition(r).Keyword()
	var kw any = matchBKW(k)
	return kw.(BindKeyword)
}

/*
Operator wraps go-stackage's Condition.Operator method.
*/
func (r BindRule) Operator() ComparisonOperator {
	x := castAsCop(castAsCondition(r).Operator().(ComparisonOperator))
	return ComparisonOperator(x)
}

/*
Expression wraps go-stackage's Condition.Expression method.
*/
func (r BindRule) Expression() any {
	return castAsCondition(r).Expression()
}

/*
IsZero wraps go-stackage's Condition.IsZero method.
*/
func (r BindRule) IsZero() bool {
	return castAsCondition(r).IsZero()
}

func (r BindRules) isBindContextQualifier() bool {
	return true
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
		NoPadding(!StackPadding).
		SetCategory(TargetAttr.String())

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
		NoPadding(!StackPadding).
		SetCategory(TargetAttr.String())

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
		NoPadding(!StackPadding).
		SetCategory(TargetAttr.String())

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
func convertBindRulesHierarchy(stack any) (BindRules, bool) {
	orig, _ := castAsStack(stack)
	if orig.Len() == 0 {
		return badBindRules, false
	}

	var clean BindRules

	// Obtain the kind string from the
	// original stack.
	k := trimS(orig.Kind())

	clean, ok := wordToStack(k)
	if !ok {
		return badBindRules, ok
	}

	// transfer (copy) references from
	// orig into clean.
	z, _ := castAsStack(clean)
	orig.Transfer(z)
	clean = BindRules(z)

	// Iterate the newly-populated clean
	// instance, performing type-casting
	// as needed, possibly in recursion.
	for i := 0; i < clean.Len(); i++ {

		// perform a type switch upon the
		// slice member @ index i. There
		// are two (2) valid types we may
		// encounter ...
		switch tv := clean.Index(i).(type) {

		// slice is a stackage.Condition.
		// We want to cast to a BindRule
		// instance, and update the string
		// value(s) to be housed within a
		// value-appropriate type defined
		// by go-aci.
		case BindRule:
			ntv := tv

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
			if err := ntv.assertExpressionValue(); err != nil {
				return badBindRules, false
			}

			// overwrite old (tv @ index i) with new (ntv)
			clean.replace(ntv, i)

		// slice is a stackage.Stack instance.
		// We want to cast to a BindRules type
		// instance, but in order to do that,
		// we'll recurse into this same function
		// using this slice as the subordinate
		// 'orig' input value.
		case BindRules:
			stk, _ := castAsStack(tv)
			sub, subok := convertBindRulesHierarchy(stk)
			if !subok {
				return badBindRules, false
			}
			clean.replace(sub, i) // overwrite
		}
	}

	// A cheap and easy means of ensuring
	// the content really did transfer and
	// [re]cast properly, and that nothing
	// was missed.
	ok = orig.String() == clean.String()
	return clean, ok
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
func (r *BindRule) SetKeyword(kw any) *BindRule {
	x := castAsCondition(*r)
	x.SetKeyword(kw)
	*r = BindRule(*x)

	return r
}

/*
SetOperator wraps go-stackage's Condition.SetOperator method.
*/
func (r *BindRule) SetOperator(op ComparisonOperator) *BindRule {
	if !keywordAllowsComparisonOperator(r.Keyword(), op) {
		return r
	}

	x := castAsCondition(*r)
	x.SetOperator(castAsCop(op))
	*r = BindRule(*x)

	return r
}

/*
SetExpression wraps go-stackage's Condition.SetExpression method.
*/
func (r *BindRule) SetExpression(expr any) *BindRule {
	x := castAsCondition(*r)
	x.SetExpression(expr)
	*r = BindRule(*x)

	return r
}

/*
SetQuoteStyle allows the election of a particular multivalued
quotation style offered by the various adopters of the ACIv3
syntax.

See the const definitions for MultivalOuterQuotes (default)
and MultivalSliceQuotes for details.

Note that this will only have an effect on BindRule instances
that bear the userdn, roledn or groupdn keyword contexts AND
contain a stack-based expression containing more than one (1)
slice elements.
*/
func (r BindRule) SetQuoteStyle(style int) BindRule {
	_r := castAsCondition(r)

	switch key := r.Keyword(); key {
	case BindUDN, BindGDN, BindRDN:
		if isStackageStack(_r.Expression()) {
			tv, _ := castAsStack(_r.Expression())
			BindDistinguishedNames(tv).setQuoteStyle(style)

			// Toggle the individual value quotation scheme
			// to the INVERSE of the Stack quotation scheme
			// set above
			if style == MultivalSliceQuotes {
				_r.Encap()
			} else {
				_r.Encap(`"`)
			}
		}
	}

	return r
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Stack.String method.
*/
func (r BindRules) String() string {
	_b, _ := castAsStack(r)
	return _b.String()
}

/*
setCategory wraps go-stackage's Stack.SetCategory method.
*/
func (r BindRules) setCategory(cat string) {
	_b, _ := castAsStack(r)
	_b.SetCategory(cat)
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r BindRules) IsZero() bool {
	_b, _ := castAsStack(r)
	return _b.IsZero()
}

/*
reset wraps go-stackage's Stack.Reset method. This is a private
method in go-aci.
*/
func (r BindRules) reset() {
	_b, _ := castAsStack(r)
	_b.Reset()
}

/*
ID wraps go-stackage's Stack.ID method.
*/
func (r BindRules) ID() string {
	if r.IsZero() {
		return ``
	}
	_b, _ := castAsStack(r)
	return _b.ID()
}

/*
Category wraps go-stackage's Stack.Category method.
*/
func (r BindRules) Category() string {
	if r.IsZero() {
		return ``
	}
	_b, _ := castAsStack(r)
	return _b.Category()
}

/*
setID wraps go-stackage's Stack.SetID method.
*/
func (r BindRules) setID(id string) {
	_b, _ := castAsStack(r)
	_b.SetID(id)
	//r = BindRules(_b)
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r BindRules) Len() int {
	if r.IsZero() {
		return 0
	}
	_b, _ := castAsStack(r)
	return _b.Len()
}

/*
IsNesting wraps go-stackage's Stack.IsNesting method.
*/
func (r BindRules) IsNesting() bool {
	if r.IsZero() {
		return false
	}
	_b, _ := castAsStack(r)
	return _b.IsNesting()
}

/*
Keyword wraps go-stackage's Stack.Category method and
resolves the raw value into a BindKeyword. Failure to
do so will return a bogus Keyword.
*/
func (r BindRules) Keyword() Keyword {
	if meth := getCategoryFunc(r); meth != nil {
		var kw any = matchBKW(meth())
		return kw.(BindKeyword)
	}

	return nil
}

/*
Push wraps go-stackage's Stack.Push method.
*/
func (r BindRules) Push(x ...BindContext) BindRules {
	_r, _ := castAsStack(r)

	// iterate variadic input arguments
	for i := 0; i < len(x); i++ {
		_r.Push(x[i])
	}

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
	_r, _ := castAsStack(r)
	x, _ := _r.Pop()

	var z any
	switch tv := x.(type) {
	case BindRule:
		z = tv
		return z.(*BindRule)
	case BindRules:
		z = tv
		return z.(BindRules)
	}

	return nil
}

/*
remove wraps go-stackage's Stack.Remove method.
*/
func (r BindRules) remove(idx int) bool {
	_b, _ := castAsStack(r)
	_, ok := _b.Remove(idx)
	return ok
}

/*
Index wraps go-stackage's Stack.Index method.
*/
func (r BindRules) Index(idx int) BindContext {
	_r, _ := castAsStack(r)
	y, _ := _r.Index(idx)

	var z any
	switch tv := y.(type) {
	case BindRule:
		z = tv
		return z.(*BindRule)
	case BindRules:
		z = tv
		return z.(BindRules)
	}

	return nil
}

/*
ReadOnly wraps go-stackage's Stack.ReadOnly method.
*/
func (r BindRules) ReadOnly(state ...bool) BindRules {
	_r, _ := castAsStack(r)
	_r.ReadOnly(state...)

	return r
}

/*
Paren wraps go-stackage's Stack.Paren method.
*/
func (r BindRules) Paren(state ...bool) BindRules {
	_r, _ := castAsStack(r)
	_r.Paren(state...)

	return r
}

/*
Fold wraps go-stackage's Stack.Fold method to allow the case
folding of logical Boolean 'AND', 'OR' and 'AND NOT' WORD
operators to 'and', 'or' and 'and not' respectively, or vice
versa.
*/
func (r BindRules) Fold(state ...bool) BindRules {
	_r, _ := castAsStack(r)
	_r.Fold(state...)

	return r
}

/*
insert wraps go-stackage's Stack.Insert method.
*/
func (r BindRules) insert(x any, left int) (ok bool) {
	_t, _ := castAsStack(r)

	switch tv := x.(type) {
	case BindRule, BindRules:
		ok = _t.Insert(tv, left)
	default:
		return
	}

	return
}

/*
replace wraps go-stackage's Stack.Replace method.
*/
func (r BindRules) replace(x any, idx int) (ok bool) {
	_b, _ := castAsStack(r)

	switch tv := x.(type) {
	case BindRule, BindRules:
		ok = _b.Replace(tv, idx)
	default:
		return
	}

	return
}

/*
NoPadding wraps go-stackage's Stack.NoPadding method.
*/
func (r BindRules) NoPadding(state ...bool) BindRules {
	_b, _ := castAsStack(r)
	_b.NoPadding(state...)

	return r
}

/*
Traverse wraps go-stackage's Stack.Traverse method.
*/
func (r BindRules) Traverse(indices ...int) (any, bool) {
	_b, _ := castAsStack(r)
	if br, ok := _b.Traverse(indices...); ok {
		var x any
		if x, ok = br.(BindContext); ok {
			return x, ok
		}
	}

	return nil, false
}

/*
Valid wraps go-stackage's Stack.Valid method.
*/
func (r BindRules) Valid() (err error) {
	_b, _ := castAsStack(r)
	err = _b.Valid()
	return
}

/*
pushPolicy conforms to the PushPolicy signature defined
within go-stackage. This function will be called privately
whenever an instance is pushed into a particular Stack (or
Stack alias) type instance.

Only BindContext qualifiers are to be cleared for push.
*/
func (r BindRules) pushPolicy(x any) (err error) {
	// perform type switch upon input value
	// x to determine suitability for push.
	switch tv := x.(type) {

	case BindContext:
		// BindContext match is qualified
		// through instances of either
		// BindRule or BindRules types.
		if tv.IsZero() {
			err = pushErrorNilOrZero(r, tv, matchBKW(r.Category()))
		}

	default:
		// unsuitable candidate per type
		err = pushErrorBadType(r, tv, matchBKW(r.Category()))
	}

	return
}

/*
ParseBindRule returns an instance of Condition alongside an error instance.

The returned Condition instance shall contain
*/
func ParseBindRule(raw string) (BindRule, error) {
	return parseBindRule(raw)
}

func parseBindRule(raw string) (BindRule, error) {
	_c, err := parser.ParseBindRule(raw)
	return BindRule(_c), err
}

/*
ParseBindRules returns an instance of Rule alongside an error instance.

The returned Rule instance shall contain a complete hierarchical stack
structure that represents the abstract rule (raw) input by the user.
*/
func ParseBindRules(raw string) (BindRules, error) {
	return parseBindRules(raw)
}

func parseBindRules(raw string) (BindRules, error) {
	// send the raw textual bind rules
	// statement(s) to our sister package
	// go-antlraci, call ParseBindRules.
	_b, err := parser.ParseBindRules(raw)
	if err != nil {
		return badBindRules, err
	}

	// Process the hierarchy, converting
	// Stack to BindRules and Condition
	// to BindRule. In addition, we'll
	// replace the parser.ExpressionValue
	// type with more appropriate types
	// defined in this package.
	n, ok := convertBindRulesHierarchy(_b)
	if !ok {
		return badBindRules, parseBindRulesHierErr(_b, n)
	}

	return n, nil
}

/*
assertExpressionValue will update the underlying go-antlraci temporary type with a
proper value-appropriate type defined within the go-aci package. An error is returned
upon processing completion.
*/
func (r BindRule) assertExpressionValue() (err error) {

	// grab the raw value from the receiver. If it is
	// NOT an instance of parser.RuleExpression, then
	// bail out.
	expr, ok := r.Expression().(parser.RuleExpression)
	if !ok {
		err = parseBindRuleInvalidExprTypeErr(r, expr, parser.RuleExpression{})
		return
	}

	// our proper type-converted expression
	// value(s) shall reside as an any, as
	// stackage.Condition allows this and
	// will make things simpler.
	var ex any

	// perform a bind keyword switch upon
	// a resolution attempt of the value.
	switch key := matchBKW(r.Keyword().String()); key {

	case BindUDN, BindRDN, BindGDN:
		// value is a userdn, groupdn or roledn
		// expressive statement. Possible multi
		// valued expression.
		ex, err = assertBindUGRDN(expr, key)

	case BindIP, BindDNS:
		// value is an IP or FQDN.
		ex, err = assertBindNet(expr, key)

	case BindUAT, BindGAT:
		// value is a userattr or groupattr
		// expressive statement.
		ex, err = assertBindUGAttr(expr, key)

	case BindDoW, BindToD:
		// value is a dayofweek or timeofday
		// expressive statement.
		ex, err = assertBindTimeDay(expr, key)

	case BindAM, BindSSF:
		// value is an authentication method
		// or a security factor expressive
		// statement.
		ex, err = assertBindSec(expr, key)

	default:
		err = badPTBRuleKeywordErr(r, bindRuleID, `BindKeyword`, key)
	}

	if err != nil {
		return
	}

	// If we got something, set it and go.
	r.SetExpression(ex)

	return
}

func assertBindTimeDay(expr parser.RuleExpression, key BindKeyword) (ex any, err error) {
	switch key {
	case BindDoW:
		// value is a dayOfWeek expressive
		// statement.
		ex, err = assertBindDayOfWeek(expr)

	case BindToD:
		// value is a timeOfDay expressive
		// statement.
		ex, err = assertBindTimeOfDay(expr)
	}

	return
}

func assertBindSec(expr parser.RuleExpression, key BindKeyword) (ex any, err error) {
	switch key {
	case BindSSF:
		// value is a security strength factor
		// expressive statement.
		ex, err = assertBindSecurityStrengthFactor(expr)

	case BindAM:
		// value is an authentication method
		// expressive statement.
		ex, err = assertBindAuthMethod(expr)
	}

	return
}

func assertBindUGAttr(expr parser.RuleExpression, key BindKeyword) (ex any, err error) {
	if err = unexpectedBindConditionValueErr(key, 1, expr.Len()); err != nil {
		return
	}

	if hasPfx(expr.Values[0], LocalScheme) {
		// value is an LDAP URI
		ex, err = parseLDAPURI(expr.Values[0], key)

	} else if hasPfx(expr.Values[0], `parent[`) {
		// value is an inheritance attributeBindTypeOrValue
		ex, err = parseInheritance(expr.Values[0])

	} else {
		// value is a standard attributeBindTypeOrValue
		ex, err = parseATBTV(expr.Values[0], key)
	}

	return
}

func assertBindTimeOfDay(expr parser.RuleExpression) (ex TimeOfDay, err error) {
	if err = unexpectedBindConditionValueErr(BindToD, 1, expr.Len()); err != nil {
		return
	}

	// extract clocktime from raw value, remove
	// quotes and any L/T WHSP
	ex = ToD(expr.Values[0])
	err = badClockTimeErr(expr.Values[0], ex.String())
	return
}

func assertBindDayOfWeek(expr parser.RuleExpression) (ex DayOfWeek, err error) {
	if err = unexpectedBindConditionValueErr(BindDoW, 1, expr.Len()); err != nil {
		return
	}

	// extract auth method from raw value, remove
	// quotes and any L/T WHSP and analyze
	ex, err = parseDoW(expr.Values[0])
	return
}

func assertBindAuthMethod(expr parser.RuleExpression) (ex AuthMethod, err error) {
	if err = unexpectedBindConditionValueErr(BindAM, 1, expr.Len()); err != nil {
		return
	}

	// extract auth method from raw value, remove
	// quotes and any L/T WHSP and analyze
	ex = matchAuthMethod(expr.Values[0])
	err = badAMErr(expr.Values[0], ex.String())
	return
}

func assertBindSecurityStrengthFactor(expr parser.RuleExpression) (ex SecurityStrengthFactor, err error) {
	if err = unexpectedBindConditionValueErr(BindSSF, 1, expr.Len()); err != nil {
		return
	}

	// extract factor from raw value, remove
	// quotes and any L/T WHSP
	fac := SSF(expr.Values[0])
	err = badSecurityStrengthFactorErr(expr.Values[0], fac.String())
	return
}

func assertBindNet(expr parser.RuleExpression, key BindKeyword) (ex any, err error) {
	if err = unexpectedBindConditionValueErr(key, 1, expr.Len()); err != nil {
		return
	}

	if key == BindIP {
		// extract IP Address(es) from raw value,
		// remove quotes and any L/T WHSP and then
		// split for iteration.
		raw := split(expr.Values[0], `,`)
		var addr IPAddr
		for ipa := 0; ipa < len(raw); ipa++ {
			addr.Set(raw[ipa])
		}

		ex = addr
		err = badIPErr(len(raw), addr.Len())
		return
	}

	// extract FQDN from raw value, remove
	// quotes and any L/T WHSP.
	fq := DNS(expr.Values[0])
	err = badDNSErr(expr.Values[0], fq.String())
	ex = fq

	return
}

/*
assertBindUGRDN is handler for all possible DN and URI values used within Bind Rule
expressive statements. In particular, this handles `userdn`, `groupdn` and `roledn`
keyword contexts.

An any-enveloped DistinguishedNames instance is returned in the event that the raw value(s)
represent one (1) or more legal LDAP Distinguished Name value.

In the event that a legal LDAP URI is found, it is returned as an instance of (any-enveloped)
LDAPURI.

Quotation schemes are supported seamlessly and either scheme shall be honored per the ANTLR4
parsed content.
*/
func assertBindUGRDN(expr parser.RuleExpression, key BindKeyword) (ex any, err error) {
	// Don't waste time if expression values
	// are nonexistent.
	if expr.Len() == 0 {
		err = noTBRuleExpressionValues(expr, bindRuleID, key)
		return
	}

	// if the value is an LDAP URI (which merely contains
	// a DN, and is not one unto itself), handle the parse
	// here instead of treating it as a DN.
	var value string = expr.Values[0]
	if hasPfx(value, LocalScheme) && contains(value, `?`) {
		ex, err = parseLDAPURI(value, key)
		return
	}

	// create an appropriate container based on the
	// Bind Rule keyword.
	bdn := stackByBDNKeyword(key)

	// Honor the established quotation scheme that
	// was observed during ANTLR4 processing.
	bdn.setQuoteStyle(expr.Style)

	// Assign the raw (DN) values to the
	// return value. If nothing was found,
	// bail out now.
	if bdn.setExpressionValues(key, expr.Values...); bdn.Len() == 0 {
		err = noTBRuleExpressionValues(expr, bindRuleID, key)
		return
	}

	// Envelope our DN stack within an
	// 'any' instance, which is returned.
	ex = bdn

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

	// Len returns the integer length of the receiver.
	// Only meaningful when run on BindRules instances.
	Len() int

	// IsNesting returns a Boolean value indicative of
	// whether the receiver contains a stack as a value.
	// Only meaningful when run on BindRules instances.
	IsNesting() bool

	// Category will report `bind` in all scenarios.
	Category() string

	// Kind will report `stack` for a BindRules instance, or
	// `condition` for a BindRule instance
	Kind() string

	// isBindContextQualifier ensures no greedy interface
	// matching outside of the realm of bind rules. It need
	// not be accessed by users, nor is it run at any time.
	isBindContextQualifier() bool
}

const bindRuleID = `bind`

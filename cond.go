package aci

/*
cond.go contains rule condition types and methods.
*/

import (
	"reflect"
	"github.com/JesseCoretta/go-stackage"
)

/*
Condition is a type alias of stackage.Condition, and provides the "<kw><op><ex>"
expressive syntax representation by itself, or as a slice element of a Rule.

Instances of this type may be created manually using the Cond function, however
it may be easier to utilize the following methods extended via some of the types
defined in this package:

• Eq

• Ne

• Lt

• Le

• Gt

• Ge

Note that occasionally some exceptions apply to the above list of methods. For
instance, non-numerical values are likely ineligible for operator methods such
as Gt, Le, etc.
*/
type Condition stackage.Condition

/*
ComparisonOperator constants defined within the stackage package are aliased
within this package for convenience, without the need for user-invoked stackage
package import.
*/
const (
	Eq stackage.ComparisonOperator = stackage.Eq // 0x1, "Equal To"
	Ne stackage.ComparisonOperator = stackage.Ne // 0x2, "Not Equal to"	// USE WITH CAUTION
	Lt stackage.ComparisonOperator = stackage.Lt // 0x3, "Less Than"
	Le stackage.ComparisonOperator = stackage.Le // 0x4, "Less Than Or Equal"
	Gt stackage.ComparisonOperator = stackage.Gt // 0x5, "Greater Than"
	Ge stackage.ComparisonOperator = stackage.Ge // 0x6, "Greater Than Or Equal"
)

var comparisonOperatorMethMap map[stackage.ComparisonOperator]string

func matchOp(op string) (stackage.Operator, bool) {
	for k, v := range comparisonOperatorMethMap {
		if eq(v,op) || eq(k.String(), op) {
			return k, true
		}
	}

	return stackage.ComparisonOperator(0), false
}

/*
Kind returns `target` or `bind` to identify the kind of *Condition.
*/
func (r Condition) Kind() string {
	return stackage.Condition(r).ID()
}

/*
Cond wraps go-stackage's Cond function. Condition instances are easier to create
using Eq, Ne, Gt, Ge, Lt and Le methods extended by the various types defined in
this package.

This particular function is really only needed by the user for special use-cases
in which Condition instances need to be manually created from scratch (rare).

A condition is represented through the following abstract structure:

	<keyword> <operator> <expression value>

An example:

	userdn = "ldap:///uid=jesse,ou=People,dc=example,dc=com"

As such, three (3) parameters are required:

• The kw input type describes an interface-encompassed Keyword type, or a string

• The ex input type is the expression (assertion) value to be declared within a Condition

• The op input type is the stackage.Operator interface definition, which allows the use of
custom operators introduced by the user, if needed (such as ~= for "approximate", et al).
However, it is unlikely any operators would be needed beyond the aforementioned go-stackage
ComparisonOperator constants.
*/
func Cond(kw, ex, op any) Condition {
	switch tv := op.(type) {

	case string:
		oper, _ := matchOp(tv)
		return Condition(stackage.Cond(kw, oper, ex)).
			NoPadding(!ConditionPadding)

	case stackage.ComparisonOperator:
		return Condition(stackage.Cond(kw, tv, ex)).
			NoPadding(!ConditionPadding)

	case stackage.Operator:
		return Condition(stackage.Cond(kw, tv, ex)).
			NoPadding(!ConditionPadding)
	}

	return Condition{}
}

/*
Valid wraps go-stackage's Condition.Valid method.
*/
func (r Condition) Valid() error {
	return stackage.Condition(r).Valid()
}

/*
Category wraps go-stackage's Condition.Category method.
*/
func (r Condition) Category() string {
	return stackage.Condition(r).Category()
}

/*
setID wraps go-stackage's Condition.SetID method.
*/
func (r Condition) setID(id string) Condition {
	stackage.Condition(r).SetID(id)
	return r
}

func (r Condition) ID() string {
	return stackage.Condition(r).ID()
}

/*
setCategory wraps go-stackage's Condition.SetCategory method.
*/
func (r Condition) setCategory(cat string) Condition {
	stackage.Condition(r).SetCategory(cat)
	return r
}

/*
IsZero wraps go-stackage's Condition.IsZero method.
*/
func (r Condition) IsZero() bool {
	return stackage.Condition(r).IsZero()
}

/*
Encap wraps go-stackage's Condition.Encap method.
*/
func (r Condition) Encap(x ...any) Condition {
	stackage.Condition(r).Encap(x...)
	return r
}

/*
Paren wraps go-stackage's Condition.Paren method.
*/
func (r Condition) Paren(x ...bool) Condition {
	stackage.Condition(r).Paren(x...)
	return r
}

/*
IsParen wraps go-stackage's Condition.IsParen method.
*/
func (r Condition) isParen() bool {
	return stackage.Condition(r).IsParen()
}

/*
NoPadding wraps go-stackage's Condition.NoPadding method.
*/
func (r Condition) NoPadding(x ...bool) Condition {
	stackage.Condition(r).NoPadding(x...)
	return r
}

/*
String wraps go-stackage's Condition.String method.
*/
func (r Condition) String() string {
	return stackage.Condition(r).String()
}

/*
Keyword wraps go-stackage's Condition.Keyword method.
*/
func (r Condition) Keyword() string {
	return stackage.Condition(r).Keyword()
}

/*
Operator wraps go-stackage's Condition.Operator method.
*/
func (r Condition) Operator() stackage.Operator {
	return stackage.Condition(r).Operator()
}

/*
Value wraps go-stackage's Condition.Value method.
*/
func (r Condition) Value() any {
	return stackage.Condition(r).Value()
}

/*
conditionByOperator uses reflect to obtain and execute the
comparison operator (op) Condition-creating method extended
by expression x. An instance of Condition, alongside an error,
are returned.
*/
func conditionByOperator(op, x any) (c Condition, err error) {
	var m, rv reflect.Value

	// See if x is nil
        if rv = reflect.ValueOf(x); rv.IsZero() {
		err = errorf("%T instance is nil, cannot initialize %T",x,c)
		return
        }

	// Perform operator switch under the assumption
	// the operator is a stackage.ComparisonOperator.
	var found bool
	switch top := op.(type) {
	case stackage.ComparisonOperator:

		// See if the operator maps to a
		// known method-named map value.
		for k, v := range comparisonOperatorMethMap {
			if found = k == top; found {
				m = rv.MethodByName(v)
				break
			}
		}

	case string:

		// string could be the literal operator (>),
		// or simply the method name (Gt); try to
		// resolve either.
		for k, v := range comparisonOperatorMethMap {
			if found = (k.String() == top); found {
				m = rv.MethodByName(v)
				break
			}
		}
	}

	if !found {
		err = errorf("Invalid or unknown method '%#v' (%T)", op, op)
		return
	}

	// Call the desired method, or fail with an error.
        meth, ok := m.Interface().(func() Condition)
	if !ok {
		err = errorf("Unable to locate suitable %T creator method from %T using operator '%s'",c,x,op)
		return
        }

	// Execute method and execute its
	// validity checker, returning the
	// error produced, nil or no.
	c = meth()
	err = c.Valid()

        return
}

func init() {
	comparisonOperatorMethMap = map[stackage.ComparisonOperator]string{
		Eq: `Eq`,
		Ne: `Ne`,
		Lt: `Lt`,
		Gt: `Gt`,
		Le: `Le`,
		Ge: `Ge`,
	}
}

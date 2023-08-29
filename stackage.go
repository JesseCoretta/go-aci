package aci

/*
stackage is a bridge to the go-stackage package
*/

import (
	"github.com/JesseCoretta/go-stackage"
)

var (
	stackOr   func(...int) stackage.Stack = stackage.Or
	stackAnd  func(...int) stackage.Stack = stackage.And
	stackNot  func(...int) stackage.Stack = stackage.Not
	stackList func(...int) stackage.Stack = stackage.List
)

/*
ComparisonOperator constants defined within the stackage package are aliased
within this package for convenience, without the need for user-invoked stackage
package import.
*/
const (
	badCop ComparisonOperator = ComparisonOperator(stackage.ComparisonOperator(0x0))

	Eq ComparisonOperator = ComparisonOperator(stackage.Eq) // 0x1, "Equal To"
	Ne ComparisonOperator = ComparisonOperator(stackage.Ne) // 0x2, "Not Equal to"     !! USE WITH CAUTION !!
	Lt ComparisonOperator = ComparisonOperator(stackage.Lt) // 0x3, "Less Than"
	Le ComparisonOperator = ComparisonOperator(stackage.Le) // 0x4, "Less Than Or Equal"
	Gt ComparisonOperator = ComparisonOperator(stackage.Gt) // 0x5, "Greater Than"
	Ge ComparisonOperator = ComparisonOperator(stackage.Ge) // 0x6, "Greater Than Or Equal"
)

type (
	/*
		ComparisonOperator is an alias for stackage.ComparisonOperator. Instances of
		this type are used to contruct TargetRule and BindRule instances, and describe
		the manner in which abstract contexts are to be evaluated.
	*/
	ComparisonOperator stackage.ComparisonOperator

	/*
		BindDistinguishedNames is an alias type for stackage.Stack, and is intended
		to house one (1) or more DistinguishedName instances for the purpose of
		expression within a BindRule instance.
	*/
	BindDistinguishedNames stackage.Stack

	/*
		TargetDistinguishedNames is an alias type for stackage.Stack, and is intended
		to house one (1) or more DistinguishedName instances for the purpose of
		expression within a TargetRule instance.
	*/
	TargetDistinguishedNames stackage.Stack

	/*
		PermissionBindRules is a stackage.Stack type alias used to store one (1)
		or more instances of PermissionBindRule. Instances of this kind are used
		in top-level Instruction (ACI) assembly.
	*/
	PermissionBindRules stackage.Stack

	/*
		ObjectIdentifiers is an alias type for stackage.Stack, and is intended
		to house one (1) or more ObjectIdentifier instances for the purpose of
		expression within a TargetRule instance.
	*/
	ObjectIdentifiers stackage.Stack

	/*
		Instructions is a go-stackage Stack alias type intended to store slices of
		Instruction instances.

		Note that the concept of a "collection" of Instruction instances does not
		come from the ACIv3 syntax per se, and is implemented here merely for the
		user's convenience. Use of this type is not required in any scenario.
	*/
	Instructions stackage.Stack

	/*
		AttributeTypes is an alias type for stackage.Stack, and is intended
		to house one (1) or more AttributeType instances for the purpose of
		expression within a BindRule or TargetRule instance.
	*/
	AttributeTypes stackage.Stack

	/*
		BindRule is a stackage.Condition type alias intended to represent
		a single Bind Rule; that is, one (1) Bind Rule keyword, one (1)
		comparison operator and one (1) or more string values (called an
		'expression').

		For example:

		        ssf >= "128"

		Instances of this type may be assembled manually by users, or may be
		created logically as a result of textual parsing. Users may also want
		to use convenient Eq, Ne, Gt, Ge, Lt, and Le methods extended through
		various types (as permitted) for simplicity.

		Instances of this type shall appear within BindRules instances and may
		or may not be parenthetical.
	*/
	BindRule stackage.Condition

	/*
		BindRules is a stackage.Stack type alias intended to store and express
		one (1) or more Bind Rule statements, with or without nesting and (at
		least usually) bound by Boolean logical WORD operators 'AND', 'OR' and
		'AND NOT'.

		For example:

		        ssf >= "128" AND ip = "192.168.*"

		Instances of this type may be assembled manually by users, or may be
		created logically as a result of textual parsing. There are also some
		convenient operator-specific methods available (And() for 'AND', Or()
		for 'OR' and Not() for 'AND NOT'.
	*/
	BindRules stackage.Stack

	/*
		TargetRule is a stackage.Condition type alias intended to represent
		a single Target Rule; that is, one (1) Target Rule keyword, one (1)
		comparison operator and one (1) or more string values (called an
		'expression').

		For example:

		        ( targetscope = "subordinate" )

		Instances of this type may be assembled manually by users, or may be
		created logically as a result of textual parsing. Users may also want
		to use convenient Eq and Ne methods extended through various types
		(as permitted) for simplicity.

		Instances of this type shall appear within TargetRules instances.

		TargetRule instances are always parenthetical. No parenthetical control
		methods exist for instances of this type.
	*/
	TargetRule stackage.Condition

	/*
		TargetRules is a stackage.Stack type alias intended to store and express
		one (1) or more Target Rule statements.

		For example:

		        ( targetscope = "subordinate" )( targetattr = "cn || sn || givenName || objectClass" )

		Instances of this type may be assembled manually by users, or may be
		created logically as a result of textual parsing. See the T function
		for easily initializing and returning instances of this type.

		Instances of this type will not allow nesting (i.e.: the addition of any
		stackage.Stack type alias instances). Only individual TargetRule instances
		may be pushed into instances of this type.
	*/
	TargetRules stackage.Stack

	/*
		AttributeFilterOperation is a stackage.Stack type alias used to store Target Attribute
		Filters expressions, specifically those used within TargetRule instances bearing the
		`targattrfilters` TargetRule keyword context.

		See also the AttributeFilterOperations type and its methods.
	*/
	AttributeFilterOperation stackage.Stack

	/*
		AttributeFilterOperations is a go-stackage Stack alias type, used for the
		storage of individual AttributeFilterOperation instances.

		Instances of this design are used in TargetRule instances which bear the
		targattrfilters keyword context.
	*/
	AttributeFilterOperations stackage.Stack
)

/*
castAsCondition merely wraps (casts, converts) and returns an
instance of BindRule -OR- TargetRule as a stackage.Condition
instance. This is useful for calling methods that have not been
extended (wrapped) in this package via go-stackage, as it may not
be needed in many cases ...

An instance submitted as x that is neither a BindRule or TargetRule
will result in an empty stackage.Condition return value.

Note this won't alter an existing BindRule or TargetRule instance,
rather a new reference is made through the stackage.Condition type
defined within go-stackage. The BindRule or TargetRule, once it has
been altered to one's satisfaction, can be sent off as intended and
this "Condition Counterpart" can be discarded, or left for GC.
*/
func castAsCondition(x any) (c *stackage.Condition) {
	switch tv := x.(type) {

	// case match is a single BindRule instance
	case BindRule:
		C := stackage.Condition(tv)
		return &C

	// case match is a single TargetRule instance
	case TargetRule:
		C := stackage.Condition(tv)
		return &C
	}

	return nil
}

/*
castAsCop wraps (casts, converts) and instance of ComparisonOperator
into a native stackage.ComparisonOperator instance. This is useful
when aliasing imported types and want to wrap and call the remote
function equivalent instead of writing your own.
*/
func castAsCop(x ComparisonOperator) stackage.ComparisonOperator {
	return stackage.ComparisonOperator(x)
}

/*
isStackageStack merely matches the input type as a stackage.Stack
type and returns the resultant Boolean value.
*/
func isStackageStack(stack any) (is bool) {
	_, is = stack.(stackage.Stack)
	return
}

/*
castAsStack merely wraps (casts, converts) and returns any type
alias of stackage.Stack as a native stackage.Stack.

This is useful for calling methods that have not been extended
(wrapped) in this package via go-stackage, as it might not be
needed in most cases ...

An instance submitted as x that is NOT a type alias of stackage.Stack
will result in an empty stackage.Stack return value.

Note this won't alter an existing values, rather a new reference is
made through the stackage.Condition type defined within go-stackage.
The alias type, once it has been altered to one's satisfaction, can be
sent off as intended and this "Stack Counterpart" can be discarded, or
left for GC.
*/
func castAsStack(u any) (S stackage.Stack, converted bool) {
	switch tv := u.(type) {

	case stackage.Stack:
		converted = true
		S = tv

	case ObjectIdentifiers:
		converted = true
		S = stackage.Stack(tv)

	case BindDistinguishedNames,
		TargetDistinguishedNames:
		S, converted = castDNRules(tv)

	case BindRules, TargetRules,
		PermissionBindRules:
		S, converted = castBTRules(tv)

	case AttributeTypes:
		converted = true
		S = stackage.Stack(tv)

	case AttributeFilterOperation,
		AttributeFilterOperations:
		S, converted = castFilterRules(tv)
	}

	return
}

func castBTRules(x any) (S stackage.Stack, converted bool) {
	switch tv := x.(type) {
	case BindRules:
		S = stackage.Stack(tv)
		converted = true
	case TargetRules:
		S = stackage.Stack(tv)
		converted = true
	case PermissionBindRules:
		S = stackage.Stack(tv)
		converted = true
	}

	return
}

func castDNRules(x any) (S stackage.Stack, converted bool) {
	switch tv := x.(type) {
	case BindDistinguishedNames:
		S = stackage.Stack(tv)
		converted = true
	case TargetDistinguishedNames:
		S = stackage.Stack(tv)
		converted = true
	}

	return
}

func castFilterRules(x any) (S stackage.Stack, converted bool) {
	switch tv := x.(type) {
	case AttributeFilterOperation:
		S = stackage.Stack(tv)
		converted = true
	case AttributeFilterOperations:
		S = stackage.Stack(tv)
		converted = true
	}

	return
}

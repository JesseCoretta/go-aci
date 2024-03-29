package aci

/*
scope.go contains LDAP Search Scope related types, constants and methods.
*/

/*
SearchScope is a type definition used to represent one of the four (4) possible LDAP Search Scope types that are eligible for use within the ACIv3 syntax specification honored by this package.

SearchScope constants are generally used for crafting TargetRule instances that bear the [TargetScope] [TargetKeyword], as well as for crafting fully-qualified LDAP Search URIs.

See the SearchScope constants defined in this package for specific scopes available.
*/
type SearchScope uint8

/*
Scope initializes, sets and returns an instance of SearchScope in one shot. Valid input types are as follows:

  - Standard scope names as string values (e.g.: `base`, `sub`, `subtree` and others)
  - Integer representations of scopes (see the predefined [SearchScope] constants for details)

This function may only be needed in certain situations where a scope needs to be parsed from values with different representations. Usually the predefined [SearchScope] constants are sufficient.
*/
func Scope(x any) (s SearchScope) {
	switch tv := x.(type) {
	case string:
		s = strToScope(tv)
	case int:
		s = intToScope(tv)
	}

	return
}

/*
SearchScope constants define four (4) known LDAP Search Scopes permitted for use per the ACIv3 syntax specification honored by this package.
*/
const (
	noScope     SearchScope = iota // 0x0 <unspecified_scope>
	BaseObject                     // 0x0, `base`
	SingleLevel                    // 0x1, `one` or `onelevel`
	Subtree                        // 0x2, `sub` or `subtree`
	Subordinate                    // 0x3, `subordinate`
)

/*
invalid value constants used as stringer method returns when something goes wrong :/
*/
const (
	badSearchScope = `<invalid_search_scope>`
)

/*
targetScope returns the "more distinguished but lesser used" naming variations for a given search scope. Generally, these are used in [Instruction] instances that support the [TargetScope] [TargetRule] instance.
*/
func (r SearchScope) targetScope() (s string) {
	switch r {
	case BaseObject:
		s = `base`
	case SingleLevel:
		s = `onelevel`
	case Subtree:
		s = `subtree`
	case Subordinate:
		s = `subordinate` // seems to be an OUD thing.
	}

	return
}

/*
standard returns the more common naming variations for a given search scope. Generally, these are used in fully-qualified LDAP Search URL statements.
*/
func (r SearchScope) standard() (s string) {
	s = `base`
	switch r {
	case SingleLevel:
		s = `one`
	case Subtree:
		s = `sub`
	}

	return
}

/*
Eq initializes and returns a new [TargetRule] instance configured to express the evaluation of the receiver value as Equal-To an [TargetScope] [TargetKeyword] context.
*/
func (r SearchScope) Eq() TargetRule {
	if r == noScope {
		return badTargetRule
	}
	return TR(TargetScope, Eq, r)
}

/*
Ne performs no useful task, as negated equality comparison does not apply to TargetRule instances that bear the [TargetScope] [TargetKeyword] context.

This method exists solely to convey this message. When executed, this method returns a bogus [TargetRule] instance.

This method SHALL NOT appear within instances of [TargetRuleMethods] that were crafted through execution of the receiver's TRM method.
*/
func (r SearchScope) Ne() TargetRule { return badTargetRule }

/*
Keyword returns the [TargetKeyword] associated with the receiver instance enveloped as a [Keyword]. In the context of this type instance, the [TargetKeyword] returned is always [TargetScope].
*/
func (r SearchScope) Keyword() Keyword {
	return TargetScope
}

/*
TRM returns an instance of [TargetRuleMethods].

Each of the return instance's key values represent a single [ComparisonOperator] that is allowed for use in the creation of [TargetRule] instances which bear the receiver instance as an expression value. The value for each key is the actual instance method to -- optionally -- use for the creation of the [TargetRule].

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus [TargetRule] instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly populated and prepared for such activity.
*/
func (r SearchScope) TRM() TargetRuleMethods {
	return newTargetRuleMethods(targetRuleFuncMap{
		Eq: r.Eq,
	})
}

/*
String is a stringer method that returns the string representation of the receiver.  In this particular case, the more succinct and standard string variant is returned, e.g.: `one` for [SingleLevel]. This will normally be used within [LDAPURI] instances.

See the [SearchScope] Target method for [TargetRule]-related scope names.
*/
func (r SearchScope) String() string {
	return r.targetScope()
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r SearchScope) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Target is a stringer method that returns the string representation of the receiver.

This method is primarily intended for creation of a new [TargetScope]-style [TargetRule] instance, and is executed automatically during that process.
*/
func (r SearchScope) Target() string {
	return r.targetScope()
}

/*
strToScope returns a SearchScope constant based on the string input. If a match does not occur, BaseObject (default) is returned.
*/
func strToScope(x string) (s SearchScope) {
	s = noScope
	switch lc(x) {
	case `base`, `baseobject`:
		s = BaseObject
	case `one`, `onelevel`:
		s = SingleLevel
	case `sub`, `subtree`:
		s = Subtree
	case `children`, `subordinate`:
		s = Subordinate
	}

	return
}

/*
intToScope returns a SearchScope constant based on the integer input. If a match does not occur, BaseObject (default) is returned.
*/
func intToScope(x int) (s SearchScope) {
	s = BaseObject //default
	switch x {
	case 1:
		s = SingleLevel
	case 2:
		s = Subtree
	case 3:
		s = Subordinate
	}

	return
}

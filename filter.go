package aci

/*
filter.go contains types, functions, methods and constants that pertain to basic
LDAP Search Filter concepts, as well as TargetRule filter-related abstracts that
bear the targattrfilters or targetfilters keyword contexts.
*/

var (
	badSearchFilter SearchFilter // for failed calls that return a SearchFilter only
	badFilter       string       = `<invalid_search_filter>`
)

/*
SearchFilter is a struct type that embeds an LDAP search filter. Instances of this type
may be used in a variety of areas, from LDAPURI composition to targetfilter rules.
*/
type SearchFilter struct {
	*searchFilter
}

/*
searchFilter is a private (pointer!) type embedded within instances of SearchFilter.
*/
type searchFilter struct {
	string
}

/*
IsZero returns a Boolean value indicative of whether the receiver is nil, or unset.
*/
func (r SearchFilter) IsZero() bool {
	if r.searchFilter == nil {
		return true
	}
	return len((*r.searchFilter).string) == 0
}

/*
Valid -- at the moment -- performs a naïve check on the receiver to determine whether
the value is defined.  This method may, in the future, introduce more sophisticated
checks to increase its value, such as counting (unescaped) parenthetical openers and
closers to verify the effective 'balance' of the expression.

For now, this method is little more than an inverse counterpart to IsZero that returns
an instance of error instead of a Boolean.
*/
func (r SearchFilter) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
	}

	//TODO - add filter checks/decompiler? maybe. maybe not.
	return
}

/*
TRM returns an instance of TargetRuleMethods.

Each of the return instance's key values represent a single instance of
the ComparisonOperator type that is allowed for use in the creation of
TargetRule instances which bear the receiver instance as an expression
value. The value for each key is the actual TargetRuleMethod instance for
OPTIONAL use in the creation of a TargetRule instance.

This is merely a convenient alternative to maintaining knowledge of which
ComparisonOperator instances apply to which types. Instances of this type
are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not
yet been initialized, the execution of ANY of the return instance's value
methods will return bogus TargetRule instances. While this is useful in unit
testing, the end user must only execute this method IF and WHEN the receiver
has been properly populated and prepared for such activity.
*/
func (r SearchFilter) TRM() TargetRuleMethods {
	return newTargetRuleMethods(targetRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
Filter initializes (and optionally sets) a new instance of SearchFilter.
Instances of this kind are used in LDAPURIs, as well as certain target
rules.
*/
func Filter(x ...string) (r SearchFilter) {
	r = SearchFilter{newSearchFilter()}
	if len(x) > 0 {
		r.searchFilter.set(x[0])
	}
	return
}

/*
newSearchFilter is a private function called by Filter during an attempt
to create a new instance of SearchFilter, which is returned as a pointer
reference.
*/
func newSearchFilter() (f *searchFilter) {
	f = new(searchFilter)
	return

}

/*
Keyword returns the Keyword associated with the receiver instance. In
the context of this type instance, the Keyword returned is always TargetFilter.
*/
func (r SearchFilter) Keyword() Keyword {
	return TargetFilter
}

/*
String is a stringer method that returns the string representation of
an LDAP Search Filter.
*/
func (r SearchFilter) String() string {
	if r.searchFilter == nil {
		return ``
	}

	return r.searchFilter.string
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r SearchFilter) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Set assigns the provided value as the LDAP Search Filter instance within the
receiver. Note that this should only be done once, as filters cannot easily
built "incrementally" by the user.
*/
func (r *SearchFilter) Set(x string) *SearchFilter {
	if r.searchFilter == nil {
		r.searchFilter = newSearchFilter()
	}

	r.searchFilter.set(x)
	return r
}

/*
set is a private method executed by SearchFilter.Set.
*/
func (r *searchFilter) set(x string) {
	if len(x) == 0 {
		return
	}
	r.string = x
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To a `targetfilter` Target Keyword
context.
*/
func (r SearchFilter) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}
	return TR(TargetFilter, Eq, r)
}

/*
Ne initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Not-Equal-To a `targetfilter` Target Keyword
context.

Negated equality TargetRule instances should be used with caution.
*/
func (r SearchFilter) Ne() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}
	return TR(TargetFilter, Ne, r)
}

/*
AttributeFilter is a struct type that embeds an AttributeTyp and filter-style Rule.
Instances of this type are a component in the creation of Target Rule definitions
based upon the targattrfilters keyword.
*/
type AttributeFilter struct {
	*atf
}

/*
atf is the embedded type (as a pointer!) within instances of AttributeFilter.
*/
type atf struct {
	AttributeType // single LDAP AttributeType
	SearchFilter  // single LDAP Search Filter
}

/*
AttributeOperation defines either an Add Operation or a Delete Operation.

Constants of this type are used in AttributeFilters instances.
*/
type AttributeOperation uint8

/*
AttributeOperation constants are used to initialize and return AttributeFilters
instances based on one (1) of the possible two (2) constants defined below.
*/
const (
	noAOp AttributeOperation = iota
	AddOp                    // add=
	DelOp                    // delete=
)

/*
AF initializes, optionally sets and returns a new instance of AttributeFilter,
which is a critical component of the `targattrfilters` Target Rule.

Input values must be either a SearchFilter or an AttributeType.
*/
func AF(x ...any) AttributeFilter {
	return AttributeFilter{newAttrFilter(x...)}
}

/*
newAttrFilter is a private function called by AF during an attempt to
create a new instance of AttributeFilter.
*/
func newAttrFilter(x ...any) *atf {
	a := new(atf)
	a.set(x...)
	return a
}

/*
Set assigns the provided address component to the receiver and
returns the receiver instance in fluent-form.

Multiple values can be provided in variadic form, or piecemeal.
*/
func (r *AttributeFilter) Set(x ...any) *AttributeFilter {
	if r.IsZero() {
		r.atf = new(atf)
	}

	r.atf.set(x...)
	return r
}

/*
AttributeType returns the underlying instance of AttributeType, or
a bogus AttributeType if unset.
*/
func (r AttributeFilter) AttributeType() AttributeType {
	if r.IsZero() {
		return badAttributeType
	}

	return r.atf.AttributeType
}

/*
SearchFilter returns the underlying instance of SearchFilter, or
a bogus SearchFilter if unset.
*/
func (r AttributeFilter) SearchFilter() SearchFilter {
	if r.IsZero() {
		return badSearchFilter
	}

	return r.atf.SearchFilter
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r AttributeFilter) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
set is a private method called by AttributeFilter.Set.
*/
func (r *atf) set(x ...any) {
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case string:
			if isIdentifier(tv) {
				r.AttributeType = AT(tv)
			} else {
				r.SearchFilter = Filter(tv)
			}
		case AttributeType:
			r.AttributeType = tv
		case SearchFilter:
			r.SearchFilter = tv
		}
	}
}

/*
String is a stringer method that returns the string representation of the
receiver instance.
*/
func (r AttributeFilter) String() string {
	if err := r.Valid(); err != nil {
		return ``
	}

	return sprintf("%s:%s", r.atf.AttributeType, r.atf.SearchFilter)
}

/*
Keyword returns the Keyword associated with the receiver instance. In
the context of this type instance, the Keyword returned is always TargetFilter.
*/
func (r AttributeFilter) Keyword() Keyword {
	return TargetAttrFilters
}

/*
Valid returns an error indicative of whether the receiver is in an aberrant
state.
*/
func (r AttributeFilter) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
		return
	}

	if r.atf.SearchFilter.IsZero() {
		err = illegalSyntaxPerTypeErr(r, TargetAttrFilters)
	} else if r.atf.AttributeType.IsZero() {
		err = illegalSyntaxPerTypeErr(r, TargetAttrFilters)
	}
	return
}

/*
IsZero returns a Boolean value indicative of whether the receiver is nil,
or unset.
*/
func (r AttributeFilter) IsZero() bool {
	if r.atf == nil {
		return true
	}
	return r.atf.SearchFilter.IsZero() &&
		r.atf.AttributeType.IsZero()
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r AttributeOperation) String() string {
	if r == DelOp {
		return `delete`
	}
	return `add`
}

/*
AFOs returns a freshly initialized instance of AttributeFilterOperations, configured
to store one (1) or more AttributeFilterOperation instances for the purpose of crafting
TargetRule instances which bear the `targattrfilters` keyword context.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
AttributeFilterOperation type (or its string equivalent) during initialization. This
is merely a more convenient alternative to separate initialization and push procedures.

Instances of this design are not generally needed elsewhere.

Values are automatically joined using stackage.List() with SetDelimiter for comma
delimitation by default. See SetDelimiter method if semicolon delimitation is
preferred.
*/
func AFOs(x ...any) (f AttributeFilterOperations) {
	// create a native stackage.Stack
	// and configure before typecast.
	_f := stackList().
		SetDelimiter(rune(44)).
		SetID(targetRuleID).
		NoPadding(!StackPadding).
		SetCategory(TargetAttrFilters.String())

	// cast _f as a proper AttributeFilterOperations
	// instance (f). We do it this way to gain
	// access to the method for the *specific
	// instance* being created (f), thus allowing
	// a custom presentation policy to be set.
	f = AttributeFilterOperations(_f)

	// Set custom Presentation/Push policies
	// per go-stackage signatures.
	//_f.SetPresentationPolicy(f.presentationPolicy).
	_f.SetPushPolicy(f.pushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_f.Push(x...)

	return
}

/*
F returns the appropriate instance creator function for crafting individual
AttributeFilterOperation instances for submission to the receiver. This is merely
a convenient alternative to maintaining knowledge as to which function applies
to the current receiver instance.

As there is only one possibility for instances of this design, the AFO function
is returned.
*/
func (r AttributeFilterOperations) F() func(...any) AttributeFilterOperation {
	return AFO
}

/*
Keyword returns the Keyword associated with the receiver instance. In
the context of this type instance, the Keyword returned is always TargetFilter.
*/
func (r AttributeFilterOperations) Keyword() Keyword {
	return TargetAttrFilters
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r AttributeFilterOperations) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Contains returns a Boolean value indicative of whether the type
and its value were located within the receiver.

Valid input types are AttributeFilterOperation or a valid string
equivalent.

Case is significant in the matching process.
*/
func (r AttributeFilterOperations) Contains(x any) bool {
	return r.contains(x)
}

/*
contains is a private method called by the AttributeFilterOperations
Contains method, et al.

Case is significant in the matching process.
*/
func (r AttributeFilterOperations) contains(x any) bool {
	if r.Len() == 0 {
		return false
	}

	var candidate string

	switch tv := x.(type) {
	case string:
		candidate = tv
	case AttributeFilterOperation:
		candidate = tv.String()
	default:
		return false
	}

	for i := 0; i < r.Len(); i++ {
		// case is significant here.
		if r.Index(i).String() == candidate {
			return true
		}
	}

	return false
}

/*
SetDelimiter controls the delimitation scheme employed by the receiver.

Some vendors use semicolon (ASCII #59) for delimitation for expressions
that include values of this kind. This alternative scheme can be set using
the AttributeFilterOperationsSemiDelim integer constant (1).

Other vendors use a comma (ASCII #44) for delimitation of the same form of
expression. This delimitation scheme represents the default (most common)
behavior, but can be set using the AttributeFilterOperationsCommaDelim
integer constant (0), or when run in niladic fashion.
*/
func (r AttributeFilterOperations) SetDelimiter(i ...int) AttributeFilterOperations {
	_r := r.cast()

	var (
		// default delimiter is a comma
		def string = string(rune(44)) // `,`

		// alternative delimiter is a semicolon
		alt string = string(rune(59)) // `;`
	)

	if len(i) == 0 {
		// caller requests the default
		// delimitation scheme (niladic
		// exec).
		_r.SetDelimiter(def)
		return r
	}

	// perform integer switch, looking
	// for a particular constant value
	switch i[0] {
	case AttributeFilterOperationsSemiDelim:
		// Caller requests alternative
		// delimitation scheme.
		_r.SetDelimiter(alt)
	default:
		// caller requests the default
		// delimitation scheme.
		_r.SetDelimiter(def)
	}

	return r
}

/*
Push wraps go-stackage's Stack.Push method. This method shall attempt to
add the provided input values (x) -- which may contain one (1) or more
instances of AttributeFilterOperation or its string equivalent -- to the
receiver instance.
*/
func (r AttributeFilterOperations) Push(x ...any) AttributeFilterOperations {
	if len(x) == 0 {
		return r
	}

	_r := r.cast()
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case string:
			if afo, err := parseAttributeFilterOperation(tv); err == nil {
				if afo.Len() > 0 {
					_r.Push(afo)
				}
			}
		default:
			_r.Push(tv)
		}
	}

	return r
}

/*
Parse is a convenient alternative to building the receiver instance using individual
instances of the needed types. This method does not use go-antlraci.

An error is returned if the parsing attempt fails for some reason. If successful, the
receiver pointer is updated (clobbered) with new information.

Parse will process the input string (raw) and attempt to split the value using a
delimiter integer identifier, if specified. See AttributeFilterOperationsCommaDelim
(default) and AttributeFilterOperationsSemiDelim const definitions for details.
*/
func (r *AttributeFilterOperations) Parse(raw string, delim ...int) (err error) {
	var d int = AttributeFilterOperationsCommaDelim
	if len(delim) > 0 {
		if delim[0] == AttributeFilterOperationsSemiDelim {
			d = delim[0]
		}
	}

	var R AttributeFilterOperations
	if R, err = parseAttributeFilterOperations(raw, d); err != nil {
		return
	}
	*r = R

	return
}

/*
Parse returns an error instance following an attempt to parse input
raw into the receiver instance. A successful parse will clobber (or
obliterate) any contents already present within the receiver.
*/
func (r *AttributeFilterOperation) Parse(raw string) error {
	afo, err := parseAttributeFilterOperation(raw)
	if err == nil {
		*r = afo
	}

	return err
}

/*
Pop wraps go-stackage's Stack.Pop method.
*/
func (r AttributeFilterOperations) Pop() (afo AttributeFilterOperation) {
	slice, _ := r.cast().Pop()
	if assert, ok := slice.(AttributeFilterOperation); ok {
		afo = assert
	}

	return
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r AttributeFilterOperations) Len() int {
	return r.cast().Len()
}

/*
Index wraps go-stackage's Stack.Index method. Note that the
Boolean OK value returned by go-stackage by default will be
shadowed and not obtainable by the caller.
*/
func (r AttributeFilterOperations) Index(idx int) (afo AttributeFilterOperation) {
	slice, _ := r.cast().Index(idx)
	if assert, ok := slice.(AttributeFilterOperation); ok {
		afo = assert
	}

	return
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r AttributeFilterOperations) IsZero() bool {
	return r.cast().IsZero()
}

/*
Valid wraps go-stackage's Stack.Valid method.
*/
func (r AttributeFilterOperations) Valid() error {
	return r.cast().Valid()
}

/*
Kind returns the categorical label assigned to the receiver.
*/
func (r AttributeFilterOperations) Kind() string {
	return TargetAttrFilters.String()
}

/*
String is a stringer method that returns the string representation of
the receiver instance.
*/
func (r AttributeFilterOperations) String() string {
	return r.cast().String()
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To a `targattrfilters` keyword context.
*/
func (r AttributeFilterOperations) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}
	return TR(TargetAttrFilters, Eq, r)
}

/*
Ne performs no useful task, as negated equality comparison does not apply
to TargetRule instances that bear the `targattrfilters` keyword context.

This method exists solely to convey this message and conform to Go's interface
qualifying signature. When executed, this method will return a bogus TargetRule.

Negated equality TargetRule instances should be used with caution.
*/
func (r AttributeFilterOperations) Ne() TargetRule { return badTargetRule }

/*
TRM returns an instance of TargetRuleMethods.

Each of the return instance's key values represent a single instance of
the ComparisonOperator type that is allowed for use in the creation of
TargetRule instances which bear the receiver instance as an expression
value. The value for each key is the actual TargetRuleMethod instance for
OPTIONAL use in the creation of a TargetRule instance.

This is merely a convenient alternative to maintaining knowledge of which
ComparisonOperator instances apply to which types. Instances of this type
are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not
yet been initialized, the execution of ANY of the return instance's value
methods will return bogus TargetRule instances. While this is useful in unit
testing, the end user must only execute this method IF and WHEN the receiver
has been properly populated and prepared for such activity.
*/
func (r AttributeFilterOperations) TRM() TargetRuleMethods {
	return newTargetRuleMethods(targetRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
pushPolicy conforms to go-stackage's PushPolicy closure signature. This method
is used to govern attempts to push instances into a stack, allowing or rejecting
attempts based upon instance type and other conditions. An error is returned to
the caller revealing the outcome of the attempt.
*/
func (r AttributeFilterOperations) pushPolicy(x ...any) (err error) {
	if len(x) == 0 {
		return
	} else if x[0] == nil {
		err = nilInstanceErr(x[0])
		return
	}

	if r.contains(x[0]) {
		err = pushErrorNotUnique(r, x[0], TargetAttrFilters)
		return
	}

	switch tv := x[0].(type) {
	case AttributeFilterOperation:
		// because codecov :/
		xerr := tv.Valid()
		err = pushErrorNilOrZero(r, tv, TargetAttrFilters, xerr)
		if xerr == nil && tv.Len() > 0 {
			err = nil
		}
	default:
		err = pushErrorBadType(r, tv, TargetAttrFilters)
	}

	return
}

/*
pushPolicy conforms to go-stackage's PushPolicy closure signature. This method
is used to govern attempts to push instances into a stack, allowing or rejecting
attempts based upon instance type and other conditions. An error is returned to
the caller revealing the outcome of the attempt.
*/
func (r AttributeFilterOperation) pushPolicy(x ...any) (err error) {
	if len(x) == 0 {
		return
	} else if x[0] == nil {
		err = nilInstanceErr(x[0])
		return
	}

	if r.contains(x[0]) {
		err = pushErrorNotUnique(r, x[0], TargetAttrFilters)
		return
	}

	switch tv := x[0].(type) {
	case string:
		if len(tv) == 0 {
			err = pushErrorNilOrZero(r, tv, TargetAttrFilters)
		}

	case AttributeFilter:
		if err = tv.Valid(); err != nil {
			err = pushErrorNilOrZero(r, tv, TargetAttrFilters, err)
		}
	default:
		err = pushErrorBadType(r, tv, TargetAttrFilters)
	}

	return
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r AttributeFilterOperation) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Push wraps go-stackage's Stack.Push method.
*/
func (r AttributeFilterOperation) Push(x ...any) AttributeFilterOperation {
	if len(x) == 0 {
		return r
	}

	_r := r.cast()
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case string:
			if af, err := parseAttributeFilter(tv); err == nil {
				_r.Push(af)
			}

		// Don't push AO; it was pushed just
		// because its a convenient way to
		// set the operation on the stack
		// receiver instance. But we can't
		// put it in our pushPolicy because
		// we'd have to return a bogus error
		// to prevent a push, which would
		// almost definitely confuse people
		// reading error logs ...
		case AttributeOperation:
			r.setCategory(tv.makeLabel())

		default:
			_r.Push(tv)
		}
	}

	return r
}

/*
Pop wraps go-stackage's Stack.Pop method.
*/
func (r AttributeFilterOperation) Pop() (af AttributeFilter) {
	slice, _ := r.cast().Pop()
	if assert, ok := slice.(AttributeFilter); ok {
		af = assert
	}

	return
}

/*
F returns the appropriate instance creator function for crafting individual
AttributeFilter instances for submission to the receiver. This is merely a
convenient alternative to maintaining knowledge as to which function applies
to the current receiver instance.

As there is only one possibility for instances of this design, the AF function
is returned.
*/
func (r AttributeFilterOperation) F() func(...any) AttributeFilter {
	return AF
}

/*
Keyword returns the Keyword associated with the receiver instance. In
the context of this type instance, the Keyword returned is always TargetFilter.
*/
func (r AttributeFilterOperation) Keyword() Keyword {
	return TargetAttrFilters
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r AttributeFilterOperation) Len() int {
	return r.cast().Len()
}

/*
Index wraps go-stackage's Stack.Index method. Note that the
Boolean OK value returned by go-stackage by default will be
shadowed and not obtainable by the caller.
*/
func (r AttributeFilterOperation) Index(idx int) (af AttributeFilter) {
	slice, _ := r.cast().Index(idx)
	if assert, ok := slice.(AttributeFilter); ok {
		af = assert
	}
	return
}

/*
Contains returns a Boolean value indicative of whether the type
and its value were located within the receiver.

Valid input types are AttributeFilter or a valid string equivalent.

Case is significant in the matching process.
*/
func (r AttributeFilterOperation) Contains(x any) bool {
	return r.contains(x)
}

/*
contains is a private method called by the AttributeFilterOperation
Contains method, et al.

Case is significant in the matching process.
*/
func (r AttributeFilterOperation) contains(x any) bool {
	if r.Len() == 0 {
		return false
	}

	var candidate string

	switch tv := x.(type) {
	case string:
		candidate = tv
	case AttributeFilter:
		candidate = tv.String()
	default:
		return false
	}

	for i := 0; i < r.Len(); i++ {
		// case is significant here.
		if r.Index(i).String() == candidate {
			return true
		}
	}

	return false
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r AttributeFilterOperation) IsZero() bool {
	return r.cast().IsZero()
}

/*
Valid wraps go-stackage's Stack.Valid method.
*/
func (r AttributeFilterOperation) Valid() (err error) {
	return r.cast().Valid()
}

/*
Kind returns the categorical label assigned to the receiver.
*/
func (r AttributeFilterOperation) Kind() string {
	// we'll just cheat and send the string form
	// of the keyword, as we leverage the underlying
	// category method for the operation. See the
	// getCategory method below...
	return TargetAttrFilters.String()
}

func (r AttributeFilterOperation) getCategory() string {
	return r.cast().Category()
}

/*
String is a stringer method that returns the string representation of
the receiver instance.
*/
func (r AttributeFilterOperation) String() (s string) {
	if r.IsZero() {
		return
	}
	_r := r.cast()
	_r.SetPresentationPolicy(nil)
	s = sprintf("%s=%s", r.Operation(), _r.String())
	_r.SetPresentationPolicy(r.presentationPolicy)
	return
}

/*
presentationPolicy -- when set via go-stackage's Stack.SetPresentationPolicy
method -- shall usurp the standard String method behavior exhibited by the
receiver in favor of the provided closure's own Stringer implementation. It
can be necessary to do this at times if go-stackage's basic String method
generates output text in a way other than what is desired.

See go-stackage's PresentationPolicy documentation for details.
*/
func (r AttributeFilterOperation) presentationPolicy(_ ...any) string {
	return r.String()
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To a `targattrfilters` keyword context.
*/
func (r AttributeFilterOperation) Eq() TargetRule {
	if r.IsZero() {
		return badTargetRule
	}
	return TR(TargetAttrFilters, Eq, r.String()) // TODO: revisit this Stringer nonsense
}

/*
Ne performs no useful task, as negated equality comparison does not apply
to TargetRule instances that bear the `targattrfilters` keyword context.

This method exists solely to convey this message and conform to Go's interface
qualifying signature. When executed, this method will return a bogus TargetRule.

Negated equality TargetRule instances should be used with caution.
*/
func (r AttributeFilterOperation) Ne() TargetRule { return badTargetRule }

/*
TRM returns an instance of TargetRuleMethods.

Each of the return instance's key values represent a single instance of
the ComparisonOperator type that is allowed for use in the creation of
TargetRule instances which bear the receiver instance as an expression
value. The value for each key is the actual TargetRuleMethod instance for
OPTIONAL use in the creation of a TargetRule instance.

This is merely a convenient alternative to maintaining knowledge of which
ComparisonOperator instances apply to which types. Instances of this type
are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not
yet been initialized, the execution of ANY of the return instance's value
methods will return bogus TargetRule instances. While this is useful in unit
testing, the end user must only execute this method IF and WHEN the receiver
has been properly populated and prepared for such activity.
*/
func (r AttributeFilterOperation) TRM() TargetRuleMethods {
	return newTargetRuleMethods(targetRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
AFO returns a freshly initialized instance of AttributeFilterOperation, configured
to store one (1) or more AttributeFilter instances for the purpose of crafting
TargetRule instances which bear the `targattrfilters` keyword context. Instances of
this design are not generally needed outside of that context.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
AttributeFilter type (or its string equivalent) during initialization. This is merely
a more convenient alternative to separate init and push procedures.

Multiple values are automatically ANDed using stackage.And() using the symbolic AND
operator (&&).

See also the AttributeFilterOperations type, and its AFOs function, for the multi-valued
incarnation of this type.
*/
func AFO(x ...any) (f AttributeFilterOperation) {
	// create a native stackage.Stack
	// and configure before typecast.
	_f := stackAnd().
		Symbol(`&&`).
		SetID(targetRuleID).
		NoPadding(!StackPadding).
		SetCategory(TargetAttrFilters.String())

	f = AttributeFilterOperation(_f)
	_f.SetPresentationPolicy(f.presentationPolicy).
		SetPushPolicy(f.pushPolicy)
	f.Push(x...)

	return
}

/*
AFO returns an instance of AttributeFilterOperation based upon the input AttributeFilter
instances.

The instance of AttributeFilterOperation contains an ANDed Rule instance using symbols (`&&`).
*/
func (r AttributeOperation) AFO(x ...any) (afo AttributeFilterOperation) {
	afo = AFO()
	afo.setCategory(r.makeLabel())
	afo.Push(x...)

	return
}

/*
makeLabel is a temporary hack to brand a particular AttributeFilterOperation stack as either
Add or Delete in its operational nature.
*/
func (r AttributeOperation) makeLabel() string {
	return sprintf("%s_%s", TargetAttrFilters, r) // TODO: Find an alternative. I really don't like this.
}

/*
setCategory assigns the categorical string label (cat) to the receiver.
*/
func (r AttributeFilterOperation) setCategory(cat string) {
	r.cast().SetCategory(cat)
}

/*
Operation returns AddOp or DelOp as extracted from the receiver's categorical
label. If invalid, an invalid AttributeOperation value is returned.
*/
func (r AttributeFilterOperation) Operation() AttributeOperation {
	switch x := trimPfx(r.getCategory(), TargetAttrFilters.String()+`_`); lc(x) {
	case `add`:
		return AddOp
	case `delete`:
		return DelOp
	}

	return noAOp
}

/*
hasAttributeFilterOperationPrefix returns a Boolean value indicative of
whether the input string value (raw) begins with a known AttributeOperation
prefix.
*/
func hasAttributeFilterOperationPrefix(raw string) bool {
	switch {
	case hasPfx(raw, `add=`):
		return true

	case hasPfx(raw, `delete=`):
		return true
	}

	return false
}

/*
parseAttributeFilterOperations processes the raw input value into an instance of
AttributeFilterOperations, which is returned alongside an error instance.
*/
func parseAttributeFilterOperations(raw string, delim int) (afos AttributeFilterOperations, err error) {
	var char rune = rune(44) // ASCII #44 [comma, default]

	// If delim is anything except one (1)
	// use the default, else use semicolon.
	if delim == 1 {
		char = rune(59) // ASCII #59 [semicolon]
	}

	// Scan the raw input value and count the number of
	// occurrences of an AttributeOperation prefix.
	var opct int
	for _, oper := range []string{
		AddOp.String() + `=`, // add=
		DelOp.String() + `=`, // delete=
	} {
		if ct := ctstr(raw, oper); ct != 0 {
			// save the count in opct
			// through an increment.
			opct += ct
		}
	}

	if opct == 0 {
		err = afoMissingPrefixErr()
		return
	}

	// Split the raw value using the specified
	// char delimiter. Verify the resultant
	// lengths to ensure a split actually did
	// occur.
	var vals []string
	if vals = split(raw, string(char)); opct != len(vals) {
		err = afosNonIdempSplitErr(delim, len(vals), opct, char)
		return
	}

	// initialize a new AttributeFilterOperations stack
	// instance. Instances of AttributeFilterOperation
	// shall be pushed into this.
	afos = AFOs()

	// iterate each of the above split string
	// slices under the assumption that each
	// is an AttributeFilterOperation instance
	//
	// e.g.: add=objectClass:(&(employeeStatus:active)(c=US))
	for i := 0; i < len(vals); i++ {
		var afo AttributeFilterOperation

		value := unquote(condenseWHSP(vals[i]))

		// each of the slices created per the
		// above char split should begin with
		// an AttributeOperator prefix, which
		// will be either `add=` or `delete=`.
		// Bail out if we find otherwise.
		err = afoMissingPrefixErr()
		if hasAttributeFilterOperationPrefix(value) {
			afo, err = parseAttributeFilterOperation(value)
		}
		if err == nil {
			// Push the verified AttributeFilterOperation
			// instance into our AttributeFilterOperations
			// stack instance.
			afos.Push(afo)
			continue
		}
		break
	}

	return
}

/*
parseAttributeFilterOperation parses the string input value (raw) and attempts to
marshal its contents into an instance of AttributeFilterOperation (afo). An error
is returned alongside afo upon completion of the attempt.
*/
func parseAttributeFilterOperation(raw string) (afo AttributeFilterOperation, err error) {
	var (
		val string
		aop AttributeOperation
		seq []string
	)

	if raw = unquote(condenseWHSP(raw)); len(raw) < 5 {
		err = nilInstanceErr(afo)
		return
	}

	if aop, val, err = parseAttrFilterOperPreamble(raw); err == nil {
		afo = aop.AFO()
		cat := sprintf("%s_%s", TargetAttrFilters, aop) // TODO: Find an alternative. I really don't like this.
		afo.setCategory(cat)
		seq = split(trimS(val), `&&`)

		for j := 0; j < len(seq); j++ {
			var af AttributeFilter
			if af, err = parseAttributeFilter(trimS(seq[j])); err == nil {
				afo.Push(af)
				continue
			}
			break
		}
	}

	return
}

/*
Parse parses the string input value (raw) and attempts to marshal its contents into the
receiver instance. An error is returned if the attempt should fail for some reason.
*/
func (r *AttributeFilter) Parse(raw string) (err error) {
	if raw = unquote(condenseWHSP(raw)); len(raw) < 5 {
		err = nilInstanceErr(r)
		return
	}

	var _r AttributeFilter
	if _r, err = parseAttributeFilter(raw); err == nil {
		*r = _r
	}

	return
}

/*
parseAttributeFilterOperation parses the string input value (raw) and attempts to
marshal its contents into an instance of AttributeFilter (af). An error is returned
alongside af upon completion of the attempt.
*/
func parseAttributeFilter(raw string) (af AttributeFilter, err error) {
	idx := idxr(raw, ':')
	if idx == -1 {
		err = afMissingDelimiterErr(af)
		return
	}

	// TODO - validity checks??
	at := AT(raw[:idx]) // cast first portion as attr
	if at.IsZero() {
		err = nilInstanceErr(at)
		return
	}

	f := Filter(raw[idx+1:]) // cast second portion as filter
	af.Set(at, f)            // assign to struct

	return
}

/*
parseAttributeFilterOperPreamble parses the string input value (raw) and attempts to
identify the prefix as a known instance of AttributeOperation. The inferred operation
identifier, which shall be either 'add=' or 'delete=' is returned as value. An error
is returned alongside aop and value upon completion of the attempt.
*/
func parseAttrFilterOperPreamble(raw string) (aop AttributeOperation, value string, err error) {
	switch {

	case hasPfx(raw, `add=`):
		aop = AddOp
		value = raw[4:]

	case hasPfx(raw, `delete=`):
		aop = DelOp
		value = raw[7:]

	default:
		err = aoBadPrefixErr()
	}

	return
}

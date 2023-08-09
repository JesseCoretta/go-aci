package aci

/*
ldap.go contains types, functions, methods and constants that pertain to basic
LDAP concepts, such as Search Scopes and DNs.
*/

/*
AttributeBindTypeOrValue contains a statement of the following syntax:

	<AttributeName>#<BindType -OR- AttributeValue>

Instances of this type are used in certain Bind Rules, particularly those that
involve user-attribute or group-attribute keywords.
*/
type AttributeBindTypeOrValue struct {
	BindKeyword
	*atbtv
}

/*
IsZero returns a boolean value indicative of whether the receiver is nil,
or unset.
*/
func (r AttributeBindTypeOrValue) IsZero() bool {
	if r.atbtv == nil {
		return true
	}

	return r.BindKeyword == 0x0
}

/*
atbtv is the embedded (POINTER!) type found within instances of AttributeBindTypeOrValue.

Slices are as follows:
  - 0: <atname> (AttributeType)
  - 1: <atv> (BindType Keyword -OR- AttributeValue)
*/
type atbtv [2]any

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
	string // to be replaced by stack-based filter (Rule)
	//Rule 	// TODO: LDAP search filter stack
	//Keyword?
}

/*
IsZero returns a boolean value indicative of whether the receiver is nil, or unset.
*/
func (r SearchFilter) IsZero() bool {
	return r.searchFilter == nil
}

/*
Filter initializes (and optionally sets) a new instance of SearchFilter.
Instances of this kind are used in LDAPURIs, as well as certain target
rules.
*/
func Filter(x ...any) (r SearchFilter) {
	r = SearchFilter{new(searchFilter)}
	r.searchFilter.set(x...)
	return
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
Set assigns the provided value as the LDAP Search Filter instance within the
receiver. Note that this should only be done once, as filters cannot easily
built "incrementally" by the user.
*/
func (r SearchFilter) Set(x ...any) SearchFilter {
	r.searchFilter.set(x...)
	return r
}

/*
set is a private method executed by SearchFilter.Set.
*/
func (r *searchFilter) set(x ...any) {
	if len(x) == 0 {
		return
	}

	if r == nil {
		r = new(searchFilter)
	}

	switch tv := x[0].(type) {
	case string:
		if len(tv) > 0 {
			r.string = tv
		}
	}
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
	AttributeType // single attributeType
	string        // single filter (TODO: replace with SearchFilter)
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
	AddOp AttributeOperation = iota // add=
	DelOp                           // delete=
)

/*
AF initializes, (partially) sets and returns a new instance of AttributeFilter,
which is a critical component of the `targattrfilters` Target Rule.

The at argument must be an attributeType string name, or an actual (valid)
AttributeType instance.
*/
func AF(at, f any) AttributeFilter {
	af := new(atf)
	// create the symbol-based
	// AND stack to act as the
	// filter "container" ...
	A := AttributeFilter{af}
	if A.setAT(at) && A.setFilter(f) {
		return A
	}

	return AttributeFilter{}
}

/*
Eq initializes and returns a new Condition instance configured to express the
evaluation of the receiver value as Equal-To a `targattrfilters`.
*/
func (r AttributeFilter) Eq() Condition {
	return Cond(TargetAttrFilters, r, Eq).
		Paren().
		Encap(`"`).
		setID(`target`).
		NoPadding(!ConditionPadding).
		setCategory(TargetAttrFilters.String())
}

/*
Ne performs no useful task, as negated equality comparison does not apply
to Condition instances that bear the `targattrfilters` keyword. This method
exists solely to convey this message and will return a bogus Condition when
executed.
*/
func (r AttributeFilter) Ne() Condition { return Condition{} }

/*
String is a stringer method that returns the string representation of the
receiver instance.
*/
func (r AttributeFilter) String() string {
	if err := r.Valid(); err != nil {
		return ``
	}

	return sprintf("%s:%s", r.atf.AttributeType, r.atf.string)
}

/*
Valid returns an error indicative of whether the receiver is in an aberrant
state.
*/
func (r AttributeFilter) Valid() (err error) {
	if r.IsZero() {
		err = errorf("%T instance is nil", r)
	}
	return
}

/*
setAT is a private method that returns a boolean value indicative of whether
the attempt to set the receiver's attributeType value (string or AttributeType)
succeeded.
*/
func (r AttributeFilter) setAT(at any) (ok bool) {
	switch tv := at.(type) {
	case string:
		if len(tv) > 0 {
			r.atf.AttributeType = ATName(tv)
			ok = true
		}
	case AttributeType:
		if !tv.IsZero() {
			r.atf.AttributeType = tv
			ok = true
		}
	}

	return
}

/*
setFilter is a private method that returns a boolean value indicative of whether
the attempt to set the receiver's filter value (string or filter Rule) succeeded.
*/
func (r AttributeFilter) setFilter(f any) (ok bool) {
	switch tv := f.(type) {
	case string:
		if len(tv) > 0 {
			// TODO: LDAP Filter decompiler / stack
			r.atf.string = tv
			ok = true
		}
		// TODO: LDAP Filter decompiler / stack
		//case Rule:
		//if tv.Category() == TargetFilter.String() && !tv.IsZero() {
		//	r.atf.Rule = tv
		//	ok = true
		//}
	}

	return
}

/*
IsZero returns a boolean value indicative of whether the receiver is nil,
or unset.
*/
func (r AttributeFilter) IsZero() bool {
	if r.atf == nil {
		return true
	}
	return len(r.atf.string) == 0 &&
		r.atf.AttributeType.IsZero()
}

/*
Kind returns one (1) of the following string values, indicating the
operational disposition of the receiver:

• `add`

• `delete`

See the AttributeOperation constants for details.
*/
func (r AttributeOperation) Kind() string {
	return r.String()
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
AF returns an instanceof AttributeFilters based upon the input AttributeFilter
instances.

The instance of AttributeFilters contains an ANDed Rule instance using symbols (`&&`)
and bears the categorical string label of `attrfilters`.
*/
func (r AttributeOperation) AF(x ...AttributeFilter) AttributeFilters {
	ppol := func(x any) (err error) {
		switch tv := x.(type) {
		case AttributeFilter:
			if tv.IsZero() {
				err = errorf("%T denied per PushPolicy method; zero length filter", tv)
			}
		default:
			err = errorf("%T denied per PushPolicy method", tv)
		}
		return
	}

	afs := AttributeFilters{new(atfs)}
	afs.atfs.AttributeOperation = r
	afs.atfs.Rule = Rule(stackageAnd().
		Symbol(`&&`).
		NoPadding(!RulePadding).
		SetPushPolicy(ppol)).
		setCategory(TargetAttrFilters.String())
	if len(x) > 0 {
		afs.atfs.set(x...)
	}
	return afs
}

/*
String is a stringer method that returns the string representation of
the receiver instance.
*/
func (r AttributeFilters) String() string {
	if r.atfs.isZero() {
		return ``
	}

	return sprintf("%s=%s",
		r.atfs.AttributeOperation,
		r.atfs.Rule)
}

/*
Eq initializes and returns a new Condition instance configured to express the
evaluation of the receiver value as Equal-To a `targattrfilters`.
*/
func (r AttributeFilters) Eq() Condition {
	return Cond(TargetAttrFilters, r, Eq).
		Paren().
		Encap(`"`).
		setID(`target`).
		NoPadding(!ConditionPadding).
		setCategory(TargetAttrFilters.String())
}

/*
Ne performs no useful task, as negated equality comparison does not apply
to Condition instances that bear the `targattrfilters` keyword. This method
exists solely to convey this message and will return a bogus Condition when
executed.
*/
func (r AttributeFilters) Ne() Condition { return Condition{} }

/*
Push pushes one (1) or more instances of AttributeFilter into the receiver
instance.
*/
func (r AttributeFilters) Push(x ...AttributeFilter) AttributeFilters {
	r.atfs.set(x...)
	return r
}

/*
isZero returns a boolean value indicative of whether the receiver is
nil, or unset.
*/
func (r *atfs) isZero() bool {
	return r == nil
}

/*
set appends slices of AttributeFilter into the receiver.
*/
func (r *atfs) set(x ...AttributeFilter) {
	for i := 0; i < len(x); i++ {
		r.Rule.Push(x[i])
	}
}

/*
AttributeFilters embeds a pointer value that reflects the following `targattrfilters`
expression syntax:

	"ldapOperation=attributeType:filter [&& attributeType:filter ...], ..."
*/
type AttributeFilters struct {
	*atfs
}

/*
opatfs is the embedded pointer type that resides within an instance of AttributeFilters.

The underlying type contains an AttributeOperation, which may be one of AddOp or DelOp,
and a list-based Rule containing a sequence of AttributFilter instances.
*/
type atfs struct {
	AttributeOperation
	Rule // ANDed AttributeFilter instances
}

/*
AttributeType embeds a pointer value that reflects a single AttributeType name
such as `manager` or `cn`.
*/
type AttributeType struct {
	*string
}

/*
AttributeValue embeds a pointer value that reflects an attributeType
assertion value.
*/
type AttributeValue struct {
	*string
}

/*
ATName initializes, sets and returns an ATName instance in one shot. The
input value x shall be a string attributeType name (e.g.: `manager`).
*/
func ATName(x string) (A AttributeType) {
	if len(x) != 0 {
		A = AttributeType{&x}
	}

	return
}

/*
String returns the string representation of the underlying
value within the receiver. The return value shall reflect
an attributeType name, such as `manager` or `cn`.
*/
func (r AttributeType) String() (s string) {
	s = badAT
	if r.string != nil {
		return (*r.string)
	}

	return
}

/*
IsZero returns a boolean value indicative of whether the receiver is nil,
or unset.
*/
func (r AttributeType) IsZero() bool {
	if r.string == nil {
		return true
	}
	return len(*r.string) == 0
}

/*
ATValue initializes, sets and returns an ATValue instance in one shot. The
input value x shall be a known BindType constant, such as USERDN, OR a raw
string attributeType value, such as `uid=bob,ou=People,dc=example,dc=com`.
*/
func ATValue(x string) (A AttributeValue) {
	if len(x) > 0 {
		A = AttributeValue{&x}
	}
	return
}

/*
String returns the string representation of the underlying value within the receiver.
The return value should be either an attributeType assertion value, or one (1) of the
five (5) possible BindType identifiers (e.g.: USERDN).
*/
func (r AttributeValue) String() (s string) {
	s = badAV
	if r.string != nil {
		return (*r.string)
	}

	return
}

/*
UAT (User-Attribute Name + Bind Type -OR- Attribute Value) returns an initialized
instance of AttributeBindTypeOrValue configured for rules that leverage the `userattr`
Bind Rule keyword.
*/
func UAT(x ...any) AttributeBindTypeOrValue {
	return userOrGroupAttr(BindUAT, x...)
}

/*
GAT (Group-Attribute Name + Bind Type -OR- Attribute Value) returns an initialized
instance of AttributeBindTypeOrValue configured for rules that leverage the `groupattr`
Bind Rule keyword.
*/
func GAT(x ...any) AttributeBindTypeOrValue {
	return userOrGroupAttr(BindGAT, x...)
}

/*
userOrGroupAttr is a private package level function called by either the GroupAttr or
UserAttr function. This function is the base initializer for the AttributeBindTypeOrValue
instance returned by said functions.
*/
func userOrGroupAttr(t BindKeyword, x ...any) (A AttributeBindTypeOrValue) {
	A = AttributeBindTypeOrValue{
		t, new(atbtv),
	}

	if len(x) != 0 {
		A.atbtv.set(x...)
	}

	return
}

/*
Set assigns value(s) x to the receiver. The value(s) must be AttributeType
and/or AttributeValue instances, created via the ATName and ATValue functions
respectively.
*/
func (r *AttributeBindTypeOrValue) Set(x ...any) AttributeBindTypeOrValue {
	r.atbtv.set(x...)
	return *r
}

/*
Eq initializes and returns a new Condition instance configured to express the
evaluation of the receiver value as Equal-To a `userattr`, `groupattr` or some
other abstraction.
*/
func (r AttributeBindTypeOrValue) Eq() (c Condition) {
	if r.atbtv.isZero() {
		return
	}

	return Cond(r.BindKeyword, r, Eq).
		Encap(`"`).
		setID(`bind`).
		NoPadding(!ConditionPadding).
		setCategory(r.BindKeyword.String())
}

/*
Ne initializes and returns a new Condition instance configured to express the
evaluation of the receiver value as Not-Equal-To a `userattr`, `groupattr` or
some other abstraction.

Negated equality Condition instances should be used with caution.
*/
func (r AttributeBindTypeOrValue) Ne() (c Condition) {
	if r.atbtv.isZero() {
		return
	}

	return Cond(r.BindKeyword, r, Ne).
		Encap(`"`).
		setID(`bind`).
		NoPadding(!ConditionPadding).
		setCategory(r.BindKeyword.String())
}

/*
isZero returns a boolean value indicative of whether the receiver is nil, or
unset.
*/
func (r *atbtv) isZero() bool {
	if r == nil {
		return true
	}
	return (r[0] == nil && r[1] == nil)
}

/*
String is a stringer method that returns the string representation of the
receiver.
*/
func (r atbtv) String() string {
	if r.isZero() {
		return ``
	}

	// AttributeType will always
	// be used.
	var at AttributeType

	// Only one (1) of the following
	// vars will be used.
	var bt BindType
	var av AttributeValue

	// Assert the attributeType value
	// or bail out.
	at, assert := r[0].(AttributeType)
	if !assert {
		return ``
	}

	// First see if the value is a BindType
	// keyword, as those are few and easily
	// identified.
	if bt, assert = r[1].(BindType); !assert || bt == BindType(0x0) {
		// If not a BindType kw, see if it
		// appears to be an AttributeValue.
		if av, assert = r[1].(AttributeValue); !assert || len(*av.string) == 0 {
			// value is neither an AttributeValue
			// nor BindType kw; bail out.
			return ``
		}

		// AttributeValue wins
		return sprintf("%s#%s", at, av)
	}

	// BindType wins
	return sprintf("%s#%s", at, bt)
}

/*
set assigns one (1) or more values (x) to the receiver. Only
AttributeType, AttributeValue and BindType instances shall be
assigned.
*/
func (r *atbtv) set(x ...any) {
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case AttributeType:
			r[0] = tv
		case AttributeValue, BindType:
			r[1] = tv
		}
	}
}

/*
String is a stringer method that returns the string representation of the
receiver.
*/
func (r AttributeBindTypeOrValue) String() string {
	if r.atbtv == nil {
		return ``
	}
	return r.atbtv.String()
}

/*
parseATBTV parses the input string (x) in an attempt to marshal its contents
into an instance of AttributeBindTypeOrValue (A), which is returned alongside
an error (err).

The optional BindKeyword argument (bkw) allows the BindGAT (groupattr) Bind
Rule keyword to be set, else the default of BindUAT (userattr) will take
precedence.
*/
func parseATBTV(x string, bkw ...any) (A AttributeBindTypeOrValue, err error) {
	// Obtain the index number for ASCII #35 (NUMBER SIGN).
	// If minus one (-1), input value x is totally bogus.
	idx := idxr(x, '#')
	if idx == -1 {
		err = badAttributeBindTypeOrValueErr(x)
		return
	}

	// Set the groupattr keyword if requested, else
	// use the default of userattr.
	kw := assertATBTVBindKeyword(bkw...)

	// If the remaining portion of the value is, in
	// fact, a known BIND TYPE keyword, pack it up
	// and ship it out.
	if bt := matchBT(x[idx+1:]); bt != BindType(0x0) {
		A = userOrGroupAttr(kw, ATName(x[:idx]), bt)
		return
	}

	// Remaining portion of the value would appear
	// to be an attribute value, so pack it up and
	// send it off.
	A = userOrGroupAttr(kw, ATName(x[:idx]), ATValue(x[idx+1:]))
	return
}

func assertATBTVBindKeyword(bkw ...any) (kw BindKeyword) {
        kw = BindUAT
        if len(bkw) == 0 {
		return
	}

        switch tv := bkw[0].(type) {
        case BindKeyword:
                if tv == BindGAT {
                        kw = tv
                }
        case int:
                if tv == 3 {
                        kw = BindGAT
                }
        case string:
                if eq(tv, BindGAT.String()) {
                        kw = BindGAT
                }
        }

	return
}

/*
SearchScope is a type definition used to represent one of the four (4) possible
LDAP Search Scope types that are eligible for use within the ACI syntax specification
honored by this package.

SearchScope constants are generally used for crafting Condition instances that bear
the `targetscope` keyword, as well as for crafting fully-qualified LDAP Search URIs.

See the SearchScope constants defined in this package for specific scopes available.
*/
type SearchScope uint8

/*
SearchScope constants define four (4) known LDAP Search Scopes permitted for use
per the ACI syntax specification honored by this package.
*/
const (
	noScope     SearchScope = iota // 0x0 <unspecified_scope>
	BaseObject                     // 0x0, `base`
	SingleLevel                    // 0x1, `one` or `onelevel`
	Subtree                        // 0x2, `sub` or `subtree`
	Subordinate                    // 0x3, `subordinate`
)

/*
Filter returns an instance of Rule set to contain the string representation of
an LDAP Search Filter.

Rule instances of this design are intended for general use, as well as for the
creation of LDAPURI instances.

See the TFilter function for a means of using a filter exclusively for a Target
Rule that bears the `targetfilter` keyword.

Rule instances of this design have a maximum capacity of one (1) imposed, thus
only one (1) successful Push of an LDAP filter can take place within any single
instance. This will change in the future to leverage a more sophisticated means
of parsing/decompiling LDAP Search Filters.
*/
/*

//// DEPRECATED, RETIRED ////

func Filter() Rule {
	ppol := func(x any) (err error) {
		switch tv := x.(type) {
		case string:
			if len(tv) == 0 {
				err = errorf("%T denied per PushPolicy method; zero length filter", tv)
			}
		default:
			err = errorf("%T denied per PushPolicy method", tv)
		}
		return
	}

	// Temporarily place the filter in a Rule
	// with a maximum capacity of one (1). This
	// will be reworked in the near future.
	return Rule(stackageList(1).
		SetPushPolicy(ppol)).
		NoPadding(!RulePadding).
		setCategory(`filter`)
}
*/

/*
TFilter returns an instance of Rule set to contain the string representation of
an LDAP Search Filter.

Rule instances of this design are primarily intended for use in composing Target
Rules that bear the `targetfilter` keyword.

See the Filter function for a general-use LDAP Filter facility.

See the AF function for creating Target Rules that bear the `targattrfilters`
keyword, as well as the AttributeOperation.AFs method for creating similar Target
Rules, but with ANDed attr:filter combinations with an ADD or DELETE disposition.

Rule instances of this design have a maximum capacity of one (1) imposed, thus
only one (1) successful Push of an LDAP filter can take place within any single
instance. This will change in the future to leverage a more sophisticated means
of parsing/decompiling LDAP Search Filters.
*/
func TFilter() Rule {
	ppol := func(x any) (err error) {
		switch tv := x.(type) {
		case string:
			if len(tv) == 0 {
				err = errorf("%T denied per PushPolicy method; zero length filter", tv)
			}
		default:
			err = errorf("%T denied per PushPolicy method", tv)
		}
		return
	}

	// Temporarily place the filter in a Rule
	// with a maximum capacity of one (1). This
	// will be reworked in the near future.
	return Rule(stackageList(1).
		SetPushPolicy(ppol)).
		setID(`target`).
		setCategory(TargetFilter.String())
}

/*
invalid value constants used as stringer method returns when
something goes wrong :/
*/
const (
	badSearchScope = `<invalid_search_scope>`
	badAT          = `<invalid_attribute_type>`
	badAV          = `<invalid_attribute_value>`
)

/*
targetScope returns the "more distinguished but lesser used"
naming variations for a given search scope. Generally, these
are used in ACIs that support the `targetscope` Target Rule
Condition value.
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
standard returns the more common naming variations
for a given search scope. Generally, these are used
in fully-qualified LDAP Search URL statements.
*/
func (r SearchScope) standard() (s string) {
	switch r {
	case BaseObject:
		s = `base`
	case SingleLevel:
		s = `one`
	case Subtree:
		s = `sub`
	}

	return
}

/*
Eq initializes and returns a new Condition instance configured to express
the evaluation of the receiver value as Equal-To a `targetscope`.
*/
func (r SearchScope) Eq() Condition {
	return Cond(TargetScope, r.Target(), Eq).
		NoPadding(!ConditionPadding).
		Encap(`"`).
		Paren()
}

/*
Ne performs no useful task, as negated equality comparison does not apply
to Condition instances that bear the `targetscope` keyword. This method
exists solely to convey this message and will return a bogus Condition
when executed.
*/
func (r SearchScope) Ne() Condition { return Condition{} }

/*
String is a stringer method that returns the string
representation of the receiver.  In this particular
case, the more succinct and standard string variant
is returned, e.g.: `one` for SingleLevel. This will
normally be used within LDAPURI instances.

See the SearchScope.Target method for Target Rule
related scope names.
*/
func (r SearchScope) String() string {
	return r.standard()
}

/*
Target is a stringer method that returns the string
representation of the receiver.

This method is primarily intended for creation of a
new `targetscope`-style Target Rule Condition, and
is executed automatically during that process.
*/
func (r SearchScope) Target() string {
	return r.targetScope()
}

/*
DistinguishedName describes an LDAP distinguished name. For example:

	ou=People,dc=example,dc=com

Executing the String method exported through instances of this type
shall automatically prepend the DN with the recommended LDAP "local
scheme" prefix of `ldap:///`, and need not be entered manually.
*/
type DistinguishedName struct {
	*distinguishedName
}

/*
distinguishedName is the embedded type (as a pointer!) within instances of
DistinguishedName.
*/
type distinguishedName struct {
	Keyword // `target`, `target_[to|from]` `userdn`, `groupdn` or `roledn`
	*string
}

/*
LDAPURI contains the components of a fully-qualified LDAP Search URI,
including:

• An LDAP Distinguished Name

• An LDAP Search Scope

• An AttributeBindTypeOrValue

• A comma-delimited list of AttributeType names

• An LDAP Search Filter
*/
type LDAPURI struct {
	*ldapURI
}

/*
ldapURI is the private embedded type (as a pointer) within instances of
LDAPURI.
*/
type ldapURI struct {
	dn     DistinguishedName
	scope  SearchScope
	avbt   AttributeBindTypeOrValue
	attrs  Rule         // list with JoinDelim(`,`)
	filter SearchFilter // TODO: LDAP filter stack
}

/*
URI initializes, (optionally) sets and returns a new instance of LDAPURI,
which represents a fully-qualified LDAP Search URI of the following syntax:

	scheme:///<dn>?<at[,...]>?<scope>?<filter>

As a practical example:

	ldap:///ou=People,dc=example,dc=com?sn,cn,givenName?one?(objectClass=*)

URIs of this format are used within UDN (userdn) and GDN (groupdn) Bind Rules.

Additionally, the ACI syntax specification honored by this package allows the
use of an AttributeBindTypeOrValue instance INSTEAD of a comma-delimited list
of attributeTypes, a search scope and a search filter:

	scheme:///<dn>?<atbtv>

As a practical example:

	ldap:///ou=People,dc=example,dc=com?owner#GROUPDN

Be advised that only UAT (userattr) Bind Rules will make use of this particular
URI format. An instance of AttributeBindTypeOrValue submitted as a URI component
bearing the GAT (groupattr) bind keyword will have this keyword IGNORED in favor
of UAT automatically.
*/
func URI(x ...any) LDAPURI {
	return LDAPURI{newLDAPURI(x...)}
}

/*
newLDAPURI is a private function called by URI.
*/
func newLDAPURI(x ...any) (l *ldapURI) {
	l = new(ldapURI)
	l.attrs = Attrs()
	l.set(x...)
	return
}

/*
parseLDAPURI reads input string x and produces an instance of LDAPURI
(L), which is returned alongside an error instance (err).

An optional Bind Keyword may be provided to supplant BindUAT in the
event of an AttributeBindTypeOrValue instance being present. Note
that only BindGAT is supported as an alternative.
*/
func parseLDAPURI(x string, bkw ...BindKeyword) (L LDAPURI, err error) {
	// URI absolutely MUST begin with the local
	// LDAP scheme (e.g.: ldap:///). If it does
	// not, fail immediately.
	if len(x) < 7 {
		err = errorf("Invalid LDAPURI string '%s'; aborting", x)
		return
	}

	// Chop the scheme off the string,
	// since it is no longer needed.
	uri := x[len(LocalScheme):]

	// initialize our embedded uri type
	l := newLDAPURI()

	var A Rule // (comma-delimited) attributeType list

	// grab all fields, whether zero
	// or not, delimited by question
	// mark characters ...
	fs := split(uri, `?`)

	// iterate fields, and assert the
	// possible types along the way.
	for i := 0; i < len(fs); i++ {

		// switch on index number i, as
		// we take special action based
		// on which index is processed.
		switch i {
		case 0:
			// field #0 is ALWAYS a DN and
			// is never zero length ...
			if len(fs[i]) == 0 {
				err = errorf("Invalid %T for %T (zero length); aborting", DistinguishedName{}, L)
				return
			}

			// Submit new value to LDAPURI instance. Note that DN
			// keyword does not matter here; any DN func would do
			// so we'll just use UDN.
			l.set(UDN(fs[i]))

		case 1:
			// Technically, a zero length
			// string value is fine ...
			if len(fs[i]) == 0 {
				continue
			}

			// field #1 is either:
			// 	1. A list of one (1) or more comma-delimited AttributeType instances
			//	... OR ...
			// 	2. A single AttributeBindTypeOrValue instance

			if contains(fs[i], `#`) {
				// Set the groupattr keyword if requested, else
				// use the default of userattr.
				kw := BindUAT
				if len(bkw) > 0 {
					if bkw[0] == BindGAT {
						kw = BindGAT
					}
				}

				var abv AttributeBindTypeOrValue
				if abv, err = parseATBTV(fs[i], kw); err != nil {
					return
				}

				// Submit new value to LDAPURI instance
				l.set(abv)
			} else {
				A = Attrs() // initialize stack for attributeType list

				// Obliterate spaces and split comma-delimited list
				// into discrete attributeType names. Finally, we'll
				// begin iteration ...
				for _, attr := range split(repAll(fs[i], ` `, ``), `,`) {
					A.Push(ATName(attr))
				}

				// Submit new value to LDAPURI instance
				l.set(A)
			}
		case 2:
			// field #2 is the LDAP Search Scope, if defined. Note that
			// while the directory server shall default to a particular
			// scope if not specified, it is not required in the value
			// and, thus, this package shall not impose the default on
			// its own.

			// targetscope value is not appropriate for LDAP URI
			// scope, and because there is no obvious alternative,
			// we won't set anything. Allow anything else.
			if sc := strToScope(fs[i]); len(fs[i]) > 0 && sc != Subordinate {
				// Submit new value to LDAPURI instance
				l.set(sc)
			}

		case 3:
			// field #3 is the LDAP Search Filter, if defined.
			filt := Filter(fs[i])
			if filt.IsZero() {
				continue
			}
			l.set(filt)
		}
	}

	// Envelope ldapURI instance and send it off
	L = LDAPURI{l}

	return
}

/*
Eq initializes and returns a new Condition instance configured to express the
evaluation of the receiver value as Equal-To one (1) of the following keywords:

• `userdn` (Bind Rule)

• `userattr` (Bind Rule)

• `groupdn` (Bind Rule)

The appropriate keyword is automatically imposed based on the following
scenarios:

• If an AttributeBindTypeOrValue (as created by UAT() package-level function)
is set within the receiver, the keyword shall be BindUAT; this is regardless
to the keyword assigned to the DistinguishedName instance.

• If an AttributeBindTypeOrValue is NOT set within the receiver, the keyword
shall fallback to that which is found within the DistinguishedName assigned to
the receiver; this keyword was set by the UDN() or GDN() package-level functions
respectively during DistinguishedName creation.
*/
func (r LDAPURI) Eq() Condition {
	return r.makeCondition()
}

/*
Ne initializes and returns a new Condition instance configured to express the
evaluation of the receiver value as Not-Equal-To one (1) of the following
keywords:

• `userdn` (Bind Rule)

• `userattr` (Bind Rule)

• `groupdn` (Bind Rule)

The appropriate keyword is automatically imposed based on the following
scenarios:

• If an AttributeBindTypeOrValue (as created by UAT() package-level function)
is set within the receiver, the keyword shall be BindUAT; this is regardless
to the keyword assigned to the DistinguishedName instance.

• If an AttributeBindTypeOrValue is NOT set within the receiver, the keyword
shall fallback to that which is found within the DistinguishedName assigned to
the receiver; this keyword was set by the UDN() or GDN() package-level functions
respectively during DistinguishedName creation.

Negated equality Condition instances should be used with caution.
*/
func (r LDAPURI) Ne() Condition {
	return r.makeCondition(true)
}

/*
makeCondition is a private method extended by LDAPURI solely to be executed
by the Eq and Ne methods during Condition assembly.
*/
func (r LDAPURI) makeCondition(negate ...bool) (c Condition) {
	// don't process a bogus receiver instance.
	if err := r.Valid(); err != nil {
		return
	}

	// Use the desired comparison operator,
	// which can be one of Eq (Equal-To), or
	// Ne (Not-Equal-To).
	var negated bool

	// Equal-To is the default for security reasons
	oper := Eq
	if len(negate) > 0 {
		negated = negate[0]
	}

	if negated {
		oper = Ne
	}

	// default is UDN (userdn) Bind Keyword.
	var kw BindKeyword = BindUDN

	// Try to fallback to GDN, if present.
	if r.dn.distinguishedName.Keyword == BindGDN {
		// only accept GDN as alt.
		kw = BindGDN
	}

	// Now examine the AttributeBindTypeOrValue
	// instance. If defined, examine its Bind
	// Keyword. Only UAT will be allowed, else
	// the above will be imposed.
	if !r.avbt.IsZero() {
		if r.avbt.BindKeyword == BindUAT {
			kw = BindUAT
		}
	}

	// Assemble our condition
	c = Cond(kw, r, oper).
		Encap(`"`).
		setID(`bind`).
		NoPadding(!ConditionPadding).
		setCategory(kw.String())

	// Done!
	return
}

/*
IsZero returns a boolean value indicative of whether the receiver
is nil, or unset.
*/
func (r LDAPURI) IsZero() bool {
	return r.ldapURI.isZero()
}

/*
isZero is a private method called by LDAPURI.IsZero.
*/
func (r *ldapURI) isZero() bool {
	return r == nil
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r LDAPURI) String() string {
	return r.ldapURI.string()
}

/*
string is a private method called by LDAPURI.String.
*/
func (r ldapURI) string() string {
	if err := r.valid(); err != nil {
		return badURI
	}

	var param string
	if !r.filter.IsZero() {
		if !r.attrs.IsZero() {
			param = sprintf("?%s", r.attrs)
		} else {
			param = "?"
		}
		param += sprintf("?%s?%s", r.scope, r.filter)
	} else if !r.avbt.IsZero() {
		param = sprintf("?%s", r.avbt)
	} else {
		return sprintf("%s???", r.dn)
	}

	return sprintf("%s%s", r.dn, param)
}

/*
Valid returns an error instance in the event the receiver is in
an aberrant state.
*/
func (r LDAPURI) Valid() error {
	if r.ldapURI == nil {
		return errorf("%T instance is nil", r)
	}
	return r.ldapURI.valid()
}

/*
valid is a private method called by LDAPURI.Valid.

TODO: add more in-depth checks?
*/
func (r ldapURI) valid() error {
	return r.dn.Valid()
}

/*
Set assigns the provided instances to the receiver. The order in which
instances are assigned to the receiver is immaterial. The semantics of
type instance assignment are as follows:

• An instance of DistinguishedName will be set as the URI DN; this is
ALWAYS required. Valid DN creator functions are UDN (userdn), and GDN
(groupdn) for instances of this type.

• An instance of SearchScope shall be set as the URI Search Scope

• An instance of AttributeBindTypeOrValue shall be set as a component
within the receiver. Please see a special remark below related to the
use of instances of this type within LDAP URIs.

• An instance of Rule, IF it bears the categorical label string value
of `attributes`  (as branded using the Attrs package-level function),
shall be set as the URI attribute(s) list.  Note that a complete list
will be appended to the receiver through iteration and is NOT used to
clobber (or overwrite) any preexisting attribute list.

• An instance of []string is regarded the same as an instance of Rule
bearing the categorical label string value of `attributes`. Instances
of this type are allowed for convenience and because they can be read
and preserved unambiguously.  A []string instance should be used even
if only one AttributeType (e.g.: `cn`) is requested.

• An instance of SearchFilter shall be set as the URI Search Filter.
Please see a special remark below related to the use of instances of
this type within LDAP URIs.

At no point will a string primitive be tolerated as an input value
for any reason.

When an AttributeBindTypeOrValue is specified, an LDAP Search Filter
MUST NOT BE SET, as it will supercede the AttributeBindTypeOrValue
instance during string representation.

In short, choose:

• DN and AttributeType(s), Scope, Filter

... OR ...

• DN and AttributeBindTypeOrValue
*/
func (r *LDAPURI) Set(x ...any) *LDAPURI {
	r.ldapURI.set(x...)
	return r
}

/*
set is a private method called by LDAPURI.Set.
*/
func (r *ldapURI) set(x ...any) {
	// Iterate each of the user-specified
	// input values ...
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {

		// Value is an LDAP Distinguished Name
		case DistinguishedName:
			r.dn = tv

		// Value is an LDAP Search Scope
		case SearchScope:
			r.scope = tv

		// Value is an AttributeBindTypeOr Value
		case AttributeBindTypeOrValue:
			r.avbt = tv

		// Value is an LDAP Search Filter
		case SearchFilter:
			r.filter = tv

		// Value must be an LDAP AttributeType,
		// or stack/slice of same.
		default:
			// append using separate function
			r.attrs.assertAppendURIAttributeTypes(tv)
		}
	}

	return
}

/*
assertAppendURIAttributeTypes appends one (1) or more attribute type
instances (src) into a destination Rule (dest). This method was made
solely to keep the ldapURI.set method's cyclomatic factor low.
*/
func (r Rule) assertAppendURIAttributeTypes(src any) {

	// AttributeType and Rule cases both
	// accomplish the same desired result;
	// in the latter case (Rule), we avoid
	// clobbering a list that was already]
	// defined in some way.
	switch tv := src.(type) {

	// source is a single AttributeType
	case AttributeType:
		// Push attributeType instance into
		// our stack.
		r.Push(tv)

	// source is a stack
	case Rule:
		// Inappropriate list category will
		// fail this process.
		if tv.Category() != `attributes` {
			break
		}

		// AttributeType instances were
		// already defined. Iterate each
		// of the slice members within
		// the stack instance provided
		// and (try to) append to the
		// receiver.
		for j := 0; j < tv.Len(); j++ {
			// call stack index as 'a'.
			// We won't need to manually
			// type-assert 'a', as Push
			// will handle that for us
			// because we're using an
			// AttributeType PushPolicy.
			a, _ := tv.Index(j)

			// Push value a (an attributeType)
			// into stack. Uniqueness will be
			// enforced automatically per our
			// AttributeType PushPolicy :)
			r.Push(a)
		}

	// source is a string slice type instance
	case []string:
		// Iterate each of the slice members
		// within the slice instance provided
		// and (try to) append to receiver.
		for j := 0; j < len(tv); j++ {

			// Push value a (an attributeType)
			// into stack. Uniqueness will be
			// enforced automatically per our
			// AttributeType PushPolicy :)
			r.Push(ATName(tv[j]))
		}
	}
}

/*
UDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Bind Rule Condition that bears the
`userdn` keyword.
*/
func UDN(x string) DistinguishedName {
	return DistinguishedName{newDistinguishedName(x, BindUDN)}
}

/*
RDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Bind Rule Condition that bears the
`roledn` keyword.
*/
func RDN(x string) DistinguishedName {
	return DistinguishedName{newDistinguishedName(x, BindRDN)}
}

/*
GDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Bind Rule Condition that bears the
`groupdn` keyword.
*/
func GDN(x string) DistinguishedName {
	return DistinguishedName{newDistinguishedName(x, BindGDN)}
}

/*
TDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Target Rule Condition that bears the
`target` keyword.
*/
func TDN(x string) DistinguishedName {
	return DistinguishedName{newDistinguishedName(x, Target)}
}

/*
TTDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Target Rule Condition that bears the
`target_to` keyword.
*/
func TTDN(x string) DistinguishedName {
	return DistinguishedName{newDistinguishedName(x, TargetTo)}
}

/*
TFDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Target Rule Condition that bears the
`target_from` keyword.
*/
func TFDN(x string) DistinguishedName {
	return DistinguishedName{newDistinguishedName(x, TargetFrom)}
}

/*
newDistinguishedName is a private function that returns a new instance of
*distinguishedName. This function is called by the UDN, GDN, TDN, TTDN and
TFDN functions.
*/
func newDistinguishedName(x string, kw Keyword) (d *distinguishedName) {
	d = new(distinguishedName)
	d.Keyword = kw

	if len(x) != 0 {
		d.string = &x
	}

	return d
}

/*
Eq initializes and returns a new Condition instance configured to express the
evaluation of the receiver value as Equal-To one (1) of the following keywords:

• `userdn` (Bind Rule),

• `roledn` (Bind Rule),

• `groupdn` (Bind Rule)

• `target` (Target Rule)

• `target_to` (Target Rule)

• `target_from` (Target Rule)

The keyword is automatically set during creation of the DistinguishedName
using one of the appropriate package level DN-manufacturing functions (e.g.:
UDN for a `userdn` Bind Rule Condition).

Note that any Condition bearing a Target Rule keyword will automatically be
configured to encapsulate within parenthesis during string representation.
*/
func (r DistinguishedName) Eq() Condition {
	if r.distinguishedName.isZero() {
		return Condition{}
	}

	switch key := r.distinguishedName.Keyword; key {
	case BindRDN, BindGDN:
		return Cond(key, r, Eq).
			Encap(`"`).
			NoPadding(!ConditionPadding).
			setCategory(key.String())

	case Target, TargetTo, TargetFrom:
		return Cond(key, r, Eq).
			Encap(`"`).
			Paren().
			NoPadding(!ConditionPadding).
			setCategory(key.String())
	}

	return Cond(BindUDN, r, Eq).
		Encap(`"`).
		NoPadding(!ConditionPadding).
		setCategory(BindUDN.String())
}

/*
Ne initializes and returns a new Condition instance configured to express the
evaluation of the receiver value as Not-Equal-To one (1) of the following
keywords:

• `userdn` (Bind Rule),

• `roledn` (Bind Rule),

• `groupdn` (Bind Rule)

• `target` (Target Rule)

• `target_to` (Target Rule)

• `target_from` (Target Rule)

The keyword is automatically set during creation of the DistinguishedName
using one of the appropriate package level DN-manufacturing functions (e.g.:
UDN for a `userdn` Bind Rule Condition).

Note that any Condition bearing a Target Rule keyword will automatically be
configured to encapsulate within parenthesis during string representation.

Negated equality Condition instances should be used with caution.
*/
func (r DistinguishedName) Ne() Condition {
	if r.distinguishedName.isZero() {
		return Condition{}
	}

	switch key := r.distinguishedName.Keyword; key {
	case BindRDN, BindGDN:
		return Cond(key, r, Ne).
			Encap(`"`).
			NoPadding(!ConditionPadding).
			setCategory(key.String())

	case Target, TargetTo, TargetFrom:
		return Cond(key, r, Ne).
			Encap(`"`).
			Paren().
			NoPadding(!ConditionPadding).
			setCategory(key.String())
	}

	return Cond(BindUDN, r, Ne).
		Encap(`"`).
		NoPadding(!ConditionPadding).
		setCategory(BindUDN.String())

}

/*
Valid returns an instance of error that reflects whether certain
required elements or value combinations were present and deemed
valid. A non-nil error indicates an undesirable receiver state.
*/
func (r DistinguishedName) Valid() (err error) {
	// TODO - add more intelligence

	if r.IsZero() {
		err = errorf("No distinguished name value found within %T", r)
	} else if len(*r.distinguishedName.string) < 3 {
		err = errorf("Distinguished name value is invalid: %v", (*r.distinguishedName.string))
	}

	return
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r DistinguishedName) String() string {
	if err := r.Valid(); err != nil {
		return ``
	}

	return sprintf("%s%s", LocalScheme, (*r.distinguishedName.string))
}

/*
IsZero returns a boolean value indicative of whether the receiver
is considered nil, or unset.
*/
func (r DistinguishedName) IsZero() bool {
	return r.distinguishedName.isZero()
}

/*
isZero is a private method called by DistinguishedName.IsZero.
*/
func (r *distinguishedName) isZero() bool {
	return r == nil
}

/*
Scope initializes, sets and returns an instance of SearchScope in one shot.
Valid input types are as follows:

• Standard scope names as string values (e.g.: `base`, `sub`, `subtree` and others)

• Integer representations of scopes (see the predefined SearchScope constants for details)

This function may only be needed in certain situations where a scope needs to be
parsed from values with different representations. Usually the predefined SearchScope
constants are sufficient.
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
strToScope returns a SearchScope constant based on the string input.
If a match does not occur, BaseObject (default) is returned.
*/
func strToScope(x string) (s SearchScope) {
	s = BaseObject //default
	switch lc(x) {
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
intToScope returns a SearchScope constant based on the integer input.
If a match does not occur, BaseObject (default) is returned.
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

/*
ldapFilterOperator is a private type for internal use which
conforms to the Operator interface definition within the
go-stackage. ldapFilterOperator includes the standard Eq, Ge
and Le ComparisonOperators, but also lends itself to the more
specialized operators in certain LDAP filters, including but
not limited to extensibleRule+Attributes, etc.
*/
type ldapFilterOperator uint8

/*
Poached go-ldap filter-related constants and global vars for
future use.
*/
const (
	equalityMatch           ldapFilterOperator = ldapFilterOperator(Eq)
	greaterThanOrEqualMatch ldapFilterOperator = ldapFilterOperator(Ge)
	lessThanOrEqualMatch    ldapFilterOperator = ldapFilterOperator(Le)
	extensibleRuleDN        ldapFilterOperator = iota + ldapFilterOperator(Ge)
	extensibleRuleDNOID
	extensibleRuleAttr
	extensibleRuleNoDN
	approximateMatch
)

var filterMatchOps map[ldapFilterOperator]string = map[ldapFilterOperator]string{
	equalityMatch:           Eq.String(),
	greaterThanOrEqualMatch: `>=`,
	lessThanOrEqualMatch:    `<=`,
	extensibleRuleDN:        `:dn:=`,
	extensibleRuleDNOID:     `:dn:`,
	extensibleRuleAttr:      `:=`,
	extensibleRuleNoDN:      `:`,
	approximateMatch:        `~=`,
}

func (r ldapFilterOperator) String() string {
	if val, found := filterMatchOps[r]; found {
		return val
	}
	return ``
}

func (r ldapFilterOperator) Context() string {
	return `filter` // TODO - return the go-ldap map names?
}

/*
TAttrs returns a new instance of Rule with an initialized embedded
stack configured to function as a simple ORed single-level stack
of attributeType names. Instances of this type are used in Target
Rule instances bearing the `targetattr` keyword.

Only valid instances of AttributeType, or non zero strings, will be
considered for push requests.
*/
func TAttrs() Rule {
	// define a push policy that limits slice candidates to
	// valid strings or bonafide AttributeType instances.
	/*
		ppol := func(x any) (err error) {
			switch tv := x.(type) {
			case string:
				if len(tv) == 0 {
					err = errorf("%T denied per PushPolicy method; zero length string", tv)
				}
			case AttributeType:
				if tv.String() == badAT {
					err = errorf("%T denied per PushPolicy method; zero length", tv)
				}
			default:
				err = errorf("%T denied per PushPolicy method", tv)
			}
			return
		}
	*/

	//SetPushPolicy(ppol)).
	return Rule(stackageOr().
		Symbol(`||`)).
		setPushPolicy().
		setID(`target`).
		NoPadding(!RulePadding).
		setCategory(TargetAttr.String())
}

/*
Attrs returns a list-based instance of Rule set to contain a series of AttributeType
instances. Generally Rule instances of this design are intended for use in fully
qualified LDAP URI instances in which one (1) or more AttributeType values are
requested.

Comma-based delimitation is automatically invoked, and uniqueness of attributeType
slice members is maintained for all push attempts.
*/
func Attrs() Rule {
	return Rule(stackageList().
		JoinDelim(`,`)).
		setPushPolicy().
		setID(`attributes`).
		NoPadding(!RulePadding).
		setCategory(`attributes`)
}

/*
Ctrls returns a new instance of Rule with an initialized embedded
stack configured to function as a simple list containing a single
level of LDAP control object identifiers. Generally instances of
this type are meant for use in Target Rule instances bearing the
`targetcontrol` keyword. See also the package level ExtOps func
for the similar `extop` keyword use case.

Only valid instances of ObjectIdentifier, or non zero strings which
translate correctly into ObjectIdentifier instances should be deemed
suitable for push requests.
*/
func Ctrls() Rule {
	// define a push policy that limits slice candidates to
	// valid strings or bonafide ObjectIdentifier instances.
	ppol := func(x any) (err error) {
		switch tv := x.(type) {
		case string:
			if len(tv) == 0 {
				err = errorf("Cannot use zero string for objectIdentifier push")
			}
		case ObjectIdentifier:
			if tv.IsZero() {
				err = errorf("%T is nil", tv)
			}
		default:
			err = errorf("%T type violates PushPolicy", tv)
		}
		return
	}

	return Rule(stackageOr().
		Symbol(`||`).
		SetPushPolicy(ppol)).
		setID(`target`).
		NoPadding(!RulePadding).
		setCategory(TargetCtrl.String())
}

/*
ExtOps returns a new instance of Rule with an initialized embedded
stack configured to function as a simple list containing a single
level of LDAP control object identifiers. Generally instances of
this type are meant for use in Target Rule instances bearing the
`extop` keyword. See also the package level Ctrls func for the
similar `targetcontrol` keyword use case.

Only valid instances of ObjectIdentifier or non zero strings which
translate correctly into ObjectIdentifier instances should be deemed
suitable for push requests.
*/
func ExtOps() Rule {
	ppol := func(x any) (err error) {
		switch tv := x.(type) {
		case string:
			if len(tv) == 0 {
				err = errorf("Cannot use zero string for objectIdentifier push")
			}
		case ObjectIdentifier:
			if tv.IsZero() {
				err = errorf("%T is nil", tv)
			}
		default:
			err = errorf("%T type violates PushPolicy", tv)
		}
		return
	}

	return Rule(stackageOr().
		Symbol(`||`).
		SetPushPolicy(ppol)).
		setID(`target`).
		NoPadding(!RulePadding).
		setCategory(TargetExtOp.String())
}

/*
TDNs returns a new instance of Rule with an initialized embedded
stack configured to function as a simple list containing a single
level of LDAP distinguished namess. Generally instances of this
type are used in Target Rule instances based upon the `target`,
`target_to` and `target_from` keywords.

Only valid instances of DistinguishedName create using the TDN,
TTDN or TFDN package level functions OR non-zero strings, are to
be considered eligible for push requests.
*/
func TDNs() Rule {
	ppol := func(x any) (err error) {
		switch tv := x.(type) {
		case string:
			if len(tv) == 0 {
				err = errorf("%T cannot use zero string for DN push", tv)
			}
		case DistinguishedName:
			if tv.IsZero() {
				err = errorf("%T is nil", tv)
			} else if !strInSlice(tv.distinguishedName.Keyword.String(), []string{
				Target.String(),
				TargetTo.String(),
				TargetFrom.String(),
			}) {
				err = errorf("%T failed during keyword verification (not a DN-based Target)", tv)
			}
		default:
			err = errorf("%T type violates PushPolicy", tv)
		}
		return
	}

	return Rule(stackageOr().
		Symbol(`||`).
		SetPushPolicy(ppol)).
		NoPadding(!RulePadding).
		setCategory(Target.String())
}

/*
UDNs returns a new instance of Rule with an initialized embedded
stack configured to function as a simple list containing a single
level of LDAP distinguished namess. Generally instances of this
type are used in Bind Rule instances bearing the `userdn` keyword.

Only valid instances of DistinguishedName create using the UDN,
OR non-zero strings, are to be considered eligible for push requests.
*/
func UDNs() Rule {
	ppol := func(x any) (err error) {
		switch tv := x.(type) {
		case string:
			if len(tv) == 0 {
				err = errorf("%T cannot use zero string for DN push", tv)
			}
		case DistinguishedName:
			if tv.IsZero() {
				err = errorf("%T is nil", tv)
			} else if tv.distinguishedName.Keyword != BindUDN {
				err = errorf("%T failed during keyword verification (not a UserDN-based Bind Rule)", tv)
			}
		default:
			err = errorf("%T type violates PushPolicy", tv)
		}
		return
	}

	return Rule(stackageOr().
		Symbol(`||`).
		SetPushPolicy(ppol)).
		NoPadding(!RulePadding).
		setCategory(BindUDN.String())
}

/*
RDNs returns a new instance of Rule with an initialized embedded
stack configured to function as a simple list containing a single
level of LDAP distinguished namess. Generally instances of this
type are used in Bind Rule instances bearing the `roledn` keyword.

Only valid instances of DistinguishedName create using the RDN,
OR non-zero strings, are to be considered eligible for push requests.
*/
func RDNs() Rule {
	ppol := func(x any) (err error) {
		switch tv := x.(type) {
		case string:
			if len(tv) == 0 {
				err = errorf("%T cannot use zero string for DN push", tv)
			}
		case DistinguishedName:
			if tv.IsZero() {
				err = errorf("%T is nil", tv)
			} else if tv.distinguishedName.Keyword != BindRDN {
				err = errorf("%T failed during keyword verification (not a RoleDN-based Bind Rule)", tv)
			}
		default:
			err = errorf("%T type violates PushPolicy", tv)
		}
		return
	}

	return Rule(stackageOr().
		Symbol(`||`).
		SetPushPolicy(ppol)).
		NoPadding(!RulePadding).
		setCategory(BindRDN.String())
}

/*
GDNs returns a new instance of Rule with an initialized embedded
stack configured to function as a simple list containing a single
level of LDAP distinguished namess. Generally instances of this
type are used in Bind Rule instances bearing the `groupdn` keyword.

Only valid instances of DistinguishedName create using the GDN,
OR non-zero strings, are to be considered eligible for push requests.
*/
func GDNs() Rule {
	ppol := func(x any) (err error) {
		switch tv := x.(type) {
		case string:
			if len(tv) == 0 {
				err = errorf("%T cannot use zero string for DN push", tv)
			}
		case DistinguishedName:
			if tv.IsZero() {
				err = errorf("%T is nil", tv)
			} else if tv.distinguishedName.Keyword != BindGDN {
				err = errorf("%T failed during keyword verification (not a GroupDN-based Bind Rule)", tv)
			}
		default:
			err = errorf("%T type violates PushPolicy", tv)
		}
		return
	}

	return Rule(stackageOr().
		Symbol(`||`).
		SetPushPolicy(ppol)).
		NoPadding(!RulePadding).
		setCategory(BindGDN.String())
}

/*
attributeValueAssertionOperator is a private type used to allow the
diverse array of LDAP AttributeValueAssertion (AVA) comparison operators
to be used to fashion new Condition instances. This is a requirement of
the go-stackage Operator interface signature and is used only during the
parsing (marshaling) of string-based LDAP Filters (et al) into usable
objects extended by the go-aci package.
*/
type attributeValueAssertionOperator string

/*
avaOperatorContext is a static string value that is used to "label" an
attributeValueAssertionOperator instance.
*/
const avaOperatorContext = `ava_op`

/*
String is a stringer method that returns the string representation
of the attributeValueAssertion comparison operator. It exists solely
to satisfy go-stackage's Operator interface signature requirements.
*/
func (r attributeValueAssertionOperator) String() string {
	return string(r)
}

/*
Context is a stringer method that returns the string representation
of the attributeValueAssertion comparison operator's context, which
should always be `ava_op`. This method exists solely to satisfy the
go-stackage Operator interface signature requirements.
*/
func (r attributeValueAssertionOperator) Context() string {
	return avaOperatorContext
}

/*
Commonly-used distinguished name expression preambles.
*/
var (
	// LocalScheme is the localhost-implicit scheme prefix for a DN. For security
	// reasons, the LDAP scheme should never be non-local within the context of an
	// ACI (even if proxy operations are involved). For instance:
	//
	//   ldap://ldap.example.com/ou=People,dc=example,dc=com??one?(objectClass=*) // BAD
	//
	//   ldap:///ou=People,dc=example,dc=com??one?(objectClass=*) // GOOD
	//
	// This constant is automatically used in any request related to the string
	// representation of DistinguishedName instances. It is exported and visible
	// to users for reference purposes only, and generally need not be accessed
	// directly.
	LocalScheme = `ldap:///`

	// AllDN is the abstraction of all known user DNs; this does not imply ANONYMOUS DNs
	AllDN DistinguishedName

	// AnyDN is the abstraction of all user DNs, known or anonymous
	AnyDN DistinguishedName

	// SelfDN is the abstraction of a user's own DN
	SelfDN DistinguishedName

	// ParentDN is the abstraction of a user's superior DN
	ParentDN DistinguishedName
)

const (
	// badURI is returned during a failed attempt at
	// string representation of an LDAPURI instance.
	badURI = `<invalid_ldap_uri>`
)

func badAttributeBindTypeOrValueErr(x string) error {
	return errorf("Invalid AttributeBindTyoeOrValue instance: must conform to '<at>#<bt_or_av>', got '%s'", x)
}

/*
init will initialize any global vars residing in this file.
*/
func init() {
	AllDN = UDN(`all`)       // ldap:///all
	AnyDN = UDN(`anyone`)    // ldap:///anyone
	SelfDN = UDN(`self`)     // ldap:///self
	ParentDN = UDN(`parent`) // ldap:///parent
}

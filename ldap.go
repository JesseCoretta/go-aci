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
atbtv is an "Any Triple" type of the following structure:

	0: <dn> <DistinguishedName>
	1: <atname> (AttributeType)
	2: <atv> (AttributeValue OR BindType Keyword constant)

Not all slices shall be populated at all times. Certain cases only call for one
or two particular slice values per use case.

Instances of this type are embedded within AttributeBindTypeOrValue instances.
*/
type atbtv [3]any

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
	Rule          // single filter
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

	return sprintf("%s:%s", r.atf.AttributeType, r.atf.Rule)
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
			r.atf.Rule = Filter().Push(tv)
			ok = true
		}
	case Rule:
		if tv.Category() == TargetFilter.String() && !tv.IsZero() {
			r.atf.Rule = tv
			ok = true
		}
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
	return r.atf.Rule.IsZero() &&
		r.atf.AttributeType.string == nil
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
assertion value, OR a BindType constant such as USERDN.
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
func ATValue(x any) (A AttributeValue) {
	switch tv := x.(type) {
	case string:
		A = AttributeValue{&tv}
	case BindType:
		if s := tv.String(); s != badBT {
			A = AttributeValue{&s}
		}
	}

	return
}

/*
Kind returns the string identifier for the kind of underlying
value within the receiver. The return value shall be one of:

• `bind_type`

• `attribute_value`
*/
func (r AttributeValue) Kind() (k string) {
	k = `attribute_value`

	for _, v := range btMap {
		if *r.string == v {
			k = `bind_type`
			break
		}
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

	return Cond(r.BindKeyword, r, Eq).Encap(`"`).setID(`bind`).setCategory(r.BindKeyword.String())
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

	return Cond(r.BindKeyword, r, Ne).Encap(`"`).setID(`bind`).setCategory(r.BindKeyword.String())
}

/*
isZero returns a boolean value indicative of whether the receiver is nil, or
unset.
*/
func (r *atbtv) isZero() bool {
	if r == nil {
		return true
	}
	return (r[0] == nil && r[1] == nil && r[2] == nil)
}

/*
String is a stringer method that returns the string representation of the
receiver.
*/
func (r atbtv) String() string {
	if r.isZero() {
		return ``
	}
	var val []string

	for i := 0; i < len(r); i++ {
		switch tv := r[i].(type) {
		case DistinguishedName:
			val = append(val, tv.String())
		case AttributeType:
			val = append(val, tv.String())
		case AttributeValue:
			val = append(val, sprintf("#%s", tv))
		case BindType:
			val = append(val, sprintf("#%s", tv))
		}
	}

	return join(val, ``)
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
			r[1] = tv
		case AttributeValue, BindType:
			r[2] = tv
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
	BaseObject  SearchScope = iota // 0x0, `base`
	SingleLevel                    // 0x1, `one` or `onelevel`
	Subtree                        // 0x2, `sub` or `subtree`
	Subordinate                    // 0x3, `subordinate` or `children`
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
	return Rule(stackageList(1).SetPushPolicy(ppol)).setCategory(`filter`)
}

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
	return Rule(stackageList(1).SetPushPolicy(ppol)).setID(`target`).setCategory(TargetFilter.String())
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
	s = `base` // default
	switch r {
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
	s = `base` // default
	switch r {
	case SingleLevel:
		s = `one`
	case Subtree:
		s = `sub`
	case Subordinate:
		s = `children` // almost never used.
	}

	return
}

/*
Eq initializes and returns a new Condition instance configured to express
the evaluation of the receiver value as Equal-To a `targetscope`.
*/
func (r SearchScope) Eq() Condition {
	return Cond(TargetScope, r.Target(), Eq).Encap(`"`).Paren()
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
is returned, e.g.: `one` for SingleLevel.

See also the Formal method for the receiver type.
*/
func (r SearchScope) String() string {
	return r.standard()
}

/*
Target is a stringer method that returns the string
representation of the receiver. Unlike the standard
String method for instances of this type, this will
return the more distinguished string name that goes
with the receiver (e.g.: `onelevel` for SingleLevel
etc).

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
	attrs  Rule // list with JoinDelim(`,`)
	filter Rule
}

/*
URI initializes, (optionally) sets and returns a new instance of LDAPURI,
which represents a fully-qualified LDAP Search URI of the following syntax:

	scheme:///<dn>?<at[,...]>?<scope>?<filter>

As a practical example:

	ldap:///ou=People,dc=example,dc=com?sn,cn,givenName?one?(objectClass=*)

Additionally, the ACI syntax specification honored by this package allows
the use of an AttributeBindTypeOrValue instance instead of a scope and
filter:

	scheme:///<dn>?<atbtv>

As a practical example:

	ldap:///ou=People,dc=example,dc=com?owner#GROUPDN

Generally, the latter case applies to `userattr` and/or `groupattr` Bind
Rules involving static groups, but may have applications elsewhere.
*/
func URI(x ...any) LDAPURI {
	return LDAPURI{newLDAPURI(x...)}
}

/*
newLDAPURI is a private function called by URI.
*/
func newLDAPURI(x ...any) (l *ldapURI) {
	l = new(ldapURI)
	l.set(x...)
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
	return r.ldapURI.valid()
}

/*
valid is a private method called by LDAPURI.Valid.

TODO: add more in-depth checks?
*/
func (r ldapURI) valid() (err error) {
	if r.isZero() {
		err = errorf("%T instance is nil", r)
		return
	}

	// Make sure the DN is sane.
	err = r.dn.Valid()
	return

}

/*
Set assigns the provided type instances to the receiver. The semantics
of type instance assignment are as follows:

• An instance of DistinguishedName shall be set as the URI DN; this is
always required

• An instance of SearchScope shall be set as the URI Search Scope

• An instance of AttributeBindTypeOrValue will be set where a Search
Filter would normally appear in an LDAP Search URI

• An instance of Rule, if it bears the categorical label string value
of `attributes`, shall be set as the URI attribute(s) list

• An instance of Rule, if it bears the categorical label string value
of `filter`, shall be set as the URI Search Filter

In the case of both AttributeBindTypeOrValue and (filter) Rule instances
being set, the filter will take precedence. Only one or the other should
be provided for any single instance of LDAPURI.

If neither a filter-based Rule nor an AttributeBindTypeOrValue are set, the
string representation process will not automatically supply the default LDAP
Search Filter of `objectClass=*`; a filter must be set explicitly in order to
appear during said stringification process.
*/
func (r *LDAPURI) Set(x ...any) *LDAPURI {
	r.ldapURI.set(x...)
	return r
}

/*
set is a private method called by LDAPURI.Set.
*/
func (r *ldapURI) set(x ...any) {
	if r.isZero() {
		r = new(ldapURI)
	}

	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case DistinguishedName:
			r.dn = tv
		case SearchScope:
			r.scope = tv
		case AttributeBindTypeOrValue:
			r.avbt = tv
		case Rule:
			switch c := tv.Category(); c {
			case `attributes`:
				r.attrs = tv
			case `filter`:
				r.filter = tv
			}
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
		return Cond(key, r, Eq).Encap(`"`).setCategory(key.String())
	case Target, TargetTo, TargetFrom:
		return Cond(key, r, Eq).Encap(`"`).Paren().setCategory(key.String())
	}

	return Cond(BindUDN, r, Eq).Encap(`"`).setCategory(BindUDN.String())
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
		return Cond(key, r, Ne).Encap(`"`).setCategory(key.String())
	case Target, TargetTo, TargetFrom:
		return Cond(key, r, Ne).Encap(`"`).Paren().setCategory(key.String())
	}

	return Cond(BindUDN, r, Ne).Encap(`"`).setCategory(BindUDN.String())
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

	return Rule(stackageOr().
		Symbol(`||`).
		SetPushPolicy(ppol)).
		setID(`target`).
		setCategory(TargetAttr.String())
}

/*
Attrs returns a list-based instance of Rule set to contain a series of AttributeType
instances. Generally Rule instances of this design are intended for use in fully
qualified LDAP URI instances in which one (1) or more AttributeType values are
requested.

Comma-based delimitation is automatically invoked.
*/
func Attrs() Rule {
	ppol := func(x any) (err error) {
		switch tv := x.(type) {
		case string:
			if len(tv) == 0 {
				err = errorf("%T denied per PushPolicy method; zero length string", tv)
			}
		case AttributeType:
			if tv.IsZero() {
				err = errorf("%T denied per PushPolicy method; zero length", tv)
			}
		default:
			err = errorf("%T denied per PushPolicy method", tv)
		}
		return
	}

	return Rule(stackageList().
		SetPushPolicy(ppol).
		JoinDelim(`,`)).
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

/*
init will initialize any global vars residing in this file.
*/
func init() {
	AllDN = UDN(`all`)       // ldap:///all
	AnyDN = UDN(`anyone`)    // ldap:///anyone
	SelfDN = UDN(`self`)     // ldap:///self
	ParentDN = UDN(`parent`) // ldap:///parent
}

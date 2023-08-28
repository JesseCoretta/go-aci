package aci

/*
uri.go contains LDAP Search URI types and methods.
*/

const (
	// badURI is returned during a failed attempt at
	// string representation of an LDAPURI instance.
	badURI = `<invalid_ldap_uri>`
)

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
	dn     BindDistinguishedName
	scope  SearchScope
	avbt   AttributeBindTypeOrValue
	attrs  AttributeTypes
	filter SearchFilter
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
	l.attrs = UAs()
	l.set(x...)
	return
}

/*
Parse is a convenient alternative to building the receiver instance using individual
instances of the needed types. This method does not use go-antlraci.

An error is returned if the parsing attempt fails for some reason. If successful, the
receiver pointer is updated (clobbered) with new information.
*/
func (r *LDAPURI) Parse(raw string) (err error) {
	var L LDAPURI
	if L, err = parseLDAPURI(raw); err != nil {
		return
	}
	*r = L

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

	// Chop the scheme off the string, since
	// it is no longer needed.
	uri := x[len(LocalScheme):]

	// initialize our embedded uri type
	l := newLDAPURI()

	// iterate each value produced through split
	// on question mark and massage values into
	// LDAP URI appropriate component values ...
	if err = l.assertURIComponents(split(uri, `?`), bkw...); err != nil {
		return
	}

	// Envelope ldapURI instance and send it off
	L = LDAPURI{l}

	return
}

/*
assertURIComponents is intended to reduce cyclomatic complexity factors associated with
the parseLDAPURI function, which is the sole caller of this function.
*/
func (r *ldapURI) assertURIComponents(vals []string, kw ...BindKeyword) (err error) {

	// iterate fields, and assert the
	// possible types along the way.
	for i := 0; i < len(vals); i++ {

		// switch on index number i, as
		// we take special action based
		// on which index is processed.
		switch i {

		// case match is LDAP Distinguished Name
		case 0:
			// Submit new value to LDAPURI instance. Note that DN
			// keyword does not matter here; any DN func would do
			// so we'll just use UDN.
			r.set(UDN(vals[i]))

		// case match is ATBTV -OR- Search Attribute(s)
		case 1:
			if err = r.uriAssertATB(vals[i], kw...); err != nil {
				return
			}

		// case match is LDAP Search Scope
		case 2:
			// field #2 is the LDAP Search Scope, if defined. Note that
			// while the directory server shall default to a particular
			// scope if not specified, it is not required in the value
			// and, thus, this package shall not impose the default on
			// its own.

			// targetscope value is not appropriate for LDAP URI
			// scope, and because there is no obvious alternative,
			// we won't set anything. Allow anything else.
			if sc := strToScope(vals[i]); len(vals[i]) > 0 {
				// Submit new value to LDAPURI instance
				r.set(sc)
			}

		// case match is LDAP Search Filter
		case 3:
			filt := Filter(vals[i])
			if filt.IsZero() {
				continue
			}
			r.set(filt)
		}
	}

	return
}

/*
uriAssertATB shall analyze the input string value (raw) and will do one (1)
of the following:

• Parse the value under the assumption it represents an AttributeBindTypeOrvalue instance
using the provided BindKeyword (kw), OR ...

• Parse the value under the assumption it represents a sequence of one or more LDAP AttributeType
values, delimited using a comma (ASCII #44) as needed

Should any errors be encountered, a non-nil error instance (err) is returned.
*/
func (r *ldapURI) uriAssertATB(raw string, bkw ...BindKeyword) (err error) {
	// Technically, a zero length
	// string value is fine ...
	if len(raw) == 0 {
		return
	}

	// raw is either:
	//      1. A list of one (1) or more comma-delimited AttributeType instances
	//      ... OR ...
	//      2. A single AttributeBindTypeOrValue instance

	// raw must be an ATBTV
	if contains(raw, `#`) {

		// Set the groupattr keyword if requested, else
		// use the default of userattr.
		kw := BindUAT
		if len(bkw) > 0 {
			if bkw[0] == BindGAT {
				kw = BindGAT
			}
		}

		var abv AttributeBindTypeOrValue
		if abv, err = parseATBTV(raw, kw); err != nil {
			return
		}

		// Submit new value to LDAPURI instance
		r.set(abv)

		// raw must be a list of one (1) or
		// more attributeType names ...
	} else {

		A := UAs() // initialize stack for attributeType list

		// Obliterate spaces and split comma-delimited list
		// into discrete attributeType names. Finally, we'll
		// begin iteration ...
		for _, attr := range split(repAll(raw, ` `, ``), `,`) {
			A.Push(AT(attr))
		}

		// Submit new value(s) to LDAPURI instance
		r.set(A)
	}

	return
}

/*
Eq initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Equal-To one (1) of the following keyword
contexts:

• `userdn` (Bind Rule)

• `userattr` (Bind Rule)

• `groupdn` (Bind Rule)

The appropriate keyword is automatically imposed based on the following
scenarios:

• If an AttributeBindTypeOrValue (as created by UAT() package-level function)
is set within the receiver, the keyword shall be BindUAT; this is regardless
to the keyword assigned to the DistinguishedName instance.

• If an AttributeBindTypeOrValue is NOT set within the receiver, the keyword
shall fallback to that which is found within the BindDistinguishedName assigned to
the receiver; this keyword was set by the UDN() or GDN() package-level functions
respectively during BindDistinguishedName creation.
*/
func (r LDAPURI) Eq() BindRule {
	return r.makeBindRule()
}

/*
Ne initializes and returns a new BindRule instance configured to express the
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
shall fallback to that which is found within the BindDistinguishedName assigned
to the receiver; this keyword was set by the UDN() or GDN() package-level functions
respectively during BindDistinguishedName creation.

Negated equality BindRule instances should be used with caution.
*/
func (r LDAPURI) Ne() BindRule {
	return r.makeBindRule(true)
}

/*
makeBindRule is a private method extended by LDAPURI solely to be executed
by the Eq and Ne methods during BindRule assembly.
*/
func (r LDAPURI) makeBindRule(negate ...bool) BindRule {
	// don't process a bogus receiver instance.
	if err := r.Valid(); err != nil {
		return badBindRule
	}

	var b BindRule

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

	// assemble our BindRule instance ...
	b.SetKeyword(kw)
	b.SetOperator(oper)
	b.SetExpression(r)

	// temporarily cast as a stackage.Condition
	// so we can apply some additional changes
	// using methods we didn't wrap because it
	// wouldn't be necessary otherwise.
	_b := castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(kw.String())

	b = BindRule(*_b)
	return b
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
	if r.IsZero() {
		return ``
	}

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
	if r.IsZero() {
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

• An instance of BindDistinguishedName will be set as the URI DN; this
is ALWAYS required. Valid DN creator functions are UDN (userdn), and GDN
(groupdn) for instances of this type.

• An instance of SearchScope shall be set as the URI Search Scope

• An instance of AttributeBindTypeOrValue shall be set as a component
within the receiver. Please see a special remark below related to the
use of instances of this type within LDAP URIs.

• An instance of AttributeTypes (created using the UAs package-level
function) shall be set as the URI attribute(s) list.

• An instance of SearchFilter shall be set as the URI Search Filter.
Please see a special remark below related to the use of instances of
this type within LDAP URIs.

• An instance of string (with no other arguments) shall result in an
LDAP URI parse operation that, if successful, shall overwrite the
receiver in its entirety. This should not be combined with other types
of input values as the results will not be compounded.

When an AttributeBindTypeOrValue is specified, an LDAP Search Filter
MUST NOT BE SET, as it will supercede the AttributeBindTypeOrValue
instance during string representation.

In short, choose:

• DN and AttributeType(s), Scope, Filter

... OR ...

• DN and AttributeBindTypeOrValue
*/
func (r *LDAPURI) Set(x ...any) *LDAPURI {
	if r == nil {
		*r = URI()
	}
	r.ldapURI.set(x...)
	return r
}

/*
set is a private method called by LDAPURI.Set.
*/
func (r *ldapURI) set(x ...any) {
	/*
		if r == nil {
			R := newLDAPURI()
			r = R
		}
	*/

	// Iterate each of the user-specified
	// input values ...
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case string:
			// Value is a complete LDAPURI in string
			// representation. If we succeed in parsing
			// the value, return (do not continue).
			if L, err := parseLDAPURI(tv); err == nil {
				*r = (*L.ldapURI)
				return
			}

		case BindDistinguishedName:
			// Value is an LDAP Distinguished Name
			r.dn = tv

		case SearchScope:
			// Value is an LDAP Search Scope
			//if tv != Subordinate {
			r.scope = tv // check elsewhere
			//}

		case AttributeBindTypeOrValue:
			// Value is an AttributeBindTypeOr Value
			r.avbt = tv

		case SearchFilter:
			// Value is an LDAP Search Filter
			r.filter = Filter(tv.String())

		case AttributeTypes:
			// Value(s) are one or more LDAP
			// Search Attributes
			tv.transfer(r.attrs)
		}
	}

	return
}

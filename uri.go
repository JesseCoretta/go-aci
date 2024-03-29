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
LDAPURI contains the components of a fully-qualified LDAP Search URI which is comprised of one (1) or more of the following:

  - One (1) [BindDistinguishedName] instance
  - One (1) [SearchScope] instance
  - One (1) [AttributeBindTypeOrValue] instance
  - One (1) [AttributeTypes] instance
  - One (1) [SearchFilter] instance
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
URI initializes, (optionally) sets and returns a new instance of [LDAPURI], which represents a fully-qualified LDAP Search URI of the following syntax:

	scheme:///<dn>?<at[,...]>?<scope>?<filter>

As a practical example:

	ldap:///ou=People,dc=example,dc=com?sn,cn,givenName?one?(objectClass=*)

URIs of this format are used within [BindUDN] and [BindGDN] [BindRules].

Additionally, the ACI syntax specification honored by this package allows the use of an [AttributeBindTypeOrValue] instance INSTEAD of a comma-delimited list of attributeTypes, a search scope and a search filter:

	scheme:///<dn>?<atbtv>

As a practical example:

	ldap:///ou=People,dc=example,dc=com?owner#GROUPDN

Be advised that only [BindUAT] (userattr) [BindRules] instances shall make use of this particular URI format. An instance of AttributeBindTypeOrValue submitted as a URI component bearing the [BindGAT] (groupattr) [BindKeyword] will have this keyword IGNORED in favor of [BindUAT] automatically.
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
Len performs no useful task, as the concept of integer length does not apply to the [LDAPURI] type in this context. Execution of this method shall always return zero (0).

This method exists solely to allow the receiver to qualify for the [DistinguishedNameContext] interface signature.
*/
func (r LDAPURI) Len() int {
	return 0
}

/*
assertURIComponents is intended to reduce cyclomatic complexity factors associated with the parseLDAPURI function, which is the sole caller of this function.
*/
func (r *ldapURI) assertURIComponents(vals []string, kw ...BindKeyword) (err error) {

	// iterate fields, and assert the
	// possible types along the way.
	for i := 0; i < len(vals); i++ {

		// switch on index number i, as
		// we take special action based
		// on which index is processed.
		switch i {

		case 0:
			// case match is LDAP Distinguished Name
			// Submit new value to LDAPURI instance. Note that DN
			// keyword does not matter here; any DN func would do
			// so we'll just use UDN.
			r.set(UDN(vals[i]))

		case 1:
			// case match is ATBTV -OR- Search Attribute(s)
			if err = r.uriAssertATB(vals[i], kw...); err != nil {
				return
			}

		case 2:
			// case match is the LDAP Search Scope, if defined. Note that
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

		case 3:
			// case match is LDAP Search Filter
			r.set(Filter(vals[i]))
		}
	}

	return
}

/*
uriAssertATB shall analyze the input string value (raw) and will do one (1) of the following:

• Parse the value under the assumption it represents an AttributeBindTypeOrValue instance using the provided BindKeyword (kw), OR ...
• Parse the value under the assumption it represents a sequence of one or more LDAP AttributeType values, delimited using a comma (ASCII #44) as needed

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
Eq initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Equal-To one (1) of the following [BindKeyword] contexts:

  - [BindUDN]
  - [BindUAT]
  - [BindGAT]

The appropriate [BindKeyword] is automatically imposed based on the following scenarios:

  - If an [AttributeBindTypeOrValue] (as created by [UAT] package-level function) is set within the receiver, the [BindKeyword] shall be [BindUAT]; this is regardless to the keyword assigned to the underying distinguished name.
  - If an [AttributeBindTypeOrValue] is NOT set within the receiver, the [BindKeyword] shall fallback to that which is found within the [BindDistinguishedName] assigned to the receiver; this keyword was set by the [UDN] or [GDN] package-level functions respectively during [BindDistinguishedName] creation.
*/
func (r LDAPURI) Eq() BindRule {
	return r.makeBindRule()
}

/*
Ne initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Not-Equal-To one (1) of the following [BindKeyword] contexts:

  - [BindUDN]
  - [BindUAT]
  - [BindGAT]

The appropriate [BindKeyword] is automatically imposed based on the following scenarios:

  - If an [AttributeBindTypeOrValue] (as created by [UAT] package-level function) is set within the receiver, the [BindKeyword] shall be [BindUAT]; this is regardless to the keyword assigned to the underying distinguished name.
  - If an [AttributeBindTypeOrValue] is NOT set within the receiver, the [BindKeyword] shall fallback to that which is found within the [BindDistinguishedName] assigned to the receiver; this keyword was set by the [UDN] or [GDN] package-level functions respectively during [BindDistinguishedName] creation.

Negated equality [BindRule] instances should be used with caution.
*/
func (r LDAPURI) Ne() BindRule {
	return r.makeBindRule(true)
}

/*
BRM returns an instance of [BindRuleMethods].

Each of the return instance's key values represent a single instance of the [ComparisonOperator] type that is allowed for use in the creation of [BindRule] instances which bear the receiver instance as an expression value. The value for each key is the actual [BindRuleMethod] instance for OPTIONAL use in the creation of a [BindRule] instance.

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus [BindRule] instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly populated and prepared for such activity.
*/
func (r LDAPURI) BRM() BindRuleMethods {
	return newBindRuleMethods(bindRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
makeBindRule is a private method extended by LDAPURI solely to be executed by the Eq and Ne methods during BindRule assembly.
*/
func (r LDAPURI) makeBindRule(negate ...bool) BindRule {
	// don't process a bogus receiver instance.
	if err := r.Valid(); err != nil {
		return badBindRule
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

	// assemble our BindRule instance ...
	return BR(kw, oper, r)
}

/*
IsZero returns a Boolean value indicative of whether the receiver is nil, or unset.
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
String is a stringer method that returns the string representation of the receiver instance.
*/
func (r LDAPURI) String() string {
	if r.IsZero() {
		return ``
	}

	return r.ldapURI.string()
}

/*
Keyword returns the [BindKeyword] associated with the receiver instance enveloped as a [Keyword]. In the context of this type instance, the [BindKeyword] returned will be one of the following:

  - [BindUDN]
  - [BindGDN]
  - [BindRDN]
  - [BindUAT]
  - [BindGAT]

Which [BindKeyword] is actually used is determined by the underlying type instances that were used to assemble the receiver.
*/
func (r LDAPURI) Keyword() (k Keyword) {
	if err := r.Valid(); err != nil {
		return
	}

	if !r.ldapURI.avbt.isZero() {
		switch kw := r.ldapURI.avbt.BindKeyword; kw {
		case BindUAT, BindGAT:
			k = kw
		}
	} else {
		switch kw := r.ldapURI.dn.Keyword(); kw {
		case BindGDN, BindRDN, BindUDN:
			k = kw
		}
	}

	return k
}

/*
Kind returns the string form of the receiver's [Keyword], if the instance is valid.
*/
func (r LDAPURI) Kind() (k string) {
	kw := r.Keyword()
	if kw != nil {
		k = kw.String()
	}
	return
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
		param = "?"
		if !r.attrs.IsZero() {
			param = sprintf("?%s", r.attrs)
		}

		// Be sure to call the standard scope here,
		// since this isn't for a targetscope rule.
		param += sprintf("?%s?%s",
			r.scope.standard(), r.filter)

	} else if !r.avbt.IsZero() {
		param = sprintf("?%s", r.avbt)
	} else {
		return sprintf("%s??%s?",
			r.dn, r.scope.standard())
	}

	return sprintf("%s%s", r.dn, param)
}

/*
Valid returns an error instance in the event the receiver is in an aberrant state.
*/
func (r LDAPURI) Valid() error {
	if r.IsZero() {
		return nilInstanceErr(r)
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
Set assigns the provided instances to the receiver. The order in which instances are assigned to the receiver is immaterial. The semantics of type instance assignment are as follows:

  - An instance of [BindDistinguishedName] will be set as the URI DN; this is ALWAYS required. Valid DN creator functions are [UDN] and [GDN] for instances of this type.
  - An instance of [SearchScope] shall be set as the URI SearchScope
  - An instance of [AttributeBindTypeOrValue] shall be set as a component within the receiver. Please see a special remark below related to the use of instances of this type within [LDAPURI] instances.
  - An instance of [AttributeTypes] (created using the UAs package-level function) shall be set as the URI attribute(s) list.
  - An instance of [SearchFilter] shall be set as the URI Search Filter. Please see a special remark below related to the use of instances of this type within [LDAPURI] instances.
  - An instance of string (with no other arguments) shall result in an [LDAPURI] parse operation that, if successful, shall overwrite the receiver in its entirety. This should not be combined with other types of input values as the results will not be compounded.

When an [AttributeBindTypeOrValue] is specified, a [SearchFilter] MUST NOT BE SET, as it will supercede the [AttributeBindTypeOrValue] instance during string representation.

In short, choose:

• [BindDistinguishedName] and [AttributeTypes], [SearchScope], [SearchFilter]

... OR ...

• [BindDistinguishedName] and [AttributeBindTypeOrValue]
*/
func (r *LDAPURI) Set(x ...any) LDAPURI {
	if r.IsZero() {
		r.ldapURI = newLDAPURI(x...)
		return *r
	}

	r.ldapURI.set(x...)
	return *r
}

/*
set is a private method called by LDAPURI.Set.
*/
func (r *ldapURI) set(x ...any) {
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
			r.scope = tv

		case AttributeBindTypeOrValue:
			// Value is an AttributeBindTypeOr Value
			r.avbt = tv

		case SearchFilter:
			// Value is an LDAP Search Filter
			r.filter = tv

		case AttributeTypes:
			// Value(s) are one or more LDAP
			// Search Attributes
			tv.transfer(r.attrs)
		}
	}

	return
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r LDAPURI) Compare(x any) bool {
	return compareHashInstance(r, x)
}

func (r LDAPURI) isDistinguishedNameContext() {}

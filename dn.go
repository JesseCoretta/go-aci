package aci

/*
dn.go contains LDAP Distinguished Name types and methods.
*/

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

	// AllDN is the BindRule abstraction of all *known user* DNs; this does not imply ANONYMOUS DNs
	AllDN BindDistinguishedName

	// AnyDN is the BindRule abstraction of all user DNs, known or anonymous
	AnyDN BindDistinguishedName

	// SelfDN is the BindRule abstraction of a user's own DN
	SelfDN BindDistinguishedName

	// ParentDN is the BindRule abstraction of a user's superior DN
	ParentDN BindDistinguishedName

	// badBindDN is an empty BindDistinguishedName struct returned when
	// a DN operation fails for some reason.
	badBindDN BindDistinguishedName

	// targetBindDN is an empty TargetDistinguishedName struct returned
	// when a DN operation fails for some reason.
	badTargetDN TargetDistinguishedName
)

/*
BindDistinguishedName describes a single LDAP distinguished name. For example:

	ou=People,dc=example,dc=com

For efficiency reasons, the LDAP "local scheme" prefix (ldap:///) is not
stored in literal form within any DistinguishedName instance, however it
will appear during string representation operations, e.g.:

	ldap:///ou=People,dc=example,dc=com

Instances of this kind can be crafted using the DN package-level function
with the appropriate Bind keyword as the input argument:

• DN(<dn>, BindUDN) for a `userdn` Distinguished Name

• DN(<dn>, BindGDN) for a `groupdn` Distinguished Name

• DN(<dn>, BindRDN) for a `roledn` Distinguished Name

In order to fashion multi-valued BindRule instances using instances of this
type, they must reside within an appropriate stack type instance. See the
BindDistinguishedNames and TargetDistinguishedNames types for details.
*/
type BindDistinguishedName struct {
	*distinguishedName
}

/*
TargetDistinguishedName describes a single LDAP distinguished name. For
example:

	ou=People,dc=example,dc=com

For efficiency reasons, the LDAP "local scheme" prefix (ldap:///) is not
stored in literal form within any DistinguishedName instance, however it
will appear during string representation operations, e.g.:

	ldap:///ou=People,dc=example,dc=com

Instances of this kind can be crafted using the DN package-level function
with the appropriate Target keyword as the input argument:

• DN(<dn>, Target) for a `target` Distinguished Name

• DN(<dn>, TargetTo) for a `target_to` Distinguished Name

• DN(<dn>, TargetFrom) for a `target_from` Distinguished Name

In order to fashion multi-valued TargetRule instances using values of this type,
they must reside within an appropriate stack type instance. For further details,
see the BindDistinguishedNames and TargetDistinguishedNames types.
*/
type TargetDistinguishedName struct {
	*distinguishedName
}

/*
distinguishedName is the embedded type (as a pointer!) within instances of
DistinguishedName.

The following TargetRule keywords allow the (indirect) use of instances
of this type:

• target

• target_to

• target_from

The following BindRule keywords allow the (indirect) use of instances
of this type:

• userdn

• roledn

• groupdn
*/
type distinguishedName struct {
	Keyword // `target`, `target_[to|from]` `userdn`, `groupdn` or `roledn`
	*string
}

/*
Valid returns an instance of error that reflects whether certain
required elements or value combinations were present and deemed
valid. A non-nil error indicates an undesirable receiver state.
*/
func (r BindDistinguishedName) Valid() (err error) {
	return validDistinguishedName(r)
}

/*
Valid returns an instance of error that reflects whether certain
required elements or value combinations were present and deemed
valid. A non-nil error indicates an undesirable receiver state.
*/
func (r TargetDistinguishedName) Valid() (err error) {
	return validDistinguishedName(r)
}

func validDistinguishedName(x any) (err error) {
	switch tv := x.(type) {
	case BindDistinguishedName:
		if tv.IsZero() {
			err = nilInstanceErr(tv)
		} else if len(*tv.distinguishedName.string) < 3 {
			err = illegalSyntaxPerTypeErr(*tv.distinguishedName, tv.Keyword())
		}
	case TargetDistinguishedName:
		if tv.IsZero() {
			err = nilInstanceErr(tv)
		} else if len(*tv.distinguishedName.string) < 3 {
			err = illegalSyntaxPerTypeErr(*tv.distinguishedName, tv.Keyword())
		}
	}

	return
}

/*
Keyword returns the Keyword (interface) assigned to the receiver instance. This shall be the keyword that appears in a BindRule bearing the receiver as a condition value.
*/
func (r BindDistinguishedName) Keyword() Keyword {
	if err := r.Valid(); err != nil {
		return nil
	}
	return r.distinguishedName.Keyword
}

/*
Keyword returns the Keyword (interface) assigned to the receiver instance. This shall be the keyword that appears in a TargetRule bearing the receiver as a condition value.
*/
func (r TargetDistinguishedName) Keyword() Keyword {
	if err := r.Valid(); err != nil {
		return nil
	}
	return r.distinguishedName.Keyword
}

/*
Kind returns the string name `bind`.
*/
func (r BindDistinguishedName) Kind() string {
	return bindRuleID
}

/*
Kind returns the string name `target`.
*/
func (r TargetDistinguishedName) Kind() string {
	return targetRuleID
}

/*
String is a stringer method that returns the string representation
of the receiver instance.

The Local LDAP scheme (ldap:///) is automatically imposed during
the string representation of the value; this is required by the
ACIv3 syntax.
*/
func (r BindDistinguishedName) String() string {
	if err := r.Valid(); err != nil {
		return ``
	}

	return sprintf("%s%s", LocalScheme, (*r.distinguishedName.string))
}

/*
String is a stringer method that returns the string representation
of the receiver instance.

The Local LDAP scheme (ldap:///) is automatically imposed during
the string representation of the value; this is required by the
ACIv3 syntax.
*/
func (r TargetDistinguishedName) String() string {
	if err := r.Valid(); err != nil {
		return ``
	}

	return sprintf("%s%s", LocalScheme, (*r.distinguishedName.string))
}

/*
IsZero returns a boolean value indicative of whether the receiver
is considered nil, or unset.
*/
func (r BindDistinguishedName) IsZero() bool {
	return r.distinguishedName.isZero()
}

/*
IsZero returns a boolean value indicative of whether the receiver
is considered nil, or unset.
*/
func (r TargetDistinguishedName) IsZero() bool {
	return r.distinguishedName.isZero()
}

/*
Set assigns value x to the receiver to represent an LDAP distinguished name in
the context of a BindRule.

This method presents an opportunity for setting a DN at a later point versus
doing so during the initialization process alone and is totally optional.

If no keyword is specified, the userdn keyword context is supplied automatically,
which may or may not be what the caller wants.
*/
func (r *BindDistinguishedName) Set(x string, kw ...Keyword) BindDistinguishedName {

	// default keyword, if unspecified by caller,
	// is the main DN context `userdn`.
	var key BindKeyword = BindUDN
	if len(kw) > 0 {
		// verify the provided keyword belongs
		// in the bindrule "family" context
		if kw[0].Kind() == bindRuleID {
			k := kw[0].(BindKeyword)

			// perform a keyword switch
			switch k {
			case BindUDN, BindGDN, BindRDN:
				// keyword is verified to be
				// related to bindrule DNs of
				// some kind.
				key = k
			default:
				return *r
			}
		} else {
			return *r
		}
	}

	// if the receiver was not initialized, do so now
	if r.IsZero() {
		*r = BindDistinguishedName{newDistinguishedName(x, key)}
	}

	// set it and go
	r.distinguishedName.set(x, key)
	return *r
}

/*
Set assigns value x to the receiver to represent an LDAP distinguished name in
the context of a TargetRule.

This method presents an opportunity for setting a DN at a later point versus
doing so during the initialization process alone and is totally optional.

If no keyword is specified, the target keyword context is supplied automatically,
which may or may not be what the caller wants.
*/
func (r *TargetDistinguishedName) Set(x string, kw ...Keyword) TargetDistinguishedName {

	// default keyword, if unspecified by caller,
	// is the main DN context `target`.
	var key TargetKeyword = Target
	if len(kw) > 0 {
		// verify the provided keyword belongs in the
		// targetrule "family" context
		if kw[0].Kind() == targetRuleID {
			k := kw[0].(TargetKeyword)

			// perform a keyword switch
			switch k {
			case Target, TargetTo, TargetFrom:
				// keyword is verified to be
				// related to targetrule DNs
				// of some kind.
				key = k
			default:
				return *r
			}
		} else {
			return *r
		}
	}

	// if the receiver was not initialized, do so now
	if r.IsZero() {
		*r = TargetDistinguishedName{newDistinguishedName(x, key)}
	}

	// set it and go
	r.distinguishedName.set(x, key)
	return *r
}

/*
isZero is a private method called by DistinguishedName.IsZero.
*/
func (r *distinguishedName) isZero() bool {
	return r == nil
}

/*
set is a private method called during the assembly of an underlying
Target or Bind (embedded) distinguishedName instance. This presents
an opportunity for setting a DN at a later point versus doing so
during the initialization process.
*/
func (r *distinguishedName) set(x string, kw Keyword) {
	_r := newDistinguishedName(x, kw)
	if len(*_r.string) == 0 {
		return
	}
	*r = *_r
}

/*
UDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Bind Rule Condition that bears the
`userdn` keyword.
*/
func UDN(x string) BindDistinguishedName {
	return BindDistinguishedName{newDistinguishedName(x, BindUDN)}
}

/*
RDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Bind Rule Condition that bears the
`roledn` keyword.
*/
func RDN(x string) BindDistinguishedName {
	return BindDistinguishedName{newDistinguishedName(x, BindRDN)}
}

/*
GDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Bind Rule Condition that bears the
`groupdn` keyword.
*/
func GDN(x string) BindDistinguishedName {
	return BindDistinguishedName{newDistinguishedName(x, BindGDN)}
}

/*
TDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Target Rule Condition that bears the
`target` keyword.
*/
func TDN(x string) TargetDistinguishedName {
	return TargetDistinguishedName{newDistinguishedName(x, Target)}
}

/*
TTDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Target Rule Condition that bears the
`target_to` keyword.
*/
func TTDN(x string) TargetDistinguishedName {
	return TargetDistinguishedName{newDistinguishedName(x, TargetTo)}
}

/*
TFDN initializes, sets and returns an instance of DistinguishedName in one shot.

An LDAP distinguished name, in string form and WITHOUT the leading `ldap:///` scheme, is required.

The return value shall be suitable for use in creating a Target Rule Condition that bears the
`target_from` keyword.
*/
func TFDN(x string) TargetDistinguishedName {
	return TargetDistinguishedName{newDistinguishedName(x, TargetFrom)}
}

/*
newDistinguishedName is a private function that returns a new instance of
*distinguishedName. This function is called by the UDN, RDN, GDN, TDN, TTDN
and TFDN functions.
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
Category wraps go-stackage's Stack.Category method.
*/
func (r BindDistinguishedNames) Category() string {
	if r.IsZero() {
		return ``
	}
	x, _ := castAsStack(r)
	return x.Category()
}

/*
Category wraps go-stackage's Stack.Category method.
*/
func (r TargetDistinguishedNames) Category() string {
	if r.IsZero() {
		return ``
	}
	x, _ := castAsStack(r)
	return x.Category()
}

/*
ID wraps go-stackage's Stack.ID method.
*/
func (r BindDistinguishedNames) ID() string {
	if r.IsZero() {
		return ``
	}
	x, _ := castAsStack(r)
	return x.ID()
}

/*
ID wraps go-stackage's Stack.ID method.
*/
func (r TargetDistinguishedNames) ID() string {
	if r.IsZero() {
		return ``
	}
	x, _ := castAsStack(r)
	return x.ID()
}

/*
setQuoteStyle shall set the receiver instance to the quotation
scheme defined by integer i.
*/
func (r BindDistinguishedNames) setQuoteStyle(style int) {
	_r, _ := castAsStack(r)
	if _r.Len() < 2 {
		_r.Encap() // not multivalued, force default
		return
	}

	if style == MultivalSliceQuotes {
		_r.Encap(`"`)
	} else {
		_r.Encap()
	}

	return
}

/*
setQuoteStyle shall set the receiver instance to the quotation
scheme defined by integer i.
*/
func (r TargetDistinguishedNames) setQuoteStyle(style int) {
	_r, _ := castAsStack(r)
	if _r.Len() < 2 {
		_r.Encap() // not multivalued, force default
		return
	}

	if style == MultivalSliceQuotes {
		_r.Encap(`"`)
	} else {
		_r.Encap()
	}

	return
}

/*
Eq initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Equal-To one (1) of the following keyword
contexts:

• `userdn`

• `groupdn`

• `roledn`
*/
func (r BindDistinguishedName) Eq() BindRule {
	x, ok := dnToCondition(r, Eq)
	if !ok {
		return badBindRule
	}

	return x.(BindRule)
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To one (1) of the following keyword
contexts:

• `target`

• `target_to`

• `target_from`
*/
func (r TargetDistinguishedName) Eq() TargetRule {
	x, ok := dnToCondition(r, Eq)
	if !ok {
		return badTargetRule
	}

	return x.(TargetRule)
}

/*
Ne initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Not-Equal-To one (1) of the following
keyword contexts:

• `userdn`

• `groupdn`

• `roledn`
*/
func (r BindDistinguishedName) Ne() BindRule {
	x, ok := dnToCondition(r, Ne)
	if !ok {
		return badBindRule
	}

	return x.(BindRule)
}

/*
Ne initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Not-Equal-To one (1) of the following
keyword contexts:

• `target`

• `target_to`

• `target_from`
*/
func (r TargetDistinguishedName) Ne() TargetRule {
	x, ok := dnToCondition(r, Ne)
	if !ok {
		return badTargetRule
	}

	return x.(TargetRule)
}

/*
Eq initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Equal-To one (1) of the following keyword
contexts:

• `userdn`

• `groupdn`

• `roledn`
*/
func (r BindDistinguishedNames) Eq() BindRule {
	x, ok := dnToCondition(r, Eq)
	if !ok {
		return badBindRule
	}

	return x.(BindRule)
}

/*
Eq initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Equal-To one (1) of the following keyword
contexts:

• `target`

• `target_to`

• `target_from`
*/
func (r TargetDistinguishedNames) Eq() TargetRule {
	x, ok := dnToCondition(r, Eq)
	if !ok {
		return badTargetRule
	}

	return x.(TargetRule)
}

/*
Ne initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Not-Equal-To one (1) of the following
keyword contexts:

• `userdn`

• `groupdn`

• `roledn`

Negated equality BindRule instances should be used with caution.
*/
func (r BindDistinguishedNames) Ne() BindRule {
	x, ok := dnToCondition(r, Ne)
	if !ok {
		return badBindRule
	}

	return x.(BindRule)
}

/*
Ne initializes and returns a new TargetRule instance configured to express the
evaluation of the receiver value as Not-Equal-To one (1) of the following
keyword contexts:

• `target`

• `target_to`

• `target_from`

Negated equality TargetRule instances should be used with caution.
*/
func (r TargetDistinguishedNames) Ne() TargetRule {
	x, ok := dnToCondition(r, Ne)
	if !ok {
		return badTargetRule
	}

	return x.(TargetRule)
}

func dnToCondition(dest any, op ComparisonOperator) (any, bool) {
	switch tv := dest.(type) {

	// case matched bind rule DN(s)
	case BindDistinguishedNames:
		return bindDNToCondition(tv, op)

	// case matched target rule DN(s)
	case TargetDistinguishedNames:
		return targetDNToCondition(tv, op)

	case BindDistinguishedName:
		if tv.Kind() == bindRuleID {
			return bindDNToCondition(tv, op)
		}
	case TargetDistinguishedName:
		if tv.Kind() == targetRuleID {
			return targetDNToCondition(tv, op)
		}
	}

	return badBindRule, false
}

func bindDNToCondition(dest any, op ComparisonOperator) (BindRule, bool) {
	var b BindRule
	switch tv := dest.(type) {
	case BindDistinguishedName:
		if tv.IsZero() {
			return badBindRule, false
		}

		if matchBKW(tv.Keyword().String()) == BindKeyword(0x0) {
			return badBindRule, false
		}

		// initialize our BindRule condition
		// with the needed keyword, operator
		// and DN value
		b.SetKeyword(tv.Keyword().String())
		b.SetOperator(op)
		b.SetExpression(tv)

	case BindDistinguishedNames:
		if tv.IsZero() {
			return badBindRule, false
		}

		if matchBKW(tv.Keyword().String()) == BindKeyword(0x0) {
			return badBindRule, false
		}

		// initialize our BindRule condition
		// with the needed keyword, operator
		// and DN values
		b.SetKeyword(tv.Keyword().String())
		b.SetOperator(op)
		b.SetExpression(tv)

	default:
		return badBindRule, false
	}

	// Cast to a stackage.Condition momentarily
	// just so we can run some needed methods
	// that are not exported due to a desire for
	// pkg simplicity.
	_b := castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(b.Category())

	return BindRule(*_b), true
}

func targetDNToCondition(dest any, op ComparisonOperator) (TargetRule, bool) {
	var t TargetRule
	switch tv := dest.(type) {
	case TargetDistinguishedName:
		if tv.IsZero() {
			return badTargetRule, false
		}

		if matchTKW(tv.Keyword().String()) == TargetKeyword(0x0) {
			return badTargetRule, false
		}

		// initialize our BindRule condition
		// with the needed keyword, operator
		// and DN value
		t.SetKeyword(tv.Keyword().String())
		t.SetOperator(op)
		t.SetExpression(tv)

	case TargetDistinguishedNames:
		if tv.IsZero() {
			return badTargetRule, false
		}

		if matchTKW(tv.Keyword().String()) == TargetKeyword(0x0) {
			return badTargetRule, false
		}

		// initialize our BindRule condition
		// with the needed keyword, operator
		// and DN values
		t.SetKeyword(tv.Keyword().String())
		t.SetOperator(op)
		t.SetExpression(tv)

	default:
		return badTargetRule, false
	}

	// Cast to a stackage.Condition momentarily
	// just so we can run some needed methods
	// that are not exported due to a desire for
	// pkg simplicity.
	_t := castAsCondition(t).
		Encap(`"`).
		Paren(true).
		SetID(targetRuleID).
		NoPadding(!RulePadding).
		SetCategory(t.Category())

	return TargetRule(*_t), true
}

/*
setExpressionValues is a private method called by assertBindUGRDN for
DN-based Bind Rules parsing.
*/
func (r BindDistinguishedNames) setExpressionValues(key Keyword, values ...string) (err error) {
	// iterate each string-based distinguishedName
	// in the values sequence ...
	for i := 0; i < len(values); i++ {

		// Identify this distinguished name value
		// as D, as referenced by index integer i.
		D := values[i]

		// Push DN into receiver. If the DN has
		// the LocalScheme (ldap:///) prefix, we
		// will chop it off as it is not needed.
		if hasPfx(D, LocalScheme) {
			r.Push(BindDistinguishedName{newDistinguishedName(D[len(LocalScheme):], key)})
		} else {
			// this will probably never match when
			// the data is coming from ANTLR ...
			r.Push(BindDistinguishedName{newDistinguishedName(D, key)})
		}
	}

	return
}

/*
setExpressionValues is a private method called by assertTargetTFDN for
DN-based Target Rules parsing.
*/
func (r TargetDistinguishedNames) setExpressionValues(key Keyword, values ...string) (err error) {
	// iterate each string-based distinguishedName
	// in the values sequence ...
	for i := 0; i < len(values); i++ {

		// Identify this distinguished name value
		// as D, as referenced by index integer i.
		D := values[i]

		// Push DN into receiver. If the DN has
		// the LocalScheme (ldap:///) prefix, we
		// will chop it off as it is not needed.
		if hasPfx(D, LocalScheme) {
			r.Push(TargetDistinguishedName{newDistinguishedName(D[len(LocalScheme):], key)})
		} else {
			// this will probably never match when
			// the data is coming from ANTLR ...
			r.Push(TargetDistinguishedName{newDistinguishedName(D, key)})
		}
	}

	return
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r BindDistinguishedNames) IsZero() bool {
	_r, _ := castAsStack(r)
	return _r.IsZero()
}

/*
IsZero wraps go-stackage's Stack.IsZero method.
*/
func (r TargetDistinguishedNames) IsZero() bool {
	_r, _ := castAsStack(r)
	return _r.IsZero()
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r BindDistinguishedNames) Len() int {
	_r, _ := castAsStack(r)
	return _r.Len()
}

/*
Len wraps go-stackage's Stack.Len method.
*/
func (r TargetDistinguishedNames) Len() int {
	_r, _ := castAsStack(r)
	return _r.Len()
}

/*
Index wraps go-stackage's Stack.Index method. Note that the
Boolean OK value returned by go-stackage's Stack.Index method
by default will be shadowed and not obtainable by the caller.
*/
func (r BindDistinguishedNames) Index(idx int) BindDistinguishedName {
	if assert, ok := distinguishedNameIndex(r, idx).(BindDistinguishedName); ok {
		return assert
	}

	return badBindDN
}

/*
Index wraps go-stackage's Stack.Index method. Note that the
Boolean OK value returned by go-stackage by default will be
shadowed and not obtainable by the caller.
*/
func (r TargetDistinguishedNames) Index(idx int) TargetDistinguishedName {
	if assert, ok := distinguishedNameIndex(r, idx).(TargetDistinguishedName); ok {
		return assert
	}

	return badTargetDN
}

/*
distinguishedNameIndex is a private index function called through instances
of BindDistinguishedName and TargetDistinguishedName.
*/
func distinguishedNameIndex(r any, idx int) any {
	_r, _ := castAsStack(r)
	y, _ := _r.Index(idx)
	if assert, ok := y.(BindDistinguishedName); ok {
		return assert
	} else if assert2, ok2 := y.(TargetDistinguishedName); ok2 {
		return assert2
	}

	return nil
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Stack.String method.
*/
func (r BindDistinguishedNames) String() string {
	_r, _ := castAsStack(r)
	return _r.String()
}

/*
String is a stringer method that returns the string
representation of the receiver instance.

This method wraps go-stackage's Stack.String method.
*/
func (r TargetDistinguishedNames) String() string {
	_r, _ := castAsStack(r)
	return _r.String()
}

/*
Keyword returns the Keyword (interface) assigned to the receiver instance.
This shall be the keyword that appears in a BindRule bearing the receiver
as a condition value.
*/
func (r BindDistinguishedNames) Keyword() Keyword {
	return keywordFromCategory(r)
}

/*
Keyword returns the Keyword (interface) assigned to the receiver instance.
This shall be the keyword that appears in a TargetRule bearing the receiver
as a condition value.
*/
func (r TargetDistinguishedNames) Keyword() Keyword {
	return keywordFromCategory(r)
}

/*
Push wraps go-stackage's Stack.Push method. Valid input types
are string and BindDistinguishedName.

In the case of a string value, it is automatically cast as an
instance of BindDistinguishedName using the appropriate keyword,
so long as the raw string is of a non-zero length.
*/
func (r BindDistinguishedNames) Push(x ...any) BindDistinguishedNames {
	_r, _ := castAsStack(r)
	kw := keywordFromCategory(r)
	if kw == nil {
		// not initialized?!
		return r
	}

	// iterate variadic input arguments
	for i := 0; i < len(x); i++ {
		// verify DN value and (possibly) cast from
		// string->dn automatically.
		dn, ok := pushBindDistinguishedNames(kw, x[i])
		if !ok {
			return r
		}

		// Push it!
		_r.Push(dn)
	}

	return r
}

/*
Contains returns a boolean value indicative of whether value x,
if a string or BindDistinguishedName instance, already resides
within the receiver instance.

Case is not significant in the matching process.
*/
func (r BindDistinguishedNames) Contains(x any) bool {
	return r.contains(x)
}

/*
contains is a private method called by BindDistinguishedNames.Contains.
*/
func (r BindDistinguishedNames) contains(x any) bool {
	if r.Len() == 0 {
		return false
	}

	var candidate string

	switch tv := x.(type) {
	case string:
		candidate = tv
	case BindDistinguishedName:
		candidate = tv.String()
	default:
		return false
	}

	if len(candidate) == 0 {
		return false
	}

	for i := 0; i < r.Len(); i++ {
		// case is not significant here.
		if eq(r.Index(i).String(), candidate) {
			return true
		}
	}

	return false
}

/*
Push wraps go-stackage's Stack.Push method. Valid input types
are string and TargetDistinguishedName.

In the case of a string value, it is automatically cast as an
instance of TargetDistinguishedName using the appropriate keyword,
so long as the raw string is of a non-zero length.
*/
func (r TargetDistinguishedNames) Push(x ...any) TargetDistinguishedNames {
	_r, _ := castAsStack(r)
	kw := keywordFromCategory(r)
	if kw == nil {
		// not initialized?!
		return r
	}

	// iterate variadic input arguments
	for i := 0; i < len(x); i++ {
		// verify DN value and (possibly) cast from
		// string->dn automatically.
		dn, ok := pushTargetDistinguishedNames(kw, x[i])
		if !ok {
			return r
		}

		// Push it!
		_r.Push(dn)
	}

	return r
}

/*
Contains returns a boolean value indicative of whether value x,
if a string or TargetDistinguishedName instance, already resides
within the receiver instance.

Case is not significant in the matching process.
*/
func (r TargetDistinguishedNames) Contains(x any) bool {
	return r.contains(x)
}

/*
contains is a private method called by TargetDistinguishedNames.Contains.
*/
func (r TargetDistinguishedNames) contains(x any) bool {
	if r.Len() == 0 {
		return false
	}

	var candidate string

	switch tv := x.(type) {
	case string:
		candidate = tv
	case TargetDistinguishedName:
		candidate = tv.String()
	default:
		return false
	}

	if len(candidate) == 0 {
		return false
	}

	for i := 0; i < r.Len(); i++ {
		// case is not significant here.
		if eq(r.Index(i).String(), candidate) {
			return true
		}
	}

	return false
}

func pushBindDistinguishedNames(kw Keyword, x any) (BindDistinguishedName, bool) {
	// perform an input type switch,
	// allowing evaluation of the
	// value.
	switch tv := x.(type) {

	// case match is a DistinguishedName in string form.
	// In this case, it will be marshaled into a proper
	// instance of DistinguishedName bearing the same
	// ke yword as the destination receiver.
	case string:
		if len(tv) == 0 {
			return badBindDN, false
		}

		return BindDistinguishedName{newDistinguishedName(tv, kw)}, true

	// case match is a proper BindDistinguishedName instance.
	// Both keywords (that of the BindDistinguishedName, and
	// that of the destination receiver) must match.
	case BindDistinguishedName:
		if tv.IsZero() {
			return badBindDN, false
		}

		// attempting to push a DN that bears a
		// different keyword than the receiver.
		// will stop the show.
		if dkw := keywordFromCategory(tv); dkw.String() != kw.String() {
			return badBindDN, false
		}

		return tv, true
	}

	return badBindDN, false
}

func pushTargetDistinguishedNames(kw Keyword, x any) (TargetDistinguishedName, bool) {
	// perform an input type switch,
	// allowing evaluation of the
	// value.
	switch tv := x.(type) {

	// case match is a DistinguishedName in string form.
	// In this case, it will be marshaled into a proper
	// instance of DistinguishedName bearing the same
	// ke yword as the destination receiver.
	case string:
		if len(tv) == 0 {
			return badTargetDN, false
		}

		return TargetDistinguishedName{newDistinguishedName(tv, kw)}, true

	// case match is a proper TargetDistinguishedName instance.
	// Both keywords (that of the TargetDistinguishedName, and
	// that of the destination receiver) must match.
	case TargetDistinguishedName:
		if tv.IsZero() {
			return badTargetDN, false
		}

		// attempting to push a DN that bears a
		// different keyword than the receiver.
		// will stop the show.
		if tv.Keyword() != kw {
			return badTargetDN, false
		}

		return tv, true
	}

	return badTargetDN, false
}

/*
Pop wraps go-stackage's Stack.Pop method and performs type
assertion to return a proper BindDistinguishedName instance.
*/
func (r BindDistinguishedNames) Pop() BindDistinguishedName {
	_r, _ := castAsStack(r)

	y, _ := _r.Pop()
	if assert, asserted := y.(BindDistinguishedName); asserted {
		return assert
	}

	return badBindDN
}

/*
Pop wraps go-stackage's Stack.Pop method and performs type
assertion to return a proper TargetDistinguishedName instance.
*/
func (r TargetDistinguishedNames) Pop() TargetDistinguishedName {
	_r, _ := castAsStack(r)

	y, _ := _r.Pop()
	if assert, asserted := y.(TargetDistinguishedName); asserted {
		return assert
	}

	return badTargetDN
}

/*
uDNPushPolicy is a private function that conforms to go-stackage's
PushPolicy interface signature. This is called during Push attempts
to a stack containing BindRule userdn distinguished name instances.
*/
func (r BindDistinguishedNames) uDNPushPolicy(x any) error {
	if r.contains(x) {
		return pushErrorNotUnique(r, x, r.Keyword())
	}
	return distinguishedNamesPushPolicy(r, x, BindUDN)
}

/*
gDNPushPolicy is a private function that conforms to go-stackage's
PushPolicy interface signature. This is called during Push attempts
to a stack containing BindRule groupdn distinguished name instances.
*/
func (r BindDistinguishedNames) gDNPushPolicy(x any) error {
	if r.contains(x) {
		return pushErrorNotUnique(r, x, r.Keyword())
	}
	return distinguishedNamesPushPolicy(r, x, BindGDN)
}

/*
rDNPushPolicy is a private function that conforms to go-stackage's
PushPolicy interface signature. This is called during Push attempts
to a stack containing BindRule roledn distinguished name instances.
*/
func (r BindDistinguishedNames) rDNPushPolicy(x any) error {
	if r.contains(x) {
		return pushErrorNotUnique(r, x, r.Keyword())
	}
	return distinguishedNamesPushPolicy(r, x, BindRDN)
}

/*
tToDNPushPolicy is a private function that conforms to go-stackage's
PushPolicy interface signature. This is called during Push attempts
to a stack containing TargetRule target_to distinguished name instances.
*/
func (r TargetDistinguishedNames) tToDNPushPolicy(x any) error {
	if r.contains(x) {
		return pushErrorNotUnique(r, x, r.Keyword())
	}
	return distinguishedNamesPushPolicy(r, x, TargetTo)
}

/*
tFromDNPushPolicy is a private function that conforms to go-stackage's
PushPolicy interface signature. This is called during Push attempts to
a stack containing TargetRule target_from distinguished name instances.
*/
func (r TargetDistinguishedNames) tFromDNPushPolicy(x any) error {
	if r.contains(x) {
		return pushErrorNotUnique(r, x, r.Keyword())
	}
	return distinguishedNamesPushPolicy(r, x, TargetFrom)
}

/*
tDNPushPolicy is a private function that conforms to go-stackage's PushPolicy
interface signature. This is called during Push attempts to a stack containing
TargetRule target distinguished name instances.
*/
func (r TargetDistinguishedNames) tDNPushPolicy(x any) error {
	if r.contains(x) {
		return pushErrorNotUnique(r, x, r.Keyword())
	}
	return distinguishedNamesPushPolicy(r, x, Target)
}

/*
distinguishedNamesPushPolicy is the backend worker called by all of
the keyword-specific DN pushPolicy functions above. This function
handles any type of DN value and DN stack.
*/
func distinguishedNamesPushPolicy(r, x any, kw Keyword) (err error) {
	switch tv := x.(type) {

	case string:
		if len(tv) == 0 {
			err = pushErrorNilOrZero(r, tv, kw)
		}

	case BindDistinguishedName:
		if tv.IsZero() {
			err = pushErrorNilOrZero(r, tv, kw)

		} else if tv.distinguishedName.Keyword != kw {
			err = badPTBRuleKeywordErr(tv, bindRuleID, kw, tv.distinguishedName.Keyword)
		}

	case TargetDistinguishedName:
		if tv.IsZero() {
			err = pushErrorNilOrZero(r, tv, kw)

		} else if tv.distinguishedName.Keyword != kw {
			err = badPTBRuleKeywordErr(tv, targetRuleID, kw, tv.distinguishedName.Keyword)
		}

	default:
		err = pushErrorBadType(r, tv, kw)
	}

	return
}

/*
UDNs returns a new instance of BindDistinguishedNames with an initialized
embedded stack configured to function as a simple ORed list containing
a single level of LDAP distinguished names. The BindUDN (userdn) keyword
is automatically assigned to the return value.

Only valid instances of BindDistinguishedName which bear the BindUDN keyword
are to be considered eligible for push requests. If the input value is a
string, it will be accepted and properly branded with the keyword.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
BindDistinguishedName type (or its string equivalent) for push during initialization.
This is merely a more convenient alternative to separate init and push procedures.

See also the RDNs and GDNs functions for BindRDN and BindGDN keywords
respectively.
*/
func UDNs(x ...any) (d BindDistinguishedNames) {
	// create a native stackage.Stack
	// and configure before typecast.
	_d := stackOr().
		Symbol(`||`).
		NoNesting(true).
		SetID(bindRuleID).
		NoPadding(!StackPadding).
		SetCategory(BindUDN.String())

	// cast _d as a proper BindDistinguishedName
	// instance (d). We do it this way to gain
	// access to the method for the *specific
	// instance* being created (d), thus allowing
	// things like uniqueness checks, etc., to
	// occur during push attempts, providing more
	// helpful and non-generalized feedback.
	d = BindDistinguishedNames(_d)
	_d.SetPushPolicy(d.uDNPushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_d.Push(x...)

	return
}

/*
RDNs returns a new instance of BindDistinguishedNames with an initialized
embedded stack configured to function as a simple ORed list containing
a single level of LDAP distinguished names. The BindRDN (roledn) keyword
is automatically assigned to the return value.

Only valid instances of BindDistinguishedName which bear the BindRDN keyword
are to be considered eligible for push requests. If the input value is a
string, it will be accepted and properly branded with the keyword.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
BindDistinguishedName type (or its string equivalent) for push during initialization.
This is merely a more convenient alternative to separate init and push procedures.

See also the UDNs and GDNs functions for BindUDN and BindGDN keywords
respectively.
*/
func RDNs(x ...any) (d BindDistinguishedNames) {
	// create a native stackage.Stack
	// and configure before typecast.
	_d := stackOr().
		Symbol(`||`).
		NoNesting(true).
		SetID(bindRuleID).
		NoPadding(!StackPadding).
		SetCategory(BindRDN.String())

	// cast _d as a proper BindDistinguishedName
	// instance (d). We do it this way to gain
	// access to the method for the *specific
	// instance* being created (d), thus allowing
	// things like uniqueness checks, etc., to
	// occur during push attempts, providing more
	// helpful and non-generalized feedback.
	d = BindDistinguishedNames(_d)
	_d.SetPushPolicy(d.rDNPushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_d.Push(x...)

	return
}

/*
GDNs returns a new instance of BindDistinguishedNames with an initialized
embedded stack configured to function as a simple ORed list containing
a single level of LDAP distinguished names. The BindGDN (groupdn) keyword
is automatically assigned to the return value.

Only valid instances of BindDistinguishedName which bear the BindGDN keyword
are to be considered eligible for push requests. If the input value is a
string, it will be accepted and properly branded with the keyword.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
BindDistinguishedName type (or its string equivalent) for push during initialization.
This is merely a more convenient alternative to separate init and push procedures.

See also the UDNs and RDNs functions for BindUDN and BindRDN keywords
respectively.
*/
func GDNs(x ...any) (d BindDistinguishedNames) {
	// create a native stackage.Stack
	// and configure before typecast.
	_d := stackOr().
		Symbol(`||`).
		NoNesting(true).
		SetID(bindRuleID).
		NoPadding(!StackPadding).
		SetCategory(BindGDN.String())

	// cast _d as a proper BindDistinguishedName
	// instance (d). We do it this way to gain
	// access to the method for the *specific
	// instance* being created (d), thus allowing
	// things like uniqueness checks, etc., to
	// occur during push attempts, providing more
	// helpful and non-generalized feedback.
	d = BindDistinguishedNames(_d)
	_d.SetPushPolicy(d.gDNPushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_d.Push(x...)

	return
}

/*
TDNs returns a new instance of TargetDistinguishedNames with an initialized
embedded stack configured to function as a simple ORed list containing
a single level of LDAP distinguished names. The Target (target) keyword is
automatically assigned to the return value.

Only valid instances of TargetDistinguishedName which bear the Target keyword
are to be considered eligible for push requests. If the input value is a
string, it will be accepted and properly branded with the keyword.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
TargetDistinguishedName type (or its string equivalent) for push during initialization.
This is merely a more convenient alternative to separate init and push procedures.

See also the TTDNs and TFDNs functions for TargetTo and TargetFrom keywords
respectively.
*/
func TDNs(x ...any) (d TargetDistinguishedNames) {
	// create a native stackage.Stack
	// and configure before typecast.
	_d := stackOr().
		Symbol(`||`).
		NoNesting(true).
		SetID(targetRuleID).
		NoPadding(!StackPadding).
		SetCategory(Target.String())

	// cast _d as a proper TargetDistinguishedName
	// instance (d). We do it this way to gain
	// access to the method for the *specific
	// instance* being created (d), thus allowing
	// things like uniqueness checks, etc., to
	// occur during push attempts, providing more
	// helpful and non-generalized feedback.
	d = TargetDistinguishedNames(_d)
	_d.SetPushPolicy(d.tDNPushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_d.Push(x...)

	return
}

/*
TTDNs returns a new instance of TargetDistinguishedNames with an initialized
embedded stack configured to function as a simple ORed list containing
a single level of LDAP distinguished names. The TargetTo (target_to) keyword
is automatically assigned to the return value.

Only valid instances of TargetDistinguishedName which bear the TargetTo keyword
are to be considered eligible for push requests. If the input value is a
string, it will be accepted and properly branded with the keyword.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
TargetDistinguishedName type (or its string equivalent) for push during initialization.
This is merely a more convenient alternative to separate init and push procedures.

See also the TDNs and TFDNs functions for Target and TargetFrom keywords
respectively.
*/
func TTDNs(x ...any) (d TargetDistinguishedNames) {
	// create a native stackage.Stack
	// and configure before typecast.
	_d := stackOr().
		Symbol(`||`).
		NoNesting(true).
		SetID(targetRuleID).
		NoPadding(!StackPadding).
		SetCategory(TargetTo.String())

	// cast _d as a proper TargetDistinguishedName
	// instance (d). We do it this way to gain
	// access to the method for the *specific
	// instance* being created (d), thus allowing
	// things like uniqueness checks, etc., to
	// occur during push attempts, providing more
	// helpful and non-generalized feedback.
	d = TargetDistinguishedNames(_d)
	_d.SetPushPolicy(d.tToDNPushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_d.Push(x...)

	return
}

/*
TFDNs returns a new instance of TargetDistinguishedNames with an initialized
embedded stack configured to function as a simple ORed list containing a single
level of LDAP distinguished names. The TargetFrom (target_from) keyword will be
automatically assigned to the return value.

Optionally, the caller may choose to submit one (1) or more (valid) instances of the
TargetDistinguishedName type (or its string equivalent) for push during initialization.
This is merely a more convenient alternative to separate init and push procedures.

Only valid instances of TargetDistinguishedName which bear the TargetFrom
keyword are to be considered eligible for push requests. If the input value
is a string, it will be accepted and properly branded with the keyword.

See also the TDNs and TTDNs functions for Target and TargetTo keywords
respectively.
*/
func TFDNs(x ...any) (d TargetDistinguishedNames) {
	// create a native stackage.Stack
	// and configure before typecast.
	_d := stackOr().
		Symbol(`||`).
		NoNesting(true).
		SetID(targetRuleID).
		NoPadding(!StackPadding).
		SetCategory(TargetFrom.String())

	// cast _d as a proper TargetDistinguishedName
	// instance (d). We do it this way to gain
	// access to the method for the *specific
	// instance* being created (d), thus allowing
	// things like uniqueness checks, etc., to
	// occur during push attempts, providing more
	// helpful and non-generalized feedback.
	d = TargetDistinguishedNames(_d)
	_d.SetPushPolicy(d.tFromDNPushPolicy)

	// Assuming one (1) or more items were
	// submitted during the call, (try to)
	// push them into our initialized stack.
	// Note that any failed push(es) will
	// have no impact on the validity of
	// the return instance.
	_d.Push(x...)

	return
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

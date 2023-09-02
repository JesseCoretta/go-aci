package aci

/*
sec.go contains types, methods and constants that relate to
LDAP authentication methods and confidentiality control
through the SSF type.
*/

/*
AuthenticationMethod is a uint8 type that manifests through predefined
package constants, each describing a supported means of LDAP
authentication.
*/
type AuthenticationMethod uint8

var (
	authMap   map[int]AuthenticationMethod
	authNames map[string]AuthenticationMethod
)

/*
AuthenticationMethodLowerCase allows control over the case folding of
AuthenticationMethod string representation.

A value of true shall force lowercase normalization, while
a value of false (default) forces uppercase normalization.
*/
var AuthenticationMethodLowerCase bool

/*
AuthenticationMethod contants define the available LDAP authentication
mechanisms that are recognized within the ACI syntax honored
by this package.

NOTE: Supported SASL mechanisms vary per impl.
*/
const (
	noAuth    AuthenticationMethod = iota // invalid
	Anonymous                             // 0
	Simple                                // 1
	SSL                                   // 2
	SASL                                  // 3
	EXTERNAL                              // 4
	DIGESTMD5                             // 5
	GSSAPI                                // 6
)

/*
BRF returns an instance of BindRuleMethods.

Each of the return instance's key values represent a single instance of the
ComparisonOperator type that is allowed for use in the creation of BindRule
instances which bear the receiver instance as an expression value. The value
for each key is the actual BindRuleMethod instance for OPTIONAL use in the
creation of a BindRule instance.

This is merely a convenient alternative to maintaining knowledge of which
ComparisonOperator instances apply to which types. Instances of this type
are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not
yet been initialized, the execution of ANY of the return instance's value
methods will return bogus BindRule instances. While this is useful in unit
testing, the end user must only execute this method IF and WHEN the receiver
has been properly populated and prepared for such activity.
*/
func (r AuthenticationMethod) BRF() BindRuleMethods {
	return newBindRuleMethods(bindRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
Eq initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Equal-To the `authmethod` Bind keyword
context.
*/
func (r AuthenticationMethod) Eq() BindRule {
	if r == noAuth {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(BindAM)
	b.SetOperator(Eq)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindAM.String())

	return b
}

/*
Ne initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Not-Equal-To the `authmethod` Bind keyword
context.
*/
func (r AuthenticationMethod) Ne() BindRule {
	if r == noAuth {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(BindAM)
	b.SetOperator(Ne)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindAM.String())

	return b
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r AuthenticationMethod) String() (am string) {
	for k, v := range authNames {
		if v == r {
			am = foldAuthenticationMethod(k)
			break
		}
	}

	return
}

/*
Compare returns a Boolean indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r AuthenticationMethod) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
SecurityStrengthFactor embeds a pointer to uint8. A nil uint8 value indicates an effective
security strength factor of zero (0). A non-nil uint8 value expresses uint8 + 1, thereby
allowing a range of 0-256 "within" a uint8 instance.
*/
type SecurityStrengthFactor struct {
	*ssf
}

type ssf struct {
	*uint8
}

/*
SSF initializes, sets and returns a new instance of SecurityStrengthFactor in one
shot. This function is an alternative to separate assignment and set procedures.
*/
func SSF(factor ...any) SecurityStrengthFactor {
	s := SecurityStrengthFactor{new(ssf)}
	if len(factor) > 0 {
		s.ssf.set(factor[0])
	}
	return s
}

/*
Keyword returns the Keyword (interface) assigned to the receiver instance. This
shall be the keyword that appears in a BindRule containing the receiver instance
as the expression value.
*/
func (r SecurityStrengthFactor) Keyword() Keyword {
	return BindSSF
}

/*
value is a private method that returns uint8 + 1, or 0 if uint8 is nil.
*/
func (r SecurityStrengthFactor) value() int {
	if r.isZero() {
		return 0
	}
	return int(*(r.ssf.uint8)) + 1 // offset for 256 max.
}

/*
IsZero returns a boolean value indicative of whether the receiver is
nil, or unset.
*/
func (r SecurityStrengthFactor) IsZero() bool {
	if r.ssf == nil {
		return true
	}

	return r.ssf.isZero()
}

func (r *ssf) isZero() bool {
	if r == nil {
		return true
	}
	return r.uint8 == nil
}

/*
Eq initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Equal-To the `ssf` Bind keyword context.
*/
func (r SecurityStrengthFactor) Eq() BindRule {

	var b BindRule
	b.SetKeyword(BindSSF)
	b.SetOperator(Eq)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindSSF.String())

	return b
}

/*
Ne initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Not-Equal-To the `ssf` Bind keyword
context.

Negated equality BindRule instances should be used with caution.
*/
func (r SecurityStrengthFactor) Ne() BindRule {

	var b BindRule
	b.SetKeyword(BindSSF)
	b.SetOperator(Ne)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindSSF.String())

	return b
}

/*
Lt initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Less-Than the `ssf` Bind keyword context.
*/
func (r SecurityStrengthFactor) Lt() BindRule {

	var b BindRule
	b.SetKeyword(BindSSF)
	b.SetOperator(Lt)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindSSF.String())

	return b
}

/*
Le initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Less-Than-Or-Equal to the `ssf` Bind
keyword context.
*/
func (r SecurityStrengthFactor) Le() BindRule {

	var b BindRule
	b.SetKeyword(BindSSF)
	b.SetOperator(Le)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindSSF.String())

	return b
}

/*
Gt initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Greater-Than the `ssf` Bind keyword
context.
*/
func (r SecurityStrengthFactor) Gt() BindRule {

	var b BindRule
	b.SetKeyword(BindSSF)
	b.SetOperator(Gt)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindSSF.String())

	return b
}

/*
Ge initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Greater-Than-Or-Equal to the `ssf` Bind
keyword context.
*/
func (r SecurityStrengthFactor) Ge() BindRule {

	var b BindRule

	b.SetKeyword(BindSSF)
	b.SetOperator(Ge)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindSSF.String())

	return b
}

/*
BRF returns an instance of BindRuleMethods.

Each of the return instance's key values represent a single instance of the
ComparisonOperator type that is allowed for use in the creation of BindRule
instances which bear the receiver instance as an expression value. The value
for each key is the actual BindRuleMethod instance for OPTIONAL use in the
creation of a BindRule instance.

This is merely a convenient alternative to maintaining knowledge of which
ComparisonOperator instances apply to which types. Instances of this type
are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not
yet been initialized, the execution of ANY of the return instance's value
methods will return bogus BindRule instances. While this is useful in unit
testing, the end user must only execute this method IF and WHEN the receiver
has been properly populated and prepared for such activity.
*/
func (r SecurityStrengthFactor) BRF() BindRuleMethods {
	return newBindRuleMethods(bindRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
		Lt: r.Lt,
		Le: r.Le,
		Gt: r.Gt,
		Ge: r.Ge,
	})
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r SecurityStrengthFactor) String() string {
	if r.isZero() {
		return `0`
	}
	return sprintf("%d", int((*r.ssf.uint8))+1)
}

/*
Compare returns a Boolean indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r SecurityStrengthFactor) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Valid returns a boolean value indicative of whether the receiver represents
a security strength factor greater than zero (0).
*/
func (r SecurityStrengthFactor) Valid() (err error) {
	if r.value() == 0 {
		err = nilInstanceErr(r)
	}

	return
}

func (r SecurityStrengthFactor) clear() {
	if r.ssf.isZero() {
		return
	}
	r.ssf.clear()
}

func (r *ssf) clear() {
	r.uint8 = nil
}

/*
Set modifies the receiver to reflect the desired security strength factor (SSF),
which can represent any numerical value between 0 (off) and 256 (max).

Valid input types are int, string and nil.

A value of nil wipes out any previous value, making the SSF effectively zero (0).

A string value of `full` or `max` sets the SSF to its maximum value. A value of `none`
or `off` has the same effect as when providing a nil value. A numerical string value
is cast as int and (if valid) will be resubmitted silently. Case is not significant
during the string matching process.

An int value less than or equal to zero (0) has the same effect as when providing a
nil value. A value between 1 and 256 is acceptable and will be used. A value greater
than 256 will be silently reduced back to the maximum.
*/
func (r *SecurityStrengthFactor) Set(factor any) SecurityStrengthFactor {
	if r.ssf == nil {
		r.ssf = new(ssf)
		r.ssf.uint8 = new(uint8)
	}
	r.ssf.set(factor)
	return *r
}

/*
set is called by SecurityStrengthFactor.Set to modify the underlying uint8 pointer
in order to represent a security strength factor value.
*/
func (r *ssf) set(factor any) {
	switch tv := factor.(type) {
	case nil:
		r.clear()
	case string:
		i := stringToIntSSF(tv)
		if i == 0 {
			r.clear()
			return
		}
		r.set(i)
	case int:
		if tv > 256 {
			tv = 256
		} else if tv <= 0 {
			r.clear()
			return
		}

		v := uint8(tv - 1)
		r.uint8 = &v
	}

	return
}

func stringToIntSSF(x string) (i int) {
	switch lc(x) {
	case `full`, `max`:
		i = 256
	case `none`, `off`:
		i = 0
	default:
		i, _ = atoi(x)
	}

	return
}

/*
matchAuthenticationMethod resolves a given authentication method
based on an integer or string input (x). If no match,
Anonymous is returned.
*/
func matchAuthenticationMethod(x any) (am AuthenticationMethod) {
	am = Anonymous // anonymous is default

	switch tv := x.(type) {
	case int:
		for k, v := range authMap {
			if k == tv {
				am = v
				break
			}
		}
	case string:
		for k, v := range authNames {
			if eq(k, tv) {
				am = v
				break
			}
		}
	}

	return
}

/*
foldAuthenticationMethod executes the string representation
case-folding, per whatever value is assigned to the
global AuthenticationMethodLowerCase variable.
*/
func foldAuthenticationMethod(x string) string {
	if AuthenticationMethodLowerCase {
		return lc(x)
	}
	return uc(x)
}

func init() {

	// authMap facilitates lookups of AuthenticationMethod
	// instances using their underlying numerical const
	// value; this is mostly used internally.
	authMap = map[int]AuthenticationMethod{
		0: Anonymous,
		1: Simple,
		2: SSL,
		3: SASL,
		5: EXTERNAL,
		4: DIGESTMD5,
		6: GSSAPI,
	}

	// authNames facilities lookups of AuthenticationMethod
	// instances using their string representation. as the
	// lookup key.
	//
	// NOTE: case is not significant during string
	// *matching* (resolution); this is regardless
	// of the state of AuthenticationMethodLowerCase.
	authNames = map[string]AuthenticationMethod{
		`none`:   Anonymous, // anonymous is ALWAYS default
		`simple`: Simple,    // simple auth (DN + Password); no confidentiality is implied
		`ssl`:    SSL,       // authentication w/ confidentiality; SSL (LDAPS) and TLS (LDAP + STARTTLS)

		// NOTE: Supported SASL methods vary per impl.
		`sasl`:            SASL,      // *any* SASL mechanism
		`sasl EXTERNAL`:   EXTERNAL,  // only SASL/EXTERNAL mechanism, e.g.: TLS Client Auth w/ personal cert
		`sasl DIGEST-MD5`: DIGESTMD5, // only SASL/DIGEST-MD5 mechanism, e.g.: password encipherment
		`sasl GSSAPI`:     GSSAPI,    // only SASL/GSSAPI mechanism, e.g.: Kerberos Single Sign-On
	}
}

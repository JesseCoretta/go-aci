package aci

/*
sec.go contains types, methods and constants that relate to
LDAP authentication methods and confidentiality control
through the SSF type.
*/

/*
AuthMethod is a uint8 type that manifests through predefined
package constants, each describing a supported means of LDAP
authentication.
*/
type AuthMethod uint8

/*
AuthMethod contants define the available LDAP authentication
mechanisms that are recognized within the ACI syntax.
*/
const (
	Anonymous AuthMethod = iota // default
	Simple
	SSL
	SASL
)

/*
Eq initializes and returns a new *Condition instance configured
to evaluate AuthMethod as Equal-To the the request address.
*/
func (r AuthMethod) Eq() Condition {
	return Cond(BindAM, r, Eq).
		Encap(`"`).
		NoPadding(!ConditionPadding).
		setCategory(`authmethod`)
}

/*
Ne initializes and returns a new *Condition instance configured
to evaluate AuthMethod as Not-Equal-To the the request address.
*/
func (r AuthMethod) Ne() Condition {
	return Cond(BindAM, r, Ne).
		Encap(`"`).
		NoPadding(!ConditionPadding).
		setCategory(`authmethod`)
}

/*
String is a stringer method that returns the string representation
of the receiver instance.
*/
func (r AuthMethod) String() string {
	var am string = `none` // anon
	switch r {
	case Simple:
		am = `simple`
	case SSL:
		am = `SSL`
	case SASL:
		am = `SASL`
	}

	return am
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
Eq initializes and returns a new *Condition instance configured
to evaluate SSF as Equal-To the the request confidentility factor.

Value encapsulation in double-quotes (") is imposed.
*/
func (r SecurityStrengthFactor) Eq() Condition {
	return Cond(BindSSF, r, Eq).
		Encap(`"`).
		setCategory(`ssf`)
}

/*
Ne initializes and returns a new *Condition instance configured
to evaluate SSF as Not-Equal-To the the request confidentility factor.

Value encapsulation in double-quotes (") is imposed.
*/
func (r SecurityStrengthFactor) Ne() Condition {
	return Cond(BindSSF, r, Ne).
		Encap(`"`).
		NoPadding(!ConditionPadding).
		setCategory(`ssf`)
}

/*
Lt initializes and returns a new *Condition instance configured
to evaluate SSF as Less-Than the request confidentility factor.
*/
func (r SecurityStrengthFactor) Lt() Condition {
	return Cond(BindSSF, r, Lt).
		Encap(`"`).
		NoPadding(!ConditionPadding).
		setCategory(`ssf`)
}

/*
Gt initializes and returns a new *Condition instance configured
to evaluate SSF as Greather-Than the request confidentility factor.
*/
func (r SecurityStrengthFactor) Gt() Condition {
	return Cond(BindSSF, r, Gt).
		Encap(`"`).
		NoPadding(!ConditionPadding).
		setCategory(`ssf`)
}

/*
Le initializes and returns a new *Condition instance configured to
evaluate SSF as Less-Than-Or-Equal to the request confidentility factor.
*/
func (r SecurityStrengthFactor) Le() Condition {
	return Cond(BindSSF, r, Le).
		Encap(`"`).
		NoPadding(!ConditionPadding).
		setCategory(`ssf`)
}

/*
Ge initializes and returns a new *Condition instance configured to
evaluate SSF as Greater-Than-Or-Equal to the request confidentility
factor.
*/
func (r SecurityStrengthFactor) Ge() Condition {
	return Cond(BindSSF, r, Ge).
		Encap(`"`).
		NoPadding(!ConditionPadding).
		setCategory(`ssf`)
}

func (r SecurityStrengthFactor) String() string {
	if r.isZero() {
		return `0`
	}
	return sprintf("%d", int((*r.ssf.uint8))+1)
}

/*
Valid returns a boolean value indicative of whether the receiver represents
a security strength factor greater than zero (0).
*/
func (r SecurityStrengthFactor) Valid() (err error) {
	if r.value() == 0 {
		err = errorf("%T instance SSF is zero")
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

func matchAuthMethod(x string) AuthMethod {
	switch lc(x) {
	case Simple.String():
		return Simple
	case SSL.String():
		return SSL
	case SASL.String():
		return SASL
	}

	return Anonymous
}

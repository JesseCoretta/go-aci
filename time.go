package aci

/*
time.go contains temporal methods involving days and clock values for use in ACI composition.
*/

import (
	"time"
)

/*
Day constants can be shifted into an instance of DayOfWeek, allowing effective expressions such as "Sun,Tues". See the DayOfWeek.Set method.
*/
const (
	noDay Day = 0         // 0 <invalid_day>
	Sun   Day = 1 << iota // 1
	Mon                   // 2
	Tues                  // 4
	Wed                   // 8
	Thur                  // 16
	Fri                   // 32
	Sat                   // 64
)

const badDoW = `<invalid_days>`

/*
Day represents the numerical abstraction of a single day of the week, such
as Sunday (1).
*/
type Day uint8

/*
DayOfWeek contains embedded left-shifted bits that collectively represent one or
more days of the week in a "dayofweek" Bind Rule condition.
*/
type DayOfWeek struct {
	*days
}

type days uint8

/*
iterate a comma-delimited list and verify
each slice as a day of the week. return a
DayOfWeek instance alongside a Boolean
value indicative of success.
*/
func parseDoW(d string) (D DayOfWeek, err error) {
	X := split(repAll(d, ` `, ``), `,`)
	for i := 0; i < len(X); i++ {
		dw := matchStrDoW(X[i])
		if dw == noDay {
			return
		}
		D.Shift(dw)
	}
	err = D.Valid()
	return
}

func matchDoW(d any) Day {
	switch tv := d.(type) {
	case int:
		return matchIntDoW(tv)
	case string:
		return matchStrDoW(tv)
	case Day:
		return tv
	}

	return noDay
}

func matchStrDoW(d string) Day {
	switch lc(d) {
	case `sun`, `sunday`, `1`:
		return Sun
	case `mon`, `monday`, `2`:
		return Mon
	case `tues`, `tuesday`, `3`:
		return Tues
	case `wed`, `wednesday`, `4`:
		return Wed
	case `thur`, `thurs`, `thursday`, `5`:
		return Thur
	case `fri`, `friday`, `6`:
		return Fri
	case `sat`, `saturday`, `7`:
		return Sat
	}

	return noDay
}

func matchIntDoW(d int) Day {
	switch d {
	case 1:
		return Sun
	case 2:
		return Mon
	case 3:
		return Tues
	case 4:
		return Wed
	case 5:
		return Thur
	case 6:
		return Fri
	case 7:
		return Sat
	}

	return noDay
}

/*
DoW initializes, shifts and returns a new instance of DayOfWeek in one shot. This
function an alternative to separate assignment and set procedures.
*/
func DoW(x ...any) (D DayOfWeek) {
	d := DayOfWeek{new(days)}

	// assert each dow's type and analyze.
	// If deemed a valid dow, left-shift
	// into D.
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case int, string:
			if dw := matchDoW(tv); dw != noDay {
				d.Shift(dw)
			}
		case Day:
			if tv.String() != badDoW {
				d.Shift(tv)
			}
		default:
			printf("UNHANDLED %T\n", tv)
		}
	}

	D.days = d.days
	return
}

func (r DayOfWeek) Keyword() Keyword {
	return BindDoW
}

/*
Len returns the abstract integer length of the receiver, quantifying
the number of Day instances currently being expressed. For example,
if the receiver instance has its Mon and Fri Day bits enabled, this
would represent an abstract length of two (2).
*/
func (r DayOfWeek) Len() int {
	var D int
	for i := 0; i < bitSize(noDay); i++ {
		if d := Day(1 << i); r.Positive(d) {
			D++
		}
	}

	return D
}

/*
Shift shifts the first (1st) byte within the receiver instance of DayOfWeek to
include Day x, if not already present.
*/
func (r *DayOfWeek) Shift(x Day) *DayOfWeek {
	if r.IsZero() {
		r.days = new(days)
	}

	(*r.days) |= days(x)
	return r
}

/*
Positive returns a boolean value indicative of whether the receiver instance
of DayOfWeek includes Day x.
*/
func (r DayOfWeek) Positive(x Day) bool {
	if r.IsZero() {
		return false
	}
	return (*r.days)&days(x) > 0
}

/*
Unshift right-shifts the first (1st) byte within the receiver instance of DayOfWeek
to remove Day x, if present.
*/
func (r *DayOfWeek) Unshift(x Day) *DayOfWeek {
	if r.IsZero() {
		return r
	}

	(*r.days) = (*r.days) &^ days(x)
	return r
}

/*
IsZero returns a boolean value indicative of whether the receiver is nil, or unset.
*/
func (r DayOfWeek) IsZero() bool {
	if &r == nil {
		return true
	}
	return r.days == nil
}

/*
String is a stringer method that returns the string representation of the receiver
instance. At least one Day's bits should register as positive in order for a valid
string return to ensue.
*/
func (r DayOfWeek) String() string {
	if r.IsZero() {
		return badDoW
	}

	var dows []string
	for i := 0; i < bitSize(noDay); i++ {
		day := Day(1 << i)
		if r.Positive(day) {
			dows = append(dows, day.String())
		}
	}

	if len(dows) == 0 {
		return badDoW
	}

	return join(dows, `,`)
}

/*
Valid returns a boolean value indicative of whether the receiver contains one or
more valid bits representing known Day values. At least one Day must be positive
within the receiver.
*/
func (r DayOfWeek) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
		return
	}

	if r.String() == badDoW {
		err = dowBadTimeErr()
	}

	return
}

/*
Eq initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Equal-To the `dayofweek` Bind keyword
context.
*/
func (r DayOfWeek) Eq() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(BindDoW)
	b.SetOperator(Eq)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(r.Keyword().String())

	return b
}

/*
Ne initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Not-Equal-To the `dayofweek` Bind keyword
context.

Negated equality BindRule instances should be used with caution.
*/
func (r DayOfWeek) Ne() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(BindDoW)
	b.SetOperator(Ne)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(r.Keyword().String())

	return b
}

/*
BRF returns an instance of BindRuleFuncs.

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
func (r DayOfWeek) BRF() BindRuleFuncs {
	return newBindRuleFuncs(bindRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
String is a stringer method that returns a single string name value for
receiver instance of Day.
*/
func (r Day) String() (day string) {
	switch r {
	case noDay:
		day = badDoW
	case Sun:
		day = `Sun`
	case Mon:
		day = `Mon`
	case Tues:
		day = `Tues`
	case Wed:
		day = `Wed`
	case Thur:
		day = `Thur`
	case Fri:
		day = `Fri`
	case Sat:
		day = `Sat`
	}

	return
}

////////////////////////////////////////////////////////////////
// Begin Time / TimeOfDay
////////////////////////////////////////////////////////////////

/*
TimeOfDay is a [2]byte type used to represent a specific point in 24-hour time using hours and minutes (such as 1215 for 12:15 PM, or 1945 for 7:45 PM). Instances of this type contain a big endian unsigned 16-bit integer value, one that utilizes the first (1st) and second (2nd) slices. The value is used within "timeofday" Bind Rule statements.
*/
type TimeOfDay struct {
	*timeOfDay
}

/*
ToD initializes, sets and returns a new instance of TimeOfDay in one shot. This
function is an alternative to separate assignment and set procedures.
*/
func ToD(x ...any) TimeOfDay {
	return newTimeOfDay(x...)
}

func newTimeOfDay(x ...any) TimeOfDay {
	t := new(timeOfDay)
	if len(x) > 0 {
		t.set(x[0])
	}
	return TimeOfDay{t}
}

type timeOfDay [2]byte

const badToD = `<invalid_timeofday>`

/*
Timeframe is a convenience function that returns a BindRules instance for the
purpose of expressing a timeframe during which access may (or may not) be
granted. This is achieved by combining the two (2) TimeOfDay input values in
a Boolean "AND stack".

The notBefore input value defines the so-called "start" of the timeframe. It
should be chronologically earlier than notAfter. This value will be used to
craft a Greater-Than-Or-Equal (Ge) BindRule expressive statement.

The notAfter input value defines the so-called "end" of the timeframe. It
should be chronologically later than notBefore. This value will be used to
craft a Less-Than (Lt) BindRule expressive statement.
*/
func Timeframe(notBefore, notAfter TimeOfDay) (window BindRules) {
	window = And()
	window.Push(
		notBefore.Ge(), // greater than or equal
		notAfter.Lt(),  // less than
	)
	return
}

func (r TimeOfDay) Keyword() Keyword {
	return BindToD
}

/*
Eq initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Equal-To the `timeofday` Bind keyword
context.
*/
func (r TimeOfDay) Eq() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(BindToD)
	b.SetOperator(Eq)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(r.Keyword().String())

	return b
}

/*
Ne initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Not-Equal-To the `timeofday` Bind keyword
context.

Negated equality BindRule instances should be used with caution.
*/
func (r TimeOfDay) Ne() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(BindToD)
	b.SetOperator(Ne)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(r.Keyword().String())

	return b
}

/*
Lt initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Less-Than the `timeofday` Bind keyword
context.
*/
func (r TimeOfDay) Lt() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(BindToD)
	b.SetOperator(Lt)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindToD.String())

	return b
}

/*
Le initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Less-Than-Or-Equal to the `timeofday` Bind
keyword context.
*/
func (r TimeOfDay) Le() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(BindToD)
	b.SetOperator(Le)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindToD.String())

	return b
}

/*
Gt initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Greater-Than the `timeofday` Bind keyword
context.
*/
func (r TimeOfDay) Gt() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(BindToD)
	b.SetOperator(Gt)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindToD.String())

	return b
}

/*
Ge initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Greater-Than-Or-Equal to the `timeofday`
Bind keyword context.
*/
func (r TimeOfDay) Ge() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}

	var b BindRule
	b.SetKeyword(BindToD)
	b.SetOperator(Ge)
	b.SetExpression(r)

	castAsCondition(b).
		Encap(`"`).
		SetID(bindRuleID).
		NoPadding(!RulePadding).
		SetCategory(BindToD.String())

	return b
}

/*
BRF returns an instance of BindRuleFuncs.

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
func (r TimeOfDay) BRF() BindRuleFuncs {
	return newBindRuleFuncs(bindRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
		Lt: r.Lt,
		Le: r.Le,
		Gt: r.Gt,
		Ge: r.Ge,
	})
}

/*
String is a stringer method that returns the string representation of the
receiver instance.
*/
func (r TimeOfDay) String() string {
	return r.timeOfDay.string()
}

/*
string is a private stringer called by TimeOfDay.String.
*/
func (r *timeOfDay) string() string {
	if r == nil {
		return badToD
	}
	return sprintf("%04d", uint16g([]byte{(*r)[0], (*r)[1]}))
}

/*
Valid returns a boolean value indicative of whether
the receiver is believed to be in a valid state.
*/
func (r TimeOfDay) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
	}
	return
}

/*
IsZero returns a boolean value indicative of whether
the receiver is nil, or unset.
*/
func (r TimeOfDay) IsZero() bool {
	return r.timeOfDay.isZero()
}

/*
Set encodes the specified 24-hour (a.k.a.: military) time value into the receiver instance.

Valid input types are string and time.Time. The effective hour and minute values, when combined,
should ALWAYS fall within the valid clock range of 0000 up to and including 2400.  Bogus values
within said range, such as 0477, will return an error.
*/
func (r *TimeOfDay) Set(t any) *TimeOfDay {
	if r == nil {
		*r = newTimeOfDay(t)
		return r
	} else if r.timeOfDay.isZero() {
		*r = newTimeOfDay(t)
		return r
	}

	r.timeOfDay.set(t)
	return r
}

func (r *timeOfDay) isZero() bool {
	return r == nil
}

func (r *timeOfDay) set(t any) {
	if r.isZero() {
		r = new(timeOfDay)
	}
	assertToD(r, t)
}

/*
assertToD is called by timeOfDay.set for the purpose of
handling a potential clock time value for use in a Bind
Rule statement.

TODO: handle pure int w/o interpolation as binary.
*/
func assertToD(r *timeOfDay, t any) {
	switch tv := t.(type) {
	case time.Time:
		// time.Time input results in a recursive
		// run of this method.
		if tv.IsZero() {
			break
		}
		r.set(sprintf("%02d%02d", tv.Hour(), tv.Minute()))
	case string:
		// Handle discrepancy between ACI time, which ends
		// at 2400, and Golang Time, which ends at 2359.
		var offset int
		if tv == `2400` {
			tv = `2359` // so time.Parse doesn't flip
			offset = 41 // so we can use it as intended per ACI time syntax.
		}

		if _, err := time.Parse(`1504`, tv); err == nil {
			if n, err := atoi(tv); err == nil {
				x := make([]byte, 2)
				uint16p(x, uint16(n+offset))
				for i := 0; i < 2; i++ {
					(*r)[i] = x[i]
				}
			}
		}
	}
}

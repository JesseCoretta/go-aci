package aci

/*
time.go contains temporal methods involving days and clock values for use in ACI composition.
*/

import (
	"time"
)

var now func() time.Time = time.Now

/*
Day constants can be shifted into an instance of [DayOfWeek], allowing effective expressions such as [Sun],[Tues]. See the [DayOfWeek.Shift] and [DayOfWeek.Unshift] methods.
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
Day represents the numerical abstraction of a single day of the week, such as Sunday (1).
*/
type Day uint8

/*
parseDoW will iterate a comma-delimited list and verify each slice as a day of the week and return a [DayOfWeek] instance alongside a Boolean value indicative of success.
*/
func parseDoW(dow string) (d DayOfWeek, err error) {
	d = newDoW()
	X := split(repAll(dow, ` `, ``), `,`)
	for i := 0; i < len(X); i++ {
		dw := matchStrDoW(X[i])
		if dw == noDay {
			err = dowBadDayErr(X[i])
			return
		}
		d.Shift(dw)
	}
	err = d.Valid()
	return
}

func matchDoW(d any) (D Day) {
	D = noDay
	switch tv := d.(type) {
	case int:
		D = matchIntDoW(tv)
	case string:
		D = matchStrDoW(tv)
	case Day:
		D = tv
	}

	return
}

func matchStrDoW(d string) (D Day) {
	D = noDay
	switch lc(d) {
	case `sun`, `sunday`, `1`:
		D = Sun
	case `mon`, `monday`, `2`:
		D = Mon
	case `tues`, `tuesday`, `3`:
		D = Tues
	case `wed`, `wednesday`, `4`:
		D = Wed
	case `thur`, `thurs`, `thursday`, `5`:
		D = Thur
	case `fri`, `friday`, `6`:
		D = Fri
	case `sat`, `saturday`, `7`:
		D = Sat
	}

	return
}

func matchIntDoW(d int) (D Day) {
	D = noDay
	switch d {
	case 1:
		D = Sun
	case 2:
		D = Mon
	case 3:
		D = Tues
	case 4:
		D = Wed
	case 5:
		D = Thur
	case 6:
		D = Fri
	case 7:
		D = Sat
	}

	return
}

/*
DoW initializes, shifts and returns a new instance of [DayOfWeek] in one shot. This function an alternative to separate assignment and set procedures.
*/
func DoW(x ...any) (d DayOfWeek) {
	d = newDoW()
	d.Shift(x...)

	return
}

/*
Keyword returns the [BindToD] [BindKeyword].
*/
func (r DayOfWeek) Keyword() Keyword {
	return BindDoW
}

/*
Len returns the abstract integer length of the receiver, quantifying the number of [Day] instances currently being expressed.

For example, if the receiver instance has its [Mon] and [Fri] [Day] bits enabled, this would represent an abstract length of two (2).
*/
func (r DayOfWeek) Len() int {
	var D int
	for i := 0; i < r.cast().Size(); i++ {
		if d := Day(1 << i); r.cast().Positive(d) {
			D++
		}
	}

	return D
}

/*
Weekdays is a convenient prefabricator function that returns an instance of [BindRule] automatically assembled to express a sequence of weekdays. The sequence "[Mon] through [Fri]" can also be expressed via the bit-shifted value of sixty-two (62). See the [Day] constants for the specific numerals used for summation in this manner.

Supplying an invalid or nonapplicable [ComparisonOperator] to this method shall return a bogus [BindRule] instance.
*/
func Weekdays(cop any) (b BindRule) {
	if c, meth := DoW(Mon, Tues, Wed, Thur, Fri).BRM().index(cop); c.Valid() == nil {
		b = meth()
	}
	return
}

/*
Weekend is a convenient prefabricator function that returns an instance of [BindRule] automatically assembled to express a sequence of [Sun] and [Sat] [Day] instances. This sequence can also be expressed via the bit-shifted value of sixty-five (65). See the [Day] constants for the specific numerals used for summation in this manner.

Supplying an invalid or nonapplicable [ComparisonOperator] to this method shall return a bogus [BindRule] instance.
*/
func Weekend(cop any) (b BindRule) {
	if c, meth := DoW(Sun, Sat).BRM().index(cop); c.Valid() == nil {
		b = meth()
	}
	return
}

/*
Shift wraps [shifty.BitValue.Shift] method to allow for bit-shifting of the receiver (r) instance using various representations of any number of days (string, int or [Day]).
*/
func (r *DayOfWeek) Shift(x ...any) DayOfWeek {
	// initialize receiver r if zero.
	if r.IsZero() {
		*r = newDoW()
	}

        // assert each dow's type and analyze.
        // If deemed a valid dow, left-shift
        // into d.
        for i := 0; i < len(x); i++ {
                switch tv := x[i].(type) {
                case int, string:
                        if dw := matchDoW(tv); dw != noDay {
				r.cast().Shift(dw)
                        }
                case Day:
			r.cast().Shift(tv)
                }
        }

	return *r
}

/*
Positive wraps the [shifty.BitValue.Positive] method.
*/
func (r DayOfWeek) Positive(x Day) (posi bool) {
	if !r.IsZero() {
		posi = r.cast().Positive(x)
	}
	return
}

/*
Unshift wraps [shifty.BitValue.Unshift] method to allow for bit-unshifting of the receiver (r) instance using various representations of any number of days (string, int or [Day]).
*/
func (r *DayOfWeek) Unshift(x ...any) DayOfWeek {
	// can't unshift from nothing
	if r.IsZero() {
		return *r
	}

        // assert each dow's type and analyze.
        // If deemed a valid dow, right-shift
        // out of d.
        for i := 0; i < len(x); i++ {
                switch tv := x[i].(type) {
                case int, string:
                        if dw := matchDoW(tv); dw != noDay {
                                r.cast().Unshift(dw)
                        }
                case Day:
                        r.cast().Unshift(tv)
                }
        }

	return *r
}

/*
IsZero returns a Boolean value indicative of whether the receiver is nil, or unset.
*/
func (r DayOfWeek) IsZero() bool {
	return r.cast().Kind() == 0x0
}

/*
String is a stringer method that returns the string representation of the receiver instance. At least one [Day] should register as positive in order for a valid string return to ensue.
*/
func (r DayOfWeek) String() (s string) {
	s = badDoW

	var dows []string
	for i := 0; i < r.cast().Size(); i++ {
		day := Day(1 << i)
		if r.cast().Positive(day) {
			dows = append(dows, day.String())
		}
	}

	if len(dows) > 0 {
		s = join(dows, `,`)
	}

	return
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r DayOfWeek) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Valid returns a Boolean value indicative of whether the receiver contains one or more valid bits representing known [Day] values.

At least one [Day] must be positive within the receiver.
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
Eq initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Equal-To the [BindDoW] [BindKeyword] context.
*/
func (r DayOfWeek) Eq() (b BindRule) {
	if err := r.Valid(); err == nil {
		b = BR(BindDoW, Eq, r)
	}
	return
}

/*
Ne initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Not-Equal-To the [BindDoW] [BindKeyword] context.

Negated equality [BindRule] instances should be used with caution.
*/
func (r DayOfWeek) Ne() (b BindRule) {
	if err := r.Valid(); err == nil {
		b = BR(BindDoW, Ne, r)
	}
	return
}

/*
BRM returns an instance of [BindRuleMethods].

Each of the return instance's key values represent a single instance of the [ComparisonOperator] type that is allowed for use in the creation of [BindRule] instances which bear the receiver instance as an expression value. The value for each key is the actual [BindRuleMethod] instance for OPTIONAL use in the creation of a [BindRule] instance.

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus [BindRule] instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly populated and prepared for such activity.
*/
func (r DayOfWeek) BRM() BindRuleMethods {
	return newBindRuleMethods(bindRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
String is a stringer method that returns a single string name value for receiver instance of [Day].
*/
func (r Day) String() (day string) {
	day = badDoW
	switch r {
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

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r Day) Compare(x any) bool {
	return compareHashInstance(r, x)
}

////////////////////////////////////////////////////////////////
// Begin Time / TimeOfDay
////////////////////////////////////////////////////////////////

/*
TimeOfDay is a [2]byte type used to represent a specific point in 24-hour time using hours and minutes (such as 1215 for 12:15 PM, or 1945 for 7:45 PM). Instances of this type contain a big endian unsigned 16-bit integer value, one that utilizes the first (1st) and second (2nd) slices. The value is used within [BindToD]-based [BindRule] statements.
*/
type TimeOfDay struct {
	*timeOfDay
}

/*
ToD initializes, sets and returns a new instance of [TimeOfDay] in one shot. This function is an alternative to separate assignment and set procedures.
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
Timeframe is a convenience function that returns a [BindRules] instance for the purpose of expressing a timeframe during which access may (or may not) be granted. This is achieved by combining the two (2) [TimeOfDay] input values in a Boolean "AND stack".

The notBefore input value defines the so-called "start" of the timeframe. It should be chronologically earlier than notAfter. This value will be used to craft a Greater-Than-Or-Equal (Ge) [BindRule] expressive statement.

The notAfter input value defines the so-called "end" of the timeframe. It should be chronologically later than notBefore. This value will be used to craft a Less-Than (Lt) [BindRule] expressive statement.
*/
func Timeframe(notBefore, notAfter TimeOfDay) (window BindRules) {
	window = And()
	window.Push(
		notBefore.Ge(), // greater than or equal
		notAfter.Lt(),  // less than
	)
	return
}

/*
Keyword wraps the [stackage.Condition.Keyword] method and resolves the raw value into a [BindKeyword]. Failure to do so will return a bogus [Keyword].
*/
func (r TimeOfDay) Keyword() Keyword {
	return BindToD
}

/*
Eq initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Equal-To the [BindToD] [BindKeyword] context.
*/
func (r TimeOfDay) Eq() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(BindToD, Eq, r)
}

/*
Ne initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Not-Equal-To the [BindToD] [BindKeyword] context.

Negated equality [BindRule] instances should be used with caution.
*/
func (r TimeOfDay) Ne() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(BindToD, Ne, r)
}

/*
Lt initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Less-Than the [BindToD] [BindKeyword] context.
*/
func (r TimeOfDay) Lt() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(BindToD, Lt, r)
}

/*
Le initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Less-Than-Or-Equal to the [BindToD] [BindKeyword] context.
*/
func (r TimeOfDay) Le() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(BindToD, Le, r)
}

/*
Gt initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Greater-Than the [BindToD] [BindKeyword] context.
*/
func (r TimeOfDay) Gt() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(BindToD, Gt, r)
}

/*
Ge initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Greater-Than-Or-Equal to the [BindToD] [BindKeyword] context.
*/
func (r TimeOfDay) Ge() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(BindToD, Ge, r)
}

/*
BRM returns an instance of [BindRuleMethods].

Each of the return instance's key values represent a single instance of the [ComparisonOperator] type that is allowed for use in the creation of [BindRule] instances which bear the receiver instance as an expression value. The value for each key is the actual [BindRuleMethod] instance for OPTIONAL use in the creation of a [BindRule] instance.

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus [BindRule] instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly populated and prepared for such activity.
*/
func (r TimeOfDay) BRM() BindRuleMethods {
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
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r TimeOfDay) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
String is a stringer method that returns the string representation of the receiver instance.
*/
func (r TimeOfDay) String() string {
	return r.timeOfDay.string()
}

/*
string is a private stringer called by [TimeOfDay]'s stringer.
*/
func (r *timeOfDay) string() (s string) {
	s = badToD
	if r != nil {
		s = sprintf("%04d", uint16g([]byte{(*r)[0], (*r)[1]}))
	}
	return
}

/*
Valid returns a Boolean value indicative of whether the receiver is believed to be in a valid state.
*/
func (r TimeOfDay) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
	}
	return
}

/*
IsZero returns a Boolean value indicative of whether the receiver is nil, or unset.
*/
func (r TimeOfDay) IsZero() bool {
	return r.timeOfDay.isZero()
}

/*
Set encodes the specified 24-hour (a.k.a.: military) time value into the receiver instance.

Valid input types are string and [time.Time]. The effective hour and minute values, when combined, should ALWAYS fall within the valid clock range of 0000 up to and including 2400.  Bogus values within said range, such as 0477, will return an error.
*/
func (r *TimeOfDay) Set(t any) TimeOfDay {
	*r = newTimeOfDay(t)
	return *r
}

func (r *timeOfDay) isZero() bool {
	return r == nil
}

func (r *timeOfDay) set(t any) {
	assertToD(r, t)
}

/*
assertToD is called by timeOfDay.set for the purpose of handling a potential clock time value for use in a [BindRule] statement.
*/
func assertToD(r *timeOfDay, t any) {
	switch tv := t.(type) {
	case time.Time:
		// time.Time input results in a recursive
		// run of this method.
		if !tv.IsZero() {
			r.set(sprintf("%02d%02d", tv.Hour(), tv.Minute()))
		}
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

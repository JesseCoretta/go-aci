package aci

import (
	"fmt"
	"testing"
	"time"
)

func TestToD(t *testing.T) {
	// results we expect
	//	TIME	VALID
	times := map[string]bool{
		`3414`: false, // 0
		`2155`: true,  // 1
		`2995`: false, // 2
		`6668`: false, // 3
		`5159`: false, // 4
		`0540`: true,  // 5
		`5197`: false, // 6
		`2400`: true,  // 7	// go time.Time will skew this by -1
		`0159`: true,  // 8
		`5601`: false, // 9
		`0101`: true,  // 10
		`6674`: false, // 11
		`9392`: false, // 12
		`0359`: true,  // 13
		`7808`: false, // 14
		`8671`: false, // 15
		`1000`: true,  // 16
		`1349`: true,  // 17
		`6746`: false, // 18
		`0505`: true,  // 19
		`1365`: false, // 20
		`1338`: true,  // 21
		`0881`: false, // 22
		`0000`: true,  // 23
		`0560`: false, // 24
		`0003`: true,  // 25
		`5544`: false, // 26
		`8582`: false, // 27
		`4139`: false, // 28
		`1206`: true,  // 29
	}

	for thyme, want := range times {
		var (
			err  error
			typ  string = BindToD.String()
			cops map[ComparisonOperator]func() BindRule
		)

		got := ToD(thyme)
		if got.String() != thyme && want {
			err = unexpectedStringResult(typ, thyme, got.String())
			t.Errorf("%s failed: %v", t.Name(), err)
		}

		if err = handleToDGoTime(thyme, typ, want); err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
		}

		times[`2400`] = false

		// tod qualifies for all comparison operators
		// due to its numerical nature.
		cops = map[ComparisonOperator]func() BindRule{
			Eq: got.Eq,
			Ne: got.Ne,
			Lt: got.Lt,
			Le: got.Le,
			Gt: got.Gt,
			Ge: got.Ge,
		}

		// try every comparison operator supported in
		// this context ...
		for c := 1; c < len(cops)+1; c++ {
			cop := ComparisonOperator(c)
			wcop := sprintf("%s %s %q", got.Keyword(), cop, got)

			// create bindrule B using comparison
			// operator (cop).
			if B := cops[cop](); B.String() != wcop {
				err = unexpectedStringResult(typ, wcop, B.String())
			}

			if err != nil && want {
				t.Errorf("%s failed: %v", t.Name(), err)
			}
		}
	}
}

func handleToDGoTime(thyme, typ string, want bool) (err error) {
	// convert thyme into an bonafide time.Time
	// instance, and retry the operation.
	var T time.Time
	if T, err = time.Parse(`1504`, thyme); err != nil && want {
		err = generalErr(typ, err)
	}

	if got := ToD(T); got.String() != thyme && want {
		err = unexpectedStringResult(typ, thyme, got.String())
	}

	if !(err != nil && want && thyme != `2400`) {
		err = nil
	}

	return
}

func TestDoW(t *testing.T) {
	// length results we expect
	//	LEN	INDEX
	lens := []int{
		5, // 0
		4, // 1
		3, // 2
		4, // 3
		3, // 4
		6, // 5
		2, // 6
		4, // 7
		1, // 8
		2, // 9
		3, // 10
		2, // 11
		1, // 12
	}

	// string results we expect
	//	STRING			INDEX
	strings := []string{
		`Mon,Tues,Wed,Thur,Fri`,    // 0
		`Mon,Tues,Wed,Fri`,         // 1
		`Sun,Mon,Sat`,              // 2
		`Mon,Wed,Thur,Fri`,         // 3
		`Tues,Wed,Fri`,             // 4
		`Sun,Mon,Tues,Wed,Fri,Sat`, // 5
		`Mon,Wed`,                  // 6
		`Sun,Thur,Fri,Sat`,         // 7
		`Mon`,                      // 8
		`Sun,Sat`,                  // 9
		`Sun,Tues,Fri`,             // 10
		`Tues,Wed`,                 // 11
		`Tues`,                     // 12
	}

	// iterate a series of test sequences, each slice
	// representing a single DayOfWeek instance.
	//	TEST VALUES				INDEX
	for idx, _days := range []any{
		[]any{Mon, Tues, Wed, Thur, Fri},       // 0
		[]any{Mon, Tues, Wed, Fri},             // 1
		[]any{`sun`, `mOn`, 7},                 // 2
		[]any{Mon, Wed, Thur, Fri},             // 3
		[]any{Tues, Wed, Fri},                  // 4
		[]any{`1`, `2`, 3, `wEd`, `6`, `7`},    // 5
		[]any{Mon, Wed, `fryDay`},              // 6
		[]any{`thur`, `fRi`, Sun, `SAT`},       // 7
		[]any{Mon},                             // 8
		[]any{Sun, Sat},                        // 9
		[]any{1, Tues, 6},                      // 10
		[]any{`tues`, `WEDNESDAY`},             // 11
		[]any{Tues, `tues`, `TUESday`, 3, `3`}, // 12
	} {
		var err error
		var typ string = BindDoW.String()
		var got DayOfWeek

		// try to parse and marshal the above 'any'
		// slices. The DoW function should be able
		// to support a host of different "day of
		// the week" expressions ...

		if got = DoW(_days.([]any)...); got.Len() != lens[idx] {
			// the "number" of days are different
			// than what we expected ...
			err = unexpectedValueCountErr(typ, lens[idx], got.Len())

		} else if got.String() != strings[idx] {
			// the string representation of the instance
			// does not match what we expected ...
			err = unexpectedStringResult(typ, strings[idx], got.String())
		}

		if err != nil {
			t.Errorf("%s failed [slice:%d]: %v", t.Name(), idx, err)
		}

		wantEq := sprintf("%s = %q", got.Keyword(), got)
		wantNe := sprintf("%s != %q", got.Keyword(), got)

		if B := got.Eq(); B.String() != wantEq {
			err = unexpectedStringResult(typ, wantEq, B.String())
		} else if B = got.Ne(); B.String() != wantNe {
			err = unexpectedStringResult(typ, wantNe, B.String())
		}

		if err != nil {
			t.Errorf("%s failed [slice:%d]: %v", t.Name(), idx, err)
		}
	}
}

func ExampleDoW() {
	fmt.Printf("%s", DoW(Thur, Sat))
	// Output: Thur,Sat
}

func ExampleTime() {
	fmt.Printf("%s", ToD(`2301`))
	// Output: 2301
}

func ExampleTime_setLater() {
	t := ToD()
	t.Set(`2301`)
	fmt.Printf("%s", t)
	// Output: 2301
}

func ExampleTimeOfDay_setLater() {
	var t TimeOfDay
	t.Set(`2301`)
	fmt.Printf("%s", t)
	// Output: 2301
}

/*
This example demonstrates the creation of a timeframe BindRules instance using
the convenient Timeframe package-level function.

The return value is a Boolean AND stack object containing the start (notBefore)
and end (notAfter) TimeOfDay instances.
*/
func ExampleTimeframe() {
	notBefore := ToD(`1730`)
	notAfter := ToD(`2400`)
	tfr := Timeframe(notBefore, notAfter).Paren(true)
	fmt.Printf("%s", tfr)
	// Output: ( timeofday >= "1730" AND timeofday < "2400" )
}

func TestParseDoW(t *testing.T) {
	failOK := func(x int) bool {
		for _, val := range []int{
			2,
			6,
			7,
		} {
			if x == val {
				return true
			}
		}
		return false
	}

	for idx, d := range []string{
		`Sunday`,
		`Monday,Saturday`,
		`humpday`,
		`sUN,mon,tues,Wed,THURSDAY,friDAY,SATurDAy`,
		`tuesday`,
		`thursday,monday`,
		``,
		`_+))`,
	} {
		if dow, err := parseDoW(d); err != nil {
			if !failOK(idx) {
				t.Errorf("%s failed to parse %T '%s': %v",
					t.Name(), d, d, err)
			}

		} else if dow.String() == badDoW {
			if !failOK(idx) {
				t.Errorf("%s failed: want '%T', got '%s'",
					t.Name(), d, dow)
			}
		}
	}
}

func TestMatchDoW(t *testing.T) {
	failOK := func(x int) bool {
		for _, val := range []int{
			3,
			8,
			11,
		} {
			if x == val {
				return true
			}
		}
		return false
	}

	for idx, d := range []any{
		1,
		`sunDAY`,
		`1`,
		8,
		Sat,
		`thur`,
		`thurs`,
		`3`,
		-1,
		Tues,
		6,
		`SQUAtcOBbl3r`,
		5,
		7,
		4,
		3,
		2,
		`Monday`,
		Mon,
	} {
		if dow := matchDoW(d); dow == noDay {
			if !failOK(idx) {
				t.Errorf("%s failed [match %d]: want '%T', got '%s'",
					t.Name(), idx, Day(0), dow)
			}

		} else if D := DoW(d); D.IsZero() {
			if !failOK(idx) {
				t.Errorf("%s failed [make %d]: want '%T', got '%s'",
					t.Name(), idx, Day(0), D)
			}
		}
	}
}

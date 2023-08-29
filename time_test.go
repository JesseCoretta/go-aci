package aci

import (
	"fmt"
	"testing"
)

func TestDoW(t *testing.T) {
	want := `Thur,Sat`
	got := DoW(Thur, Sat).String()
	if want != got {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

func ExampleDoW() {
	fmt.Printf("%s", DoW(Thur, Sat))
	// Output: Thur,Sat
}

func TestToD(t *testing.T) {
	want := `2301`
	got := ToD(`2301`)
	if got.String() != want {
		t.Errorf("1 %s failed: want '%s', got '%s'", t.Name(), want, got)
	}

	// reset our instance without
	// allocating a new one
	got.Set(`0021`)
	want = `0021`
	if got.String() != want {
		t.Errorf("2 %s failed: want '%s', got '%s'", t.Name(), want, got)
	}
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
	for _, d := range []string{
		`Sunday`,
		`Monday,Saturday`,
		`sUN,mon,tues,Wed,THURSDAY,friDAY,SATurDAy`,
		`tuesday`,
		`thursday,monday`,
	} {
		if dow, err := parseDoW(d); err != nil {
			t.Errorf("%s failed to parse %T '%s': %v",
				t.Name(), d, d, err)
		} else if dow.String() == badDoW {
			t.Errorf("%s failed: want '%T', got '%s'",
				t.Name(), d, dow)
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
		`Monday`,
		Mon,
	} {
		if dow := matchDoW(d); dow == noDay {
			if !failOK(idx) {
				t.Errorf("%s failed [test idx %d]: want '%T', got '%s'",
					t.Name(), idx, Day(0), dow)
			}
		}
	}
}

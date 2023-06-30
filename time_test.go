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

func ExampleTimeOfDay_timeFrameCondition() {
	// Make an ANDed stack, enable
	// parenthetical encapsulation,
	// then push two (2) Conditions
	timeframe := And().Paren().Push(
		ToD(`1730`).Ge(),	// greater than or equal, i.e.: not before
		ToD(`2400`).Lt(),	// less than, i.e.: not after
	)
	fmt.Printf("%s", timeframe)
	// Output: ( timeofday >= "1730" AND timeofday < "2400" )
}

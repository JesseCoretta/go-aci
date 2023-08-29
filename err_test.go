package aci

import (
	"testing"
)

func TestErrorf(t *testing.T) {
	var results []bool = []bool{true, true, false}
	for idx, msg := range []any{
		`this is an error`,
		errorf(`this is also an error`),
		nil, // but this is not.
	} {
		err := errorf(msg)
		got := err != nil
		want := results[idx]
		if want != got {
			t.Errorf("%s failed: unexpected errorf result; want '%t', got '%t'",
				t.Name(), want, got)
		}
	}
}

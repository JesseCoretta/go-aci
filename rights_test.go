package aci

import (
	"fmt"
	"testing"
)

func TestAllow(t *testing.T) {
	G := Allow(ReadAccess, CompareAccess)
	want := `allow(read,compare)`
	got := G.String()
	if want != got {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

func TestAllow_all(t *testing.T) {
	G := Allow(ReadAccess,
		CompareAccess,
		SearchAccess,
		ImportAccess,
		ExportAccess,
		SelfWriteAccess,
		DeleteAccess,
		AddAccess,
		WriteAccess)
	want := `allow(all)`
	got := G.String()
	if want != got {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

/*
This example demonstrates the granting (allowance) of read, compare and search access.
*/
func ExampleAllow() {
	// allow read, compare and search privileges ...
	G := Allow(ReadAccess, CompareAccess, SearchAccess)

	// order is always fixed for string representation
	// regardless of order-of-input ...
	fmt.Printf("%s", G)
	// Output: allow(read,search,compare)
}

func TestDeny(t *testing.T) {
	G := Deny(AllAccess)
	want := `deny(all)`
	got := G.String()
	if want != got {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

/*
This example demonstrates the withholding (denial) of all privileges.
*/
func ExampleDeny() {
	// allow read, compare and search privileges ...
	D := Deny(AllAccess)

	// order is always fixed for string representation
	// regardless of order-of-input ...
	fmt.Printf("%s", D)
	// Output: deny(all)
}

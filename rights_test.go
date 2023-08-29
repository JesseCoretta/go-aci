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
This example demonstrates the withholding (denial) of all privileges except proxy.
*/
func ExampleDeny() {
	// deny everything (this does not include proxy privilege)
	D := Deny(AllAccess)
	fmt.Printf("%s", D)
	// Output: deny(all)
}

func TestRights_lrShift(t *testing.T) {
	var p Permission = Allow(NoAccess)
	if !p.Positive(0) || !p.Positive(`none`) {
		t.Errorf("%s failed: cannot identify 'none' permission", t.Name())
	}

	// three iterations, one per supported
	// Right type
	for i := 0; i < 3; i++ {

		// iterate each of the rights in the
		// rights/names map
		for k, v := range rightsMap {

			if k == 0 {
				continue
			}

			term, typ := testGetRightsTermType(i, k, v)

			shifters := map[int]func(any) Permission{
				0: p.Shift,
				1: p.Unshift,
			}

			for j := 0; j < len(shifters); j++ {
				mode, phase := testGetRightsPhase(j)
				if shifters[j](term); p.Positive(term) != phase {
					t.Errorf("%s failed: %T %s %s failed [key:%d; term:%v] (value:%v)",
						t.Name(), p, typ, mode, k, term, p)
				}
			}
		}
	}
}

func testGetRightsPhase(j int) (mode string, phase bool) {
	mode = `shift`
	if phase = (0 == j); !phase {
		mode = `un` + mode
	}

	return
}

func testGetRightsTermType(i int, k Right, v string) (term any, typ string) {
	term = k // default
	switch i {
	case 1:
		term = v // string name (e.g.: read)
	case 2:
		term = Right(k) // Right
	}
	typ = sprintf("%T", term) // label for err

	return
}

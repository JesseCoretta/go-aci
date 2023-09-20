package aci

import (
	"fmt"
	"testing"
)

func TestAllow(t *testing.T) {
	G := Allow(ReadAccess, CompareAccess)
	t.Logf("WHAT THE FUCK %d\n", G.permission.rights.cast().Int())

	want := `allow(read,compare)`
	got := G.String()
	if want != got {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
		return
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
		return
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
		return
	}
}

func ExampleRight_String() {
	// iterate all of the known Right definitions
	// defined as constants in this package.
	for idx, privilege := range []Right{
		NoAccess,
		ReadAccess,
		WriteAccess,
		AddAccess,
		DeleteAccess,
		SearchAccess,
		CompareAccess,
		SelfWriteAccess,
		ProxyAccess,
		ImportAccess,
		ExportAccess,
		AllAccess, // does NOT include proxy access !
	} {
		fmt.Printf("Privilege %02d/%d: %s (bit:%d)\n", idx+1, 12, privilege, int(privilege))
	}
	// Output:
	// Privilege 01/12: none (bit:0)
	// Privilege 02/12: read (bit:1)
	// Privilege 03/12: write (bit:2)
	// Privilege 04/12: add (bit:4)
	// Privilege 05/12: delete (bit:8)
	// Privilege 06/12: search (bit:16)
	// Privilege 07/12: compare (bit:32)
	// Privilege 08/12: selfwrite (bit:64)
	// Privilege 09/12: proxy (bit:128)
	// Privilege 10/12: import (bit:256)
	// Privilege 11/12: export (bit:512)
	// Privilege 12/12: all (bit:895)
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

func ExamplePermission_IsZero() {
	var priv Permission
	fmt.Printf("Privileges are undefined: %t", priv.IsZero())
	// Output: Privileges are undefined: true
}

func ExamplePermission_Valid() {
	var priv Permission
	fmt.Printf("%T is ready for use: %t", priv, priv.Valid() == nil)
	// Output: aci.Permission is ready for use: false
}

func ExamplePermission_Disposition() {
	priv := Allow(`read`, `write`, `compare`, `selfwrite`)
	fmt.Printf("%s", priv.Disposition())
	// Output: allow
}

func ExamplePermission_String() {
	priv := Allow(`read`, `write`, `compare`, `selfwrite`)
	fmt.Printf("%s", priv)
	// Output: allow(read,write,compare,selfwrite)
}

func ExamplePermission_Shift() {
	var priv Permission = Allow() // you MUST initialize Permission explicitly using Allow or Deny funcs
	priv.Shift(ReadAccess, ProxyAccess)
	fmt.Printf("Allows proxy: %t", priv.Positive(`proxy`))
	// Output: Allows proxy: true
}

func ExamplePermission_Len() {
	var priv Permission = Deny() // you MUST initialize Permission explicitly using Allow or Deny funcs
	priv.Shift(ReadAccess, WriteAccess, CompareAccess, SearchAccess, ProxyAccess)
	fmt.Printf("Number of privileges denied: %d", priv.Len())
	// Output: Number of privileges denied: 5
}

func ExamplePermission_Positive() {
	var priv Permission = Deny(`read`, `write`, `proxy`, `search`) // you MUST initialize Permission explicitly using Allow or Deny funcs
	fmt.Printf("Forbids read access: %t", priv.Positive(`read`))
	// Output: Forbids read access: true
}

func ExamplePermission_Unshift() {
	var priv Permission = Deny() // you MUST initialize Permission explicitly using Allow or Deny funcs
	priv.Shift(ReadAccess, WriteAccess, CompareAccess, SearchAccess, ProxyAccess)

	priv.Unshift(CompareAccess) // remove the negated compare privilege

	fmt.Printf("Forbids compare: %t", priv.Positive(`compare`))
	// Output: Forbids compare: false
}

func TestRights_bogus(t *testing.T) {
	var p Permission
	if err := p.Valid(); err == nil {
		t.Errorf("%s failed: invalid %T returned no validity error",
			t.Name(), p)
		return
	}

	if p.String() != badPerm {
		t.Errorf("%s failed: invalid %T returned no bogus string warning",
			t.Name(), p)
		return
	}

	p.Unshift(`all`)
	p.Shift(-1985)       //underflow
	p.Shift(45378297659) //overflow
	if !p.IsZero() {
		t.Errorf("%s failed: overflow or underflow shift value accepted for %T",
			t.Name(), p)
		return
	}

	p.Unshift(-5)     //underflow
	p.Unshift(134559) //overflow
	if !p.IsZero() {
		t.Errorf("%s failed: overflow or underflow unshift value accepted for %T",
			t.Name(), p)
		return
	}

}

func TestRights_lrShift(t *testing.T) {
	var p Permission = Allow(NoAccess)
	if !p.Positive(0) || !p.Positive(`none`) || !p.positive(NoAccess) {
		t.Errorf("%s failed: cannot identify 'none' permission", t.Name())
		return
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

			shifters := map[int]func(...any) Permission{
				0: p.Shift,
				1: p.Unshift,
			}

			for j := 0; j < len(shifters); j++ {
				mode, phase := testGetRightsPhase(j)
				if shifters[j](term); p.Positive(term) != phase {
					t.Errorf("%s failed: %T %s %s failed [key:%d; term:%v] (value:%v)",
						t.Name(), p, typ, mode, k, term, p)
					return
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

func TestPermission_codecov(t *testing.T) {
	if ReadAccess.Compare(WriteAccess) {
		t.Errorf("%s failed: %s wrongly equal to %s",
			t.Name(), ReadAccess, WriteAccess)
		return
	}

	var p Permission
	_ = p.Len()
	_ = p.Valid()
	_ = p.Disposition()
	_ = p.Shift()
	_ = p.Shift(nil)
	_ = p.Positive(nil)
	_ = p.Shift(4)
	_ = p.Shift(4547887935)
	_ = p.Shift(-45478879)
	_ = p.Unshift()
	_ = p.Unshift(nil)
	_ = p.Unshift(4)
	_ = p.Unshift(4547887935)
	_ = p.Unshift(-45478879)
	_ = p.Parse(`alow(red,rite)`)
	_ = p.Disposition()
	_ = p.Positive(ProxyAccess)
	p = Allow(ReadAccess, WriteAccess)
	d := Deny(ReadAccess, WriteAccess)
	_ = p.Unshift()
	_ = p.Unshift(nil)
	_ = p.Unshift(4)
	_ = p.Shift(4547887935)
	_ = p.Positive(4547887935)
	_ = p.Shift(-45478879)
	_ = p.Unshift(4547887935)
	_ = p.Unshift(-45478879)
	_ = d.Unshift()
	_ = d.Unshift(nil)
	_ = d.Unshift(4)
	_ = d.Shift(4547887935)
	_ = d.Positive(4547887935)
	_ = d.Shift(-45478879)
	_ = d.Unshift(4547887935)
	if p.Compare(d) {
		t.Errorf("%s failed: %s wrongly equal to %s",
			t.Name(), p, d)
		return
	}

	p.permission = new(permission)
	_ = p.Valid()
}

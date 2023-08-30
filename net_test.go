package aci

import (
	"fmt"
	"testing"
)

func TestFQDN(t *testing.T) {
	var f FQDN = DNS()
	var typ string = f.Keyword().String()

	if f.len() != 0 {
		t.Errorf("%s failed: unexpected %T length: want '%d', got '%d'",
			t.Name(), f, 0, f.len())
	}

	if err := f.Valid(); err == nil {
		t.Errorf("%s failed: empty %T deemed valid", t.Name(), f)
	}

	f.Set(`www`, `example`, `com`)

	want := `www.example.com`
	got := f.String()

	if want != got {
		t.Errorf("%s failed; want '%s', got '%s'", t.Name(), want, got)
	}

	absurd := `eeeeeeeeeeeeeeeeeeeeeeeee#eee^eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeexample`

	if validLabel(absurd) {
		t.Errorf("%s failed: bogus %T label accepted as valid (%s)",
			t.Name(), absurd, absurd)
	}

	var F FQDN
	if F.String() != badFQDN {
		t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
			t.Name(), badFQDN, F)
	}

	F.Set(`www`).Set(`$&^#*(`).Set(absurd).Set(`example`).Set(``).Set(`com`)
	if llen := F.Len(); llen != 3 {
		t.Errorf("%s failed; want '%d', got '%d'", t.Name(), 3, llen)
	}

	// FQDN qualifies for equality and negated equality
	// comparison operators.
	cops := map[ComparisonOperator]func() BindRule{
		Eq: F.Eq,
		Ne: F.Ne,
	}

	// try every comparison operator supported in
	// this context ...
	for c := 1; c < len(cops)+1; c++ {
		cop := ComparisonOperator(c)
		wcop := sprintf("%s %s %q", F.Keyword(), cop, got)

		// create bindrule B using comparison
		// operator (cop).
		if B := cops[cop](); B.String() != wcop {
			err := unexpectedStringResult(typ, wcop, B.String())
			t.Errorf("%s failed [%s rule]: %v", t.Name(), typ, err)
		}

	}
}

func TestDNS_alternativeFQDN(t *testing.T) {
	want := `www.example.com`
	got := DNS(`www.example.com`)

	if want != got.String() {
		t.Errorf("%s failed; want '%s', got '%s'", t.Name(), want, got)
	}
}

func TestIPAddr_C(t *testing.T) {
	var i IPAddr
	if !i.IsZero() {
		t.Errorf("%s failed: non-zero %T instance", t.Name(), i)
	}

	if got := i.String(); got != badAddr {
		t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
			t.Name(), badAddr, got)
	}

	i.Set(`192.168.0`)
	i.Set(`12.3.45.*`)
	i.Set(`12.3.45.*`) // duplicate
	i.Set(`10.0.0.0/8`)

	if lens := i.Len(); lens != 3 {
		t.Errorf("%s failed: bad %T length; want '%d', got '%d'", t.Name(), i, 3, lens)
	}

	if cond := i.Ne(); cond.IsZero() {
		t.Errorf("%s failed: nil %T instance!", t.Name(), cond)
	}
}

func ExampleFQDN_Eq() {
	var f FQDN = DNS()

	// Can be set incrementally ...
	f.Set(`www`)
	f.Set(`example`)
	f.Set(`com`)

	// OR as a whole value ...
	// f.Set(`www.example.com`)

	fmt.Printf("%s", f.Eq())
	// Output: dns = "www.example.com"
}

func ExampleFQDN_Eq_oneShot() {
	fmt.Printf("%s", DNS(`www.example.com`).Eq())
	// Output: dns = "www.example.com"
}

func ExampleIPAddr_Ne() {
	var i IPAddr
	i.Set(`192.168.0`).Set(`12.3.45.*`).Set(`10.0.0.0/8`)
	neq := i.Ne()
	neq.Paren()
	fmt.Printf("%s", neq)
	// Output: ( ip != "192.168.0,12.3.45.*,10.0.0.0/8" )
}

func ExampleIPAddr_Eq_oneShot() {
	fmt.Printf("%s", IP(`192.168.0`, `12.3.45.*`, `10.0.0.0/8`).Eq())
	// Output: ip = "192.168.0,12.3.45.*,10.0.0.0/8"
}

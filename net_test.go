package aci

import (
	"fmt"
	"testing"
)

func TestFQDN(t *testing.T) {
	var f FQDN = DNS()
	_ = f.Len()
	_ = f.Keyword()
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

	// try every comparison operator supported in
	// this context ...
	brf := F.BRF()
	for i := 0; i < brf.Len(); i++ {
		cop, meth := brf.Index(i + 1)
		wcop := sprintf("( %s %s \"www.example.com\" )", f.Keyword(), cop)
		if T := meth(); T.Paren().String() != wcop {
			err := unexpectedStringResult(F.String(), wcop, T.String())
			t.Errorf("%s [%s] multival failed [%s rule]; %s, %s: %v",
				t.Name(), F.Keyword(), cop.Context(), cop.Description(), typ, err)
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

func TestIPAddr_BRF(t *testing.T) {
	var i IPAddr
	_ = i.Len()
	_ = i.Keyword()

	if !i.IsZero() {
		t.Errorf("%s failed: non-zero %T instance", t.Name(), i)
	}

	if got := i.String(); got != badAddr {
		t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
			t.Name(), badAddr, got)
	}

	var typ string = i.Keyword().String()

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

	// try every comparison operator supported in
	// this context ...
	brf := i.BRF()
	for j := 0; j < brf.Len(); j++ {
		cop, meth := brf.Index(j + 1)
		if meth == nil {
			t.Errorf("%s [%s] multival failed: expected %s method (%T), got nil",
				t.Name(), i.Keyword(), cop.Context(), meth)
		}

		wcop := sprintf("( %s %s %q )", i.Keyword(), cop, i)
		if T := meth(); T.Paren().String() != wcop {
			err := unexpectedStringResult(i.String(), wcop, T.String())
			t.Errorf("%s [%s] multival failed [%s rule]: %v",
				t.Name(), i.Keyword(), typ, err)
		}
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

func ExampleFQDN_Ne() {
	var f FQDN = DNS()

	// Can be set incrementally ...
	f.Set(`www`)
	f.Set(`example`)
	f.Set(`com`)

	// OR as a whole value ...
	// f.Set(`www.example.com`)

	fmt.Printf("%s", f.Ne())
	// Output: dns != "www.example.com"
}

func ExampleIPAddr_Set() {
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

/*
This example demonstrates the creation of an instance of IPAddr, which
is used in a variety of contexts.

In this example, a string name is fed to the package level IP function to form
a complete IPAddr instance, which is then shown in string representation.
*/
func ExampleIP() {
	ip := IP(`10.0.0.1`)
	fmt.Printf("%s", ip)
	// Output: 10.0.0.1
}

/*
This example demonstrates the string representation of the receiver instance.
*/
func ExampleIPAddr_String() {
	fmt.Printf("%s", IP(`192.168.56.7`))
	// Output: 192.168.56.7
}

/*
This example demonstrates the use of the useless Keyword method, as IPAddr
instances do not have any knowledge of Keywords at this time.
*/
func ExampleIPAddr_Keyword() {
	fmt.Printf("%v", IP(`10.0`).Keyword())
	// Output: ip
}

/*
This example demonstrates the use of the useless Kind method, as this information
is normally derived from a Keyword, which the receiver does not have.
*/
func ExampleIPAddr_Kind() {
	fmt.Printf("%s", IP(`192.168.1`).Kind())
	// Output: ip
}

/*
This example demonstrates the use of the useless Len method, as this information
is only made available to satisfy Go's interface signature requirements as they
pertain to the IPAddrContext interface.
*/
func ExampleIPAddr_Len() {
	fmt.Printf("%d", IP(`10.8.`).Len())
	// Output: 1
}

/*
This example demonstrates a check of the receiver for "nilness".
*/
func ExampleIPAddr_IsZero() {
	fmt.Printf("%t", IP(`192.168.7.*,10.,172.12.*,8.8.8.8`).IsZero())
	// Output: false
}

/*
This example demonstrates a check of the receiver for an aberrant state.
*/
func ExampleIPAddr_Valid() {
	ip := IP(``)
	fmt.Printf("Valid: %t", ip.Valid() == nil)
	// Output: Valid: false
}

func ExampleIPAddr_Eq() {
	var i IPAddr
	i.Set(`192.8.`).Set(`10.7.0`)

	fmt.Printf("%s", i.Eq())
	// Output: ip = "192.8.,10.7.0"
}

func ExampleIPAddr_Ne() {
	var i IPAddr
	i.Set(`10.8.`)

	fmt.Printf("%s", i.Ne())
	// Output: ip != "10.8."
}

func ExampleFQDN_Set() {
	var i FQDN
	i.Set(`*`).Set(`example`).Set(`com`)
	fmt.Printf("%s", i)
	// Output: *.example.com
}

func ExampleFQDN_Eq_oneShot() {
	fmt.Printf("%s", DNS(`www`, `example`, `com`).Eq())
	// Output: dns = "www.example.com"
}

/*
This example demonstrates the creation of an instance of FQDN, which
is used in a variety of contexts.

In this example, a string name is fed to the package level IP function to form
a complete FQDN instance, which is then shown in string representation.
*/
func ExampleDNS() {
	i := DNS(`example.com`)
	fmt.Printf("%s", i)
	// Output: example.com
}

/*
This example demonstrates the string representation of the receiver instance.
*/
func ExampleFQDN_String() {
	fmt.Printf("%s", DNS(`example.com`))
	// Output: example.com
}

/*
This example demonstrates the use of the useless Keyword method, as FQDN
instances do not have any knowledge of Keywords at this time.
*/
func ExampleFQDN_Keyword() {
	fmt.Printf("%v", DNS(`example.com`).Keyword())
	// Output: dns
}

/*
This example demonstrates a check of the receiver for "nilness".
*/
func ExampleFQDN_IsZero() {
	fmt.Printf("%t", DNS(`www.www.www.www.www.example.com`).IsZero())
	// Output: false
}

/*
This example demonstrates a check of the receiver for an aberrant state.
*/
func ExampleFQDN_Valid() {
	i := DNS(``)
	fmt.Printf("Valid: %t", i.Valid() == nil)
	// Output: Valid: false
}

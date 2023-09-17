package aci

import (
	"fmt"
	"testing"
)

func TestFQDN(t *testing.T) {
	var f FQDN
	_ = f.Len()
	_ = f.Keyword()
	_ = f.Eq()
	_ = f.Ne()
	_ = f.Valid()
	f = DNS()
	var typ string = f.Keyword().String()

	if f.len() != 0 {
		t.Errorf("%s failed: unexpected %T length: want '%d', got '%d'",
			t.Name(), f, 0, f.len())
		return
	}

	if err := f.Valid(); err == nil {
		t.Errorf("%s failed: empty %T deemed valid", t.Name(), f)
		return
	}

	f.Set()
	f.Set(``)
	f.Set(`www`, `example`, `com`)

	want := `www.example.com`
	got := f.String()

	if want != got {
		t.Errorf("%s failed; want '%s', got '%s'", t.Name(), want, got)
		return
	}

	absurd := `eeeeeeeeeeeeeeeeeeeeeeeee#eee^eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeexample`

	if validLabel(absurd) {
		t.Errorf("%s failed: bogus %T label accepted as valid (%s)",
			t.Name(), absurd, absurd)
		return
	}

	var F FQDN
	if F.String() != badFQDN {
		t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
			t.Name(), badFQDN, F)
		return
	}

	F.Set(`www`).Set(`$&^#*(`).Set(absurd).Set(`example`).Set(``).Set(`com`)
	if llen := F.Len(); llen != 3 {
		t.Errorf("%s failed; want '%d', got '%d'", t.Name(), 3, llen)
		return
	}

	// try every comparison operator supported in
	// this context ...
	brm := F.BRM()
	for i := 0; i < brm.Len(); i++ {
		cop, meth := brm.Index(i + 1)
		wcop := sprintf("( %s %s \"www.example.com\" )", f.Keyword(), cop)
		if T := meth(); T.Paren().String() != wcop {
			err := unexpectedStringResult(F.String(), wcop, T.String())
			t.Errorf("%s [%s] multival failed [%s rule]; %s, %s: %v",
				t.Name(), F.Keyword(), cop.Context(), cop.Description(), typ, err)
			return
		}
	}
}

func TestDNS_alternativeFQDN(t *testing.T) {
	want := `www.example.com`
	got := DNS(`www.example.com`)

	if want != got.String() {
		t.Errorf("%s failed; want '%s', got '%s'", t.Name(), want, got)
		return
	}
}

func TestIPAddr_BRM(t *testing.T) {
	var i IPAddr
	_ = i.Len()
	_ = i.Valid()
	_ = i.Keyword()

	if !i.IsZero() {
		t.Errorf("%s failed: non-zero %T instance", t.Name(), i)
		return
	}

	if got := i.String(); got != badAddr {
		t.Errorf("%s failed: unexpected string result; want '%s', got '%s'",
			t.Name(), badAddr, got)
		return
	}

	var typ string = i.Keyword().String()

	if !i.unique(`192.168.0`) {
		t.Errorf("%s failed; uniqueness check returned bogus result",
			t.Name())
		return
	}
	i.Set(`192.168.0`)
	i.Set(`12.3.45.*`)
	i.Set(`12.3.45.*`) // duplicate
	i.Set(`10.0.0.0/8`)

	if lens := i.Len(); lens != 3 {
		t.Errorf("%s failed: bad %T length; want '%d', got '%d'", t.Name(), i, 3, lens)
		return
	}

	if cond := i.Ne(); cond.IsZero() {
		t.Errorf("%s failed: nil %T instance!", t.Name(), cond)
		return
	}

	// try every comparison operator supported in
	// this context ...
	brm := i.BRM()
	for j := 0; j < brm.Len(); j++ {
		cop, meth := brm.Index(j + 1)
		if meth == nil {
			t.Errorf("%s [%s] multival failed: expected %s method (%T), got nil",
				t.Name(), i.Keyword(), cop.Context(), meth)
			return
		}

		wcop := sprintf("( %s %s %q )", i.Keyword(), cop, i)
		if T := meth(); T.Paren().String() != wcop {
			err := unexpectedStringResult(i.String(), wcop, T.String())
			t.Errorf("%s [%s] multival failed [%s rule]: %v",
				t.Name(), i.Keyword(), typ, err)
			return
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

/*
This example demonstrates the SHA-1 hash comparison between two (2)
IPAddr instances using the Compare method.
*/
func ExampleIPAddr_Compare() {
	addr1 := IP(`10.1.,192.168.`)
	addr2 := IP(`10.1.,192.168.1.`)

	fmt.Printf("Hashes are equal: %t", addr1.Compare(addr2))
	// Output: Hashes are equal: false
}

/*
This example demonstrates the SHA-1 hash comparison between two (2)
FQDN instances using the Compare method.
*/
func ExampleFQDN_Compare() {
	addr1 := DNS(`www`, `example`, `com`)
	addr2 := DNS(`www.example.com`)

	fmt.Printf("Hashes are equal: %t", addr1.Compare(addr2))
	// Output: Hashes are equal: true
}

func ExampleFQDN_BRM() {
	var host FQDN
	host.Set(`www.example.com`)
	cops := host.BRM()
	fmt.Printf("%T allows Eq: %t", host, cops.Contains(`=`))
	// Output: aci.FQDN allows Eq: true
}

func ExampleFQDN_Len() {
	var host FQDN
	host.Set(`www`)
	host.Set(`example`)
	host.Set(`com`)
	//host.Set(`www.example.com`)	// same!

	fmt.Printf("%T contains %d DNS labels", host, host.Len())
	// Output: aci.FQDN contains 3 DNS labels
}

func ExampleIPAddr_BRM() {
	var address IPAddr
	address.Set(`192.168.0`)
	cops := address.BRM()
	fmt.Printf("%T allows Eq: %t", address, cops.Contains(`=`))
	// Output: aci.IPAddr allows Eq: true
}

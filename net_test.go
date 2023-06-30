package aci

import (
	"fmt"
	"testing"
)

func TestFQDN(t *testing.T) {
	var f FQDN = DNS()
	f.Set(`www`, `example`, `com`)

	want := `www.example.com`
	got := f.String()

	if want != got {
		t.Errorf("%s failed; want '%s', got '%s'", t.Name(), want, got)
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
	i.Set(`192.168.0`)
	i.Set(`12.3.45.*`)
	i.Set(`10.0.0.0/8`)

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
	fmt.Printf("%s", i.Ne())
	// Output: ip != "192.168.0,12.3.45.*,10.0.0.0/8"
}

func ExampleIPAddr_Eq_oneShot() {
	fmt.Printf("%s", IP(`192.168.0`, `12.3.45.*`, `10.0.0.0/8`).Eq())
	// Output: ip = "192.168.0,12.3.45.*,10.0.0.0/8"
}

package aci

import (
	"testing"
)

func TestDistinguishedName_codecov(t *testing.T) {
	var (
		bd BindDistinguishedName
		td TargetDistinguishedName
	)

	user := `cn=Jesse Coretta,ou=People,dc=example,dc=com`
	if td.Set(user, Target); td.IsZero() {
		t.Errorf("%s failed [%T.Set]: instance is nil",
			t.Name(), td)
	}
	if bd.Set(user, BindUDN); bd.IsZero() {
		t.Errorf("%s failed [%T.Set]: instance is nil",
			t.Name(), bd)
	}

	want := LocalScheme + user
	if got := td.String(); want != got {
		t.Errorf("%s failed [%T.String compare]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), td, want, got)
	}

	if got := bd.String(); want != got {
		t.Errorf("%s failed [%T.String compare]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), bd, want, got)
	}
}

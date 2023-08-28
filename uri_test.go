package aci

import (
	"testing"
)

func TestLDAPURI_Parse(t *testing.T) {
	want := `ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(employeeStatus=active))`
	var l LDAPURI
	if err := l.Parse(want); err != nil {
		t.Errorf("%s failed [LDAPURI.Parse()]: %v",
			t.Name(), err)
	}

	if got := l.String(); want != got {
		t.Errorf("%s failed: [LDAPURI.Parse(compare)]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), want, got)
	}
}

func TestURI_initParse(t *testing.T) {
	want := `ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(employeeStatus=active))`
	if got := URI(want); want != got.String() {
		t.Errorf("%s failed: [LDAPURI.initParse(compare)]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), want, got)
	}
}

func TestURI_bindRules(t *testing.T) {
	var x LDAPURI

	// codecov
	if !x.Eq().IsZero() {
		t.Errorf("%s failed: want 'true', got 'false'", t.Name())
	} else if !x.Ne().IsZero() {
		t.Errorf("%s failed: want 'true', got 'false'", t.Name())
	}

	x = URI()
	x.Set(
		SingleLevel,
		UDN(`ou=People,dc=example,dc=com`),
		Filter(`(&(objectClass=employee)(employeeStatus=active))`),
	)

	want := `userdn = "ldap:///ou=People,dc=example,dc=com??one?(&(objectClass=employee)(employeeStatus=active))"`
	if got := x.Eq().String(); got != want {
		t.Errorf("%s failed: [LDAPURI.piecemeal(compare)]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), want, got)
	}

	// overwrite UserDN with GroupDN
	x.Set(GDN(`ou=Groups,dc=example,dc=com`))
	// overwrite user filter with group filter
	x.Set(Filter(`(&(objectClass=groupOfNames)(ownerStatus=active))`))
	want = `( groupdn != "ldap:///ou=Groups,dc=example,dc=com??one?(&(objectClass=groupOfNames)(ownerStatus=active))" )`
	if got := x.Ne().Paren().String(); got != want {
		t.Errorf("%s failed: [LDAPURI.piecemeal(compare)]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), want, got)
	}
}

func TestURI_piecemeal(t *testing.T) {
	piecemeal := URI()
	piecemeal.Set(
		SingleLevel,
		UDN(`ou=People,dc=example,dc=com`),
		Filter(`(&(objectClass=employee)(employeeStatus=active))`),
	)

	want := `ldap:///ou=People,dc=example,dc=com??one?(&(objectClass=employee)(employeeStatus=active))`
	if got := piecemeal; want != got.String() {
		t.Errorf("%s failed: [LDAPURI.piecemeal(compare)]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), want, got)
	}

	/*
		// FIX ME
			var piecedessert LDAPURI
			piecedessert.Set(Filter(`(&(objectClass=employee)(employeeStatus=active))`))
			piecedessert.Set(UDN(`ou=People,dc=example,dc=com`))
			piecedessert.Set(SingleLevel)

			if got := piecedessert; want != got.String() {
				t.Errorf("%s failed: [LDAPURI.piecemeal(alt. compare)]:\nwant: '%s'\ngot:  '%s'",
		                        t.Name(), want, got)
			}
	*/
}

/*
put codecov tests here that may be awkward and out-of-place
anywhere else...
*/
func TestURI_codecov(t *testing.T) {
	var l LDAPURI
	if wat := l.String(); len(wat) != 0 {
		t.Errorf("%s failed: unexpected value; want '', got '%s'",
			t.Name(), wat)
	}

	if !l.Eq().IsZero() {
		t.Errorf("%s failed: want 'true', got 'false'", t.Name())

	} else if !l.Ne().IsZero() {
		t.Errorf("%s failed: want 'true', got 'false'", t.Name())
	}

	// missing scheme pfx
	loser := `ou=People,dc=example,dc=com???`
	if err := l.Parse(loser); err == nil {
		t.Errorf("%s failed [missing URI scheme]: want 'error', got 'nil'",
			t.Name())
	}

	// someone is trying to do a remote URI.
	// get bent.
	loser = `ldap://ldap.example.com/dc=example,dc=com?objectClass?sub?(objectClass=*)`
	if err := l.Parse(loser); err == nil {
		t.Errorf("%s failed [SERIOUS VULNERABILITY]: non-local URI returned no error",
			t.Name())
	}

	// nice try dingus
	loser = `http:///ou=People,dc=example,dc=com???`
	if err := l.Parse(loser); err == nil {
		t.Errorf("%s failed [missing URI scheme]: want 'error', got 'nil'",
			t.Name())
	}

	// atbtv
	atbtval := `ldap:///ou=People,dc=example,dc=com?owner#GROUPDN`
	if err := l.Parse(atbtval); err != nil {
		t.Errorf("%s failed [atbtv URI parse]: %v", t.Name(), err)
	}

	if got := l.String(); atbtval != got {
		t.Errorf("%s failed [atbtv compare]: want '%s', got '%s'",
			t.Name(), atbtval, got)
	}

}

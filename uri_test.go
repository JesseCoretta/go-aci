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

	x.Set(GDN(`ou=Groups,dc=example,dc=com`))
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

package aci

import (
	"fmt"
	"testing"
)

func TestLDAPURI_Parse(t *testing.T) {

	for _, want := range []string{
		`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(employeeStatus=active))`,
		`ldap:///ou=People,dc=example,dc=com?manager#SELFDN`,
		`ldap:///ou=People,dc=example,dc=com??sub?`,
	} {
		var x LDAPURI
		if err := x.Parse(want); err != nil {
			t.Errorf("%s failed [LDAPURI.Parse()]: %v",
				t.Name(), err)
			return
		}
		_ = x.String()
		_ = x.Valid()
		_ = x.Eq()
		_ = x.Ne()
		_ = x.Kind()
		x.isDistinguishedNameContext()

		if got := x.String(); want != got {
			t.Errorf("%s failed: [LDAPURI.Parse(compare)]:\nwant: '%s'\ngot:  '%s'",
				t.Name(), want, got)
			return
		}
	}

	donotwant := `ldap:///ou=People,dc=example,dc=com?#SELFDN`
	var x LDAPURI
	if err := x.Parse(donotwant); err == nil {
		t.Errorf("%s failed: no error where one was expected",
			t.Name())
		return
	}
}

func TestURI_initParse(t *testing.T) {
	want := `ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(employeeStatus=active))`
	if got := URI(want); want != got.String() {
		t.Errorf("%s failed: [LDAPURI.initParse(compare)]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), want, got)
		return
	}
}

func TestURI_bindRules(t *testing.T) {
	var x LDAPURI
	_ = x.Keyword()
	_ = x.String()
	_ = x.Valid()

	// codecov
	if !x.Eq().IsZero() {
		t.Errorf("%s failed: want 'true', got 'false'", t.Name())
		return
	} else if !x.Ne().IsZero() {
		t.Errorf("%s failed: want 'true', got 'false'", t.Name())
		return
	}

	x = URI()
	x.Set(
		SingleLevel,
		UDN(`ou=People,dc=example,dc=com`),
		Filter(`(&(objectClass=employee)(employeeStatus=active))`),
	)

	want := `userdn = "ldap:///ou=People,dc=example,dc=com??one?(&(objectClass=employee)(employeeStatus=active))"`
	if got := x.Eq().String(); got != want {
		t.Errorf("%s failed: [LDAPURI(%s).piecemeal(compare)]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), x.Keyword(), want, got)
		return
	}

	// overwrite UserDN with GroupDN
	x.Set(GDN(`ou=Groups,dc=example,dc=com`))
	// overwrite user filter with group filter
	x.Set(Filter(`(&(objectClass=groupOfNames)(ownerStatus=active))`))
	want = `( groupdn != "ldap:///ou=Groups,dc=example,dc=com??one?(&(objectClass=groupOfNames)(ownerStatus=active))" )`
	if got := x.Ne().Paren().String(); got != want {
		t.Errorf("%s failed: [LDAPURI(%s).piecemeal(compare)]:\nwant: '%s'\ngot:  '%s'",
			t.Name(), x.Keyword(), want, got)
		return
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
		return
	}

	/*
		// TODO FIX ME
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
		return
	}

	if !l.Eq().IsZero() {
		t.Errorf("%s failed: want 'true', got 'false'", t.Name())
		return

	} else if !l.Ne().IsZero() {
		t.Errorf("%s failed: want 'true', got 'false'", t.Name())
		return
	}

	// missing scheme pfx
	loser := `ou=People,dc=example,dc=com???`
	if err := l.Parse(loser); err == nil {
		t.Errorf("%s failed [missing URI scheme]: want 'error', got 'nil'",
			t.Name())
		return
	}

	// someone is trying to do a remote URI.
	// get bent.
	loser = `ldap://ldap.example.com/dc=example,dc=com?objectClass?sub?(objectClass=*)`
	if err := l.Parse(loser); err == nil {
		t.Errorf("%s failed [SERIOUS VULNERABILITY]: non-local URI returned no error",
			t.Name())
		return
	}

	// nice try dingus
	loser = `http:///ou=People,dc=example,dc=com???`
	if err := l.Parse(loser); err == nil {
		t.Errorf("%s failed [missing URI scheme]: want 'error', got 'nil'",
			t.Name())
		return
	}

	// atbtv
	atbtval := `ldap:///ou=People,dc=example,dc=com?owner#GROUPDN`
	if err := l.Parse(atbtval); err != nil {
		t.Errorf("%s failed [atbtv URI parse]: %v", t.Name(), err)
		return
	}

	if got := l.String(); atbtval != got {
		t.Errorf("%s failed [atbtv compare]: want '%s', got '%s'",
			t.Name(), atbtval, got)
		return
	}

}

func ExampleURI() {
	dn := UDN(`ou=People,o=example`)
	filter := Filter(`(objectClass=employee)`)
	scope := Subtree

	uri := URI(dn, filter, scope)

	fmt.Printf("%s", uri)
	// Output: ldap:///ou=People,o=example??sub?(objectClass=employee)
}

func ExampleLDAPURI_Parse() {
	raw := `ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))`

	var uri LDAPURI
	if err := uri.Parse(raw); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", uri)
	// Output: ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))
}

func ExampleLDAPURI_Set() {
	var uri LDAPURI
	uri.Set(`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))`)
	fmt.Printf("%s", uri)
	// Output: ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))
}

func ExampleLDAPURI_IsZero() {
	var uri LDAPURI
	fmt.Printf("Zero: %t", uri.IsZero())
	// Output: Zero: true
}

func ExampleLDAPURI_Valid() {
	var uri LDAPURI
	fmt.Printf("Valid: %t", uri.Valid() == nil)
	// Output: Valid: false
}

func ExampleLDAPURI_Keyword() {
	dn := GDN(`ou=Groups,dc=example,dc=com`)
	filter := Filter(`(&(objectClass=distributionList)(status=active))`)
	uri := URI(dn, filter)

	fmt.Printf("Keyword: %s", uri.Keyword())
	// Output: Keyword: groupdn
}

func ExampleLDAPURI_Kind() {
	dn := GDN(`ou=Groups,dc=example,dc=com`)
	filter := Filter(`(&(objectClass=distributionList)(status=active))`)
	uri := URI(dn, filter)

	fmt.Printf("Keyword: %s", uri.Keyword())
	// Output: Keyword: groupdn
}

func ExampleLDAPURI_Len() {
	dn := GDN(`ou=Groups,dc=example,dc=com`)
	filter := Filter(`(&(objectClass=distributionList)(status=active))`)
	uri := URI(dn, filter)

	fmt.Printf("Len: %d", uri.Len())
	// Output: Len: 0
}

func ExampleLDAPURI_String() {
	var uri LDAPURI
	uri.Set(`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))`)
	fmt.Printf("%s", uri)
	// Output: ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))
}

func ExampleLDAPURI_Eq() {
	dn := GDN(`ou=Groups,dc=example,dc=com`)
	filter := Filter(`(&(objectClass=distributionList)(status=active))`)
	attrs := UAs(`cn`, `sn`, `givenName`, `objectClass`, `uid`)
	uri := URI(dn, attrs, filter, SingleLevel)

	fmt.Printf("%s", uri.Eq())
	// Output: groupdn = "ldap:///ou=Groups,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=distributionList)(status=active))"
}

func ExampleLDAPURI_Ne() {
	raw := `ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))`
	uri := URI(raw)
	fmt.Printf("%s", uri.Ne())
	// Output: userdn != "ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))"
}

func ExampleLDAPURI_BRM() {
	var uri LDAPURI
	fmt.Printf("%d available aci.BindRuleMethod instances", uri.BRM().Len())
	// Output: 2 available aci.BindRuleMethod instances
}

/*
This example demonstrates the SHA-1 hash comparison between two (2)
Inheritance instances using the Compare method.
*/
func ExampleLDAPURI_Compare() {
	raw := `ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))`
	uri1 := URI(raw)

	raw = `ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?sub?(&(objectClass=distributionList)(status=active))`
	uri2 := URI(raw)

	fmt.Printf("Hashes are equal: %t", uri1.Compare(uri2))
	// Output: Hashes are equal: false
}

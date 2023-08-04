package aci

import (
	"fmt"
	"testing"
)

/*
This example demonstrates the creation of a userattr Bind Rule Condition bearing
an attributeType paired with the LDAPURL BindType keyword.
*/
func ExampleUAT_userDN() {
	at := ATName(`manager`)
	uat := UAT(at, USERDN)
	fmt.Printf("%s", uat.Eq())
	// Output: userattr = "manager#USERDN"
}

/*
This example demonstrates the creation of a userattr Bind Rule Condition bearing
an attributeType paired with the GROUPDN BindType keyword.
*/
func ExampleUAT_groupDN() {
	at := ATName(`manager`)
	uat := UAT(at, GROUPDN)
	fmt.Printf("%s", uat.Eq())
	// Output: userattr = "manager#GROUPDN"
}

/*
This example demonstrates the creation of a userattr Bind Rule Condition bearing
an attributeType paired with the ROLEDN BindType keyword.
*/
func ExampleUAT_roleDN() {
	at := ATName(`manager`)
	uat := UAT(at, ROLEDN)
	fmt.Printf("%s", uat.Eq())
	// Output: userattr = "manager#ROLEDN"
}

/*
This example demonstrates the creation of a userattr Bind Rule Condition bearing
an attributeType paired with the SELFDN BindType keyword.
*/
func ExampleUAT_selfDN() {
	at := ATName(`owner`)
	uat := UAT(at, SELFDN)
	fmt.Printf("%s", uat.Eq())
	// Output: userattr = "owner#SELFDN"
}

/*
This example demonstrates the creation of a userattr Bind Rule Condition bearing
an attributeType paired with the LDAPURL BindType keyword.
*/
func ExampleUAT_lDAPURL() {
	at := ATName(`authorizationURI`)
	uat := UAT(at, LDAPURL)
	fmt.Printf("%s", uat.Eq())
	// Output: userattr = "authorizationURI#LDAPURL"
}

/*
This example demonstrates the creation of a groupattr Bind Rule Condition bearing
an attributeType paired with the USERDN BindType keyword.
*/
func ExampleGAT_userDN() {
	at := ATName(`manager`)
	gat := GAT(at, USERDN)
	fmt.Printf("%s", gat.Eq())
	// Output: groupattr = "manager#USERDN"
}

/*
This example demonstrates the creation of a groupattr Bind Rule Condition bearing
an attributeType paired with the GROUPDN BindType keyword.
*/
func ExampleGAT_groupDN() {
	at := ATName(`manager`)
	gat := GAT(at, GROUPDN)
	fmt.Printf("%s", gat.Eq())
	// Output: groupattr = "manager#GROUPDN"
}

/*
This example demonstrates the creation of a groupattr Bind Rule Condition bearing
an attributeType paired with the ROLEDN BindType keyword.
*/
func ExampleGAT_roleDN() {
	at := ATName(`manager`)
	gat := GAT(at, ROLEDN)
	fmt.Printf("%s", gat.Eq())
	// Output: groupattr = "manager#ROLEDN"
}

/*
This example demonstrates the creation of a groupattr Bind Rule Condition bearing
an attributeType paired with the SELFDN BindType keyword.
*/
func ExampleGAT_selfDN() {
	at := ATName(`owner`)
	gat := GAT(at, SELFDN)
	fmt.Printf("%s", gat.Eq())
	// Output: groupattr = "owner#SELFDN"
}

/*
This example demonstrates the creation of a groupattr Bind Rule Condition bearing
an attributeType paired with the LDAPURL BindType keyword.
*/
func ExampleGAT_lDAPURL() {
	at := ATName(`authorizationURI`)
	gat := GAT(at, LDAPURL)
	fmt.Printf("%s", gat.Eq())
	// Output: groupattr = "authorizationURI#LDAPURL"
}

func TestFilter(t *testing.T) {
	// this is a temporary and admittedly naïve filter "parser". A future
	// release of this package shall utilize a proper LDAP Search Filter
	// decompiler that transfers conditions into a go-stackage Stack.
	want := `(&(objectClass=employee)(|(objectClass=shareholder)(objectClass=engineeringLead)))`
	got := Filter(want)
	if want != got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

func TestUAT(t *testing.T) {
	uat := UAT(ATName(`manager`), USERDN).Eq()
	want := `userattr = "manager#USERDN"`
	got := uat.String()
	if want != got {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

func TestGAT(t *testing.T) {
	gat := GAT(ATName(`owner`), LDAPURL).Eq()
	want := `groupattr = "owner#LDAPURL"`
	got := gat.String()
	if want != got {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

func TestAttributeFilter(t *testing.T) {
	filter := `(&(objectClass=employee)(|(objectClass=shareholder)(objectClass=engineeringLead)))`
	attr := `objectClass`

	af := AF(attr, filter)
	want := sprintf("%s:%s", attr, filter)
	got := af.String()
	if want != got {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

func TestAttributeFilters(t *testing.T) {
	attrs := []string{
		`objectClass`,
		`objectClass`,
	}

	filters := []string{
		`(&(objectClass=employee)(|(objectClass=shareholder)(objectClass=engineeringLead)))`,
		`(&(objectClass=executive)(|(objectClass=shareholder)(objectClass=legal)))`,
	}

	ops := []AttributeOperation{
		AddOp,
		DelOp,
	}

	wants := []string{
		`add=objectClass:(&(objectClass=employee)(|(objectClass=shareholder)(objectClass=engineeringLead)))`,
		`delete=objectClass:(&(objectClass=executive)(|(objectClass=shareholder)(objectClass=legal)))`,
	}

	for i := 0; i < 2; i++ {
		at := attrs[i]
		fr := filters[i]
		op := ops[i]
		af := AF(at, fr)
		afs := op.AF(af)
		if wants[i] != afs.String() {
			t.Errorf("%s failed: want '%s', got '%s'", t.Name(), wants[i], afs)
		}
	}
}

func TestAttributeFilters_multiple(t *testing.T) {

	want := `add=objectClass:(&(objectClass=employee)(|(objectClass=shareholder)(objectClass=engineeringLead))) && homeDirectory:(&(objectClass=executive)(|(objectClass=shareholder)(objectClass=legal)))`
	got := AddOp.AF(
		AF(`objectClass`, `(&(objectClass=employee)(|(objectClass=shareholder)(objectClass=engineeringLead)))`),
		AF(`homeDirectory`, `(&(objectClass=executive)(|(objectClass=shareholder)(objectClass=legal)))`),
	)

	if want != got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

/*
This example demonstrates the creation of a Condition that expresses an AttributeFilters Target Rule.
*/
func ExampleAttributeFilters_Eq() {

	afs := AddOp.AF(
		AF(`objectClass`, `(&(objectClass=employee)(|(objectClass=shareholder)(objectClass=engineeringLead)))`),
		AF(`homeDirectory`, `(&(objectClass=executive)(|(objectClass=shareholder)(objectClass=legal)))`),
	)

	fmt.Printf("%s", afs.Eq())
	// Output: ( targattrfilters = "add=objectClass:(&(objectClass=employee)(|(objectClass=shareholder)(objectClass=engineeringLead))) && homeDirectory:(&(objectClass=executive)(|(objectClass=shareholder)(objectClass=legal)))" )
}

func TestURI_standardURI(t *testing.T) {
	// dn, filter, attrs and scope
	got := URI(
		UDN(`ou=Contractors,ou=People,dc=example,dc=com`),
		SingleLevel,
		Attrs().Push(`cn`, `sn`, `givenName`, `objectClass`, `drink`),
		Filter(`(objectClass=*)`),
	)

	want := `ldap:///ou=Contractors,ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,drink?one?(objectClass=*)`
	if want != got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}

	// dn, filter, no attrs, no scope
	got = URI(
		UDN(`ou=Contractors,ou=People,dc=example,dc=com`),
		Filter(`(objectClass=*)`),
	)
	want = `ldap:///ou=Contractors,ou=People,dc=example,dc=com???(objectClass=*)`
	if want != got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}

	// dn and atbtv, no attrs, no scope, no filter
	got = URI(
		UDN(`ou=Groups,dc=example,dc=com`),
		GAT(ATName(`manager`), GROUPDN),
	)

	want = `ldap:///ou=Groups,dc=example,dc=com?manager#GROUPDN`
	if want != got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

/*
This example demonstrates the creation of an LDAP Search URI bearing a scope, dn, filter and attributeType list.
*/
func ExampleURI_standard() {
	// dn, filter, attrs and scope
	u := URI(
		UDN(`ou=Contractors,ou=People,dc=example,dc=com`),
		SingleLevel,
		Attrs().Push(`cn`, `sn`, `givenName`, `objectClass`, `drink`),
		Filter(`(objectClass=*)`),
	)
	fmt.Printf("%s", u)
	// Output: ldap:///ou=Contractors,ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,drink?one?(objectClass=*)
}

/*
This example demonstrates the creation of an LDAP Search URI bearing a DN, Filter and default scope w/o attributeType enumeration.
*/
func ExampleURI_dnAndFilterNoAttributes() {
	u := URI(
		UDN(`ou=Contractors,ou=People,dc=example,dc=com`),
		Subtree,
		Filter(`(objectClass=*)`),
	)

	// Note that no explicit scope will introduce the default
	// of base.

	fmt.Printf("%s", u)
	// Output: ldap:///ou=Contractors,ou=People,dc=example,dc=com??sub?(objectClass=*)
}

/*
This example demonstrates the creation of an LDAP Search URI bearing a DN and AttributeBindTypeOrValue only.
*/
func ExampleURI_dnAndAttributeBindTypeOrValueOnly() {
	u := URI(
		UDN(`ou=Groups,dc=example,dc=com`),
		GAT(ATName(`manager`), GROUPDN),
	)

	fmt.Printf("%s", u)
	// Output: ldap:///ou=Groups,dc=example,dc=com?manager#GROUPDN
}

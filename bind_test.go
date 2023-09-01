package aci

import (
	"fmt"
	"testing"
)

func TestParseBindRule(t *testing.T) {
	want := `userdn = "ldap:///cn=Jesse Coretta,ou=People,dc=example,dc=com"`

	c, err := ParseBindRule(want)
	if err != nil {
		return
	}

	if want != c.String() {
		t.Errorf("%s failed:\nwant '%s'\ngot '%s'", t.Name(), want, c)
	}
}

func TestParseBindRules(t *testing.T) {
	want := `( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) )`

	r, err := ParseBindRules(want)
	if err != nil {
		return
	}

	if want != r.String() {
		t.Errorf("%s failed:\nwant '%s',\ngot  '%s'", t.Name(), want, r)
	}
}

/*
This example demonstrates the indexing, iteration and execution of the available
comparison operator methods for the BindDistinguishedName type.
*/
func ExampleBindRuleFuncs() {
	var dn BindDistinguishedName = GDN(`cn=X.500 Administrators,ou=Groups,dc=example,dc=com`)
	brf := dn.BRF()

	for i := 0; i < brf.Len(); i++ {
		cop, meth := brf.Index(i + 1)                              // zero (0) should never be accessed, start at 1
		fmt.Printf("[%s] %s\n", cop.Description(), meth().Paren()) // enable parentheticals, because why not
	}
	// Output:
	// [Equal To] ( groupdn = "ldap:///cn=X.500 Administrators,ou=Groups,dc=example,dc=com" )
	// [Not Equal To] ( groupdn != "ldap:///cn=X.500 Administrators,ou=Groups,dc=example,dc=com" )
}

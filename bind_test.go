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

/*
This example demonstrates the union between a group distinguished name and a timeframe,
expressed as a BindRules instance. Parenthetical encapsulation is enabled for inner stack
elements, but not the outer (AND) stack itself.
*/
func ExampleAnd() {
	and := And(
		GDN(`cn=X.500 Administrators,ou=Groups,dc=example,dc=com`).Eq().Paren(),
		Timeframe(ToD(`1730`), ToD(`2330`)).Paren(),
	)
	fmt.Printf("%s", and)
	// Output: ( groupdn = "ldap:///cn=X.500 Administrators,ou=Groups,dc=example,dc=com" ) AND ( timeofday >= "1730" AND timeofday < "2330" )
}

/*
This example demonstrates a logical OR between a value matching bind rule and an LDAP URI
bearing the userdn keyword context. Parentheticals are enabled at every point.
*/
func ExampleOr() {
	or := Or(
		UAT(`manager`, `LDAPURL`).Eq().Paren(),
		URI(UDN(`ou=People,dc=example,dc=com`), Subtree).Eq().Paren(),
	)
	fmt.Printf("%s", or.Paren())
	// Output: ( ( userattr = "manager#LDAPURL" ) OR ( userdn = "ldap:///ou=People,dc=example,dc=com??sub?" ) )
}

/*
This example demonstrates a logical NOT that excludes a value matching userattr context or
an LDAPURI bearing the userdn key context. The NOT operation should generally encompass one
(1) or more conditions within an OR context.  Additionally, NOT operations generally reside
within an outer AND context as shown. YMMV.
*/
func ExampleNot() {
	and := And(
		IP(`192.168.`).Eq().Paren(),
		Not(Or(
			UAT(`manager`, `LDAPURL`).Eq().Paren(),
			URI(UDN(`ou=People,dc=example,dc=com`), Subtree).Eq().Paren(),
		).Paren()),
	)
	fmt.Printf("%s", and)
	// Output: ( ip = "192.168." ) AND NOT ( ( userattr = "manager#LDAPURL" ) OR ( userdn = "ldap:///ou=People,dc=example,dc=com??sub?" ) )
}

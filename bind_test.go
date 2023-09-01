package aci

import (
	"fmt"
	"testing"
)

func TestParseBindRuleFuncs(t *testing.T) {
	brf := newBindRuleFuncs(bindRuleFuncMap{})

	_ = brf.Len()
	_ = brf.IsZero()
	_, _ = brf.Index(0)

	ssf := SSF(256)
	brf = ssf.BRF()

	for i := 0; i < brf.Len(); i++ {
		if cop, meth := brf.Index(i + 1); meth().String() != fmt.Sprintf("ssf %s %q", cop, `256`) {
			t.Errorf("%s failed: failed to call index %d [%s] non-nil %T", t.Name(), i, cop.Context(), brf)
		}
	}
}

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

func ExampleBindRuleFuncs_Index() {
	ssf := SSF(256)
	brf := ssf.BRF()

	for i := 0; i < brf.Len(); i++ {
		// IMPORTANT: Do not call index 0. Either adjust your
		// loop variable (i) to begin at 1, and terminate at
		// brf.Len()+1 --OR-- simply +1 the index call as we
		// are doing here (seems easier). The reason for this
		// is because there is no valid ComparisonOperator
		// with an underlying uint8 value of zero (0). See
		// the ComparisonOperator constants for details.
		idx := i + 1
		cop, meth := brf.Index(idx)

		// create the bindrule, and make it parenthetical
		rule := meth().Paren()

		// grab the raw string output
		fmt.Printf("[%d] %T instance [%s] execution returned %T: %s\n", idx, meth, cop.Context(), rule, rule)
	}
	// Output:
	// [1] aci.BindRuleMethod instance [Eq] execution returned aci.BindRule: ( ssf = "256" )
	// [2] aci.BindRuleMethod instance [Ne] execution returned aci.BindRule: ( ssf != "256" )
	// [3] aci.BindRuleMethod instance [Lt] execution returned aci.BindRule: ( ssf < "256" )
	// [4] aci.BindRuleMethod instance [Gt] execution returned aci.BindRule: ( ssf > "256" )
	// [5] aci.BindRuleMethod instance [Le] execution returned aci.BindRule: ( ssf <= "256" )
	// [6] aci.BindRuleMethod instance [Ge] execution returned aci.BindRule: ( ssf >= "256" )
}

func ExampleBindRuleFuncs_IsZero() {
	var brf BindRuleFuncs
	fmt.Printf("Zero: %t", brf.IsZero())
	// Output: Zero: true
}

func ExampleBindRuleFuncs_Valid() {
	var brf BindRuleFuncs
	fmt.Printf("Error: %v", brf.Valid())
	// Output: Error: aci.BindRuleFuncs instance is nil
}

func ExampleBindRuleFuncs_Len() {
	// Note: we need not populate the value to get a
	// BRF list, but the methods in that list won't
	// actually work until the instance (ssf) is in
	// an acceptable state. Since all we're doing
	// here is checking the length, a receiver that
	// is nil/zero is totally fine.
	var ssf SecurityStrengthFactor // not init'd
	total := ssf.BRF().Len()

	fmt.Printf("There are %d available comparison operator methods for %T BindRules", total, ssf)
	// Output: There are 6 available comparison operator methods for aci.SecurityStrengthFactor BindRules
}

func ExampleBindRuleMethod() {
	ssf := SSF(256)
	brf := ssf.BRF()

	// verify that the receiver (ssf) is copacetic
	// and will produce a legal expression if meth
	// is executed
	if err := brf.Valid(); err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < brf.Len(); i++ {
		// IMPORTANT: Do not call index 0. Either adjust your
		// loop variable (i) to begin at 1, and terminate at
		// brf.Len()+1 --OR-- simply +1 the index call as we
		// are doing here (seems easier). The reason for this
		// is because there is no valid ComparisonOperator
		// with an underlying uint8 value of zero (0). See
		// the ComparisonOperator constants for details.
		idx := i + 1
		cop, meth := brf.Index(idx)

		// execute meth and modify the returned BindRule
		// instance to make it parenthetical in one shot.
		rule := meth().Paren()

		// grab the raw string output
		fmt.Printf("[%d] %T instance [%s] execution returned %T: %s\n", idx, meth, cop.Context(), rule, rule)
	}
	// Output:
	// [1] aci.BindRuleMethod instance [Eq] execution returned aci.BindRule: ( ssf = "256" )
	// [2] aci.BindRuleMethod instance [Ne] execution returned aci.BindRule: ( ssf != "256" )
	// [3] aci.BindRuleMethod instance [Lt] execution returned aci.BindRule: ( ssf < "256" )
	// [4] aci.BindRuleMethod instance [Gt] execution returned aci.BindRule: ( ssf > "256" )
	// [5] aci.BindRuleMethod instance [Le] execution returned aci.BindRule: ( ssf <= "256" )
	// [6] aci.BindRuleMethod instance [Ge] execution returned aci.BindRule: ( ssf >= "256" )
}

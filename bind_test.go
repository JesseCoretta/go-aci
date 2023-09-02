package aci

import (
	"fmt"
	"testing"
)

func TestParseBindRuleMethods(t *testing.T) {
	var brf BindRuleMethods
	_ = brf.Len()
	_ = brf.IsZero()
	_, _ = brf.Index(0)

	brf = newBindRuleMethods(bindRuleFuncMap{})

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

// mainly this exists to satisfy codecov, but also
// aid in identifying panic points.
func TestBindRule_bogus(t *testing.T) {
	var br BindRule
	_ = br.ID()
	_ = br.Category()
	_ = br.IsZero()
	_ = br.Len()
	_ = br.Valid()
	_ = br.Paren()
	_ = br.Operator()
	_ = br.Expression()
	_ = br.Keyword()
	_ = br.String()
}

// mainly this exists to satisfy codecov, but also
// aid in identifying panic points.
func TestBindRules_bogus(t *testing.T) {
	var br BindRules
	_ = br.ID()
	_ = br.Category()
	_ = br.IsZero()
	_ = br.Len()
	_ = br.Fold()
	_ = br.Valid()
	_ = br.Paren()
	_ = br.ReadOnly()
	_ = br.NoPadding()
	_ = br.String()
	_ = br.Index(-100)
	_, _ = br.Traverse([]int{1, 2, 3, 4}...)
}

func TestParseBindRule(t *testing.T) {
	want := `userdn = "ldap:///cn=Jesse Coretta,ou=People,dc=example,dc=com"`

	var b BindRule
	var err error
	b.isBindContextQualifier() // just to satisfy codecov.
	_ = b.Kind()
	_ = b.IsNesting()
	_ = b.Operator()
	_ = b.Keyword()
	_ = b.Expression()
	_ = b.SetQuoteStyle(1)

	if b, err = ParseBindRule(want); err != nil {
		return
	}

	b.isBindContextQualifier()
	_ = b.Kind()
	_ = b.IsNesting()
	_ = b.Operator()
	_ = b.Keyword()
	_ = b.Expression()
	_ = b.SetQuoteStyle(0)

	if want != b.String() {
		t.Errorf("%s failed:\nwant '%s'\ngot '%s'", t.Name(), want, b)
	}
}

func TestParseBindRules(t *testing.T) {
	want := `( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) )`

	var r BindRules
	var err error

	r.isBindContextQualifier()
	_ = r.Kind()
	_ = r.IsNesting()

	if r, err = ParseBindRules(want); err != nil {
		return
	}

	_ = r.Kind()
	_ = r.IsNesting()

	if want != r.String() {
		t.Errorf("%s failed:\nwant '%s',\ngot  '%s'", t.Name(), want, r)
	}
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

/*
This example demonstrates the indexing, iteration and execution of the available
comparison operator methods for the BindDistinguishedName type.
*/
func ExampleBindRuleMethods() {
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

func ExampleBindRuleMethods_Index() {
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

		// execute method to create the bindrule, while
		// also enabling the (optional) parenthetical bit
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

func ExampleBindRuleMethods_Index_byText() {
	ssf := SSF(256)
	brf := ssf.BRF()

	// Here, we demonstrate calling a particular BindRuleMethod
	// not by its numerical index, but rather by its actual
	// "symbolic" operator representation. Keep in mind these
	// options for text-based searches:
	//
	// - symbols (e.g.: '=', '>') are available via ComparisonOperator.String()
	// - func names (e.g.: 'Eq', 'Gt') are available via ComparisonOperator.Context()
	// - descriptions (e.g.: 'Not Equal To', 'Less Than') are available via ComparisonOperator.Description()
	//
	// As such, feel free to replace these list items with one of the above methods,
	// but keep in mind that text based searches are more resource intensive than as
	// compared to direct ComparisonOperator numeric calls. If you have performance
	// concerns, avoid this text based procedure.
	for i, term := range []string{
		`=`,
		`!=`,
		`<`,
		`>`,
		`<=`,
		`>=`,
	} {
		// IMPORTANT: Do not call index 0. Either adjust your
		// loop variable (i) to begin at 1, and terminate at
		// brf.Len()+1 --OR-- simply +1 the index call as we
		// are doing here (seems easier). The reason for this
		// is because there is no valid ComparisonOperator
		// with an underlying uint8 value of zero (0). See
		// the ComparisonOperator constants for details.
		cop, meth := brf.Index(term)

		// execute method to create the bindrule, while
		// also enabling the (optional) parenthetical bit
		rule := meth().Paren()

		// grab the raw string output
		fmt.Printf("[%d] %T instance [%s] execution returned %T: %s\n", i+1, meth, cop.Context(), rule, rule)
	}
	// Output:
	// [1] aci.BindRuleMethod instance [Eq] execution returned aci.BindRule: ( ssf = "256" )
	// [2] aci.BindRuleMethod instance [Ne] execution returned aci.BindRule: ( ssf != "256" )
	// [3] aci.BindRuleMethod instance [Lt] execution returned aci.BindRule: ( ssf < "256" )
	// [4] aci.BindRuleMethod instance [Gt] execution returned aci.BindRule: ( ssf > "256" )
	// [5] aci.BindRuleMethod instance [Le] execution returned aci.BindRule: ( ssf <= "256" )
	// [6] aci.BindRuleMethod instance [Ge] execution returned aci.BindRule: ( ssf >= "256" )
}

func ExampleBindRuleMethods_IsZero() {
	var brf BindRuleMethods
	fmt.Printf("Zero: %t", brf.IsZero())
	// Output: Zero: true
}

func ExampleBindRuleMethods_Valid() {
	var brf BindRuleMethods
	fmt.Printf("Error: %v", brf.Valid())
	// Output: Error: aci.BindRuleMethods instance is nil
}

func ExampleBindRuleMethods_Len() {
	// Note: we need not populate the value to get a
	// BRF list, but the methods in that list won't
	// actually work until the instance (ssf) is in
	// an acceptable state. Since all we're doing
	// here is checking the length, a receiver that
	// is nil/zero is totally fine.
	var ssf SecurityStrengthFactor // not init'd
	total := ssf.BRF().Len()

	fmt.Printf("There are %d available aci.BindRuleMethod instances for creating %T BindRules", total, ssf)
	// Output: There are 6 available aci.BindRuleMethod instances for creating aci.SecurityStrengthFactor BindRules
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

		// execute method to create the bindrule, while
		// also enabling the (optional) parenthetical bit
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

package aci

import (
	"fmt"
	"testing"
)

func TestBindRuleMethods(t *testing.T) {
	var brm BindRuleMethods
	_ = brm.Len()
	_ = brm.IsZero()
	_, _ = brm.Index(0)

	brm = newBindRuleMethods(bindRuleFuncMap{})

	_ = brm.Len()
	_ = brm.IsZero()
	_, _ = brm.Index(0)

	ssf := SSF(256)
	brm = ssf.BRM()

	for i := 0; i < brm.Len(); i++ {
		if cop, meth := brm.Index(i + 1); meth().String() != fmt.Sprintf("ssf %s %q", cop, `256`) {
			t.Errorf("%s failed: failed to call index %d [%s] non-nil %T", t.Name(), i, cop.Context(), brm)
		}
	}
}

func TestBindRules_wordToStack(t *testing.T) {
	for word, expect := range map[string]bool{
		`BAND`:                         false,
		`AND`:                          true,
		`aNd`:                          true,
		`and`:                          true,
		`andy`:                         false,
		`OR`:                           true,
		`or`:                           true,
		`oR`:                           true,
		`orwellian`:                    false,
		`NOT`:                          true,
		`not`:                          true,
		`nOT`:                          true,
		`AND NOT`:                      true,
		`and not`:                      true,
		`aNd NoT`:                      true,
		`andnot`:                       true,
		`and not a moment too soon...`: false,
		``:                             false,
		`&*#($`:                        false,
	} {
		if _, result := wordToStack(word); expect != result {
			t.Errorf("%s unexpected result: want %t, got %t",
				t.Name(), expect, result)
		}
	}
}

// mainly this exists to satisfy codecov, but also
// aid in identifying panic points.
func TestBindRule_bogus(t *testing.T) {
	var br BindRule
	br.isBindContextQualifier() // just to satisfy codecov.
	_ = br.ID()
	_ = br.Category()
	_ = br.IsZero()
	_ = br.Paren()
	_ = br.Len()
	_ = br.Kind()
	_ = br.Valid()
	_ = br.Paren()
	_ = br.Operator()
	_ = br.IsNesting()
	_ = br.NoPadding()
	_ = br.Expression()
	_ = br.Keyword()
	_ = br.SetQuoteStyle(1)
	_ = br.String()
}

// mainly this exists to satisfy codecov, but also
// aid in identifying panic points.
func TestBindRules_bogus(t *testing.T) {
	var br BindRules
	br.isBindContextQualifier() // just to satisfy codecov.
	_ = br.ID()
	_ = br.Category()
	_ = br.IsNesting()
	_ = br.IsZero()
	_ = br.Len()
	_ = br.Fold()
	_ = br.Valid()
	_ = br.Paren()
	_ = br.ReadOnly()
	_ = br.NoPadding()
	_ = br.String()
	_ = br.Push(nil)
	_ = br.Pop()
	_ = br.Index(-100)
	_ = br.Traverse([]int{1, 2, 3, 4}...)
	br.reset()

	br = And(SSF(128).Eq, UDN("uid=jesse,ou=People,dc=example,dc=com").Eq())
	_ = br.Kind()
	_ = br.IsNesting()
	_ = br.Keyword()
	_ = br.Pop()
	_ = br.Push()

}

/*
This example demonstrates the useless nature of the Traverse method for a receiver
that is an instance of BindRule. As BindRule is logically "singular", there is no
structure in which a traversal would be possible. The Traverse method only exists
to satisfy Go's interface signature requirements as they pertain to the BindContext
type, and this test exists only to maintain code coverage and to convey this message.

Execution of this method simply returns the receiver.
*/
func ExampleBindRule_Traverse() {
	br := SSF(71).Eq()
	fmt.Printf("%T", br.Traverse(1, 2, 3, 4, 5))
	// Output: aci.BindRule
}

/*
This example demonstrates the traversal of a BindRules structure to obtain a specific
nested element, in this case the 'ssf >= "128"' expression.
*/
func ExampleBindRules_Traverse() {
	// And() is the variable for the outermost stack, don't count it as an index.
	// Rather, begin counting its children instead.
	rules := And(
		GDN(`cn=X.500 Administrators,ou=Groups,dc=example,dc=com`).Eq().Paren(),
		Timeframe(ToD(`1730`), ToD(`2330`)).Paren(),

		// Enter the Or stack by descending within the third element (AND slice #2)
		Or(
			UAT(`manager`, `LDAPURL`).Eq().Paren(),
			GAT(`owner`, SELFDN).Eq().Paren(),
			URI(UDN(`ou=People,dc=example,dc=com`), Subtree).Eq().Paren(),
			// OR slice #3
			And(
				// Inner AND slice #0
				SSF(128).Ge(),
			),
		),
	)

	ssf := rules.Traverse(2, 3, 0)
	fmt.Printf("%s", ssf)
	// Output: ssf >= "128"
}

/*
This is an identical scenario to the above Traverse example, except in this scenario we
will be writing to the slice returned.
*/
func ExampleBindRules_Traverse_andWrite() {
	// And() is the variable for the outermost stack, don't count it as an index.
	// Rather, begin counting its children instead.
	rules := And(
		GDN(`cn=X.500 Administrators,ou=Groups,dc=example,dc=com`).Eq().Paren(),
		Timeframe(ToD(`1730`), ToD(`2330`)).Paren(),

		// Enter the Or stack by descending within the third element (AND slice #2)
		Or(
			UAT(`manager`, `LDAPURL`).Eq().Paren(),
			GAT(`owner`, SELFDN).Eq().Paren(),
			URI(UDN(`ou=People,dc=example,dc=com`), Subtree).Eq().Paren(),
			// OR slice #3
			And(
				// Inner AND slice #0
				SSF(128).Ge(),
			),
		),
	)

	// Call the specific stack/slice we want. Remember,
	// Traverse returns a BindContext interface, which
	// will be either BindRule OR BindRules. In this
	// demonstration, we know what it is, thus we need
	// not perform type switching in a case statement.
	raw := rules.Traverse(2, 3, 0)
	asserted, ok := raw.(BindRule)
	if !ok {
		fmt.Printf("Failed to assert %T", asserted)
		return
	}

	// Because go-stackage is so heavily pointer-based,
	// we need not worry about "writing the updated SSF
	// value back to the stack". Just make the changes
	// to the condition itself, and be done with it.
	asserted.SetExpression(SSF(164)) // set to 164 because I'm so arbitrary

	// Do the stack walk again to see if the pointer updated ...
	fmt.Printf("%s", rules.Traverse(2, 3, 0))
	// Output: ssf >= "164"
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
	brm := dn.BRM()

	for i := 0; i < brm.Len(); i++ {
		cop, meth := brm.Index(i + 1)                              // zero (0) should never be accessed, start at 1
		fmt.Printf("[%s] %s\n", cop.Description(), meth().Paren()) // enable parentheticals, because why not
	}
	// Output:
	// [Equal To] ( groupdn = "ldap:///cn=X.500 Administrators,ou=Groups,dc=example,dc=com" )
	// [Not Equal To] ( groupdn != "ldap:///cn=X.500 Administrators,ou=Groups,dc=example,dc=com" )
}

func ExampleBindRuleMethods_Index() {
	ssf := SSF(256)
	brm := ssf.BRM()

	for i := 0; i < brm.Len(); i++ {
		// IMPORTANT: Do not call index 0. Either adjust your
		// loop variable (i) to begin at 1, and terminate at
		// brm.Len()+1 --OR-- simply +1 the index call as we
		// are doing here (seems easier). The reason for this
		// is because there is no valid ComparisonOperator
		// with an underlying uint8 value of zero (0). See
		// the ComparisonOperator constants for details.
		idx := i + 1
		cop, meth := brm.Index(idx)

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
	brm := ssf.BRM()

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
		// brm.Len()+1 --OR-- simply +1 the index call as we
		// are doing here (seems easier). The reason for this
		// is because there is no valid ComparisonOperator
		// with an underlying uint8 value of zero (0). See
		// the ComparisonOperator constants for details.
		cop, meth := brm.Index(term)

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
	var brm BindRuleMethods
	fmt.Printf("Zero: %t", brm.IsZero())
	// Output: Zero: true
}

func ExampleBindRuleMethods_Valid() {
	var brm BindRuleMethods
	fmt.Printf("Error: %v", brm.Valid())
	// Output: Error: aci.BindRuleMethods instance is nil
}

func ExampleBindRuleMethods_Len() {
	// Note: we need not populate the value to get a
	// BRM list, but the methods in that list won't
	// actually work until the instance (ssf) is in
	// an acceptable state. Since all we're doing
	// here is checking the length, a receiver that
	// is nil/zero is totally fine.
	var ssf SecurityStrengthFactor // not init'd
	total := ssf.BRM().Len()

	fmt.Printf("There are %d available aci.BindRuleMethod instances for creating %T BindRules", total, ssf)
	// Output: There are 6 available aci.BindRuleMethod instances for creating aci.SecurityStrengthFactor BindRules
}

func ExampleBindRuleMethod() {
	ssf := SSF(256)
	brm := ssf.BRM()

	// verify that the receiver (ssf) is copacetic
	// and will produce a legal expression if meth
	// is executed
	if err := brm.Valid(); err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < brm.Len(); i++ {
		// IMPORTANT: Do not call index 0. Either adjust your
		// loop variable (i) to begin at 1, and terminate at
		// brm.Len()+1 --OR-- simply +1 the index call as we
		// are doing here (seems easier). The reason for this
		// is because there is no valid ComparisonOperator
		// with an underlying uint8 value of zero (0). See
		// the ComparisonOperator constants for details.
		idx := i + 1
		cop, meth := brm.Index(idx)

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

func ExampleBindRule_Compare() {
	ssf1 := SSF(128).Eq()
	ssf2 := SSF(127).Eq()
	fmt.Printf("Equal: %t", ssf1.Compare(ssf2))
	// Output: Equal: false
}

func ExampleBindRules_Compare() {
	tf1 := Timeframe(ToD(`0130`), ToD(`1605`))
	tf2 := Timeframe(ToD(`1215`), ToD(`1605`))

	fmt.Printf("Equal: %t", tf1.Compare(tf2))
	// Output: Equal: false
}

func ExampleBindRule_Category() {
	fmt.Printf("%s", SSF(71).Eq().Category())
	// Output: ssf
}

func ExampleBindRules_Category() {
	fmt.Printf("%s", And(SSF(71).Eq()).Category())
	// Output: and
}

func ExampleBindRule_Keyword() {
	fmt.Printf("%s", SSF(0).Ne().Keyword())
	// Output: ssf
}

func ExampleBindRule_Kind() {
	var tr BindRule
	fmt.Printf("%s", tr.Kind())
	// Output: condition
}

func ExampleBindRule_ID() {
	fmt.Printf("%s", IP(`192.168.`).Ne().ID())
	// Output: bind
}

func ExampleBindRule_Expression() {
	dn := `uid=jesse,ou=Contractors,ou=People,dc=example,dc=com`
	fmt.Printf("%s", UDN(dn).Eq().Expression())
	// Output: ldap:///uid=jesse,ou=Contractors,ou=People,dc=example,dc=com
}

func ExampleBindRule_IsZero() {
	var tr BindRule
	fmt.Printf("Zero: %t", tr.IsZero())
	// Output: Zero: true
}

func ExampleBindRule_Valid() {
	var tr BindRule
	fmt.Printf("Valid: %t", tr.Valid() == nil)
	// Output: Valid: false
}

func ExampleBindRule_SetQuoteStyle() {
	var tgt BindRule
	tgt.SetKeyword(BindUDN)
	tgt.SetOperator(Ne)
	tgt.SetExpression(UDNs(
		UDN(`ldap:///uid=jesse,ou=People,dc=example,dc=com`),
		UDN(`ldap:///uid=courtney,ou=People,dc=example,dc=com`),
		UDN(`ldap:///uid=jimmy,ou=People,dc=example,dc=com`),
	))

	tgt.Paren() // optional

	tgt.SetQuoteStyle(0)
	style1 := tgt.String()

	tgt.SetQuoteStyle(1)
	style2 := tgt.String()

	fmt.Printf("\n0: %s\n1: %s", style1, style2)
	// Output:
	// 0: ( userdn != "ldap:///uid=jesse,ou=People,dc=example,dc=com" || "ldap:///uid=courtney,ou=People,dc=example,dc=com" || "ldap:///uid=jimmy,ou=People,dc=example,dc=com" )
	// 1: ( userdn != "ldap:///uid=jesse,ou=People,dc=example,dc=com || ldap:///uid=courtney,ou=People,dc=example,dc=com || ldap:///uid=jimmy,ou=People,dc=example,dc=com" )
}

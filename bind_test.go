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
			return
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
			return
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

	br = And(
		SSF(128).Ge(),
		GDN(`cn=Executives,ou=Groups,dc=example,dc=com`).Ne(),
		UDNs("uid=jesse,ou=People,dc=example,dc=com", "uid=courtney,ou=People,dc=example,dc=com").Eq(),
	)
	_ = br.Kind()
	_ = br.IsNesting()
	_ = br.Keyword()
	_ = br.Pop()
	_ = br.Push()
	_ = br.insert(`fnord`, 0)
	_ = br.Index(1)
	replacer := GDN(`cn=Executive Assistants,ou=Groups,dc=example,dc=com`).Ne()
	br.Replace(replacer, 1)

	replaced := br.Index(1)
	if replacer.String() != replaced.String() {
		t.Errorf("%s failed: %T.Replace did not replace specified slice value, want '%s', got '%s'",
			t.Name(), br, replacer, replaced)
		return
	}

	ctx := br.Traverse(1)
	if ctx.String() != replaced.String() {
		t.Errorf("%s failed: %T.Traverse did not return replaced slice value, want '%s', got '%s'",
			t.Name(), br, replaced, ctx)
		return
	}
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

/*
This example demonstrates the various capabilities of a BindRules instance, as well as
the use of some so-called "prefabricator" functions for additional convenience.

Note that BindRules should always be initialized before use. The reason for this is
because it may be required to set a logical Boolean operating mode, such as 'AND',
'OR' and '(AND) NOT'. This governs how the individual slices (which may descend into
other stacks) are string represented in a Boolean expressive statement.
*/
func ExampleBindRules() {
	// Outer BindRules structure (a.k.a.: the "top")
	// is a non-parenthetical ORed BindRules stack.
	var brs BindRules = Or(

		// First ORed condition (slice #0) is a
		// non-parenthetical ANDed BindRules stack
		// with operator folding enabled.
		And(
			// A single non-parenthetical groupdn
			// equality assertion BindRule occupies
			// slice #0 of the current stack. Note
			// that in this case, calling the needed
			// operator method is necessary, as GDN
			// is not a prefab function.
			//
			// Result:
			GDN(`cn=Accounting,ou=Groups,dc=example,dc=com`).Eq(), // groupdn

			// Timeframe is a BindRules prefab function
			// used to conveniently produce a pair of
			// `timeofday` conditions within an ANDed
			// logical context in order to specify a
			// "time-window" during which access will
			// be granted or withheld in some manner.
			//
			// This expression occupies slice #1 of the
			// current stack and is parenthetical. There
			// is no need to execute the operator method
			// (e.g.: Eq) manually, as this function is
			// a prefabricator and does this for us 😎.
			// Additionally, toggle the Fold bit for the
			// return instance.
			//
			// Result: ( timeofday >= "0900" and timeofday < "1830" )
			Timeframe(
				ToD(`0900`), // notBefore
				ToD(`1830`), // notAfter
			).Paren().Fold(),

			// Weekdays is a BindRule prefab to conveniently
			// produce a sequence of Day instances expressing
			// Monday through Friday. See also Weekend for an
			// inverse of this functionality.
			//
			// This expression occupies slice #2 of the current
			// stack and is parenthetical. There is no need to
			// execute the operator method (e.g:. Eq) manually,
			// as this function automatically crafts the BindRule.
			//
			// Result: ( dayofweek = "Mon,Tues,Wed,Thur,Fri" )
			Weekdays(`=`).Paren(), // comparison operator could also be 1, `eq` or `equal to`
		).Fold(),

		// Second ORed condition (slice #1) is a
		// parenthetically ANDed BindRules stack.
		And(
			// Two (2) individual BindRule instances
			// circumscribed within a parenthetical
			// AND context. Note that in this case,
			// calling the comparison method manually
			// is required, as RDN and GAT (among
			// others) are not prefab functions.
			//
			// result: ( roledn = "ldap:///cn=Superusers,ou=Groups,dc=example,dc=com`" AND groupattr = "privilegeLevel#GAMMA5" )
			RDN(`cn=Superusers,ou=Groups,dc=example,dc=com`).Eq(), // roledn
			GAT(`privilegeLevel`, `GAMMA5`).Eq(),                  // groupattr
		).Paren(),
	)

	// Let's make some arbitrary changes ...
	//
	// Make the outer stack (OR) parenthetical
	brs.Paren() // no arg = toggle state, else use true/false for explicit set.

	fmt.Printf("%s", brs)
	// Output: ( groupdn = "ldap:///cn=Accounting,ou=Groups,dc=example,dc=com" and ( timeofday >= "0900" and timeofday < "1830" ) and ( dayofweek = "Mon,Tues,Wed,Thur,Fri" ) OR ( roledn = "ldap:///cn=Superusers,ou=Groups,dc=example,dc=com" AND groupattr = "privilegeLevel#GAMMA5" ) )
}

/*
This example demonstrates the toggling of Boolean WORD operator case-folding
of a BindRules instance.
*/
func ExampleBindRules_Fold() {
	strong := And(SSF(128).Ge(), EXTERNAL.Eq())
	strong.Fold() // we want `AND` to be `and`

	fmt.Printf("%s", strong)
	// Output: ssf >= "128" and authmethod = "SASL EXTERNAL"
}

func ExampleBindRules_ID() {
	var brs BindRules
	fmt.Printf("%s", brs.ID())
	// Output: bind
}

func ExampleBindRules_IsZero() {
	var brs BindRules
	fmt.Printf("Zero: %t", brs.IsZero())
	// Output: Zero: true
}

func ExampleBindRules_Valid() {
	var brs BindRules
	fmt.Printf("Valid: %t", brs.Valid() == nil)
	// Output: Valid: false
}

func ExampleBindRules_String() {
	strong := And(SSF(128).Ge(), EXTERNAL.Eq())

	fmt.Printf("%s", strong)
	// Output: ssf >= "128" AND authmethod = "SASL EXTERNAL"
}

/*
This example demonstrates the selective replacement of
a specific BindRules stack slice.
*/
func ExampleBindRules_Replace() {
	strong := And(
		SSF(128).Ge(),  // slice #0
		DIGESTMD5.Eq(), // slice #1
	)

	// Replace awful Digest-MD5 with the
	// SASL/EXTERNAL mechanism.
	strong.Replace(EXTERNAL.Eq(), 1) // <x> replace slice #1

	fmt.Printf("%s", strong)
	// Output: ssf >= "128" AND authmethod = "SASL EXTERNAL"
}

/*
This example demonstrates an attempt to modify a BindRules
stack instance while its ReadOnly bit is enabled.
*/
func ExampleBindRules_ReadOnly() {
	strong := And(
		SSF(128).Ge(),  // slice #0
		DIGESTMD5.Eq(), // slice #1
	)

	strong.ReadOnly()

	// Try to replace awful Digest-MD5 with
	// the SASL/EXTERNAL mechanism.
	strong.Replace(EXTERNAL.Eq(), 1)

	fmt.Printf("%s", strong)
	// Output: ssf >= "128" AND authmethod = "SASL DIGEST-MD5"
}

/*
This example demonstrates the addition of new slice elements
to a BindRules instance using its Push method.
*/
func ExampleBindRules_Push() {
	// create a single BindRules instance
	// with only one (1) slice (Timeframe
	// BindRule) ...
	brs := And(
		Timeframe(
			ToD(`0900`),
			ToD(`1830`),
		),
	)

	// Add a Weekdays BindRule prefab
	brs.Push(Weekdays(Eq))

	fmt.Printf("%s", brs)
	// Output: timeofday >= "0900" AND timeofday < "1830" AND dayofweek = "Mon,Tues,Wed,Thur,Fri"
}

/*
This example demonstrates the removal of a single slice
within a BindRules instance in LIFO fashion using its
Pop method.
*/
func ExampleBindRules_Pop() {
	// create a single BindRules instance
	// with two (2) slices ...
	brs := And(
		Timeframe(
			ToD(`0900`), // slice #0
			ToD(`1830`), // slice #1
		),
		Weekdays(Eq),
	)

	// Remove (by Pop) the Weekdays slice (slice #1)
	popped := brs.Pop()

	fmt.Printf("%s", popped)
	// Output: dayofweek = "Mon,Tues,Wed,Thur,Fri"
}

/*
This example demonstrates the interrogation of a BindRules
instance to determine whether any of its immediate slice
members are other stack elements, thereby indicating that
a "nesting condition" is in effect.
*/
func ExampleBindRules_IsNesting() {
	brs := And(
		Timeframe(
			ToD(`0900`),
			ToD(`1830`),
		),
		Weekdays(Eq),
	)

	fmt.Printf("Contains nesting elements: %t", brs.IsNesting())
	// Output: Contains nesting elements: true
}

/*
This example demonstrates the calling of a specific slice
member by its numerical index using the BindRules.Index
method.
*/
func ExampleBindRules_Index() {
	brs := And(
		Timeframe(
			ToD(`0900`),
			ToD(`1830`),
		),
		Weekdays(Eq),
	)

	fmt.Printf("%s", brs.Index(0))
	// Output: timeofday >= "0900" AND timeofday < "1830"
}

/*
This example demonstrates the use of the NoPadding method
to remove the outer padding of a BindRules instance. Note
that parentheticals are enabled for visual aid.
*/
func ExampleBindRules_NoPadding() {
	brs := And(
		Timeframe(
			ToD(`0900`),
			ToD(`1830`),
		),
		Weekdays(Eq),
	)

	brs.Paren().NoPadding()

	fmt.Printf("%s", brs)
	// Output: (timeofday >= "0900" AND timeofday < "1830" AND dayofweek = "Mon,Tues,Wed,Thur,Fri")
}

/*
This example demonstrates the interrogation of a BindRules
instance to determine its integer length using its Len
method.  The return value describes the number of slice
elements present within the instance.
*/
func ExampleBindRules_Len() {
	brs := And(
		Timeframe(
			ToD(`0900`),
			ToD(`1830`),
		),
		Weekdays(Eq),
	)

	fmt.Printf("Contains %d elements", brs.Len())
	// Output: Contains 2 elements
}

/*
This example demonstrates the use of the Kind method to
reveal the underlying type of the receiver. This is for
use during the handling of a combination of BindRules
and BindRule instances under the BindContext interface
context.
*/
func ExampleBindRules_Kind() {
	brs := And(
		Timeframe(
			ToD(`0900`),
			ToD(`1830`),
		),
		Weekdays(Eq),
	)

	fmt.Printf("%s", brs.Kind())
	// Output: stack
}

/*
This example demonstrates the use of the useless Keyword method
that returns a bogus keyword in string representation. Keywords
are only directly applicable to BindRule instances.
*/
func ExampleBindRules_Keyword() {
	tf := Timeframe(
		ToD(`0900`),
		ToD(`1830`),
	)

	fmt.Printf("%s", tf.Keyword())
	// Output: <invalid_bind_keyword>
}

/*
This example demonstrates the use of the Paren method to
enable the parenthetical setting for the receiver instance.
*/
func ExampleBindRules_Paren() {
	brs := And(
		Timeframe(
			ToD(`0900`),
			ToD(`1830`),
		),
		Weekdays(Eq),
	)

	brs.Paren()

	fmt.Printf("%s", brs)
	// Output: ( timeofday >= "0900" AND timeofday < "1830" AND dayofweek = "Mon,Tues,Wed,Thur,Fri" )
}

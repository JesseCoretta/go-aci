package aci

import (
	"fmt"
	"testing"
)

func TestTargetRuleMethods(t *testing.T) {
	var trm TargetRuleMethods
	_ = trm.Len()
	_ = trm.IsZero()
	_, _ = trm.Index(0)
	_, _ = trm.Index(140)

	trm = newTargetRuleMethods(targetRuleFuncMap{})

	_ = trm.Len()
	_ = trm.IsZero()
	_, _ = trm.Index(0)

	attrs := TAs(`cn`, `sn`, `givenName`, `objectClass`, `uid`, `homeDirectory`)
	trm = attrs.TRM()
	if trm.Len() != 2 {
		t.Errorf("%s failed: unexpected %T length: want %d, got %d", t.Name(), trm, 2, trm.Len())
	}
	_, _ = trm.Index(-100)
	_, _ = trm.Index(100)
}

func TestCtrls(t *testing.T) {
	L := Ctrls()
	o1 := Ctrl(`1.3.6.1.4.1.56521.101.2.1.1`)
	o2 := Ctrl(`1.3.6.1.4.1.56521.101.2.2.2`)
	o3 := Ctrl(`1.3.6.1.4.1.56521.101.3.1`)

	L.Push(o1, o2, o3)

	want := `1.3.6.1.4.1.56521.101.2.1.1 || 1.3.6.1.4.1.56521.101.2.2.2 || 1.3.6.1.4.1.56521.101.3.1`
	got := L.String()
	if want != got {
		t.Errorf("%s failed [oidORs]:\nwant '%s'\ngot  '%s'", t.Name(), want, got)
	}

	c := L.Eq()
	want = `( targetcontrol = "1.3.6.1.4.1.56521.101.2.1.1 || 1.3.6.1.4.1.56521.101.2.2.2 || 1.3.6.1.4.1.56521.101.3.1" )`
	if got = c.String(); got != want {
		t.Errorf("%s failed [makeTargetRule]:\nwant '%s'\ngot  '%s'", t.Name(), want, got)
	}
}

func TestTargetKeyword_Set_targetScope(t *testing.T) {
	got := SingleLevel.Eq()
	want := `( targetscope = "onelevel" )`
	if want != got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

// mainly this exists to satisfy codecov, but also
// aid in identifying panic points.
func TestTargetRule_bogus(t *testing.T) {
	var tr TargetRule
	_ = tr.ID()
	_ = tr.Kind()
	_ = tr.NoPadding()
	_ = tr.Category()
	_ = tr.IsZero()
	_ = tr.Valid()
	_ = tr.Operator()
	_ = tr.Expression()
	_ = tr.Keyword()
	_ = tr.String()
}

// mainly this exists to satisfy codecov, but also
// aid in identifying panic points.
func TestTargetRules_bogus(t *testing.T) {
	var tr TargetRules
	_ = tr.ID()
	_ = tr.Category()
	_ = tr.IsZero()
	_ = tr.Len()
	_ = tr.Valid()
	_ = tr.ReadOnly()
	_ = tr.NoPadding()
	_ = tr.String()
	_ = tr.Pop()
	_ = tr.Index(100)
	_ = tr.remove(14)
	tr.reset()
}

/*
This example demonstrates a similar scenario to the one described in the above example, but with
an alternative means of quotation demonstrated. Additionally, string primitives are used instead
of proper ExtOp style OIDs.
*/
func ExampleExtOps_alternativeQuotationScheme() {
	// Here we set double-quote encapsulation
	// upon the Rule instance created by the
	// ExtOps function.
	ext := ExtOps().Push(
		// These aren't real control OIDs.
		`1.3.6.1.4.1.56521.999.5`,
		`1.3.6.1.4.1.56521.999.6`,
		`1.3.6.1.4.1.56521.999.7`,
	)

	fmt.Printf("%s", ext.Eq().SetQuoteStyle(0)) // see MultivalSliceQuotes const for details
	// Output: ( extop = "1.3.6.1.4.1.56521.999.5" || "1.3.6.1.4.1.56521.999.6" || "1.3.6.1.4.1.56521.999.7" )
}

/*
This example demonstrates how to use a single Target DN to craft a Target Rule Equality
Condition.
*/
func ExampleTargetDistinguishedName_Eq_target() {
	dn := TDN(`uid=jesse,ou=People,dc=example,dc=com`)
	fmt.Printf("%s", dn.Eq())
	// Output: ( target = "ldap:///uid=jesse,ou=People,dc=example,dc=com" )
}

/*
This example demonstrates how a list of Target DNs can be used to create a single Target
Rule. First, create a Rule using TDNs().Parens(), then push N desired TDN (Target DN)
values into the Rule.
*/
func ExampleTargetDistinguishedNames_Eq() {
	tdns := TDNs().Push(
		TDN(`uid=jesse,ou=People,dc=example,dc=com`),
		TDN(`uid=courtney,ou=People,dc=example,dc=com`),
	)

	// Craft an equality Condition
	fmt.Printf("%s", tdns.Eq())
	// Output: ( target = "ldap:///uid=jesse,ou=People,dc=example,dc=com || ldap:///uid=courtney,ou=People,dc=example,dc=com" )
}

func TestAttrs_attrList(t *testing.T) {
	ats := TAs().Push(
		AT(`cn`),
		AT(`sn`),
		AT(`givenName`),
		AT(`homeDirectory`),
		AT(`uid`),
	) //.SetQuoteStyle(0)

	// Style #0 (MultivalOuterQuotes)
	//want := `( targetattr = "cn || sn || givenName || homeDirectory || uid" )`

	// Style #1 (MultivalSliceQuotes)
	want := `( targetattr = "cn" || "sn" || "givenName" || "homeDirectory" || "uid" )`

	got := ats.Eq().SetQuoteStyle(MultivalSliceQuotes).String()

	if got != want {
		t.Errorf("%s failed [attrList]:\nwant '%s'\ngot  '%s'", t.Name(), want, got)
	}
}

/*
This example demonstrates how to create a Target Attributes Rule using a list of AttributeType instances.
*/
func ExampleTAs() {
	attrs := TAs().Push(
		AT(`cn`),
		AT(`sn`),
		AT(`givenName`),
	)
	fmt.Printf("%s", attrs)
	// Output: cn || sn || givenName
}

/*
This example demonstrates how to create a Target Attributes Rule Equality Condition using a list of
AttributeType instances.
*/
func ExampleAttributeTypes_Eq_targetAttributes() {
	attrs := TAs().Push(
		AT(`cn`),
		AT(`sn`),
		AT(`givenName`),
	)
	fmt.Printf("%s", attrs.Eq())
	// Output: ( targetattr = "cn || sn || givenName" )
}

/*
This example demonstrates how to craft a Target Scope Rule Condition for a onelevel Search Scope.
*/
func ExampleSearchScope_Eq_targetScopeOneLevel() {
	fmt.Printf("%s", SingleLevel.Eq())
	// Output: ( targetscope = "onelevel" )
}

/*
This example demonstrates how to craft a Target Rule Condition bearing the `targetfilter` keyword
and an LDAP Search Filter.
*/
func ExampleFilter() {
	tf := Filter(`(&(uid=jesse)(objectClass=*))`)
	fmt.Printf("%s", tf.Eq())
	// Output: ( targetfilter = "(&(uid=jesse)(objectClass=*))" )
}

/*
This example demonstrates how to craft a set of Target Rule Conditions.
*/
func ExampleTRs() {
	t := TRs().Push(
		TDN(`uid=jesse,ou=People,dc=example,dc=com`).Eq(),
		Filter(`(&(uid=jesse)(objectClass=*))`).Eq(),
		ExtOp(`1.3.6.1.4.1.56521.999.5`).Eq(),
	)
	fmt.Printf("%s", t)
	// Output: ( target = "ldap:///uid=jesse,ou=People,dc=example,dc=com" ) ( targetfilter = "(&(uid=jesse)(objectClass=*))" ) ( extop = "1.3.6.1.4.1.56521.999.5" )
}

/*
This example demonstrates the indexing, iteration and execution of the available
TargetRuleMethod instances for the TargetDistinguishedName type.
*/
func ExampleTargetRuleMethods() {
	var tdn TargetDistinguishedName = TTDN(`uid=*,ou=People,dc=example,dc=com`)
	trm := tdn.TRM()

	for i := 0; i < trm.Len(); i++ {
		cop, meth := trm.Index(i + 1) // zero (0) should never be accessed, start at 1
		fmt.Printf("[%s] %s\n", cop.Description(), meth())
	}
	// Output:
	// [Equal To] ( target_to = "ldap:///uid=*,ou=People,dc=example,dc=com" )
	// [Not Equal To] ( target_to != "ldap:///uid=*,ou=People,dc=example,dc=com" )
}

func ExampleTargetRuleMethods_Index() {
	var dn TargetDistinguishedName = TFDN(`uid=*,ou=People,dc=example,dc=com`)
	trm := dn.TRM()

	for i := 0; i < trm.Len(); i++ {
		// IMPORTANT: Do not call index 0. Either adjust your
		// loop variable (i) to begin at 1, and terminate at
		// trm.Len()+1 --OR-- simply +1 the index call as we
		// are doing here (seems easier). The reason for this
		// is because there is no valid ComparisonOperator
		// with an underlying uint8 value of zero (0). See
		// the ComparisonOperator constants for details.
		idx := i + 1
		cop, meth := trm.Index(idx)

		// execute method to create the targetrule
		rule := meth()

		// grab the raw string output
		fmt.Printf("[%d] %T instance [%s] execution returned %T: %s\n", idx, meth, cop.Context(), rule, rule)
	}
	// Output:
	// [1] aci.TargetRuleMethod instance [Eq] execution returned aci.TargetRule: ( target_from = "ldap:///uid=*,ou=People,dc=example,dc=com" )
	// [2] aci.TargetRuleMethod instance [Ne] execution returned aci.TargetRule: ( target_from != "ldap:///uid=*,ou=People,dc=example,dc=com" )
}

func ExampleTargetRuleMethods_Index_byText() {
	attrs := TAs(`cn`, `sn`, `givenName`, `objectClass`, `uid`, `homeDirectory`)
	trm := attrs.TRM()

	// Here, we demonstrate calling a particular TargetRuleMethod
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
	} {
		// IMPORTANT: Do not call index 0. Either adjust your
		// loop variable (i) to begin at 1, and terminate at
		// trm.Len()+1 --OR-- simply +1 the index call as we
		// are doing here (seems easier). The reason for this
		// is because there is no valid ComparisonOperator
		// with an underlying uint8 value of zero (0). See
		// the ComparisonOperator constants for details.
		cop, meth := trm.Index(term)

		// execute method to create the TargetRule, while
		// enabling the so-called "Slice Quotation scheme"
		// for one of the iterations (just for fun!)
		rule := meth().SetQuoteStyle(i)

		// grab the raw string output
		fmt.Printf("[%d] %T instance [%s] execution returned %T: %s\n", i+1, meth, cop.Context(), rule, rule)
	}
	// Output:
	// [1] aci.TargetRuleMethod instance [Eq] execution returned aci.TargetRule: ( targetattr = "cn" || "sn" || "givenName" || "objectClass" || "uid" || "homeDirectory" )
	// [2] aci.TargetRuleMethod instance [Ne] execution returned aci.TargetRule: ( targetattr != "cn || sn || givenName || objectClass || uid || homeDirectory" )
}

func ExampleTargetRuleMethods_IsZero() {
	var trm TargetRuleMethods
	fmt.Printf("Zero: %t", trm.IsZero())
	// Output: Zero: true
}

func ExampleTargetRuleMethods_Valid() {
	var trm TargetRuleMethods
	fmt.Printf("Error: %v", trm.Valid())
	// Output: Error: aci.TargetRuleMethods instance is nil
}

func ExampleTargetRuleMethods_Len() {
	// Note: we need not populate the value to get a
	// TRM list, but the methods in that list won't
	// actually work until the instance (ssf) is in
	// an acceptable state. Since all we're doing
	// here is checking the length, a receiver that
	// is nil/zero is totally fine.
	var sco SearchScope = SingleLevel // any would do
	total := sco.TRM().Len()

	fmt.Printf("There is one (%d) available aci.TargetRuleMethod instance for creating %T TargetRules", total, sco)
	// Output: There is one (1) available aci.TargetRuleMethod instance for creating aci.SearchScope TargetRules
}

func ExampleTargetRuleMethod() {
	tfil := Filter(`(&(objectClass=employee)(terminated=FALSE))`)
	trm := tfil.TRM()

	// verify that the receiver (ssf) is copacetic
	// and will produce a legal expression if meth
	// is executed
	if err := trm.Valid(); err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < trm.Len(); i++ {
		// IMPORTANT: Do not call index 0. Either adjust your
		// loop variable (i) to begin at 1, and terminate at
		// trm.Len()+1 --OR-- simply +1 the index call as we
		// are doing here (seems easier). The reason for this
		// is because there is no valid ComparisonOperator
		// with an underlying uint8 value of zero (0). See
		// the ComparisonOperator constants for details.
		idx := i + 1
		cop, meth := trm.Index(idx)

		// execute method to create the targetrule
		rule := meth()

		// grab the raw string output
		fmt.Printf("[%d] %T instance [%s] execution returned %T: %s\n", idx, meth, cop.Context(), rule, rule)
	}
	// Output:
	// [1] aci.TargetRuleMethod instance [Eq] execution returned aci.TargetRule: ( targetfilter = "(&(objectClass=employee)(terminated=FALSE))" )
	// [2] aci.TargetRuleMethod instance [Ne] execution returned aci.TargetRule: ( targetfilter != "(&(objectClass=employee)(terminated=FALSE))" )
}

/*
This example demonstrates the creation of a TargetRule in a completely manual way. Users
will almost certainly want to use the (far easier) methods for Eq, Ne, etc., extended via
the very type instances intended for representation within a rule.
*/
func ExampleTR() {
	var rule TargetRule = TR(TargetScope, Eq, SingleLevel)
	fmt.Printf("%s", rule)
	// Output: ( targetscope = "onelevel" )
}

func ExampleTargetRules_Contains_alternative() {
	tdns := TRs(
		TDN(`uid=jesse,ou=People,dc=example,dc=com`),
		Subordinate.Eq(),
	)
	fmt.Printf("Contains: %t", tdns.Contains(TargetScope))
	// Output: Contains: true
}

func ExampleTargetRule_Compare() {
	tdn1 := TDN(`uid=jesse,ou=People,dc=example,dc=com`).Eq()
	tdn2 := TDN(`uid=jesse,ou=People,dc=example,dc=com`).Ne()
	fmt.Printf("Equal: %t", tdn1.Compare(tdn2))
	// Output: Equal: false
}

func ExampleTargetRules_Compare() {
	trs1 := TRs()
	tdn1 := TDN(`uid=jesse,ou=People,dc=example,dc=com`).Eq()
	tfdn1 := TFDN(`ou=People,dc=example,dc=com`).Eq()

	trs2 := TRs()
	tf2 := Filter(`(objectClass=*`).Eq()
	tsc2 := SingleLevel.Eq()

	trs1.Push(tdn1, tfdn1)
	trs2.Push(tf2, tsc2)

	fmt.Printf("Equal: %t", trs1.Compare(trs2))
	// Output: Equal: false
}

func ExampleTargetRule() {
	var tgt TargetRule = TR(
		TargetScope,
		Eq,
		Subordinate,
	)

	fmt.Printf("%s", tgt)
	// Output: ( targetscope = "subordinate" )
}

func ExampleTargetRule_String_traditional() {
	var tgt TargetRule = TR(
		TargetFrom,
		Ne,
		TFDN(`uid=*,ou=Contractors,ou=People,dc=example,dc=com`),
	)

	fmt.Printf("%s", tgt)
	// Output: ( target_from != "ldap:///uid=*,ou=Contractors,ou=People,dc=example,dc=com" )
}

func ExampleTargetRule_String_targetScope() {

	fmt.Printf("%s", BaseObject.Eq())
	// Output: ( targetscope = "base" )
}

func ExampleTargetRule_String_negatedTargetFilter() {

	fmt.Printf("%s", Filter(`(&(objectClass=*)(employeeStatus=ACTIVE))`).Ne())
	// Output: ( targetfilter != "(&(objectClass=*)(employeeStatus=ACTIVE))" )
}

func ExampleTargetRule_NoPadding() {
	f := `(&(objectClass=*)(employeeStatus=ACTIVE))`

	fmt.Printf("%s", Filter(f).Ne().NoPadding())
	// Output: (targetfilter!="(&(objectClass=*)(employeeStatus=ACTIVE))")
}

func ExampleTargetRule_SetKeyword() {
	var tgt TargetRule
	tgt.SetKeyword(TargetAttr)
	tgt.SetOperator(Ne)
	tgt.SetExpression(AT(`aci`))

	fmt.Printf("%s", tgt)
	// Output: ( targetattr != "aci" )
}

func ExampleTargetRule_SetOperator() {
	var tgt TargetRule
	tgt.SetKeyword(TargetAttr)
	tgt.SetOperator(Ne)
	tgt.SetExpression(AT(`aci`))

	fmt.Printf("%s", tgt)
	// Output: ( targetattr != "aci" )
}

func ExampleTargetRule_Operator() {
	cond := Filter(`(&(objectClass=*)(status=ACTIVE))`).Ne()
	fmt.Printf("%s", cond.Operator())
	// Output: !=
}

func ExampleTargetRule_SetExpression() {
	var tgt TargetRule
	tgt.SetKeyword(TargetAttr)
	tgt.SetOperator(Ne)
	tgt.SetExpression(AT(`aci`))

	fmt.Printf("%s", tgt)
	// Output: ( targetattr != "aci" )
}

func ExampleTargetRule_Category() {
	fmt.Printf("%s", BaseObject.Eq().Category())
	// Output: targetscope
}

func ExampleTargetRule_Keyword() {
	fmt.Printf("%s", BaseObject.Eq().Keyword())
	// Output: targetscope
}

func ExampleTargetRule_Kind() {
	var tr TargetRule
	fmt.Printf("%s", tr.Kind())
	// Output: condition
}

func ExampleTargetRule_ID() {
	fmt.Printf("%s", BaseObject.Eq().ID())
	// Output: target
}

func ExampleTargetRule_Expression() {
	tfdn := `uid=*,ou=Contractors,ou=People,dc=example,dc=com`
	fmt.Printf("%s", TFDN(tfdn).Eq().Expression())
	// Output: ldap:///uid=*,ou=Contractors,ou=People,dc=example,dc=com
}

func ExampleTargetRule_IsZero() {
	var tr TargetRule
	fmt.Printf("Zero: %t", tr.IsZero())
	// Output: Zero: true
}

func ExampleTargetRule_Valid() {
	var tr TargetRule
	fmt.Printf("Valid: %t", tr.Valid() == nil)
	// Output: Valid: false
}

func ExampleTargetRule_SetQuoteStyle() {
	var tgt TargetRule
	tgt.SetKeyword(Target)
	tgt.SetOperator(Ne)
	tgt.SetExpression(TDNs(
		TDN(`ldap:///uid=jesse,ou=People,dc=example,dc=com`),
		TDN(`ldap:///uid=courtney,ou=People,dc=example,dc=com`),
		TDN(`ldap:///uid=jimmy,ou=People,dc=example,dc=com`),
	))

	tgt.SetQuoteStyle(0)
	style1 := tgt.String()

	tgt.SetQuoteStyle(1)
	style2 := tgt.String()

	fmt.Printf("\n0: %s\n1: %s", style1, style2)
	// Output:
	// 0: ( target != "ldap:///uid=jesse,ou=People,dc=example,dc=com" || "ldap:///uid=courtney,ou=People,dc=example,dc=com" || "ldap:///uid=jimmy,ou=People,dc=example,dc=com" )
	// 1: ( target != "ldap:///uid=jesse,ou=People,dc=example,dc=com || ldap:///uid=courtney,ou=People,dc=example,dc=com || ldap:///uid=jimmy,ou=People,dc=example,dc=com" )
}

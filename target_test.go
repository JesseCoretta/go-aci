package aci

import (
	"fmt"
	"testing"
)

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

/*
This example demonstrates a similar scenario to the one described in the above example, but with
an alternative means of quotation demonstrated.
*/
func ExampleExtOps_alternativeQuotationScheme() {
	// Here we set double-quote encapsulation
	// upon the Rule instance created by the
	// ExtOps function.
	ext := ExtOps().Push(
		// These aren't real control OIDs.
		ExtOp(`1.3.6.1.4.1.56521.999.5`),
		ExtOp(`1.3.6.1.4.1.56521.999.6`),
		ExtOp(`1.3.6.1.4.1.56521.999.7`),
	)

	fmt.Printf("%s", ext.Eq())
	// Output: ( extop = "1.3.6.1.4.1.56521.999.5 || 1.3.6.1.4.1.56521.999.6 || 1.3.6.1.4.1.56521.999.7" )
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

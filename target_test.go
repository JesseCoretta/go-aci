package aci

import "testing"

func TestCtrls(t *testing.T) {
        L := Ctrls().Paren()
        o1 := Ctrl(`1.3.6.1.4.1.56521.101.2.1.1`)
        o2 := Ctrl(`1.3.6.1.4.1.56521.101.2.2.2`)
        o3 := Ctrl(`1.3.6.1.4.1.56521.101.3.1`)

        L.Push(o1, o2, o3)

        want := `( 1.3.6.1.4.1.56521.101.2.1.1 || 1.3.6.1.4.1.56521.101.2.2.2 || 1.3.6.1.4.1.56521.101.3.1 )`
        got := L.String()
        if want != got {
                t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
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
This example demonstrates how to craft a Target Control Rule using a list
of LDAP Control OIDs.
*/
func ExampleCtrls() {
        ctrls := Ctrls().Paren().Push(
                // These aren't real control OIDs.
                Ctrl(`1.3.6.1.4.1.56521.999.5`),
                Ctrl(`1.3.6.1.4.1.56521.999.6`),
                Ctrl(`1.3.6.1.4.1.56521.999.7`),
        )
        fmt.Printf("%s", ctrls.Eq())
        // Output: ( targetcontrol = "1.3.6.1.4.1.56521.999.5 || 1.3.6.1.4.1.56521.999.6 || 1.3.6.1.4.1.56521.999.7" )
}


/*
This example demonstrates how to craft a Target ExtOp Rule using a list
of LDAP Extended Operation OIDs.
*/
func ExampleExtOps() {
        ext := ExtOps().Paren().Push(
                // These aren't real control OIDs.
                ExtOp(`1.3.6.1.4.1.56521.999.5`),
                ExtOp(`1.3.6.1.4.1.56521.999.6`),
                ExtOp(`1.3.6.1.4.1.56521.999.7`),
        )
        fmt.Printf("%s", ext.Eq())
        // Output: ( extop = "1.3.6.1.4.1.56521.999.5 || 1.3.6.1.4.1.56521.999.6 || 1.3.6.1.4.1.56521.999.7" )
}

/*
This example demonstrates a similar scenario to the one described in the above example, but with
an alternative means of quotation demonstrated.
*/
func ExampleExtOps_alternativeQuotationScheme() {
        // Here we set double-quote encapsulation
        // upon the Rule instance created by the
        // ExtOps function.
        ext := ExtOps().Encap(`"`).Paren().Push(
                // These aren't real control OIDs.
                ExtOp(`1.3.6.1.4.1.56521.999.5`),
                ExtOp(`1.3.6.1.4.1.56521.999.6`),
                ExtOp(`1.3.6.1.4.1.56521.999.7`),
        )

        // Note we UNset encapsulation for the
        // Condition instance returned by Eq.
        fmt.Printf("%s", ext.Eq().Encap())
        // Output: ( extop = "1.3.6.1.4.1.56521.999.5" || "1.3.6.1.4.1.56521.999.6" || "1.3.6.1.4.1.56521.999.7" )
}

/*
This example demonstrates how to use a single Target DN to craft a Target Rule Equality
Condition.
*/
func ExampleDistinguishedName_Eq_target() {
        dn := TDN(`uid=jesse,ou=People,dc=example,dc=com`)
        fmt.Printf("%s", dn.Eq())
        // Output: ( target = "ldap:///uid=jesse,ou=People,dc=example,dc=com" )
}

/*
This example demonstrates how a list of Target DNs can be used to create a single Target
Rule. First, create a Rule using TDNs().Parens(), then push N desired TDN (Target DN)
values into the Rule.
*/
func ExampleRule_Eq_targetDNs() {
        tdns := TDNs().Paren().Push(
                TDN(`uid=jesse,ou=People,dc=example,dc=com`),
                TDN(`uid=courtney,ou=People,dc=example,dc=com`),
        )
        // Craft an equality Condition
        fmt.Printf("%s", tdns.Eq())
        // Output: ( target = "ldap:///uid=jesse,ou=People,dc=example,dc=com || ldap:///uid=courtney,ou=People,dc=example,dc=com" )
}


func TestAttrs_attrList(t *testing.T) {
        got := TAttrs().Paren().Push(
                ATName(`cn`),
                ATName(`sn`),
                ATName(`givenName`),
                ATName(`homeDirectory`),
                ATName(`uid`),
        )
        want := `( cn || sn || givenName || homeDirectory || uid )`

        if want != got.String() {
                t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
        }
}

/*
This example demonstrates how to create a Target Attributes Rule using a list of AttributeType instances.
*/
func ExampleTAttrs() {
        attrs := TAttrs().Push(
                ATName(`cn`),
                ATName(`sn`),
                ATName(`givenName`),
        )
        fmt.Printf("%s", attrs)
        // Output: cn || sn || givenName
}

/*
This example demonstrates how to create a Target Attributes Rule Equality Condition using a list of
AttributeType instances.
*/
func ExampleRule_Eq_targetAttr() {
        attrs := TAttrs().Paren().Push(
                ATName(`cn`),
                ATName(`sn`),
                ATName(`givenName`),
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
func ExampleTFilter() {
        tf := TFilter().Push(`(&(uid=jesse)(objectClass=*))`)
        fmt.Printf("%s", tf.Eq())
        // Output: ( targetfilter = "(&(uid=jesse)(objectClass=*))" )
}

/*
This example demonstrates how to craft a set of Target Rule Conditions.
*/
func ExampleT() {
        t := T().Push(
                TDN(`uid=jesse,ou=People,dc=example,dc=com`).Eq(),
                TFilter().Push(`(&(uid=jesse)(objectClass=*))`).Eq(),
                ExtOp(`1.3.6.1.4.1.56521.999.5`).Eq(),
        )
        fmt.Printf("%s", t)
        // Output: ( target = "ldap:///uid=jesse,ou=People,dc=example,dc=com" ) ( targetfilter = "(&(uid=jesse)(objectClass=*))" ) ( extop = "1.3.6.1.4.1.56521.999.5" )
}


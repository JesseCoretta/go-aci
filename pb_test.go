package aci

import (
	"fmt"
	"testing"
)

/*
This example demonstrates the creation of a PermissionBindRule using the PBR
package level function.
*/
func ExamplePBR() {

	// create a granting (allow)
	// permissive statement, in
	// which the read, search and
	// compare privileges are
	// bestowed.
	grant := Allow(
		`read`,
		`search`,
		`compare`,
	)

	// create a timeframe rule to reference.
	rule := Timeframe(
		ToD(`1400`), // notBefore
		ToD(`2300`), // notAfter
	)

	// Assemble the PermissionBindRule
	// using the above components. Note
	// that using the PBR function (as
	// opposed to assembling the struct
	// manually) automatically executes
	// the Valid method for us.
	pbr := PBR(grant, rule)

	fmt.Printf("%s", pbr)
	// Output: allow(read,search,compare) timeofday >= "1400" AND timeofday < "2300";

}

func ExamplePermissionBindRule_String() {
	pbr := PermissionBindRule{
		Allow(NoAccess),
		UDN(`uid=disgruntled_employee,ou=People,dc=example,dc=com`).Eq(),
	}

	fmt.Printf("%s", pbr)
	// Output: allow(none) userdn = "ldap:///uid=disgruntled_employee,ou=People,dc=example,dc=com";
}

func ExamplePermissionBindRule_Compare() {
	pbr1 := PermissionBindRule{
		Allow(NoAccess),
		UDN(`uid=disgruntled_employees,ou=Group,dc=example,dc=com`).Eq(),
	}

	pbr2 := PermissionBindRule{
		Allow(NoAccess),
		UDN(`uid=disgruntled_employee,ou=People,dc=example,dc=com`).Eq(),
	}

	fmt.Printf("%t", pbr1.Compare(pbr2))
	// Output: false
}

func ExamplePermissionBindRule_IsZero() {
	var pbr PermissionBindRule
	fmt.Printf("Zero: %t", pbr.IsZero())
	// Output: Zero: true
}

func ExamplePermissionBindRule_Valid() {
	var pbr PermissionBindRule
	fmt.Printf("Valid: %t", pbr.Valid() == nil)
	// Output: Valid: false
}

func ExamplePermissionBindRule_Kind() {
	var pbr PermissionBindRule
	fmt.Printf("%s", pbr.Kind())
	// Output: pbr
}

/*
This example demonstrates the creation of a PermissionBindRules instance using the PBRs
package level function.
*/
func ExamplePBRs() {
	rule1 := PermissionBindRule{
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	}

	rule2 := PermissionBindRule{
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	}

	// Init/Push in one shot
	pbrs := PBRs(rule1, rule2)

	fmt.Printf("%d %T instances found within %T", pbrs.Len(), rule1, pbrs)
	// Output: 2 aci.PermissionBindRule instances found within aci.PermissionBindRules
}

func ExamplePermissionBindRules_Kind() {
	var pbrs PermissionBindRules
	fmt.Printf("%s", pbrs.Kind())
	// Output: pbrs
}

func ExamplePermissionBindRules_IsZero() {
	var pbrs PermissionBindRules
	fmt.Printf("Zero: %t", pbrs.IsZero())
	// Output: Zero: true
}

func ExamplePermissionBindRules_Valid() {
	var pbrs PermissionBindRules
	fmt.Printf("Valid: %t", pbrs.Valid() == nil)
	// Output: Valid: false
}

func ExamplePermissionBindRules_Len() {
	var pbrs PermissionBindRules
	fmt.Printf("Length: %d", pbrs.Len())
	// Output: Length: 0
}

func ExamplePermissionBindRules_Push() {
	rule1 := PermissionBindRule{
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	}

	rule2 := PermissionBindRule{
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	}

	// Init/Push in one shot
	pbrs := PBRs()
	pbrs.Push(rule1)
	pbrs.Push(rule2)

	fmt.Printf("%d %T instances found within %T", pbrs.Len(), rule1, pbrs)
	// Output: 2 aci.PermissionBindRule instances found within aci.PermissionBindRules
}

func ExamplePermissionBindRules_Compare() {
	rule1 := PermissionBindRule{
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	}

	rule2 := PermissionBindRule{
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	}

	// Init/Push in one shot
	pbrs1 := PBRs()
	pbrs1.Push(rule1, rule2)

	rule1 = PermissionBindRule{
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=onboard_employees,ou=Groups,dc=example,dc=com`).Eq(),
	}

	rule2 = PermissionBindRule{
		Allow(AllAccess),
		UDN(`cn=Jesse Coretta,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	}

	pbrs2 := PBRs()
	pbrs2.Push(rule1, rule2)

	fmt.Printf("%t", pbrs1.Compare(pbrs2))
	// Output: false
}

func ExamplePermissionBindRules_Index() {
	rule1 := PermissionBindRule{
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	}

	rule2 := PermissionBindRule{
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	}

	// Init/Push in one shot
	pbrs := PBRs()
	pbrs.Push(rule1)
	pbrs.Push(rule2)

	for i := 0; i < pbrs.Len(); i++ {
		slice := pbrs.Index(i)
		fmt.Printf("%s\n", slice)
	}
	// Output:
	// deny(all,proxy) groupdn = "ldap:///cn=disgruntled_employees,ou=Groups,dc=example,dc=com";
	// allow(all) userdn = "ldap:///cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com";
}

func ExamplePermissionBindRules_Contains() {
	rule1 := PermissionBindRule{
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	}

	rule2 := PermissionBindRule{
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	}

	// Init/Push in one shot
	pbrs := PBRs(rule1, rule2)

	fmt.Printf("%T contains rule2: %t", pbrs, pbrs.Contains(rule2))
	// Output: aci.PermissionBindRules contains rule2: true
}

func ExamplePermissionBindRules_Pop() {
	rule1 := PermissionBindRule{
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	}

	rule2 := PermissionBindRule{
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	}

	// Init/Push in one shot
	pbrs := PBRs()
	pbrs.Push(rule1)
	pbrs.Push(rule2)

	// Pop the most recent (LIFO) slice
	// for interrogation. Note that this
	// REMOVED it from the above stack.
	popped := pbrs.Pop()

	fmt.Printf("Popped: %s", popped)
	// Output: Popped: allow(all) userdn = "ldap:///cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com";
}

func ExamplePermissionBindRules_String() {
	rule1 := PermissionBindRule{
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	}

	rule2 := PermissionBindRule{
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	}

	// Init/Push in one shot
	pbrs := PBRs()
	pbrs.Push(rule1)
	pbrs.Push(rule2)

	fmt.Printf("%s", pbrs)
	// Output: deny(all,proxy) groupdn = "ldap:///cn=disgruntled_employees,ou=Groups,dc=example,dc=com"; allow(all) userdn = "ldap:///cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com";
}

func TestPermissionBindRule_codecov(t *testing.T) {
	var pb PermissionBindRule
	_ = pb.IsZero()
	_ = pb.Valid()
	pb.B = And(UDN(``))   // bogus dn
	pb.P = Allow(`moddn`) // bogus right
	_ = pb.Valid()
}

func TestPermissionBindRules_codecov(t *testing.T) {
	var pbs PermissionBindRules
	_ = pbs.IsZero()
	_ = pbs.Valid()

	var pb *PermissionBindRule = new(PermissionBindRule)
	pb.B = And()
	pbs.Push(pb)
	_ = pbs.Len()
	_ = pbs.Valid()

	pbs = PBRs()

	rule0 := PermissionBindRule{
		P: Deny(AllAccess, ProxyAccess),
	}

	pbs.Push(rule0)

	rule1 := PermissionBindRule{
		B: GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	}

	pbs.Push(rule1)

	rule2 := PermissionBindRule{
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	}

	pbs.Push(rule2)
	if pbs.Len() != 1 {
		t.Errorf("%s failed: unexpected slice count, want '%d', got '%d'",
			t.Name(), 1, pbs.Len())
		return
	}
}

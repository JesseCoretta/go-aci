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
	var pbr PermissionBindRule = PBR(
		Allow(NoAccess),
		UDN(`uid=disgruntled_employee,ou=People,dc=example,dc=com`).Eq(),
	)

	fmt.Printf("%s", pbr)
	// Output: allow(none) userdn = "ldap:///uid=disgruntled_employee,ou=People,dc=example,dc=com";
}

func ExamplePermissionBindRule_Compare() {
	var pbr1 PermissionBindRule = PBR(
		Allow(NoAccess),
		UDN(`uid=disgruntled_employees,ou=Group,dc=example,dc=com`).Eq(),
	)

	var pbr2 PermissionBindRule = PBR(
		Allow(NoAccess),
		UDN(`uid=disgruntled_employee,ou=People,dc=example,dc=com`).Eq(),
	)

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
	var rule1 PermissionBindRule = PBR(
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	)

	var rule2 PermissionBindRule = PBR(
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	)

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
	var rule1 PermissionBindRule = PBR(
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	)

	var rule2 PermissionBindRule = PBR(
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	)

	// Init/Push in one shot
	pbrs := PBRs()
	pbrs.Push(rule1)
	pbrs.Push(rule2)

	fmt.Printf("%d %T instances found within %T", pbrs.Len(), rule1, pbrs)
	// Output: 2 aci.PermissionBindRule instances found within aci.PermissionBindRules
}

func ExamplePermissionBindRules_Compare() {
	var rule1 PermissionBindRule = PBR(
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	)

	var rule2 PermissionBindRule = PBR(
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	)

	// Init/Push in one shot
	pbrs1 := PBRs()
	pbrs1.Push(rule1, rule2)

	rule1 = PBR(
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=onboard_employees,ou=Groups,dc=example,dc=com`).Eq(),
	)

	rule2 = PBR(
		Allow(AllAccess),
		UDN(`cn=Jesse Coretta,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	)

	pbrs2 := PBRs()
	pbrs2.Push(rule1, rule2)

	fmt.Printf("%t", pbrs1.Compare(pbrs2))
	// Output: false
}

func ExamplePermissionBindRules_Index() {
	var rule1 PermissionBindRule = PBR(
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	)

	var rule2 PermissionBindRule = PBR(
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	)

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
	var rule1 PermissionBindRule = PBR(
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	)

	var rule2 PermissionBindRule = PBR(
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	)

	// Init/Push in one shot
	pbrs := PBRs(rule1, rule2)

	fmt.Printf("%T contains rule2: %t", pbrs, pbrs.Contains(rule2))
	// Output: aci.PermissionBindRules contains rule2: true
}

func ExamplePermissionBindRules_Pop() {
	var rule1 PermissionBindRule = PBR(
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	)

	var rule2 PermissionBindRule = PBR(
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	)

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
	var rule1 PermissionBindRule = PBR(
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	)

	var rule2 PermissionBindRule = PBR(
		Allow(AllAccess),
		UDN(`cn=Courtney Tolana,ou=Admin,ou=People,dc=example,dc=com`).Eq(),
	)

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
	pb = PBR(
		Allow(`moddn`), // bogus right
		And(UDN(``)),   // bogus dn
	)
	_ = pb.Valid()
}

func TestPermissionBindRules_codecov(t *testing.T) {
	var pbs PermissionBindRules
	_ = pbs.IsZero()
	_ = pbs.Valid()
	pbs.Push()
	pbs.Contains(``)
	pbs.Contains(nil)
	pbs.Contains('a')
	pbs.Push(nil, nil)
	pbs.Push(``)
	pbs.Contains(PermissionBindRule{})
	pbs.Push(PermissionBindRule{})
	pbs.Push(`fartknocker`)
	_ = pbs.pushPolicy()
	_ = pbs.pushPolicy('a')
	_ = pbs.pushPolicy(`baljfg`)
	_ = pbs.pushPolicy(nil, nil)

	var pb PermissionBindRule
	_ = pb.IsZero()
	_ = pb.Valid()
	pb = PBR(
		Permission{nil},
		And(),
	)
	pbs.Push(pb)
	pbs.Push()
	pbs.Contains(``)
	pbs.Contains(nil)
	pbs.Contains('a')
	pbs.Push(nil, nil)
	pbs.Push(``)
	pbs.Contains(PermissionBindRule{})
	pbs.Push(PermissionBindRule{})
	pbs.Push(`fartknocker`)
	_ = pb.IsZero()
	_ = pbs.Len()
	_ = pbs.Valid()

	pbs = PBRs()

	rule0 := PBR(Permission{nil}, nil)
	_ = rule0.IsZero()
	_ = rule0.Valid()

	pbs.Push(rule0)

	rule1 := PBR(Deny(ProxyAccess), nil)
	_ = rule1.IsZero()
	_ = rule1.Valid()
	pbs.Push(rule1)

	var rule2a PermissionBindRule
	var badbind BindContext = GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq()
	rule2a.Set(
		Permission{nil},
		badbind,
	)
	_ = rule2a.IsZero()
	_ = rule2a.Valid()
	pbs.Push(rule2a)

	var rule2b PermissionBindRule
	badbind = RDN(``).Ne()
	rule2b.Set(
		Allow(`read`, `search`, `compare`),
		badbind,
	)
	_ = rule2b.IsZero()
	_ = rule2b.Valid()
	pbs.Push(rule2b)

	var rule2c PermissionBindRule
	rule2c.Set(`pspspsppspspsp`)
	rule2c.Set(badBindDN, `pspspsppspspsp`)
	_ = rule2c.IsZero()
	_ = rule2c.Valid()
	pbs.Push(rule2c)
	pbs.Push(`pspsppspspspps`)

	rule3 := PBR(
		Permission{nil},
		nil,
	)
	_ = rule3.IsZero()
	_ = rule3.Valid()
	pbs.Push(rule3)

	rule4 := PBR(
		Permission{nil},
		BindRules(stackAnd().SetID(`bonedrule`)).Push(SSF(128).Eq()),
	)
	_ = rule4.IsZero()
	_ = rule4.Valid()
	pbs.Push(rule4, float32(1.234))

	var rule5 PermissionBindRule
	rule5.Set(
		Deny(AllAccess, ProxyAccess),
		GDN(`cn=disgruntled_employees,ou=Groups,dc=example,dc=com`).Eq(),
	)
	pbs.Push(rule5)

	if pbs.Len() != 1 {
		t.Errorf("%s failed: unexpected slice count, want '%d', got '%d'",
			t.Name(), 1, pbs.Len())
		return
	}

	pbs.Push(rule5.String())
	if pbs.Len() != 1 {
		t.Errorf("%s failed: %T allowed duplicate push",
			t.Name(), pbs)
		return
	}
	pbs.Pop()
	pbs.Push(rule5)
	pbs.Push(rule5.String())
	pbs.Contains(rule5)
	pbs.Contains(rule5.String())
}

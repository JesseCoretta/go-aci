package aci

import (
	"fmt"
	//"testing"
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

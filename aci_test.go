package aci

import (
	"fmt"
	"testing"
)

func TestACI(t *testing.T) {
	// Make a target rule that encompasses any account
	// with a DN syntax of "uid=<userid>,ou=People,dc=example,dc=com"
	C := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()

	// push into a new instance of Rule automatically
	// configured to store Target Rule Condition instances.
	tgt := T().Push(C)

	// define a timeframe for our PermissionBindRule
	// using two Condition instances
	notBefore := ToD(`1730`).Ge()                    // Condition: greater than or equal to time
	notAfter := ToD(`2400`).Lt()                     // Condition: less than time
	brule := And().Paren().Push(notBefore, notAfter) // our actual bind rule expression

	// Define the permission (rights).
	perms := Allow(ReadAccess, CompareAccess, SearchAccess)

	// Make our PermissionBindRule instance, which defines the
	// granting of access within a particular timeframe.
	pbrule := PB(perms, brule)

	// The ACI's effective name (should be unique within the directory)
	acl := `Limit people access to timeframe`

	// Finally, craft the Instruction instance
	aci := ACI(acl, tgt, pbrule)

	want := `(target = "ldap:///uid=*,ou=People,dc=example,dc=com")(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" );)`
	if want != aci.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, aci)
	}
}

func ExampleInstruction_build() {
	// Make a target rule that encompasses any account
	// with a DN syntax of "uid=<userid>,ou=People,dc=example,dc=com"
	C := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()

	// push into a new instance of Rule automatically
	// configured to store Target Rule Condition instances.
	tgt := T().Push(C)

	// define a timeframe for our PermissionBindRule
	// using two Condition instances
	notBefore := ToD(`1730`).Ge()                    // Condition: greater than or equal to time
	notAfter := ToD(`2400`).Lt()                     // Condition: less than time
	brule := And().Paren().Push(notBefore, notAfter) // our actual bind rule expression

	// Define the permission (rights).
	perms := Allow(ReadAccess, CompareAccess, SearchAccess)

	// Make our PermissionBindRule instance, which defines the
	// granting of access within a particular timeframe.
	pbrule := PB(perms, brule)

	// The ACI's effective name (should be unique within the directory)
	acl := `Limit people access to timeframe`

	// Finally, craft the Instruction instance
	var aci Instruction
	aci.Set(acl, tgt, pbrule)

	fmt.Printf("%s", aci)
	// Output: (target = "ldap:///uid=*,ou=People,dc=example,dc=com")(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" );)
}

func ExampleInstruction_buildNested() {
	// Make a target rule that encompasses any account
	// with a DN syntax of "uid=<userid>,ou=People,dc=example,dc=com"
	C := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()

	// push into a new instance of Rule automatically
	// configured to store Target Rule Condition instances.
	tgt := T().Push(C)

	ors := Or().Paren().Push(
		UDN(`uid=jesse,ou=admin,dc=example,dc=com`),
		UDN(`uid=courtney,ou=admin,dc=example,dc=com`),
	)

	nots := Not().Paren().Encap(`"`).Push(
		UAT(ATName(`terminated`), ATValue(`TRUE`)),
	)

	// define a timeframe for our PermissionBindRule
	// using two Condition instances
	brule := And().Paren().Push(
		And().Paren().Push(
			ToD(`1730`).Ge(), // Condition: greater than or equal to time
			ToD(`2400`).Lt(), // Condition: less than time
		),
		ors,
		nots,
	)

	// Define the permission (rights).
	perms := Allow(ReadAccess, CompareAccess, SearchAccess)

	// Make our PermissionBindRule instance, which defines the
	// granting of access within a particular timeframe.
	pbrule := PB(perms, brule)

	// The ACI's effective name (should be unique within the directory)
	acl := `Limit people access to timeframe`

	// Finally, craft the Instruction instance
	var aci Instruction
	aci.Set(acl, tgt, pbrule)

	fmt.Printf("%s", aci)
	// Output: (target = "ldap:///uid=*,ou=People,dc=example,dc=com")(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) (( timeofday >= "1730" AND timeofday < "2400" ) AND ( ldap:///uid=jesse,ou=admin,dc=example,dc=com OR ldap:///uid=courtney,ou=admin,dc=example,dc=com ) AND NOT ( "terminated#TRUE" ));)
}

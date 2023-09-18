package aci

import (
	"fmt"
	"testing"
)

func TestInstruction(t *testing.T) {
	var Ins Instruction
	_ = Ins.Valid()
	_ = Ins.IsZero()
	_ = Ins.String()

	// define a timeframe for our PermissionBindRule
	// using two Condition instances
	notBefore := ToD(`1730`).Ge()                    // Condition: greater than or equal to time
	notAfter := ToD(`2400`).Lt()                     // Condition: less than time
	brule := And().Paren().Push(notBefore, notAfter) // our actual bind rule expression

	// Define the permission (rights).
	perms := Allow(ReadAccess, CompareAccess, SearchAccess)

	// Make our PermissionBindRule instance, which defines the
	// granting of access within a particular timeframe.
	pbrule := PBR(perms, brule)

	// Make a target rule that encompasses any account
	// with a DN syntax of "uid=<userid>,ou=People,dc=example,dc=com"
	Tar := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()

	badACL := ``

	ACI(badACL, Tar, pbrule)
}

func TestACIs(t *testing.T) {
	var Ins Instructions
	_ = Ins.Valid()
	_ = Ins.Push()
	_ = Ins.Push(nil)
	_ = Ins.Push(Instruction{})
	_ = Ins.Push(``)
	_ = Ins.Push('a')
	_ = Ins.IsZero()
	_ = Ins.String()
	_ = Ins.Len()
	_ = Ins.Index(0)

	// Make a target rule that encompasses any account
	// with a DN syntax of "uid=<userid>,ou=People,dc=example,dc=com"
	C := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()

	// push into a new instance of Rule automatically
	// configured to store Target Rule Condition instances.
	tgt := TRs(C)

	// define a timeframe for our PermissionBindRule
	// using two Condition instances
	notBefore := ToD(`1730`).Ge()                    // Condition: greater than or equal to time
	notAfter := ToD(`2400`).Lt()                     // Condition: less than time
	brule := And().Paren().Push(notBefore, notAfter) // our actual bind rule expression

	// Define the permission (rights).
	perms := Allow(ReadAccess, CompareAccess, SearchAccess)

	// Make our PermissionBindRule instance, which defines the
	// granting of access within a particular timeframe.
	pbrule := PBR(perms, brule)

	// The ACI's effective name (should be unique within the directory)
	acl := `Limit people access to timeframe`

	// Finally, craft the Instruction instance
	var i Instruction
	_ = i.TRs()
	_ = i.PBRs()
	_ = i.ACL()
	_ = i.Valid()
	_ = i.String()

	i = ACI(acl, tgt, pbrule)
	i.Set(tgt)
	_ = i.TRs()
	_ = i.PBRs()
	_ = i.ACL()

	want := `( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" );)`
	if want != i.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, i)
		return
	}

	Ins = ACIs()
	Ins.Push(i)
	popped := Ins.Pop()
	Ins.Push(popped)
	Ins.F()
	Ins.Push(popped.String())
	Ins.Push(`<3 <3 <3`)
	if Ins.Len() != 1 {
		t.Errorf("%s failed to push %T into %T, len:%d, want:%d\n%s", t.Name(), i, Ins, Ins.Len(), 1, Ins)
		return
	}

	if Ins.String() != Ins.Index(0).String() {
		t.Errorf("%s strcmp fail", t.Name())
		return
	}
}

func ExampleInstruction_build() {
	// Make a target rule that encompasses any account
	// with a DN syntax of "uid=<userid>,ou=People,dc=example,dc=com"
	C := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()

	// push into a new instance of Rule automatically
	// configured to store Target Rule Condition instances.
	tgt := TRs(C).Push(C)

	// define a timeframe for our PermissionBindRule
	// using two Condition instances
	notBefore := ToD(`1730`).Ge()                    // Condition: greater than or equal to time
	notAfter := ToD(`2400`).Lt()                     // Condition: less than time
	brule := And().Paren().Push(notBefore, notAfter) // our actual bind rule expression

	// Define the permission (rights).
	perms := Allow(ReadAccess, CompareAccess, SearchAccess)

	// Make our PermissionBindRule instance, which defines the
	// granting of access within a particular timeframe.
	pbrule := PBR(perms, brule)

	// The ACI's effective name (should be unique within the directory)
	acl := `Limit people access to timeframe`

	// Finally, craft the Instruction instance
	var i Instruction
	i.Set(acl, tgt, pbrule)

	fmt.Printf("%s", i)
	// Output: ( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" );)
}

func ExampleInstruction_buildNested() {
	// Make a target rule that encompasses any account
	// with a DN syntax of "uid=<userid>,ou=People,dc=example,dc=com"
	C := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()

	// push into a new instance of Rule automatically
	// configured to store Target Rule Condition instances.
	tgt := TRs().Push(C)

	// create an ORed stack, pushing the two specified
	// userdn equality conditions into its collection.
	ors := Or().Paren().Push(
		UDN(`uid=jesse,ou=admin,dc=example,dc=com`).Eq(),
		UDN(`uid=courtney,ou=admin,dc=example,dc=com`).Eq(),
	)

	// create our AttributeBindTypeOrValue instance,
	// setting the AttributeType as `ninja`, and the
	// AttributeValue as `FALSE`
	attr := AT(`ninja`)       // attributeType
	aval := AV(`FALSE`)       // attributeValue
	userat := UAT(attr, aval) // attributeBindTypeOrValue

	// create a negated (NOT) stack, pushing
	// our AttributeBindTypeOrValue BindRule
	// (Eq()) instance into its collection.
	// Also, make stack parenthetical.
	nots := Not().Paren().Push(userat.Eq())

	// define a timeframe for our PermissionBindRule
	// using two Condition instances. Make both AND
	// stacks parenthetical, and push our OR and NOT
	// stacks defined above.
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
	pbr := PBR(perms, brule)

	// The ACI's effective name (should be unique within the directory)
	acl := `Limit people access to timeframe`

	// Finally, craft the Instruction instance
	var i Instruction
	i.Set(acl, tgt, pbr)

	fmt.Printf("%s", i)
	// Output: ( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( ( timeofday >= "1730" AND timeofday < "2400" ) AND ( userdn = "ldap:///uid=jesse,ou=admin,dc=example,dc=com" OR userdn = "ldap:///uid=courtney,ou=admin,dc=example,dc=com" ) AND NOT ( userattr = "ninja#FALSE" ) );)
}

/*
This example demonstrates doing a literal search for an ACI within a stack of ACIs
using its Contains method. Case is not significant in the matching process.
*/
func ExampleInstructions_Contains() {
	raw1 := `( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe for those ninjas"; allow(read,search,compare) ( ( timeofday >= "1730" AND timeofday < "2400" ) AND ( userdn = "ldap:///uid=jesse,ou=admin,dc=example,dc=com" OR userdn = "ldap:///uid=courtney,ou=admin,dc=example,dc=com" ) AND NOT ( userattr = "ninja#FALSE" ) );)`
	raw2 := `( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" );)`

	acis := ACIs(
		raw1,
		raw2,
	)

	fmt.Printf("%T contains raw1: %t", acis, acis.Contains(raw1))
	// Output: aci.Instructions contains raw1: true
}

/*
This example demonstrates use of the F method to obtain the
package-level function appropriate for the creation of new
stack elements.
*/
func ExampleInstructions_F() {
	var acis Instructions
	funk := acis.F()

	ins := funk() // normally you'd want to supply some type instances
	fmt.Printf("%T", ins)
	// Output: aci.Instruction

}

func ExampleACIs() {
	raw1 := `( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe for those ninjas"; allow(read,search,compare) ( ( timeofday >= "1730" AND timeofday < "2400" ) AND ( userdn = "ldap:///uid=jesse,ou=admin,dc=example,dc=com" OR userdn = "ldap:///uid=courtney,ou=admin,dc=example,dc=com" ) AND NOT ( userattr = "ninja#FALSE" ) );)`
	raw2 := `( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" );)`

	acis := ACIs(
		raw1,
		raw2,
	)

	fmt.Printf("%T contains %d Instruction instances", acis, acis.Len())
	// Output: aci.Instructions contains 2 Instruction instances
}

/*
This example demonstrates use of the ACI package-level function. An instance
of Instruction is created using manually assembled type instances.
*/
func ExampleACI() {
	t := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()
	tgt := TRs().Push(t)

	userat := UAT(
		AT(`ninja`),
		AV(`FALSE`),
	)

	ors := Or().Paren().Push(
		UDN(`uid=jesse,ou=admin,dc=example,dc=com`).Eq(),
		UDN(`uid=courtney,ou=admin,dc=example,dc=com`).Eq(),
	)

	nots := Not().Paren().Push(userat.Eq())
	brule := And().Paren().Push(
		Timeframe(
			ToD(`1730`),
			ToD(`2400`),
		).Paren(),
		ors,
		nots,
	)

	perms := Allow(
		ReadAccess,
		SearchAccess,
		CompareAccess,
	)

	pbr := PBR(perms, brule)
	acl := `Limit people access to timeframe`

	var i Instruction
	i.Set(acl, tgt, pbr)

	fmt.Printf("%s", i)
	// Output: ( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( ( timeofday >= "1730" AND timeofday < "2400" ) AND ( userdn = "ldap:///uid=jesse,ou=admin,dc=example,dc=com" OR userdn = "ldap:///uid=courtney,ou=admin,dc=example,dc=com" ) AND NOT ( userattr = "ninja#FALSE" ) );)
}

func ExampleInstruction_String() {
	t := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()
	tgt := TRs().Push(t)

	userat := UAT(AT(`ninja`), AV(`FALSE`))
	ors := Or().Paren().Push(
		UDN(`uid=jesse,ou=admin,dc=example,dc=com`).Eq(),
		UDN(`uid=courtney,ou=admin,dc=example,dc=com`).Eq(),
	)
	nots := Not().Paren().Push(userat.Eq())
	brule := And().Paren().Push(
		Timeframe(
			ToD(`1730`),
			ToD(`2400`),
		).Paren(),
		ors,
		nots,
	)

	perms := Allow(
		ReadAccess,
		SearchAccess,
		CompareAccess,
	)

	pbr := PBR(perms, brule)
	acl := `Limit people access to timeframe`

	var i Instruction
	i.Set(acl, tgt, pbr)

	fmt.Printf("%s", i)
	// Output: ( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( ( timeofday >= "1730" AND timeofday < "2400" ) AND ( userdn = "ldap:///uid=jesse,ou=admin,dc=example,dc=com" OR userdn = "ldap:///uid=courtney,ou=admin,dc=example,dc=com" ) AND NOT ( userattr = "ninja#FALSE" ) );)
}

func ExampleInstruction_TRs() {
	t := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()
	tgt := TRs().Push(t)

	userat := UAT(AT(`ninja`), AV(`FALSE`))
	ors := Or().Paren().Push(
		UDN(`uid=jesse,ou=admin,dc=example,dc=com`).Eq(),
		UDN(`uid=courtney,ou=admin,dc=example,dc=com`).Eq(),
	)
	nots := Not().Paren().Push(userat.Eq())
	brule := And().Paren().Push(
		Timeframe(
			ToD(`1730`),
			ToD(`2400`),
		).Paren(),
		ors,
		nots,
	)

	perms := Allow(
		ReadAccess,
		SearchAccess,
		CompareAccess,
	)

	pbr := PBR(perms, brule)
	acl := `Limit people access to timeframe`

	var i Instruction
	i.Set(acl, tgt, pbr)

	fmt.Printf("%s", i.TRs())
	// Output: ( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )
}

func ExampleInstruction_PBRs() {
	t := TDN(`uid=*,ou=People,dc=example,dc=com`).Eq()
	tgt := TRs().Push(t)

	userat := UAT(AT(`ninja`), AV(`FALSE`))
	ors := Or().Paren().Push(
		UDN(`uid=jesse,ou=admin,dc=example,dc=com`).Eq(),
		UDN(`uid=courtney,ou=admin,dc=example,dc=com`).Eq(),
	)
	nots := Not().Paren().Push(userat.Eq())
	brule := And().Paren().Push(
		Timeframe(
			ToD(`1730`),
			ToD(`2400`),
		).Paren(),
		ors,
		nots,
	)

	perms := Allow(
		ReadAccess,
		SearchAccess,
		CompareAccess,
	)

	pbr := PBR(perms, brule)
	acl := `Limit people access to timeframe`

	var i Instruction
	i.Set(acl, tgt, pbr)

	fmt.Printf("%s", i.PBRs())
	// Output: allow(read,search,compare) ( ( timeofday >= "1730" AND timeofday < "2400" ) AND ( userdn = "ldap:///uid=jesse,ou=admin,dc=example,dc=com" OR userdn = "ldap:///uid=courtney,ou=admin,dc=example,dc=com" ) AND NOT ( userattr = "ninja#FALSE" ) );
}

func ExampleInstruction_ACL() {
	var i Instruction
	acl := `This is an access control label`
	i.Set(acl)
	fmt.Printf("%s", i.ACL())
	// Output: This is an access control label
}

func ExampleInstruction_IsZero() {
	var i Instruction
	fmt.Printf("Zero: %t", i.IsZero())
	// Output: Zero: true
}

func ExampleInstruction_Valid() {
	var i Instruction
	fmt.Printf("Valid: %t", i.Valid() == nil)
	// Output: Valid: false
}

func ExampleInstruction_Set() {
	var i Instruction
	acl := `This is an access control label`
	i.Set(acl)
	fmt.Printf("%s", i.ACL())
	// Output: This is an access control label
}

func ExampleInstructions_IsZero() {
	var i Instructions
	fmt.Printf("Zero: %t", i.IsZero())
	// Output: Zero: true
}

func ExampleInstructions_Len() {
	var i Instructions
	fmt.Printf("Length: %d", i.Len())
	// Output: Length: 0
}

func ExampleInstructions_Valid() {
	var i Instructions
	fmt.Printf("Valid: %t", i.Valid() == nil)
	// Output: Valid: false
}

func ExampleInstructions_Pop() {
	raw := []string{
		`( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" );)`,
		`( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) ); deny(proxy,selfwrite) ( userdn = "ldap:///all" );)`,
	}

	ins := ACIs()
	for i := 0; i < len(raw); i++ {
		ins.Push(raw[i])
	}

	popped := ins.Pop()
	fmt.Printf("%s", popped)
	// Output: ( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) ); deny(selfwrite,proxy) ( userdn = "ldap:///all" );)
}

func ExampleInstructions_Push() {
	raw := []string{
		`( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" ); )`,
		`( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) ); deny(proxy,selfwrite) ( userdn = "ldap:///all" ); )`,
	}

	ins := ACIs()
	for i := 0; i < len(raw); i++ {
		ins.Push(raw[i])
	}
	fmt.Printf("Length: %d", ins.Len())
	// Output: Length: 2
}

func ExampleInstructions_Index() {
	raw := []string{
		`( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" );)`,
		`( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) ); deny(selfwrite,proxy) ( userdn = "ldap:///all" );)`,
	}

	ins := ACIs()
	for i := 0; i < len(raw); i++ {
		ins.Push(raw[i])
	}

	fmt.Printf("%s", ins.Index(1))
	// Output: ( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) ); deny(selfwrite,proxy) ( userdn = "ldap:///all" );)
}

func ExampleInstructions_String() {
	raw := []string{
		`( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" ); )`,
		`( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) ); deny(proxy,selfwrite) ( userdn = "ldap:///all" ); )`,
	}

	ins := ACIs()
	for i := 0; i < len(raw); i++ {
		ins.Push(raw[i])
	}

	fmt.Printf("%s", ins)
	// Output:
	// ( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" );)
	// ( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) ); deny(selfwrite,proxy) ( userdn = "ldap:///all" );)
}

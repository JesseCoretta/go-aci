package aci

import (
	"fmt"
	"testing"
)

func TestInheritance(t *testing.T) {
	inh := Inherit(UAT(AT(`manager`), USERDN), Level0, Level1, Level2, Level8)
	want := `userattr = "parent[0,1,2,8].manager#USERDN"`
	got := inh.Eq()
	if want != got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}

func ExampleInherit_uSERDN() {
	attr := AT(`manager`)
	uat := UAT(attr, AV(`USERDN`))
	inh := Inherit(uat, 0, 1, 2, 8)
	fmt.Printf("%s", inh.Eq())
	// Output: userattr = "parent[0,1,2,8].manager#USERDN"
}

func ExampleInherit_uAT() {
	attr := AT(`manager`)
	uat := UAT(attr, AV(`uid=frank,ou=People,dc=example,dc=com`))
	inh := Inherit(uat, 1, 3)
	fmt.Printf("%s", inh.Eq())
	// Output: userattr = "parent[1,3].manager#uid=frank,ou=People,dc=example,dc=com"
}

func ExampleInherit_gAT() {
	attr := AT(`owner`)
	gat := GAT(attr, USERDN)
	inh := Inherit(gat, 3, 4)
	fmt.Printf("%s", inh.Eq())
	// Output: groupattr = "parent[3,4].owner#USERDN"
}

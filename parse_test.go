package aci

import (
	_ "fmt"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	want := ` (target = "ldap:///ipatokenuniqueid=*,cn=otp,dc=example,dc=com")(targetfilter = "(objectClass=ipaToken)")(version 3.0;acl "token-add-delete"; allow (add) userattr = "ipatokenOwner#SELFDN";)`
	//got, err := parseInstruction(want)
	//if err != nil {
		//t.Errorf("%s failed: %v", t.Name(), err)
	//}

	if want != `` { //got.String() {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(),want,``)
	}

}

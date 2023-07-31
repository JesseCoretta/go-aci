package aci

import "testing"

/*
func TestParseFilter(t *testing.T) {
	// this is an illogical but legal filter,
	// and aims to provide as diverse a test
	// as possible.
	want := `(&(objectClass=employee)(objectClass=engineering)(|(&(team>=5)(givenName~=Dave)(givenName:caseExactMatch:=John)(:dn:2.5.13.5:=Jesse)(:caseExactMatch:=Courtney)(color;lang-fr=bleu)(string=rr*a))(cn=Jesse Coretta)(:1.3.6.1.4.1.123456:=Value)(team=B))(!(terminated=TRUE)(dumbass=TRUE)))`

	r, err := parseFilter(want)
	if err != nil {
		t.Errorf("%s failed: %v", t.Name(), err)
	}

	if got := r.String(); want != got {
		t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, got)
	}
}
*/

var testInstructions []string = []string{
	`(targetfilter = "(&(objectClass=employee)(objectClass=engineering))")(targetscope = "onelevel")(version 3.0; acl "Allow anonymous onelevel searches for engineering employees"; allow(read,search,compare) userdn = "ldap:///anyone";)`,
	`(version 3.0; acl "Allow read and compare for anyone using less than 128 SSF"; allow(read,compare) userdn = "ldap:///anyone" AND ssf < "128";)`,
	`(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF"; allow(read,write) ( userdn = "ldap:///anyone" AND ssf >= "128" ) AND NOT dayofweek = "Fri";)`,
	`(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF"; allow(read,write) (( userdn = "ldap:///anyone" AND ssf >= "128" ) AND NOT dayofweek = "Fri");)`,
	`(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF"; allow(read,write) userdn = "ldap:///anyone" AND ssf >= "128" AND NOT dayofweek = "Fri";)`,
	`(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF"; allow(read,write) (( (userdn = "ldap:///anyone") AND (ssf >= "128") ) AND NOT (dayofweek = "Fri"));)`,
	`(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF"; allow(read,write) userdn = "ldap:///anyone" AND ssf >= "128" AND NOT ( dayofweek = "Fri" OR dayofweek = "Sun" );)`,
}

func TestParseInstruction(t *testing.T) {
	for i := 0; i < len(testInstructions); i++ {
		want := testInstructions[i]
		//t.Logf("WANT [%d]: %s\n", i, want)
		a, err := parseInstruction(want)
		if err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
		}

		if got := a.String(); want != got {
			t.Errorf("%s failed [testInstructions:%d]: want '%s', got '%s'", t.Name(), i, want, got)
		}
		t.Logf("%s\n", a)
	}
}

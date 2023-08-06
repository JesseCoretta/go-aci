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
	`( target = "ldap:///ipatokenuniqueid=*,cn=otp,dc=example,dc=com" ) ( targetfilter = "(objectClass=ipaToken)" )(version 3.0; acl "token-add-delete"; allow(add) userattr = "ipatokenOwner#SELFDN";)`,
	`( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" ) ( targetscope = "onelevel" )(version 3.0; acl "Allow anonymous onelevel searches for engineering employees"; allow(read,search,compare) userdn = "ldap:///anyone";)`,
	`( targetattr = "*" )(version 3.0; acl "Allow read,search "; allow(read,search) ( userattr = "aciurl#LDAPURL" );)`,
	`( targetattr = "*" )(version 3.0; acl "Allow anonymous onelevel searches for engineering employees"; allow(read,search,compare) ( userdn = "ldap:///anyone || ldap:///parent || ldap:///self" );)`,
	`( targetattr = "*" )(version 3.0; acl "Profile access"; allow(read,search) userattr = "parent[0,1].owner#USERDN";)`,
	`( target = "ldap:///ou=Groups,dc=subdomain1,dc=hostedCompany1,dc=example,dc=com" ) ( targetattr = "*" )(version 3.0; acl "Domain access"; allow(read,search) groupdn = "ldap:///cn=DomainAdmins,ou=Groups,dc=subdomain1,dc=hostedCompany1,dc=example,dc=com";)`,
	`( targetattr = "ou || cn" ) ( targetfilter = "(ou=Engineering)" )(version 3.0; acl "Allow uid=user to search and read engineering attributes"; allow(read,search) ( userdn = "ldap:///uid=user,ou=People,dc=example,dc.com" );)`,
	`( targetattr = "*" )(version 3.0; acl "Deny 192.0.2.0/24"; deny(all) ( userdn = "ldap:///anyone" ) AND ( ip != "192.0.2." );)`,
	`( target = "ldap:///ou=People,dc=example,dc=com" )(version 3.0; acl "Allow users to read and search attributes of own entry"; allow(read,search) ( userdn = "ldap:///self" );)`,
	`( targetattr = "*" )(version 3.0; acl "Deny 2001:db8::/64"; deny(all) ( userdn = "ldap:///anyone" ) AND ( ip != "2001:db8::" );)`,
	`( targetattr = "*" )(version 3.0; acl "Deny client.example.com"; deny(all) ( userdn = "ldap:///anyone" ) AND ( dns != "client.example.com" );)`,
	`( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" ) ( targetcontrol = "1.3.6.1.4.1.56521.999.5 || 1.3.6.1.4.1.56521.999.6" ) ( targetscope = "onelevel" )(version 3.0; acl "Allow anonymous onelevel searches for engineering employees"; allow(read,search,compare) ( userdn = "ldap:///anyone" );)`,
	`( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" ) ( targetcontrol = "1.3.6.1.4.1.56521.999.5" || "1.3.6.1.4.1.56521.999.6" ) ( targetscope = "onelevel" )(version 3.0; acl "Allow anonymous onelevel searches for engineering employees; alt. quotation"; allow(read,search,compare) ( userdn = "ldap:///anyone" );)`,
	`( target = "ldap:///anyone" || "ldap:///uid=jesse,ou=People,dc=example.com" ) ( targetscope = "subordinate" )(version 3.0; acl "Allow subordinate searches of any account by authenticated users"; allow(read,search,compare) ( userdn = "ldap:///all" );)`,
	`( targetattr = "*" )(version 3.0; acl "Deny example.com"; deny(all) ( userdn = "ldap:///anyone" ) AND ( dns != "*.example.com" );)`,
	`(version 3.0; acl "Allow read and compare for anyone using less than 128 SSF"; allow(read,compare) userdn = "ldap:///anyone" AND ssf < "128";)`,
	`( target = "ldap:///ou=Groups,($dn),dc=example,dc=com" ) ( targetattr = "*" )(version 3.0; acl "Domain access"; allow(read,search) groupdn = "ldap:///cn=DomainAdmins,ou=Groups,[$dn],dc=example,dc=com";)`,
	`( targetattr = "userPassword" )(version 3.0; acl "Allow users updating own userPassword"; allow(write) ( userdn = "ldap:///self" ) AND ( ssf >= "128" );)`,
	`( targetfilter = "(|(department=Engineering)(department=Sales)" )(version 3.0; acl "Allow HR updating engineering and sales entries"; allow(write) ( groupdn = "ldap:///cn=Human Resources,dc=example,dc.com" );)`,
	`(version 3.0; acl "Deny access on Saturdays and Sundays"; deny(all) ( userdn = "ldap:///uid=user,ou=People,dc=example,dc=com" ) AND ( dayofweek = "Sun,Sat" );)`,
	`(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF everyday EXCEPT Friday"; allow(read,write) ( ( userdn = "ldap:///anyone" AND ssf >= "128" ) AND NOT dayofweek = "Fri" );)`,

	// ANTLR err ?
	//`(version 3.0; acl "Deny access between 6pm and 0am"; deny(all) ( userdn = "ldap:///uid=user,ou=People,dc=example,dc=com" ) AND ( timeofday >= "1800" AND timeofday < "2400" );)`,

	// parenthetical nesting issue
	`(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" ) );)`,

	`(version 3.0; acl "Deny all access without certificate"; deny(all) ( authmethod = "NONE" OR authmethod = "SIMPLE" );)`,
	`( target = "ldap:///cn=*,ou=people,dc=example,dc=com" )(version 3.0; acl "Deny modrdn rights to the example group"; deny(write) groupdn = "ldap:///cn=example,ou=groups,dc=example,dc=com";)`,
	`( target = "ldap:///ou=*,($dn),dc=example,dc=com" ) ( targetattr = "*" )(version 3.0; acl "Domain access"; allow(read,search) groupdn = "ldap:///cn=DomainAdmins,ou=Groups,($dn),dc=example,dc=com";)`,
	`( targetattr = "manager" )(version 3.0; acl "Allow manager role to update manager attribute"; allow(read,search) roledn = "ldap:///cn=Human Resources,ou=People,dc=example,dc=com";)`,
	`( target = "ldap:///ou=People,dc=example,dc=com" )(version 3.0; acl "Allow members of administrators and operators group to manage users"; allow(read,write,add,delete,search) groupdn = "ldap:///cn=Administrators,ou=Groups,dc=example,com" AND groupdn = "ldap:///cn=Operators,ou=Groups,dc=example,com";)`,
	`( target = "ldap:///ou=Accounting,dc=example,dc=com" ) ( targetattr = "*" )(version 3.0; acl "test acl"; allow(read,search,compare) ( userdn = "ldap:///anyone" );)`,
	`( targetattr = "sn" || "givenName" || "telephoneNumber" )(version 3.0; acl "Anonymous read, search for names and phone numbers"; allow(read,search) userdn = "ldap:///anyone";)`,
	`( targetattr = "member" )(version 3.0; acl "Allow users to add/remove themselves from example group"; allow(selfwrite) userdn = "ldap:///all";)`,
	`( targetattr = "userPassword" )(version 3.0; acl "Allow users updating their password"; allow(write) userdn = "ldap:///self";)`,
	`( targetattr = "manager" )(version 3.0; acl "Allow cn=user to update manager attributes"; allow(write) userdn = "ldap:///parent";)`,
	`( targetattr = "manager" )(version 3.0; acl "Allow example group to read manager attribute"; allow(read,search) groupdn = "ldap:///cn=example,ou=Groups,dc=example,dc=com";)`,
	`( targetattr = "manager" )(version 3.0; acl "Allow uid=admin reading manager attribute"; allow(read,search) userdn = "ldap:///uid=admin,ou=People,dc=example,dc=com";)`,
	`( targetattr = "homePostalAddress" )(version 3.0; acl "Allow HR setting homePostalAddress"; allow(write) userdn = "ldap:///ou=People,dc=example,dc=com??sub?(department=Human Resources)";)`,


	`( targetattr = "userPassword" )(version 3.0; acl "Allow users updating own userPassword"; allow(write) ( userdn = "ldap:///self" );)`,
	`(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - no nesting"; allow(read,write) userdn = "ldap:///anyone" AND ssf >= "128" AND NOT dayofweek = "Fri";)`,
	`( target = "ldap:///ou=People,dc=example,dc=com" )(version 3.0; acl "Allow users to read and search attributes of own entry"; allow(read,search) ( userdn = "ldap:///self" );)`,
	`( targetattr = "homePostalAddress" )(version 3.0; acl "Allow manager=example setting homePostalAddress"; allow(write) userdn = "ldap:///dc=example,dc=com??sub?(manager=example)";)`,
	`( targetattr = "telephoneNumber" )(version 3.0; acl "Manager: telephoneNumber"; allow(all) userattr = "manager#USERDN";)`,
	`( target_from = "ldap:///uid=*,cn=staging,dc=example,dc=com" ) ( target_to = "ldap:///cn=People,dc=example,dc=com" )(version 3.0; acl "Allow modrdn by user"; allow(write) userdn = "ldap:///uid=user,dc=example,dc=com";)`,
	`(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - not/or nesting"; allow(read,write) userdn = "ldap:///anyone" AND ssf >= "128" AND NOT ( dayofweek = "Fri" OR dayofweek = "Sun" );)`,
}

func TestParseInstruction(t *testing.T) {
	for i := 0; i < len(testInstructions); i++ {
		want := testInstructions[i]

		a, err := parseInstruction(want)
		if err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
		}

		if got := a.String(); want != got {
			t.Errorf("%s failed [testInstructions:%d]:\nwant '%s'\ngot  '%s'", t.Name(), i, want, got)
		}
	}
}

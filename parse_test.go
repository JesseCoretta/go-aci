package aci

import (
	_ "fmt"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	want := `(target = "ldap:///ipatokenuniqueid=*,cn=otp,dc=example,dc=com")(targetfilter = "(objectClass=ipaToken)")(version 3.0; acl "token-add-delete"; allow(add) userattr = "ipatokenOwner#SELFDN";)`
        got, _, err := parseRule(want[154:])
	if err != nil {
		t.Errorf("%s failed: %v", t.Name(), err)
	}

        if want != got.String() {
                t.Errorf("%s failed:\nwant '%s' (%d)\ngot  '%s' (%d)",
			t.Name(),
			want,
			len(want),
			got,
			len(got.String()))
        }

}

func TestParseInstruction_nest(t *testing.T) {
        want := `(target = "ldap:///ipatokenuniqueid=*,cn=otp,dc=example,dc=com")(targetfilter = "(objectClass=ipaToken)")(version 3.0; acl "token-add-delete"; allow(add) userattr = "ipatokenOwner#SELFDN" AND userdn = "ldap:///uid=jesse,ou=People,dc=example,dc=com";)`
        got, _, err := parseRule(want[154:])
        if err != nil {
                t.Errorf("%s failed: %v", t.Name(), err)
        }

        if want != got.String() {
                t.Errorf("%s failed:\nwant '%s' (%d)\ngot  '%s' (%d)",
			t.Name(),
			want,
			len(want),
			got,
			len(got.String()))
        }

}

func TestParseInstruction_nest2(t *testing.T) {
        want := `(target = "ldap:///ipatokenuniqueid=*,cn=otp,dc=example,dc=com")(targetfilter = "(objectClass=ipaToken)")(version 3.0; acl "token-add-delete"; allow(add) userattr = "ipatokenOwner#SELFDN" AND ( userdn = "ldap:///uid=jesse,ou=People,dc=example,dc=com" OR userdn = "ldap:///uid=courtney,ou=People,dc=example,dc=com" ) AND NOT userattr = "terminated:TRUE";)`
        got, _, err := parseRule(want[154:])
        if err != nil {
                t.Errorf("%s failed: %v", t.Name(), err)
        }

        if want != got.String() {
                t.Errorf("%s failed:\nwant '%s' (%d)\ngot  '%s' (%d)",
			t.Name(),
			want,
			len(want),
			got,
			len(got.String()))
        }

}

func TestParseInstruction_nest3(t *testing.T) {
        want := `(target = "ldap:///ipatokenuniqueid=*,cn=otp,dc=example,dc=com")(targetfilter = "(objectClass=ipaToken)")(version 3.0; acl "token-add-delete"; allow(add) userattr = "ipatokenOwner#SELFDN" AND( userdn = "ldap:///uid=jesse,ou=People,dc=example,dc=com" OR userdn = "ldap:///uid=courtney,ou=People,dc=example,dc=com" )AND NOT( userattr = "terminated:TRUE" OR userattr = "pendingTermination:TRUE" );)`
        got, _, err := parseRule(want[154:])
        if err != nil {
                t.Errorf("%s failed: %v", t.Name(), err)
        }

        if want != got.String() {
                t.Errorf("%s failed:\nwant '%s' (%d)\ngot  '%s' (%d)",
                        t.Name(),
                        want,
                        len(want),
                        got,
                        len(got.String()))
        }

}

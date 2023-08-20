package aci

import "testing"

func TestParseTargetRule(t *testing.T) {
        want := `( target = "ldap:///ou=Accounting,dc=example,dc=com" )`

        c, err := ParseTargetRule(want)
        if err != nil {
                return
        }

        if want != c.String() {
                t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, c)
        }
}

func TestParseTargetRules(t *testing.T) {
        want := `( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetscope = "onelevel" )( targetattr = "cn || givenName" )`

        c, err := ParseTargetRules(want)
        if err != nil {
                return
        }

        if want != c.String() {
                t.Errorf("%s failed:\nwant '%s',\ngot  '%s'", t.Name(), want, c)
        }
}

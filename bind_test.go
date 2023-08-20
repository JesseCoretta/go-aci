package aci

import "testing"

func TestParseBindRule(t *testing.T) {
        want := `userdn = "ldap:///cn=Jesse Coretta,ou=People,dc=example,dc=com"`

        c, err := ParseBindRule(want)
        if err != nil {
                return
        }

        if want != c.String() {
                t.Errorf("%s failed: want '%s', got '%s'", t.Name(), want, c)
        }
}

func TestParseBindRules(t *testing.T) {
        want := `( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) )`

        r, err := ParseBindRules(want)
        if err != nil {
                return
        }

        if want != r.String() {
                t.Errorf("%s failed:\nwant '%s',\ngot  '%s'", t.Name(), want, r)
        }
}

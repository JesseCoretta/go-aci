package aci

import (
	"testing"
)

var bogusKeywords []string = []string{
	`bagels`,
	`63`,
	``,
	`userdnssf`,
}

func TestKeyword_bogusMatches(t *testing.T) {
	for _, bogus := range bogusKeywords {
		if bt := matchBT(bogus); bt != BindType(0x0) {
			t.Errorf("%s failed: '%s' matched bogus %T",
				t.Name(), bogus, bt)
		}

		if tk := matchTKW(bogus); tk != TargetKeyword(0x0) {
			t.Errorf("%s failed: '%s' matched bogus %T",
				t.Name(), bogus, tk)
		}

		if bk := matchBKW(bogus); bk != BindKeyword(0x0) {
			t.Errorf("%s failed: '%s' matched bogus %T",
				t.Name(), bogus, bk)
		}

		if key, ok := idKW(bogus); ok || key != nil {
			t.Errorf("%s failed: '%s' matched %s interface context",
				t.Name(), bogus, key)
		}
	}
}

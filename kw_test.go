package aci

import (
	"fmt"
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
			return
		}

		if tk := matchTKW(bogus); tk != TargetKeyword(0x0) {
			t.Errorf("%s failed: '%s' matched bogus %T",
				t.Name(), bogus, tk)
			return
		}

		if bk := matchBKW(bogus); bk != BindKeyword(0x0) {
			t.Errorf("%s failed: '%s' matched bogus %T",
				t.Name(), bogus, bk)
			return
		}

		if key, ok := idKW(bogus); ok || key != nil {
			t.Errorf("%s failed: '%s' matched %s interface context",
				t.Name(), bogus, key)
			return
		}
	}
}

// Let's print out each BindType constant
// defined in this package.
func ExampleBindType() {
	for idx, bt := range []BindType{
		USERDN,
		GROUPDN,
		ROLEDN,
		SELFDN,
		LDAPURL,
	} {
		fmt.Printf("%T %d/%d: %s\n",
			bt, idx+1, 5, bt)
	}
	// Output:
	// aci.BindType 1/5: USERDN
	// aci.BindType 2/5: GROUPDN
	// aci.BindType 3/5: ROLEDN
	// aci.BindType 4/5: SELFDN
	// aci.BindType 5/5: LDAPURL
}

/*
This example demonstrates the interrogation of BindKeyword const
definitions. This type qualifies for the Keyword interface type.

There are a total of eleven (11) such BindKeyword definitions.
*/
func ExampleBindKeyword() {
	for idx, bk := range []BindKeyword{
		BindUDN,
		BindRDN,
		BindGDN,
		BindUAT,
		BindGAT,
		BindIP,
		BindDNS,
		BindDoW,
		BindToD,
		BindAM,
		BindSSF,
	} {
		fmt.Printf("[%s] %02d/%d: %s\n",
			bk.Kind(), idx+1, 11, bk)
	}
	// Output:
	// [bind] 01/11: userdn
	// [bind] 02/11: roledn
	// [bind] 03/11: groupdn
	// [bind] 04/11: userattr
	// [bind] 05/11: groupattr
	// [bind] 06/11: ip
	// [bind] 07/11: dns
	// [bind] 08/11: dayofweek
	// [bind] 09/11: timeofday
	// [bind] 10/11: authmethod
	// [bind] 11/11: ssf
}

/*
This example demonstrates the interrogation of TargetKeyword const
definitions. This type qualifies for the Keyword interface type.

There are a total of nine (9) such TargetKeyword definitions.
*/
func ExampleTargetKeyword() {
	for idx, tk := range []TargetKeyword{
		Target,
		TargetTo,
		TargetAttr,
		TargetCtrl,
		TargetFrom,
		TargetScope,
		TargetFilter,
		TargetAttrFilters,
		TargetExtOp,
	} {
		fmt.Printf("[%s] %d/%d: %s\n",
			tk.Kind(), idx+1, 9, tk)
	}
	// Output:
	// [target] 1/9: target
	// [target] 2/9: target_to
	// [target] 3/9: targetattr
	// [target] 4/9: targetcontrol
	// [target] 5/9: target_from
	// [target] 6/9: targetscope
	// [target] 7/9: targetfilter
	// [target] 8/9: targattrfilters
	// [target] 9/9: extop
}

/*
This example demonstrates the interrogation of qualifiers of
the Keyword interface type (BindKeyword and TargetKeyword
const definitions).

There are a total of twenty (20) qualifying instances (spanning
two (2) distinct types) of this interface.
*/
func ExampleKeyword() {
	for idx, k := range []Keyword{
		BindUDN,
		BindRDN,
		BindGDN,
		BindUAT,
		BindGAT,
		BindIP,
		BindDNS,
		BindDoW,
		BindToD,
		BindAM,
		BindSSF,
		Target,
		TargetTo,
		TargetAttr,
		TargetCtrl,
		TargetFrom,
		TargetScope,
		TargetFilter,
		TargetAttrFilters,
		TargetExtOp,
	} {
		fmt.Printf("[%s] %02d/%d: %s\n",
			k.Kind(), idx+1, 20, k)
	}
	// Output:
	// [bind] 01/20: userdn
	// [bind] 02/20: roledn
	// [bind] 03/20: groupdn
	// [bind] 04/20: userattr
	// [bind] 05/20: groupattr
	// [bind] 06/20: ip
	// [bind] 07/20: dns
	// [bind] 08/20: dayofweek
	// [bind] 09/20: timeofday
	// [bind] 10/20: authmethod
	// [bind] 11/20: ssf
	// [target] 12/20: target
	// [target] 13/20: target_to
	// [target] 14/20: targetattr
	// [target] 15/20: targetcontrol
	// [target] 16/20: target_from
	// [target] 17/20: targetscope
	// [target] 18/20: targetfilter
	// [target] 19/20: targattrfilters
	// [target] 20/20: extop
}

func ExampleBindKeyword_String() {
	fmt.Printf("%s", BindUDN)
	// Output: userdn
}

func ExampleBindKeyword_Kind() {
	fmt.Printf("%s", BindUDN.Kind())
	// Output: bind
}

func ExampleTargetKeyword_String() {
	fmt.Printf("%s", TargetScope)
	// Output: targetscope
}

func ExampleTargetKeyword_Kind() {
	fmt.Printf("%s", TargetAttrFilters.Kind())
	// Output: target
}

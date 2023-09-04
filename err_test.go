package aci

import (
	"testing"
)

func TestErrorf(t *testing.T) {
	var results []bool = []bool{true, true, false}
	for idx, msg := range []any{
		`this is an error`,
		errorf(`this is also an error`),
		nil, // but this is not.
	} {
		err := errorf(msg)
		got := err != nil
		want := results[idx]
		if want != got {
			t.Errorf("%s failed: unexpected errorf result; want '%t', got '%t'",
				t.Name(), want, got)
			return
		}
	}
}

func TestErrorFNs_codecov(t *testing.T) {
	_ = badPTBRuleKeywordErr(BindRule{}, `blarg`, `userdn`, `farts`)
	_ = badObjectIdentifierKeywordErr(TargetExtOp)
	_ = unexpectedKindErr(PermissionBindRule{}, `pbr`, `...`)
	_ = illegalSliceTypeErr(AttributeFilter{}, `(objectClass=*)`, `35^`)
	_ = illegalSyntaxPerTypeErr(BindDistinguishedName{},
		BindUDN,
		errorf("This is an error"))
	_ = afoMissingPrefixErr()
	_ = aoBadPrefixErr()
	_ = instructionNoLabelErr()
	_ = dowBadTimeErr()
	_ = badCopErr(badCop)
	_ = noPermissionDispErr()
	_ = fqdnInvalidLabelErr(domainLabel(`__`))
	_ = parseBindRulesHierErr(BindDistinguishedNames{}, BindRules{})
	_ = unexpectedStringResult(`faily mcfailface`, `mad skillz`, `pwned`)
	_ = generalErr(`unknown`, errorf("This is another error"))
	_ = pushError(
		TargetDistinguishedNames{},
		TargetDistinguishedName{},
		Target, "You're in trouble",
		errorf("Yet another error"))
}

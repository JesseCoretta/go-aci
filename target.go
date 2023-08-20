package aci

import (
	parser "github.com/JesseCoretta/go-antlraci"
	//"github.com/JesseCoretta/go-stackage"
)

func NewTargetRules() Rule {
	return T()
	/*
	func NewTargetRules() TargetRules {
	return TargetRules(stackageList(9)).
		setID(`target`).
		setCategory(`target`).
		NoPadding(!RulePadding).
		setPushPolicy()
	*/
}

func ParseTargetRule(raw string) (Condition, error) {
	return parseTargetRule(raw)
}

func parseTargetRule(raw string) (Condition, error) {
        t, err := parser.ParseTargetRule(raw)
	c := Condition(t)

	//c, err = assertTargetRule(t.Values.Values, t.Values.Style, t.Keyword, t.Operator)

	return c, err
}

func ParseTargetRules(raw string) (t Rule, err error) {
	return parseTargetRules(raw)
}

func parseTargetRules(raw string) (Rule, error) {
	_t, err := parser.ParseTargetRules(raw)
	t := Rule(_t)

	return t, err
}

func assertTargetRule(vals []string, qt int, kw, op string) (c Condition, err error) {
        // Begin with an assertion switch upon the target keyword
        // (which we already vetted as sane) ...
        switch key := matchTKW(kw); key {

        case TargetScope, TargetFilter:
                if len(vals) != 1 {
                        err = errorf("Unexpected number of %s values; want %d, got %d",
                                key, 1, len(vals))
                        return
                }

                if key == TargetScope {
                        c, err = assertTargetScope(unquote(vals[0]), op)
                } else {
                        f := unquote(vals[0])
                        c = TFilter().Push(f).Eq()
                }

        case TargetAttr:
                c, err = assertTargetAttributes(vals, qt, op)

        case TargetCtrl, TargetExtOp:
                c, err = assertTargetOID(vals, qt, op, key)

        case Target, TargetTo, TargetFrom:
                c, err = assertTargetDN(vals, qt, op, key)

        case TargetAttrFilters:
                // TODO
                //if len(vals) != 1 {
                //      err = errorf("Target Rule keyword '%s' supports single values only, but %d values were found: %v",
                //              kw,len(vals),vals)
                //      return
                //}

        default:
                err = errorf("Unhandled target rule type '%s'", key)
        }

        return
}

func assertTargetScope(value string, op string) (c Condition, err error) {
        if len(value) == 0 {
                err = errorf("Zero-length LDAP Search Scope detected; aborting")
                return
        }

        scn := unquote(value)
        sc := strToScope(scn)

        // base is a fallback for a bogus scope, so
        // if the user did not originally request
        // base, we know they requested something
        // totally unsupported.
        if sc == noScope {
                err = errorf("Bogus %s value: '%s'", TargetScope, scn)
                return
        }

        c, err = conditionByOperator(op, sc)

        return
}

func assertTargetOID(vals []string, qt int, op string, key TargetKeyword) (c Condition, err error) {
        var toid Rule
        if key == TargetExtOp {
                toid = ExtOps()
        } else {
                toid = Ctrls()
        }

        if qt == 1 {
                toid.Encap()
        } else {
                toid.Encap(`"`)
        }

        for x := 0 ; x < len(vals); x++ {
                o, _ := newObjectID(key, vals[x])
		toid.Push(ObjectIdentifier{o})
        }

        c, err = conditionByOperator(op, toid)
        if qt == 1 {
                c.Encap(`"`)
        } else {
                c.Encap()
	}

        return
}

func assertTargetDN(vals []string, qt int, op string, key TargetKeyword) (c Condition, err error) {
        var tdnr Rule = TDNs().setCategory(key.String())

	if qt == 1 {
		tdnr.Encap()
	} else {
		tdnr.Encap(`"`)
	}

	for x := 0 ; x < len(vals); x++ {
		D := vals[x]
	        if !hasPfx(D, LocalScheme) {
	                err = errorf("Illegal %s distinguishedName slice: [index:%d;value:%s] missing LDAP local scheme (%s)",
	                        key, x, D, LocalScheme)
	                return
	        }
	        tdnr.Push(DistinguishedName{newDistinguishedName(D[len(LocalScheme):], key)})
	}

        c, err = conditionByOperator(op, tdnr)
        if qt == 1 {
                c.Encap(`"`)
        } else {
                c.Encap()
        }

        return
}

func assertTargetAttributes(vals []string, qt int, op string) (c Condition, err error) {
        var tat Rule = TAttrs().Encap()

        if qt == 1 {
                tat.Encap()
        } else {
                tat.Encap(`"`)
        }

        for x := 0 ; x < len(vals); x++ {
                tat.Push(ATName(vals[x]))
        }

        c, err = conditionByOperator(op, tat)
	if qt == 1 {
		c.Encap(`"`)
	} else {
		c.Encap()
	}

        return
}



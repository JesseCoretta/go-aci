package aci

/*
parseTR reads and processes tokens into a sequence of zero (0) or more
Target Rule instances, which are returned alongside an error.

The chop return value informs the calling method which index integer
from which to continue, thus avoiding tokens already processed.

Target Rules are not nested; if a non-zero targetRules stack is returned,
it shall contain a like-number of Condition instances, each bearing a
target keyword (e.g.: `targetattr`), an equality (=) OR negated (!=)
comparison operator and a doubled-quoted expression value. Note that Target
Rule conditions are **ALWAYS** parenthetical.
*/
func parseTR(tokens []string) (chop int, targetRules Rule, err error) {

        // Initialize our target rule stack; max capacity is nine (9)
        targetRules = T()

        var (
                kw, // target rule keyword
                cop, // target rule comparison operator
                last, // previous token
                next string // upcoming token

                // vals stores a sequence of value (and value RELATED)
                // tokens when detected. This var is purged of its
                // contents whenever the end of the value sequence is
                // reached.
                vals []string = make([]string, 0)

                done,
                cready bool // condition ready for assembly
        )

        // iterate tokens, looking for target rule elements.
        for index, token := range tokens {
                // get previous and upcoming tokens, if possible.
                if index > 0 {
                        last = tokens[index-1]
                }

                if index+1 < len(tokens) {
                        // look-ahead to the next iteration's token
                        next = tokens[index+1]
                }

                _, istok := matchOp(token)

                switch {

                // token is a target rule keyword
                case matchTKW(lc(token)) != TargetKeyword(0x0):
                        kw = lc(token)

                // token is a target rule comparison operator
                case istok:
                        cop = token
                        cready = len(kw) > 0

                // token is anchor, meaning there are no more
                // target rules to process.
                case token == `version 3.0; acl`:
                        if chop = index; chop > 0 {
                                done = true
                                break
                        }

                // generalized token fallback will capture quoted values as well as
                // logical symbolic operators (||, &&) between quoted values (or as
                // part of a quoted value).
                default:
                        if cready {
                                // condition is ready for assembly AND a non-zero
                                // double-quoted value is the current token. We
                                // will also accept symbolic operators if we're
                                // dealing with a multi-valued expression.
                                if err = targetExprValueInvalid(token, last, next, index); err != nil {
                                        return
                                }

                                // strip quotes, as go-stackage provides encaps
                                // without the need for literal storage of such
                                // characters.

                                // increment chop index by one (1)
                                chop++

                                // Save this value; we don't yet know if this
                                // value is merely one (1) of multiple values
                                // as opposed to a single value alone.
                                vals = append(vals, token)

                                // Look ahead to see what is coming next. If
                                // another quoted value or symbolic operator
                                // are detected, we know we're not done yet.
                                // In that case, break out of this case to
                                // continue at the next for-loop iteration.
                                if targetRuleNotDone(next) {
                                        break
                                }

                                var (
                                        // Prepare our condition for target rule creation
                                        c Condition
                                )

                                // This is the last (or only!) value component. We can
                                // now analyze the keyword and the value(s) to ascertain
                                // the appropriate instance type for condition storage
                                // (and to perform other context-specific sanity checks).

                                if c, err = assertTargetRule(vals, kw, cop); err != nil {
                                        return
                                }
                                targetRules.Push(c)

                                // Reset for next target rule condition, if any
                                kw = ``
                                cop = ``
                                cready = false
                                vals = make([]string, 0)
                        }
                }

                if done {
                        break
                }
        }

        return
}

func targetExprValueInvalid(token, last, next string, index int) (err error) {
        if !isQuoted(token) && (token != `||` && token != `&&`) {
                err = errorf("Bogus Target Rule condition expression between '%s' [%d] and '%s' [%d]; value must be a non-zero string enclosed within double quotes, or a symbolic list (||,&&) of same",
                        last, index-1, next, index+1)
        }
        return
}

func targetExprReady(kw, op string) bool {
        return len(kw) > 0 && len(op) > 0
}

func targetRuleNotDone(next string) bool {
        return next == `||` || next == `&&` || isQuoted(next)
}

func assertTargetRule(vals []string, kw, op string) (c Condition, err error) {
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
                c, err = assertTargetAttributes(vals, op)

        case TargetCtrl, TargetExtOp:
                c, err = assertTargetOID(vals, op, key)

        case Target, TargetTo, TargetFrom:
                c, err = assertTargetDN(vals, op, key)

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

func assertTargetOID(vals []string, op string, key TargetKeyword) (c Condition, err error) {
        var vencap bool
        var toid Rule
        if key == TargetExtOp {
                toid = ExtOps()
        } else {
                toid = Ctrls()
        }

        // target rule is either or both of the following:
        // A: one (1) double-quoted OID
        // B: one (1) double-quoted LIST of unquoted OIDs in symbolic OR context
        for x := 0; x < len(vals); x++ {
                var value string = vals[x]

                if contains(value, `||`) {

                        // Type-B confirmed
                        for ix, O := range split(unquote(value), `||`) {
                                if len(O) == 0 {
                                        continue
                                }

                                if x == 0 && ix == 0 {
                                        if !isQuoted(vals[x]) && isQuoted(O) {
                                                vencap = true
                                                toid.Encap()
                                        } else if !isQuoted(O) {
                                                toid.Encap(`"`)
                                        }
                                }

                                value = trimS(unquote(O))
                                o, _ := newObjectID(key, value)
                                toid.Push(ObjectIdentifier{o})
                        }
                } else {
                        if x == 0 {
                                if isQuoted(value) {
                                        vencap = true
                                        toid.Encap(`"`)
                                }
                        }

                        // Type-A confirmed
                        value = trimS(unquote(value))
                        o, _ := newObjectID(key, value)
                        toid.Push(ObjectIdentifier{o})
                }
        }

        c, err = conditionByOperator(op, toid)
        if !vencap {
                c.Encap(`"`)
                return
        }
        c.Encap()

        return
}

func assertTargetDN(vals []string, op string, key TargetKeyword) (c Condition, err error) {
        var vencap bool
        var tdnr Rule = TDNs().setCategory(key.String())

        // target rule is either or both of the following:
        // A: one (1) double-quoted DN
        // B: one (1) double-quoted LIST of unquoted DNs in symbolic OR context
        for x := 0; x < len(vals); x++ {
                var value string = vals[x]
                if contains(value, `||`) {

                        // Type-B confirmed
                        for ix, O := range split(unquote(value), `||`) {
                                if len(O) == 0 {
                                        continue
                                }

                                if x == 0 && ix == 0 {
                                        if !isQuoted(vals[x]) && isQuoted(O) {
                                                vencap = true
                                                tdnr.Encap()
                                        } else if !isQuoted(O) {
                                                tdnr.Encap(`"`)
                                        }
                                }

                                D := trimS(unquote(O))
                                if !hasPfx(D, LocalScheme) {
                                        err = errorf("Illegal %s distinguishedName slice: [index:%d;value:%s] missing LDAP local scheme (%s)",
                                                key, x, D, LocalScheme)
                                        return
                                }

                                tdnr.Push(DistinguishedName{newDistinguishedName(D[len(LocalScheme):], key)})
                        }

                } else {

                        // Type-A confirmed
                        if x == 0 {
                                if isQuoted(value) {
                                        vencap = true
                                        tdnr.Encap(`"`)
                                }
                        }

                        D := unquote(value)
                        if !hasPfx(D, LocalScheme) {
                                err = errorf("Illegal %s distinguishedName: [index:%d;value:%s] missing LDAP local scheme (%s)",
                                        key, x, D, LocalScheme)
                                return
                        }

                        tdnr.Push(DistinguishedName{newDistinguishedName(D[len(LocalScheme):], key)})
                }
        }

        c, err = conditionByOperator(op, tdnr)
        if !vencap {
                c.Encap(`"`)
                return
        }
        c.Encap()

        return
}

func assertTargetAttributes(vals []string, op string) (c Condition, err error) {
        var vencap bool
        var tat Rule = TAttrs().Encap()

        // target rule is either or both of the following:
        // A: one (1) double-quoted AT
        // B: one (1) double-quoted LIST of unquoted ATs in symbolic OR context
        for x := 0; x < len(vals); x++ {
                var value string = vals[x]

                if contains(value, `||`) {

                        // Type-B confirmed
                        for ix, O := range split(unquote(value), `||`) {
                                if len(O) == 0 {
                                        continue
                                }

                                if x == 0 && ix == 0 {
                                        if !isQuoted(vals[x]) && isQuoted(O) {
                                                vencap = true
                                                tat.Encap()
                                        } else if !isQuoted(O) {
                                                tat.Encap(`"`)
                                        }
                                }

                                tat.Push(ATName(trimS(unquote(O))))
                        }

                } else {
                        // Type-A confirmed
                        if x == 0 {
                                if isQuoted(value) {
                                        vencap = true
                                        tat.Encap(`"`)
                                }
                        }
                        tat.Push(ATName(trimS(unquote(value))))
                }
        }

        c, err = conditionByOperator(op, tat)
        if !vencap {
                c.Encap(`"`)
                return
        }
        c.Encap()

        return
}



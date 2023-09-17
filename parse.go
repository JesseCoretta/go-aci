package aci

/*
parse.go is a bridge to the go-antlraci package
*/

import (
	parser "github.com/JesseCoretta/go-antlraci"
)

/*
ParseBindRule returns an instance of BindRule alongside an error instance.

This function calls the imported antlraci.ParseBindRule function, delegating
parsing responsibilities there.
*/
func ParseBindRule(raw string) (BindRule, error) {
	return parseBindRule(raw)
}

func (r *BindRule) Parse(raw string) error {
	_r, err := parseBindRule(raw)
	if err != nil {
		return err
	}
	*r = _r

	return nil
}

func parseBindRule(raw string) (BindRule, error) {
	_r, err := parser.ParseBindRule(raw)
	return BindRule(_r), err
}

/*
ParseBindRules returns an instance of BindContext alongside an error
instance. BindContext may represent either a BindRule or BindRules
instance, depending on that which was parsed.
*/
func ParseBindRules(raw string) (BindContext, error) {
	return parseBindRules(raw)
}

/*
Parse returns an error based upon an attempt to parse the raw input
value into the receiver instance. If successful, any contents within
the receiver instance would be obliterated, replaced irrevocably by
the freshly parsed values.

Both this method, and the package-level ParseBindRule function, call
antlraci's ParseBindRule function in similar fashion. The only real
difference here is the process of writing to a receiver, versus writing
to an uninitialized variable declaration.
*/
func (r *BindRules) Parse(raw string) error {
	_r, err := parseBindRules(raw)
	if err != nil {
		return err
	}
	switch tv := _r.(type) {
	case BindRule:
		r.reset()
		r.Push(tv)
	case BindRules:
		*r = tv
	}

	return nil
}

/*
parseBindRules communicates with the backend parser (antlraci)
package for the purpose of parsing an instance of BindRules,
which is returned alongside an error.
*/
func parseBindRules(raw string) (BindContext, error) {
	// In case the input has bizarre
	// contiguous whsp, etc., remove
	// it safely.
	raw = condenseWHSP(raw)

	// send the raw textual bind rules
	// statement(s) to our sister package
	// go-antlraci, call ParseBindRules.
	_b, err := parser.ParseBindRules(raw)
	if err != nil {
		return badBindRules, err
	}

	// Process the hierarchy, converting
	// Stack to BindRules and Condition
	// to BindRule. In addition, we'll
	// replace the parser.ExpressionValue
	// type with more appropriate types
	// defined in this package.
	n, ok := convertBindRulesHierarchy(_b)

	// for codecov
	if err = parseBindRulesHierErr(_b, n); ok {
		err = nil
	}

	return n, err
}

/*
assertParserRuleExpr an assertion on parser.RuleExpression
instances and is largely for convenience.
*/
func assertParserRuleExpr(x any) (expr parser.RuleExpression, is bool) {
	expr, is = x.(parser.RuleExpression)
	return
}

/*
makeParserRuleExpr is for testing purposes only.
*/
func makeParserRuleExpr(style int, x ...string) parser.RuleExpression {
	return parser.RuleExpression{
		Style:  style,
		Values: append([]string{}, x...),
	}
}

func assertBindTimeDay(expr parser.RuleExpression, key BindKeyword) (ex any, err error) {
	switch key {
	case BindDoW:
		// value is a dayOfWeek expressive
		// statement.
		ex, err = assertBindDayOfWeek(expr)

	case BindToD:
		// value is a timeOfDay expressive
		// statement.
		ex, err = assertBindTimeOfDay(expr)
	}

	return
}

func assertBindSec(expr parser.RuleExpression, key BindKeyword) (ex any, err error) {
	switch key {
	case BindSSF:
		// value is a security strength factor
		// expressive statement.
		ex, err = assertBindSecurityStrengthFactor(expr)

	case BindAM:
		// value is an authentication method
		// expressive statement.
		ex, err = assertBindAuthenticationMethod(expr)
	}

	return
}

func assertBindUGAttr(expr parser.RuleExpression, key BindKeyword) (ex any, err error) {
	if err = unexpectedBindConditionValueErr(key, 1, expr.Len()); err != nil {
		return
	}

	value := unquote(condenseWHSP(expr.Values[0]))

	if hasPfx(value, LocalScheme) {
		// value is an LDAP URI
		ex, err = parseLDAPURI(value, key)

	} else if hasPfx(value, `parent[`) {
		// value is an inheritance attributeBindTypeOrValue
		ex, err = parseInheritance(value)

	} else {
		// value is a standard attributeBindTypeOrValue
		ex, err = parseATBTV(value, key)
	}

	return
}

func assertBindTimeOfDay(expr parser.RuleExpression) (ex TimeOfDay, err error) {
	if err = unexpectedBindConditionValueErr(BindToD, 1, expr.Len()); err != nil {
		return
	}

	// extract clocktime from raw value, remove
	// quotes and any L/T WHSP
	unq := unquote(condenseWHSP(expr.Values[0]))
	ex = ToD(unq)
	err = badClockTimeErr(unq, ex.String())
	return
}

func assertBindDayOfWeek(expr parser.RuleExpression) (ex DayOfWeek, err error) {
	if err = unexpectedBindConditionValueErr(BindDoW, 1, expr.Len()); err != nil {
		return
	}

	// extract auth method from raw value, remove
	// quotes and any L/T WHSP and analyze
	unq := unquote(condenseWHSP(expr.Values[0]))
	ex, err = parseDoW(unq)
	return
}

func assertBindAuthenticationMethod(expr parser.RuleExpression) (ex AuthenticationMethod, err error) {
	if err = unexpectedBindConditionValueErr(BindAM, 1, expr.Len()); err != nil {
		return
	}

	// extract auth method from raw value, remove
	// quotes and any L/T WHSP and analyze
	unq := unquote(condenseWHSP(expr.Values[0]))
	ex = matchAuthenticationMethod(unq)
	err = badAMErr(unq, ex.String())
	return
}

func assertBindSecurityStrengthFactor(expr parser.RuleExpression) (ex SecurityStrengthFactor, err error) {
	if err = unexpectedBindConditionValueErr(BindSSF, 1, expr.Len()); err != nil {
		return
	}

	// extract factor from raw value, remove
	// quotes and any L/T WHSP
	unq := unquote(condenseWHSP(expr.Values[0]))
	ex = SSF(unq)
	err = badSecurityStrengthFactorErr(unq, ex.String())
	return
}

func assertBindNet(expr parser.RuleExpression, key BindKeyword) (ex any, err error) {
	if err = unexpectedBindConditionValueErr(key, 1, expr.Len()); err != nil {
		return
	}

	unq := unquote(condenseWHSP(expr.Values[0]))

	if key == BindIP {
		// extract IP Address(es) from raw value,
		// remove quotes and any L/T WHSP and then
		// split for iteration.
		raw := split(unq, `,`)
		var addr IPAddr
		for ipa := 0; ipa < len(raw); ipa++ {
			addr.Set(raw[ipa])
		}

		ex = addr
		err = badIPErr(len(raw), addr.Len())
		return
	}

	// extract FQDN from raw value, remove
	// quotes and any L/T WHSP.
	fq := DNS(unq)
	err = badDNSErr(unq, fq.String())
	ex = fq

	return
}

/*
assertBindUGRDN is handler for all possible DN and URI values used within Bind Rule
expressive statements. In particular, this handles `userdn`, `groupdn` and `roledn`
keyword contexts.

An any-enveloped DistinguishedNames instance is returned in the event that the raw value(s)
represent one (1) or more legal LDAP Distinguished Name value.

In the event that a legal LDAP URI is found, it is returned as an instance of (any-enveloped)
LDAPURI.

Quotation schemes are supported seamlessly and either scheme shall be honored per the ANTLR4
parsed content.
*/
func assertBindUGRDN(expr parser.RuleExpression, key BindKeyword) (ex any, err error) {
	// Don't waste time if expression values
	// are nonexistent.
	if expr.Len() == 0 {
		err = noTBRuleExpressionValues(expr, bindRuleID, key)
		return
	}

	// if the value is an LDAP URI (which merely contains
	// a DN, and is not one unto itself), handle the parse
	// here instead of treating it as a DN.
	var value string = unquote(condenseWHSP(expr.Values[0]))
	if len(value) < 3 {
		err = illegalSyntaxPerTypeErr(value, key)
		return
	}

	// create an appropriate container based on the
	// Bind Rule keyword.
	var bdn BindDistinguishedNames
	switch key {
	case BindRDN:
		bdn = RDNs()
	case BindGDN:
		bdn = GDNs()
	default:
		bdn = UDNs()
	}

	// Honor the established quotation scheme that
	// was observed during ANTLR4 processing.
	bdn.setQuoteStyle(expr.Style)

	// Assign the raw (DN) values to the
	// return value. If nothing was found,
	// bail out now.
	if err = bdn.setExpressionValues(key, expr.Values...); err != nil {
		return
	}

	// Envelope our DN stack within an
	// 'any' instance, which is returned.
	ex = bdn

	return
}

/*
assertExpressionValue will update the underlying go-antlraci temporary type with a
proper value-appropriate type defined within the go-aci package. An error is returned
upon processing completion.
*/
func (r BindRule) assertExpressionValue() (err error) {

	// grab the raw value from the receiver. If it is
	// NOT an instance of parser.RuleExpression, then
	// bail out.
	expr, ok := assertParserRuleExpr(r.Expression())
	if !ok {
		err = parseBindRuleInvalidExprTypeErr(r, expr, expr)
		return
	}

	// our proper type-converted expression
	// value(s) shall reside as an any, as
	// stackage.Condition allows this and
	// will make things simpler.
	var ex any

	// perform a bind keyword switch upon
	// a resolution attempt of the value.
	switch key := matchBKW(r.Keyword().String()); key {

	case BindUDN, BindRDN, BindGDN:
		// value is a userdn, groupdn or roledn
		// expressive statement. Possible multi
		// valued expression.
		ex, err = assertBindUGRDN(expr, key)

	case BindIP, BindDNS:
		// value is an IP or FQDN.
		ex, err = assertBindNet(expr, key)

	case BindUAT, BindGAT:
		// value is a userattr or groupattr
		// expressive statement.
		ex, err = assertBindUGAttr(expr, key)

	case BindDoW, BindToD:
		// value is a dayofweek or timeofday
		// expressive statement.
		ex, err = assertBindTimeDay(expr, key)

	case BindAM, BindSSF:
		// value is an authentication method
		// or a security factor expressive
		// statement.
		ex, err = assertBindSec(expr, key)

	default:
		err = badPTBRuleKeywordErr(r, bindRuleID, `BindKeyword`, key)
	}

	if err != nil {
		return
	}

	// If we got something, set it and go.
	r.SetExpression(ex)
	r.SetQuoteStyle(expr.Style)

	return
}

/*
ParseTargetRule processes the raw input string value,
which should represent a single Target Rule expressive
statement, into an instance of TargetRule. This, along
with an error instance, are returned upom completion
of processing.
*/
func ParseTargetRule(raw string) (TargetRule, error) {
	return parseTargetRule(raw)
}

/*
parseTargetRule is a private function which converts the
stock stackage.Condition instance assembled by go-antlraci
and casts as a go-aci TargetRule instance, which will be
returned alongside an error upon completion of processing.
*/
func parseTargetRule(raw string) (TargetRule, error) {
	_t, err := parser.ParseTargetRule(raw)
	t := TargetRule(_t)
	t.assertExpressionValue()
	return t, err
}

/*
Parse returns an error based upon an attempt to parse the raw input
value into the receiver instance. If successful, any contents within
the receiver instance would be obliterated, replaced irrevocably by
the freshly parsed values.

Both this method, and the package-level ParseTargetRule function,
call antlraci's ParseTargetRule function in similar fashion. The only
real difference here is the process of writing to a receiver, versus
writing to an uninitialized variable declaration.
*/
func (r *TargetRule) Parse(raw string) error {
	_r, err := parseTargetRule(raw)
	if err != nil {
		return err
	}
	*r = _r

	return nil
}

/*
Parse returns an error based upon an attempt to parse the raw input
value into the receiver instance. If successful, any contents within
the receiver instance would be obliterated, replaced irrevocably by
the freshly parsed values.

Both this method, and the package-level ParseTargetRules function,
call antlraci's ParseTargetRules function in similar fashion. The only
real difference here is the process of writing to a receiver, versus
writing to an uninitialized variable declaration.
*/
func (r *TargetRules) Parse(raw string) error {
	_r, err := parseTargetRules(raw)
	if err != nil {
		return err
	}
	*r = _r

	return nil
}

/*
ParseTargetRules processes the raw input string value,
which should represent one (1) or more valid Target Rule
expressive statements, into an instance of TargetRules.
This, alongside an error instance, are returned at the
completion of processing.
*/
func ParseTargetRules(raw string) (TargetRules, error) {
	return parseTargetRules(raw)
}

/*
parseTargetRules is a private function which converts the
stock stackage.Stack instance assembled by go-antlraci and
coaxes the raw string values into proper value-appropriate
type instances made available by go-aci.
*/
func parseTargetRules(raw string) (TargetRules, error) {
	// In case the input has bizarre
	// contiguous whsp, etc., remove
	// it safely.
	raw = condenseWHSP(raw)

	// Call our go-antlraci (parser) package's
	// ParseTargetRules function, and get the
	// results (or bail if error).
	_t, err := parser.ParseTargetRules(raw)
	if err != nil {
		return badTargetRules, err
	}

	return processTargetRules(_t)
}

func processTargetRules(stack any) (TargetRules, error) {
	var err error

	_z, ok := castAsStack(stack)
	if !ok {
		err = errorf("Invalid input type %T; expecting stackage.Stack", stack)
		return badTargetRules, err
	}

	// create our (eventual) return object.
	t := TRs()

	// transfer raw contents into new TargetRules
	// instance.
	for i := 0; i < _z.Len(); i++ {
		slice, _ := _z.Index(i)
		t.Push(TargetRule(derefC(slice)))
	}

	// iterate our (new) target rule slice members,
	// identifying each by integer index i. Try to
	// marshal the parser.RuleExpression contents
	// into the appropriate go-aci type.
	for i := 0; i < t.Len(); i++ {
		trv := t.Index(i)

		// Extract individual expression value
		// from TargetRule (ntv), and recreate it
		// using the proper type, replacing the
		// original. For example, a `target_to`
		// (DN) Target Rule with a RuleExpression
		// value of:
		//
		//   []string{<dn1>,<dn2>,<dn3>}
		//
		// ... shall be replaced with:
		//
		//   <stack alias type>-idx#------val-
		//   DistinguishedNames[<N1>] -> <dn1>
		//                     [<N2>] -> <dn2>
		//                     [<N3>] -> <dn3>
		if err = trv.assertExpressionValue(); err != nil {
			return badTargetRules, err
		}
	}

	return t, err
}

/*
assertExpressionValue will update the underlying go-antlraci temporary expression type
with a proper value-appropriate type defined within the go-aci package.

An error is returned upon processing completion.
*/
func (r *TargetRule) assertExpressionValue() (err error) {
	// grab the raw value from the receiver. If it is
	// NOT an instance of parser.RuleExpression, then
	// bail out.
	expr, ok := assertParserRuleExpr(r.Expression())
	if !ok {
		err = parseBindRuleInvalidExprTypeErr(r, expr, expr)
		return
	}

	// our proper type-converted expression
	// value(s) shall reside as an any, as
	// stackage.Condition allows this and
	// will make things simpler.
	var ex any

	// perform a target keyword switch upon
	// a resolution attempt of the value.
	switch key := matchTKW(r.Keyword().String()); key {

	case TargetScope, TargetFilter, TargetAttrFilters:
		// value is a targetscope, targetfilter or a
		// targattrfilters expressive statement. We
		// handle them here because they're strictly
		// single-valued.

		if key == TargetScope {
			// value is a target scope
			ex, err = assertTargetScope(expr)

		} else if key == TargetAttrFilters {
			// value is a targattrfilters
			ex, err = assertTargetAttrFilters(expr)

		} else {
			// value (seems to be) an LDAP Search Filter
			// TODO - assertion func
			ex, err = assertTargetFilter(expr)
		}

	case TargetAttr:
		// value is a targetattr expressive statement,
		// possibly multi-valued.
		ex, err = assertTargetAttributes(expr)

	case TargetCtrl, TargetExtOp:
		// value is a targetcontrol or extop expressive
		// statement, possibly multi-valued.
		ex, err = assertTargetOID(expr, key)

	case Target, TargetTo, TargetFrom:
		// value is a target, target_to or target_from
		// expressive statement, possibly multi-valued
		ex, err = assertTargetTFDN(expr, key)

	default:
		// value is ... bogus
		err = badPTBRuleKeywordErr(expr, targetRuleID, `TargetKeyword`, key)
	}

	if err != nil {
		return
	}

	r.SetExpression(ex)
	r.SetQuoteStyle(expr.Style)

	return
}

func assertTargetFilter(expr parser.RuleExpression) (ex SearchFilter, err error) {
	if expr.Len() != 1 {
		err = unexpectedValueCountErr(TargetAttrFilters.String(), 1, expr.Len())
		return
	}

	value := unquote(condenseWHSP(expr.Values[0]))
	if len(value) < 3 {
		err = nilInstanceErr(ex)
		return
	}

	ex = Filter(value)
	return
}

/*
assertTargetOID is handler for all possible OID values used within Target Rule expressive
statements. In particular, this handles `targetcontrol` and `extop`.

An ObjectIdentifiers instance is returned in the event that the raw value(s) represent one
(1) or more legal ASN.1 Object Identifiers in "dot notation".

Quotation schemes are supported seamlessly and either scheme shall be honored per the ANTLR4
parsed content.
*/
func assertTargetOID(expr parser.RuleExpression, key TargetKeyword) (ex ObjectIdentifiers, err error) {
	// Don't waste time if expression values
	// are nonexistent.
	if expr.Len() == 0 {
		err = noValueErr(ex, key.String())
		return
	}

	// create an appropriate container based on the
	// Target Rule keyword.
	switch key {
	case TargetExtOp:
		ex = ExtOps()
	default:
		ex = Ctrls()
	}

	// Honor the established quotation scheme that
	// was observed during ANTLR4 processing.
	ex.setQuoteStyle(expr.Style)

	// Assign the raw (DN) values to the
	// return value. If nothing was found,
	// bail out now.
	if ex.setExpressionValues(key, expr.Values...); ex.Len() == 0 {
		err = noValueErr(ex, `targetcontrol/extop`)
		return
	}

	return
}

/*
assertTargetTFDN is handler for all possible DN values used within Target Rule expressive
statements. In particular, this handles `target`, `target_to` and `target_from` keyword
contexts.

A DistinguishedNames instance is returned in the event that the raw value(s) represent one
(1) or more legal LDAP Distinguished Name value.

Quotation schemes are supported seamlessly and either scheme shall be honored per the ANTLR4
parsed content.
*/
func assertTargetTFDN(expr parser.RuleExpression, key TargetKeyword) (ex any, err error) {
	// Don't waste time if expression values
	// are nonexistent.
	if expr.Len() == 0 {
		err = noValueErr(ex, key.String())
		return
	}

	// create an appropriate container based on the
	// Target Rule keyword.
	var tdn TargetDistinguishedNames
	switch key {
	case TargetTo:
		tdn = TTDNs()
	case TargetFrom:
		tdn = TFDNs()
	default:
		tdn = TDNs()
	}

	// Honor the established quotation scheme that
	// was observed during ANTLR4 processing.
	tdn.setQuoteStyle(expr.Style)

	// Assign the raw (DN) values to the
	// return value. If nothing was found,
	// bail out now.
	if err = tdn.setExpressionValues(key, expr.Values...); err != nil {
		return
	}

	ex = tdn

	return
}

/*
assertTargetAttributes is a private functions called during the processing of a TargetRule
expressive statement bearing the `targetattr` keyword context. An instance of AttributeTypes
is returned.
*/
func assertTargetAttributes(expr parser.RuleExpression) (ex AttributeTypes, err error) {
	// Don't waste time if expression values
	// are nonexistent.
	if expr.Len() == 0 {
		err = noValueErr(ex, TargetAttr.String())
		return
	}

	ex = TAs()
	ex.setQuoteStyle(expr.Style)

	for i := 0; i < expr.Len(); i++ {
		value := unquote(condenseWHSP(expr.Values[0]))
		if len(value) == 0 {
			err = nilInstanceErr(AttributeType{})
			return
		}
		attr := AT(value)
		if attr.IsZero() {
			err = nilInstanceErr(attr)
			return
		}
		ex.Push(attr)
	}

	return
}

/*
assertTargetAttrFilters is a private function called during the processing of a TargetRule
expressive statement bearing the `targattrfilters` keyword context. An instance of the
AttributeFilterOperations type is returned, alongside an error instance, when processing is
complete.
*/
func assertTargetAttrFilters(expr parser.RuleExpression) (ex AttributeFilterOperations, err error) {
	if expr.Len() != 1 {
		err = unexpectedValueCountErr(TargetAttrFilters.String(), 1, expr.Len())
		return
	}

	value := unquote(condenseWHSP(expr.Values[0]))
	if len(value) < 3 {
		err = nilInstanceErr(ex)
		return
	}

	if idx := idxr(value, ','); idx != -1 {
		// First, try to split on a comma rune (ASCII #44).
		// This is the default, and is the most common char
		// for use in delimiting values of this form.
		ex, err = parseAttributeFilterOperations(expr.Values[0], 0)

	} else if idx = idxr(value, ';'); idx != -1 {
		// If no comma was found, try semicolon (ASCII #59).
		ex, err = parseAttributeFilterOperations(value, 1)

	} else if hasAttributeFilterOperationPrefix(value) {
		// Still nothing? Try AttributeFilterOperation (whether
		// multivalued or not).
		var afo AttributeFilterOperation
		if afo, err = parseAttributeFilterOperation(value); err != nil {
			return
		}
		ex = AFOs(afo)

	} else {
		// The only other thing it could be is a bare AttributeFilter.
		var af AttributeFilter
		if af, err = parseAttributeFilter(value); err != nil {
			return
		}

		ex = AFOs(AddOp.AFO(af)) // we have to choose one, 'add' seems safer than 'delete'
	}

	return
}

/*
assertTargetScope processes the raw expression value (expr) provided by go-antlraci
into a proper instance of SearchScope (ex), which is returned alongside an instance of
error (err).
*/
func assertTargetScope(expr parser.RuleExpression) (ex SearchScope, err error) {
	if expr.Len() != 1 {
		err = unexpectedValueCountErr(TargetScope.String(), 1, expr.Len())
		return
	}
	value := unquote(condenseWHSP(expr.Values[0]))
	if len(value) == 0 {
		err = nilInstanceErr(ex)
		return
	}

	var temp SearchScope
	// base is a fallback for a bogus scope, so
	// if the user did not originally request
	// base, we know they requested something
	// totally unsupported.
	if temp = strToScope(value); temp == noScope {
		err = bogusValueErr(TargetScope.String(), value)
		return
	}
	ex = temp

	return
}

/*
parsePermission is a private function called by Permission.Parse, et al.
*/
func parsePermission(raw string) (*permission, error) {
	perm, err := parser.ParsePermission(raw)
	if err != nil {
		return nil, err
	}

	return unpackageAntlrPermission(perm)
}

func unpackageAntlrPermission(perm parser.Permission) (*permission, error) {
	p := &permission{
		bool:  new(bool),  // disposition (ptr to bool)
		Right: new(Right), // rights specifiers (ptr to uint8 alias)
	}

	// !! WARNING - EXTREME SECURITY RISK !!
	//
	// A disposition has two (2) official settings and, thus,
	// is considered to be a MuTeX:
	//
	// - allow, which is expressed through a bool value of true
	// - deny, which is expressed through a bool value of false
	//
	// Ones initial thinking might lead to the conclusion that
	// a default of false is perfectly fine. But it shouldn't
	// take long for them to rethink that position, given the
	// following expression (or similar):
	//
	//   deny(none)
	//
	// Given the right (or wrong?) context, this could be bad.
	// Really, really bad. The above expression could return
	// as a result of parsing an instruction if a bogus, or
	// outright absent disposition was perceived and (as a
	// result of this failure) the default "Rights" specifiers
	// default to "none", which is only logical, right?
	//
	// BUT, because Golang (and most languages) defines implicit
	// defaults for certain types -- such as 0 for int and false
	// for bool -- any default is a bad idea here.
	//
	// Therefore a POINTER to a bool is used, both here in go-aci
	// AND within its sister package go-antlraci. go-antlraci will
	// evaluate/set the pointer using a double MuTeX case statement,
	// which allows only specific mutual-exclusion permutations that
	// are certain to avoid the above scenario.
	//
	// The ultimate disposition decision made by go-antlraci in this
	// case can be trusted, so long as the imported build is not some
	// fork from a source you don't know and trust.
	switch x := perm.Allow; x {
	case nil:
		// ambiguous result = fatal error
		return nil, noPermissionDispErr()
	default:
		(*p.bool) = *x
	}

	// process each permission one at a time
	var bits int // temporary storage for verification of bitshifted permission values
	var err error
	for i := 0; i < len(perm.Rights); i++ {
		rite := lc(perm.Rights[i])
		r, ok := rightsNames[rite]
		if !ok {
			err = rightNotfound(rite)
			return nil, err
		}
		bits |= int(r)
		p.shift(perm.Rights[i])
	}

	// The result of the above shifts MUST match the
	// same resulting bit value that would occur if
	// parsing was not involved. This accounts for
	// special values like 'none' and 'all' -- not
	// by simply looking for their presence as the
	// string literals that go-antlraci returns, but
	// rather through bit summation of the underlying
	// values defined in go-aci as part of its attempt
	// to be memory efficient.
	if bits != int(*p.Right) {
		err = unexpectedValueCountErr(`permission bits`, bits, int(*p.Right))
		return nil, err
	}

	return p, nil
}

func parsePermissionBindRule(raw string) (PermissionBindRule, error) {
	pbr, err := parser.ParsePermissionBindRule(raw)
	if err != nil {
		return badPermissionBindRule, err
	}

	return processPermissionBindRule(pbr)
}

func processPermissionBindRule(pbr parser.PermissionBindRule) (PermissionBindRule, error) {
	perm, err := unpackageAntlrPermission(pbr.P)
	if err != nil {
		return badPermissionBindRule, err
	}

	// traverse the native stackage.Stack instance returned
	// by antlraci and marshal its contents into proper
	// BindRule/BindRules instances, etc.
	rules, ok := convertBindRulesHierarchy(pbr.B)
	if !ok {
		return badPermissionBindRule, parseBindRulesHierErr(pbr.B, rules)
	}

	return PermissionBindRule{
		&permissionBindRule{
			P: Permission{perm},
			B: rules,
		},
	}, nil
}

func processPermissionBindRules(stack any) (PermissionBindRules, error) {
	var err error

	_pbrs, ok := castAsStack(stack)
	if !ok {
		err = errorf("Invalid input type %T; expecting stackage.Stack", stack)
		return badPermissionBindRules, err
	}

	var pbrs PermissionBindRules = PBRs()
	for i := 0; i < _pbrs.Len(); i++ {
		slice, _ := _pbrs.Index(i)
		_pbr, asserted := slice.(parser.PermissionBindRule)
		if !asserted {
			err = illegalSliceTypeErr(PermissionBindRules{}, `pbrs`, nil)
			return badPermissionBindRules, err
		}
		var pbr PermissionBindRule
		if pbr, err = processPermissionBindRule(_pbr); err != nil {
			return badPermissionBindRules, err
		}
		pbrs.Push(pbr)
	}

	return pbrs, nil
}

/*
Parse wraps go-antlraci's ParsePermissionBindRule function, writing
valid data into the receiver, or returning an error instance should
processing fail.
*/
func (r *PermissionBindRule) Parse(raw string) error {
	_r, err := parsePermissionBindRule(raw)
	if err != nil {
		return err
	}
	*r = _r

	return nil
}

/*
Parse wraps go-antlraci's ParsePermissionBindRules function, writing
valid data into the receiver, or returning an error instance should
processing fail.
*/
func (r *PermissionBindRules) Parse(raw string) error {
	_pbrs, err := parser.ParsePermissionBindRules(raw)
	if err != nil {
		return err
	}

	var _r PermissionBindRules
	if _r, err = processPermissionBindRules(_pbrs); err != nil {
		return err
	}
	*r = _r

	return err
}

/*
Parse wraps go-antlraci's ParseInstruction package-level function,
writing data into the receiver, or returning a non-nil instance of
error if processing should fail.

WARNING: Note that the act of successfully parsing an ACIv3 instruction
statement will clobber (overwrite) all of the contents present within the
receiver, if any.
*/
func (r *Instruction) Parse(raw string) (err error) {
	raw = condenseWHSP(raw) // get rid of leading/trailing/contiguous whitespace, newlines, et al.

	var (
		_r parser.Instruction  // instance returned by antlraci
		_i Instruction         // temporary container for assembly
		t  TargetRules         // stack for zero (0) or more TargetRule instances
		a  string              // Access Control Label (unique string label)
		p  PermissionBindRules // stack for one (1) or more PermissionBindRule instances
	)

	// hand the raw content to go-antlraci, where
	// the top-level instruction parser shall be
	// invoked, returning a struct containing the
	// three (2+) critical components for our new
	// ACIv3 instruction expression.
	if _r, err = parser.ParseInstruction(raw); err != nil {
		return
	}

	// obtain the ACL (string) value; if not
	// defined, fail the entire process.
	if a = _r.L.String(); len(a) == 0 {
		err = nilInstanceErr(_r.L)
		return
	}

	// process zero (0) or more TargetRules
	if t, err = processTargetRules(_r.T); err != nil {
		return
	}

	// process one (1) or more PermissionBindRules
	if p, err = processPermissionBindRules(_r.PB); err != nil {
		return
	}

	// set the target rules, acl and
	// pbr(s) within the temporary
	// Instruction instance.
	_i.Set(
		t,
		a,
		p,
	)

	// clobber receiver
	*r = _i

	return
}

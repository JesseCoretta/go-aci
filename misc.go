package aci

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/JesseCoretta/go-stackage"
)

const (
	// MultivalOuterQuotes represents the default quotation style
	// used by this package. In cases where a multi-valued BindRule
	// or TargetRule expression involving LDAP distinguished names,
	// ASN.1 Object Identifiers (in dot notation) and LDAP Attribute
	// Type names is being created, this constant will enforce only
	// outer-most double-quotation of the whole sequence of values.
	//
	// Example: keyword = "<val> || <val> || <val>"
	MultivalOuterQuotes = 0

	// MultivalSliceQuotes represents an alternative quotation scheme
	// offered by this package. In cases where a multi-valued BindRule
        // or TargetRule expression involving LDAP distinguished names,
        // ASN.1 Object Identifiers (in dot notation) and LDAP Attribute
        // Type names is being created, this constant shall disable outer
	// most quotation and will, instead, quote individual values. This
	// will NOT enclose symbolic OR (||) delimiters within quotations.
	//
	// Example: keyword = "<val>" || "<val>" || "<val>"
	MultivalSliceQuotes = 1
)

/*
ComparisonOperator constants defined within the stackage package are aliased
within this package for convenience, without the need for user-invoked stackage
package import.
*/
const (
	Eq stackage.ComparisonOperator = stackage.Eq // 0x1, "Equal To"
	Ne stackage.ComparisonOperator = stackage.Ne // 0x2, "Not Equal to"     !! USE WITH CAUTION !!
	Lt stackage.ComparisonOperator = stackage.Lt // 0x3, "Less Than"
	Le stackage.ComparisonOperator = stackage.Le // 0x4, "Less Than Or Equal"
	Gt stackage.ComparisonOperator = stackage.Gt // 0x5, "Greater Than"
	Ge stackage.ComparisonOperator = stackage.Ge // 0x6, "Greater Than Or Equal"
)

/*
ConditionPadding is a global variable that will be applies to ALL
Condition instances assembled during package operations. This is
a convenient alternative to manually invoking the NoPadding method
on a case-by-case basis.

Padding is enabled by default, and can be disabled here globally,
or overridden for individual Condition instances as needed.

Note that altering this value will not impact instances that were
already created; this only impacts the creation of new instances.
*/
var ConditionPadding bool = true

/*
RulePadding is a global variable that will be applies to ALL Rule
instances assembled during package operations. This is a convenient
alternative to manually invoking the NoPadding method on a case by
case basis.

Padding is enabled by default, and can be disabled here globally,
or overridden for individual Rule instances as needed.

Note that altering this value will not impact instances that were
already created; this only impacts the creation of new instances.
*/
var RulePadding bool = true

/*
frequently-accessed import function aliases.
*/
var (
	lc       func(string) string                 = strings.ToLower
	uc       func(string) string                 = strings.ToUpper
	eq       func(string, string) bool           = strings.EqualFold
	idxf     func(string, func(rune) bool) int   = strings.IndexFunc
	idxr     func(string, rune) int              = strings.IndexRune
	idxs     func(string, string) int            = strings.Index
	hasPfx   func(string, string) bool           = strings.HasPrefix
	hasSfx   func(string, string) bool           = strings.HasSuffix
	repAll   func(string, string, string) string = strings.ReplaceAll
	contains func(string, string) bool           = strings.Contains
	split    func(string, string) []string       = strings.Split
	trimS    func(string) string                 = strings.TrimSpace
	trimPfx  func(string, string) string         = strings.TrimPrefix
	join     func([]string, string) string       = strings.Join
	printf   func(string, ...any) (int, error)   = fmt.Printf
	sprintf  func(string, ...any) string         = fmt.Sprintf
	atoi     func(string) (int, error)           = strconv.Atoi
	isDigit  func(rune) bool                     = unicode.IsDigit
	isLetter func(rune) bool                     = unicode.IsLetter
	isLower  func(rune) bool                     = unicode.IsLower
	isUpper  func(rune) bool                     = unicode.IsUpper
	uint16g  func([]byte) uint16                 = binary.BigEndian.Uint16
	uint16p  func([]byte, uint16)                = binary.BigEndian.PutUint16
	valOf    func(x any) reflect.Value           = reflect.ValueOf
	typOf    func(x any) reflect.Type            = reflect.TypeOf

	stackOr   func(...int) stackage.Stack = stackage.Or
	stackAnd  func(...int) stackage.Stack = stackage.And
	stackNot  func(...int) stackage.Stack = stackage.Not
	stackList func(...int) stackage.Stack = stackage.List
)

/*
Version defines the ACI syntax version number implemented by
this package.
*/
const Version float32 = 3.0

/*
unquote removes leading and trailing quotation characters from
str.

This function considers any of ASCII #34 ("), ASCII #39 (') and
ASCII #96 (`) to be eligible candidates for truncation, though
only matches of the first and final slices are considered.
*/
func unquote(str string) string {
	if len(str) <= 2 {
		return str
	}

	// remove leading candidate
	switch c := rune(str[0]); c {
	case '"', '\'', '`':
		str = str[1:]
	}

	// remove trailing candidate
	switch c := rune(str[len(str)-1]); c {
	case '"', '\'', '`':
		str = str[:len(str)-1]
	}

	return str
}

/*
isQuoted looks at the first and final indices of the input
string value to determine whether BOTH are quotation ASCII
characters, and returns a Boolean value indicative of the
result. Any values found within unbalanced (e.g.: "Jesse')
OR incomplete (e.g.: "Jesse) quotation schemes will result
in a false return value.

This function considers any of ASCII #34 ("), ASCII #39 (')
and ASCII #96 (`) to be eligible candidates for quotation.
*/
func isQuoted(str string) bool {
	if len(str) < 2 {
		return false
	}

	var char rune

	// Perform a rune (Unicode char) switch
	switch q := rune(str[0]); q {

	// We've encountered a valid rune, make
	// a note of it.
	case '"', '\'', '`':
		char = q
	default:
		return false
	}

	// return the evaluation of the first rune
	// we found with the last one. They should
	// match if properly quoted.
	return char == rune(str[len(str)-1])
}

/*
isNumber returns a boolean value indicative of whether the provided value (which can be string or []byte instances)
represents a decimal number.
*/
func isNumber(val any) bool {
	var v []byte
	switch tv := val.(type) {
	case []byte:
		v = tv
	case string:
		v = []byte(tv)
	default:
		return false
	}

	if len(v) == 0 {
		return false
	}

	for i := 0; i < len(v); i++ {
		if !('0' <= rune(v[i]) && rune(v[i]) <= '9') {
			return false
		}
	}

	return true
}

func isAlnum(r rune) bool {
	return isLower(r) || isUpper(r) || isDigit(r)
}

/*
isIdentifier scans the input string val and judges whether
it appears to qualify as an identifier, in that:

- it begins with a lower alpha
- it contains only alphanumeric characters, hyphens or semicolons

This is used, specifically, it identify an LDAP attributeType (with
or without a tag), or an LDAP matchingRule.
*/
func isIdentifier(val string) bool {
	if len(val) == 0 {
		return false
	}

	// must begin with lower alpha.
	if !isLower(rune(val[0])) {
		return false
	}

	// can only end in alnum.
	if !isAlnum(rune(val[len(val)-1])) {
		return false
	}

	for i := 0; i < len(val); i++ {
		ch := rune(val[i])
		switch {
		case isAlnum(ch):
			// ok
		case ch == ';', ch == '-':
			// ok
		default:
			return false
		}
	}

	return true
}

/*
version returns the string version label for the ACI syntax.
*/
func version() string {
	return sprintf("version %.1f", Version)
}

/*
errorf wraps errors.New and returns a non-nil instance of error
based upon a non-nil/non-zero msg input value with optional args.
*/
func errorf(msg any, x ...any) error {
	switch tv := msg.(type) {
	case string:
		if len(tv) > 0 {
			return errors.New(sprintf(tv, x...))
		}
	case error:
		if tv != nil {
			return errors.New(sprintf(tv.Error(), x...))
		}
	}

	return nil
}

/*
strInSlice returns a boolean value indicative of whether the
specified string (str) is present within slice. Please note
that case is a significant element in the matching process.
*/
func strInSlice(str string, slice []string) bool {
	for i := 0; i < len(slice); i++ {
		if str == slice[i] {
			return true
		}
	}
	return false
}

/*
condenseWHSP returns input string b with all contiguous
WHSP characters condensed into single space characters.

WHSP is qualified through space or TAB chars (ASCII #32
and #9 respectively).
*/
func condenseWHSP(b string) (a string) {
	// remove leading and trailing
	// WHSP characters ...
	b = trimS(b)

	var last bool
	for i := 0; i < len(b); i++ {
		c := rune(b[i])
		switch c {
		// match space (32) or tab (9)
		case rune(32), rune(9):
			if !last {
				last = true
				a += string(c)
			}
		default:
			if last {
				last = false
			}
			a += string(c)
		}
	}

	return
}

/*
assertToD is called by timeOfDay.set for the purpose of
handling a potential clock time value for use in a Bind
Rule statement.

TODO: handle pure int w/o interpolation as binary.
*/
func assertToD(r *timeOfDay, t any) {
	switch tv := t.(type) {
	case time.Time:
		// time.Time input results in a recursive
		// run of this method.
		if tv.IsZero() {
			break
		}
		r.set(sprintf("%02d%02d", tv.Hour(), tv.Minute()))
	case string:
		// Handle discrepancy between ACI time, which ends
		// at 2400, and Golang Time, which ends at 2359.
		var offset int
		if tv == `2400` {
			tv = `2359` // so time.Parse doesn't flip
			offset = 41 // so we can use it as intended per ACI time syntax.
		}

		if _, err := time.Parse(`1504`, tv); err == nil {
			if n, err := atoi(tv); err == nil {
				x := make([]byte, 2)
				uint16p(x, uint16(n+offset))
				for i := 0; i < 2; i++ {
					(*r)[i] = x[i]
				}
			}
		}
	}
}

func chopACITerm(def string) string {
	if !hasSfx(def, `;)`) {
		return def
	}

	return def[:len(def)-2]
}

func isPowerOfTwo(x int) bool {
	return x&(x-1) == 0
}

// decom
func isParenthetical(x string) bool {
	if len(x) < 2 {
		return false
	}

	return x[0] == '(' && ')' == x[len(x)-1]
}

/*
keywordFromCategory attempts to locate the Category
method from input value r and, if found, runs it.

If a value is obtained, a resolution is attempted in
order to identify it as a BindKeyword or TargetKeyword
instance, which is then returned. Nil is returned in
all other cases.
*/
func keywordFromCategory(r any) Keyword {
	if r == nil {
		return nil
	}

	// if the instance has the Category
	// func, use reflect to get it.
	meth := getCategoryFunc(r)
	if meth == nil {
		return nil
	}

	var kw any

	// Try to match the category as a target rule
	// keyword context ...
	if tk := matchTKW(meth()); tk != TargetKeyword(0x0) {
		kw = tk
		return kw.(TargetKeyword)

		// Try to match the category as a bind rule
		// keyword context ...
	} else if bk := matchBKW(meth()); bk != BindKeyword(0x0) {
		kw = bk
		return kw.(BindKeyword)
	}

	return nil
}

/*
stackByDNKeyword returns an instance of BindDistinguishedNames based on the
following:

• BindGDN (groupdn) keyword returns the BindDistinguishedNames instance created by GDNs()

• BindRDN (roledn) keyword returns the BindDistinguishedNames instance created by RDNs()

• BindUDN (userdn) keyword returns the BindDistinguishedNames instance created by UDNs()

This function is private and is used only during parsing of bind and target rules
which permit a list of DNs as a single logical value. It exists mainly to keep
cyclomatics low.
*/
func stackByBDNKeyword(key Keyword) BindDistinguishedNames {
	// prepare a stack for our DN value(s)
	// based on the input keyword (key)
	switch key {
	case BindRDN:
		return RDNs()
	case BindGDN:
		return GDNs()
	}

	return UDNs()
}

/*
stackByDNKeyword returns an instance of TargetDistinguishedNames based on the
following:

• Target (target) keyword returns the TargetDistinguishedNames instance created by TDNs()

• TargetTo (target_to) keyword returns the TargetDistinguishedNames instance created by TTDNs()

• TargetFrom (target_from) keyword returns the TargetDistinguishedNames instance created by TFDNs()

This function is private and is used only during parsing of bind and target rules
which permit a list of DNs as a single logical value. It exists mainly to keep
cyclomatics low.
*/
func stackByTDNKeyword(key Keyword) TargetDistinguishedNames {
	// prepare a stack for our DN value(s)
	// based on the input keyword (key)
	switch key {
	case TargetTo:
		return TTDNs()
	case TargetFrom:
		return TFDNs()
	}

	return TDNs()
}

/*
stackByOIDKeyword returns an instance of ObjectIdentifiers based on the following:

• TargetExtOp (extop) keyword returns the ObjectIdentifiers instance created by ExtOps()

• TargetCtrl (targetcontrol) keyword returns the ObjectIdentifiers instance created by Ctrls()
*/
func stackByOIDKeyword(key Keyword) ObjectIdentifiers {
	// prepare a stack for our OID value(s)
	// based on the input keyword (key)
	switch key {
	case TargetExtOp:
		return ExtOps()
	}

	return Ctrls()
}

/*
castAsCondition merely wraps (casts, converts) and returns an
instance of BindRule -OR- TargetRule as a stackage.Condition
instance. This is useful for calling methods that have not been
extended (wrapped) in this package via go-stackage, as it may not
be needed in many cases ...

An instance submitted as x that is neither a BindRule or TargetRule
will result in an empty stackage.Condition return value.

Note this won't alter an existing BindRule or TargetRule instance,
rather a new reference is made through the stackage.Condition type
defined within go-stackage. The BindRule or TargetRule, once it has
been altered to one's satisfaction, can be sent off as intended and
this "Condition Counterpart" can be discarded, or left for GC.
*/
func castAsCondition(x any) (c *stackage.Condition) {
	switch tv := x.(type) {

	// case match is a single BindRule instance
	case BindRule:
		C := stackage.Condition(tv)
		return &C

	// case match is a single TargetRule instance
	case TargetRule:
		C := stackage.Condition(tv)
		return &C
	}

	return nil
}

/*
castAsStack merely wraps (casts, converts) and returns any type
alias of stackage.Stack as a native stackage.Stack.

This is useful for calling methods that have not been extended
(wrapped) in this package via go-stackage, as it might not be
needed in most cases ...

An instance submitted as x that is NOT a type alias of stackage.Stack
will result in an empty stackage.Stack return value.

Note this won't alter an existing values, rather a new reference is
made through the stackage.Condition type defined within go-stackage.
The alias type, once it has been altered to one's satisfaction, can be
sent off as intended and this "Stack Counterpart" can be discarded, or
left for GC.
*/
func castAsStack(u any) (S stackage.Stack, converted bool) {
	switch tv := u.(type) {

	case ObjectIdentifiers:
		converted = true
		S = stackage.Stack(tv)

	case BindDistinguishedNames,
		TargetDistinguishedNames:
		S, converted = castDNRules(tv)

	case BindRules, TargetRules,
		PermissionBindRules:
		S, converted = castBTRules(tv)

	case AttributeTypes:
		converted = true
		S = stackage.Stack(tv)

	case AttributeFilterOperation,
		AttributeFilterOperations:
		S, converted = castFilterRules(tv)
	}

	return
}

func castBTRules(x any) (S stackage.Stack, converted bool) {
	switch tv := x.(type) {
	case BindRules:
		S = stackage.Stack(tv)
		converted = true
	case TargetRules:
		S = stackage.Stack(tv)
		converted = true
	case PermissionBindRules:
		S = stackage.Stack(tv)
		converted = true
	}

	return
}

func castDNRules(x any) (S stackage.Stack, converted bool) {
	switch tv := x.(type) {
	case BindDistinguishedNames:
		S = stackage.Stack(tv)
		converted = true
	case TargetDistinguishedNames:
		S = stackage.Stack(tv)
		converted = true
	}

	return
}

func castFilterRules(x any) (S stackage.Stack, converted bool) {
	switch tv := x.(type) {
	case AttributeFilterOperation:
		S = stackage.Stack(tv)
		converted = true
	case AttributeFilterOperations:
		S = stackage.Stack(tv)
		converted = true
	}

	return
}

/*
getCategoryFunc uses reflect to obtain and return a given
type instance's Category method, if present. If not, nil
is returned.
*/
func getCategoryFunc(x any) func() string {
	v := valOf(x)
	if v.IsZero() {
		return nil
	}

	method := v.MethodByName(`Category`)
	if method.Kind() == reflect.Invalid {
		return nil
	}

	if meth, ok := method.Interface().(func() string); ok {
		return meth
	}

	return nil
}

/*
BindContext is a convenient interface type that is qualified by
the following types:

• BindRule

• BindRules

The qualifying methods shown below are intended to make the
handling of a structure of (likely nested) BindRules instances
slightly easier without an absolute need for type assertion at
every step. These methods are inherently read-only in nature
and represent only a subset of the available methods exported
by the underlying qualifier types.

To alter the underlying value, or to gain access to all of a
given type's methods, type assertion shall be necessary.
*/
type BindContext interface {
	// String returns the string representation of the
	// receiver instance.
	String() string

	// Keyword returns the BindKeyword, enveloped as a
	// Keyword interface value. If the receiver is an
	// instance of BindRule, the value is derived from
	// the Keyword method. If the receiver is an instance
	// of BindRules, the value is derived (and resolved)
	// using the Category method.
	Keyword() Keyword

	// IsZero returns a Boolean value indicative of the
	// receiver instance being nil, or unset.
	IsZero() bool

	// Len returns the integer length of the receiver.
	// Only meaningful when run on BindRules instances.
	Len() int

	// IsNesting returns a Boolean value indicative of
	// whether the receiver contains a stack as a value.
	// Only meaningful when run on BindRules instances.
	IsNesting() bool

	// Category will report `bind` in all scenarios.
	Category() string

	// Kind will report `stack` for a BindRules instance, or
	// `condition` for a BindRule instance
	Kind() string

	// isBindContextQualifier ensures no greedy interface
	// matching outside of the realm of bind rules. It need
	// not be accessed by users, nor is it run at any time.
	isBindContextQualifier() bool
}

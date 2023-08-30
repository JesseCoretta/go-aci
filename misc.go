package aci

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// AttributeFilterOperationsCommaDelim represents the default
// delimitation scheme offered by this package. In cases where
// the AttributeFilterOperations type is used to represent any
// TargetRule bearing the targattrfilters keyword context, two
// (2) different delimitation characters may be permitted for
// use (depending on the product in question).
//
// Use of this constant allows the use of a comma (ASCII #44) to
// delimit the slices in an AttributeFilterOperations instance as
// opposed to the alternative delimiter (semicolon, ASCII #59).
//
// This constant may be fed to the SetDelimiter method that is
// extended through the AttributeFilterOperations type.
const AttributeFilterOperationsCommaDelim = 0

// AttributeFilterOperationsSemiDelim represents an alternative
// delimitation scheme offered by this package. In cases where
// the AttributeFilterOperations type is used to represent any
// TargetRule bearing the targattrfilters keyword context, two
// (2) different delimitation characters may be permitted for
// use (depending on the product in question).
//
// Use of this constant allows the use of a semicolon (ASCII #59)
// to delimit the slices in an AttributeFilterOperations instance
// as opposed to the default delimiter (comma, ASCII #44).
//
// This constant may be fed to the SetDelimiter method that is
// extended through the AttributeFilterOperations type.
const AttributeFilterOperationsSemiDelim = 1

// MultivalOuterQuotes represents the default quotation style
// used by this package. In cases where a multi-valued BindRule
// or TargetRule expression involving LDAP distinguished names,
// ASN.1 Object Identifiers (in dot notation) and LDAP Attribute
// Type names is being created, this constant will enforce only
// outer-most double-quotation of the whole sequence of values.
//
// Example: keyword = "<val> || <val> || <val>"
//
// This constant may be fed to the SetQuoteStyle method that is
// extended through eligible types.
const MultivalOuterQuotes = 0

// MultivalSliceQuotes represents an alternative quotation scheme
// offered by this package. In cases where a multi-valued BindRule
// or TargetRule expression involving LDAP distinguished names,
// ASN.1 Object Identifiers (in dot notation) and LDAP Attribute
// Type names is being created, this constant shall disable outer
// most quotation and will, instead, quote individual values. This
// will NOT enclose symbolic OR (||) delimiters within quotations.
//
// Example: keyword = "<val>" || "<val>" || "<val>"
//
// This constant may be fed to the SetQuoteStyle method that is
// extended through eligible types.
const MultivalSliceQuotes = 1

/*
RulePadding is a global variable that will be applies to ALL
TargetRule and Bindrule instances assembled during package operations.
This is a convenient alternative to manually invoking the NoPadding
method on a case-by-case basis.

Padding is enabled by default, and can be disabled here globally,
or overridden for individual TargetRule/BindRule instances as needed.

Note that altering this value will not impact instances that were
already created; this only impacts the creation of new instances.
*/
var RulePadding bool = true

/*
StackPadding is a global variable that will be applies to ALL Stack
instances assembled during package operations. This is a convenient
alternative to manually invoking the NoPadding method on a case by
case basis.

Padding is enabled by default, and can be disabled here globally,
or overridden for individual Stack instances as needed.

Note that altering this value will not impact instances that were
already created; this only impacts the creation of new instances.
*/
var StackPadding bool = true

/*
frequently-accessed import function aliases.
*/
var (
	lc       func(string) string                 = strings.ToLower
	uc       func(string) string                 = strings.ToUpper
	eq       func(string, string) bool           = strings.EqualFold
	ctstr    func(string, string) int            = strings.Count
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
	itoa     func(int) string                    = strconv.Itoa
	atoi     func(string) (int, error)           = strconv.Atoi
	isDigit  func(rune) bool                     = unicode.IsDigit
	isLetter func(rune) bool                     = unicode.IsLetter
	isLower  func(rune) bool                     = unicode.IsLower
	isUpper  func(rune) bool                     = unicode.IsUpper
	uint16g  func([]byte) uint16                 = binary.BigEndian.Uint16
	uint16p  func([]byte, uint16)                = binary.BigEndian.PutUint16
	valOf    func(x any) reflect.Value           = reflect.ValueOf
	typOf    func(x any) reflect.Type            = reflect.TypeOf
)

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

func isPowerOfTwo(x int) bool {
	return x&(x-1) == 0
}

/*
keywordFromCategory attempts to locate the Category
method from input value r and, if found, runs it.

If a value is obtained, a resolution is attempted in
order to identify it as a BindKeyword or TargetKeyword
instance, which is then returned. Nil is returned in
all other cases.

DECOM me (someday)
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
isPtr returns a boolean value indicative of whether kind
reflection revealed the presence of a pointer type.
*/
func isPtr(x any) bool {
	return typOf(x).Kind() == reflect.Ptr
}

/*
isIUint returns a Boolean value of true if x represents a
member of the integer / unsigned integer "family". Any size
is allowed, so long as it is a built-in primitive.

If a (valid) member is a pointer reference, it is dereferenced
and examined just the same.

Floats and complexes are ineligible and will return false as they
are not used in this package. Additionally, non-numerical types
shall return false. This would include structs, strings, maps, etc.
*/
func isIUint(x any) (is bool) {
	// create a reflect.Type abstract
	// instance using raw input x.
	X := typOf(x)

	// disenvelop the instance if
	// it is a pointer reference.
	if isPtr(x) {
		X = X.Elem()
	}

	// perform a reflect.Kind switch upon
	// reflect.Type instance X ...
	switch k := X.Kind(); k {

	// allow only the following "kinds":
	case reflect.Int, reflect.Uint,
		reflect.Int8, reflect.Uint8,
		reflect.Int16, reflect.Uint16,
		reflect.Int32, reflect.Uint32,
		reflect.Int64, reflect.Uint64,
		reflect.Uintptr:
		is = true
	}

	return
}

/*
getBitSize returns the max bit length capacity
for a given type.

Note this will only return a meaningful value if
x represents a numerical type, such as Day, Right
or Level (all of which are subject to bit shifts).
Passing inappropriate type instances, such as a
struct, string, etc., will return zero (0).

This function uses the reflect.Size method (and
thus unsafe.Sizeof) to obtain a uintptr, which
will be cast as an int, multiplied by eight (8)
and finally returned.
*/
func bitSize(x any) (bits int) {
	// create a reflect.Type abstract
	// instance using raw input x.
	X := typOf(x)

	// disenvelop the instance if
	// it is a pointer reference.
	if isPtr(x) {
		X = X.Elem()
	}

	// see if the instance is an int
	// or uint (or a variant of same)
	if isIUint(x) {
		bits = int(X.Size()) * 8
	}

	return
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

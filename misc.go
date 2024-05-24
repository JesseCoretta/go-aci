package aci

/*
misc.go contains fundamental "worker functions" used elsewhere in
this package, as well as certain constants and variables unsuitable
for placement anywhere else.
*/

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

/*
AttributeFilterOperationsCommaDelim shall invoke the default delimitation character offered by this package for use within instances of the [AttributeFilterOperations] type.

In cases where the [AttributeFilterOperations] type is used to represent any [TargetRule] bearing the [TargetAttrFilters] [TargetKeyword] context, one (1) of two (2) different delimiter characters MAY be permitted for use, depending on which directory product is in use.

Use of this constant allows the use of a comma (ASCII #44) to delimit the slices in an [AttributeFilterOperations] instance as opposed to the alternative delimiter (semicolon, ASCII #59).

This constant may be fed to the [AttributeFilterOperations.SetDelimiter] method.

This is the default setting for the above scenario, and likely need to be specified manually unless reverting back from the alternative character.
*/
const AttributeFilterOperationsCommaDelim = 0

/*
[AttributeFilterOperationsSemiDelim] invokes the alternative delimitation character offered by this package for use within instances of the [AttributeFilterOperations] type.

In cases where the [AttributeFilterOperations] type is used to represent any [TargetRule] bearing the [TargetAttrFilters] [TargetKeyword] context, one (1) of two (2) different delimiter characters MAY be permitted for use, depending on which directory product is in use.

Use of this constant allows the use of a semicolon (ASCII #59) to delimit the slices in an [AttributeFilterOperations] instance as opposed to the default delimiter (comma, ASCII #44).

This constant may be fed to the [AttributeFilterOperations.SetDelimiter] method.
*/
const AttributeFilterOperationsSemiDelim = 1

/*
MultivalOuterQuotes represents the alternate quotation style used by this package. In cases where a multi-valued [BindRule] or [TargetRule] expression involving distinguished names, ASN.1 Object Identifiers (in dot notation) and LDAP Attribute Type names is being created, this constant will enforce only outer-most double-quotation of the whole sequence of values, including the delimiters.

	Example: keyword = "<val> || <val> || <val>"

This constant may be fed to the SetQuoteStyle method that is extended through eligible types.
*/
const MultivalOuterQuotes = 1

/*
MultivalSliceQuotes represents the standard quotation scheme offered by this package. In cases where a multi-valued [BindRule] or [TargetRule] expression involving distinguished names, ASN.1 Object Identifiers (in dot notation) and LDAP Attribute Type names is being created, this constant shall disable outermost quotation and will, instead, quote individual values. This will NOT enclose symbolic OR (||) delimiters within quotations.

	Example: keyword = "<val>" || "<val>" || "<val>"

This constant may be fed to the SetQuoteStyle method that is extended through eligible types.
*/
const MultivalSliceQuotes = 0

/*
RulePadding is a global variable that will be applies to ALL [TargetRule] and [BindRule] instances assembled during package operations. This is a convenient alternative to manually invoking the NoPadding method on a case-by-case basis.

Padding is enabled by default, and can be disabled here globally, or overridden for individual [TargetRule]/[BindRule] instances as needed.

Note that altering this value will not impact instances that were already created; this only impacts the creation of new instances.
*/
var RulePadding bool = true

/*
StackPadding is a global variable that will be applies to ALL [stackage.Stack] instances assembled during package operations. This is a convenient alternative to manually invoking the NoPadding method on a case by case basis.

Padding is enabled by default, and can be disabled here globally, or overridden for individual [stackage.Stack] instances as needed.

Note that altering this value will not impact instances that were already created; this only impacts the creation of new instances.
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

/*
isAlnum returns a Boolean value indicative of whether rune r represents
an alphanumeric character. Specifically, one (1) of the following ranges
must evaluate as true:

  - 0-9 (ASCII characters 48 through 57)
  - A-Z (ASCII characters 65 through 90)
  - a-z (ASCII characters 97 through 122)
*/
func isAlnum(r rune) bool {
	return isLower(r) || isUpper(r) || isDigit(r)
}

/*
isIdentifier scans the input string val and judges whether
it appears to qualify as an identifier, in that:

- it begins with a lower alpha
- it contains only alphanumeric characters, hyphens or semicolons (for tagged attributes)

This is used, specifically, it identify an LDAP attributeType (with
or without a tag), or an LDAP matchingRule.
*/
func isIdentifier(val string) bool {
	if len(val) == 0 {
		return false
	}

	// must begin with lower alpha.
	if !isLetter(rune(val[0])) {
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
strInSlice returns a Boolean value indicative of whether the
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
strInSliceFold returns a Boolean value indicative of whether
the specified string (str) is present within slice. Case is
not significant in the matching process.
*/
func strInSliceFold(str string, slice []string) bool {
	for i := 0; i < len(slice); i++ {
		if eq(str, slice[i]) {
			return true
		}
	}
	return false
}

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
isPowerOfTwo returns a Boolean value indicative of whether int x
represents a value that is a power of two (2). Mostly used in the
act of bitshifting certain (eligible) composite values.
*/
func isPowerOfTwo(x int) bool {
	return x&(x-1) == 0
}

/*
Hash computes a SHA-1 hash value, derived from the String method output (or string value) of input value x.

The hash, if generated, is cast as a string prior to being returned alongside an error.

Input value x must qualify as one (1) of the following:

  - Must be a string in of itself, OR ...
  - Must be a type instance that has its own stringer method

Failing the above, a non-nil error instance is returned.

This package-level function is the basis for all Compare methods extended by myriad types throughout this package. In most cases, using an available Compare method is easier than using this function directly.

The hash evaluation offered by this package is meant to act as a supplement in a change review process or similar.  The return value should not be used to gauge "validity", "nilness" or "initialization status" of an instance. In certain cases, two (2) dissimilar (and invalid!) instances of the same type shall evaluate as "equal". When string representation yields the same effective value for two (2) type instances, this is both guaranteed and expected behavior.
*/
func Hash(x any) (string, error) {
	return hashInstance(x)
}

func compareHashInstance(r, x any) bool {
	var rh, xh string
	var err error

	if rh, err = hashInstance(r); err != nil {
		return false
	} else if xh, err = hashInstance(x); err != nil {
		return false
	}

	return rh == xh
}

/*
hashInstance is a private function called by the Hash package
level function. It uses crypto/sha1 to compute a hash value
derived from the string representation of input value x, (in
the event it possesses its own String method, or if the value
itself is a string).

A string representation of the hash value alongside an error
instance are returned.
*/
func hashInstance(x any) (s string, err error) {
	var _s string
	switch tv := x.(type) {
	case string:
		_s = tv
	default:
		meth := getStringFunc(x)
		if meth == nil {
			err = errorf("%T instance has no String method; cannot compute hash", x)
			return
		}

		if _s = meth(); len(_s) == 0 {
			err = errorf("%T instance produced a zero length string, cannot compute hash", x)
			return
		}
	}

	s = uc(sprintf("%x", sha1.Sum([]byte(_s))))

	return
}

/*
getStringFunc uses reflect to obtain and return a given
type instance's String method, if present. If not, nil
is returned.
*/
func getStringFunc(x any) (meth func() string) {
	if x == nil {
		return
	}

	if v := valOf(x); !v.IsZero() {

		method := v.MethodByName(`String`)
		if method.Kind() == reflect.Invalid {
			return nil
		}

		if _meth, ok := method.Interface().(func() string); ok {
			meth = _meth
		}
	}

	return
}

/*
isPtr returns a Boolean value indicative of whether kind
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
	if x == nil {
		return
	}

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
condenseWHSP returns input string b with all contiguous
WHSP characters condensed into single space characters.

WHSP is qualified through space or TAB chars (ASCII #32
and #9 respectively).
*/
func condenseWHSP(b string) (a string) {
	// remove leading and trailing
	// WHSP characters ...
	b = trimS(b)
	b = repAll(b, string(rune(10)), string(rune(32)))

	var last bool
	for i := 0; i < len(b); i++ {
		c := rune(b[i])
		switch c {
		// match space (32) or tab (9)
		case rune(9), rune(10), rune(32):
			if !last {
				last = true
				a += string(rune(32))
			}
		default:
			if last {
				last = false
			}
			a += string(c)
		}
	}

	a = trimS(a) //once more
	return
}

package aci

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
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
isAVAOperatorComponent returns a boolean value indicative of whether
the string input val conforms to a valid LDAP AttributeValueAssertion
comparison operator character (e.g.: `=`) or perhaps a complete token,
e.g.: '~=', ':dn:2.5.14.3:=', et al. Note that the return value should
not be taken as a definitive answer, considering values or keywords
might contain some of the very same characters. This function merely
determines whether val qualifies the desired "signature", but does NOT
know (in context) that this is its true nature.
*/
func isAVAOperatorComponent(val string) bool {
	if len(val) == 0 {
		return false
	}

	// equality, order and approx,
	for _, v := range []string{
		`=`,
		`<=`,
		`>=`,
		`~=`,
		`:`,
		`dn`,
	} {
		if val == v {
			return true
		}
	}

	switch {
	case isDotNot(val), isIdentifier(val):
		return true

	}

	// does not conform
	return false
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


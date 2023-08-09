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

var ops []string = []string{
	`AND NOT`,
	`NOT`, `!`,
	`OR`, `||`, `|`,
	`AND`, `&&`, `&`,
}

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
np (no parenthesis) strips all left "(" and right ")" parenthesis
from input value e, and returns the altered string.
*/
func np(e string) (up string) {
	for i := 0; i < len(e); i++ {
		c := rune(e[i])
		switch c {
		case '(', ')':
			continue
		default:
			up += string(e[i])
		}
	}

	return
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
runeInStr scans the string input (str) for rune r. A boolean
value is returned indicative of a match.
*/
func runeInStr(r rune, str string) bool {
	for i := 0; i < len(str); i++ {
		if rune(str[i]) == r {
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
removeNewlines replaces all newline characters (ASCII #10)
with a single WHSP character (ASCII #32).
*/
func removeNewlines(b string) string {
	return repAll(b,
		string(rune(10)),
		string(rune(32)))
}

/*
nextRune will attempt to return idx+1's rune as extracted
from the raw input string. If successful, a non-null rune
along with a positive index are returned. If unsuccessful
a null rune and negative one (-1) index are returned.
*/
func nextRune(raw string, idx int) (rune, int) {
	n := -1
	if idx >= len(raw)-1 {
		return rune(0), n
	}

	return rune(raw[idx+1]), idx + 1
}

/*
nextRune will attempt to return idx-1's rune as extracted
from the raw input string. If successful, a non-null rune
along with a positive index are returned. If unsuccessful
a null rune and negative one (-1) index are returned.
*/
func lastRune(raw string, idx int) (rune, int) {
	n := -1
	if idx <= 0 {
		return rune(0), n
	}

	return rune(raw[idx-1]), idx - 1
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

/*
trimLRParen returns the input string minus outer parentheses.
If parentheses were actually trimmed, a boolean of true is
returned. Leading/Trailing WHSP characters are removed in any
scenario.
*/
func trimLRParen(raw string) (val string, trimmed bool) {
	if len(raw) < 2 {
		// below minimum usable length
		val = trimS(raw)
		return
	} else if rune(raw[0]) != '(' || rune(raw[len(raw)-1]) != ')' {
		// just as a courtesy, trim L/T WHSP chars
		val = trimS(raw)
		trimmed = len(val) < len(raw)
		return
	}

	// trim any L/T WHSP chars left over ...
	val = trimS(raw[1 : len(raw)-1])

	// we should be shorter than the original.
	trimmed = len(val) < len(raw)
	return
}

/*
keywords are strictly alphabetical (and, normalized to lower
case). This function returns a boolean value which indicates
whether rune c is a lower alpha char.
*/
func kwIdxFunc(c rune) bool {
	return isLower(c)
}

func isBoolOp(c any) bool {
	switch tv := c.(type) {
	case string:
		return tv == `&` || tv == `|` || tv == `!`
	case rune:
		return tv == '&' || tv == '|' || tv == '!'
	}

	return false
}

func isExprDelim(s string) bool {
	return s == `||` || s == `&&`
}

func isWordOp(w string) bool {
	switch W := lc(w); W {
	case `and`, `or`, `and not`:
		return true
	}

	return false
}

/*
parenState analyzes the parenthetical nature of input string x.
It will scan for parenthetical characters (ASCII #40 and #41
respectively) and will return information describing its findings.

Scans shall not occur within quoted values, and such values are
skipped until quotation ends.

Integer variables for total and pairs, along with Boolean values
for outer and balanced, are returned.

A total value of zero (0) indicates no parentheticals were found.

A pairs value of zero (0) indicates no parentheticals were found. A
pairs value of one (1) indicates that a single left parenthesis '('
and right parenthesis ')' found found, thus indicating a "pair".

Neither total nor pairs shall ever be negative.

An outer value of true indicates that, in addition to at least one
(1) balanced pair of parentheticals being present, the input string
x is parenthetical on its outermost ends (in other words, x begins
with a left parenthetical, and ends with a right parenthetical).

A balanced value of true indicates the count is even numbered, and
thus no unmatched parentheticals were observed.
*/
func parenState(x string) (total, pairs int, outer, balanced bool) {
	// L and R integers record the progressive
	// counts of left and right parenthetical
	// character occurrences.
	var L, R int

	// Can't process a null string, so bail out
	// early and declare balance.
	if len(x) == 0 {
		balanced = true
		return
	}

	// Keep track of our quotation state.
	var q bool

	// Iterate each character found within
	// input string x.
	for i := 0; i < len(x); i++ {

		// Perform a rune switch (c) upon
		// the current character to begin
		// case matching ...
		switch c := rune(x[i]); c {

		// We found a left parenthetical
		case '(':
			if !q {
				if L++; L == R {
					pairs++
				}
			}

		// We found a right parenthetical
		case ')':
			if !q {
				if R++; R == L {
					pairs++
				}
			}

		// Quote characters detected. We
		// will toggle the quote state. A
		// state of true shall result in
		// parentheticals being ignores
		// until the state returns to false.
		case '"', '\'', '`':
			q = !q
		}
	}

	// Sum our total for all parenthetical occurrences
	// regardless of balance. Don't bother going any
	// further if the sum total turns out to be zero
	// (0), as none of the other checks would return
	// a meaningful result ...
	if total = L + R; total == 0 {
		balanced = true
		return
	}

	// See if we found at least one (1) parenthetical
	// pair. If so, perform outer and balance checks.
	if pairs > 0 {

		// Assuming we found at least one (1) balanced pair
		// of parentheticals, check to see if the entirety
		// of input string x is parenthetical itself.
		outer = (x[0] == '(' && x[len(x)-1] == ')')
		balanced = (L == R)
	}

	return
}

/*
operators can conceivably contain any characters other than
WHSP. This function returns a boolean value which indicates
whether rune c is such a char.  A value of true indicates a
WHSP character was encountered, false otherwise. WHSP chars
should be interpreted as an indication that the operator is
done.
*/
func opIdxFunc(c rune) bool {
	return c == rune(32)
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

func getLastNextToken(tokens []string, index int) (last, next string) {
	if index != 0 && index-1 < len(tokens) {
		last = tokens[index-1]
	}

	if index+1 < len(tokens) {
		next = tokens[index+1]
	} else {
		printf("Can't get token for next index %d [%d]\n", index, len(tokens))
	}

	return
}

func tokenIsAnchor(token string, chop, index int) (anchor bool) {
	anchor = token == `version 3.0; acl`
	return
}

/*
isBS is a simple function used to identify a parenthetical boolean stack
(as opposed to a single condition) in a Bind Rule.

A parenthetical boolean stack MUST begin with an opening parenthesis,
and have a sequence of one (1) or more condition statements delimited
by a boolean word operator. The length of the stack does not matter,
so long as it follows said format:

	( cond word cond word ... )
*/
func isBS(t []string) (ct int, ok bool) {
	if len(t) < 5 {
		return
	}

	isKW := func(o string) bool {
		return matchBKW(lc(o)) != BindKeyword(0x0)
	}

	// convenient comparison operator token
	// recognizer func.
	isTokOp := func(o string) (ok bool) {
		_, ok = matchOp(o)
		return
	}

	var last string
	var is bool
	for i := 0; i < len(t); i++ {
		tok := t[i]
		if i > 0 {
			last = t[i-1]
		}

		switch {
		case tok == `)`:
			if is = ct > 0; !is {
				return
			}
			if isQuoted(last) {
				ct++
			}
		case tok == `(`:
			// ok
		case isKW(tok):
			if !(last == `(` || isWordOp(last) || i == 0) {
				return
			}
		case isTokOp(tok):
			if !isKW(last) {
				return
			}
		case isExprDelim(tok):
			if !isQuoted(last) {
				return
			}
		case isWordOp(tok):
			if !isQuoted(last) {
				return
			}
			ct++ // complete condition +1
		case isQuoted(tok):
			if !(isTokOp(last) || isExprDelim(last)) {
				return
			}
		default:
			return
		}

		if is {
			break
		}
	}

	ok = is // closer is handled in 'is'
	return
}

/*
isBPC is a simple function used to identify a parenthetical condition
(as opposed to a STACK) in a Bind Rule. This function need not be used
for Target Rule parsing.

isBPC looks at the provided token stream (t) and, assuming it contains
five (5) or more slices, attempts to determine whether the first token
is an opening parenthesis (ASCII #40), the second token is a known bind
keyword, and the fifth token is the closing parenthesis (ASCII #41). A
boolean value indicative of ALL of these conditions evaluating as true
is returned.

Assuming a closing parenthesis is NOT the fifth character as expected,
this function will take extra measures to ensure that the expression
value is not a sequence of quoted values. In this case, all of this
values should be [re]interpreted as a single value. When this happens,
the closing parenthetical check is run again.

In short:

	"value || value || value"     // this is definitely a single value, and poses no trouble
	"value" || "value" || "value" // this SHOULD be single, but requires extra handling
*/
func isBPC(t []string) bool {
	if len(t) < 5 {
		return false
	}

	isCloser := func(t []string) bool {
		skip, _ := scanMultival(t)
		return t[4] == `)` || t[skip] == `)`
	}

	// collect the evaluation results
	opener := t[0] == `(`                        // is opening parenthesis
	bindkw := matchBKW(t[1]) != BindKeyword(0x0) // is (legit) bind keyword
	_, oper := matchOp(t[2])                     // is comparison operator
	closer := isCloser(t)                        // is closing parenthesis

	return (opener && bindkw && oper && closer) // return the combined result
}

/*
isBC returns a boolean value indicative of whether the provided
token stream (t) initially contains a bind rule condition. This
is opposed to a parenthetical bind rule condition.
*/
func isBC(t []string) bool {
	// chop the terminator, as we won't need it
	if t[len(t)-1] == `)` && t[len(t)-2] == `;` {
		t = t[:len(t)-2]
	}

	if len(t) < 3 {
		// nothing left over, or too
		// small to begin with.
		return false
	}

	// collect the evaluation results
	bindkw := matchBKW(t[0]) != BindKeyword(0x0) // is (legit) bind keyword
	_, oper := matchOp(t[1])                     // is comparison operator
	quotedv := isQuoted(t[2])                    // at least one (1) quoted value

	return bindkw && oper && quotedv // return the combiend result
}

/*
scanMultival reads the provided token stream in an attempt to find
the closing parenthetical character of the current expression. This
function is solely used by isPC and exists only to keep cyclomatics
low.
*/
func scanMultival(t []string) (skip int, ok bool) {
	// Begin at three (3) since we're fairly
	// certain that is where the first value
	// resides.
	for i := 3; i < len(t); i++ {
		// Perform token switch
		switch tok := t[i]; tok {

		// found it!
		case `)`:
			ok = true
			return

			// Increment for delimiter, but
			// continue looping.
		case `||`:
			skip++

			// this HAS to be a value. If not
			// something is screwy and we can
			// go no further.
		default:
			// Yeah, something screwy is
			// going on.
			if !isQuoted(tok) || isWordOp(tok) {
				skip = 0
				return
			}

			// Value was quoted, increment
			// counter and continue.
			skip++
		}
	}

	skip = 0
	return
}

func tokensAreTerminators(t []string) bool {
	if len(t) != 2 {
		return false
	}
	return t[1] == `)` &&
		t[0] == `;`
}

// Find (and remember) the first (1st)
// Boolean WORD operator encountered.
//
// If no Boolean WORD operator was
// encountered, just fallback to AND
// for convenience.
func hasOp(t []string) (op string, found bool) {
	op = `AND`
	for i := 0; i < len(t); i++ {
		if found = isWordOp(t[i]); found {
			// a known Boolean WORD operator has
			// been found; create the outer stack
			// accordingly.
			op = t[i]
			break
		}
	}
	return
}

func readQuotedValues(t []string) (stop int, values []string) {
	for i := 0 ;i < len(t); i++ {
		switch t[i]{
		case `||`:
			continue
		default:
			if !isQuoted(t[i]) {
				stop = i-1
				return
			}
			values = append(values, t[i])
		}
	}

	return
}

/*
objectString is a stringer caller, allowing the execution and
return of a stringer method value without manual type assertion.
*/
func objectString(x any) (str string) {
        var m, rv reflect.Value

        // See if x is nil
        if rv = reflect.ValueOf(x); rv.IsZero() {
                return
        }

        // Call the desired method, or fail with an error.
	m = rv.MethodByName(`String`)
        meth, ok := m.Interface().(func() string)
        if !ok {
                return
        }

        str = meth()
	return
}

/*
objectString is a stringer caller, allowing the execution and
return of the Category method value without manual type assertion.
*/
func objectCategory(x any) (str string) {
        var m, rv reflect.Value

        // See if x is nil
        if rv = reflect.ValueOf(x); rv.IsZero() {
                return
        }

        // Call the desired method, or fail with an error.
        m = rv.MethodByName(`Category`)
        meth, ok := m.Interface().(func() string)
        if !ok {
                return
        }

        str = meth()
        return
}

/*
objectIdent is a stringer caller, allowing the execution and
return of the ID method value without manual type assertion.
*/
func objectIdent(x any) (str string) {
        var m, rv reflect.Value

        // See if x is nil
        if rv = reflect.ValueOf(x); rv.IsZero() {
                return
        }

        // Call the desired method, or fail with an error.
        m = rv.MethodByName(`ID`)
        meth, ok := m.Interface().(func() string)
        if !ok {
                return
        }

        str = meth()
        return
}

func unwrapRule(x Rule) (r Rule) {
	if x.Len() != 1 {
		r = x
		return
	}

	sl, _ := x.Index(0)
	switch tv := sl.(type) {
	case Rule:
		if tv.Len() == 1 {
			sl2, _ := tv.Index(0)
			switch uv := sl2.(type) {
			case Rule:
				r = unwrapRule(uv)
				r.setCategory(x.Category())
			}
		} else if tv.Len() > 1 {
			return tv
		}
	}

	if r.Len() ==0 {
		r = x
	}

	return
}

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
frequently-accessed import function aliases.
*/
var (
	lc       func(string) string                  = strings.ToLower
	uc       func(string) string                  = strings.ToUpper
	eq       func(string,string) bool             = strings.EqualFold
	idxf     func(string, func(rune) bool) int    = strings.IndexFunc
	idxr     func(string, rune) int               = strings.IndexRune
	idxs     func(string, string) int             = strings.Index
	hasPfx   func(string, string) bool            = strings.HasPrefix
	hasSfx   func(string, string) bool            = strings.HasSuffix
	repAll   func(string, string, string) string  = strings.ReplaceAll
	contains func(string, string) bool            = strings.Contains
	split    func(string, string) []string        = strings.Split
	trimS    func(string) string                  = strings.TrimSpace
	join     func([]string, string) string        = strings.Join
	printf   func(string, ...any) (int, error)    = fmt.Printf
	sprintf  func(string, ...any) string          = fmt.Sprintf
	atoi     func(string) (int, error)            = strconv.Atoi
	isDigit  func(rune) bool                      = unicode.IsDigit
	isLetter func(rune) bool                      = unicode.IsLetter
	isLower  func(rune) bool		      = unicode.IsLower
	isUpper  func(rune) bool		      = unicode.IsUpper
	uint16g  func([]byte) uint16                  = binary.BigEndian.Uint16
	uint16p  func([]byte, uint16)		      = binary.BigEndian.PutUint16
)

var ops []string = []string{
	`AND NOT`,
        `NOT`, `!(`,
        `OR`, `||`, `|(`,
        `AND`, `&&`, `&(`,
}

/*
Version defines the ACI syntax version number implemented by
this package.
*/
const Version float32 = 3.0

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

        return rune(raw[idx+1]), idx+1
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

        return rune(raw[idx-1]), idx-1
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
	val = trimS(raw[1:len(raw)-1])

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

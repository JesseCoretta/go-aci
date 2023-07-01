package aci

/*
cond.go contains rule condition types and methods.
*/

import (
	"github.com/JesseCoretta/go-stackage"
)

/*
Condition is a type alias of stackage.Condition, and provides the "<kw><op><ex>"
expressive syntax representation by itself, or as a slice element of a Rule.

Instances of this type may be created manually using the Cond function, however
it may be easier to utilize the following methods extended via some of the types
defined in this package:

• Eq

• Ne

• Lt

• Le

• Gt

• Ge

Note that occasionally some exceptions apply to the above list of methods. For
instance, non-numerical values are likely ineligible for operator methods such
as Gt, Le, etc.
*/
type Condition stackage.Condition

/*
ComparisonOperator constants defined within the stackage package are aliased
within this package for convenience, without the need for user-invoked stackage
package import.
*/
const (
	Eq stackage.ComparisonOperator = stackage.Eq	// 0x1, "Equal To"
	Ne stackage.ComparisonOperator = stackage.Ne	// 0x2, "Not Equal to"	// USE WITH CAUTION
	Lt stackage.ComparisonOperator = stackage.Lt	// 0x3, "Less Than"
	Le stackage.ComparisonOperator = stackage.Le	// 0x4, "Less Than Or Equal"
	Gt stackage.ComparisonOperator = stackage.Gt	// 0x5, "Greater Than"
	Ge stackage.ComparisonOperator = stackage.Ge	// 0x6, "Greater Than Or Equal"
)

/*
idOp attempts to identify an operator based on string input.
*/
func idOp(raw string) (op stackage.Operator, ok bool) {
	// try compops first
	for i := 0x1; i < 0x6; i++ {
		if raw == stackage.ComparisonOperator(i).String() {
			op = stackage.ComparisonOperator(i)
			ok = true
			return
		}
	}

	// TODO - add LDAP Search Filter operators
	// for fallback ...

	return
}

/*
Kind returns `target` or `bind` to identify the kind of *Condition.
*/
func (r Condition) Kind() string {
	return stackage.Condition(r).ID()
}

/*
Cond wraps go-stackage's Cond function. Condition instances are easier to create
using Eq, Ne, Gt, Ge, Lt and Le methods extended by the various types defined in
this package.

This particular function is really only needed by the user for special use-cases
in which Condition instances need to be manually created from scratch (rare).

A condition is represented through the following abstract structure:

 <keyword> <operator> <expression value>

An example:

 userdn = "ldap:///uid=jesse,ou=People,dc=example,dc=com"

As such, three (3) parameters are required:

• The kw input type describes an interface-encompassed Keyword type, or a string

• The ex input type is the expression (assertion) value to be declared within a Condition

• The op input type is the stackage.Operator interface definition, which allows the use of
custom operators introduced by the user, if needed (such as ~= for "approximate", et al).
However, it is unlikely any operators would be needed beyond the aforementioned go-stackage
ComparisonOperator constants.
*/
func Cond(kw, ex any, op stackage.Operator) Condition {
	return Condition(stackage.Cond(kw, op, ex))
}

/*
Valid wraps go-stackage's Condition.Valid method.
*/
func (r Condition) Valid() error {
	return stackage.Condition(r).Valid()
}

/*
Category wraps go-stackage's Condition.Category method.
*/
func (r Condition) Category() string {
	return stackage.Condition(r).Category()
}

/*
setID wraps go-stackage's Condition.SetID method.
*/
func (r Condition) setID(id string) Condition {
        stackage.Condition(r).SetID(id)
        return r
}

/*
setCategory wraps go-stackage's Condition.SetCategory method.
*/
func (r Condition) setCategory(cat string) Condition {
	stackage.Condition(r).SetCategory(cat)
	return r
}

/*
IsZero wraps go-stackage's Condition.IsZero method.
*/
func (r Condition) IsZero() bool {
	return stackage.Condition(r).IsZero()
}

/*
Encap wraps go-stackage's Condition.Encap method.
*/
func (r Condition) Encap(x ...any) Condition {
	stackage.Condition(r).Encap(x...)
	return r
}

/*
Paren wraps go-stackage's Condition.Paren method.
*/
func (r Condition) Paren(x ...bool) Condition {
	stackage.Condition(r).Paren(x...)
	return r
}

/*
NoPadding wraps go-stackage's Condition.NoPadding method.
*/
func (r Condition) NoPadding(x ...bool) Condition {
	stackage.Condition(r).NoPadding(x...)
	return r
}

/*
String wraps go-stackage's Condition.String method.
*/
func (r Condition) String() string {
	return stackage.Condition(r).String()
}

/*
Keyword wraps go-stackage's Condition.Keyword method.
*/
func (r Condition) Keyword() string {
	return stackage.Condition(r).Keyword()
}

/*
Operator wraps go-stackage's Condition.Operator method.
*/
func (r Condition) Operator() stackage.Operator {
	return stackage.Condition(r).Operator()
}

/*
Value wraps go-stackage's Condition.Value method.
*/
func (r Condition) Value() any {
	return stackage.Condition(r).Value()
}

/*
Parse shall marshal the provided raw input string into the
receiver. An error is returned should the process fail for
some reason.
*/
func (r Condition) Parse(raw string) error {
	_, err := r.parse(raw)
	return err
}

/*
parse is a private method called during the parsing of a
string value believed to be a conditional expression of
some form. An error is returned if the receiver was not
properly populated, or if the raw input value does not
conform to the known structure of a Condition.

Note that if a non-zero receiver (i.e.: the receiver 
already contains values that express a conditional
statement), those contents shall be wiped out if the
parsing process should succeed. No changes are made
if errors are encountered at any point.
*/
func (r *Condition) parse(raw string) (rest string, err error) {
	if len(raw) < 4 {
		err = errorf("%T parsing failed: raw conditional expression is below the minimum required length", r)
		return
	}

	// remove parenthetical encapsulation and
	// remove leading/trailing WHSP afterwards.
	if raw, _ = trimLRParen(raw); len(raw) == 0 {
		err = errorf("%T parsing failed: invalid or zero length raw value", r)
		return
	}

	// cidx defines the current index integer
	// value during each of the three (3)
	// stages of Condition parsing.
	var cidx int

	// Scan raw value rune by rune using the
	// idxf alias for strings.IndexFunc with
	// kwIdxFunc passed as the closure value.
	// We are looking for a keyword, which we
	// know should always be lower alphas.
	//
	// A return value of true during iteration
	// indicates such a char was found, false
	// otherwise, which breaks the loop. 
	var kw string
	for i := 0; i < len(raw); i++ {
		// != -1 true means keep going, we're
		// still reading keyword.
		if idx := idxf(string(raw[i]), kwIdxFunc); idx != -1 {
			kw += string(raw[i])
			continue
		}

		cidx = i
		break
	}

	// Bail out if no kw found
	if len(kw) == 0 {
		err = errorf("%T parsing failed: no keyword detected", r)
		return
	}

	// Attempt to identify what kind of keyword
	// we found by iterating each of the three
	// (3) keyword maps found in kw.go looking
	// for a stringer match.
	k, ok := idKW(kw)
	if !ok {
		err = errorf("%T parsing failed: unidentifiable or invalid keyword '%s'", r, kw)
		return
	}

        // Scan raw value rune by rune using the
        // idxf alias for strings.IndexFunc with
        // opIdxFunc passed as the closure value.
        // We are looking for an operator, which
        // can contain 
        //
        // A return value of true during iteration
        // indicates such a char was found, false
        // otherwise, which breaks the loop.
        var op string
        for i := cidx+1; i < len(raw[cidx+1:]); i++ {
		// false means keep going, we're
		// still reading operator.
                if idx := idxf(string(raw[i]), opIdxFunc); idx == -1 {
                        op += string(raw[i])
                        continue
                }

		cidx = i
                break
        }

	// Bail out if no op found
	o, opOK := idOp(op)
	if !opOK {
		err = errorf("%T parsing failed: zero length or unidentifiable operator '%s'", r, op)
		return
	}

        // Scan the remaining raw value rune by rune
	// until double-quote, right parenthesis or
	// EOL.
        var (
		ex any
		exstr string
		quoted bool
		ur Rule
	)

	raw = trimS(raw[cidx+1:])
        for i := 0; i < len(raw); i++ {
		var end bool
		switch char := rune(raw[i]); char {
		case lpar:
			if !quoted {
				if err = ur.parse(raw[i:]); err != nil {
					break
				}
				ex = ur
				end = true
			} else {
				exstr += string(raw[i])
			}
		case dqot:
			if i != 0 {
				end = true
				break
			}
			quoted = !quoted
		case whsp:
			if i != 0 {
				end = true
				break
			}
		default:
			if i == len(raw)-1 {
				end = true
				break
			}
			exstr += string(raw[i])
		}

		if end {
			cidx = i
			break
		} else if err != nil {
			break
		}
        }

	if err == nil {
		if cidx+1 < len(raw) {
			rest = raw[cidx+1:]
		}
		if len(exstr) > 0 {
			ex = exstr
		}

		*r = Cond(k, ex, o).
			Encap(`"`).
			Paren().
			setID(k.Kind()).
			setCategory(k.String())
	}

	return
}

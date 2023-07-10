package aci

/*
parse.go contains functions related to the interpolation of text based ACIs
into proper Instruction instances.
*/

/*
char constants used throughout the parsing process.
*/
const (
	ntbs rune = rune(0)   // NULL
	htab rune = rune(9)   // \t
	newl rune = rune(10)  // \n
	whsp rune = rune(32)  // \s
	excl rune = rune(33)  // !
	dqot rune = rune(34)  // "
	nsgn rune = rune(35)  // #
	dsgn rune = rune(36)  // $
	ampr rune = rune(38)  // &
	lpar rune = rune(40)  // (
	rpar rune = rune(41)  // )
	star rune = rune(42)  // *
	cmma rune = rune(44)  // ,
	dash rune = rune(45)  // -
	fstp rune = rune(46)  // .
	slds rune = rune(47)  // /
	coln rune = rune(58)  // ;
	semi rune = rune(59)  // ;
	lthn rune = rune(60)  // <
	eqls rune = rune(61)  // =
	gthn rune = rune(62)  // >
	lbrk rune = rune(91)  // [
	rbrk rune = rune(93)  // ]
	pipe rune = rune(124) // |
)

func parseInstruction(raw string) (aci Instruction, err error) {
	if len(raw) == 0 {
		err = errorf("%T string definition is zero", aci)
		return
	}

	// for parsing ease, replace `) (` with `)(`
	def := repAll(condenseWHSP(raw), `) (`, `)(`)

	var (
		vers  string = sprintf("%s;", version())
		label string
		ct    int
		x     int = idxs(def,vers)+1
		pbr,t Rule
	)

	// parse targets into instance of Rule
	if t, ct, err = parseInstructionTargets(def[:x]); ct == -1 {
		// no tgts (not fatal)
	} else if err != nil {
		return
	}

	// Obtain the label
	if label, ct = parseInstructionName(def[x:]); ct == -1 {
		err = errorf("%T parsing failed: no name (acl) found within ACI definition; a name is always required", aci)
		return
	}

	// Assign the label to the Instruction
	// return instance.
	aci.Set(label, t)
	offset := ct+x+len(label)+1
	if offset > len(def) {
		err = errorf("%T parsing failed: unexpected end of input definition", aci)
		return
	}

	// parse one (1) or more PB instances
	if pbr, err = parsePermissionBindRules(trimS(def[offset:])); err != nil {
		return
	}
	aci.Set(pbr)

	return
}

/*
parsePermissionBindRules reads the ACI definition string value and
extracts one (1) or more PermissionBindRule instance values.
*/
func parsePermissionBindRules(def string) (pbr Rule, err error) {
	var (
		P Permission	// allow/deny(...)
		p *permission	// embedding for P
		B Rule		// Bind Rule(s)
		bidx int	// bind "rest" index for continuous processing of expressive statements
	)

	// First parse the privilege keyword(s)
	// along with the disposition. This is
	// always required in any given PBR.
	if p, bidx, err = p.parse(def); err != nil {
		return
	}
	P = Permission{p}

	if B, _, err = parseRule(def[bidx+1:]); err != nil {
		return
	}

	pbr = pbrule()
	pB := PB(P,B)
	pbr.Push(pB)

	return
}

/*
parseInstructionTargets reads the ACI definition string value
and extracts up to nine (9) Target Rule definitions. A Rule
containing the extracted values, along with a count of the
extracted values, is returned upon completion. A count of minus
one (-1) means there was an error. A count of zero (0) means
no Target Rule definitions were found (which is not an error
necessarily).

Use of Target Rule conditions is not required (not all ACIs use
them). However if at least one (1) such rule if detected, and it
is perceived to be invalid in some way, the entire parsing process
will fail with an error.
*/
func parseInstructionTargets(def string) (t Rule, c int, err error) {
	var targs []string
	t = T() // prepare a Target Rule container

	// Begin looping and don't stop until
	// done, or an error was encountered.
	for {
		// Each iteration, when successful, obtains one (1)
		// Target Rule condition statement (<kw><op><ex>).
		// We want to count the number of iterations thus
		// far, as one can only have so many such rules
		// if they have any at all ...
		c++

		// Use neighboring R/L paren sequence as an
		// indication of one TR ending, and another
		// beginning ...
		if idx := idxs(def, `)(`); idx != -1 {

			var (
				//cnd     Condition
				invalid,
				kwcReady,
				kwcValid bool
				kwc     string
			)

			// First, extract and verify the Target Keyword
			// by iterating each character rune individually.
			for i := 0; i < len(def); i++ {
				switch char := rune(def[i]); char {
				case lpar:
					// LEFT PARENTHESIS
					// We should be at the very beginning
					// of a target rule. paren not needed
					continue
				case whsp, rpar, eqls:
					// WHITESPACE, EQUALS or RIGHT PARENTHESIS 
					// Declare "done" for this current Target
					// Rule keyword
					kwcReady = true
				default:
					// ANY RUNE
					// TargetKeyword instances are strictly
					// alphabetical, so we only accept alpha
					// chars (and only those that are lower
					// case normalized) ...
					if isLower(char) {
						kwc += string(char)
					} else {
						// flip out if we found anything
						// else not handled in this case
						// switching routine ...
						kwc += string(char)	// include the offending rune for err info.
						invalid = true
						break
					}
				}

				// We're either done processing what we believe to
				// be a Target Rule keyword, OR we encountered an
				// error in said keyword ...
				if kwcReady && !invalid {
					// See if we can match the candiate kw
					// to an entry in our global tkwMap var.
					// If not, declare our error.
					if kwcValid = matchTKW(kwc).String() == kwc; !kwcValid {
						invalid = !kwcValid
						break
					}

					// Make sure
					kwcReady = false
					kwc = ``
				}

				// Whether bogus or zero length, we're boned. Prepare
				// an error instance for the invalid keyword ...
				if invalid {
					err = errorf("Target Rule candidate found with bogus keyword '%s'", kwc)
					break
				}
			}

			if err != nil {
				break
			}

			targs = append(targs, def[:idx+1])
			def = def[idx+1:] // chop left components starting at idx for each iteration
		} else if idx == -1 {
			break // done processing
		} else if c > 9 {
			err = errorf("Target Rule count exceeds maximum possible size of nine (9)")
			// too many targets!
			break
		}
	}

	if err == nil {
		// Loop through our complete condition
		// expression statements, and parse each
		// into a new instance of Condition, which
		// is then pushed into our Target Rule (t)
		// container created earlier.
		for _, targ := range targs {
			var Cx Condition
			if Cx, _, err = parseCondition(targ); err != nil {
				break
			}
			t.Push(Cx)
		}
	}

	return
}

/*
parseInstructionName extracts the name (acl) label from the
definition and returns it.
*/
func parseInstructionName(def string) (n string, c int) {
	n = badACI
	s := idxs(def, `acl "`)
	if s == -1 {
		return
	}

	offset := s+4 // 4 = index pt for `acl "`
	if c = idxr(def[offset:], ';'); c != -1 {
		n = def[offset:offset+c]
		if len(n) < 2 {
			// value was less than two (2) bytes
			c = -1
		} else if n[1] != '"' && n[len(n)-1] != '"' {
			// value was unquoted?!
			c = -1
		} else if n = n[1:len(n)-1]; len(n) == 0 {
			// seems OK, trim quotes in
			// preparation for return.
			c = -1
		}
	}

	return
}


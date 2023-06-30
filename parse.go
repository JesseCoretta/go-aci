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

func parseInstruction(def string) (aci Instruction, err error) {
	var (
		//sec int // Sections: 0=T[s],1=Vers/Lbl,2=PB[s]
		idx int

		char,
		next,
		last rune

		tL,
		tN,
		oP,
		oD,
		oQ bool

		//kwc,		// keyword candidate
		//opc,		// operator candidate
		//exc string	// expression candidate
		val string
		raw string = condenseWHSP(removeNewlines(def))

		cnd [3]string	// Condition *components* (candidate values)
	)

	for i := 0; i < len(raw); i++ {
		next, idx = nextRune(raw, i)
		tL = idx != -1

		last, idx = lastRune(raw, i)
		tN = idx != -1

		if tL {
			printf("Last %T is '%c'\n", last, last)
		}

		if tN {
			printf("Next %T is '%c'\n", next, next)
		}

		char = rune(raw[i])
		printf("Current %T is '%c'\n", char, char)

		printf("Double-quotation: %t\n", oQ)
		printf("Active parenthesis: %t\n", oP)

		switch char {
		case whsp:
			if len(cnd[2]) == 0 && len(val) > 0 && !oD {
				cnd[2] = val
			} else if len(cnd[1]) == 0 && len(val) > 0 && !oD {
				cnd[1] = val
			} else if oD {
				cnd[2] = val
			}
			val = ``
		case excl, eqls, gthn, lthn, coln, slds:
			if oD && oQ {
				val += string(char)
			} else if oQ {
				val += string(char)
			} else {
				printf("Not sure what to do with '%c'\n", char)
			}
		case dqot:
			oQ = !oQ
			oD = oQ
		case lpar:
			oP = true
		case rpar:
			oP = false
			printf("CND now: '%v' - '%v' - '%v'\n", cnd[0],cnd[1],cnd[2])
		default:
			if oD {
				val += string(char)
			} else if oP && oD {
				val += string(char)
			} else if !oP {
				val += string(char)
			}
		}
	}

	return
}

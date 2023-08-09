package aci

/*
levels.go contains methods pertaining to types based upon concepts of "inheritance"
within the ACIv3 standard.
*/

// Maps for resolving level instances
var (
	levelMap     = make(map[int]Level, 0)
	levelNumbers = make(map[string]Level, 0)
)

/*
Level uint16 constants are left-shifted into an instance of Levels
to define a range of vertical (depth) rule statements.
*/
const (
	noLvl  Level = 0         //   0 - <no levels>
	Level0 Level = 1 << iota //   1 - base  (0) (current Object)
	Level1                   //   2 - one   (1) level below baseObject
	Level2                   //   4 - two   (2) levels below baseObject
	Level3                   //   8 - three (3) levels below baseObject
	Level4                   //  16 - four  (4) levels below baseObject
	Level5                   //  32 - five  (5) levels below baseObject
	Level6                   //  64 - six   (6) levels below baseObject
	Level7                   // 128 - seven (7) levels below baseObject
	Level8                   // 256 - eight (8) levels below baseObject
	Level9                   // 512 - nine  (9) levels below baseObject

	AllLevels Level = Level(1023) // ALL levels; zero (0) through nine (9)
)

/*
badInhErr returns an error describing the appropriate syntax and displaying the offending value.
*/
func badInhErr(bad string) error {
	return errorf("Bad Inheritance value '%s'; must conform to 'parent[0-4+].<at>#<bt_or_av>'", bad)
}

/*
Inheritance describes an inherited Bind Rule syntax, allowing access
control over child entry enumeration below the specified parent.
*/
type Inheritance struct {
	*inheritance
}

/*
inheritance is the private embedded (POINTER!) type found within
instances of Inheritance. It contains a Level bit container and
an AttributeBindTypeOrValue instance.
*/
type inheritance struct {
	Level
	AttributeBindTypeOrValue
}

/*
Inherit creates a new instance of Inheritance bearing the provided
AttributeBindTypeOrValue instance, as well as zero (0) or more Level
instances for shifting.
*/
func Inherit(x AttributeBindTypeOrValue, lvl ...any) Inheritance {
	return Inheritance{newInheritance(x, lvl...)}
}

/*
newInheritance initializes and sets a new instance of *inheritance,
which is embedded within a new instance of Inheritance. This function
is called by Inherit.
*/
func newInheritance(x AttributeBindTypeOrValue, lvl ...any) (i *inheritance) {
	i = new(inheritance)
	i.shift(lvl...)
	i.AttributeBindTypeOrValue = x

	return
}

/*
Level describes a discrete numerical abstract of a subordinate level.
Level describes one (1) or more additive values of Level that collectively
express the subordinate level condition within a given Bind Rule that involves
inheritance.

Valid levels are level zero (0) through level nine (9), though this may
vary across implementations.
*/
type Level uint16

/*
IsZero returns a boolean value indicative of whether the receiver instance
is nil, or unset.
*/
func (r Inheritance) IsZero() bool {
	return r.inheritance.isZero()
}

/*
Eq initializes and returns a new Condition instance configured to express the
evaluation of the receiver value as Equal-To a `userattr`.
*/
func (r Inheritance) Eq() Condition {
	if r.IsZero() {
		return Condition{}
	}
	kw := r.inheritance.AttributeBindTypeOrValue.BindKeyword
	return Cond(kw, r, Eq).
		Encap(`"`).
		setID(`bind`).
		NoPadding(!ConditionPadding).
		setCategory(kw.String())
}

/*
Ne initializes and returns a new Condition instance configured to express the
evaluation of the receiver value as Not-Equal-To a `userattr`.
*/
func (r Inheritance) Ne() Condition {
	if r.IsZero() {
		return Condition{}
	}
	kw := r.inheritance.AttributeBindTypeOrValue.BindKeyword
	return Cond(kw, r, Ne).
		Encap(`"`).
		setID(`bind`).
		NoPadding(!ConditionPadding).
		setCategory(kw.String())
}

/*
isZero is a private method called by Inheritance.IsZero.
*/
func (r *inheritance) isZero() bool {
	if r == nil {
		return true
	}
	return r.AttributeBindTypeOrValue.IsZero()
}

/*
parseInheritance is a private function that reads the input string (inh)
and attempts to marshal its contents into an instance of Inheritance (I),
which is returned alongside an error (err).

This function is called during the bind rule parsing phase if and when
an inheritance-related userattr/groupattr rule is encountered.
*/
func parseInheritance(inh string) (I Inheritance, err error) {
	// Bail out immediately if the prefix is
	// non conformant.
	if !hasPfx(lc(inh), `parent[`) {
		err = badInhErr(inh)
		return
	}

	// chop off the 'parent[' prefix; we don't need
	// to preserve it following the presence check.
	raw := inh[7:]

	// Grab the sequence of level identifiers up to
	// and NOT including the right (closing) bracket.
	// The integer index (idx) marks the boundary of
	// the identifier sequence.
	idx := idxr(raw, ']')
	if idx == -1 {
		// non conformant!
		err = badInhErr(inh)
		return
	}

	// Initialize our return instance, as we're about
	// to begin storing things in it.
	I = Inheritance{new(inheritance)}

	// Iterate the split sequence of level identifiers.
	// Also, obliterate any ASCII #32 (SPACE) chars
	// (e.g.: ', ' -> ',').
	for _, r := range split(repAll(raw[:idx], ` `, ``), `,`) {
		I.inheritance.shift(r) // left shift
	}

	// Bail if nothing was found (do not fall
	// back to default when parsing).
	if I.inheritance.Level == noLvl {
		// bogus or unsupported identifiers?
		err = errorf("No level identifiers parsed; aborting")
		return
	}

	// Call our AttributeBindTypeOrValue parser
	// and marshal a new instance to finish up.
	// At this phase, we begin value parsing
	// one (1) character after the identifier
	// boundary (see above).
	var abv AttributeBindTypeOrValue

	if abv, err = parseATBTV(raw[idx+2:]); err != nil {
		// non conformant ATBTV
		return
	}
	I.inheritance.AttributeBindTypeOrValue = abv
	return
}

/*
String is a stringer method that returns the string name value for receiver instance of Inheritance.

The return value(s) are enclosed within square-brackets, comma-delimited and prefixed with "parent".
*/
func (r Inheritance) String() string {
	if r.IsZero() {
		return `parent[0]`
	}

	var levels []string
	if r.inheritance.Level == noLvl {
		// No levels? default to level 0 (baseObject)
		levels = append(levels, Level0.String())
	} else {
		for i := 1; i < 12; i++ {
			shift := Level(1 << i)
			if r.Positive(shift) {
				levels = append(levels, shift.String())
			}
		}
	}

	return sprintf("parent[%s].%s", join(levels, `,`),
		r.inheritance.AttributeBindTypeOrValue)
}

/*
String is a stringer method that returns a single string name value for receiver instance of Level.
*/
func (r Level) String() (lvl string) {
	for k, v := range levelNumbers {
		if r == v {
			lvl = k
			return
		}
	}

	return
}

/*
Shift shifts the receiver instance of Levels to include Level x, if not already present.
*/
func (r *Inheritance) Shift(x ...any) *Inheritance {
	r.inheritance.shift(x...)
	return r
}

func (r *inheritance) shift(x ...any) {
	for i := 0; i < len(x); i++ {
		var lvl Level
		switch tv := x[i].(type) {
		case Level:
			if tv == noLvl {
				continue
			}
			lvl = tv
		case int:
			if lvl = assertIntInheritance(tv); lvl == noLvl {
				continue
			}
		case string:
			if lvl = assertStrInheritance(tv); lvl == noLvl {
				continue
			}
		default:
			continue
		}

		(*r).Level |= lvl
	}
}

/*
assertStrInheritance returns the appropriate Level instance
logically associated with the string value (x) input by the
user. Valid levels are zero (0) through four (4), else noLvl
is returned.
*/
func assertStrInheritance(x string) (lvl Level) {
	lvl = noLvl
	for k, v := range levelNumbers {
		if x == k {
			lvl = v
			break
		}
	}

	return
}

/*
assertIntInheritance returns the appropriate Level instance
logically associated with the integer value (x) input by the
user. Valid levels are zero (0) through four (4), else noLvl
is returned.
*/
func assertIntInheritance(x int) (lvl Level) {
	lvl = noLvl
	if L, found := levelMap[x]; found {
		lvl = L
		return
	}

	return
}

/*
Positive returns a boolean value indicative of whether the receiver instance of Levels includes Level x.
*/
func (r Inheritance) Positive(x any) bool {
	return r.inheritance.positive(x)
}

func (r *inheritance) positive(x any) (posi bool) {
	if r.isZero() {
		return
	}

	var lvl Level
	switch tv := x.(type) {
	case Level:
		if tv == noLvl {
			return
		}
		lvl = tv
	case int:
		if lvl = assertIntInheritance(tv); lvl == noLvl {
			return
		}
	case string:
		if lvl = assertStrInheritance(tv); lvl == noLvl {
			return
		}
	default:
		return
	}

	posi = (r.Level & lvl) > 0
	return
}

/*
Unshift right-shifts the receiver instance of Levels to remove Level x, if present.
*/
func (r *Inheritance) Unshift(x ...any) *Inheritance {
	r.inheritance.unshift(x...)
	return r
}

func (r *inheritance) unshift(x ...any) {
	for i := 0; i < len(x); i++ {
		var lvl Level
		switch tv := x[0].(type) {
		case Level:
			if tv == noLvl {
				continue
			}
			lvl = tv
		case int:
			if lvl = assertIntInheritance(tv); lvl == noLvl {
				continue
			}
		case string:
			if lvl = assertStrInheritance(tv); lvl == noLvl {
				continue
			}
		default:
			continue
		}

		r.Level = r.Level &^ lvl
	}

	return
}

func init() {
	levelMap = map[int]Level{
		0: Level0,
		1: Level1,
		2: Level2,
		3: Level3,
		4: Level4,
		5: Level5,
		6: Level6,
		7: Level7,
		8: Level8,
		9: Level9,
	}

	levelNumbers = map[string]Level{
		`0`: Level0,
		`1`: Level1,
		`2`: Level2,
		`3`: Level3,
		`4`: Level4,
		`5`: Level5,
		`6`: Level6,
		`7`: Level7,
		`8`: Level8,
		`9`: Level9,
	}
}

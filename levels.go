package aci

/*
levels.go contains methods pertaining to types based upon concepts of "inheritance"
within the ACIv3 standard.
*/

// Maps for resolving level instances
var (
	levelBitIter = bitSize(Levels(0)) - 4 // we don't use all of uint16, no sense iterating the whole thing
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

	AllLevels Level = Level(2046) // ALL levels; zero (0) through nine (9)
)

/*
Inheritance describes an inherited BindRule syntax, allowing access
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
	Levels
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
Level describes a discrete numerical abstract of a subordinate level. Level
describes any single Level definition. Level constants are intended for "storage"
within an instance of Levels -- the compound counterpart of this type.

Valid Level constants are level zero (0) through level nine (9), though this
will vary across implementations.
*/
type Level uint16

/*
Levels is the compound (bitshifted) counterpart to the scalar Level type.

Levels can express all combinations Level values; for example, if a Levels
instance has an underlying value of three (3), this describes the presence
of the individual Level0 (1) and Level1 (2) definitions within said value
(i.e.: 1+2=3).
*/
type Levels uint16

/*
IsZero returns a Boolean value indicative of whether the receiver instance
is nil, or unset.
*/
func (r Inheritance) IsZero() bool {
	return r.inheritance.isZero()
}

/*
Valid returns an error indicative of whether the receiver is in an aberrant state.
*/
func (r Inheritance) Valid() (err error) {
	if r.IsZero() {
		return nilInstanceErr(r)
	}

	err = nilInstanceErr(r.inheritance.AttributeBindTypeOrValue)
	if !r.inheritance.AttributeBindTypeOrValue.IsZero() {
		err = nil
	}

	return
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r Inheritance) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
BRM returns an instance of BindRuleMethods.

Each of the return instance's key values represent a single instance of
the ComparisonOperator type that is allowed for use in the creation of
BindRule instances which bear the receiver instance as an expression
value. The value for each key is the actual BindRuleMethod instance for
OPTIONAL use in the creation of a BindRule instance.

This is merely a convenient alternative to maintaining knowledge of which
ComparisonOperator instances apply to which types. Instances of this type
are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not
yet been initialized, the execution of ANY of the return instance's value
methods will return bogus BindRule instances. While this is useful in unit
testing, the end user must only execute this method IF and WHEN the receiver
has been properly populated and prepared for such activity.
*/
func (r Inheritance) BRM() BindRuleMethods {
	return newBindRuleMethods(bindRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
Eq initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Equal-To the `userattr` or `groupattr`
Bind keyword context.
*/
func (r Inheritance) Eq() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(r.inheritance.AttributeBindTypeOrValue.BindKeyword, Eq, r)
}

/*
Ne initializes and returns a new BindRule instance configured to express the
evaluation of the receiver value as Not-Equal-To the `userattr` or `groupattr`
Bind keyword context.
*/
func (r Inheritance) Ne() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(r.inheritance.AttributeBindTypeOrValue.BindKeyword, Ne, r)
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

	// make sure the dot delimiter
	// comes immediately after the
	// closing square bracket.
	if raw[idx+1] != '.' {
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
	if I.inheritance.Levels == Levels(noLvl) {
		// bogus or unsupported identifiers?
		err = levelsNotFoundErr()
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
Len returns the abstract integer length of the receiver, quantifying
the number of Level instances currently being expressed. For example,
if the receiver instance has its Level1 and Level5 bits enabled, this
would represent an abstract length of two (2).
*/
func (r Inheritance) Len() int {
	var D int
	for i := 0; i < bitSize(noLvl); i++ {
		if d := Level(1 << i); r.Positive(d) {
			D++
		}
	}

	return D
}

/*
Levels returns the compound Levels instance within the receiver.
*/
func (r Inheritance) Levels() Levels {
	if r.IsZero() {
		return Levels(noLvl)
	}
	return r.inheritance.Levels
}

/*
Keyword returns the Keyword associated with the receiver instance. In
the context of this type instance, the Keyword returned will be either
BindUAT or BindGAT.
*/
func (r Inheritance) Keyword() (kw Keyword) {
	if err := r.Valid(); err != nil {
		return nil
	}

	k := r.inheritance.AttributeBindTypeOrValue.BindKeyword
	switch k {
	case BindGAT, BindUAT:
		kw = k
	}

	return
}

/*
String is a stringer method that returns the string name value for
receiver instance of Inheritance.

The return value(s) are enclosed within square-brackets, followed
by comma delimitation and are prefixed with "parent" before being
returned.
*/
func (r Inheritance) String() string {
	if err := r.Valid(); err != nil {
		return badInheritance
	}

	// string representation of Levels sequence
	lvls := r.inheritance.Levels.String()

	return sprintf("parent[%s].%s", lvls,
		r.inheritance.AttributeBindTypeOrValue)
}

/*
Len returns the abstract integer length of the receiver, quantifying
the number of Level instances currently being expressed. For example,
if the receiver instance has its Level4 and Level7 bits enabled, this
would represent an abstract length of two (2).
*/
func (r Levels) Len() (l int) {
	for i := 0; i < levelBitIter; i++ {
		shift := Level(1 << i)
		if r.Positive(shift) {
			l++
		}
	}

	return
}

/*
String is a string method that returns the string
representation of the receiver instance.
*/
func (r Levels) String() string {
	var levels []string
	if r == Levels(noLvl) {
		// No levels? default to level 0 (baseObject)
		levels = append(levels, Level0.String())
	} else {
		for i := 0; i < levelBitIter; i++ {
			shift := Level(1 << i)
			if r.Positive(shift) {
				levels = append(levels, shift.String())
			}
		}
	}

	return join(levels, `,`)
}

/*
String is a stringer method that returns a single string name value for receiver instance of Level.
*/
func (r Level) String() (lvl string) {
	for k, v := range levelNumbers {
		if r == v {
			lvl = k
			break
		}
	}

	return
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r Level) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Shift wraps Levels.Shift via the underlying Levels value
found within the receiver instance.
*/
func (r Inheritance) Shift(x ...any) Inheritance {
	if r.IsZero() {
		r = Inheritance{new(inheritance)}
	}

	r.inheritance.Levels.shift(x...)
	return r
}

/*
Shift shifts the receiver instance of Levels to include Level
x, if not already present.
*/
func (r *Levels) Shift(x ...any) *Levels {
	r.shift(x...)
	return r
}

/*
shift is a private method called by the Shift method.
*/
func (r *Levels) shift(x ...any) {
	for i := 0; i < len(x); i++ {
		var lvl Level
		switch tv := x[i].(type) {
		case Level:
			if tv != noLvl {
				lvl = tv
			}
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

		(*r) |= Levels(lvl)
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
Positive returns a Boolean value indicative of whether
the receiver instance of Levels includes Level x.
*/
func (r Inheritance) Positive(x any) bool {
	if r.IsZero() {
		r = Inheritance{new(inheritance)}
	}
	return r.inheritance.Levels.positive(x)
}

/*
Positive returns a Boolean value indicative of whether
the receiver has the appropriate bits enabled to include
abstract value x, which describes a Level definition.
*/
func (r Levels) Positive(x any) bool {
	return r.positive(x)
}

/*
IsZero returns a Boolean value indicative of whether the
receiver is in an aberrant state.
*/
func (r Levels) IsZero() bool {
	return int(r) == 0
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r Levels) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Valid returns an error instance that describes the undesirable
state of the receiver, if applicable. A nil error is returned
otherwise.
*/
func (r Levels) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
	}

	// TODO: additional checks?
	return
}

/*
positive is a private method executed by the Positive method.
*/
func (r Levels) positive(x any) (posi bool) {
	if r.IsZero() {
		return
	}

	var lvl Level
	switch tv := x.(type) {
	case Level:
		if tv != noLvl {
			lvl = tv
		}
	case int:
		lvl = assertIntInheritance(tv)
	case string:
		lvl = assertStrInheritance(tv)
	}

	posi = (r & Levels(lvl)) > 0
	return
}

/*
Unshift wraps Levels.Unshift via the underlying Levels value
found within the receiver instance.
*/
func (r Inheritance) Unshift(x ...any) Inheritance {
	if r.IsZero() {
		r = Inheritance{new(inheritance)}
	}

	r.inheritance.Levels.unshift(x...)
	return r
}

/*
Unshift right-shifts the receiver instance of Levels to
remove Level x, if present.
*/
func (r *Levels) Unshift(x ...any) *Levels {
	r.unshift(x...)
	return r
}

/*
unshift is a private method called by the Unshift method.
*/
func (r *Levels) unshift(x ...any) {
	for i := 0; i < len(x); i++ {
		var lvl Level
		switch tv := x[0].(type) {
		case Level:
			if tv != noLvl {
				lvl = tv
			}
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

		(*r) = (*r) &^ Levels(lvl)
	}

	return
}

const badInheritance = `<invalid_inheritance>`

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

package aci

/*
levels.go contains methods pertaining to types based upon concepts of "inheritance"
within the ACIv3 standard.
*/

/*
Level uint8 constants are left-shifted into an instance of Levels
to define a range of vertical (depth) rule statements.
*/
const (
	noLvl  Level = 0         //  0 - <no levels>
	Level0 Level = 1 << iota //  1 - baseObject
	Level1                   //  2 - one (1) level below baseObject
	Level2                   //  4 - two (2) levels below baseObject
	Level3                   //  8 - three (3) levels below baseObject
	Level4                   // 16 - four (4) levels below baseObject

	AllLevels Level = Level(31)
)

/*
Inheritance describes an inherited Bind Rule syntax, allowing access
control over child entry enumeration below the specified parent.
*/
type Inheritance struct {
	*inheritance
}

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
	for j := 0 ; j < len(lvl); j++ {
		i.shift(lvl[j])
	}
	i.AttributeBindTypeOrValue = x

	return
}

/*
Level describes a discrete numerical abstract of a subordinate level.
Level describes one (1) or more additive values of Level that collectively
express the subordinate level condition within a given Bind Rule that involves
inheritance.
*/
type Level uint8

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
        return Cond(kw, r, Eq).Encap(`"`).setID(`bind`).setCategory(kw.String())
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
        return Cond(kw, r, Ne).Encap(`"`).setID(`bind`).setCategory(kw.String())
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
String is a stringer method that returns the string name value for receiver instance of Inheritance.

The return value(s) are enclosed within square-brackets, comma-delimited and prefixed with "parent".
*/
func (r Inheritance) String() string {
	if r.IsZero() {
		return `parent[0]`
	}

	var levels []string
	if r.inheritance.Level == noLvl {
		return `parent[0]`
	} else {
		for i := 1; i < 6; i++ {
			shift := Level(1 << i)
			if r.Positive(shift) {
				levels = append(levels, shift.String())
			}
		}
	}

	return sprintf("parent[%s].%s",
		join(levels, `,`),
		r.inheritance.AttributeBindTypeOrValue)
}

/*
String is a stringer method that returns a single string name value for receiver instance of Level.
*/
func (r Level) String() (lvl string) {
	switch r {
	case Level0:
		lvl = `0`
	case Level1:
		lvl = `1`
	case Level2:
		lvl = `2`
	case Level3:
		lvl = `3`
	case Level4:
		lvl = `4`
	}

	return
}

/*
Shift shifts the receiver instance of Levels to include Level x, if not already present.
*/
func (r *Inheritance) Shift(x any) *Inheritance {
	r.inheritance.shift(x)
	return r
}

func (r *inheritance) shift(x any) {
	var lvl Level
	switch tv := x.(type) {
	case Level:
		lvl = tv
	case int:
		switch tv {
		case 0:
			lvl = Level0
		case 1:
			lvl = Level1
		case 2:
			lvl = Level2
		case 3:
			lvl = Level3
		case 4:
			lvl = Level4
		}
	default:
		return
	}

	(*r).Level |= lvl
}

/*
Positive returns a boolean value indicative of whether the receiver instance of Levels includes Level x.
*/
func (r Inheritance) Positive(x any) bool {
	return r.inheritance.positive(x)
}

func (r *inheritance) positive(x any) bool {
	if r.isZero() {
		return false
	}

	var lvl Level
	switch tv := x.(type) {
	case Level:
		lvl = tv
	case int:
		if !(0 <= tv && tv <= 4) {
			return false
		}
		lvl = Level(uint8(tv + 1))
	default:
		return false
	}

	return (r.Level & lvl) > 0
}

/*
Unshift right-shifts the receiver instance of Levels to remove Level x, if present.
*/
func (r *Inheritance) Unshift(x any) *Inheritance {
	r.inheritance.unshift(x)
	return r
}

func (r *inheritance) unshift(x any) {
	var lvl Level
	switch tv := x.(type) {
	case Level:
		lvl = tv
	case int:
		if !(0 <= tv && tv <= 4) {
			return
		}
		lvl = Level(uint8(tv + 1))
	default:
		return
	}

	r.Level = r.Level &^ lvl
	return
}

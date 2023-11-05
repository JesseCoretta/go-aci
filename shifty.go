package aci

/*
shifty.go is a bridge to the go-shifty bit-shifting package's types and methods.
*/

import (
	"github.com/JesseCoretta/go-shifty"
)

func newDoW() DayOfWeek {
	return DayOfWeek(shifty.New(shifty.Uint8))
}

func newLvls() *levels {
	l := levels(shifty.New(shifty.Uint16))
	return &l
}

func newRights() *rights {
	r := rights(shifty.New(shifty.Uint16))
	return &r
}

func (r DayOfWeek) cast() shifty.BitValue {
	return shifty.BitValue(r)
}

func (r levels) cast() shifty.BitValue {
	return shifty.BitValue(r)
}

func (r rights) cast() shifty.BitValue {
	return shifty.BitValue(r)
}

type (
	// [DayOfWeek] is a type alias of shifty.BitValue, and is used
	// to construct a dayofweek [BindRule].
	DayOfWeek shifty.BitValue // 8-bit

	// rights is a private type alias of shifty.BitValue, and is
	// used in the construction an instance of [Permission].
	rights shifty.BitValue // 16-bit

	// levels is a private type alias of shifty.BitValue, and is
	// used in the construction an inheritance-based userattr or
	// groupattr BindRule by embedding.
	levels shifty.BitValue // 16-bit

)

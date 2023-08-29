package aci

/*
rights.go contains methods pertaining to the use of discrete privilege identifiers.
*/

/*
Right constants are discrete left-shifted privilege aggregates that can be used in an additive (or subtractive) manner to form a complete Permission.
*/
const (
	ReadAccess      Right = 1 << iota // 1
	WriteAccess                       // 2
	AddAccess                         // 4
	DeleteAccess                      // 8
	SearchAccess                      // 16
	CompareAccess                     // 32
	SelfWriteAccess                   // 64
	ProxyAccess                       // 128
	ImportAccess                      // 256
	ExportAccess                      // 512

	NoAccess  Right = 0
	AllAccess Right = 895 // DOES NOT INCLUDE "proxy"
)

const badPerm = `<invalid_permission>`

var rightsMap map[Right]string
var rightsNames map[string]Right

/*
Right contains the specific bit value of a single user privilege. Constants of this
type are intended for submission to the Permission.Shift, Permission.Unshift and
Permission.Positive methods.
*/
type Right uint16

/*
Permission defines a level of access bestowed (or withheld) by an ACI.
*/
type Permission struct {
	*permission
}

type permission struct {
	*bool
	*Right
}

/*
Allow returns a granting Permission instance bearing the provided
instances of Right.
*/
func Allow(x ...any) Permission {
	return Permission{newPermission(true, x...)}
}

/*
Deny returns a withholding Permission instance bearing the provided
instances of Right.
*/
func Deny(x ...any) Permission {
	return Permission{newPermission(false, x...)}
}

/*
newPermission returns a newly initialized instance of *permission
bearing the provided disposition and Right instance(s).
*/
func newPermission(disp bool, x ...any) (p *permission) {
	p = new(permission)
	p.bool = &disp
	p.Right = new(Right)
	p.shift(x...)
	return p
}

func (r *permission) shift(x ...any) {
	if r.isZero() {
		return
	}

	// iterate through the sequence of "anys"
	// and assert to a Right (or the abstraction
	// of a Right).
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case int:
			r.shiftIntRight(tv)
		case string:
			r.shiftStrRight(tv)
		case Right:
			r.shiftIntRight(int(tv))
		}
	}
}

func (r *permission) unshift(x ...any) {
	if r.isZero() {
		return
	}

	// iterate through the sequence of "anys"
	// and assert to a Right (or the abstraction
	// of a Right).
	for i := 0; i < len(x); i++ {
		switch tv := x[i].(type) {
		case int:
			r.unshiftIntRight(tv)
		case string:
			r.unshiftStrRight(tv)
		case Right:
			r.unshiftIntRight(int(tv))
		}
	}
}

func (r *permission) shiftIntRight(i int) {
	// avoid over/under flow
	if !(0 <= i && i <= int(^uint16(0))) {
		return
	}

	if _, matched := noneOrFullAccessRights(Right(i)); matched {
		(*r.Right) = Right(i)
		return
	}

	if rightsPowerOfTwo(i) {
		(*r.Right) |= Right(i)
	}
}

func (r *permission) unshiftIntRight(i int) {
	// avoid over/under flow
	if !(0 <= i && i <= int(^uint16(0))) {
		return
	}

	if R, matched := noneOrFullAccessRights(Right(i)); matched {
		(*r.Right) = (*r.Right) &^ R
		return
	}

	if rightsPowerOfTwo(i) {
		(*r.Right) = (*r.Right) &^ Right(i)
	}
}

func (r *permission) shiftStrRight(s string) {
	if priv, found := rightsNames[lc(s)]; found {
		(*r.Right) |= priv
	}
}

func (r *permission) unshiftStrRight(s string) {
	if priv, found := rightsNames[lc(s)]; found {
		(*r.Right) = (*r.Right) &^ priv
	}
}

func rightsPowerOfTwo(x int) bool {
	return isPowerOfTwo(x) && (0 <= x && x <= int(^uint16(0)))
}

func (r permission) positive(x any) bool {
	if r.isZero() {
		return false
	}
	switch tv := x.(type) {
	case int:
		return r.positiveIntRight(tv)

	case string:
		return r.positiveStrRight(tv)

	case Right:
		return r.positiveIntRight(int(tv))
	}

	// unsupported type?
	return false
}

func (r permission) positiveIntRight(i int) bool {
	// avoid over/under flow
	if !(0 <= i && i <= int(^uint16(0))) {
		return false
	}

	if _, matched := noneOrFullAccessRights(Right(i)); matched {
		if i == 0 {
			return int(*r.Right) == i
		}
		return ((*r.Right) & AllAccess) > 0
	}

	if rightsPowerOfTwo(i) {
		return ((*r.Right) & Right(i)) > 0
	}

	return false
}

func (r permission) positiveStrRight(s string) bool {
	if R, matched := noneOrFullAccessString(lc(s), (*r.Right)); matched {
		if R == NoAccess {
			return int(*r.Right) == 0
		}
		return ((*r.Right) & AllAccess) > 0
	}

	// Resolve the name of a Right into a Right.
	if priv, found := rightsNames[lc(s)]; found {
		return ((*r.Right) & priv) > 0
	}

	return false
}

func noneOrFullAccessRights(x Right) (Right, bool) {
	var matched bool
	if x == NoAccess {
		// NoAccess should stop the party
		// dead in its tracks, regardless
		// of any other iterated value.
		matched = true

	} else if x == AllAccess {
		// AllAccess should stop the party
		// dead in its tracks, regardless
		// of any other iterated value.
		matched = true
	}

	return x, matched
}

func noneOrFullAccessString(x string, r Right) (R Right, matched bool) {
	switch {
	case x == NoAccess.String():
		R = NoAccess
		matched = int(r) == 0 // 0
	case x == AllAccess.String():
		R = AllAccess
		matched = int(r) == 895 // *895
	}

	return
}

/*
String is a stringer method that returns a single string name value for receiver instance of Right.
*/
func (r Right) String() (p string) {
	switch r {
	case NoAccess:
		return rightsMap[0]
	case AllAccess:
		return rightsMap[895]
	}

	if kw, found := rightsMap[r]; found {
		p = kw
	}
	return
}

/*
String is a stringer method that returns the string representation of
the receiver instance.
*/
func (r Permission) String() string {
	if r.IsZero() {
		return badPerm
	}

	var rights []string
	if (*r.permission.Right) == AllAccess {
		rights = append(rights, AllAccess.String())
		return r.sprintf(rights)
	} else if (*r.permission.Right) == NoAccess {
		rights = append(rights, NoAccess.String())
		return r.sprintf(rights)
	}

	for i := 0; i < 16; i++ {
		right := Right(1 << i)
		if r.Positive(right) {
			rights = append(rights, right.String())
		}
	}
	return r.sprintf(rights)
}

func (r Permission) sprintf(rights []string) string {
	return sprintf("%s(%s)", r.Disposition(), join(rights, `,`))
}

/*
Disposition returns the string disposition `allow`
or 'deny', depending on the state of the receiver.
*/
func (r Permission) Disposition() string {
	return r.permission.disposition()
}

func (r permission) disposition() string {
	if r.isZero() {
		return `<unknown_disposition>`
	}

	if *r.bool {
		return `allow`
	}
	return `deny`
}

/*
Positive returns a boolean value indicative of whether a particular bit is positive (is set). Negation implies negative, or unset.
*/
func (r Permission) Positive(x any) bool {
	if err := r.Valid(); err != nil {
		return false
	}
	return r.permission.positive(x)
}

/*
Shift left-shifts the receiver instance to include Right x, if not already present.
*/
func (r Permission) Shift(x any) Permission {
	if err := r.Valid(); err != nil {
		return r
	}
	r.permission.shift(x)
	return r
}

/*
Unshift right-shifts the receiver instance to remove Right x, if present.
*/
func (r Permission) Unshift(x any) Permission {
	if err := r.Valid(); err != nil {
		return r
	}
	r.permission.unshift(x)
	return r
}

/*
IsZero returns a boolean value indicative of whether the receiver
is nil, or unset.
*/
func (r Permission) IsZero() bool {
	return r.permission.isZero()
}

func (r *permission) isZero() bool {
	if r == nil {
		return true
	}

	return r.bool == nil && r.Right == nil
}

/*
Valid returns a non-error instance if the receiver fails to pass
basic validity checks.
*/
func (r Permission) Valid() (err error) {
	if r.IsZero() {
		err = nilInstanceErr(r)
		return
	}

	if r.permission.bool == nil {
		err = noPermissionDispErr()
	}

	return
}

func init() {
	rightsMap = map[Right]string{
		NoAccess:        `none`,
		ReadAccess:      `read`,
		WriteAccess:     `write`,
		AddAccess:       `add`,
		DeleteAccess:    `delete`,
		SearchAccess:    `search`,
		CompareAccess:   `compare`,
		SelfWriteAccess: `selfwrite`,
		AllAccess:       `all`,
		ProxyAccess:     `proxy`,
		ImportAccess:    `import`,
		ExportAccess:    `export`,
	}

	// we want to resolve the name
	// of a Right into an actual
	// Right instance.
	rightsNames = make(map[string]Right, 0)
	for k, v := range rightsMap {
		rightsNames[v] = k
	}

}

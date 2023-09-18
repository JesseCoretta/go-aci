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

var badPermission Permission
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
	if !r.isZero() {
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
}

func (r *permission) unshift(x ...any) {
	if !r.isZero() {
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

func (r permission) positive(x any) (posi bool) {
	if !r.isZero() {
		switch tv := x.(type) {
		case int:
			posi = r.positiveIntRight(tv)

		case string:
			posi = r.positiveStrRight(tv)

		case Right:
			posi = r.positiveIntRight(int(tv))
		}
	}
	return
}

func (r permission) positiveIntRight(i int) (posi bool) {
	// avoid over/under flow
	if !(0 <= i && i <= int(^uint16(0))) {
		return
	}

	if _, matched := noneOrFullAccessRights(Right(i)); matched {
		if i == 0 {
			posi = int(*r.Right) == i
		} else {
			posi = ((*r.Right) & AllAccess) > 0
		}
		return
	}

	if rightsPowerOfTwo(i) {
		posi = ((*r.Right) & Right(i)) > 0
	}

	return
}

func (r permission) positiveStrRight(s string) (posi bool) {
	if R, matched := noneOrFullAccessString(lc(s), (*r.Right)); matched {
		if R == NoAccess {
			posi = int(*r.Right) == 0
		} else {
			posi = ((*r.Right) & AllAccess) > 0
		}
		return
	}

	// Resolve the name of a Right into a Right.
	if priv, found := rightsNames[lc(s)]; found {
		posi = ((*r.Right) & priv) > 0
	}

	return
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
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r Right) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Len returns the abstract integer length of the receiver, quantifying
the number of Right instances currently being expressed. For example,
if the receiver instance has its Read and Delete Right bits enabled,
this would represent an abstract length of two (2).
*/
func (r Permission) Len() int {
	if r.IsZero() {
		return 0
	}

	return r.permission.len()
}

func (r permission) len() int {
	var D int
	for i := 0; i < bitSize(NoAccess); i++ {
		if d := Right(1 << i); r.positive(d) {
			D++
		}
	}

	return D
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
	} else if (*r.permission.Right) == Right(1023) {
		rights = append(rights, AllAccess.String())
		rights = append(rights, ProxyAccess.String())
		return r.sprintf(rights)
	} else if (*r.permission.Right) == NoAccess {
		rights = append(rights, NoAccess.String())
		return r.sprintf(rights)
	}

	for i := 0; i < bitSize(NoAccess); i++ {
		right := Right(1 << i)
		if r.Positive(right) {
			rights = append(rights, right.String())
		}
	}
	return r.sprintf(rights)
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison
between the receiver (r) and input value x.
*/
func (r Permission) Compare(x any) bool {
	return compareHashInstance(r, x)
}

func (r Permission) sprintf(rights []string) string {
	return sprintf("%s(%s)", r.Disposition(), join(rights, `,`))
}

/*
Disposition returns the string disposition `allow`
or 'deny', depending on the state of the receiver.
*/
func (r Permission) Disposition() string {
	if r.permission == nil {
		return `<unknown_disposition>`
	}
	return r.permission.disposition()
}

func (r permission) disposition() (disp string) {
	disp = `<unknown_disposition>`
	if *r.bool {
		disp = `allow`
	} else if !*r.bool {
		disp = `deny`
	}
	return
}

/*
Positive returns a Boolean value indicative of whether a particular bit is positive (is set). Negation implies negative, or unset.
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
func (r Permission) Shift(x ...any) Permission {
	if err := r.Valid(); err != nil {
		return r
	}

	for i := 0; i < len(x); i++ {
		r.permission.shift(x[i])
	}
	return r
}

/*
Unshift right-shifts the receiver instance to remove Right x, if present.
*/
func (r Permission) Unshift(x ...any) Permission {
	if err := r.Valid(); err != nil {
		return r
	}

	for i := 0; i < len(x); i++ {
		r.permission.unshift(x[i])
	}
	return r
}

/*
IsZero returns a Boolean value indicative of whether the receiver
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
Parse wraps go-antlraci's ParsePermission function, writing
valid data into the receiver, or returning an error instance
if processing fails.
*/
func (r *Permission) Parse(raw string) (err error) {
	var perm *permission
	if perm, err = parsePermission(raw); err != nil {
		return
	}

	r.permission = perm
	return
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

	err = noPermissionDispErr()
	if r.permission.bool != nil {
		err = nil
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

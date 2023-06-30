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

var rightsMap map[Right]string

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

func Allow(x ...Right) Permission {
	return newPermission(true, x...)
}

func Deny(x ...Right) Permission {
	return newPermission(false, x...)
}

func newPermission(disp bool, x ...Right) Permission {
	p := new(permission)
	p.bool = &disp
	p.Right = new(Right)
	p.shift(x...)
	return Permission{p}
}

func (r *permission) shift(x ...Right) {
	if r.isZero() {
		return
	}

	for i := 0; i < len(x); i++ {
		if x[i] == NoAccess {
			(*r.Right) = NoAccess
			continue
		}
		(*r.Right) |= x[i]
	}
}

func (r *permission) unshift(x ...Right) {
	if r.isZero() {
		return
	}

	for i := 0; i < len(x); i++ {
		(*r.Right) = (*r.Right) &^ x[i]
	}
}

func (r permission) positive(x Right) bool {
	if r.isZero() {
		return false
	}
	return ((*r.Right) & x) > 0
}

/*
String is a stringer method that returns a single string name value for receiver instance of Right.
*/
func (r Right) String() (p string) {
	p = `none`
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
		return `<invalid_permission>`
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
	//return sprintf("%s(%s)", r.Kind(), join(rights, `,`))
	return r.sprintf(rights)
}

func (r Permission) sprintf(rights []string) string {
	return sprintf("%s(%s)", r.Kind(), join(rights, `,`))
}

/*
Kind returns the string disposition `allow`.
*/
func (r Permission) Kind() string {
	return r.permission.kind()
}

func (r permission) kind() string {
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
func (r Permission) Positive(x Right) bool {
	if err := r.Valid(); err != nil {
		return false
	}
	return r.permission.positive(x)
}

/*
Shift left-shifts the receiver instance to include Right x, if not already present.
*/
func (r Permission) Shift(x Right) Permission {
	if err := r.Valid(); err != nil {
		return r
	}
	r.permission.shift(x)
	return r
}

/*
Unshift right-shifts the receiver instance to remove Right x, if present.
*/
func (r Permission) Unshift(x Right) Permission {
	if err := r.Valid(); err != nil {
		return r
	}
	r.permission.unshift(x)
	return r
}

func (r *permission) isZero() bool {
	if r == nil {
		return true
	}

	return r.bool == nil && r.Right == nil
}

/*
IsZero returns a boolean value indicative of whether the receiver
is nil, or unset.
*/
func (r Permission) IsZero() bool {
	return r.permission.isZero()
}

/*
Valid returns a non-error instance if the receiver fails to pass
basic validity checks.
*/
func (r Permission) Valid() (err error) {
	if r.IsZero() {
		err = errorf("%T instance is nil", r)
		return
	}

	if r.permission.bool == nil {
		err = errorf("%T has no disposition (allow/deny)", r)
		return
	}

	return
}

func init() {
	rightsMap = map[Right]string{
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
}

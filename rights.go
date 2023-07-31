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
			// so long as the integer magnitude does
			// not fall out of uint8 bounds AND is a
			// power of two (2), we can interpolate
			// as a Right.
			if isPowerOfTwo(tv) && ( 0 <= tv && tv <= int(^uint16(0)) ) {
				(*r.Right) |= Right(tv)
			}
		case string:
			// Resolve the name of a Right into a Right.
			if priv, found := rightsNames[lc(tv)]; found {
				(*r.Right) |= priv
			}
		case Right:
			if tv == NoAccess {
				// NoAccess should stop the party
				// dead in its tracks, regardless
				// of any other iterated value.
				(*r.Right) = NoAccess
				return
			} else if tv == AllAccess {
				// AllAccess should stop the party
				// dead in its tracks, regardless
				// of any other iterated value.
				(*r.Right) = AllAccess
				return
			}

			(*r.Right) |= tv
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
                        // so long as the integer magnitude does
                        // not fall out of uint8 bounds AND is a
                        // power of two (2), we can interpolate
                        // as a Right.
                        if isPowerOfTwo(tv) && ( 0 <= tv && tv <= int(^uint16(0)) ) {
				(*r.Right) = (*r.Right) &^ Right(tv)
                        }
                case string:
                        if lc(tv) == NoAccess.String() {
				// Asking to remove NoAccess will
				// not accomplish anything definitive.
                                return
                        } else if lc(tv) == AllAccess.String() {
                                // Asking to remove all access
				// privileges is the same as
				// setting NoAccess outright.
                                (*r.Right) = NoAccess
                                return
                        }

                        // Resolve the name of a Right into a Right.
                        if priv, found := rightsNames[lc(tv)]; found {
				(*r.Right) = (*r.Right) &^ priv
                        }
                case Right:
                        if tv == NoAccess {
				// Asking to remove NoAccess will
				// not accomplish anything definitive.
                                return
                        } else if tv == AllAccess {
                                // Asking to remove all access
				// privileges is the same as
				// setting NoAccess outright.
                                (*r.Right) = NoAccess
                                return
                        }
			(*r.Right) = (*r.Right) &^ tv
		}
	}
}

func (r permission) positive(x any) bool {
	if r.isZero() {
		return false
	}
	switch tv := x.(type) {
	case int:
		if isPowerOfTwo(tv) && ( 0 <= tv && tv <= int(^uint16(0)) ) {
			return ((*r.Right) & Right(tv)) > 0
		}
	case string:
	        if lc(tv) == NoAccess.String() {
			// NoAccess always equals zero (0)
	                return int(*r.Right) == 0
	        } else if lc(tv) == AllAccess.String() {
			// See if the effective bit value
			// is equalTo the known "all"
			// compound value.
	                return int(*r.Right) == 895
	        }

		// Resolve the name of a Right into a Right.
	        if priv, found := rightsNames[lc(tv)]; found {
			return ((*r.Right) & priv) > 0
	        }
	case Right:
		return ((*r.Right) & tv) > 0
	}

	// unsupported type?
	return false
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
idRight will attempt to resolve a string value into a known privilege type (e.g.: `read` -> ReadAccess).
If found, the proper Right instance is returned alongside a success-indicative boolean. If nothing was
matched, NoAccess is returned alongside false bool.
*/
func idRight(def string) (r Right, ok bool) {
	r = NoAccess
	for k, v := range rightsMap {
		if eq(def,v) {
			r = k
			ok = uint8(k) != 0x0
			break
		}
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

        // we want to resolve the name
        // of a Right into an actual
        // Right instance.
	rightsNames = make(map[string]Right, 0)
        for k, v := range rightsMap {
                rightsNames[v] = k
        }

}

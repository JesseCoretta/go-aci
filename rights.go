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
			if rightsPowerOfTwo(tv) {
				(*r.Right) |= Right(tv)
			}
		case string:
			// Resolve the name of a Right into a Right.
			if priv, found := rightsNames[lc(tv)]; found {
				(*r.Right) |= priv
			}
		case Right:
			if R, matched := noneOrFullAccessRights(tv); matched {
				(*r.Right) = R
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
			if rightsPowerOfTwo(tv) {
				(*r.Right) = (*r.Right) &^ Right(tv)
			}
		case string:
			//if matched := noneOrFullAccessString(lc(tv), (*r.Right)); matched {
			//        return
			//}

			// Resolve the name of a Right into a Right.
			if priv, found := rightsNames[lc(tv)]; found {
				(*r.Right) = (*r.Right) &^ priv
			}
		case Right:
			if _, matched := noneOrFullAccessRights(tv); matched {
				return
			}
			(*r.Right) = (*r.Right) &^ tv
		}
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
		if rightsPowerOfTwo(tv) {
			return ((*r.Right) & Right(tv)) > 0
		}
	case string:
		if matched := noneOrFullAccessString(lc(tv), (*r.Right)); matched {
			return matched
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

func noneOrFullAccessRights(x Right) (Right, bool) {
	var matched bool
	if x == NoAccess {
		// NoAccess should stop the party
		// dead in its tracks, regardless
		// of any other iterated value.
		x = NoAccess
		matched = true

	} else if x == AllAccess {
		// AllAccess should stop the party
		// dead in its tracks, regardless
		// of any other iterated value.
		x = AllAccess
		matched = true
	}

	return x, matched
}

func noneOrFullAccessString(x string, r Right) (matched bool) {
	switch {
	case x == NoAccess.String():
		matched = int(r) == 0 // 0
	case x == AllAccess.String():
		matched = int(r) == 895 // *895
	}

	return
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
		if eq(def, v) {
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
	//return r.sprintf(rights)
	return sprintf("%s(%s)", r.Disposition(), join(rights, `,`))
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
		err = errorf("%T instance is nil", r)
		return
	}

	if r.permission.bool == nil {
		err = errorf("%T has no disposition (allow/deny)", r)
		return
	}

	return
}

/*
parsePerm reads and processes a token stream into an instance of Permission, which
is the first (1st) component of a PermissionBindRule instance. It, alongside an
error and chop index, are returned when processing stops or completes.
*/
func parsePerm(tokens []string) (chop int, perm Permission, err error) {
	var disp string
	var privs []any
	var done bool

	for _, token := range tokens {
		// closing paren during perm
		// mode means perm has ended
		// and at least one (1) bind
		// rule is beginning.
		switch lc(token) {
		case `allow`, `deny`:
			disp = lc(token)
		case `;`, `(`, `,`:
			// do nothing
		case `)`:
			done = true
		default:
			privs = append(privs, lc(token))
		}

		chop++
		if done {
			break
		}
	}

	// assemble permission
	perm = assemblePermissionByDisposition(disp, privs)

	return
}

/*
assemblePermissionByDisposition shifts the return value to include all privilege
abstraction found within privs, and will initialize said return value based on
the disposition (allow or deny) selected by the user.
*/
func assemblePermissionByDisposition(disp string, privs []any) (perm Permission) {
	if disp == `allow` {
		if len(privs) == 0 {
			perm = Allow(`none`)
		} else {
			perm = Allow(privs...)
		}
		return
	}

	if len(privs) == 0 {
		perm = Deny(`all`, `proxy`)
	} else {
		perm = Deny(privs...)
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

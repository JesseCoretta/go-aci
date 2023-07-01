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
func Allow(x ...Right) Permission {
	return Permission{newPermission(true, x...)}
}

/*
Deny returns a withholding Permission instance bearing the provided
instances of Right.
*/
func Deny(x ...Right) Permission {
	return Permission{newPermission(false, x...)}
}

/*
newPermission returns a newly initialized instance of *permission
bearing the provided disposition and Right instance(s).
*/
func newPermission(disp bool, x ...Right) (p *permission) {
	p = new(permission)
	p.bool = &disp
	p.Right = new(Right)
	p.shift(x...)
	return p
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
Parse returns an instance of error following an attempt to initialize and set
the receiver based upon the  string-based privilege input statement (def).
*/
func (r Permission) Parse(def string) error {
	var err error
        r.permission, _, err = r.permission.parse(def)
	return err
}

/*
parsePermission reads the ACI definition string value and extracts
the first permission statement encountered. The receiver instance is
updated with the information, and an error is returned in the event
the attempt fails.

The return integer value (last) defines the index value at which the
privilege string value is believed to terminate. This is useful for
efficient top-level ACI parsing to know where to begin scanning for
one (1) or more bind rules immediately after a privilege statement.
*/
func (r *permission) parse(def string) (p *permission, last int, err error) {

        var allow bool
	// First attempt to find the disposition. It can
	// only be allow or deny.
        if hasPfx(def, `allow`) {
                allow = true
        } else if !hasPfx(def, `deny`){
		// if it wasn't allow above, and it isn't
		// deny here, we can't go any further.
                err = errorf("%T parsing failed: indeterminate disposition (must be 'allow' or 'deny')",
			Permission{})
                return
        }

	// get left paren idx
        lidx := idxr(def,'(');
        if lidx == -1 {
                err = errorf("%T parsing failed: unopened rights definition",
			Permission{})
                return
        }

	// get right paren idx
        ridx := idxr(def,')')
        if ridx == -1 {
                err = errorf("%T parsing failed: unclosed rights definition",
			Permission{})
                return
        }

	// Save righthanded index for subsequent PB
	// rule parsing if this involves a toplevel
	// ACI parsing attempt.
	last = ridx

	// extract the (apparent) privilege
	// statement(s) from between the two
	// above L/R indices. Additionally,
	// replace all WHSP with nothing,
	// and split the parenthetical value
	// on its commas into a string slice.
        rights := split(repAll(def[lidx+1:ridx], string(whsp), ``), `,`)
        if len(rights) == 0 {
		// Even if `none` is desired, you'd
		// have to actually put the `none`
		// string into the ACI definition.
		// We found nothing, so we can't go
		// further.
                err = errorf("%T parsing failed: no specific privileges specified",
			Permission{})
                return
        }

	// Begin new private instance
	p = newPermission(allow)

        // Iterate each string-based privilege keyword,
        // and attempt to resolve into a known Right
        // constant. If found, shift the perms bits to
	// include said constant.
        for i := 0; i < len(rights); i++ {

		if i == len(rightsMap) {
			// Since you can't add the same
			// priv twice, only a certain
			// iteration count makes sense.
			// Don't spin wheels forever.
			break
		}

		// Begin resolution attempt
                if pr, ok := idRight(rights[i]); ok {
			// We found *something* as a result of
			// the resolution attempt. Examine it.
                        if pr == NoAccess && !r.isZero() {
				// If they requested NoAccess and
				// there were already privileges
				// set, then wipe them out and bail.
                                p.shift(NoAccess)
                                break
                        }
			// Add the verified privilege.
                        p.shift(pr)
                } else {
			// Resolution attempt failed
			err = errorf("%T failed: unresolvable privilege keyword '%s'", rights[i])
			break
		}
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

package aci

/*
kw.go contains Keyword / BindTypes methods and constants.
*/

/*
Keyword describes the effective "type" within the context of a given [BindRule] or [TargetRule]. The available [Keyword] instances vary based on the rule type in which a given [Keyword] resides.

See the [Keyword] constants defined in this package for a complete list.
*/
type Keyword interface {
	String() string
	Kind() string
}

/*
private keyword maps exist only to keep cyclomatics down.
*/
var (
	bkwMap map[BindKeyword]string
	tkwMap map[TargetKeyword]string
	btMap  map[BindType]string
)

/*
BindKeyword contains the value describing a particular [BindKeyword] to be used within a [BindRule].
*/
type BindKeyword uint8

/*
TargetKeyword contains the value describing a particular Target Keyword to be used within a [TargetRule].
*/
type TargetKeyword uint8

/*
BindType describes one (1) of five (5) possible contexts used in certain [BindRule] instances:

  - [USERDN]
  - [GROUPDN]
  - [ROLEDN]
  - [SELFDN]
  - [LDAPURL]
*/
type BindType uint8

/*
keyword/type placeholders for bad definitions.
*/
const (
	badBT  = `<invalid_bind_type>`
	badBKW = `<invalid_bind_keyword>`
	badTKW = `<invalid_target_keyword>`
)

/*
BindType keyword constants are used in value matching definitions that utilizes either the [BindUAT] (userattr) or [BindGAT] (groupattr) [BindKeyword] constant within a [BindRule] instance.
*/
const (
	_ BindType = iota // <invalid_bind_type>
	USERDN
	GROUPDN
	ROLEDN
	SELFDN
	LDAPURL
)

/*
BindKeyword constants are intended for singular use within a [BindRule] instance.
*/
const (
	_       BindKeyword = iota // <invalid_bind_keyword>
	BindUDN                    // `userdn`
	BindRDN                    // `roledn`
	BindGDN                    // `groupdn`
	BindUAT                    // `userattr`
	BindGAT                    // `groupattr`
	BindIP                     // `ip`
	BindDNS                    // `dns`
	BindDoW                    // `dayofweek`
	BindToD                    // `timeofday`
	BindAM                     // `authmethod`
	BindSSF                    // `ssf`
)

/*
TargetKeyword constants are intended for singular use within a [TargetRule] instance.
*/
const (
	_                 TargetKeyword = iota // <invalid_target_keyword>
	Target                                 // 0x1, target
	TargetTo                               // 0x2, target_to
	TargetAttr                             // 0x3, targetattr
	TargetCtrl                             // 0x4, targetcontrol
	TargetFrom                             // 0x5, target_from
	TargetScope                            // 0x6, targetscope
	TargetFilter                           // 0x7, targetfilter
	TargetAttrFilters                      // 0x8, targattrfilters (yes, "targ". As in "wild Klingon boars").
	TargetExtOp                            // 0x9, extop
)

/*
String is a stringer method that returns the string representation of the receiver instance of [BindType].
*/
func (r BindType) String() (b string) {
	b = badBT
	if kw, found := btMap[r]; found {
		b = kw
	}
	return
}

/*
Kind returns the static string literal `bind` identifying the instance as a [BindKeyword].
*/
func (r BindKeyword) Kind() string {
	return bindRuleID
}

/*
Kind returns the static string literal `target` identifying the instance as a [TargetKeyword].
*/
func (r TargetKeyword) Kind() string {
	return targetRuleID
}

/*
String is a stringer method that returns the string representation of the receiver instance of [BindKeyword].
*/
func (r BindKeyword) String() (k string) {
	k = badBKW
	if kw, found := bkwMap[r]; found {
		k = kw
	}
	return
}

/*
String is a stringer method that returns the string representation of the receiver instance of [TargetKeyword].
*/
func (r TargetKeyword) String() (k string) {
	k = badTKW
	if kw, found := tkwMap[r]; found {
		k = kw
	}
	return
}

func assertATBTVBindKeyword(bkw ...any) (kw BindKeyword) {
	kw = BindUAT
	if len(bkw) == 0 {
		return
	}

	switch tv := bkw[0].(type) {
	case BindKeyword:
		if tv == BindGAT {
			kw = tv
		}
	}

	return
}

/*
matchTKW will return the matching TargetKeyword constant for the input kw string value.
*/
func matchTKW(kw string) TargetKeyword {
	for k, v := range tkwMap {
		if eq(kw, v) {
			return k
		}
	}

	return TargetKeyword(0x0)
}

/*
matchBKW will return the matching BindKeyword constant for the input kw string value.
*/
func matchBKW(kw string) BindKeyword {
	for k, v := range bkwMap {
		if eq(kw, v) {
			return k
		}
	}

	return BindKeyword(0x0)
}

/*
matchBT will return the matching BindType constant for the input kw string value.
*/
func matchBT(kw string) BindType {
	for k, v := range btMap {
		if eq(kw, v) {
			return k
		}
	}

	return BindType(0x0)
}

func init() {
	// bindkeyword map
	bkwMap = map[BindKeyword]string{
		BindUDN: `userdn`,
		BindRDN: `roledn`,
		BindGDN: `groupdn`,
		BindUAT: `userattr`,
		BindGAT: `groupattr`,
		BindIP:  `ip`,
		BindDNS: `dns`,
		BindDoW: `dayofweek`,
		BindToD: `timeofday`,
		BindAM:  `authmethod`,
		BindSSF: `ssf`,
	}

	// targetkeyword map
	tkwMap = map[TargetKeyword]string{
		Target:            `target`,
		TargetTo:          `target_to`,
		TargetAttr:        `targetattr`,
		TargetCtrl:        `targetcontrol`,
		TargetFrom:        `target_from`,
		TargetScope:       `targetscope`,
		TargetFilter:      `targetfilter`,
		TargetAttrFilters: `targattrfilters`,
		TargetExtOp:       `extop`,
	}

	// bindtype map
	btMap = map[BindType]string{
		USERDN:  `USERDN`,
		ROLEDN:  `ROLEDN`,
		SELFDN:  `SELFDN`,
		GROUPDN: `GROUPDN`,
		LDAPURL: `LDAPURL`,
	}
}

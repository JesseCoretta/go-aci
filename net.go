package aci

/*
net.go contains types, methods and constants that relate to the use of IP addresses and DNS names within Bind Rules.
*/

/*
IPAddr embeds slices of address values, allowing simple composition of flexible IP-based [BindRule] instances.
*/
type IPAddr struct {
	*ipAddrs
}

/*
IP initializes, sets and returns a new instance of [IPAddr] in one shot. This function is an alternative to separate assignment and set procedures.
*/
func IP(addr ...string) IPAddr {
	return newIPAddr(addr...)
}

func newIPAddr(addr ...string) IPAddr {
	x := new(ipAddrs)
	if len(addr) > 0 {
		x.set(addr...)
	}
	return IPAddr{x}
}

type ipAddrs []ipAddr
type ipAddr string

const (
	badAddr = `<invalid_address_list>`
)

/*
Keyword returns the [BindKeyword] instance assigned to the receiver instance as a [Keyword]. This shall be the [BindKeyword] that appears in a [BindRule] containing the receiver instance as the expression value.
*/
func (r FQDN) Keyword() Keyword {
	return BindDNS
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r FQDN) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Eq initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Equal-To the [BindIP] [BindKeyword] context.
*/
func (r IPAddr) Eq() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(BindIP, Eq, r)
}

/*
Ne initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Not-Equal-To the [BindIP] [BindKeyword] context.

Negated equality [BindRule] instances should be used with caution.
*/
func (r IPAddr) Ne() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(BindIP, Ne, r)
}

/*
BRM returns an instance of [BindRuleMethods].

Each of the return instance's key values represent a single instance of the [ComparisonOperator] type that is allowed for use in the creation of [BindRule] instances which bear the receiver instance as an expression value. The value for each key is the actual [BindRuleMethod] instance for OPTIONAL use in the creation of a [BindRule] instance.

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus BindRule instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly populated and prepared for such activity.
*/
func (r IPAddr) BRM() BindRuleMethods {
	return newBindRuleMethods(bindRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
Compare returns a Boolean value indicative of a SHA-1 comparison between the receiver (r) and input value x.
*/
func (r IPAddr) Compare(x any) bool {
	return compareHashInstance(r, x)
}

/*
Len returns the integer length of the receiver instance.
*/
func (r IPAddr) Len() int {
	if r.ipAddrs == nil {
		return 0
	}
	return len(*r.ipAddrs)
}

/*
Keyword returns the [BindKeyword] assigned to the receiver instance. This shall be the keyword that appears in a [BindRule] containing the receiver instance as the expression value.
*/
func (r IPAddr) Keyword() Keyword {
	return BindIP
}

/*
Kind returns the string representation of the receiver's kind.
*/
func (r IPAddr) Kind() string {
	return BindIP.String()
}

/*
Set assigns the provided address component to the receiver and returns the receiver instance in fluent-form.

Multiple values can be provided in variadic form, or piecemeal.
*/
func (r *IPAddr) Set(addr ...string) *IPAddr {
	if r.ipAddrs == nil {
		r.ipAddrs = new(ipAddrs)
	}

	r.ipAddrs.set(addr...)
	return r
}

func (r *ipAddrs) set(addr ...string) {
	for i := 0; i < len(addr); i++ {
		if len(addr[i]) > 0 && r.unique(addr[i]) {
			if isValidIP(addr[i]) {
				*r = append(*r, ipAddr(addr[i]))
			}
		}
	}
}

func isValidIP(x string) bool {
	return isV4(x) || isV6(x)
}

func isV4(x string) bool {
	if len(x) <= 1 {
		return false
	}

	for c := 0; c < len(x); c++ {
		char := rune(byte(lc(string(x[c]))[0]))
		if !isValidV4Char(char) {
			return false
		}
	}

	return true
}

func isValidV4Char(char rune) bool {
	return ('0' <= char && char <= '9') || char == '.' || char == '*' || char == '/'
}

func isV6(x string) bool {
	if len(x) <= 1 {
		return false
	}

	for c := 0; c < len(x); c++ {
		char := rune(byte(lc(string(x[c]))[0]))
		if !isValidV6Char(char) {
			return false
		}
	}

	return true
}

func isValidV6Char(char rune) bool {
	return ('0' <= char && char <= '9') || ('a' <= char && char <= 'f') || char == ':' || char == '*' || char == '/'
}

/*
IsZero returns a Boolean value indicative of whether the receiver is considered nil, or unset.
*/
func (r IPAddr) IsZero() bool {
	if r.ipAddrs == nil {
		return true
	}

	return r.ipAddrs.isZero()
}

/*
Valid returns an error indicative of whether the receiver is in an aberrant state.
*/
func (r IPAddr) Valid() error {
	if r.IsZero() {
		return nilInstanceErr(r)
	}

	if r.Len() == 0 {
		return nilInstanceErr(r)
	}

	return nil
}

func (r *ipAddrs) isZero() bool {
	return r == nil
}

/*
unique scans the receiver to verify whether the addr input value is not already present within the receiver.
*/
func (r IPAddr) unique(addr string) bool {
	if r.IsZero() {
		return true
	}

	return r.ipAddrs.unique(addr)
}

func (r ipAddrs) unique(addr string) bool {
	var addrs []string
	for i := 0; i < len(r); i++ {
		addrs = append(addrs, string(r[i]))
	}

	return !strInSlice(addr, addrs)
}

/*
String is a stringer method that returns the string representation of an IP address.
*/
func (r IPAddr) String() string {
	if r.isZero() {
		return badAddr
	}

	var str []string
	for i := 0; i < len(*r.ipAddrs); i++ {
		str = append(str, string((*r.ipAddrs)[i]))
	}
	return sprintf("%s", join(str, `,`))
}

//////////////////////////////////////////////////////////////////////////////////
// Begin DNS/FQDN
//////////////////////////////////////////////////////////////////////////////////

/*
domainLabel represents a single component within a fully-qualified domain name. Multiple occurrences of ordered instances of this type represent a complete FQDN, which may include wildcards (*), to be used in DNS-based ACIs.
*/
type domainLabel []byte
type labels []domainLabel

/*
FQDN contains ordered domain labels that form a fully-qualified domain name.
*/
type FQDN struct {
	*labels
}

/*
DNS initializes, sets and returns a new instance of [FQDN] in one shot. This function is an alternative to separate assignment and set procedures.
*/
func DNS(label ...string) FQDN {
	x := FQDN{new(labels)}
	if len(label) > 0 {
		x.Set(label...)
	}
	return x
}

/*
zeroDomain is used in as a return instance for failed [FQDN]-related Set operations.
*/
var zeroDomain FQDN

const (
	fqdnMax  = 253
	labelMax = 63
	badFQDN  = `<invalid_fqdn_or_label>`
)

/*
Len returns the abstract integer length of the receiver. The value returned represents the number of valid DNS labels within a given instance of FQDN. For example, `www.example.com` has three (3) such labels.
*/
func (r FQDN) Len() int {
	if r.labels == nil {
		return 0
	}
	return len(*r.labels)
}

/*
Set appends one or more domain labels to the receiver. The total character length of a single label CANNOT exceed sixty-three (63) characters.  When added up, all domain label instances present within the receiver SHALL NOT collectively exceed two hundred fifty-three (253) characters.

Valid characters within labels:

  - a-z
  - A-Z
  - 0-9
  - Hyphen ('-', limited to [1:length-1] slice range)
  - Asterisk ('*', use with care for wildcard DNS-based ACI [BindRule] expressions)
  - Full Stop ('.', see below for remarks on this character)

Users need not enter full stops (.) manually, given this method supports the use of variadic expressions, i.e.:

	Set(`www`,`example`,`com`)

However, should full stops (.) be used within input values:

	Set(`www.example.com`)

... the parser shall split the input into label components and add them to the receiver piecemeal in the intended order.

Please note that it is not necessary to include a NULL terminating full stop character (.) at the end (TLD?) of the intended [FQDN].
*/
func (r *FQDN) Set(label ...string) *FQDN {
	if r.IsZero() {
		*r = FQDN{new(labels)}
	}

	r.labels.set(label...)
	return r
}

func (r *labels) set(label ...string) {
	if len(label) == 0 {
		return
	}

	/*
		if r.isZero() {
			r = new(labels)
		}
	*/

	dl, c, ok := processLabel(label...)
	if !ok {
		return
	}

	// Only update the receiver if
	// we haven't breached the high
	// water mark ...
	if len(*r)+c <= fqdnMax {
		for l := 0; l < len(dl); l++ {
			*r = append(*r, dl[l])
		}
	}

	return
}

func processLabel(label ...string) (dl labels, c int, ok bool) {
	for i := 0; i < len(label); i++ {
		if !validLabel(label[i]) {
			return
		}

		if idx := idxr(label[i], '.'); idx != -1 {
			sp := split(label[i], `.`)
			for j := 0; j < len(sp); j++ {
				// null label doesn't
				// need to stop the
				// show.
				if !validLabel(sp[j]) {
					return
				}
				c += len(sp[j])
				dl = append(dl, domainLabel(sp[j]))
			}
		} else {
			c += len(label[i])
			dl = append(dl, domainLabel(label[i]))
		}
	}

	ok = c > 0 && len(dl) > 0
	return
}

/*
String is a stringer method that returns the string representation of a fully-qualified domain name.
*/
func (r FQDN) String() string {
	if err := r.Valid(); err != nil {
		return badFQDN
	}

	var str []string

	for i := 0; i < len(*r.labels); i++ {
		str = append(str, string((*r.labels)[i]))
	}

	return join(str, `.`)
}

/*
Eq initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Equal-To the [BindDNS] [BindKeyword] context.
*/
func (r FQDN) Eq() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(BindDNS, Eq, r)
}

/*
Ne initializes and returns a new [BindRule] instance configured to express the evaluation of the receiver value as Not-Equal-To the [BindDNS] [BindKeyword] context.

Negated equality [BindRule] instances should be used with caution.
*/
func (r FQDN) Ne() BindRule {
	if err := r.Valid(); err != nil {
		return badBindRule
	}
	return BR(BindDNS, Ne, r)
}

/*
BRM returns an instance of [BindRuleMethods].

Each of the return instance's key values represent a single instance of the [ComparisonOperator] type that is allowed for use in the creation of [BindRule] instances which bear the receiver instance as an expression value. The value for each key is the actual [BindRuleMethod] instance for OPTIONAL use in the creation of a [BindRule] instance.

This is merely a convenient alternative to maintaining knowledge of which [ComparisonOperator] instances apply to which types. Instances of this type are also used to streamline package unit tests.

Please note that if the receiver is in an aberrant state, or if it has not yet been initialized, the execution of ANY of the return instance's value methods will return bogus [BindRule] instances. While this is useful in unit testing, the end user must only execute this method IF and WHEN the receiver has been properly populated and prepared for such activity.
*/
func (r FQDN) BRM() BindRuleMethods {
	return newBindRuleMethods(bindRuleFuncMap{
		Eq: r.Eq,
		Ne: r.Ne,
	})
}

/*
IsZero returns a Boolean value indicative of whether the receiver is nil, or unset.
*/
func (r FQDN) IsZero() bool {
	return r.labels.isZero()
}

func (r *labels) isZero() bool {
	return r == nil
}

/*
Valid returns a Boolean value indicative of whether the receiver contents represent a legal fully-qualified domain name value.
*/
func (r FQDN) Valid() (err error) {
	L := r.len()

	if !(0 < L && L <= fqdnMax) || len(*r.labels) < 2 {
		err = fqdnInvalidLenErr(L)
		return
	}

	// seems legit
	return
}

/*
Len returns the integer length of the receiver in terms of character count.
*/
func (r FQDN) len() int {
	if r.labels == nil {
		return 0
	}

	var c int
	for i := 0; i < len(*r.labels); i++ {
		for j := 0; j < len(*r.labels); j++ {
			c++
		}
	}

	return c
}

/*
validLabel returns a Boolean value indicative of whether the input value (label) represents a valid label component for use within a fully-qualified domain.
*/
func validLabel(label string) bool {
	// Cannot exceed maximum component lengths!
	if !(0 < len(label) && len(label) <= labelMax) {
		return false
	}

	for i := 0; i < len(label); i++ {
		if ok := labelCharsOK(rune(label[i]), i, len(label)-1); !ok {
			return ok
		}
	}

	// seems legit
	return true
}

func labelCharsOK(c rune, i, l int) (ok bool) {
	// Cannot contain unsupported characters!
	if !isDigit(c) && !isLetter(c) &&
		c != '.' && c != '*' && c != '-' {
		return
	}

	// Cannot begin or end with hyphen!
	if c == '-' && (i == 0 || i == l) {
		return
	}

	ok = true
	return
}

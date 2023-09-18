package aci

import (
	"fmt"
	"testing"
)

func TestParseAttributeBindTypeOrValue(t *testing.T) {
	var avb AttributeBindTypeOrValue
	_ = avb.Valid()
	_ = avb.IsZero()
	raw := `manager#SELFDN`

	for _, kw := range []any{
		3,
		`groupattr`,
		BindGAT,
	} {
		if err := avb.Parse(raw, kw); err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		}
	}
}

func TestParseBindRule(t *testing.T) {
	want := `userdn = "ldap:///cn=Jesse Coretta,ou=People,dc=example,dc=com" || "ldap:///anyone"`

	var b BindRule
	var err error
	_, _ = ParseBindRule(``)
	_ = b.Parse(``)

	if b, err = ParseBindRule(want); err != nil {
		return
	}

	if want != b.String() {
		t.Errorf("%s failed:\nwant '%s'\ngot '%s'", t.Name(), want, b)
		return
	}

	if err = b.Parse(want); err != nil {
		return
	}

	if want != b.String() {
		t.Errorf("%s failed:\nwant '%s'\ngot '%s'", t.Name(), want, b)
		return
	}
}

func TestParseBindRules_codecov(t *testing.T) {
	single := `userdn = "ldap:///anyone"`
	var b BindRules
	var err error
	if err = b.Parse(single); err != nil {
		return
	}

	if b.String() != single {
		t.Errorf("%s failed:\nwant '%s',\ngot  '%s'", t.Name(), single, b)
		return
	}
}

func TestParseBindRules(t *testing.T) {
	want := `( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) )`

	var r BindContext
	var err error

	_, _ = ParseBindRules(``)

	if r, err = ParseBindRules(want); err != nil {
		return
	}

	if want != r.String() {
		t.Errorf("%s failed:\nwant '%s',\ngot  '%s'", t.Name(), want, r)
		return
	}

	if r.Keyword() == nil {
		t.Errorf("%s failed: unidentified %T", t.Name(), r.Keyword())
		return
	}

	if !r.IsNesting() {
		t.Errorf("%s failed: nesting not detected", t.Name())
		return
	}

	bl := r.Len()
	orig := r.String()

	R, _ := r.(BindRules)

	R.Push(BindRules{})

	var ctx BindContext = BindRule{}

	if R.Push(ctx); r.Len() != bl {
		t.Errorf("%s failed: bogus enveloped content was pushed into %T", t.Name(), r)
		return
	}

	popped := R.Pop()
	bl = r.Len()
	if popped.String() != orig {
		t.Errorf("%s failed: unexpected element popped; want '%s', got '%s'", t.Name(), orig, popped)
		return
	}

	R.Push(popped)
	R.remove(r.Len() - 1)
	if r.Len() != bl {
		t.Errorf("%s failed: content not removed from %T", t.Name(), r)
		return
	}

	R.insert(popped, 0)
	if r.Len() == bl {
		t.Errorf("%s failed: content not inserted into %T", t.Name(), r)
		return
	}
}

func TestBindRules_Parse_codecov(t *testing.T) {
	var br BindRules
	_ = br.Parse(``)
	_ = br.Parse(`%$#^&*iB%^*O%&^#G*%^*U(&%^S#&*%^#&*`)
}

func ExampleBindRules_Parse_messy() {
	raw := `(
                        (
                                ( userdn = "ldap:///anyone" ) AND
                                ( ssf >= "71" )

                        ) AND NOT (
                                dayofweek = "Wed" OR
                                dayofweek = "Fri"
                        )
        )`

	br, err := ParseBindRules(raw)
	if err != nil {
		fmt.Println(err) // always check your parser errors
		return
	}

	called := br.Traverse(0, 0, 0)
	fmt.Printf("%s", called)
	// Output: ( userdn = "ldap:///anyone" )
}

func ExampleBindRules_Parse() {
	raw := `( ssf >= "128" AND ( authmethod = "SASL" OR authmethod = "SSL" ) )`
	var brs BindRules
	if err := brs.Parse(raw); err != nil {
		fmt.Println(err) // always check your parser errors
		return
	}

	fmt.Printf("%s", brs.Traverse(0, 1, 0))
	// Output: authmethod = "SASL"
}

func ExampleBindRule_Parse() {
	raw := `ssf >= "128"`
	var br BindRule
	if err := br.Parse(raw); err != nil {
		fmt.Println(err) // always check your parser errors
		return
	}

	br.NoPadding(true)
	br.Paren(true)

	fmt.Printf("%s", br)
	// Output: (ssf>="128")
}

/////////////////////////////////////////////////////////////////////
/// begin TargetRule tests
/////////////////////////////////////////////////////////////////////

/*
This example demonstrates the imported ANTLR4-based go-antlraci parser capabilities as
they pertain to the handling of raw target rule text.
*/
func ExampleParseTargetRule() {

	// NOTE: padding manually stripped out, and an
	// extraneous horizontal tab (ASCII #9) added
	// for purely demonstrative reasons ...
	raw := `(target_to=     "ldap:///cn=*,ou=Contractors,ou=People,dc=example,dc=com")`
	tr, err := ParseTargetRule(raw)
	if err != nil {
		fmt.Println(err) // always check your parser errors.
		return
	}
	fmt.Printf("%s", tr)
	// Output: ( target_to = "ldap:///cn=*,ou=Contractors,ou=People,dc=example,dc=com" )
}

/*
This example demonstrates the imported ANTLR4-based go-antlraci parser capabilities as
they pertain to the handling of raw target rule text that contains multiple values with
specific delimiters and standard quotation.

Additionally, upon receiving the returned value, we'll disable padding just for fun.
*/
func ExampleParseTargetRule_multiValuedWithStandardQuotation() {

	raw := `(target_to="ldap:///cn=*,ou=Contractors,ou=People,dc=example,dc=com"||"ldap:///cn*,ou=X.500 Administrators,ou=People,dc=example,dc=com")`
	tr, err := ParseTargetRule(raw)
	if err != nil {
		fmt.Println(err) // always check your parser errors.
		return
	}
	fmt.Printf("%s", tr.NoPadding(true))
	// Output: (target_to="ldap:///cn=*,ou=Contractors,ou=People,dc=example,dc=com" || "ldap:///cn*,ou=X.500 Administrators,ou=People,dc=example,dc=com")
}

/*
This example demonstrates the imported ANTLR4-based go-antlraci parser capabilities as
they pertain to the handling of raw target rule text that contains multiple values with
specific delimiters and alternative quotation.
*/
func ExampleParseTargetRule_multiValuedWithAlternativeQuotation() {

	raw := `(target_to="ldap:///cn=*,ou=Contractors,ou=People,dc=example,dc=com||ldap:///cn=*,ou=X.500 Administrators,ou=People,dc=example,dc=com")`
	tr, err := ParseTargetRule(raw)
	if err != nil {
		fmt.Println(err) // always check your parser errors.
		return
	}
	fmt.Printf("%s", tr)
	// Output: ( target_to = "ldap:///cn=*,ou=Contractors,ou=People,dc=example,dc=com || ldap:///cn=*,ou=X.500 Administrators,ou=People,dc=example,dc=com" )
}

/*
This example demonstrates the imported ANTLR4-based go-antlraci parser capabilities as
they pertain to the handling of a sequence of raw target rule text values. Note in this
example, we've added awkward spacing mixed-in with fair attempts to make the sequence of
TargetRule expressions easier to read. This includes newline characters (ASCII #10) to
really try and mess things up. 😈
*/
func ExampleParseTargetRules() {

	omg := `(
                target_to=
                        "ldap:///cn=*,ou=Contractors,ou=People,dc=example,dc=com"               ||
                        "ldap:///cn=*,ou=X.500 Administrators,ou=People,dc=example,dc=com"      ||
                        "ldap:///cn=*,ou=Executives,ou=People,dc=example,dc=com"
                )

                ( targetscope="subordinate"  )

                (
                        targattrfilters =
                                "add=nsroleDN:(!(nsroledn=cn=X.500 Administrator)) && employeeStatus:(!(drink=beer)) && telephoneNumber:(telephoneNumber=612*)"
                )`

	tr, err := ParseTargetRules(omg)
	if err != nil {
		fmt.Println(err) // always check your parser errors.
		return
	}

	fmt.Printf("%s", tr)
	// Output: ( target_to = "ldap:///cn=*,ou=Contractors,ou=People,dc=example,dc=com" || "ldap:///cn=*,ou=X.500 Administrators,ou=People,dc=example,dc=com" || "ldap:///cn=*,ou=Executives,ou=People,dc=example,dc=com" )( targetscope = "subordinate" )( targattrfilters = "add=nsroleDN:(!(nsroledn=cn=X.500 Administrator)) && employeeStatus:(!(drink=beer)) && telephoneNumber:(telephoneNumber=612*)" )
}

/*
This example is the same as the TargetRules example, except with the alternative
quotation scheme in effect.
*/
func ExampleParseTargetRules_alternativeQuotation() {

	omg := `(
                target_to=
                        "ldap:///cn=*,ou=Contractors,ou=People,dc=example,dc=com                ||
                         ldap:///cn=*,ou=X.500 Administrators,ou=People,dc=example,dc=com       ||
                         ldap:///cn=*,ou=Executives,ou=People,dc=example,dc=com"
                )

                ( targetscope=
                                "subordinate"
                )

                ( targattrfilters =
                                "add=nsroleDN:(!(nsroledn=cn=X.500 Administrator))      &&
                                 employeeStatus:(!(drink=beer))                         &&
                                 telephoneNumber:(telephoneNumber=612*)"
                )`

	tr, err := ParseTargetRules(omg)
	if err != nil {
		fmt.Println(err) // always check your parser errors.
		return
	}

	fmt.Printf("%s", tr)
	// Output: ( target_to = "ldap:///cn=*,ou=Contractors,ou=People,dc=example,dc=com || ldap:///cn=*,ou=X.500 Administrators,ou=People,dc=example,dc=com || ldap:///cn=*,ou=Executives,ou=People,dc=example,dc=com" )( targetscope = "subordinate" )( targattrfilters = "add=nsroleDN:(!(nsroledn=cn=X.500 Administrator)) && employeeStatus:(!(drink=beer)) && telephoneNumber:(telephoneNumber=612*)" )
}

func TestParseBindRule_postANTLR(t *testing.T) {
	var Br BindRule
	_ = Br.assertExpressionValue()

	tests := map[string]map[string]map[int][]string{
		`valid`: {
			`authmethod`: {
				0: {`simple`},
				1: {`none`},
				2: {`sasl`},
				3: {`ssl`},
				4: {`sasl DIGEST-MD5`},
				5: {`sasl GSSAPI`},
			},
			`ssf`: {
				0: {`0`},
				1: {`1`},
				2: {`50`},
				3: {`71`},
				4: {`128`},
				5: {`164`},
				6: {`256`},
			},
			`ip`: {
				0: {`192.168.1.100`},
				1: {`10.8.0`},
				2: {`fe80:19da:004a:1212::`},
				3: {`10.8.*`},
				4: {`10.8.`},
				5: {`172.16.5`},
				6: {`10`},
				7: {`10.1.*`},
				8: {`2001:47a:ee4::`},
				9: {`2001:47a:*`},
			},
			`dns`: {
				0: {`*.example.com`},
				1: {`www.example.com`},
				2: {`www.*.example.com`},
				3: {`www.example.*`},
			},
			`userattr`: {
				0: {`manager#SELFDN`},
				1: {`owner#GROUPDN`},
				2: {`parent[0,1,4].manager#LDAPURL`},
				3: {`parent[1,3].manager#uid=frank,ou=People,dc=example,dc=com`},
				4: {`parent[0].manager#SELFDN`},
				5: {`manager#LDAPURL`},
				6: {`parent[0,1,2,8].manager#USERDN`},
				7: {`ldap:///ou=People,dc=example,dc=com?manager#USERDN`},
			},
			`groupattr`: {
				0: {`manager#SELFDN`},
				1: {`owner#GROUPDN`},
				2: {`parent[0,1,4].manager#LDAPURL`},
				3: {`ldap:///ou=People,dc=example,dc=com?manager#USERDN`},
				4: {`parent[1,3].manager#uid=frank,ou=People,dc=example,dc=com`},
				5: {`parent[0].manager#SELFDN`},
				6: {`manager#LDAPURL`},
				7: {`parent[0,1,2,8].manager#USERDN`},
			},
			`userdn`: {
				0: {`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))`},
				1: {`ldap:///dc=example,dc=com??one?(&(objectClass=employee)(cn=*))`},
				2: {`uid=jesse,ou=People,dc=example,dc=com`},
			},
			`groupdn`: {
				0: {`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))`},
				1: {`ldap:///dc=example,dc=com??one?(&(objectClass=employee)(cn=*))`},
				2: {`uid=jesse,ou=People,dc=example,dc=com`},
			},
			`roledn`: {
				0: {`ldap:///cn=role,ou=Roles,dc=example,dc=com`},
			},
			`timeofday`: {
				0: {`1702`},
				1: {`0001`},
				2: {`0106`},
				3: {`1301`},
				4: {`2359`},
				5: {`0900`},
			},
			`dayofweek`: {
				0: {`Mon,Wed`},
				1: {`Tues,Wed`},
				2: {`Sun,Sat`},
				3: {`Fri`},
				4: {`Wed`},
				5: {`Sat,Sun`},
			},
		},

		`invalid`: {
			`authmethod`: {
				0: {`pimple`, `sample`},
				1: {`noone`},
				2: {`sizzle`},
				3: {`sslssllsll`},
				4: {`sasl INDIGESTION-MD5`},
				5: {`sasl GASSY`},
				6: {``},
			},
			`userdn`: {
				0: {``},
				1: {},
				2: {`           i   `},
			},
			`groupdn`: {
				0: {``},
			},
			`roledn`: {
				0: {``},
			},
			`userattr`: {
				0: {``},
				1: {},
				2: {`           i   `},
				3: {``, `ldap:///uid=jesse,ou=People,dc=example,dc=com`},
				4: {`parent[]].manager#USERDN`},
				5: {`parent[55555555]].manager#USERDN`},
			},
			`groupattr`: {
				0: {``},
				1: {},
				2: {`           i   `},
				3: {``, `ldap:///uid=jesse,ou=People,dc=example,dc=com`},
				4: {`parent[75].manager#`},
			},
			`timeofday`: {
				0: {`11702`},
				1: {`2500`},
				2: {`2401`},
				3: {},
				4: {`A`, ``},
				5: {``},
				6: {`A`},
				7: {`:)`},
			},
			`ssf`: {
				0: {`-1`},
				1: {`farts`},
				2: {`^`},
				3: {`257`},
				4: {`512`},
				5: {},
				6: {`100-`},
				7: {`a`},
				8: {``},
				9: {`X`, `:)`},
			},
			`dayofweek`: {
				0: {``},
				1: {`Toes`},
				2: {`Sub,Sad`, `banana`},
				3: {`Fry`},
				4: {`          `},
				5: {},
				6: {`Wad`},
				7: {`-14`},
				8: {`8`},
				9: {`Sal,Sum`},
			},
			`ip`: {
				0: {``},
				1: {`10?8.0`},
				2: {`.`},
				3: {`??`, `10.1.9/24`},
				4: {`@10.8.`},
				5: {`(172.16.5`},
				6: {`z748,`},
				7: {`10.a.*`},
				8: {},
			},
			`dns`: {
				0: {`___`},
				1: {`%al;`},
				2: {`][`},
				3: {`..*`},
				4: {`192.168.*`, `10/8`},
				5: {},
			},
		},
	}

	for typ, kwtests := range tests {
		for kw, typtests := range kwtests {
			for idx, value := range typtests {
				var br BindRule
				br.Init()
				br.SetKeyword(kw)
				for _, cop := range []ComparisonOperator{
					Eq, Ne,
				} {
					expr := makeParserRuleExpr(1, value...)
					br.SetOperator(cop).SetExpression(expr)
					err := br.assertExpressionValue()

					if err != nil && typ == `valid` {
						t.Errorf("%s [%s;%s::%d (%s)] failed: %v [%v]",
							t.Name(), kw, typ, idx, cop, err, expr)
						return

					} else if err == nil && typ == `invalid` {
						t.Errorf("%s [%s;%s::%d (%s)] failed: no error for bogus value [%v]",
							t.Name(), kw, typ, idx, cop, expr)
						return
					}
				}
			}
		}
	}
}

func TestParseTargetRule_postANTLR_codecov(t *testing.T) {
	_, _ = ParseTargetRule(``)
	want := `( target_to = "ldap:///ou=People,dc=example,dc=com" )`

	var _t TargetRule
	var err error
	_, _ = ParseBindRule(``)
	_ = _t.Parse(``)

	if _t, err = ParseTargetRule(want); err != nil {
		return
	}

	if want != _t.String() {
		t.Errorf("%s failed:\nwant '%s'\ngot '%s'", t.Name(), want, _t)
		return
	}

	if err = _t.Parse(want); err != nil {
		return
	}

	if want != _t.String() {
		t.Errorf("%s failed:\nwant '%s'\ngot '%s'", t.Name(), want, _t)
		return
	}
}

func TestParseTargetRule_postANTLR_extended(t *testing.T) {

	tests := map[string]map[string]map[int][]string{
		`valid`: {
			`targetscope`: {
				0: {`bAse`},
				1: {`one`},
				2: {`sUBtree`},
				3: {`sUBordinaTE`},
				4: {`sub`},
				5: {`onelevel`},
			},
			`targetattr`: {
				0: {`cn`, `sn`, `givenName`, `objectClass`, `uidNumber`, `gidNumber`, `uid`, `homeDirectory`, `gecos`, `description`, `loginShell`},
				1: {`aci`},
				2: {`cn;lang-jp`},
				3: {`objectclass`},
				4: {`objectClass`},
			},
			`targetfilter`: {
				0: {`(&(objectClass=employee)(cn=Jesse Coretta))`},
				1: {`(objectClass=account)`},
				2: {`(&(objectClass=accounting)(terminated=FALSE))`},
			},
			`targattrfilters`: {
				0: {`add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) &&
					gecos:(|(objectClass=contractor)(objectClass=intern))`},
				1: {`add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) &&
						gecos:(|(objectClass=contractor)(objectClass=intern)),
					delete=uidNumber:(&(objectClass=accounting)(terminated=FALSE)) 	    &&
						gidNumber:(objectClass=account)`},
				2: {`delete=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))`},

				// ASCII #59 delim (semi)
				3: {`add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) &&
						gecos:(|(objectClass=contractor)(objectClass=intern));
					delete=uidNumber:(&(objectClass=accounting)(terminated=FALSE)) 	    &&
						gidNumber:(objectClass=account)`},
			},
			`targetcontrol`: {
				0: {`1.3.6.1.4.1.56521.999.5`, `1.3.6.1.4.1.56521.999.100.1`},
				1: {`1.3.6.1.4.1.56521.999.5`},
			},
			`extop`: {
				0: {`1.3.6.1.4.1.56521.999.5`, `1.3.6.1.4.1.56521.999.100.1`},
				1: {`1.3.6.1.4.1.56521.999.5`},
			},
			`target`: {
				0: {`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))`},
				1: {`ldap:///cn=*,ou=People,dc=example,dc=com`, `ldap:///self`},
				2: {`uid=jesse,ou=People,dc=example,dc=com`, `uid=jimmy,ou=People,dc=example,dc=com`},
				3: {`uid=courtney,ou=People,dc=example,dc=com`},
				4: {`ldap:///ou=People,dc=example,dc=com??one?(&(objectClass=employee)(terminated=FALSE))`},
				5: {`ldap:///anyone`, `ldap:///parent`, `ldap:///all`},
			},
			`target_to`: {
				0: {`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))`},
				1: {`ldap:///cn=*,ou=People,dc=example,dc=com`, `ldap:///self`},
				2: {`uid=jesse,ou=People,dc=example,dc=com`, `uid=jimmy,ou=People,dc=example,dc=com`},
				3: {`uid=courtney,ou=People,dc=example,dc=com`},
				4: {`ldap:///ou=People,dc=example,dc=com??one?(&(objectClass=employee)(terminated=FALSE))`},
				5: {`ldap:///anyone`, `ldap:///parent`, `ldap:///all`},
			},
			`target_from`: {
				0: {`ldap:///ou=People,dc=example,dc=com?cn,sn,givenName,objectClass,uid?one?(&(objectClass=employee)(terminated=FALSE))`},
				1: {`ldap:///cn=*,ou=People,dc=example,dc=com`, `ldap:///self`},
				2: {`uid=jesse,ou=People,dc=example,dc=com`, `uid=jimmy,ou=People,dc=example,dc=com`},
				3: {`uid=courtney,ou=People,dc=example,dc=com`},
				4: {`ldap:///ou=People,dc=example,dc=com??one?(&(objectClass=employee)(terminated=FALSE))`},
				5: {`ldap:///anyone`, `ldap:///parent`, `ldap:///all`},
			},
		},

		`invalid`: {
			`target`: {
				0: {``, ``},
				1: {``},
				2: {},
			},
			`targetfilter`: {
				0: {``, ``},
				1: {``},
				2: {},
			},
			`targattrfilters`: {
				0: {``, ``},
				1: {``},
				2: {},
				3: {`cn?(objectClass=*)`},
			},
			`targetscope`: {
				0: {``, ``},
				1: {``},
				2: {},
				3: {`subtrap`},
				4: {`1level`},
			},
			`targetcontrol`: {
				0: {``, ``},
				1: {``},
				2: {},
				3: {`8.b.0.d`},
			},
			`extop`: {
				0: {``, ``},
				1: {``},
				2: {},
				3: {`8.b.0.d`},
			},
			`targetattr`: {
				0: {``, ``},
				1: {``},
				2: {},
				3: {`8cn`, ``, `object Class`, `Uid`},
			},
			`target_to`: {
				0: {``, ``},
				1: {``},
				2: {},
			},
			`target_from`: {
				0: {``, ``},
				1: {``},
				2: {},
			},
		},
	}

	for typ, kwtests := range tests {
		for kw, typtests := range kwtests {
			for idx, value := range typtests {
				var tr TargetRule
				tr.Init()
				tr.SetKeyword(kw)
				_ = tr.assertExpressionValue() // codecov
				// TODO: replace with TRM
				for _, cop := range []ComparisonOperator{
					Eq, Ne,
				} {
					expr := makeParserRuleExpr(1, value...)
					tr.SetOperator(cop).SetExpression(expr)
					err := tr.assertExpressionValue()

					if err != nil && typ == `valid` {
						t.Errorf("%s [%s;%s::%d (%s)] failed: %v [%v]",
							t.Name(), kw, typ, idx, cop, err, expr)
						return

					} else if err == nil && typ == `invalid` {
						t.Errorf("%s [%s;%s::%d (%s)] failed: no error for bogus value [%v]",
							t.Name(), kw, typ, idx, cop, expr)
						return
					}
				}
			}
		}
	}
}

func TestParseTargetRules(t *testing.T) {
	//var trs TargetRules
	_, _ = ParseTargetRules(``) // codecov
	bad_single := `target = "ldap:///uid=jesse,ou=People,dc=example,dc=com"`
	single := `( target = "ldap:///uid=jesse,ou=People,dc=example,dc=com" )`
	var b TargetRules
	var err error
	if err = b.Parse(bad_single); err == nil {
		t.Errorf("%s failed: no error where one was expected",
			t.Name())
		return
	}

	if err = b.Parse(single); err != nil {
		t.Errorf("%s failed: %v", t.Name(), err)
		return
	}

	if b.String() != single {
		t.Errorf("%s failed:\nwant '%s',\ngot  '%s'", t.Name(), single, b)
		return
	}
}

func TestParsePermission(t *testing.T) {
	var raw string = `allow(read,write,compare,selfwrite)`

	var perm Permission
	if err := perm.Parse(raw); err != nil {
		t.Errorf("%s failed: %v",
			t.Name(), illegalSyntaxPerTypeErr(perm, nil))
		return
	}

	if !perm.Positive(ReadAccess) {
		t.Errorf("%s failed; could not parse raw privileges (%s) into valid %T",
			t.Name(), raw, perm)
		return
	}

	if perm.String() != raw {
		t.Errorf("%s failed; bad result: want '%s', got '%s'",
			t.Name(), raw, perm)
		return
	}
}

func TestParsePermissionBindRule(t *testing.T) {
	var pbr PermissionBindRule
	_ = pbr.IsZero()
	_ = pbr.Valid()

	if err := pbr.Parse(``); err == nil {
		t.Errorf("%s bogus %T attempt returned no error",
			t.Name(), pbr)
		return
	}

	var privs string = `allow(read,write,search,compare)`
	var rules string = `( ( timeofday >= "0900" AND timeofday < "1830" ) AND ( dayofweek = "Mon,Tues,Wed,Thur,Fri" ) )`
	var raw string = sprintf("%s %s;", privs, rules)

	if err := pbr.Parse(raw); err != nil {
		t.Errorf("%s failed: %v",
			t.Name(), illegalSyntaxPerTypeErr(pbr, nil))
		return
	}

	if pbr.String() != raw {
		t.Errorf("%s failed; bad result: want '%s', got '%s'",
			t.Name(), raw, pbr)
		return
	}
}

func TestParsePermissionBindRules(t *testing.T) {
	var pbrs PermissionBindRules
	_ = pbrs.IsZero()
	_ = pbrs.Valid()

	if err := pbrs.Parse(``); err == nil {
		t.Errorf("%s bogus %T attempt returned no error",
			t.Name(), pbrs)
		return
	}

	// yeah
	var raw string = `allow(read,write) ( groupdn = "ldap:///cn=Human Resources,dc=example,dc=com" ); allow(read,write,delete,search,compare) ( userdn = "ldap:///all" ); deny(all) ( userdn = "ldap:///anyone" ) AND ( ip != "192.0.2." ); deny(proxy,export) ( userdn = "ldap:///self" ); deny(all,proxy) ( userdn = "ldap:///uid=user,ou=People,dc=example,dc=com" ); allow(read,search,compare,selfwrite) ( userdn = "ldap:///uid=user,ou=People,dc=example,dc=com" ) AND ( timeofday >= "1800" AND timeofday < "2400" ); deny(write,selfwrite,import) groupdn = "ldap:///cn=DomainAdmins,ou=Groups,[$dn],dc=example,dc=com"; allow(all,proxy) groupdn = "ldap:///cn=example,ou=groups,dc=example,dc=com"; deny(none) userattr = "owner#USERDN"; allow(none) userattr = "parent[1].manager#USERDN"; deny(all,proxy) userdn = "ldap:///anyone" || "ldap:///self" || "ldap:///cn=Admin"; deny(add,delete,selfwrite) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" ) ); deny(delete,compare,export) ( authmethod = "NONE" OR authmethod = "SIMPLE" ); allow(write,compare) groupdn = "ldap:///cn=Administrators,ou=Groups,dc=example,com" AND groupdn = "ldap:///cn=Operators,ou=Groups,dc=example,com"; allow(write,search) userattr = "manager#USERDN"; allow(proxy,import,export) userdn = "ldap:///anyone" AND ssf >= "128" AND NOT dayofweek = "Fri"; allow(search,compare) userdn = "ldap:///cn=Courtney Tolana,dc=example,dc=com"; deny(read) userdn = "ldap:///ou=People,dc=example,dc=com??sub?(department=Human Resources)"; deny(write,import) ( userdn = "ldap:///anyone" ) AND ( dns != "client.example.com" ); allow(selfwrite,proxy) ( userdn = "ldap:///anyone" ) AND NOT ( dns != "client.example.com" );`

	if err := pbrs.Parse(raw); err != nil {
		t.Errorf("%s failed: %v",
			t.Name(), illegalSyntaxPerTypeErr(pbrs, nil))
		return
	}

	if pbrs.String() != raw {
		t.Errorf("%s failed; bad result:\nwant '%s'\ngot  '%s'",
			t.Name(), raw, pbrs)
		return
	}
}

/*
This example demonstrates the acts of parsing a sequence of multiple PermissionBindRule
expressive statements. Each individual PermissionBindRule must be valid unto itself, and
in particular, must be terminated with ASCII #59 (;).
*/
func ExamplePermissionBindRules_Parse() {
	var pbrs PermissionBindRules
	// this is a sequence of three (3) PermissionBindRule
	// expressions in raw text format.
	var raw string = `allow(read,write) ( groupdn = "ldap:///cn=Human Resources,dc=example,dc=com" ); allow(read,write,delete,search,compare) ( userdn = "ldap:///all" ); deny(all) ( userdn = "ldap:///anyone" ) AND ( ip != "192.0.2." );`

	if err := pbrs.Parse(raw); err != nil {
		fmt.Println(err) // always check your parser errors
		return
	}

	fmt.Printf("%T instance contains %d slices", pbrs, pbrs.Len())
	// Output: aci.PermissionBindRules instance contains 3 slices
}

func ExamplePermission_Parse_granting() {
	var perm Permission

	raw := `allow(read,write,compare,selfwrite)`

	if err := perm.Parse(raw); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Read privileges granted: %t", perm.Positive(ReadAccess) && perm.Disposition() == `allow`)
	// Output: Read privileges granted: true

}

func ExamplePermission_Parse_withholding() {
	var perm Permission

	raw := `deny(all,proxy)`

	if err := perm.Parse(raw); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Proxy privileges denied: %t", perm.Positive(ProxyAccess) && perm.Disposition() == `deny`)
	// Output: Proxy privileges denied: true

}

/*
This example demonstrates the complete parsing of a composite ACIv3
component: the PermissionBindRule. A PermissionBindRule is a single
permission statement followed by a BindRule or BindRules statement,
and terminated by a semicolon (ASCII #59).
*/
func ExamplePermissionBindRule_Parse() {
	var privs string = `allow(read,write,search,compare)`
	var rules string = `( ( timeofday >= "0900" AND timeofday < "1830" ) AND ( dayofweek = "Mon,Tues,Wed,Thur,Fri" ) )`

	// combine the above into one single statement
	// with proper termination.
	var raw string = sprintf("%s %s;", privs, rules)

	// Prepare the 'container' for our new
	// PermissionBindRule components.
	var pbr PermissionBindRule
	if err := pbr.Parse(raw); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", pbr)
	// Output: allow(read,write,search,compare) ( ( timeofday >= "0900" AND timeofday < "1830" ) AND ( dayofweek = "Mon,Tues,Wed,Thur,Fri" ) );
}

/*
This example demonstrates a basic assembly of a PermissionBindRule using
actual Permission and BindContext instances assigned through the Set
method.
*/
func ExamplePermissionBindRule_Set() {
	// Prepare the 'container' for our new
	// PermissionBindRule components.
	var pbr PermissionBindRule
	pbr.Set(
		Allow(NoAccess),
		UDN(`ldap:///uid=disgruntled_employee,ou=People,dc=example,dc=com`).Eq(),
	)

	fmt.Printf("%s", pbr)
	// Output: allow(none) userdn = "ldap:///uid=disgruntled_employee,ou=People,dc=example,dc=com";
}

/*
This example demonstrates the same outcome as the PermissionBindRule.Parse
example, except this time using the Set method on a nil instance.
*/
func ExamplePermissionBindRule_Set_withParse() {
	var privs string = `allow(read,write,search,compare)`
	var rules string = `( ( timeofday >= "0900" AND timeofday < "1830" ) AND ( dayofweek = "Mon,Tues,Wed,Thur,Fri" ) )`

	// combine the above into one single statement
	// with proper termination.
	var raw string = sprintf("%s %s;", privs, rules)

	// Prepare the 'container' for our new
	// PermissionBindRule components.
	var pbr PermissionBindRule
	pbr.Set(raw)

	fmt.Printf("%s", pbr)
	// Output: allow(read,write,search,compare) ( ( timeofday >= "0900" AND timeofday < "1830" ) AND ( dayofweek = "Mon,Tues,Wed,Thur,Fri" ) );
}

/*
This example demonstrates a basic parse of an ACIv3 instruction in string representation into
a proper instance of Instruction using the Parse method.
*/
func ExampleInstruction_Parse() {
	raw := `( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" ); )`

	// define a variable into which the
	// parser shall deposit data
	var ins Instruction
	if err := ins.Parse(raw); err != nil {
		fmt.Println(err) // always check your parser errors
		return
	}

	fmt.Printf("%s", ins)
	// Output: ( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( timeofday >= "1730" AND timeofday < "2400" );)
}

/*
This example demonstrates the parsing of a single BindRule condition.

A textual Bind Keyword, an appropriate ComparisonOperator as well as
an appropriate expression value are submitted through the ParseBindRule
package-level function, which in turn calls a similarly named antlraci
function through which parsing is delegated.
*/
func ExampleParseBindRule() {
	raw := `( userdn = "ldap:///cn=Jesse Coretta,ou=People,dc=example,dc=com" || "ldap:///cn=Courtney Tolana,ou=People,dc=example,dc=com" )`
	br, err := ParseBindRule(raw)
	if err != nil {
		fmt.Println(err) // always check your parser errors
		return
	}
	fmt.Printf("%T is parenthetical: %t", br, br.IsParen())
	// Output: aci.BindRule is parenthetical: true
}

/*
This example demonstrates the parsing of a TargetRules expressive statement
containing multiple TargetRule conditions.
*/
func ExampleTargetRules_Parse() {
	raw := `( targetscope="base" )(targetfilter="(&(objectClass=employee)(status=terminated))")(targetattr != "aci")"`
	var trs TargetRules
	if err := trs.Parse(raw); err != nil {
		fmt.Println(err) // always check your parser errors
		return
	}

	fmt.Printf("%s", trs.Index(1))
	// Output: ( targetfilter = "(&(objectClass=employee)(status=terminated))" )
}

/*
This example demonstrates the parsing of a single TargetRule condition.
*/
func ExampleTargetRule_Parse() {
	raw := `(targetattr != "aci")"`
	var tr TargetRule
	if err := tr.Parse(raw); err != nil {
		fmt.Println(err) // always check your parser errors
		return
	}

	fmt.Printf("%s", tr.Expression())
	// Output: aci
}

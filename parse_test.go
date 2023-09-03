package aci

import (
	"fmt"
	"testing"
)

func TestParseBindRule(t *testing.T) {
	want := `userdn = "ldap:///cn=Jesse Coretta,ou=People,dc=example,dc=com" || "ldap:///anyone"`

	var b BindRule
	var err error
	_, _ = ParseBindRule(``)

	if b, err = ParseBindRule(want); err != nil {
		return
	}

	if want != b.String() {
		t.Errorf("%s failed:\nwant '%s'\ngot '%s'", t.Name(), want, b)
	}
}

func TestParseBindRules(t *testing.T) {
	want := `( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) )`

	var r BindRules
	var err error

	_, _ = ParseBindRules(``)

	if r, err = ParseBindRules(want); err != nil {
		return
	}

	if want != r.String() {
		t.Errorf("%s failed:\nwant '%s',\ngot  '%s'", t.Name(), want, r)
	}

	if r.Keyword() == nil {
		t.Errorf("%s failed: unidentified %T", t.Name(), r.Keyword())
	}

	if !r.IsNesting() {
		t.Errorf("%s failed: nesting not detected", t.Name())
	}

	bl := r.Len()
	orig := r.String()

	r.Push(BindRules{})

	var ctx BindContext = BindRule{}

	if r.Push(ctx); r.Len() != bl {
		t.Errorf("%s failed: bogus enveloped content was pushed into %T", t.Name(), r)
	}

	popped := r.Pop()
	bl = r.Len()
	if popped.String() != orig {
		t.Errorf("%s failed: unexpected element popped; want '%s', got '%s'", t.Name(), orig, popped)
	}

	r.Push(popped)
	r.remove(r.Len() - 1)
	if r.Len() != bl {
		t.Errorf("%s failed: content not removed from %T", t.Name(), r)
	}

	r.insert(popped, 0)
	if r.Len() == bl {
		t.Errorf("%s failed: content not inserted into %T", t.Name(), r)
	}
}

func ExampleParseBindRules_messy() {
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
		fmt.Println(err)
		return
	}

	called := br.Traverse(0, 0, 0)
	fmt.Printf("%s", called)
	// Output: ( userdn = "ldap:///anyone" )
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
				br := new(BindRule).SetKeyword(kw)
				for _, cop := range []ComparisonOperator{
					Eq, Ne,
				} {
					expr := makeParserRuleExpr(1, value...)
					br.SetOperator(cop).SetExpression(expr)
					err := br.assertExpressionValue()

					if err != nil && typ == `valid` {
						t.Errorf("%s [%s;%s::%d (%s)] failed: %v [%v]",
							t.Name(), kw, typ, idx, cop, err, expr)

					} else if err == nil && typ == `invalid` {
						t.Errorf("%s [%s;%s::%d (%s)] failed: no error for bogus value [%v]",
							t.Name(), kw, typ, idx, cop, expr)
					}
				}
			}
		}
	}
}

func TestParseTargetRule_postANTLR(t *testing.T) {
	_, _ = ParseTargetRule(``) // codecov

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
				var tr *TargetRule
				tr = new(TargetRule).SetKeyword(kw)
				_ = tr.assertExpressionValue() // codecov
				for _, cop := range []ComparisonOperator{
					Eq, Ne,
				} {
					expr := makeParserRuleExpr(1, value...)
					tr.SetOperator(cop).SetExpression(expr)
					err := tr.assertExpressionValue()

					if err != nil && typ == `valid` {
						t.Errorf("%s [%s;%s::%d (%s)] failed: %v [%v]",
							t.Name(), kw, typ, idx, cop, err, expr)

					} else if err == nil && typ == `invalid` {
						t.Errorf("%s [%s;%s::%d (%s)] failed: no error for bogus value [%v]",
							t.Name(), kw, typ, idx, cop, expr)
					}
				}
			}
		}
	}
}

func TestParseTargetRules(t *testing.T) {
	//var trs TargetRules
	_, _ = ParseTargetRule(``) // codecov
}

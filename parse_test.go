package aci

import (
	"fmt"
	"testing"
)

func TestParseBindRule(t *testing.T) {
	want := `userdn = "ldap:///cn=Jesse Coretta,ou=People,dc=example,dc=com" || "ldap:///anyone"`

	b, err := ParseBindRule(want)
	if err != nil {
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

/*
// use for something else ...
func ExampleParseTargetRules_Contains() {

        omg1 := `(
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

        tr1, err := ParseTargetRules(omg1)
        if err != nil {
                fmt.Println(err) // always check your parser errors.
                return
        }

        fmt.Printf("%t", tr1.Contains(TargetTo))
        // Output: true
}
*/

func TestParseBindRule_postANTLR(t *testing.T) {
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
			},
			`groupattr`: {
				0: {`manager#SELFDN`},
				1: {`owner#GROUPDN`},
				2: {`parent[0,1,4].manager#LDAPURL`},
				3: {`parent[1,3].manager#uid=frank,ou=People,dc=example,dc=com`},
				4: {`parent[0].manager#SELFDN`},
				5: {`manager#LDAPURL`},
				6: {`parent[0,1,2,8].manager#USERDN`},
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
			/*
						`targetfilter`: map[int][]string{
							0: []string{ `(&(objectClass=employee)(cn=Jesse Coretta))` },
							1: []string{ `(objectClass=account)` },
							2: []string{ `(&(objectClass=accounting)(terminated=FALSE))` },
						},
						`targattrfilters`: map[int][]string{
			                                0: []string{ `add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern))` },
							1: []string{ `add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern)),delete=uidNumber:(&(objectClass=accounting)(terminated=FALSE)) && gidNumber:(objectClass=account)` },
							2: []string{ `delete=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))` },
						},
			*/
		},

		`invalid`: {
			`authmethod`: {
				0: {`pimple`},
				1: {`noone`},
				2: {`sizzle`},
				3: {`sslssllsll`},
				4: {`sasl INDIGESTION-MD5`},
				5: {`sasl GASSY`},
				6: {``},
			},
			`userdn`: {
				0: {``},
				1: {`           i   `},
			},
			`groupdn`: {
				0: {``},
			},
			`roledn`: {
				0: {``},
			},
			`timeofday`: {
				0: {`11702`},
				1: {`2500`},
				2: {`2401`},
				3: {``},
				4: {`A`},
				5: {`:)`},
			},
			`ssf`: {
				0: {`-1`},
				1: {`farts`},
				2: {`^`},
				3: {`257`},
				4: {`512`},
				5: {`100-`},
				6: {`a`},
				7: {``},
			},
			`dayofweek`: {
				0: {``},
				1: {`Toes`},
				2: {`Sub,Sad`},
				3: {`Fry`},
				4: {`          `},
				5: {`Wad`},
				6: {`-14`},
				7: {`8`},
				8: {`Sal,Sum`},
			},
			`ip`: {
				0: {``},
				1: {`10?8.0`},
				2: {`.`},
				3: {`??`},
				4: {`@10.8.`},
				5: {`(172.16.5`},
				6: {`z748,`},
				7: {`10.a.*`},
			},
			`dns`: {
				0: {``},
				1: {`%al;`},
				2: {`][`},
				3: {`..*`},
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

func TestParseBindRule_postANTLRold(t *testing.T) {
	for k, v := range map[int]*BindRule{
		0: new(BindRule).SetKeyword(`ssf`).SetOperator(Lt).SetExpression(makeParserRuleExpr(1, []string{
			`175`,
		}...)),
		1: new(BindRule).SetKeyword(`userdn`).SetOperator(Ne).SetExpression(makeParserRuleExpr(0, []string{
			`ldap:///cn=Jesse Coretta,ou=Contractors,ou=People,dc=example,dc=com`,
			`ldap:///cn=Courtney Tolana,ou=Contractors,ou=People,dc=example,dc=com`,
			`ldap:///cn=Dr. Doctor Steve Brule,ou=Contractors,ou=People,dc=example,dc=com`,
		}...)),
		2: new(BindRule).SetKeyword(`groupdn`).SetOperator(Ne).SetExpression(makeParserRuleExpr(1, []string{
			`ldap:///cn=Executives,ou=Groups,dc=example,dc=com`,
			`ldap:///cn=Engineering,ou=Groups,dc=example,dc=com`,
			`ldap:///cn=Payroll,ou=Groups,dc=example,dc=com`,
			`ldap:///cn=Research and Development,ou=Engineering,ou=Groups,dc=example,dc=com`,
		}...)),
		3: new(BindRule).SetKeyword(`dns`).SetOperator(Eq).SetExpression(makeParserRuleExpr(1, []string{
			`*.example.com`,
		}...)),
		4: new(BindRule).SetKeyword(`ip`).SetOperator(Eq).SetExpression(makeParserRuleExpr(1, []string{
			`192.168.,10.,172.16.`,
		}...)),
		5: new(BindRule).SetKeyword(`authmethod`).SetOperator(Ne).SetExpression(makeParserRuleExpr(1, []string{
			`simple`,
		}...)),
		6: new(BindRule).SetKeyword(`timeofday`).SetOperator(Le).SetExpression(makeParserRuleExpr(1, []string{
			`1701`,
		}...)),
		7: new(BindRule).SetKeyword(`dayofweek`).SetOperator(Eq).SetExpression(makeParserRuleExpr(1, []string{
			`Wed,Fri,Sat`,
		}...)),
		8: new(BindRule).SetKeyword(`userattr`).SetOperator(Eq).SetExpression(makeParserRuleExpr(1, []string{
			`manager#SELFDN`,
		}...)),
		9: new(BindRule).SetKeyword(`groupattr`).SetOperator(Ne).SetExpression(makeParserRuleExpr(1, []string{
			`owner#GROUPDN`,
		}...)),
	} {
		if err := v.assertExpressionValue(); err != nil {
			t.Errorf("%s [%d] failed: %v", t.Name(), k, err)
		}
	}
}

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

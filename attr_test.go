package aci

import (
	"fmt"
	"testing"
)

/*
This example demonstrates the creation of an instance of AttributeType, which
is used in a variety of contexts.

In this example, a string name is fed to the package level AT function to form
a complete AttributeType instance, which is then shown in string representation.
*/
func ExampleAT() {
	atb := AT(`homeDirectory`)
	fmt.Printf("%s", atb)
	// Output: homeDirectory
}

/*
This example demonstrates the string representation of the receiver instance.
*/
func ExampleAttributeType_String() {
	fmt.Printf("%s", AT(`owner`))
	// Output: owner
}

/*
This example demonstrates the SHA-1 hash comparison between two (2) AttributeType
instances using the Compare method.
*/
func ExampleAttributeType_Compare() {
	attr := AT(`cACertificate`)
	oattr := AT(`cacertificate`)

	fmt.Printf("Hashes are equal: %t", oattr.Compare(attr))
	// Output: Hashes are equal: false
}

/*
This example demonstrates the use of the useless Keyword method, as AttributeType
instances do not have any knowledge of Keywords at this time.
*/
func ExampleAttributeType_Keyword() {
	fmt.Printf("%v", AT(`owner`).Keyword())
	// Output: <nil>
}

/*
This example demonstrates the use of the useless Kind method, as this information
is normally derived from a Keyword, which the receiver does not have.
*/
func ExampleAttributeType_Kind() {
	fmt.Printf("'%s'", AT(`owner`).Kind())
	// Output: ''
}

/*
This example demonstrates the use of the useless Len method, as this information
is only made available to satisfy Go's interface signature requirements as they
pertain to the AttributeTypeContext interface.
*/
func ExampleAttributeType_Len() {
	fmt.Printf("%d", AT(`owner`).Len())
	// Output: 1
}

/*
This example demonstrates a check of the receiver for "nilness".
*/
func ExampleAttributeType_IsZero() {
	fmt.Printf("%t", AT(`owner`).IsZero())
	// Output: false
}

/*
This example demonstrates a check of the receiver for an aberrant state.
*/
func ExampleAttributeType_Valid() {
	fmt.Printf("Valid: %t", AT(`owner`).Valid() == nil)
	// Output: Valid: true
}

/*
This example demonstrates how a caller can determine the number of comparison
operator methods are available for use by the receiver instance.
*/
func ExampleAttributeType_TRM() {
	var at AttributeType
	fmt.Printf("%d available aci.BindRuleMethod instances", at.TRM().Len())
	// Output: 2 available aci.BindRuleMethod instances
}

/*
This example demonstrates the creation of an equality TargetRule (targetattr)
using the receiver instance as input:
*/
func ExampleAttributeType_Eq() {
	attr := AT(`*`)
	fmt.Printf("%s", attr.Eq())
	// Output: ( targetattr = "*" )
}

/*
This example demonstrates the creation of a negated equality TargetRule
(targetattr) using the receiver instance as input:
*/
func ExampleAttributeType_Ne() {
	attr := AT(`aci`)
	fmt.Printf("%s", attr.Ne())
	// Output: ( targetattr != "aci" )
}

func ExampleABTV() {
	var atb AttributeBindTypeOrValue = ABTV(BindUAT)
	atb.Set(AT(`owner`), GROUPDN)
	fmt.Printf("%T: %s", atb, atb)
	// Output: aci.AttributeBindTypeOrValue: owner#GROUPDN
}

/*
This example demonstrates the creation of an instance of AttributeBindTypeOrValue.
*/
func ExampleUAT() {
	attr := AT(`manager`)
	btype := USERDN

	atb := UAT(attr, btype)

	fmt.Printf("%T: %s", atb, atb)
	// Output: aci.AttributeBindTypeOrValue: manager#USERDN
}

/*
This example demonstrates the creation of an instance of AttributeBindTypeOrValue.
*/
func ExampleGAT() {
	attr := AT(`owner`)
	btype := SELFDN

	atb := GAT(attr, btype)

	fmt.Printf("%T: %s", atb, atb)
	// Output: aci.AttributeBindTypeOrValue: owner#SELFDN
}

/*
This example demonstrates the creation of an instance of AttributeBindTypeOrValue.

In this example, a raw string representation of an AttributeBindTypeOrValue instance
is used for parser input.
*/
func ExampleAttributeBindTypeOrValue_Parse() {
	txt := `manager#USERDN`
	var atb AttributeBindTypeOrValue
	err := atb.Parse(txt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%T: %s", atb, atb)
	// Output: aci.AttributeBindTypeOrValue: manager#USERDN
}

/*
This example demonstrates the creation of an instance of AttributeBindTypeOrValue followed
by a call of its AttributeType method.

The return type, AttributeType, is shown in string representation.
*/
/*
func ExampleAttributeBindTypeOrValue_AttributeType() {
	aftxt := `homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))`
	var atb AttributeBindTypeOrValue
	_ = atb.Parse(aftxt) // shadow error for brevity

	fmt.Printf("%s", atb.AttributeType())
	// Output: homeDirectory
}
*/

/*
This example demonstrates a check of the receiver for "nilness".
*/
func ExampleAttributeBindTypeOrValue_IsZero() {
	var atb AttributeBindTypeOrValue

	fmt.Printf("Is zero: %t", atb.IsZero())
	// Output: Is zero: true
}

/*
This example demonstrates the interrogation of the receiver in order to
discern the appropriate Keyword.

Its string representation, along with the name of the Keyword type, is
shown.
*/
func ExampleAttributeBindTypeOrValue_Keyword() {
	var atb AttributeBindTypeOrValue

	fmt.Printf("Keyword is '%s' (type:%T)", atb.Keyword(), atb.Keyword())
	// Output: Keyword is 'userattr' (type:aci.BindKeyword)
}

/*
This example demonstrates the creation of an instance of AttributeBindTypeOrValue followed
by a call of its String method.

The return value is the entirely of the receiver in string representation.
*/
func ExampleAttributeBindTypeOrValue_String() {
	var atb AttributeBindTypeOrValue
	atb.Set(AT(`manager`), USERDN)
	fmt.Printf("%s", atb)
	// Output: manager#USERDN
}

/*
This example demonstrates the creation of an instance of AttributeBindTypeOrValue followed
by a call of its String method through fmt.Printf.

The return value is the entirely of the receiver in string representation.
*/
func ExampleAttributeBindTypeOrValue_Set() {
	var atb AttributeBindTypeOrValue = ABTV(BindUAT)
	atb.Set(AT(`manager`), USERDN)
	fmt.Printf("%s value is %s", atb.Keyword(), atb)
	// Output: userattr value is manager#USERDN
}

/*
This example demonstrates the SHA-1 hash comparison between two (2) AttributeBindTypeOrValue
instances using the Compare method.
*/
func ExampleAttributeBindTypeOrValue_Compare() {
	// these will eval as true (match) because the underlying
	// string call does not include the Keyword. We need the
	// bind rule to see that ...
	attr := ABTV(BindUAT, `cACertificate`, `USERDN`)
	oattr := ABTV(BindGAT, `cACertificate`, USERDN)

	eqaBr := attr.Eq()
	eqoBr := oattr.Eq()

	fmt.Printf("Hashes are equal: %t", eqaBr.Compare(eqoBr))
	// Output: Hashes are equal: false
}

/*
This example demonstrates the creation of an instance of AttributeBindTypeOrValue followed
by a call of its String method through fmt.Printf. In this example, the receiver instance
is populated using only string values.

The return value is the entirely of the receiver in string representation.
*/
func ExampleAttributeBindTypeOrValue_Set_alt() {
	var atb AttributeBindTypeOrValue
	atb.Set(`manager`, `USERDN`)
	fmt.Printf("%s", atb)
	// Output: manager#USERDN
}

/*
This example demonstrates the creation of an instance of AttributeBindTypeOrValue followed
by a call of its Valid method for the purpose of sanity checking the receiver.

An error is reported and printed to STDOUT.
*/
func ExampleAttributeBindTypeOrValue_Valid() {
	var (
		atb AttributeBindTypeOrValue
		err error
	)

	if err = atb.Valid(); err != nil {
		fmt.Println(err)
	}
	// Output: aci.AttributeBindTypeOrValue instance is nil
}

/*
This example demonstrates the use of the AttributeBindTypeOrValue type's Eq method,
allowing for the creation of a TargetRule instance containing the receiver value,
and bearing the `targattrfilters` keyword context.
*/
func ExampleAttributeBindTypeOrValue_Eq() {
	attr := AT(`manager`)
	btype := USERDN

	atb := UAT(attr, btype)

	fmt.Printf("%s", atb.Eq())
	// Output: userattr = "manager#USERDN"
}

/*
This example demonstrates the use of the AttributeBindTypeOrValue type's Ne method
*/
func ExampleAttributeBindTypeOrValue_Ne() {
	attr := AT(`manager`)
	btype := USERDN

	atb := UAT(attr, btype)

	fmt.Printf("%s", atb.Ne())
	// Output: userattr != "manager#USERDN"
}

func ExampleAttributeBindTypeOrValue_BRM() {
	var atb AttributeBindTypeOrValue
	fmt.Printf("%d available aci.BindRuleMethod instances", atb.BRM().Len())
	// Output: 2 available aci.BindRuleMethod instances
}

func TestAttributeTypes(t *testing.T) {

	for keyword, atfn := range map[Keyword]func(...any) AttributeTypes{
		BindUDN:    UAs,
		TargetAttr: TAs,
	} {
		var attrs AttributeTypes = atfn()
		_ = attrs.Eq()
		_ = attrs.Ne()
		_ = attrs.Len()
		attrs.isAttributeTypeContext()
		attrs.reset()
		attrs.Push(keyword)
		attrs.resetKeyword(keyword)
		attrs.resetKeyword(keyword.String())
		attrs.setQuoteStyle(1)
		attrs.setQuoteStyle(0)
		attrs.Contains(3.14159)
		attrs.Push('𝝅')

		_ = attrs.Keyword()
		_ = attrs.Kind()
		_ = attrs.Valid()

		for _, raw := range []string{
			`cn`,
			`givenName`,
			`color;lang-fr`,
			`objectClass`,
			`drink`,
		} {
			var (
				al   int                        = attrs.Len()
				atn  func(string) AttributeType = attrs.F()
				attr AttributeType
				actx AttributeTypeContext
			)
			attr.isAttributeTypeContext()

			if err := testEmptyAttrContext(t, keyword, attr, attrs, al); err != nil {
				t.Errorf(err.Error())
			}

			// process attribute
			attr = atn(raw)
			al = attrs.Len()

			if attrs.Push(attr); !attrs.Contains(raw) {
				t.Errorf("%s [%s] multival failed: valid %T[%s] instance (%s) not pushed into %T[%s; len:%d]",
					t.Name(), keyword, attr, keyword, attr, attrs, attrs.Keyword(), al)
			}
			popped := attrs.Pop()
			attrs.Push(popped)
			attrs.Push(popped.String())
			attrs.Push(3.14159)
			attrs.Push('𝝅')

			attrs.setQuoteStyle(0)
			attrs.setQuoteStyle(1)

			actx = testMakeAttrContext(1, attr, attrs)
			trm := attrs.TRM()
			for i := 0; i < trm.Len(); i++ {
				cop, meth := trm.Index(i + 1)
				if meth == nil {
					t.Errorf("%s [%s] multival failed: expected %s method (%T), got nil",
						t.Name(), attrs.Keyword(), cop.Context(), meth)
				}

				wcop := sprintf("( %s %s %q )", attrs.Keyword(), cop, actx.String())
				if T := meth(); T.String() != wcop {
					err := unexpectedStringResult(actx.String(), wcop, T.String())
					t.Errorf("%s multival failed [%s rule]: %v",
						t.Name(), attrs.Keyword(), err)
				}
			}
		}
	}
}

func testMakeAttrContext(phase int, attr AttributeType, attrs AttributeTypes) (actx AttributeTypeContext) {
	if phase == 0 {
		actx = attr
		return
	}

	actx = attrs
	return
}

func testEmptyAttrContext(t *testing.T, kw Keyword, attr AttributeType, attrs AttributeTypes, ol int) (err error) {
	_ = attr.Kind()
	_ = attr.Valid()
	_ = attrs.Keyword()
	_ = attrs.Kind()
	_ = attrs.Valid()

	err = attr.Valid()
	if err != nil {
		if err.Error() != `aci.AttributeType instance is nil` {
			err = errorf("%s [%s] multival failed: invalid %T returned no validity error",
				t.Name(), kw, attr)
		} else {
			err = nil
		}
	} else {
		err = errorf("%s [%s] multival failed: invalid %T returned no validity error",
			t.Name(), kw, attr)
	}

	if attr.String() != badAT {
		err = errorf("%s [%s] multival failed: unexpected string result; want '%s', got '%s'",
			t.Name(), kw, badAT, attr)
	}

	if attrs.Push(attr); attrs.Len() > ol {
		err = errorf("%s [%s] multival failed (len): invalid %T (%s) pushed into %T without error",
			t.Name(), kw, attr, attr, attrs)
	}

	if attrs.Contains(attr) {
		err = errorf("%s [%s] multival failed (contains): invalid %T instance pushed into %T without error",
			t.Name(), kw, attr, attrs)
	}

	return
}

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

/*
This example demonstrates the creation of an AttributeTypes instance
suitable for use in the assembly of a TargetRule bearing the `targetattr`
keyword context.
*/
func ExampleAttributeTypes_targetAttributes() {

	// the TAs function allows AttributeType
	// instances in string representation, or
	// cast as proper AttributeType instances.
	//
	// Ordering is always preserved.
	attrs := TAs(
		`cn`,
		AT(`givenName`),
		`sn`,
		`objectClass`,
		AT(`l`),
	)

	fmt.Printf("%s", attrs.Eq())
	// Output: ( targetattr = "cn || givenName || sn || objectClass || l" )
}

/*
This example demonstrates the creation of an AttributeTypes
instance suitable for use in an LDAPURI instance.
*/
func ExampleAttributeTypes_uRIAttributes() {

	// the UAs function allows AttributeType
	// instances in string representation, or
	// cast as proper AttributeType instances.
	//
	//
	// Ordering is always preserved.
	attrs := UAs(
		`cn`,
		AT(`givenName`),
		`sn`,
		`objectClass`,
		AT(`l`),
	)

	fmt.Printf("%s", attrs)
	// Output: cn,givenName,sn,objectClass,l
}

func ExampleAttributeTypes_Len() {
	fmt.Printf("%d attributeTypes", TAs(`cn`, `sn`, `uid`).Len())
	// Output: 3 attributeTypes
}

func ExampleAttributeTypes_IsZero() {
	fmt.Printf("Empty stack: %t", TAs(`cn`, `sn`, `uid`).IsZero())
	// Output: Empty stack: false
}

func ExampleAttributeTypes_Valid() {
	fmt.Printf("Empty stack: %t", TAs(`cn`, `sn`, `uid`).Valid() == nil)
	// Output: Empty stack: true
}

func ExampleAttributeTypes_Contains() {
	attrs := TAs(
		`cn`,
		AT(`givenName`),
		`sn`,
		`objectClass`,
		AT(`l`),
	)

	fmt.Printf("Contains 'l': %t", attrs.Contains(`l`))
	// Output: Contains 'l': true
}

/*
This example demonstrates a basic SHA-1 hash comparison between two
like instances of the receiver's type.
*/
func ExampleAttributeTypes_Compare_likeInstances() {
	attrs1 := TAs(`cn`, `givenName`, `sn`, `objectClass`, `l`)
	attrs2 := TAs(`cn`, `givenName`, `sn`, `objectClass`, `l`)

	fmt.Printf("%T instances match: %t", attrs1, attrs1.Compare(attrs2))
	// Output: aci.AttributeTypes instances match: true
}

/*
This example demonstrates why two seemingly identical instances, though
created by separate functions, fail to evaluate as equal when multi-valued.
The reason for this is due to the nature of the String method behavior for
each of the instances. Use of the TAs package level function guarantees a
delimitation scheme using the symbolic OR (||) symbol, while use of the UAs
package level function uses comma-delimitation.
*/
func ExampleAttributeTypes_Compare_multiValueFalse() {
	attrs1 := TAs(`cn`, `givenName`, `sn`, `objectClass`, `l`) // <x> || <x> || ...
	attrs2 := UAs(`cn`, `givenName`, `sn`, `objectClass`, `l`) // x,x,x ...

	fmt.Printf("%T instances match: %t", attrs1, attrs1.Compare(attrs2))
	// Output: aci.AttributeTypes instances match: false
}

/*
This example demonstrates the contrary condition to that demonstrated in the
AttributeTypes CompareMultiValueFalse example. Because the two instances to
be evaluated are single-valued, no delimiter scheme comes into play. As such,
the two instances produce identical String output.
*/
func ExampleAttributeTypes_Compare_singleValueTrue() {
	attrs1 := TAs(`cn`)
	attrs2 := UAs(`cn`)

	fmt.Printf("%T instances match: %t", attrs1, attrs1.Compare(attrs2))
	// Output: aci.AttributeTypes instances match: true
}

func ExampleAttributeTypes_Eq() {
	attrs := TAs(
		`cn`,
		AT(`givenName`),
		`sn`,
		`objectClass`,
		AT(`l`),
	)

	fmt.Printf("%s", attrs.Eq())
	// Output: ( targetattr = "cn || givenName || sn || objectClass || l" )
}

func ExampleAttributeTypes_Ne() {
	fmt.Printf("%s", TAs(AT(`aci`)).Ne())
	// Output: ( targetattr != "aci" )
}

func ExampleAttributeTypes_F() {
	attrs := TAs(
		`l`,
		`cn`,
		`sn`,
		`givenName`,
		`objectClass`,
	)

	attr := attrs.F()(`homeDirectory`) // We've just executed 'AT' function without having to find it.
	fmt.Printf("attr is %T", attr)
	// Output: attr is aci.AttributeType
}

func ExampleAttributeTypes_Index() {
	attrs := TAs(
		`cn`,
		AT(`givenName`),
		`sn`,
		`objectClass`,
		AT(`l`),
	)

	fmt.Printf("%s", attrs.Index(2))
	// Output: sn
}

/*
This example demonstrates the addition of new slice elements
to an AttributeTypes instance using its Push method.
*/
func ExampleAttributeTypes_Push() {
	attrs := TAs(
		`cn`,
		AT(`givenName`),
		`sn`,
		`objectClass`,
	)
	attrs.Push(`l`)

	fmt.Printf("%s", attrs)
	// Output: cn || givenName || sn || objectClass || l
}

/*
This example demonstrates the string representation of the
receiver instance using its String method.
*/
func ExampleAttributeTypes_String() {
	attrs := TAs(
		`cn`,
		AT(`givenName`),
		`sn`,
		`objectClass`,
	)

	fmt.Printf("%s", attrs)
	// Output: cn || givenName || sn || objectClass
}

/*
This example demonstrates the use of the receiver's TRM
method in order to determine available comparison operator
driven methods available in this context.
*/
func ExampleAttributeTypes_TRM_targetAttributes() {
	attrs := TAs(
		`cn`,
		AT(`givenName`),
		`sn`,
		`objectClass`,
	)

	fmt.Printf("%d available comparison operator methods", attrs.TRM().Len())
	// Output: 2 available comparison operator methods
}

/*
This example demonstrates the useless nature of the receiver's
TRM method in situations where the receiver is intended for use
within an LDAPURI instance, rather than as a rule condition unto
itself. As a result, a bogus TargetRuleMethods instance will be
returned if the receiver was created through any means *OTHER
THAN* by execution of the TAs package level function.
*/
func ExampleAttributeTypes_TRM_uRIAttributes() {
	attrs := UAs(
		`cn`,
		AT(`givenName`),
		`sn`,
		`objectClass`,
	)

	fmt.Printf("%d available comparison operator methods", attrs.TRM().Len())
	// Output: 0 available comparison operator methods
}

/*
This example demonstrates the removal of a single slice element
from an AttributeTypes instance in LIFO fashion using its Pop
method.
*/
func ExampleAttributeTypes_Pop() {
	attrs := TAs(
		`cn`,
		AT(`givenName`),
		`sn`,
		`objectClass`,
		`l`,
	)
	popped := attrs.Pop()

	fmt.Printf("%s", popped)
	// Output: l
}

func ExampleAttributeTypes_Keyword_targetAttributes() {
	attrs := TAs(
		`l`,
		`cn`,
		`sn`,
		`givenName`,
		`objectClass`,
	)

	fmt.Printf("%s", attrs.Keyword())
	// Output: targetattr
}

func ExampleAttributeTypes_Kind_targetAttributes() {
	attrs := TAs(
		`l`,
		`cn`,
		`sn`,
		`givenName`,
		`objectClass`,
	)

	fmt.Printf("%s", attrs.Kind())
	// Output: targetattr
}

func ExampleAttributeTypes_Kind_uRIAttributes() {
	attrs := UAs(
		`l`,
		`cn`,
		`sn`,
		`givenName`,
		`objectClass`,
	)

	fmt.Printf("%s", attrs.Kind())
	// Output: <uri_search_attributes>
}

func ExampleAttributeTypes_Kind_uninitialized() {
	var attrs AttributeTypes
	fmt.Printf("%s", attrs.Kind())
	// Output: <uninitialized>
}

/*
This example demonstrates the incompatibility with an
AttributeTypes stack intended for use within an LDAPURI
instance. Because this incarnation of AttributeTypes has
no direct application within a BindRule or a TargetRule
(rather it resides WITHIN another type eligible for such
use),  the Keyword cannot be inferred and returns a nil
instance. This is expected behavior.
*/
func ExampleAttributeTypes_Keyword_uRIAttributes() {
	attrs := UAs(
		`l`,
		`cn`,
		`sn`,
		`givenName`,
		`objectClass`,
	)

	fmt.Printf("Keyword: %s", attrs.Keyword())
	// Output: Keyword: targetfilter
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

func ExampleAV() {
	attr := AT(`homeDirectory`)
	value := AV(`/home/jesse`)

	fmt.Printf("My %s is '%s'", attr, value)
	// Output: My homeDirectory is '/home/jesse'
}

func ExampleAttributeValue() {
	attr := AT(`homeDirectory`)
	value := AV(`/home/jesse`)

	fmt.Printf("My %s is '%s'", attr, value)
	// Output: My homeDirectory is '/home/jesse'
}

func ExampleAttributeValue_String() {
	attr := AT(`homeDirectory`)
	value := AV(`/home/jesse`)

	fmt.Printf("My %s is '%s'", attr, value)
	// Output: My homeDirectory is '/home/jesse'
}

/*
This example demonstrates a SHA-1 hash comparison between the receiver and
a test value. Note that the SHA-1 comparison is driven by the type instance's
stringer output for simplicity in this package.
*/
func ExampleAttributeValue_Compare() {

	var (
		// granted, password values aren't an ACIv3 thing in
		// this context, but its still a fair example ...
		myPassword    AttributeValue = AV(`ALAA¢<ý²áßNß%a.)_ÿ3÷`)
		notMyPassword AttributeValue = AV(`ALAA¢<ýªáßNß%a.)_ÿ3÷`)
	)

	fmt.Printf("These passwords match: %t", myPassword.Compare(notMyPassword))
	// Output: These passwords match: false
}

func TestAttrs_codecov(t *testing.T) {

	var atv AttributeBindTypeOrValue
	atv.Set(AT(``))
	_ = atv.Eq()
	_ = atv.Ne()
	_ = atv.Valid()
	_ = atv.String()

	if err := atv.Parse(`fartknocker`); err == nil {
		t.Errorf("%s failed: no error where one was expected", t.Name())
		return
	}

	if atv = ABTV(BindGAT, AT(`manager`), SELFDN); !atv.Compare(`manager#SELFDN`) {
		t.Errorf("%s failed: hash comparison error", t.Name())
		return
	}
	_ = atv.Keyword()
	_ = atv.Valid()
	_ = atv.String()

	var at AttributeType
	_ = at.Eq()
	_ = at.Ne()
	_ = at.Len()
}

func TestAttributeTypes(t *testing.T) {

	for keyword, atfn := range map[Keyword]func(...any) AttributeTypes{
		BindUAT:    UAs, /// TODO clean this up
		BindGAT:    UAs, /// TODO clean this up
		TargetAttr: TAs,
	} {
		var attrs AttributeTypes
		_ = attrs.Eq()
		_ = attrs.Ne()
		_ = attrs.Len()
		_ = attrs.Valid()
		attrs.isAttributeTypeContext()
		attrs.reset()
		attrs.Push()
		attrs.Push(AT(``))
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

		attrs = atfn()

		attrs.Push()
		attrs.Push(AT(``))
		attrs.Push(keyword)
		attrs.resetKeyword(keyword)
		attrs.resetKeyword(keyword.String())
		attrs.setQuoteStyle(1)
		attrs.setQuoteStyle(0)
		attrs.Contains(3.14159)
		attrs.Push('𝝅')

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
				return
			}

			// process attribute
			attr = atn(raw)
			al = attrs.Len()

			if attrs.Push(attr); !attrs.Contains(raw) {
				t.Errorf("%s [%s] multival failed: valid %T[%s] instance (%s) not pushed into %T[%s; len:%d]",
					t.Name(), keyword, attr, keyword, attr, attrs, attrs.Keyword(), al)
				return
			}
			popped := attrs.Pop()
			attrs.Push(popped)
			attrs.resetKeyword(TargetCtrl)
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
					return
				}

				wcop := sprintf("( %s %s %q )", attrs.Keyword(), cop, actx.String())
				if T := meth(); T.String() != wcop {
					err := unexpectedStringResult(actx.String(), wcop, T.String())
					t.Errorf("%s multival failed [%s rule]: %v",
						t.Name(), attrs.Keyword(), err)
					return
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

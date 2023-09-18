package aci

import (
	"fmt"
	"testing"
)

func TestSearchFilter_setFromVar(t *testing.T) {
	want := `(&(objectClass=employee)(cn=Jesse Coretta))`
	var f SearchFilter

	_ = f.Eq()
	_ = f.Ne()

	// for codecov (check panic potential)
	if empty := f.String(); empty != `` {
		t.Errorf("%T failed: [%T.String]; should be empty",
			t.Name(), f)
		return
	}
	// for codecov (zero string set)
	f.Set(``)
	if err := f.Eq().Valid(); err == nil {
		t.Errorf("%T failed: [%T.Eq]; no error",
			t.Name(), f)
		return
	}

	if err := f.Ne().Valid(); err == nil {
		t.Errorf("%T failed: [%T.Eq]; no error",
			t.Name(), f)
		return
	}

	f.Set(want)

	if want != f.String() {
		t.Errorf("%T failed [Filter]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, f)
		return
	}
}

func TestFilter(t *testing.T) {
	raw := `(&(objectClass=employee)(cn=Jesse Coretta))`
	want := raw
	f := Filter(raw)

	if want != f.String() {
		t.Errorf("%T failed [Filter]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, f)
		return
	}

	c := f.Eq()
	want = sprintf("( targetfilter = %q )", raw)
	if want != c.String() {
		t.Errorf("%T failed [SearchFilter.Eq]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, f)
		return
	}

	c = f.Ne()
	want = sprintf("( targetfilter != %q )", raw)
	if want != c.String() {
		t.Errorf("%T failed [SearchFilter.Ne]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, f)
		return
	}
}

func TestFilter_Set(t *testing.T) {
	want := `(&(objectClass=employee)(cn=Jesse Coretta))`
	f := Filter()
	f.Set(want)

	if want != f.String() {
		t.Errorf("%T failed [SearchFilter.Set]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, f)
		return
	}
}

func TestAttributeFilter(t *testing.T) {
	var _afo AttributeFilter
	_ = _afo.Parse(` (       hi        ) `)
	_, _, _ = parseAttrFilterOperPreamble(` (       hi        )`)

	afo := AF(AT(`cn`))
	_ = afo.Valid()
	_ = afo.IsZero()
	_ = afo.AttributeType()
	_ = afo.SearchFilter()
	_ = afo.String()
	afo = AF(Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`))
	_ = afo.Valid()
	_ = afo.IsZero()
	_ = afo.AttributeType()
	_ = afo.SearchFilter()
	_ = afo.String()

	var af AttributeFilter
	_ = af.AttributeType()
	_ = af.SearchFilter()
	_ = af.String()
	_ = af.Valid()
	_ = af.IsZero()

	want := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	af = AF(`objectClass`, `(&(objectClass=employee)(cn=Jesse Coretta))`)
	if want != af.String() {
		t.Errorf("%T failed [AttributeFilter]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, af)
		return
	}
}

func TestAttributeFilterOperation_byStringParse(t *testing.T) {
	straf := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	want := `add=` + straf

	var afo AttributeFilterOperation

	afo.Push()
	_ = afo.Valid()
	_ = afo.String()
	_ = afo.Contains('𝝅')
	_ = afo.Push('𝝅')
	_ = afo.Push(nil)
	_ = afo.Eq()
	_ = afo.Pop()
	_ = afo.Valid()

	adder := AddOp.AFO(straf)

	afo = AFO()
	afo.Push()
	_ = afo.String()
	_ = afo.Valid()
	_ = afo.Contains('𝝅')
	_ = afo.Push('𝝅')
	_ = afo.Push(nil)
	_ = afo.Push(adder)
	_ = afo.Push(adder.String())
	_ = afo.Eq()
	_ = afo.Pop()
	_ = afo.Valid()

	var af AttributeFilter
	af.Parse(straf)

	_ = hasAttributeFilterOperationPrefix(``)
	_ = hasAttributeFilterOperationPrefix(af.String())

	_ = afo.Contains(``)
	_ = afo.Contains(af)
	afo.Push(af)
	_ = afo.Push(af.String())

	if afo.String() != want {
		t.Errorf("%T failed [AttributeFilterOperation.AFO]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, afo)
		return
	}
}

func TestAttributeFilterOperation_byType(t *testing.T) {
	at := AT(`objectClass`)
	sf := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	afo := AF(at, sf)

	want := `delete=` + at.String() + `:` + sf.String()

	dafo := DelOp.AFO(afo)
	if dafo.String() != want {
		t.Errorf("%T failed [AttributeFilterOperation.AFO]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, dafo)
		return
	}
}

func TestAttributeFilterOperation_addMultiVal(t *testing.T) {
	af1 := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	af2 := `homeDirectory:(&(objectClass=accountant)(cn=Courtney Tolana))`
	want := `add=` + af1 + ` && ` + af2

	afo := AddOp.AFO(af1, af2)
	if afo.String() != want {
		t.Errorf("%T failed [AttributeFilterOperation.AFO(Add)]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, afo)
		return
	}

	af2a := afo.Pop()
	if af2a.String() != af2 {
		t.Errorf("%T failed [AttributeFilterOperation.Pop]:\nwant '%s'\ngot  '%s'",
			t.Name(), af2, af2a)
		return
	}
}

func TestAttributeFilterOperation_delMultiVal(t *testing.T) {
	af1 := `nsroledn:(!(nsroledn=cn=X.500 Administrator))`
	af2 := `telephoneNumber:(telephoneNumber=456*)`
	want := `delete=` + af1 + ` && ` + af2

	afo := DelOp.AFO(af1, af2)
	if afo.String() != want {
		t.Errorf("%T failed [AttributeFilterOperation.AFO(Delete)]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, afo)
		return
	}

	_ = afo.Contains(nil)
	_ = afo.Contains(``)
	_ = afo.pushPolicy()
	_ = afo.pushPolicy(``)
	_ = afo.pushPolicy(AddOp)
	_ = afo.pushPolicy(nil, nil)
	_ = afo.pushPolicy(nil)
	_ = afo.pushPolicy(AF(``))

}

func TestAttributeFilterOperations_byStringParse(t *testing.T) {

	badaf1 := `add==,objcectlcass*`
	_, _ = parseAttributeFilterOperations(badaf1, 0)

	af1 := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	af2 := `homeDirectory:(&(objectClass=accountant)(cn=Courtney Tolana))`
	af3 := `nsroledn:(!(nsroledn=cn=X.500 Administrator))`
	af4 := `telephoneNumber:(telephoneNumber=456*)`
	want := `add=` + af1 + ` && ` + af2 + `,` + `delete=` + af3 + ` && ` + af4

	// Parse the entirety of the want literal.
	var afos AttributeFilterOperations
	_ = afos.Contains('𝝅')

	if err := afos.Parse(want); err != nil {
		t.Errorf("%s failed [AttributeFilterOperations.Parse(raw)]: %v",
			t.Name(), err)
		return
	}
	_ = afos.Contains(nil)
	_ = afos.pushPolicy()
	_ = afos.pushPolicy(``)
	_ = afos.pushPolicy(nil, nil)
	_ = afos.pushPolicy(nil)
	_ = afos.pushPolicy(AFO())

	// verify the complete string representation
	// matches that of the above want literal.
	if afos.String() != want {
		t.Errorf("%s failed [AttributeFilterOperations.Parse(compare)]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, afos)
		return
	}

	// verify the top-level stack's length.
	if afos.Len() != 2 {
		t.Errorf("%s failed [AttributeFilterOperations.Parse(chk AFOs len)]: want 'len:%d', got 'len:'%d'",
			t.Name(), 2, afos.Len())
		return
	}

	// scan the sub slices, verify those
	// lengths as well.
	for i := 0; i < 2; i++ {
		if afo := afos.Index(i); afo.Len() != 2 {
			t.Errorf("%s failed [AttributeFilterOperations.Parse(chk AFO len)]: want 'len:%d', got 'len:'%d'",
				t.Name(), 2, afo.Len())
			return
		}
	}

	// try to fool the parser by specifying a semi delimiter, but w/o
	// updating the above want string literal accordingly ...
	if err := afos.Parse(want, AttributeFilterOperationsSemiDelim); err == nil {
		t.Errorf("%s failed [AttributeFilterOperations.Parse(raw, alt delim)]: incorrect delimiter caused no error (but should have)",
			t.Name())
		return
	}
}

func TestAttributeFilterOperations_byTypes(t *testing.T) {
	af1 := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	af2 := `homeDirectory:(&(objectClass=accountant)(cn=Courtney Tolana))`
	af3 := `nsroledn:(!(nsroledn=cn=X.500 Administrator))`
	af4 := `telephoneNumber:(telephoneNumber=456*)`

	afos := AFOs(
		AddOp.AFO(af1, af2),
		DelOp.AFO(af3, af4),
	)

	want1 := `add=` + af1 + ` && ` + af2
	want2 := `delete=` + af3 + ` && ` + af4
	want := want1 + `,` + want2

	if afos.String() != want {
		t.Errorf("%T failed [AttributeFilterOperations.AFOs]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, afos)
		return
	}

	af2a := afos.Pop()
	if af2a.String() != want2 {
		t.Errorf("%T failed [AttributeFilterOperations.Pop]:\nwant '%s'\ngot  '%s'",
			t.Name(), want2, af2a)
		return
	}
}

func TestAttributeFilterOperation_toTargetRule(t *testing.T) {
	af := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	want := `( targattrfilters = "add=` + af + `" )`

	rule := AddOp.AFO(af).Eq()
	if rule.String() != want {
		t.Errorf("%T failed [AttributeFilterOperation.Eq]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, rule)
		return
	}
}

func TestAttributeFilterOperations_toTargetRule(t *testing.T) {
	af1 := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	af2 := `homeDirectory:(&(objectClass=accountant)(cn=Courtney Tolana))`
	af3 := `nsroledn:(!(nsroledn=cn=X.500 Administrator))`
	af4 := `telephoneNumber:(telephoneNumber=456*)`

	var afos AttributeFilterOperations
	afos.Push()
	_ = afos.Valid()
	_ = afos.Eq()

	// for codecov
	if !afos.IsZero() {
		t.Errorf("%s failed [%T.IsZero]:\nwant 'true'\ngot 'false'",
			t.Name(), afos)
		return
	}
	_ = afos.Contains('𝝅')

	adder := AddOp.AFO(af1, af2)
	deler := DelOp.AFO(af3, af4)

	afos = AFOs(
		adder,
		deler,
	)

	_ = afos.Push('𝝅')
	_ = afos.Push(nil)
	_ = afos.Push(adder)
	_ = afos.Push(adder.String())
	_ = afos.Push('𝝅')
	_ = afos.Push(``)
	_ = afos.Push(nil)
	_ = afos.Push(AddOp.AFO())
	_ = afos.SetDelimiter(1)
	_ = afos.SetDelimiter(0)
	_ = afos.SetDelimiter()
	_ = afos.Valid()

	rule := afos.Eq()

	want1 := `add=` + af1 + ` && ` + af2
	want2 := `delete=` + af3 + ` && ` + af4
	want := `( targattrfilters = "` + want1 + `,` + want2 + `" )`

	if rule.String() != want {
		t.Errorf("%T failed [AttributeFilterOperations.Eq]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, rule)
		return
	}
}

/*
This example demonstrates the creation of an instance of AttributeFilter, which
is strictly intended for use within instances of AttributeFilterOperation.

In this example, proper type instances are fed to the package level AF function
to form a complete AttributeFilter instance.

The return type, AttributeFilter, is then shown in string representation.
*/
func ExampleAF() {
	af := AF(
		AT(`homeDirectory`),
		Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`),
	)

	fmt.Printf("%s", af)
	// Output: homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))
}

/*
This example demonstrates the creation of an instance of AttributeFilter, which
is strictly intended for use within instances of AttributeFilterOperation.

In this example, a raw string representation of an AttributeFilter is used for
parser input.

The return type, AttributeFilter, is then interrogated by way of the AttributeType
and SearchFilter methods it makes available to the user. An alternative to this
approach is to simply use its String method to get the whole value, if desired.
*/
func ExampleAttributeFilter_Parse() {
	aftxt := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	var af AttributeFilter
	_ = af.Parse(`4537895439h`)
	err := af.Parse(aftxt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("AT:%q, FILTER:%q", af.AttributeType(), af.SearchFilter())
	// Output: AT:"objectClass", FILTER:"(&(objectClass=employee)(cn=Jesse Coretta))"
}

/*
This example demonstrates the creation of an instance of AttributeFilter, which
is strictly intended for use within instances of AttributeFilterOperation.

In this example, proper type instances are fed to the Set method to form a
complete AttributeFilter instance.

The return type, AttributeFilter, is then shown in string representation.
*/
func ExampleAttributeFilter_Set() {
	var af AttributeFilter // see also the package level AF function
	af.Set(
		AT(`homeDirectory`),
		Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`),
	)

	fmt.Printf("%s", af)
	// Output: homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))
}

/*
This example demonstrates the creation of an instance of AttributeFilter followed
by a call of its AttributeType method.

The return type, AttributeType, is shown in string representation.
*/
func ExampleAttributeFilter_AttributeType() {
	aftxt := `homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))`
	var af AttributeFilter
	_ = af.Parse(aftxt) // shadow error for brevity

	fmt.Printf("%s", af.AttributeType())
	// Output: homeDirectory
}

/*
This example demonstrates the creation of an instance of AttributeFilter followed
by a call of its SearchFilter method.

The return type, SearchFilter, is shown in string representation.
*/
func ExampleAttributeFilter_SearchFilter() {
	aftxt := `homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))`
	var af AttributeFilter
	_ = af.Parse(aftxt) // shadow error for brevity

	fmt.Printf("%s", af.SearchFilter())
	// Output: (&(objectClass=employee)(cn=Jesse Coretta))
}

/*
This example demonstrates a check of the receiver for "nilness".
*/
func ExampleAttributeFilter_IsZero() {
	var af AttributeFilter

	fmt.Printf("Is zero: %t (obviously)", af.IsZero())
	// Output: Is zero: true (obviously)
}

/*
This example demonstrates the interrogation of the receiver in order to
discern the appropriate Keyword.

Its string representation, along with the name of the Keyword type, is
shown.
*/
func ExampleAttributeFilter_Keyword() {
	var af AttributeFilter

	fmt.Printf("Keyword is '%s' (type:%T)", af.Keyword(), af.Keyword())
	// Output: Keyword is 'targattrfilters' (type:aci.TargetKeyword)
}

/*
This example demonstrates the creation of an instance of AttributeFilter followed
by a call of its String method.

The return value is the entirely of the receiver in string representation.
*/
func ExampleAttributeFilter_String() {
	aftxt := `homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))`
	var af AttributeFilter
	_ = af.Parse(aftxt) // shadow error for brevity

	fmt.Printf("%s", af)
	// Output: homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))
}

/*
This example demonstrates the creation of an instance of AttributeFilter followed
by a call of its Valid method for the purpose of sanity checking the receiver.

An error is reported and printed to STDOUT.
*/
func ExampleAttributeFilter_Valid() {
	var (
		af  AttributeFilter
		err error
	)

	if err = af.Valid(); err != nil {
		fmt.Println(err)
	}
	// Output: aci.AttributeFilter instance is nil
}

/*
This example demonstrates the creation of an instance of AttributeFilterOperation,
which is strictly intended for use within instances of AttributeFilterOperations.

In this example, proper type instances are fed to the package level AFO function
to form a complete AttributeFilterOperation instance.

The return type, AttributeFilterOperation, is then shown in string representation.
*/
func ExampleAFO() {
	// define the desired attributeType
	attr := AT(`homeDirectory`)

	// define the filter expression
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)

	// create an AttributeFilter instance
	// using the above components
	aF := AF(attr, filter)

	// When using the package level AFO function, it
	// is necessary to feed it an AttributeOperation
	// instance (either AddOp or DelOp) to define the
	// disposition of the new instance.
	aFO := AFO(AddOp, aF)

	fmt.Printf("%s", aFO)
	// Output: add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))
}

/*
This example demonstrates an alternative to the AFO package level example. Instead
of feeding an instance of AttributeOperation into the function, one can also use
the AttributeOperation constant itself to generate the type instance needed. This
may be useful in situations which require portability of certain functionality.
*/
func ExampleAttributeFilterOperation_byAttributeOperationConstants() {
	// define the desired attributeType
	attr := AT(`homeDirectory`)

	// define the filter expression
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)

	// create an AttributeFilter instance
	// using the above components
	aF := AF(attr, filter)

	// We'll use the Delete operation (DelOp)
	// package constant to spawn a new instance
	// of AttributeFilterOperation. This will
	// produce the same result as the AFO example
	// demonstrated earlier, except this time we
	// will impose the Delete operation.
	aFO := DelOp.AFO(aF)

	fmt.Printf("%s", aFO)
	// Output: delete=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Contains
method, allowing for basic text searches of the receiver.

Note that case is significant in the matching process for instances of this type.
*/
func ExampleAttributeFilterOperation_Contains() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	aFO := AddOp.AFO(aF1, aF2)

	fmt.Printf("%t", aFO.Contains(`homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))`))
	// Output: true
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Eq method,
allowing for the creation of a TargetRule instance containing the receiver value,
and bearing the `targattrfilters` keyword context.
*/
func ExampleAttributeFilterOperation_Eq() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	aFO := AddOp.AFO(aF1, aF2)

	fmt.Printf("%s", aFO.Eq())
	// Output: ( targattrfilters = "add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern))" )
}

/*
This example demonstrates the use of the AttributeFilterOperation type's F
method, which returns the appropriate slice building function for convenience.
*/
func ExampleAttributeFilterOperation_F() {
	var aFO AttributeFilterOperation

	// this returns the package-level AF function
	Func := aFO.F()

	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF := Func(attr, filter)

	fmt.Printf("%s", aF)
	// Output: homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Index
method to allow iteration of the receiver's contents.
*/
func ExampleAttributeFilterOperation_Index() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	aFO := AddOp.AFO(aF1, aF2)

	idx := aFO.Index(1) // 2nd index in stack

	fmt.Printf("%s", idx.AttributeType())
	// Output: gecos
}

/*
This example demonstrates a check of the receiver for "nilness" using the
AttributeFilterOperation type's IsZero method.
*/
func ExampleAttributeFilterOperation_IsZero() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF := AF(attr, filter)
	aFO := AddOp.AFO(aF)

	fmt.Printf("%t", aFO.IsZero())
	// Output: false
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Len
method to report the integer length of the receiver.
*/
func ExampleAttributeFilterOperation_Len() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	aFO := AddOp.AFO(aF1, aF2)

	fmt.Printf("Length: %d", aFO.Len())
	// Output: Length: 2
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Keyword
method.
*/
func ExampleAttributeFilterOperation_Keyword() {
	var aFO AttributeFilterOperation

	fmt.Printf("%s", aFO.Keyword())
	// Output: targattrfilters
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Kind
method.
*/
func ExampleAttributeFilterOperation_Kind() {
	var aFO AttributeFilterOperation

	fmt.Printf("%s", aFO.Kind())
	// Output: targattrfilters
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Ne method,
which is not intended to be used in any situation ever. See the comments for this
method for details.
*/
func ExampleAttributeFilterOperation_Ne() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	af := AF(attr, filter)
	aFO := AddOp.AFO(af)

	bogus := aFO.Ne()

	fmt.Printf("%t", bogus.IsZero())
	// Output: true
}

/*
This example demonstrates a check of the receiver's operational disposition
using the Operation method.
*/
func ExampleAttributeFilterOperation_Operation() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF := AF(attr, filter)
	aFO := AddOp.AFO(aF)

	fmt.Printf("%s", aFO.Operation())
	// Output: add
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Pop
method to remove the last stack slice per LIFO ordering.
*/
func ExampleAttributeFilterOperation_Pop() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	aFO := AddOp.AFO(aF1, aF2)

	popped := aFO.Pop()

	fmt.Printf("%s", popped.AttributeType())
	// Output: gecos
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Push
method to append a new (eligible) instance to the receiver.
*/
func ExampleAttributeFilterOperation_Push() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	aFO := AddOp.AFO(aF1)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	aFO.Push(aF2)

	fmt.Printf("%s", aFO)
	// Output: add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern))
}

/*
This example demonstrates the creation of an instance of AttributeFilterOperation
followed by a call of its String method.

The return value is the entirely of the receiver in string representation.
*/
func ExampleAttributeFilterOperation_String() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	aFO := AddOp.AFO(aF1, aF2)

	fmt.Printf("%s", aFO)
	// Output: add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern))
}

/*
This example demonstrates the creation of an instance of AttributeFilterOperation
followed by a call of its Valid method for the purpose of sanity checking the receiver.

No error is reported and printed to STDOUT in this case, as the instance is valid.
*/
func ExampleAttributeFilterOperation_Valid() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF := AF(attr, filter)
	aFO := AddOp.AFO(aF)

	fmt.Printf("%t", aFO.Valid() == nil)
	// Output: true
}

func ExampleAttributeFilterOperation_TRM() {
	var afo AttributeFilterOperation
	fmt.Printf("%d available aci.TargetRuleMethod instances", afo.TRM().Len())
	// Output: 2 available aci.TargetRuleMethod instances
}

/*
This example demonstrates the creation of an instance of AttributeFilterOperations,
which is used to store individual AttributeFilterOperation instances.

In this example, proper type instances are fed to the package level AFOs function
to form a complete AttributeFilterOperations instance.

The return type, AttributeFilterOperations, is then shown in string representation.
*/
func ExampleAFOs() {
	// define the desired attributeType
	// and filter for the first element
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	// Create the AttributeFilterOperation
	// instance (aFO)
	aFO := AddOp.AFO(aF1, aF2)

	// prepare our AttributeFilterOperations
	// instance stack (aFOs) using the AFOs
	// package level function.
	aFOs := AFOs(aFO)

	fmt.Printf("%s", aFOs)
	// Output: add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern))
}

/*
This example demonstrates the creation of an instance of AttributeFilterOperations,
which is used to store individual AttributeFilterOperation instances.

In this example, proper type instances are fed to the package level AFOs function
to form a complete AttributeFilterOperations instance.

The return type, AttributeFilterOperations, is then shown in string representation.
*/
func ExampleAttributeFilterOperations_Contains() {
	// define the desired attributeType
	// and filter for the first element
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	// Create the first AttributeFilterOperation
	// instance (aFO1)
	aFO1 := AddOp.AFO(aF1, aF2)

	attr = AT(`uidNumber`)
	filter = Filter(`(&(objectClass=accounting)(terminated=FALSE))`)
	aF3 := AF(attr, filter)

	attr = AT(`gidNumber`)
	filter = Filter(`(objectClass=account)`)
	aF4 := AF(attr, filter)

	// Create the second AttributeFilterOperation
	// instance (aFO2)
	aFO2 := DelOp.AFO(aF3, aF4)

	// prepare our AttributeFilterOperations
	// instance stack (aFOs) using the AFOs
	// package level function. Push both of
	// the above AFO instances.
	aFOs := AFOs(aFO1, aFO2)

	fmt.Printf("%t", aFOs.Contains(`delete=uidNumber:(&(objectClass=accounting)(terminated=FALSE)) && gidNumber:(objectClass=account)`))
	// Output: true
}

func ExampleAttributeFilterOperations_TRM() {
	var afos AttributeFilterOperations
	fmt.Printf("%d available aci.TargetRuleMethod instances", afos.TRM().Len())
	// Output: 2 available aci.TargetRuleMethod instances
}

/*
This example demonstrates the use of the AttributeFilterOperations type's Eq method,
allowing for the creation of a TargetRule instance containing the receiver value,
and bearing the `targattrfilters` keyword context.
*/
func ExampleAttributeFilterOperations_Eq() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	aFO := AddOp.AFO(aF1, aF2)

	// all of the above was copied verbatim
	// from the AttributeFilterOperation Eq
	// example. All we're really doing here
	// is enveloping it in another stack
	aFOs := AFOs(aFO)

	fmt.Printf("%s", aFOs.Eq())
	// Output: ( targattrfilters = "add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern))" )
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Ne method,
which is not intended to be used in any situation ever. See the comments for this
method for details.
*/
func ExampleAttributeFilterOperations_Ne() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	af := AF(attr, filter)
	aFO := AddOp.AFO(af)

	// all of the above was copied verbatim
	// from the AttributeFilterOperation Ne
	// example. All we're really doing here
	// is enveloping it in another stack
	aFOs := AFOs(aFO)
	bogus := aFOs.Ne()

	fmt.Printf("%t", bogus.IsZero())
	// Output: true
}

/*
This example demonstrates the use of the AttributeFilterOperations type's F
method, which returns the appropriate slice building function for convenience.
*/
func ExampleAttributeFilterOperations_F() {
	var aFOs AttributeFilterOperations

	// this returns the package-level AFO function
	Func := aFOs.F()

	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF := Func(attr, filter)
	aFO := Func(aF)

	fmt.Printf("%T", aFO)
	// Output: aci.AttributeFilterOperation
}

/*
This example demonstrates the use of the AttributeFilterOperations type's Index
method to allow iteration of the receiver's contents.
*/
func ExampleAttributeFilterOperations_Index() {
	// define the desired attributeType
	// and filter for the first element
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	// Create the first AttributeFilterOperation
	// instance (aFO1)
	aFO1 := AddOp.AFO(aF1, aF2)

	attr = AT(`uidNumber`)
	filter = Filter(`(&(objectClass=accounting)(terminated=FALSE))`)
	aF3 := AF(attr, filter)

	attr = AT(`gidNumber`)
	filter = Filter(`(objectClass=account)`)
	aF4 := AF(attr, filter)

	// Create the second AttributeFilterOperation
	// instance (aFO2)
	aFO2 := DelOp.AFO(aF3, aF4)

	// prepare our AttributeFilterOperations
	// instance stack (aFOs) using the AFOs
	// package level function. Push both of
	// the above AFO instances.
	aFOs := AFOs(aFO1, aFO2)

	slice := aFOs.Index(1)

	fmt.Printf("%s", slice)
	// Output: delete=uidNumber:(&(objectClass=accounting)(terminated=FALSE)) && gidNumber:(objectClass=account)
}

/*
This example demonstrates a check of the receiver for "nilness" using the
AttributeFilterOperations type's IsZero method.
*/
func ExampleAttributeFilterOperations_IsZero() {
	var afo AttributeFilterOperations
	fmt.Printf("%t", afo.IsZero())
	// Output: true
}

/*
This example demonstrates the use of the AttributeFilterOperation type's Len
method to report the integer length of the receiver.
*/
func ExampleAttributeFilterOperations_Len() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	aFOs := AFOs(AddOp.AFO(aF1, aF2))

	fmt.Printf("Length: %d", aFOs.Len())
	// Output: Length: 1
}

/*
This example demonstrates the use of the AttributeFilterOperations type's Keyword
method.
*/
func ExampleAttributeFilterOperations_Keyword() {
	var afos AttributeFilterOperations

	fmt.Printf("%s", afos.Keyword())
	// Output: targattrfilters
}

/*
This example demonstrates the use of the AttributeFilterOperations type's Kind
method.
*/
func ExampleAttributeFilterOperations_Kind() {
	var aFOs AttributeFilterOperations

	fmt.Printf("%s", aFOs.Kind())
	// Output: targattrfilters
}

/*
This example demonstrates the creation of an instance of AttributeFilterOperations.

In this example, a raw string representation of an AttributeFilterOperations is used
for parser input. The resultant output from the instance's String method should be
identical to that which was fed into the Parse method.
*/
func ExampleAttributeFilterOperations_Parse() {
	aftxt := `delete=objectClass:(&(objectClass=employee)(cn=Jesse Coretta)) && homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta))`
	var afos AttributeFilterOperations
	err := afos.Parse(aftxt)
	if err != nil {
		fmt.Println(err)
		return
	}
	got := afos.String()

	fmt.Printf("Values match: %t", got == aftxt)
	// Output: Values match: true
}

/*
This example demonstrates the use of the AttributeFilterOperations type's Push
method to append a new (eligible) instance to the receiver.
*/
func ExampleAttributeFilterOperations_Push() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	var afos AttributeFilterOperations = AFOs()
	afos.Push(AddOp.AFO(aF1, aF2))

	fmt.Printf("%d", afos.Len())
	// Output: 1
}

/*
This example demonstrates the use of the AttributeFilterOperations type's Push
method to append a new (eligible) instance to the receiver.
*/
func ExampleAttributeFilterOperations_Push_byString() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	var afos AttributeFilterOperations = AFOs()
	ad := AddOp.AFO(aF1, aF2)
	afos.Push(ad.String())

	fmt.Printf("%d", afos.Index(0).Len())
	// Output: 2
}

/*
This example demonstrates the use of the AttributeFilterOperations type's Pop
method to remove the last stack slice per LIFO ordering.
*/
func ExampleAttributeFilterOperations_Pop() {
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	aFOs := AFOs(DelOp.AFO(aF1, aF2))

	popped := aFOs.Pop()

	fmt.Printf("%s", popped)
	// Output: delete=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern))
}

func ExampleAttributeFilterOperations_SetDelimiter() {
	// define the desired attributeType
	// and filter for the first element
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	// Create the first AttributeFilterOperation
	// instance (aFO1)
	aFO1 := AddOp.AFO(aF1, aF2)

	attr = AT(`uidNumber`)
	filter = Filter(`(&(objectClass=accounting)(terminated=FALSE))`)
	aF3 := AF(attr, filter)

	attr = AT(`gidNumber`)
	filter = Filter(`(objectClass=account)`)
	aF4 := AF(attr, filter)

	// Create the second AttributeFilterOperation
	// instance (aFO2)
	aFO2 := DelOp.AFO(aF3, aF4)

	// prepare our AttributeFilterOperations
	// instance stack (aFOs) using the AFOs
	// package level function. Push both of
	// the above AFO instances, and set a
	// delimiter other than the default...
	delimiter := AttributeFilterOperationsSemiDelim

	aFOs := AFOs(aFO1, aFO2).SetDelimiter(delimiter)
	fmt.Printf("%s", aFOs)
	// Output: add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern));delete=uidNumber:(&(objectClass=accounting)(terminated=FALSE)) && gidNumber:(objectClass=account)
}

func ExampleAttributeFilterOperations_String() {
	// define the desired attributeType
	// and filter for the first element
	attr := AT(`homeDirectory`)
	filter := Filter(`(&(objectClass=employee)(cn=Jesse Coretta))`)
	aF1 := AF(attr, filter)

	attr = AT(`gecos`)
	filter = Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF2 := AF(attr, filter)

	// Create the first AttributeFilterOperation
	// instance (aFO1)
	aFO1 := AddOp.AFO(aF1, aF2)

	attr = AT(`uidNumber`)
	filter = Filter(`(&(objectClass=accounting)(terminated=FALSE))`)
	aF3 := AF(attr, filter)

	attr = AT(`gidNumber`)
	filter = Filter(`(objectClass=account)`)
	aF4 := AF(attr, filter)

	// Create the second AttributeFilterOperation
	// instance (aFO2)
	aFO2 := DelOp.AFO(aF3, aF4)

	// prepare our AttributeFilterOperations
	// instance stack (aFOs) using the AFOs
	// package level function. Push both of
	// the above AFO instances

	aFOs := AFOs(aFO1, aFO2)
	fmt.Printf("%s", aFOs)
	// Output: add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern)),delete=uidNumber:(&(objectClass=accounting)(terminated=FALSE)) && gidNumber:(objectClass=account)
}

/*
This example demonstrates the creation of an instance of AttributeFilterOperations
followed by a call of its Valid method for the purpose of sanity checking the receiver.
*/
func ExampleAttributeFilterOperations_Valid() {
	attr := AT(`uidNumber`)
	filter := Filter(`(&(objectClass=accounting)(terminated=FALSE))`)
	aF1 := AF(attr, filter)

	attr = AT(`gidNumber`)
	filter = Filter(`(objectClass=account)`)
	aF2 := AF(attr, filter)

	// Create the second AttributeFilterOperation
	// instance (aFO2)
	var afos AttributeFilterOperations = AFOs()
	afos.Push(DelOp.AFO(aF1, aF2))

	fmt.Printf("Valid: %t", afos.Valid() == nil)
	// Output: Valid: true
}

/*
This example demonstrates the creation of an instance of AttributeFilterOperation,
which is strictly intended for use within instances of AttributeFilterOperations.

In this example, proper type instances are fed to the AttributeOperation AFO method
to form a complete AttributeFilterOperation instance.

The return type, AttributeFilterOperation, has its type shown in text, along with
the current stack length.
*/
func ExampleAttributeOperation_AFO() {
	attr := AT(`gecos`)
	filter := Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	aF := AF(attr, filter)

	// Create the first AttributeFilterOperation
	// instance (aFO1)
	aFO := AddOp.AFO(aF)

	fmt.Printf("%T [len:%d]", aFO, aFO.Len())
	// Output: aci.AttributeFilterOperation [len:1]
}

/*
This example demonstrates the SHA-1 hash comparison between two (2) AttributeFilterOperation
instances using the Compare method.
*/
func ExampleAttributeFilterOperation_Compare() {
	attr := AT(`uidNumber`)
	filter := Filter(`(&(objectClass=accounting)(terminated=FALSE))`)
	aF1 := AF(attr, filter)
	aFO1 := AddOp.AFO(aF1)

	attr = AT(`gidNumber`)
	filter = Filter(`(&(objectClass=accounting)(terminated=FALSE))`)
	aF2 := AF(attr, filter)
	aFO2 := DelOp.AFO(aF2)

	fmt.Printf("Hashes are equal: %t", aFO1.Compare(aFO2))
	// Output: Hashes are equal: false
}

/*
This example demonstrates the SHA-1 hash comparison between two (2) AttributeFilterOperations
instances using the Compare method.

The comparison returns false, as the compared instances are ordered differently.
*/
func ExampleAttributeFilterOperations_Compare() {
	attr := AT(`uidNumber`)
	filter := Filter(`(&(objectClass=accounting)(terminated=FALSE))`)
	aF1 := AF(attr, filter)
	aFO1 := AddOp.AFO(aF1)

	attr = AT(`gidNumber`)
	filter = Filter(`(&(objectClass=accounting)(terminated=FALSE))`)
	aF2 := AF(attr, filter)
	aFO2 := DelOp.AFO(aF2)

	// Create the second AttributeFilterOperation
	// instance (aFO2)
	var afos1 AttributeFilterOperations = AFOs()
	afos1.Push(aFO1, aFO2)

	var afos2 AttributeFilterOperations = AFOs()
	afos2.Push(aFO2, aFO1)

	fmt.Printf("Hashes are equal: %t", afos1.Compare(afos2))
	// Output: Hashes are equal: false
}

/*
This example demonstrates the string representation of an AttributeOperation
constant.

In this example, the instance is shown in string representation.
*/
func ExampleAttributeOperation_String() {
	fmt.Printf("%s and %s", AddOp, DelOp)
	// Output: add and delete
}

func ExampleSearchFilter_IsZero() {
	filter := Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	fmt.Printf("%t", filter.IsZero())
	// Output: false
}

func ExampleSearchFilter_Keyword() {
	filter := Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	fmt.Printf("%s", filter.Keyword())
	// Output: targetfilter
}

func ExampleSearchFilter_Set() {
	var filter SearchFilter
	filter.Set(`(|(objectClass=contractor)(objectClass=intern))`)
	fmt.Printf("%t", filter.IsZero())
	// Output: false
}

func ExampleSearchFilter_String() {
	var filter SearchFilter
	filter.Set(`(|(objectClass=contractor)(objectClass=intern))`)
	fmt.Printf("%s", filter)
	// Output: (|(objectClass=contractor)(objectClass=intern))
}

func ExampleSearchFilter_Eq() {
	var filter SearchFilter
	filter.Set(`(|(objectClass=contractor)(objectClass=intern))`)
	fmt.Printf("%s", filter.Eq())
	// Output: ( targetfilter = "(|(objectClass=contractor)(objectClass=intern))" )
}

func ExampleSearchFilter_TRM() {
	var filter SearchFilter
	filter.Set(`(|(objectClass=contractor)(objectClass=intern))`)
	fmt.Printf("%d available aci.TargetRuleMethod instances", filter.TRM().Len())
	// Output: 2 available aci.TargetRuleMethod instances
}

func ExampleSearchFilter_Ne() {
	var filter SearchFilter
	filter.Set(`(&(objectClass=contractor)(objectClass=intern))`)
	fmt.Printf("%s", filter.Ne())
	// Output: ( targetfilter != "(&(objectClass=contractor)(objectClass=intern))" )
}

func ExampleSearchFilter_Valid() {
	var filter SearchFilter
	fmt.Printf("%v", filter.Valid())
	// Output: aci.SearchFilter instance is nil
}

/*
This example demonstrates the SHA-1 hash comparison between two (2) SearchFilter
instances using the Compare method.
*/
func ExampleSearchFilter_Compare() {
	f1 := Filter(`(|(objectClass=contractor)(objectClass=intern))`)
	f2 := Filter(`(|(objectClass=contractor)(objectClass=intern))`)

	fmt.Printf("Hashes are equal: %t", f1.Compare(f2))
	// Output: Hashes are equal: true
}

/*
This example demonstrates the SHA-1 hash comparison between two (2) AttributeFilter
instances using the Compare method.
*/
func ExampleAttributeFilter_Compare() {
	f1 := AF(`objectClass`, `(|(objectClass=contractor)(objectClass=intern))`)
	f2 := AF(`homeDirectory`, `(|(objectClass=contractor)(objectClass=intern))`)

	fmt.Printf("Hashes are equal: %t", f1.Compare(f2))
	// Output: Hashes are equal: false
}

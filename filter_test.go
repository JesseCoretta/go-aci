package aci

import "testing"

func TestSearchFilter_setFromVar(t *testing.T) {
	want := `(&(objectClass=employee)(cn=Jesse Coretta))`
	var f SearchFilter

	// for codecov (check panic potential)
	if empty := f.String(); empty != `` {
		t.Errorf("%T failed: [%T.String]; should be empty",
			t.Name(), f)
	}
	// for codecov (zero string set)
	f.Set(``)
	if err := f.Eq().Valid(); err != nil {
		t.Errorf("%T failed: [%T.Eq];\nerror: %v",
			t.Name(), f, err)
	}

	if err := f.Ne().Valid(); err != nil {
		t.Errorf("%T failed: [%T.Ne];\nerror: %v",
			t.Name(), f, err)
	}

	f.Set(want)

	if want != f.String() {
		t.Errorf("%T failed [Filter]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, f)
	}
}

func TestFilter(t *testing.T) {
	raw := `(&(objectClass=employee)(cn=Jesse Coretta))`
	want := raw
	f := Filter(raw)

	if want != f.String() {
		t.Errorf("%T failed [Filter]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, f)
	}

	c := f.Eq()
	want = sprintf("( targetfilter = %q )", raw)
	if want != c.String() {
		t.Errorf("%T failed [SearchFilter.Eq]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, f)
	}

	c = f.Ne()
	want = sprintf("( targetfilter != %q )", raw)
	if want != c.String() {
		t.Errorf("%T failed [SearchFilter.Ne]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, f)
	}
}

func TestFilter_Set(t *testing.T) {
	want := `(&(objectClass=employee)(cn=Jesse Coretta))`
	f := Filter()
	f.Set(want)

	if want != f.String() {
		t.Errorf("%T failed [SearchFilter.Set]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, f)
	}
}

func TestAttributeFilter(t *testing.T) {
	want := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	af := AF(`objectClass`, `(&(objectClass=employee)(cn=Jesse Coretta))`)
	if want != af.String() {
		t.Errorf("%T failed [AttributeFilter]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, af)
	}
}

func TestAttributeFilterOperation_byStringParse(t *testing.T) {
	straf := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	want := `add=` + straf

	afo := AddOp.AFO(straf)
	if afo.String() != want {
		t.Errorf("%T failed [AttributeFilterOperation.AFO]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, afo)
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
	}

	af2a := afo.Pop()
	if af2a.String() != af2 {
		t.Errorf("%T failed [AttributeFilterOperation.Pop]:\nwant '%s'\ngot  '%s'",
			t.Name(), af2, af2a)
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
	}
}

func TestAttributeFilterOperations_byStringParse(t *testing.T) {
	af1 := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	af2 := `homeDirectory:(&(objectClass=accountant)(cn=Courtney Tolana))`
	af3 := `nsroledn:(!(nsroledn=cn=X.500 Administrator))`
	af4 := `telephoneNumber:(telephoneNumber=456*)`
	want := `add=` + af1 + ` && ` + af2 + `,` + `delete=` + af3 + ` && ` + af4

	// Parse the entirety of the want literal.
	var afos AttributeFilterOperations
	if err := afos.Parse(want); err != nil {
		t.Errorf("%s failed [AttributeFilterOperations.Parse(raw)]: %v",
			t.Name(), err)
	}

	// verify the complete string representation
	// matches that of the above want literal.
	if afos.String() != want {
		t.Errorf("%s failed [AttributeFilterOperations.Parse(compare)]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, afos)
	}

	// verify the top-level stack's length.
	if afos.Len() != 2 {
		t.Errorf("%s failed [AttributeFilterOperations.Parse(chk AFOs len)]: want 'len:%d', got 'len:'%d'",
			t.Name(), 2, afos.Len())
	}

	// scan the sub slices, verify those
	// lengths as well.
	for i := 0; i < 2; i++ {
		if afo := afos.Index(i); afo.Len() != 2 {
			t.Errorf("%s failed [AttributeFilterOperations.Parse(chk AFO len)]: want 'len:%d', got 'len:'%d'",
				t.Name(), 2, afo.Len())
		}
	}

	// try to fool the parser by specifying a semi delimiter, but w/o
	// updating the above want string literal accordingly ...
	if err := afos.Parse(want, AttributeFilterOperationsSemiDelim); err == nil {
		t.Errorf("%s failed [AttributeFilterOperations.Parse(raw, alt delim)]: incorrect delimiter caused no error (but should have)",
			t.Name())
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
	}

	af2a := afos.Pop()
	if af2a.String() != want2 {
		t.Errorf("%T failed [AttributeFilterOperations.Pop]:\nwant '%s'\ngot  '%s'",
			t.Name(), want2, af2a)
	}
}

func TestAttributeFilterOperation_toTargetRule(t *testing.T) {
	af := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	want := `( targattrfilters = "add=` + af + `" )`

	rule := AddOp.AFO(af).Eq()
	if rule.String() != want {
		t.Errorf("%T failed [AttributeFilterOperation.Eq]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, rule)
	}
}

func TestAttributeFilterOperations_toTargetRule(t *testing.T) {
	af1 := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	af2 := `homeDirectory:(&(objectClass=accountant)(cn=Courtney Tolana))`
	af3 := `nsroledn:(!(nsroledn=cn=X.500 Administrator))`
	af4 := `telephoneNumber:(telephoneNumber=456*)`

	var afos AttributeFilterOperations

	// for codecov
	if !afos.IsZero() {
		t.Errorf("%s failed [%T.IsZero]:\nwant 'true'\ngot 'false'",
			t.Name(), afos)
	}

	afos = AFOs(
		AddOp.AFO(af1, af2),
		DelOp.AFO(af3, af4),
	)
	rule := afos.Eq()

	want1 := `add=` + af1 + ` && ` + af2
	want2 := `delete=` + af3 + ` && ` + af4
	want := `( targattrfilters = "` + want1 + `,` + want2 + `" )`

	if rule.String() != want {
		t.Errorf("%T failed [AttributeFilterOperations.Eq]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, rule)
	}
}

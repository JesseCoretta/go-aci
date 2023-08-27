package aci

import "testing"

func TestSearchFilter_setFromVar(t *testing.T) {
        want := `(&(objectClass=employee)(cn=Jesse Coretta))`
	var f SearchFilter
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
	af := AF(`objectClass`,`(&(objectClass=employee)(cn=Jesse Coretta))`)
	if want != af.String() {
		t.Errorf("%T failed [AttributeFilter]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, af)
	}
}

func TestAttributeFilterOperation(t *testing.T) {
	af := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
	want := `add=` + af

	afo := AddOp.AFO(af)
	if afo.String() != want {
		t.Errorf("%T failed [AttributeFilter.AFO]:\nwant '%s'\ngot  '%s'",
			t.Name(), want, afo)
	}
}

func TestAttributeFilterOperation_addMultiVal(t *testing.T) {
        af1 := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
        af2 := `homeDirectory:(&(objectClass=accountant)(cn=Courtney Tolana))`
        want := `add=` + af1 + ` && ` + af2

        afo := AddOp.AFO(af1,af2)
        if afo.String() != want {
                t.Errorf("%T failed [AttributeFilterOperation.AFO(Add)]:\nwant '%s'\ngot  '%s'",
                        t.Name(), want, afo)
        }
}


func TestAttributeFilterOperation_delMultiVal(t *testing.T) {
	af1 := `nsroledn:(!(nsroledn=cn=X.500 Administrator))`
	af2 := `telephoneNumber:(telephoneNumber=456*)`
        want := `delete=` + af1 + ` && ` + af2

        afo := DelOp.AFO(af1,af2)
        if afo.String() != want {
                t.Errorf("%T failed [AttributeFilterOperation.AFO(Delete)]:\nwant '%s'\ngot  '%s'",
                        t.Name(), want, afo)
        }
}

func TestAttributeFilterOperations(t *testing.T) {
        af1 := `objectClass:(&(objectClass=employee)(cn=Jesse Coretta))`
        af2 := `homeDirectory:(&(objectClass=accountant)(cn=Courtney Tolana))`
	af3 := `nsroledn:(!(nsroledn=cn=X.500 Administrator))`
	af4 := `telephoneNumber:(telephoneNumber=456*)`

	afos := AFOs().Push(
		AddOp.AFO(af1,af2),
		DelOp.AFO(af3,af4),
	)

        want1 := `add=` + af1 + ` && ` + af2
	want2 := `delete=` + af3 + ` && ` + af4
	want := want1 + `,` + want2

        if afos.String() != want {
                t.Errorf("%T failed [AttributeFilterOperations.AFOs]:\nwant '%s'\ngot  '%s'",
                        t.Name(), want, afos)
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

        afos := AFOs().Push(
                AddOp.AFO(af1,af2),
                DelOp.AFO(af3,af4),
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

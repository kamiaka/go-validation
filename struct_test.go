package validation

import (
	"fmt"
	"reflect"
	"testing"
)

type SubStructParams struct {
	A string
	B string
}

type StructParams struct {
	Data *SubStructParams
}

func TestStruct(t *testing.T) {
	cases := []struct {
		params      *StructParams
		wantErrMsgs []string
	}{
		{
			params:      &StructParams{},
			wantErrMsgs: []string{fmt.Sprintf(MsgRequiredFormat, "data")},
		},
		{
			params: &StructParams{Data: &SubStructParams{}},
			wantErrMsgs: []string{
				fmt.Sprintf(MsgRequiredFormat, "a"),
				fmt.Sprintf(MsgRequiredFormat, "b"),
			},
		},
	}
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("failed to new validator")
	}
	for i, tc := range cases {
		err := v.Validate(tc.params, Field("data", &tc.params.Data, Required,
			DeepStructRuleFunc(func(v FieldValue) (rules []StructRule, err error) {
				p := v.Interface().(*SubStructParams)
				return append(rules,
					Field("a", &p.A, Required),
					Field("b", &p.B, Required),
				), nil
			}),
		))
		if err != nil {
			errs, ok := err.(Errors)
			if !ok {
				t.Fatalf("#%d: Validate(%#v) returns error: %s", i, tc.params, err)
			}
			msgs := make([]string, len(errs))
			for i, err := range errs {
				msgs[i] = err.Error()
			}
			if !reflect.DeepEqual(msgs, tc.wantErrMsgs) {
				t.Errorf("#%d: Validate(%#v) returns errors: %#v, want: %#v", i, tc.params, msgs, tc.wantErrMsgs)
			}
		}
	}
}

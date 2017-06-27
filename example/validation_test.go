package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	validation "github.com/kamiaka/go-validation"
)

type CreateUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	UsesMail  *bool  `json:"usesMail"`
	MailQuota int64  `json:"mailQuota"`
}

func jsonFieldName(field *reflect.StructField) string {
	name := field.Tag.Get("json")
	if name != "" {
		return strings.SplitN(name, ",", 2)[0]
	}
	return field.Name
}

func TestValidate(t *testing.T) {
	r := &CreateUserRequest{
		Password: "ng",
	}
	v, _ := validation.NewValidator(validation.FieldNameFunc(jsonFieldName))

	err := v.Validate(
		r,
		validation.Field("username", &r.Username, validation.Required, validation.MaxLength(4)),
		validation.Field("password", &r.Password, validation.Required, validation.Length(4, 16)),
		validation.Field("using mail", &r.UsesMail, validation.Required),
		validation.Field("quota", &r.MailQuota, mustUsesMail, validation.FieldRuleFunc(func(fi validation.FieldInfo, e validation.ErrorFunc) error {
			// set custom rule
			if fi.IsEmpty() {
				return nil
			}
			if r.UsesMail == nil || *r.UsesMail != true {
				e("Setting %[1]v requires %[2]v as truthy", fi.Label(), "using mail")
			}
			return nil
		}), validation.Max(1024)),
		validation.StructRuleFunc(func(si validation.StructInfo, e validation.ErrorFunc) error {
			// set custom struct level validation.
			if r.Username == r.Password {
				e("You are foolish!")
			}
			return nil
		}),
	)
	fmt.Printf("err: %+v\n", err)
}

var mustUsesMail validation.FieldRuleFunc = func(fi validation.FieldInfo, e validation.ErrorFunc) error {
	if validation.IsEmpty(fi.Value()) {
		return nil
	}
	s := fi.Parent().Interface().(CreateUserRequest)
	if s.UsesMail != nil && *s.UsesMail != true {
		return e("%s is bad!", fi.Label())
	}
	return nil
}

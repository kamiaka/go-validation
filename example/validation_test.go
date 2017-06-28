package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	validation "github.com/kamiaka/go-validation"
	"github.com/kamiaka/go-validation/is"
)

type CreateUserRequest struct {
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	UsesMail  *bool    `json:"usesMail"`
	MailQuota int64    `json:"mailQuota"`
	Domains   []string `json:"domains"`
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
		Domains:  []string{"", "."},
	}
	v, _ := validation.NewValidator(validation.FieldNameFunc(jsonFieldName))

	err := v.Validate(
		r,
		validation.Field("username", &r.Username, validation.Required, validation.MaxLength(4)),
		validation.Field("password", &r.Password, validation.Required, validation.Length(4, 16)),
		validation.Field("using mail", &r.UsesMail, validation.Required),
		validation.Field("ftp users", &r.Domains, validation.Required, validation.MaxLength(2), validation.Repeat(
			validation.Required.SetErrorFormat("value of %[1]v is required"), is.DNSName,
		)),
		validation.Field("quota", &r.MailQuota, mustUsesMail, validation.FieldRuleFunc(func(fi validation.FieldValue, e validation.ErrorFunc) error {
			// set custom rule
			if fi.IsEmpty() {
				return nil
			}
			if r.UsesMail == nil || *r.UsesMail != true {
				e("Setting %[1]v requires %[2]v as truthy", fi.Label(), "using mail")
			}
			return nil
		}), validation.Max(1024)),
		validation.StructRuleFunc(func(v validation.Value, e validation.ErrorFunc) error {
			// set custom struct level validation.
			if r.Username == r.Password {
				e("You are foolish!")
			}
			return nil
		}),
	)
	if err != nil {
		if errs, ok := err.(validation.Errors); ok {
			for _, e := range errs {
				if fe, ok := e.(validation.Error); ok {
					fmt.Printf("%s: %#v\n", fe.Value().Namespace(), fe.Error())
				} else {
					fmt.Printf("struct: %#v\n", e.Error())
				}
			}
		}
	}
}

var mustUsesMail validation.FieldRuleFunc = func(fi validation.FieldValue, e validation.ErrorFunc) error {
	if validation.IsEmpty(fi.Value()) {
		return nil
	}
	s := fi.Parent().Interface().(CreateUserRequest)
	if s.UsesMail != nil && *s.UsesMail != true {
		return e("%s is bad!", fi.Label())
	}
	return nil
}

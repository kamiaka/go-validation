package main

import (
	"fmt"

	validation "github.com/kamiaka/go-validation"
	"github.com/kamiaka/go-validation/is"
)

type request struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Mail     *mail    `json:"mail"`
	Domains  []string `json:"domains"`
	Type     int      `json:"type"`
}

type mail struct {
	IsEnabled *bool `json:"isEnabled"`
	Quota     int64 `json:"quota"`
}

func main() {
	// print "ok"
	validateAndPrintError(&request{
		Username: "太郎",
		Password: "abcdefg",
		Domains:  []string{"example.com"},
		Type:     42,
	})

	// print `
	// 	password: password must between 4 and 16
	// 	mail.isEnabled: using mail is required
	// 	mail.quota: setting quota requires using mail
	// 	domains[0]: value of my domains is required
	// 	domains[1]: my domains must be a valid DNS name
	// 	type: type must be a valid value
	// 	: You are foolish!
	// `
	validateAndPrintError(&request{
		Username: "invalid",
		Password: "ng",
		Domains:  []string{"", "."},
		Mail: &mail{
			Quota: 200,
		},
		Type: 43,
	})
}

func validateAndPrintError(r *request) {
	v, _ := validation.NewValidator()
	err := v.Validate(
		r,
		validation.Field("username", &r.Username, validation.Required, validation.CharMaxLength(4)),
		validation.Field("password", &r.Password, validation.Required, validation.StringLength(4, 16)),
		validation.Field("mail", &r.Mail, validation.DeepStructRuleFunc(func(v validation.FieldValue) (rules []validation.StructRule, err error) {
			p := v.Interface().(*mail)
			return append(rules,
				validation.Field("using mail", &p.IsEnabled, validation.Required),
				validation.Field("quota", &p.Quota, mustEnableWhenInput),
			), nil
		})),
		validation.Field("my domains", &r.Domains, validation.Required, validation.MaxLength(2), validation.Repeat(
			validation.Required.SetErrorFormat("value of %[1]v is required"), is.DNSName,
		)),
		validation.Field("type", &r.Type, validation.In("ok", "foo", 42)),
		validation.StructRuleFunc(func(v validation.Value, e validation.ErrorFunc) error {
			// set custom struct level validation.
			if r.Username == r.Password {
				return e("You are foolish!")
			}
			return nil
		}),
	)
	if err != nil {
		if errs, ok := err.(validation.Errors); ok {
			for _, e := range errs {
				fmt.Printf("%s: %s\n", e.Value().Namespace(), e.Error())
			}
		} else {
			fmt.Printf("internal error while validate: %s", err)
		}
	} else {
		fmt.Println("ok")
	}
}

var mustEnableWhenInput = validation.FieldRuleFunc(func(v validation.FieldValue, e validation.ErrorFunc) error {
	if v.IsEmpty() {
		return nil
	}
	s := v.Parent().Interface().(*mail)
	if s.IsEnabled == nil || *s.IsEnabled != true {
		return e("setting %s requires using mail", v.Label())
	}
	return nil
})

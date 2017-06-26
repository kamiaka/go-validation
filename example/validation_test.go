package main

import (
	"fmt"
	"strings"
	"testing"

	validation "github.com/kamiaka/go-validation"
)

type UpdatePasswordRequest struct {
	NewPassword string `json:"new_password"`
	OldPassword string ``
}

func jsonFieldName(field reflect.Field) string {
	name := field.Tag.Get("json")
	if name != "" {
		return strings.SplitN(name, ",", 2)[0]
	}
	return field.Name
}

func TestValidate(t *testing.T) {
	r := &UpdatePasswordRequest{
		NewPassword: "NEW_PASS",
		OldPassword: "OLD_PASS",
	}
	v := validation.NewValidator(FieldNameFunc(jsonFieldName))
	fmt.Printf("v: %#v\n", v)
}

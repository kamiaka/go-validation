package validation

import (
	"fmt"
)

type (
	// FieldRule is validation rule for field.
	FieldRule interface {
		Validate(FieldLevel) FieldError
	}

	// StructRule is validation rule for pointer of a struct.
	StructRule interface {
		Validate(StructLevel) StructError
	}
	StructLevel interface {
		Value() interface{}
		Parent() interface{}
	}
	FieldLevel interface {
		Label() string
		Key() string
		Value() interface{}
		Parent() StructLevel
	}
	fieldRules struct {
		label    string
		fieldPtr interface{}
		rules    []FieldRule
	}
)

func (b *fieldRules) Validate(sl StructLevel) error {
	field := &field{
		label: b.label,
		sl:    sl,
	}
	for _, rule := range b.rules {
		err := rule.Validate(field)
		if err != nil {
			return err
		}
	}
	return nil
}

func ValidateField(label string, fieldPtr interface{}, rules ...FieldRule) StructRule {
	return &fieldRules{
		label:    label,
		fieldPtr: fieldPtr,
		rules:    rules,
	}
}

// Validate struct.
func Validate(structPtr interface{}, validators ...StructRule) error {
	for _, v := range validators {
		fmt.Printf("validatable: %#v\n", v)
		err := v.Validate()
		if err == nil {
			continue
		}

		if _, ok := err.(FieldError); ok {
			continue
		}

		if _, ok := err.(StructError); ok {
			continue
		}

		return err
	}
	return nil
}

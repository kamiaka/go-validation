package validation

import "fmt"

type arrayValidator struct {
	rules []FieldRule
}

func Repeat(rules ...FieldRule) FieldRule {
	return &arrayValidator{
		rules: rules,
	}
}

func (v *arrayValidator) Apply(fi FieldInfo) error {
	fmt.Printf("WIP repeat: ")
	return nil
}

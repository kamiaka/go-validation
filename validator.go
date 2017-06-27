package validation

import (
	"fmt"
	"reflect"
)

type fieldNameFunc func(*reflect.StructField) string

type validator struct {
	config ValidatorConfig
}

type Validator interface {
	Validate(structPtr interface{}, vs ...StructRule) error
}

type Validatable interface {
	Validate(StructInfo) error
}

type ValidatorConfig struct {
	fieldNameFunc fieldNameFunc
}

func getFieldName(field *reflect.StructField) string {
	return field.Name
}

func defaultConfig() *ValidatorConfig {
	return &ValidatorConfig{
		fieldNameFunc: getFieldName,
	}
}

// NewValidator returns implementation of Validator.
func NewValidator(opts ...ValidatorOption) (Validator, error) {
	config := defaultConfig()

	for _, opt := range opts {
		err := opt(config)
		if err != nil {
			return nil, err
		}
	}

	return &validator{
		config: *config,
	}, nil
}

func (v *validator) Validate(structPtr interface{}, rules ...StructRule) error {
	rv := reflect.ValueOf(structPtr)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("Only a pointer to struct can be validated.")
	}
	rv = rv.Elem()

	si := &structInfo{
		rv: rv,
	}

	errs := []Error{}

	for _, rule := range rules {
		err := rule.Apply(si)
		if err == nil {
			continue
		}
		if e, ok := err.(Errors); ok {
			errs = append(errs, e...)
			continue
		}
		return err
	}

	if len(errs) == 0 {
		return nil
	}
	return Errors(errs)
}

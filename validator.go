package validation

import (
	"reflect"
)

type fieldNameFunc func(reflect.StructField) string

type validator struct {
	config ValidatorConfig
}
type Validator interface {
	Validate(structPtr interface{}, rules ...StructRule) error
}

type ValidatorConfig struct {
	fieldNameFunc fieldNameFunc
}

func getFieldName(field reflect.StructField) string {
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
	return nil
}

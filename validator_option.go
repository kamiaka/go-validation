package validation

type ValidatorOption func(*ValidatorConfig) error

func FieldNameFunc(fn fieldNameFunc) ValidatorOption {
	return func(config *ValidatorConfig) error {
		config.fieldNameFunc = fn
	}
}

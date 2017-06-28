package validation

// ValidatorOption is option for new validator.
type ValidatorOption func(*validatorConfig) error

// FieldNameFunc set fieldNameFunc to validator config.
func FieldNameFunc(fn fieldNameFunc) ValidatorOption {
	return func(config *validatorConfig) error {
		config.fieldNameFunc = fn
		return nil
	}
}

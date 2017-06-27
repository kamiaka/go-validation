package validation

// ValidatorOption is option for new validator.
type ValidatorOption func(*ValidatorConfig) error

// FieldNameFunc set fieldNameFunc to validator config.
func FieldNameFunc(fn fieldNameFunc) ValidatorOption {
	return func(config *ValidatorConfig) error {
		config.fieldNameFunc = fn
		return nil
	}
}

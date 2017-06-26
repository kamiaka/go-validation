package validation

// Required is validation rule that checks value is not empty.
var Required = &requiredRule{
	format: "%[1]s is requried",
}

type requiredRule struct {
	format string
}

func (r *requiredRule) Validate(f FieldLevel) FieldError {
	if IsEmpty(f.Value()) {
		return &fieldError{
			field:  f,
			format: r.format,
			params: []interface{}{f.Label},
		}
	}

	return nil
}

func (r *requiredRule) SetFormat(format string) *requiredRule {
	return &requiredRule{
		format: format,
	}
}

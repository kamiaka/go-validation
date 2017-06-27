package validation

// Required is validation rule that checks value is not empty.
var Required = &requiredRule{
	format: "%[1]s is requried",
}

type requiredRule struct {
	format string
}

func (r *requiredRule) Apply(f FieldInfo) error {
	if IsEmpty(f.Value()) {
		return newFieldError(f, r.format, f.Label())
	}

	return nil
}

func (r *requiredRule) Format(format string) BuiltInFieldRule {
	return &requiredRule{
		format: format,
	}
}

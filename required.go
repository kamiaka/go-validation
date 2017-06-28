package validation

// Required is validation rule that checks value is not empty.
var Required = &requiredRule{
	format: "%[1]s is requried",
}

type requiredRule struct {
	format string
}

func (r *requiredRule) Apply(f FieldValue) error {
	if IsEmpty(f.Value()) {
		return newError(f, r.format, f.Label())
	}

	return nil
}

func (r *requiredRule) ErrorFormat(format string) BuiltInFieldRule {
	return &requiredRule{
		format: format,
	}
}

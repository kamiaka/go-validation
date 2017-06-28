package validation

type stringRule struct {
	validate func(string) bool
	format   string
}

func NewStringRule(validator func(string) bool, format string) BuiltInFieldRule {
	return &stringRule{
		validate: validator,
		format:   format,
	}
}

func (r *stringRule) ErrorFormat(format string) BuiltInFieldRule {
	return &stringRule{
		validate: r.validate,
		format:   format,
	}
}

func (r *stringRule) Apply(fi FieldValue) error {
	if fi.IsEmpty() {
		return nil
	}

	str, err := EnsureString(fi.Value())
	if err != nil {
		return err
	}

	if !r.validate(str) {
		return newError(fi, r.format, fi.Label())
	}

	return nil
}

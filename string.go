package validation

type stringValidator func(string) bool

type stringRule struct {
	validate stringValidator
	format   string
}

func NewStringRule(validator stringValidator) BuiltInFieldRule {
	return &stringRule{
		validate: validator,
		format:   "%[1]v is not match rule.",
	}
}

func (r *stringRule) ErrorFormat(format string) BuiltInFieldRule {
	return &stringRule{
		validate: r.validate,
		format:   format,
	}
}

func (r *stringRule) Apply(fi FieldInfo) error {
	if fi.IsEmpty() {
		return nil
	}

	str, err := EnsureString(fi.Value())
	if err != nil {
		return err
	}

	if !r.validate(str) {
		return newFieldError(fi, r.format, fi.Label())
	}

	return nil
}

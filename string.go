package validation

type stringRule struct {
	*rule
	validate func(string) bool
}

// NewStringRule ...
func NewStringRule(validator func(string) bool, format string) BuiltInFieldRule {
	return &stringRule{
		rule:     newRule(format),
		validate: validator,
	}
}

func (r *stringRule) SetErrorFormat(f string) BuiltInFieldRule {
	r.rule.format = f
	return r
}

func (r *stringRule) SetParamsMap(f MapParamsFunc) BuiltInFieldRule {
	r.rule.mapParams = f
	return r
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
		return r.newError(fi, fi.Label())
	}

	return nil
}

package validation

// Required message format
const (
	MsgRequiredFormat = "%[1]s is requried"
)

// Required is validation rule that checks value is not empty.
var Required = &requiredRule{
	isEnabled: true,
	format:    MsgRequiredFormat,
}

type requiredRule struct {
	isEnabled bool
	format    string
}

// RequiredWhen rules requires when args is true.
func RequiredWhen(b bool) BuiltInFieldRule {
	return &requiredRule{
		isEnabled: b,
		format:    MsgRequiredFormat,
	}
}

func (r *requiredRule) Apply(f FieldValue) error {
	if r.isEnabled && IsEmpty(f.Value()) {
		return newError(f, r.format, f.Label())
	}

	return nil
}

func (r *requiredRule) ErrorFormat() string {
	return r.format
}

func (r *requiredRule) SetErrorFormat(format string) BuiltInFieldRule {
	return &requiredRule{
		format: format,
	}
}

package validation

// Required message format
const (
	MsgRequiredFormat = "%[1]s is required"
)

// Required is validation rule that checks value is not empty.
var Required = &requiredRule{
	rule:      newRule(MsgRequiredFormat),
	isEnabled: true,
}

type requiredRule struct {
	*rule
	isEnabled bool
}

// RequiredWhen rules requires when args is true.
func RequiredWhen(b bool) BuiltInFieldRule {
	return &requiredRule{
		rule:      newRule(MsgRequiredFormat),
		isEnabled: b,
	}
}

func (r *requiredRule) Apply(f FieldValue) error {
	if r.isEnabled && IsEmpty(f.Value()) {
		return r.newError(f, f.Label())
	}

	return nil
}

func (r *requiredRule) SetErrorFormat(f string) BuiltInFieldRule {
	r.rule.format = f
	return r
}

func (r *requiredRule) SetParamsMap(f MapParamsFunc) BuiltInFieldRule {
	r.rule.mapParams = f
	return r
}

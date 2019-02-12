package validation

// FieldRule is validation rule for validate field.
type FieldRule interface {
	Apply(FieldValue) error
}

// MapParamsFunc map error params before Error.Params()
type MapParamsFunc func([]interface{}) []interface{}

// BuiltInFieldRule is validation rule for validation field that can change error format.
type BuiltInFieldRule interface {
	FieldRule
	ErrorFormat() string
	SetErrorFormat(format string) BuiltInFieldRule
	SetParamsMap(MapParamsFunc) BuiltInFieldRule
}

type fieldRule struct {
	apply func(FieldValue, ErrorFunc) error
}

// ErrorFunc returns validation error.
type ErrorFunc func(message string, params ...interface{}) error

// FieldRuleFunc type is an adapter to allow the use of ordinary functions as field validation rule.
type FieldRuleFunc func(FieldValue, ErrorFunc) error

// Apply calls apply(FieldValue, ErrorFunc)
func (apply FieldRuleFunc) Apply(v FieldValue) error {
	return apply(v, func(message string, params ...interface{}) error {
		return newError(v, message, params...)
	})
}

// StructRule is validation rule for validate struct.
type StructRule interface {
	Apply(Value) error
}

// StructRuleFunc type is an adapter to allow the use of ordinary functions as validation rule.
type StructRuleFunc func(Value, ErrorFunc) error

// Apply calls apply(Value, ErrorFunc)
func (apply StructRuleFunc) Apply(v Value) error {
	return apply(v, func(message string, params ...interface{}) error {
		return newError(v, message, params...)
	})
}

type rule struct {
	format    string
	mapParams MapParamsFunc
}

func defaultParamsMap(ls []interface{}) []interface{} {
	return ls
}

func newRule(format string) *rule {
	return &rule{
		format:    format,
		mapParams: defaultParamsMap,
	}
}

func (r *rule) newError(v Value, params ...interface{}) Errors {
	return newError(v, r.format, r.mapParams(params)...)
}

func (r *rule) ErrorFormat() string {
	return r.format
}

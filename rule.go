package validation

// FieldRule is validation rule for validate field.
type FieldRule interface {
	Apply(FieldInfo) error
}

// BuiltInFieldRule is validation rule for validation field that can change error format.
type BuiltInFieldRule interface {
	FieldRule
	ErrorFormat(format string) BuiltInFieldRule
}

type fieldRule struct {
	apply func(FieldInfo, ErrorFunc) error
}

// ErrorFunc returns validation error.
type ErrorFunc func(message string, params ...interface{}) error

// FieldRuleFunc type is an adapter to allow the use of ordinary functions as field validation rule.
type FieldRuleFunc func(FieldInfo, ErrorFunc) error

// Apply calls apply(FieldInfo, ErrorFunc)
func (apply FieldRuleFunc) Apply(fi FieldInfo) error {
	return apply(fi, func(message string, params ...interface{}) error {
		return &fieldError{
			field:  fi,
			format: message,
			params: params,
		}
	})
}

// StructRule is validation rule for validate struct.
type StructRule interface {
	Apply(StructInfo) error
}

// StructRuleFunc type is an adapter to allow the use of ordinary functions as validation rule.
type StructRuleFunc func(StructInfo, ErrorFunc) error

// Apply calls apply(StructInfo, ErrorFunc)
func (apply StructRuleFunc) Apply(si StructInfo) error {
	return apply(si, func(message string, params ...interface{}) error {
		return newStructError(si, message, params)
	})
}

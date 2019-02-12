package validation

// Error message formats.
const (
	MsgInvalidLengthFormat   = "%[1]v must between %[2]v and %[3]v"
	MsgMinLengthFormat       = "%[1]v must be %[2]v or more"
	MsgMaxLengthFormat       = "%[1]v must be %[2]v or less"
	MsgStringLengthFormat    = "%[1]v must between %[2]v and %[3]v character(s)"
	MsgStringMinLengthFormat = "%[1]v must be %[2]v character(s) or more"
	MsgStringMaxLengthFormat = "%[1]v must be %[2]v character(s) or less"
)

type lengthRule struct {
	*rule
	min    *int
	max    *int
}

func newLengthRule(format string, min, max *int) BuiltInFieldRule {
	return &lengthRule{
		rule:   newRule(format),
		min:    min,
		max:    max,
	}
}

// Length returns a validation rule that checks if a value's length is within the specified range.
func Length(min int, max int) BuiltInFieldRule {
	return newLengthRule(MsgInvalidLengthFormat, &min, &max)
}

// MinLength returns a validation rule that checks if a value's length is greater or equal than specified value.
func MinLength(min int) BuiltInFieldRule {
	return newLengthRule(MsgMinLengthFormat, &min, nil)
}

// MaxLength returns a validation rule that checks if a value's length is less or equal than specified value
func MaxLength(max int) BuiltInFieldRule {
	return newLengthRule(MsgMaxLengthFormat, nil, &max)
}

func reflectStrLength(v reflect.Value) (int, error) {
	return utf8.RuneCountInString(v.String()), nil
}

func newStrLengthRule(format string, min, max *int) BuiltInFieldRule {
	return &lengthRule{
		rule:   newRule(format),
		min:    min,
		max:    max,
	}
}

// StringLength returns a validation rule that checks if a string length is within the specified range.
func StringLength(min int, max int) BuiltInFieldRule {
	return newStrLengthRule(MsgStringLengthFormat, &min, &max)
}

// StringMinLength returns a validation rule that checks if a string length is within the specified range.
func StringMinLength(min int) BuiltInFieldRule {
	return newStrLengthRule(MsgStringMinLengthFormat, &min, nil)
}

// StringMaxLength returns a validation rule that checks if a string length is within the specified range.
func StringMaxLength(max int) BuiltInFieldRule {
	return newStrLengthRule(MsgStringMaxLengthFormat, nil, &max)
}

func (r *lengthRule) Apply(f FieldValue) error {
	if IsEmpty(f.Value()) {
		return nil
	}

	size, err := LengthOfValue(f.Value())
	if err != nil {
		return err
	}
	if r.min == nil {
		if *r.max < size {
			return r.newError(f, f.Label(), *r.max)
		}
	} else if r.max == nil {
		if size < *r.min {
			return r.newError(f, f.Label(), *r.min)
		}
	} else if size < *r.min || *r.max < size {
		return r.newError(f, f.Label(), *r.min, *r.max)
	}

	return nil
}

func (r *lengthRule) SetErrorFormat(f string) BuiltInFieldRule {
	r.rule.format = f
	return r
}

func (r *lengthRule) SetParamsMap(f MapParamsFunc) BuiltInFieldRule {
	r.rule.mapParams = f
	return r
}

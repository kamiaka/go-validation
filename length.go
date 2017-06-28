package validation

const (
	MsgInvalidLength = "%[1]v must between %[2]v and %[3]v"
	MsgMinLength = "%[1]v must be %[2]v or more"
	MsgMaxLength = "%[1]v must be %[2]v or less"
	MsgStringLength = "%[1]v must between %[2]v and %[3]v character(s)"
	MsgStringMinLength = "%[1]v must be %[2]v character(s) or more"
	MsgStringMaxLength = "%[1]v must be %[2]v character(s) or less"
)

type lengthRule struct {
	min    *int
	max    *int
	format string
}

// Length returns a validation rule that checks if a value's length is within the specified range.
func Length(min int, max int) BuiltInFieldRule {
	return &lengthRule{
		min:    &min,
		max:    &max,
		format: MsgInvalidLength,
	}
}

// MinLength returns a validation rule that checks if a value's length is greater or equal than specified value.
func MinLength(min int) BuiltInFieldRule {
	return &lengthRule{
		min:    &min,
		format: MsgMinLength,
	}
}

// MaxLength returns a validation rule that checks if a value's length is less or equal than specified value
func MaxLength(max int) BuiltInFieldRule {
	return &lengthRule{
		max:    &max,
		format: MsgMaxLength,
	}
}

// StringLength returns a validation rule that checks if a string length is within the specified range.
func StringLength(min int, max int) BuiltInFieldRule {
	return Length(min, max).SetErrorFormat(MsgStringLength)
}

// StringMinLength returns a validation rule that checks if a string length is within the specified range.
func StringMinLength(min int) BuiltInFieldRule {
	return MinLength(min).SetErrorFormat(MsgStringMinLength)
}

// StringMaxLength returns a validation rule that checks if a string length is within the specified range.
func StringMaxLength(max int) BuiltInFieldRule {
	return MaxLength(max).SetErrorFormat(MsgStringMaxLength)
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
			return newError(f, r.format, f.Label(), *r.max)
		}
	} else if r.max == nil {
		if size < *r.min {
			return newError(f, r.format, f.Label(), *r.min)
		}
	} else if size < *r.min || *r.max < size {
		return newError(f, r.format, f.Label(), *r.min, *r.max)
	}

	return nil
}

func (r *lengthRule) ErrorFormat() string {
	return r.format
}

func (r *lengthRule) SetErrorFormat(format string) BuiltInFieldRule {
	return &lengthRule{
		min:    r.min,
		max:    r.max,
		format: format,
	}
}

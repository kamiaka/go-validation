package validation

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
		format: "%[1]v must between %[2]v and %[3]v",
	}
}

// MinLength returns a validation rule that checks if a value's length is greater or equal than specified value.
func MinLength(min int) BuiltInFieldRule {
	return &lengthRule{
		min:    &min,
		format: "%[1]v must be %[2]v or more",
	}
}

// MaxLength returns a validation rule that checks if a value's length is less or equal than specified value
func MaxLength(max int) BuiltInFieldRule {
	return &lengthRule{
		max:    &max,
		format: "%[1]v must be %[2]v or less",
	}
}

// StringLength returns a validation rule that checks if a string length is within the specified range.
func StringLength(min int, max int) BuiltInFieldRule {
	return Length(min, max).ErrorFormat("%[1]v must between %[2]v and %[3]v character(s)")
}

// StringMinLength returns a validation rule that checks if a string length is within the specified range.
func StringMinLength(min int) BuiltInFieldRule {
	return MinLength(min).ErrorFormat("%[1]v must be %[2]v character(s) or more")
}

// StringMaxLength returns a validation rule that checks if a string length is within the specified range.
func StringMaxLength(max int) BuiltInFieldRule {
	return MaxLength(max).ErrorFormat("%[1]v must be %[2]v character(s) or less")
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

func (r *lengthRule) ErrorFormat(format string) BuiltInFieldRule {
	return &lengthRule{
		min:    r.min,
		max:    r.max,
		format: format,
	}
}

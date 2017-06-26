package validation

type lengthRule struct {
	min    int
	max    int
	format string
}

func Length(min int, max int) FieldRule {
	return &lengthRule{
		min:    min,
		max:    max,
		format: "%[1]v must between %[2]v and %[3]v",
	}
}

func (r *lengthRule) Validate(f FieldLevel) FieldError {
	if IsEmpty(f.Value()) {
		return nil
	}

	size := 0
	if size < r.min || r.max < size {
		return &fieldError{
			field:  f,
			format: r.format,
			params: []interface{}{f.Label, r.min, r.max},
		}
	}

	return nil
}

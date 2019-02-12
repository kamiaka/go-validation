package validation

import (
	"reflect"
	"unicode/utf8"
)

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
	length func(reflect.Value) (int, error)
}

func newLengthRule(format string, min, max *int) BuiltInFieldRule {
	return &lengthRule{
		rule:   newRule(format),
		min:    min,
		max:    max,
		length: LengthOfValue,
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

// StringLength returns a validation rule that checks if a string length is within the specified range.
func StringLength(min int, max int) BuiltInFieldRule {
	return newLengthRule(MsgStringLengthFormat, &min, &max)
}

// StringMinLength returns a validation rule that checks if a string length is within the specified range.
func StringMinLength(min int) BuiltInFieldRule {
	return newLengthRule(MsgStringMinLengthFormat, &min, nil)
}

// StringMaxLength returns a validation rule that checks if a string length is within the specified range.
func StringMaxLength(max int) BuiltInFieldRule {
	return newLengthRule(MsgStringMaxLengthFormat, nil, &max)
}

func reflectCharLength(v reflect.Value) (int, error) {
	return utf8.RuneCountInString(v.String()), nil
}

// CharLength returns a validation rule that checks if a multibyte string length is within the specified range.
func CharLength(min int, max int) BuiltInFieldRule {
	return newCharLengthRule(MsgStringLengthFormat, &min, &max)
}

// CharMinLength returns a validation rule that checks if a multibyte string length is within the specified range.
func CharMinLength(min int) BuiltInFieldRule {
	return newCharLengthRule(MsgStringMinLengthFormat, &min, nil)
}

// CharMaxLength returns a validation rule that checks if a multibyte string length is within the specified range.
func CharMaxLength(max int) BuiltInFieldRule {
	return newCharLengthRule(MsgStringMaxLengthFormat, nil, &max)
}

func newCharLengthRule(format string, min, max *int) BuiltInFieldRule {
	return &lengthRule{
		rule:   newRule(format),
		min:    min,
		max:    max,
		length: reflectCharLength,
	}
}

func (r *lengthRule) Apply(f FieldValue) error {
	if IsEmpty(f.Value()) {
		return nil
	}

	size, err := r.length(f.Value())
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

package validation

import (
	"reflect"
)

// Error format.
const (
	MsgInFormat = "%[1]v must be a valid value"
)

type inRule struct {
	values []interface{}
	format string
}

// In values rule.
func In(values ...interface{}) BuiltInFieldRule {
	return &inRule{
		values: values,
		format: MsgInFormat,
	}
}

func (r *inRule) ErrorFormat() string {
	return r.format
}

func (r *inRule) SetErrorFormat(format string) BuiltInFieldRule {
	return &inRule{
		values: r.values,
		format: format,
	}
}

func (r *inRule) Apply(f FieldValue) error {
	fv := f.Value()
	if IsEmpty(fv) {
		return nil
	}
	if fv.Kind() == reflect.Ptr {
		fv = fv.Elem()
	}

	value := f.Interface()
	for _, v := range r.values {
		if v == value {
			return nil
		}
	}
	return newError(f, r.format, f.Label())
}

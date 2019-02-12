package validation

import (
	"reflect"
)

// Error format.
const (
	MsgInFormat = "%[1]v must be a valid value"
)

type inRule struct {
	*rule
	values []interface{}
}

// In values rule.
func In(values ...interface{}) BuiltInFieldRule {
	return &inRule{
		rule:   newRule(MsgInFormat),
		values: values,
	}
}

func (r *inRule) SetErrorFormat(f string) BuiltInFieldRule {
	r.rule.format = f
	return r
}

func (r *inRule) SetParamsMap(f MapParamsFunc) BuiltInFieldRule {
	r.rule.mapParams = f
	return r
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
	return r.newError(f, f.Label())
}

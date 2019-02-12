package validation

import (
	"fmt"
	"reflect"
	"time"
)

// Error message format.
const (
	MsgGTEFormat = "%[1]v must be greater than equal %[2]v"
	MsgGTFormat  = "%[1]v must be greater than %[2]v"
	MsgLTEFormat = "%[1]v must be less than equal %[2]v"
	MsgLTFormat  = "%[1]v must be less than %[2]v"
)

const (
	greaterThan = iota
	greaterThanEqual
	lessThan
	lessThanEqual
)

type thresholdRule struct {
	*rule
	threshold interface{}
	operator  int
}

// Min returns validation rule that checks if a value is greater or equal than specified value.
// that is alias of GTE
func Min(min interface{}) BuiltInFieldRule {
	return GTE(min)
}

// Max returns validation rule that checks if a value is less or equal than specified value.
// that is alias of LTE
func Max(max interface{}) BuiltInFieldRule {
	return LTE(max)
}

// GTE returns validation rule that checks if a value is greater or equal than specified value.
func GTE(v interface{}) BuiltInFieldRule {
	return &thresholdRule{
		threshold: v,
		operator:  greaterThanEqual,
		rule:      newRule(MsgGTEFormat),
	}
}

// GT returns validation rule that checks if a value is greater than specified value.
func GT(v interface{}) BuiltInFieldRule {
	return &thresholdRule{
		threshold: v,
		operator:  greaterThan,
		rule:      newRule(MsgGTFormat),
	}
}

// LTE returns validation rule that checks if a value is less or equal than specified value.
func LTE(v interface{}) BuiltInFieldRule {
	return &thresholdRule{
		threshold: v,
		operator:  lessThanEqual,
		rule:      newRule(MsgLTEFormat),
	}
}

// LT returns validation rule that checks if a value is less than specified value.
func LT(v interface{}) BuiltInFieldRule {
	return &thresholdRule{
		threshold: v,
		operator:  lessThan,
		rule:      newRule(MsgLTFormat),
	}
}

func (r *thresholdRule) SetErrorFormat(f string) BuiltInFieldRule {
	r.rule.format = f
	return r
}

func (r *thresholdRule) SetParamsMap(f MapParamsFunc) BuiltInFieldRule {
	r.rule.mapParams = f
	return r
}

func (r *thresholdRule) Apply(f FieldValue) error {
	fv := f.Value()
	if IsEmpty(fv) {
		return nil
	}
	if fv.Kind() == reflect.Ptr {
		fv = fv.Elem()
	}

	tv := reflect.ValueOf(r.threshold)
	switch tv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if r.compareInt(tv.Int(), fv.Int()) {
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if r.compareUint(tv.Uint(), fv.Uint()) {
			return nil
		}
	case reflect.Float32, reflect.Float64:
		if r.compareFloat(tv.Float(), fv.Float()) {
			return nil
		}
	case reflect.Struct:
		t, ok := r.threshold.(time.Time)
		if !ok {
			return fmt.Errorf("threshold rule does not support type %v", tv.Type())
		}
		v, ok := fv.Interface().(time.Time)
		if !ok {
			return fmt.Errorf("cannot convert %v to time.Time", fv.Type())
		}
		if v.IsZero() || r.compareTime(t, v) {
			return nil
		}
	}
	return r.newError(f, f.Label(), r.threshold)
}

func (r *thresholdRule) compareInt(threshold, value int64) bool {
	switch r.operator {
	case greaterThan:
		return value > threshold
	case greaterThanEqual:
		return value >= threshold
	case lessThan:
		return value < threshold
	default:
		return value <= threshold
	}
}

func (r *thresholdRule) compareUint(threshold, value uint64) bool {
	switch r.operator {
	case greaterThan:
		return value > threshold
	case greaterThanEqual:
		return value >= threshold
	case lessThan:
		return value < threshold
	default:
		return value <= threshold
	}
}

func (r *thresholdRule) compareFloat(threshold, value float64) bool {
	switch r.operator {
	case greaterThan:
		return value > threshold
	case greaterThanEqual:
		return value >= threshold
	case lessThan:
		return value < threshold
	default:
		return value <= threshold
	}
}

func (r *thresholdRule) compareTime(threshold, value time.Time) bool {
	switch r.operator {
	case greaterThan:
		return value.After(threshold)
	case greaterThanEqual:
		return value.After(threshold) || value.Equal(threshold)
	case lessThan:
		return value.Before(threshold)
	default:
		return value.Before(threshold) || value.Equal(threshold)
	}
}

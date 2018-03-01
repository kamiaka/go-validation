package validation

import (
	"fmt"
	"reflect"
)

var (
	bytesType = reflect.TypeOf([]byte(nil))
)

func EnsureString(v reflect.Value) (string, error) {
	if v.Kind() == reflect.String {
		return v.String(), nil
	}
	if v.Type() == bytesType {
		return string(v.Interface().([]byte)), nil
	}

	return "", fmt.Errorf("cannot convert %v to string", v.Type())
}

func StringOrBytes(v reflect.Value) (isString bool, str string, isBytes bool, bs []byte) {
	if v.Kind() == reflect.String {
		isString = true
		str = v.String()
	} else if v.Kind() == reflect.Slice && v.Type() == bytesType {
		isBytes = true
		bs = v.Interface().([]byte)
	}
	return
}

// IsEmpty returns whether a value is empty or not.
func IsEmpty(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Invalid:
		return true
	case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Ptr:
		return rv.IsNil()
	case reflect.Interface:
		if rv.IsNil() {
			return true
		}
		return IsEmpty(rv.Elem())
	}
	return false
}

func LengthOfValue(rv reflect.Value) (int, error) {
	switch rv.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array:
		return rv.Len(), nil

	}
	return 0, fmt.Errorf("cannot get length of %s", rv.Kind().String())
}

func ToFloat(v reflect.Value) (float64, error) {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	}
	return 0, fmt.Errorf("cannot convert %v to float64", v.Kind().String())
}

func ToInt(v reflect.Value) (int64, error) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	}
	return 0, fmt.Errorf("cannot convert %v to int64", v.Kind().String())
}

func ToUint(v reflect.Value) (uint64, error) {
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint(), nil
	}
	return 0, fmt.Errorf("cannot convert %v to uint64", v.Kind().String())
}

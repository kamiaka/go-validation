package validation

import (
	"fmt"
	"reflect"
)

// IsEmpty returns whether a value is empty or not.
func IsEmpty(rv reflect.Value) bool {
	switch rv.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.String:
		return rv.String() == ""
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
	return 0, fmt.Errorf("Cannot convert %v to float64.", v.Kind().String())
}

func ToInt(v reflect.Value) (int64, error) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	}
	return 0, fmt.Errorf("Cannot convert %v to int64.", v.Kind().String())
}

func ToUint(v reflect.Value) (uint64, error) {
	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint(), nil
	}
	return 0, fmt.Errorf("Cannot convert %v to uint64.", v.Kind().String())
}

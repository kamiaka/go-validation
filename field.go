package validation

import (
	"fmt"
	"reflect"
)

type fieldValidator struct {
	label    string
	fieldPtr interface{}
	rules    []FieldRule
}

// Field returns field validator.
func Field(label string, fieldPtr interface{}, rules ...FieldRule) StructRule {
	return &fieldValidator{
		label:    label,
		fieldPtr: fieldPtr,
		rules:    rules,
	}
}

func findStructField(structValue reflect.Value, fieldValue reflect.Value) *reflect.StructField {
	ptr := fieldValue.Pointer()
	for i := structValue.NumField() - 1; i >= 0; i-- {
		sf := structValue.Type().Field(i)
		if ptr == structValue.Field(i).UnsafeAddr() {
			if sf.Type == fieldValue.Elem().Type() {
				return &sf
			}
		}
	}
	return nil
}

func (v *fieldValidator) Apply(si StructInfo) error {
	fv := reflect.ValueOf(v.fieldPtr)
	if fv.Kind() != reflect.Ptr {
		return fmt.Errorf("Field %s is not specified as a pointer.", v.label)
	}

	sf := findStructField(si.Value(), fv)
	if sf == nil {
		return fmt.Errorf("Cannot find struct field for %s", v.label)
	}

	field := &fieldInfo{
		label: v.label,
		sf:    sf,
		rv:    fv.Elem(),
		si:    si,
	}
	var errs Errors
	for _, rule := range v.rules {
		err := rule.Apply(field)
		if err == nil {
			continue
		}
		if e, ok := err.(Errors); ok {
			errs = append(errs, e...)
			continue
		}
		return err
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

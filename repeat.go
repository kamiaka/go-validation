package validation

import (
	"fmt"
	"reflect"
)

type arrayValidator struct {
	rules []FieldRule
}

func Repeat(rules ...FieldRule) FieldRule {
	return &arrayValidator{
		rules: rules,
	}
}

func (v *arrayValidator) Apply(f FieldValue) error {
	ls, ok := f.(*value)
	if !ok {
		return fmt.Errorf("cannot convert %v to *value", f.Value().Type())
	}

	fields := []FieldValue{}
	switch ls.rv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < ls.rv.Len(); i++ {
			fields = append(fields, &value{
				label:  ls.label,
				key:    i,
				rv:     ls.rv.Index(i),
				parent: ls,
			})
		}
	case reflect.Map:
		for _, k := range ls.rv.MapKeys() {
			fields = append(fields, &value{
				label:  ls.label,
				key:    k.Interface(),
				rv:     ls.rv.MapIndex(k),
				parent: ls,
			})
		}
	}

	errs := []Error{}
	for _, field := range fields {
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
	}

	if len(errs) == 0 {
		return nil
	}

	return Errors(errs)
}

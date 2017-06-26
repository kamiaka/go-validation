package validation

import "reflect"

type field struct {
	label string
	key   string
	rv    reflect.Value
	sl    StructLevel
}

func (f *field) Label() string {
	return f.label
}

func (f *field) Key() string {
	return f.label
}

func (f *field) Value() interface{} {
	return f.rv.Interface()
}

func (f *field) Parent() interface{} {
	return f.parent
}

func (f *field) Kind() reflect.Kind {
	return f.rv.Kind()
}

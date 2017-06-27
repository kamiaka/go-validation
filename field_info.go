package validation

import "reflect"

type FieldInfo interface {
	Label() string
	Name() string
	Interface() interface{}
	Value() reflect.Value
	IsEmpty() bool
	Parent() StructInfo
}

type fieldInfo struct {
	label     string
	sf        *reflect.StructField
	rv        reflect.Value
	si        StructInfo
	fieldName fieldNameFunc
}

func (f *fieldInfo) Label() string {
	return f.label
}

func (f *fieldInfo) Name() string {
	return f.fieldName(f.sf)
}

func (f *fieldInfo) Interface() interface{} {
	return f.rv.Interface()
}

func (f *fieldInfo) Value() reflect.Value {
	return f.rv
}

func (f *fieldInfo) IsEmpty() bool {
	return IsEmpty(f.Value())
}

func (f *fieldInfo) Parent() StructInfo {
	return f.si
}

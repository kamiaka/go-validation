package validation

import (
	"fmt"
	"reflect"
)

// Value ...
type Value interface {
	Value() reflect.Value
	Interface() interface{}
	Parent() Value
	Name() string
	Namespace() string
}

// FieldValue ...
type FieldValue interface {
	Value
	Label() string
	IsEmpty() bool
}

type value struct {
	label  string
	sf     *reflect.StructField
	key    interface{} // key for slice, array or mapa, if value is is an array, slice or map element.
	rv     reflect.Value
	parent Value
	config *validatorConfig
}

func (v *value) Label() string {
	return v.label
}

func (v *value) Name() string {
	if v.parent == nil {
		return ""
	}
	pv := v.parent.Value()
	if pv.Kind() == reflect.Ptr {
		pv = pv.Elem()
	}

	if pv.Kind() == reflect.Struct {
		return v.config.fieldNameFunc(v.sf)
	}
	return v.parent.Name()
}

func (v *value) Namespace() string {
	if v.parent == nil {
		return ""
	}
	ns := v.parent.Namespace()

	pv := v.parent.Value()
	if pv.Kind() == reflect.Ptr {
		pv = pv.Elem()
	}

	switch pv.Kind() {
	case reflect.Struct:
		if ns == "" {
			return v.Name()
		}
		return fmt.Sprintf("%s.%s", ns, v.Name())
	case reflect.Array, reflect.Slice, reflect.Map:
		return fmt.Sprintf("%s[%#v]", ns, v.key)
	}

	return ""
}

func (v *value) Value() reflect.Value {
	return v.rv
}

func (v *value) Interface() interface{} {
	return v.rv.Interface()
}

func (v *value) Parent() Value {
	return v.parent
}

func (v *value) IsEmpty() bool {
	return IsEmpty(v.rv)
}

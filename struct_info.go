package validation

import (
	"reflect"
)

type StructInfo interface {
	Value() reflect.Value
	Interface() interface{}
	Parent() StructInfo
}

type structInfo struct {
	rv     reflect.Value
	parent StructInfo
}

func (b *structInfo) Value() reflect.Value {
	return b.rv
}

func (b *structInfo) Interface() interface{} {
	return b.rv.Interface()
}

func (b *structInfo) Parent() StructInfo {
	return b.parent
}

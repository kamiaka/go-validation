package validation

import (
	"reflect"
)

type StructLevel interface {
	Value() interface{}
	Parent() StructLevel
}

type structLevel struct {
	rv     reflect.Value
	parent StructLevel
}

func (b *structLevel) Value() interface{} {
	return b.rv.Interface()
}

func (b *structLevel) Parent() StructLevel {
	return b.parent
}

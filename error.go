package validation

import "fmt"

type FieldError interface {
	Error() string
	Field() FieldLevel
	Params() []interface{}
}

type StructError interface {
	Error() string
	Params() []string
}

type fieldError struct {
	field  FieldLevel
	format string
	params []interface{}
}

func (e *fieldError) Error() string {
	return fmt.Sprintf(e.format, e.params...)
}

func (e *fieldError) Field() FieldLevel {
	return e.field
}

func (e *fieldError) Params() []interface{} {
	return e.params
}

func Error(field FieldLevel, format string, params ...interface{}) FieldError {
	return &fieldError{
		field:  field,
		format: format,
		params: params,
	}
}

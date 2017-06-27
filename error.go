package validation

import (
	"bytes"
	"fmt"
)

// Errors is an array of Error.
type Errors []Error

func (e Errors) Error() string {
	buf := bytes.NewBufferString("Validation error.")

	for i := 0; i < len(e); i++ {
		buf.WriteString("\n")
		buf.WriteString(e[i].Error())
	}

	return buf.String()
}

// Error is validation error, that is field or struct error.
type Error interface {
	// Error returns error message.
	Error() string

	// ErrorFormat returns format string for error.
	ErrorFormat() string

	// Params returns validation params for error message.
	Params() []interface{}
}

// FieldError is field error interface. that implements `Error`
type FieldError interface {
	Error
	Field() FieldInfo
}

type fieldError struct {
	field  FieldInfo
	format string
	params []interface{}
}

func (e *fieldError) Error() string {
	return fmt.Sprintf(e.format, e.params...)
}

func (e *fieldError) ErrorFormat() string {
	return e.format
}

func (e *fieldError) Params() []interface{} {
	return e.params
}

func (e *fieldError) Field() FieldInfo {
	return e.field
}

func newFieldError(fi FieldInfo, format string, params ...interface{}) Errors {
	return Errors{
		&fieldError{
			field:  fi,
			format: format,
			params: params,
		},
	}
}

type StructError interface {
	Error
	Struct() StructInfo
}

type structError struct {
	si     StructInfo
	format string
	params []interface{}
}

func newStructError(si StructInfo, format string, params ...interface{}) Errors {
	return Errors{
		&structError{
			si:     si,
			format: format,
			params: params,
		},
	}
}

func (e *structError) Error() string {
	return fmt.Sprintf(e.format, e.params...)
}

func (e *structError) ErrorFormat() string {
	return e.format
}

func (e *structError) Params() []interface{} {
	return e.params
}

func (e *structError) Struct() StructInfo {
	return e.si
}

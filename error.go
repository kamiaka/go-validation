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

	// Value returns field/struct level Value.
	Value() Value
}

type err struct {
	value  Value
	format string
	params []interface{}
}

func newError(v Value, format string, params ...interface{}) Errors {
	return Errors{
		&err{
			value:  v,
			format: format,
			params: params,
		},
	}
}

func (e *err) Error() string {
	return fmt.Sprintf(e.format, e.params...)
}

func (e *err) ErrorFormat() string {
	return e.format
}

func (e *err) Params() []interface{} {
	return e.params
}

func (e *err) Value() Value {
	return e.value
}

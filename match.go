package validation

import (
	"fmt"
	"regexp"
)

// error message format
const (
	MsgMatchFormat = "%[1]v must be in a valid format"
)

type matchRule struct {
	re     *regexp.Regexp
	format string
}

// Match regular expressions
func Match(re *regexp.Regexp) BuiltInFieldRule {
	return &matchRule{
		re:     re,
		format: MsgMatchFormat,
	}
}

func (r *matchRule) ErrorFormat() string {
	return r.format
}

func (r *matchRule) SetErrorFormat(format string) BuiltInFieldRule {
	return &matchRule{
		re:     r.re,
		format: format,
	}
}

func (r *matchRule) Apply(fi FieldValue) error {
	if fi.IsEmpty() {
		return nil
	}

	isStr, str, isBytes, bs := StringOrBytes(fi.Value())
	if isStr {
		if r.re.MatchString(str) {
			return nil
		}
	} else if isBytes {
		if r.re.Match(bs) {
			return nil
		}
	} else {
		return fmt.Errorf("cannot convert %v to string or bytes", fi.Value().Type())
	}

	return newError(fi, r.format, fi.Label())
}

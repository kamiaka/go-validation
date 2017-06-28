package validation

import (
	"fmt"
	"regexp"
)

type matchRule struct {
	re     *regexp.Regexp
	format string
}

func Match(re *regexp.Regexp) BuiltInFieldRule {
	return &matchRule{
		re:     re,
		format: "%[1]v must be in a valid format",
	}
}

func (r *matchRule) ErrorFormat(format string) BuiltInFieldRule {
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

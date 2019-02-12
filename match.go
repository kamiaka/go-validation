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
	*rule
	re *regexp.Regexp
}

// Match regular expressions
func Match(re *regexp.Regexp) BuiltInFieldRule {
	return &matchRule{
		rule: newRule(MsgMatchFormat),
		re:   re,
	}
}

func (r *matchRule) SetErrorFormat(f string) BuiltInFieldRule {
	r.rule.format = f
	return r
}

func (r *matchRule) SetParamsMap(f MapParamsFunc) BuiltInFieldRule {
	r.rule.mapParams = f
	return r
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

	return r.newError(fi, fi.Label())
}

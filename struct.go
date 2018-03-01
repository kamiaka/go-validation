package validation

// DeepStructRuleFunc ...
type DeepStructRuleFunc func(FieldValue) ([]StructRule, error)

// Apply rules
func (f DeepStructRuleFunc) Apply(v FieldValue) error {
	if v.IsEmpty() {
		return nil
	}
	rules, err := f(v)
	if err != nil {
		return err
	}
	var errs []Error
	for _, rule := range rules {
		err := rule.Apply(v)
		if err != nil {
			if e, ok := err.(Errors); ok {
				errs = append(errs, e...)
				continue
			}
			return err
		}
	}
	return Errors(errs)
}

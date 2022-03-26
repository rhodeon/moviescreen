package validator

// Validator
type Validator struct {
	Errors map[string][]string
}

func New() *Validator {
	return &Validator{Errors: map[string][]string{}}
}

// Valid checks if the validator has any errors.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError append the message to the specified validator field.
func (v *Validator) AddError(field string, message string) {
	v.Errors[field] = append(v.Errors[field], message)
}

// Check adds an error to the validator field if the rule fails.
func (v *Validator) Check(rule bool, field string, message string) {
	if !rule {
		v.AddError(field, message)
	}
}

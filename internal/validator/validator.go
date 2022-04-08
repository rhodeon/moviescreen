package validator

import "regexp"

// A Validator is associated with a struct and has its Errors populated by fields
// and messages which fail validation rules.
// Type of the associated struct.
type Validator struct {
	Type   string
	Errors map[string][]string
}

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func New(structType string) *Validator {
	return &Validator{Type: structType, Errors: map[string][]string{}}
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

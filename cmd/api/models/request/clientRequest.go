package request

import "github.com/rhodeon/moviescreen/internal/validator"

// ClientRequest is the generic interface which all request models have to satisfy.
type ClientRequest interface {
	// Validate handles data validation in the request against specified rules.
	// A validator instance which holds all incurred validations is returned.
	Validate() *validator.Validator
}

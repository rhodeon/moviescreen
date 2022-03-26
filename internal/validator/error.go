package validator

type ValidationError struct{}

func (ve *ValidationError) Error() string {
	return "validation errors incurred"
}

func NewError() error {
	return &ValidationError{}
}

package request

import (
	"github.com/rhodeon/moviescreen/internal/validator"
	"unicode/utf8"
)

type UserActivationRequest struct {
	Token *string `json:"token"`
}

const UserActivationFieldToken = "token"

func (u UserActivationRequest) Validate(required []string) *validator.Validator {
	v := validator.New("activate")

	if u.Token == nil {
		v.AddError(UserActivationFieldToken, "must be provided")
	} else if utf8.RuneCountInString(*u.Token) != 26 {
		v.AddError(UserActivationFieldToken, "must have exactly 26 characters")
	}

	return v
}

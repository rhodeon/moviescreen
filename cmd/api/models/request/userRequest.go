package request

import (
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/internal/validator"
	"github.com/rhodeon/moviescreen/internal/validator/rules"
	"strings"
	"unicode/utf8"
)

type UserRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

const (
	UserFieldUsername = "username"
	UserFieldEmail    = "email"
	UserFieldPassword = "password"
)

func (request *UserRequest) ToModel() models.User {
	return models.User{
		Username: *request.Username,
		Email:    *request.Email,
	}
}

func (request UserRequest) Validate(required []string) *validator.Validator {
	v := validator.New("user")

	for _, field := range required {
		switch field {
		case UserFieldUsername:
			v.Check(request.Username != nil, field, "must be provided")

		case UserFieldEmail:
			v.Check(request.Email != nil, field, "must be provided")

		case UserFieldPassword:
			v.Check(request.Password != nil, field, "must be provided")
		}
	}

	if request.Username != nil {
		v.Check(strings.TrimSpace(*request.Username) != "", UserFieldUsername, "must not be blank")
		v.Check(utf8.RuneCountInString(*request.Username) <= 500, UserFieldUsername, "must not have more than 500 characters")
	}

	if request.Email != nil {
		v.Check(strings.TrimSpace(*request.Email) != "", UserFieldEmail, "must not be blank")
		v.Check(rules.MatchesPattern(*request.Email, validator.EmailRX), UserFieldEmail, "must be a valid email address")
	}

	if request.Password != nil {
		v.Check(strings.TrimSpace(*request.Password) != "", UserFieldPassword, "must not be blank")
		v.Check(utf8.RuneCountInString(*request.Password) >= 8, UserFieldPassword, "must have at least 8 characters")
		v.Check(utf8.RuneCountInString(*request.Password) <= 72, UserFieldPassword, "must not have more than 72 characters")
	}

	return v
}

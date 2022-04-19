package mock

import (
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
	"github.com/rhodeon/moviescreen/internal/types"
	"strings"
	"time"
)

var MockDate = time.Date(2022, 4, 10, 10, 00, 00, 00, time.UTC)

var users = []models.User{
	{
		Id:       1,
		Username: "rhodeon",
		Email:    "rhodeon@dev.mail",
		Password: types.Password{
			Plaintext: common.StringLiteralPointer("rhodeonpass"),
			Hash:      []byte("rhodeonhashedpass"),
		},
		Activated: true,
		Version:   1,
		Created:   MockDate,
	},
	{
		Id:       2,
		Username: "ruona",
		Email:    "ruona@mail.com",
		Password: types.Password{
			Plaintext: common.StringLiteralPointer("ruonapass"),
			Hash:      []byte("ruonahashedpass"),
		},
		Activated: false,
		Version:   0,
		Created:   MockDate,
	},
	{
		Id:       3,
		Username: "johndoe",
		Email:    "johndoe@mail.com",
		Password: types.Password{
			Plaintext: common.StringLiteralPointer("johndoepass"),
			Hash:      []byte("johndoehashedpass"),
		},
		Activated: false,
		Version:   0,
		Created:   MockDate,
	},
}

type UserController struct{}

func (u UserController) Register(user *models.User) error {
	for _, u := range users {
		if strings.ToLower(u.Username) == strings.ToLower(user.Username) {
			return repository.ErrDuplicateUsername
		}
		if strings.ToLower(u.Email) == strings.ToLower(user.Email) {
			return repository.ErrDuplicateEmail
		}
	}

	user.Id = len(users) + 1
	user.Version = 1
	user.Created = MockDate
	return nil
}

func (u UserController) GetByEmail(email string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserController) Update(user *models.User) error {
	for i, u := range users {
		if u.Id == user.Id {
			users[i] = *user
			return nil
		}
	}
	return repository.ErrRecordNotFound
}

func (u UserController) GetByToken(plainTextToken string, scope string) (models.User, error) {
	for _, token := range tokens {
		if token.PlainText == plainTextToken && token.Scope == scope {
			for _, user := range users {
				if user.Id == token.UserId {
					return user, nil
				}
			}
		}
	}

	return models.User{}, repository.ErrRecordNotFound
}

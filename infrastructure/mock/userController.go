package mock

import (
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
	"github.com/rhodeon/moviescreen/internal/types"
	"strings"
	"time"
)

type UserController struct {
	Data []models.User
}

// NewUserController creates a UserController pointer with the data being
// a copy of the users slice to avoid persistent modification across tests.
func NewUserController() *UserController {
	newUsers := make([]models.User, len(users))
	copy(newUsers, users)
	return &UserController{Data: newUsers}
}

var MockDate = time.Date(2022, 4, 10, 10, 00, 00, 00, time.UTC)

var users = []models.User{
	{
		Id:       1,
		Username: "rhodeon",
		Email:    "rhodeon@dev.mail",
		Password: types.Password{
			Plaintext: common.StringLiteralPointer("password"),
			Hash:      []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
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
			Plaintext: common.StringLiteralPointer("password"),
			Hash:      []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
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
			Plaintext: common.StringLiteralPointer("password"),
			Hash:      []byte("$2a$10$T.olpluq6ZZAisvfJVuLuOIXnqh/bN.9RCDiEu/tnnCgBqjesMkse.sP49rm"),
		},
		Activated: false,
		Version:   0,
		Created:   MockDate,
	},
}

func (u *UserController) Register(user *models.User) error {
	for _, u := range u.Data {
		if strings.EqualFold(u.Username, user.Username) {
			return repository.ErrDuplicateUsername
		}
		if strings.EqualFold(u.Email, user.Email) {
			return repository.ErrDuplicateEmail
		}
	}

	user.Id = len(u.Data) + 1
	user.Version = 1
	user.Created = MockDate
	return nil
}

func (u *UserController) GetByEmail(email string) (models.User, error) {
	for _, user := range u.Data {
		if user.Email == email {
			return user, nil
		}
	}

	return models.User{}, repository.ErrRecordNotFound
}

func (u *UserController) Update(user *models.User) error {
	for i, savedUser := range u.Data {
		if savedUser.Id == user.Id {
			users[i] = *user
			return nil
		}
	}
	return repository.ErrRecordNotFound
}

func (u *UserController) GetByToken(plainTextToken string, scope string) (models.User, error) {
	for _, token := range tokens {
		if token.PlainText == plainTextToken && token.Scope == scope {
			for _, user := range u.Data {
				if user.Id == token.UserId {
					return user, nil
				}
			}
		}
	}

	return models.User{}, repository.ErrRecordNotFound
}

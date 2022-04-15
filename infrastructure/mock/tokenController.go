package mock

import (
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
	"time"
)

var ActivationExpiry = time.Now().Add(2 * 24 * time.Hour)

var tokens = []models.Token{
	{
		PlainText: "2QRJK3S54HAIUNIHNXEF4WSZSI",
		Hash:      []byte("84828986df43c6289a90a0001d01d2ec4fcbf045429a6bf9fe9275bb21cbaf7c"),
		UserId:    1,
		Scope:     "activation",
		Expires:   ActivationExpiry,
	},
}

type TokenController struct{}

func (t TokenController) New(userId int, scope string, lifetime time.Duration) (models.Token, error) {
	return models.Token{
		PlainText: "token",
		Hash:      []byte("hashedToken"),
		UserId:    userId,
		Scope:     scope,
		Expires:   time.Now().Add(lifetime),
	}, nil
}

func (t TokenController) Insert(token models.Token) error {
	//TODO implement me
	panic("implement me")
}

func (t TokenController) DeleteAllForUser(userId int, scope string) error {
	for i, token := range tokens {
		if token.UserId == userId && token.Scope == scope {
			tokens[i] = models.Token{}
		}
		return nil
	}
	return repository.ErrRecordNotFound
}

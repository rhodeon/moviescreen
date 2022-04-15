package mock

import (
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"github.com/rhodeon/moviescreen/domain/repository"
	"time"
)

var ActivationExpiry = time.Now().Add(2 * 24 * time.Hour)

var tokens = []common.Token{
	{
		PlainText: "2QRJK3S54HAIUNIHNXEF4WSZSI",
		Hash:      []byte("84828986df43c6289a90a0001d01d2ec4fcbf045429a6bf9fe9275bb21cbaf7c"),
		UserId:    1,
		Scope:     "activation",
		Expires:   ActivationExpiry,
	},
}

type TokenController struct{}

func (t TokenController) New(userId int, scope string, lifetime time.Duration) (common.Token, error) {
	return common.Token{
		PlainText: "token",
		Hash:      []byte("hashedToken"),
		UserId:    userId,
		Scope:     scope,
		Expires:   time.Now().Add(lifetime),
	}, nil
}

func (t TokenController) Insert(token common.Token) error {
	//TODO implement me
	panic("implement me")
}

func (t TokenController) DeleteAllForUser(userId int, scope string) error {
	for i, token := range tokens {
		if token.UserId == userId && token.Scope == scope {
			tokens[i] = common.Token{}
		}
		return nil
	}
	return repository.ErrRecordNotFound
}

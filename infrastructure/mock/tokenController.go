package mock

import (
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"time"
)

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
	//TODO implement me
	panic("implement me")
}

package repository

import (
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"time"
)

type TokenRepository interface {
	New(userId int, scope string, lifetime time.Time) (common.Token, error)
	Insert(token common.Token) error
	DeleteAllForUser(userId int, scope string) error
}

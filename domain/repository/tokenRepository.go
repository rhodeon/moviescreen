package repository

import (
	"github.com/rhodeon/moviescreen/domain/models"
	"time"
)

type TokenRepository interface {
	New(userId int, scope string, lifetime time.Duration) (models.Token, error)
	Insert(token models.Token) error
	DeleteAllForUser(userId int, scope string) error
}

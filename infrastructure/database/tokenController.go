package database

import (
	"context"
	"database/sql"
	"github.com/rhodeon/moviescreen/cmd/api/common"
	"time"
)

type TokenController struct {
	Db *sql.DB
}

// New is a shortcut to insert a new token with the given user ID, token scope and lifetime.
func (t TokenController) New(userId int, scope string, lifetime time.Duration) (common.Token, error) {
	token, err := common.GenerateToken(userId, scope, lifetime)
	if err != nil {
		return common.Token{}, err
	}

	err = t.Insert(token)
	if err != nil {
		return common.Token{}, err
	}

	return token, nil
}

// Insert adds a new token to the database
func (t TokenController) Insert(token common.Token) error {
	stmt := `INSERT INTO tokens (hash, user_id, scope, expires)
	VALUES ($1, $2, $3, $4)
`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := t.Db.ExecContext(ctx, stmt, token.Hash, token.UserId, token.Scope, token.Expires)
	return err
}

// DeleteAllForUser removes all expired tokens, and those of a user with the given scope.
func (t TokenController) DeleteAllForUser(userId int, scope string) error {
	stmt := `DELETE FROM tokens
    WHERE (user_id = $1 AND scope = $2) OR expires < now()
`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := t.Db.ExecContext(ctx, stmt, userId, scope)
	return err
}

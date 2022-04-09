package database

import (
	"context"
	"database/sql"
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/domain/repository"
	"strings"
	"time"
)

type UserController struct {
	Db *sql.DB
}

// Register creates a new user updating the details of the inputted user pointer.
func (u UserController) Register(user *models.User) error {
	stmt := `INSERT INTO users (username, email, password_hash) 
	VALUES ($1, $2, $3)
	RETURNING id, version, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := u.Db.QueryRowContext(ctx, stmt, user.Username, user.Email, user.Password.Hash)
	err := row.Scan(&user.Id, &user.Version, &user.Created)

	if err != nil {
		if strings.Contains(err.Error(), "users_username_key") {
			return repository.ErrDuplicateUsername
		}
		if strings.Contains(err.Error(), "users_email_key") {
			return repository.ErrDuplicateEmail
		}
		return err
	}

	return nil
}

func (u UserController) GetByEmail(email string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserController) Update(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

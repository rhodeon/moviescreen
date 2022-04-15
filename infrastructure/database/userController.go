package database

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
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
	stmt := `SELECT id, username, email, password_hash, activated, version, created_at FROM users
	WHERE email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user := models.User{}

	err := u.Db.QueryRowContext(ctx, stmt, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password.Hash,
		&user.Activated,
		&user.Version,
		&user.Created,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.User{}, repository.ErrRecordNotFound

		default:
			return models.User{}, err
		}
	}

	return user, nil
}

// Update replaces the data of the user in the database with those in the passed-in user.
// An "edit conflict" error is returned if the version of the user in the database does not
// match that in the parameter. This is done to prevent data races.
func (u UserController) Update(user *models.User) error {
	stmt := `UPDATE users 
	SET username = $1, email = $2, password_hash = $3, activated = $4, version = version + 1 
	WHERE id = $5 AND version = $6
	RETURNING version, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.Db.QueryRowContext(ctx, stmt, user.Username, user.Email, user.Password.Hash, user.Activated, user.Id, user.Version).Scan(
		&user.Version,
		&user.Created,
	)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "users_username_key"):
			return repository.ErrDuplicateUsername

		case strings.Contains(err.Error(), "users_email_key"):
			return repository.ErrDuplicateEmail

		case errors.Is(err, sql.ErrNoRows):
			return repository.ErrEditConflict

		default:
			return err
		}

	}

	return nil
}

// GetByToken returns the user satisfying both the plain text token and the scope
func (u UserController) GetByToken(plainTextToken string, scope string) (models.User, error) {
	// join user and token tables to check users against the tokens and scopes
	stmt := `SELECT users.id, users.username, users.email, users.password_hash, users.activated, users.version, users.created_at
	FROM users
	INNER JOIN tokens ON users.id = tokens.user_id
	WHERE tokens.hash = $1 AND tokens.scope = $2 AND tokens.expires > $3`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user := models.User{}
	tokenHash := sha256.Sum256([]byte(plainTextToken))

	err := u.Db.QueryRowContext(ctx, stmt, tokenHash[:], scope, time.Now()).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password.Hash,
		&user.Activated,
		&user.Version,
		&user.Created,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.User{}, repository.ErrRecordNotFound
		default:
			return models.User{}, nil
		}
	}

	return user, nil
}

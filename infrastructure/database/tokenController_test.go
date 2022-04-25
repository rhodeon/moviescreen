package database

import (
	"database/sql"
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/infrastructure/mock"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"reflect"
	"testing"
	"time"
)

func TestTokenController_Insert(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}

	db, teardown := newTestDb(t)
	tokenController := TokenController{Db: db}
	defer teardown()

	token := models.Token{
		Hash:    []byte("FA691ECC1287EC358F96FC19250FC10FD6E6DFC064688D8D559EA6C3BD4C1D17"),
		UserId:  2,
		Scope:   models.ScopeActivation,
		Expires: mock.MockDate,
	}

	err := tokenController.Insert(token)

	// verify no error occurred during insertion
	testhelpers.AssertError(t, err, nil)

	// query database to confirm the token was successfully inserted
	stmt := `SELECT hash, user_id, scope, expires FROM tokens 
	where user_id = $1 AND scope = $2`

	fetchedToken := models.Token{}
	row := db.QueryRow(stmt, token.UserId, token.Scope)

	err = row.Scan(&fetchedToken.Hash, &fetchedToken.UserId, &fetchedToken.Scope, &fetchedToken.Expires)
	if err != nil {
		testhelpers.AssertFatalError(t, err)
	}

	// set fetched token time to UTC to match with original token
	fetchedToken.Expires = fetchedToken.Expires.In(time.UTC)

	// compare the original and fetched tokens
	if !reflect.DeepEqual(fetchedToken, token) {
		t.Errorf("\nGot:\t%+v\nWant:\t%+v", fetchedToken, token)
	}
}

func TestTokenController_DeleteAllForUser(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}

	db, teardown := newTestDb(t)
	tokenController := TokenController{Db: db}
	defer teardown()

	userId := 1
	err := tokenController.DeleteAllForUser(userId, models.ScopeActivation)

	testhelpers.AssertError(t, err, nil)

	// query tokens to ensure all qualified ones were deleted
	stmt := `SELECT hash, user_id, scope, expires FROM tokens
	WHERE (user_id = $1 AND scope = $2) OR expires < now()`

	fetchedToken := models.Token{}
	row := tokenController.Db.QueryRow(stmt, userId, models.ScopeActivation)

	err = row.Scan(&fetchedToken.Hash, &fetchedToken.UserId, &fetchedToken.Scope, &fetchedToken.Expires)
	if err != sql.ErrNoRows {
		t.Errorf("\nGot:\t%+v\nWant: an empty token", fetchedToken)
	}
}

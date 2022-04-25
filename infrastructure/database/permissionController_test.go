package database

import (
	"github.com/rhodeon/moviescreen/domain/models"
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"reflect"
	"testing"
)

func TestPermissionController_AddForUser(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}

	db, teardown := newTestDb(t)
	permissionController := PermissionController{Db: db}
	defer teardown()

	userId := 2
	err := permissionController.AddForUser(models.User{Id: userId}, models.PermissionMoviesRead)

	testhelpers.AssertError(t, err, nil)

	// query database to ensure permission code was inserted
	stmt := `SELECT permissions.code 
	FROM permissions
	INNER JOIN users_permissions ON permissions.id = users_permissions.permission_id
	INNER JOIN users ON users_permissions.user_id = users.id
	WHERE users.id = $1`

	var code string
	err = permissionController.Db.QueryRow(stmt, userId).Scan(&code)
	if err != nil {
		testhelpers.AssertFatalError(t, err)
	}

	testhelpers.AssertEqual(t, code, models.PermissionMoviesRead)
}

func TestPermissionController_GetAllForUser(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}

	db, teardown := newTestDb(t)
	permissionController := PermissionController{Db: db}
	defer teardown()

	userId := 1
	permissions, err := permissionController.GetAllForUser(models.User{Id: userId})

	testhelpers.AssertError(t, err, nil)

	wantPermissions := models.Permissions{models.PermissionMoviesRead, models.PermissionMoviesWrite}
	if !reflect.DeepEqual(permissions, wantPermissions) {
		t.Errorf("\nGot:\t%#v\nWant:\t%#v", permissions, wantPermissions)
	}
}

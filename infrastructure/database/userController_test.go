package database

import (
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"testing"
	"time"
)

func TestUserController_Register(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}
	testcases := registerUserTestCases

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			db, teardown := newTestDb(t)
			userController := UserController{Db: db}
			defer teardown()

			err := userController.Register(&tc.user)

			// check error
			testhelpers.AssertError(t, err, tc.wantErr)

			// check user content on success
			if err == nil {
				tc.user.Created = time.Time{}
				testhelpers.AssertStruct(t, tc.user, tc.registeredUser)
			}
		})
	}
}

func TestUserController_GetByEmail(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}
	testcases := getUserByEmailTestCases

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			db, teardown := newTestDb(t)
			userController := UserController{Db: db}
			defer teardown()

			user, err := userController.GetByEmail(tc.email)

			testhelpers.AssertError(t, err, tc.wantErr)

			user.Created = time.Time{}
			testhelpers.AssertStruct(t, user, tc.wantUser)
		})
	}
}

func TestUserController_Update(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}
	testcases := updateUserTestCases

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			db, teardown := newTestDb(t)
			userController := UserController{Db: db}
			defer teardown()

			err := userController.Update(&tc.user)
			testhelpers.AssertError(t, err, tc.wantErr)

			if err == nil {
				tc.user.Created = time.Time{}
				testhelpers.AssertStruct(t, tc.user, tc.updatedUser)
			}
		})
	}
}

func TestUserController_GetByToken(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: skipping integration test")
	}
	testcases := getUserByTokenTestCases

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			db, teardown := newTestDb(t)
			userController := UserController{Db: db}
			defer teardown()

			user, err := userController.GetByToken(tc.plaintTextToken, tc.scope)
			testhelpers.AssertError(t, err, tc.wantErr)

			user.Created = time.Time{}
			testhelpers.AssertStruct(t, user, tc.wantUser)
		})
	}
}

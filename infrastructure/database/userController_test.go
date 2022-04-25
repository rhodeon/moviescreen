package database

import (
	"github.com/rhodeon/moviescreen/internal/testhelpers"
	"reflect"
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
				if !reflect.DeepEqual(tc.user, tc.registeredUser) {
					t.Errorf("\nGot:\t%+v\nWant:\t%+v", tc.user, tc.registeredUser)
				}
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
			if !reflect.DeepEqual(user, tc.wantUser) {
				t.Errorf("\nGot:\t%+v\nWant:\t%+v", user, tc.wantUser)
			}
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
				if !reflect.DeepEqual(tc.user, tc.updatedUser) {
					t.Errorf("\nGot:\t%+v\nWant:\t%+v", tc.user, tc.updatedUser)
				}
			}
		})
	}
}

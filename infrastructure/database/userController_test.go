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

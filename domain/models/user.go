package models

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/internal/types"
	"reflect"
	"time"
)

type User struct {
	Id        int
	Username  string
	Email     string
	Password  types.Password
	Activated bool
	Version   int
	Created   time.Time
}

// AnonymousUser is the user model to be used if no authentication
// token is sent as part of a request.
var AnonymousUser = User{}

func (user User) ToResponse() response.UserResponse {
	return response.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Version:   user.Version,
		Activated: user.Activated,
		Created:   user.Created,
	}
}

func (user User) IsAnonymous() bool {
	return reflect.DeepEqual(user, AnonymousUser)
}

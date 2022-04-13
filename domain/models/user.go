package models

import (
	"github.com/rhodeon/moviescreen/cmd/api/models/response"
	"github.com/rhodeon/moviescreen/internal/types"
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

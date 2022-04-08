package response

import "time"

type User struct {
	Id        int       `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Activated bool      `json:"activated,omitempty"`
	Created   time.Time `json:"created,omitempty"`
}

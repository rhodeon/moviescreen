package response

import "time"

type UserResponse struct {
	Id       int       `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
	Version  int       `json:"version,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

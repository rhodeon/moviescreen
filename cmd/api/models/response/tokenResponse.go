package response

import "time"

type TokenResponse struct {
	PlainText string    `json:"token"`
	Expires   time.Time `json:"expires"`
}

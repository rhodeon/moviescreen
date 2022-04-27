package mock

import (
	"github.com/rhodeon/moviescreen/domain/models"
	"time"
)

type TokenController struct {
	Data []models.Token
}

// NewTokenController creates a TokenController pointer with the data being
// a copy of the tokens slice to avoid persistent modification across tests.
func NewTokenController() *TokenController {
	newTokens := make([]models.Token, len(tokens))
	copy(newTokens, tokens)
	return &TokenController{Data: newTokens}
}

var ActivationExpiry = time.Now().Add(2 * 24 * time.Hour)
var AuthenticationBaseDate = time.Date(2023, 4, 10, 10, 00, 00, 00, time.UTC)

var tokens = []models.Token{
	{
		PlainText: "2QRJK3S54HAIUNIHNXEF4WSZSI",
		Hash:      []byte("84828986df43c6289a90a0001d01d2ec4fcbf045429a6bf9fe9275bb21cbaf7c"),
		UserId:    1,
		Scope:     models.ScopeActivation,
		Expires:   ActivationExpiry,
	},
	{
		PlainText: "2QRJK3S54HAIUNIHNXEF4WSZSI",
		Hash:      []byte("84828986df43c6289a90a0001d01d2ec4fcbf045429a6bf9fe9275bb21cbaf7c"),
		UserId:    1,
		Scope:     models.ScopeAuthentication,
		Expires:   ActivationExpiry,
	},
	{
		PlainText: "2QRJK3S54HAIUNIHNXEF4WSZSI",
		Hash:      []byte("84828986df43c6289a90a0001d01d2ec4fcbf045429a6bf9fe9275bb21cbaf7c"),
		UserId:    1,
		Scope:     models.ScopePasswordReset,
		Expires:   ActivationExpiry,
	},
}

func (t TokenController) New(userId int, scope string, lifetime time.Duration) (models.Token, error) {
	return models.Token{
		PlainText: "token",
		Hash:      []byte("hashedToken"),
		UserId:    userId,
		Scope:     scope,
		Expires:   AuthenticationBaseDate.Add(lifetime),
	}, nil
}

func (t TokenController) Insert(models.Token) error {
	return nil
}

func (t TokenController) DeleteAllForUser(userId int, scope string) error {
	for i, token := range tokens {
		if token.UserId == userId && token.Scope == scope {
			tokens[i] = models.Token{}
		}
	}
	return nil
}

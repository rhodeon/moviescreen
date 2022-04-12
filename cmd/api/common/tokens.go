package common

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

const (
	ScopeActivation = "activation"
)

// Token represents authentication tokens used to verify users.
type Token struct {
	PlainText string
	Hash      []byte
	UserId    int
	Scope     string
	Expires   time.Time
}

func GenerateToken(userId int, scope string, lifetime time.Time) (Token, error) {
	token := Token{
		UserId:  userId,
		Scope:   scope,
		Expires: lifetime,
	}

	// generate random token and hash
	randomBytes := make([]byte, 16)

	// fill randomBytes with data from the OS CSPRNG
	_, err := rand.Read(randomBytes)
	if err != nil {
		return Token{}, err
	}

	// derive and set the plaintext token
	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	// hash the plaintext token
	hash := sha256.Sum256([]byte(token.PlainText))
	token.Hash = hash[:]

	return token, nil
}

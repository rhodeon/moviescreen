package types

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Plaintext *string
	Hash      []byte
}

// Set populates the plaintext hash and the plaintext of the password
// with their appropriate representations from the plaintextPassword.
func (p *Password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.Plaintext = &plaintextPassword
	p.Hash = hash
	return nil
}

// Matches compares the hash of the password to the plaintextPassword.
func (p *Password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(plaintextPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

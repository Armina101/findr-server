package enc

import (
	"errors"
	"runtime"

	"github.com/alexedwards/argon2id"
)

// params: Argon2id hashing parameters for generating hashed password
var params = &argon2id.Params{
	Memory:      64 * 1024,
	Iterations:  5,
	Parallelism: uint8(runtime.NumCPU()),
	SaltLength:  32,
	KeyLength:   32,
}

// CreateHash generates a cryptographic key from the given password using Argon2id.
func CreateHash(password string) (string, error) {
	if password == "" {
		return "", errors.New("invalid input password from user")
	}

	hashedString, err := argon2id.CreateHash(password, params)
	if err != nil || hashedString == "" {
		return "", errors.New("unable to generate hash key from the password")
	}
	return hashedString, nil
}

// VerifyPassword checks and verifies the hashed password and input password.
func VerifyPassword(hashPassword, password string) (bool, error) {
	if password == "" {
		return false, errors.New("invalid input password from user")
	}

	ok, err := argon2id.ComparePasswordAndHash(password, hashPassword)
	if err != nil {
		if errors.Is(err, argon2id.ErrInvalidHash) {
			return false, errors.New("invalid login details")
		}
		return false, errors.New("invalid login details")
	}
	return ok, nil
}

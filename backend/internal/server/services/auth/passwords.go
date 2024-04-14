package auth

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/argon2"
)

func ComparePassword(password, salt, hash string) bool {
	calculatedHash, _, err := HashPassword(password, salt)
	if err != nil {
		return false
	}

	return calculatedHash == hash
}

func HashPassword(password string, salt string) (string, string, error) {
	saltBytes := make([]byte, 32)
	var err error
	if len(salt) > 0 {
		saltBytes, err = hex.DecodeString(salt)
		if err != nil {
			return "", "", err
		}
	} else {
		_, err = rand.Read(saltBytes)
		if err != nil {
			return "", "", err
		}
		salt = hex.EncodeToString(saltBytes)
	}

	// Generate hash
	hash := argon2.IDKey([]byte(password), saltBytes, 1, 64*1024, 4, 32)
	// Return the generated hash and salt used for storage.

	return hex.EncodeToString(hash), salt, nil
}

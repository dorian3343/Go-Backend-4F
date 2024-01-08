package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(input string) (string, error) {
	// Hash the input string using SHA-256
	hasher := sha256.New()
	_, err := hasher.Write([]byte(input))
	if err != nil {
		return "", err
	}

	// Get the hashed result and convert it to a hexadecimal string
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)

	return hashedString, nil
}

func GenerateUserID(login string) (string, error) {
	// Hash the login string
	hashedString, err := HashString(login)
	if err != nil {
		return "", err
	}

	// Concatenate the login and hashed string to form the user ID
	id := login + "-" + hashedString

	return id, nil
}

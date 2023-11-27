package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) (string, error) {
	size := (length * 6) / 8
	if (length*6)%8 != 0 {
		size++
	}

	randomBytes := make([]byte, size)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.RawURLEncoding.EncodeToString(randomBytes)

	randomString = randomString[:length]

	return randomString, nil
}

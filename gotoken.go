package gotoken

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

// GenerateBytes returns n random bytes.
func GenerateBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// GenerateHex returns a random hexadecimal string containing n bytes.
func GenerateHex(n int) (string, error) {
	bytes, err := GenerateBytes(n)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GenerateURLSafe returns a random URL-safe base64 string containing n bytes.
func GenerateURLSafe(n int) (string, error) {
	bytes, err := GenerateBytes(n)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

package gcrypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	/* #nosec */
	"crypto/md5"
)

// HMACSHA256 returns a SHA-256 of the given message signed by the secret.
func HMACSHA256(secret string, message string) (string, error) {
	if secret == "" {
		return "", errors.New("secret must not be empty")
	}

	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// SHA256 returns a SHA-256 of the given message.
func SHA256(message string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// MD5 returns the MD5 checksum of a given message.
func MD5(message string) (string, error) {
	/* #nosec */
	h := md5.New()
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

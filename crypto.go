package gcrypto 

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// HMACSHA256 returns a SHA-256 of the given message signed by the secret.
func HMACSHA256(secret string, message string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 returns a SHA-256 of the given message.
func SHA256(message string) string {
	h := sha256.New()
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5 returns the MD5 checksum of a given message.
func MD5(message string) string {
	h := md5.New()
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

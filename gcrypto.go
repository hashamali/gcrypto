package gcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	/* #nosec */
	"crypto/md5"
)

// GenerateRandomKey will generate a new random key.
func GenerateRandomKey(len int) ([]byte, error) {
	key := make([]byte, len)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

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

// AES256Encrypt will run AES-256 excryption on the provided message.
func AES256Encrypt(message []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("unable to generate cipher with key: %s", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("unable to generate block cipher: %s", err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("unable to generate nonce: %s", err.Error())
	}

	return gcm.Seal(nonce, nonce, message, nil), nil
}

// AES256Decrypt will run AES-256 decryption on the provided data.
func AES256Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("unable to generate cipher with key: %s", err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("unable to generate block cipher: %s", err.Error())
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	message, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to decrypt data: %s", err.Error())
	}

	return message, nil
}

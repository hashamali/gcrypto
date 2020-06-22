package gcrypto

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomKey(t *testing.T) {
	keys := map[string]bool{}
	for i := 0; i < 1000; i++ {
		key, err := GenerateRandomKey(32)
		stringKey := string(key)
		require.NoError(t, err)
		_, ok := keys[stringKey]
		require.False(t, ok)
		keys[stringKey] = true
	}
}

func TestHMACSHA256(t *testing.T) {
	_, err := HMACSHA256("", "")
	require.Error(t, err)

	secret := "secret"
	digest, err := HMACSHA256(secret, "input1")
	require.NoError(t, err)
	require.Equal(t, "b985131d6c4a56d40dece6e6413bf4b29abe48c0431d265f0ceec89a52e0f6b8", digest)

	digest, err = HMACSHA256(secret, "input2")
	require.NoError(t, err)
	require.Equal(t, "3e23a4b3f4246e49fb896778297c4d600a2c5430392a1201af8f20617a070005", digest)
}

func TestSHA256(t *testing.T) {
	digest, err := SHA256("input1")
	require.NoError(t, err)
	require.Equal(t, strings.ToLower("1EA06586B18E8FCE1B923EFF26FD8252F617F0EFD4E49820E8E9BEE0614E5792"), digest)

	digest, err = SHA256("input2")
	require.NoError(t, err)
	require.Equal(t, strings.ToLower("124D8541FF3D7A18B95432BDFBECD86816B86C8265BFF44EF629765AFB25F06B"), digest)
}

func TestMD5(t *testing.T) {
	digest, err := MD5("input1")
	require.NoError(t, err)
	require.Equal(t, strings.ToLower("da853c5826b798b82320e42024d97837"), digest)

	digest, err = MD5("input2")
	require.NoError(t, err)
	require.Equal(t, strings.ToLower("3eaa25d43fac6e39a12c3936942b72c8"), digest)
}

func TestAES256(t *testing.T) {
	// Error on invalid keys.
	_, err := AES256Encrypt([]byte(""), []byte("abc"))
	require.Error(t, err)

	_, err = AES256Decrypt([]byte(""), []byte("abc"))
	require.Error(t, err)

	// Process various message bodies.
	messages := []string{
		"",
		"alsdfkjlaskjdf",
		"1111111111111111111111",
		"12983470	9283470198723498q273 aakjdfhakbndfkl pqaeuwhf kajhdf",
	}

	key, err := GenerateRandomKey(32)
	require.NoError(t, err)

	for _, message := range messages {
		encrypted, err := AES256Encrypt([]byte(message), key)
		require.NoError(t, err)
		decrypted, err := AES256Decrypt(encrypted, key)
		require.NoError(t, err)
		require.Equal(t, message, string(decrypted))
	}
}

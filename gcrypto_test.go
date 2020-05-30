package gcrypto

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHMACSHA256(t *testing.T) {
	secret := "secret"
	digest := HMACSHA256(secret, "input1")
	assert.Equal(t, "b985131d6c4a56d40dece6e6413bf4b29abe48c0431d265f0ceec89a52e0f6b8", digest)

	digest = HMACSHA256(secret, "input2")
	assert.Equal(t, "3e23a4b3f4246e49fb896778297c4d600a2c5430392a1201af8f20617a070005", digest)
}

func TestSHA256(t *testing.T) {
	digest := SHA256("input1")
	assert.Equal(t, strings.ToLower("1EA06586B18E8FCE1B923EFF26FD8252F617F0EFD4E49820E8E9BEE0614E5792"), digest)

	digest = SHA256("input2")
	assert.Equal(t, strings.ToLower("124D8541FF3D7A18B95432BDFBECD86816B86C8265BFF44EF629765AFB25F06B"), digest)
}

func TestMD5(t *testing.T) {
	digest := MD5("input1")
	assert.Equal(t, strings.ToLower("da853c5826b798b82320e42024d97837"), digest)

	digest = MD5("input2")
	assert.Equal(t, strings.ToLower("3eaa25d43fac6e39a12c3936942b72c8"), digest)
}

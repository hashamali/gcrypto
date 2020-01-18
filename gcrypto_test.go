package gcrypto

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256(t *testing.T) {
	sha1 := SHA256("input1")
	assert.Equal(t, strings.ToLower("1EA06586B18E8FCE1B923EFF26FD8252F617F0EFD4E49820E8E9BEE0614E5792"), sha1)

	sha2 := SHA256("input2")
	assert.Equal(t, strings.ToLower("124D8541FF3D7A18B95432BDFBECD86816B86C8265BFF44EF629765AFB25F06B"), sha2)
}

func TestMD5(t *testing.T) {
	md51 := MD5("input1")
	assert.Equal(t, strings.ToLower("da853c5826b798b82320e42024d97837"), md51)

	md52 := MD5("input2")
	assert.Equal(t, strings.ToLower("3eaa25d43fac6e39a12c3936942b72c8"), md52)
}

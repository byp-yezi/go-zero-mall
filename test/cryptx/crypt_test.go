package cryptx

import (
	"go-zero-mall/common/cryptx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordCrypt(t *testing.T) {
	output := cryptx.PasswordEncrypt("123", "789")
	assert.Equal(t, "ebea6eb394c06d2800241cecbaa231f7b09edc65bec40cc4b429e8d62b1b5e23", output)
}

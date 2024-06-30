package cryptx

import (
	"go-zero-mall/common/cryptx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordCrypt(t *testing.T) {
	output := cryptx.PasswordEncrypt("HWVOFkGgPTryzICwd7qnJaZR9KQ2i8xe", "123456")
	assert.Equal(t, "ef8364020560a703cfc7aebebcd0b62d1bfea7d4a841eb8964cfbcda2ba85dd5", output)
}

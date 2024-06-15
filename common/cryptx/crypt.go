package cryptx

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(salt, password string) string {
	dk, err := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	if err != nil {
		logx.Errorf("PasswordEncrypt Failed!!!: %+v", err)
	}
	return fmt.Sprintf("%x", string(dk))

}

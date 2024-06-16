package code

import (
	"go-zero-mall/common/codex"
)

var (
	RegisterNameEmpty = codex.New(20001, "注册名字不能为空")
	UserExist         = codex.New(20002, "该用户已存在")
)

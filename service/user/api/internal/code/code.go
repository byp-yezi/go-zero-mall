package code

import (
	"go-zero-mall/common/codex"
	"go-zero-mall/common/errx"
)

var (
	RegisterMobileEmpty   = codex.New(10001, "注册手机号不能为空")
	RegisterPasswdEmpty   = codex.New(10002, "注册密码不能为空")
	VerificationCodeEmpty = codex.New(10003, "验证码不能为空")
	MobileHasRegistered   = codex.New(10004, "手机号已经注册")
	LoginMobileEmpty      = codex.New(10005, "登录手机号不能为空")
	GenderOnlyOneOrTwo    = codex.New(10006, "Gender只能为0或者1")
)

// 用户模块
var (
	ErrUserAlreadyRegister = errx.NewErrCodeMsg(20001, "该用户已注册-api")
	ErrGenderOnlyOneOrTwo  = errx.NewErrCodeMsg(20002, "Gender的值只能是0或1")
	ErrUserIdNotExist      = errx.NewErrCodeMsg(20003, "请求的用户id和登录id不同")
)

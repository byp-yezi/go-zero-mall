package codex

import (
	"context"
	"go-zero-mall/common/codex/types"
	"net/http"

	"github.com/pkg/errors"
)

func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)

	return http.StatusOK, types.Status{
		Code: int32(code.Code()),
		Message: code.Message(),
	}
}

func CodeFromError(err error) CodeX {
	// 首先，使用 `errors.Cause` 来获取错误的根本原因。`errors.Cause` 是来自 `github.com/pkg/errors` 包的一个函数，
    // 它用于获取嵌套错误链中的最根本的那个错误。
	err = errors.Cause(err)

	// 检查这个错误是否可以转换为 `CodeX` 类型。如果可以，则直接返回这个代码。
	if code, ok := err.(CodeX); ok {
		return code
	}

	// 使用 `switch` 语句根据错误类型来返回相应的 `CodeX`。
	switch err {
	case context.Canceled:
		return Canceled // 如果错误是 `context.Canceled`，则返回 `Canceled` 常量。
	case context.DeadlineExceeded:
		return Deadline // 如果错误是 `context.DeadlineExceeded`，则返回 `Deadline` 常量。
	}

	// 如果错误类型不在上述范围内，默认返回 `ServerErr` 常量。
	return ServerErr
}
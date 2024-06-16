package user

import (
	"net/http"

	"go-zero-mall/common/result"
	"go-zero-mall/service/user/api/internal/logic/user"
	"go-zero-mall/service/user/api/internal/svc"
	"go-zero-mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户登录
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			result.ParamErrorResult(r.Context(), w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			result.AuthHttpResult(r.Context(), r, w, resp, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

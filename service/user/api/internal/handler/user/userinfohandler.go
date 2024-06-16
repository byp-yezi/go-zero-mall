package user

import (
	"net/http"

	"go-zero-mall/service/user/api/internal/logic/user"
	"go-zero-mall/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户信息
func UserinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			// result.AuthHttpResult(r.Context(), r, w, resp, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

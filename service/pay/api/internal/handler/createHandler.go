package handler

import (
	"net/http"

	"go-zero-mall/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-mall/service/pay/api/internal/logic"
	"go-zero-mall/service/pay/api/internal/svc"
	"go-zero-mall/service/pay/api/internal/types"
)

// 支付创建
func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			result.ParamErrorResult(r.Context(), w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			result.HttpResult(r.Context(), r, w, resp, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

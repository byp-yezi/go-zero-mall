package handler

import (
	"net/http"

	"go-zero-mall/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-mall/service/order/api/internal/logic"
	"go-zero-mall/service/order/api/internal/svc"
	"go-zero-mall/service/order/api/internal/types"
)

// 订单更新
func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			result.ParamErrorResult(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(&req)
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			result.HttpResult(r.Context(), r, w, resp, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

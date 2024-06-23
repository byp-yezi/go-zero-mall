package handler

import (
	"net/http"

	"go-zero-mall/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-mall/service/product/api/internal/logic"
	"go-zero-mall/service/product/api/internal/svc"
	"go-zero-mall/service/product/api/internal/types"
)

// 产品列表
func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRequest
		if err := httpx.Parse(r, &req); err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			result.ParamErrorResult(r.Context(), w, err)
			return
		}

		l := logic.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		if err != nil {
			// httpx.ErrorCtx(r.Context(), w, err)
			result.HttpResult(r.Context(), r, w, resp, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

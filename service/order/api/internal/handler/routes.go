// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-mall/service/order/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 订单创建
				Method:  http.MethodPost,
				Path:    "/api/order/create",
				Handler: CreateHandler(serverCtx),
			},
			{
				// 订单详情
				Method:  http.MethodPost,
				Path:    "/api/order/detail",
				Handler: DetailHandler(serverCtx),
			},
			{
				// 订单列表
				Method:  http.MethodPost,
				Path:    "/api/order/list",
				Handler: ListHandler(serverCtx),
			},
			{
				// 订单删除
				Method:  http.MethodPost,
				Path:    "/api/order/remove",
				Handler: RemoveHandler(serverCtx),
			},
			{
				// 订单更新
				Method:  http.MethodPost,
				Path:    "/api/order/update",
				Handler: UpdateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/order/v1"),
	)
}

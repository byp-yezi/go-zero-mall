// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "go-zero-mall/service/user/api/internal/handler/user"
	"go-zero-mall/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 用户登录
				Method:  http.MethodPost,
				Path:    "/api/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				// 用户注册
				Method:  http.MethodPost,
				Path:    "/api/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 用户信息
				Method:  http.MethodPost,
				Path:    "/api/userinfo",
				Handler: user.UserinfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/user/v1"),
	)
}

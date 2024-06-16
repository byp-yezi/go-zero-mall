package svc

import (
	"go-zero-mall/service/user/api/internal/config"
	"go-zero-mall/service/user/rpc/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	userRpc := zrpc.MustNewClient(c.UserRpc /*, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor())*/)
	return &ServiceContext{
		Config:  c,
		UserRpc: userservice.NewUserService(userRpc),
	}
}

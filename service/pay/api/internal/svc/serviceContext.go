package svc

import (
	"go-zero-mall/service/pay/api/internal/config"
	"go-zero-mall/service/pay/rpc/payservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	PayRpc payservice.PayService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: payservice.NewPayService(zrpc.MustNewClient(c.PayRpc)),
	}
}

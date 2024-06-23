package svc

import (
	"go-zero-mall/service/order/api/internal/config"
	"go-zero-mall/service/order/rpc/ordersrv"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	OrderRpc ordersrv.Ordersrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		OrderRpc: ordersrv.NewOrdersrv(zrpc.MustNewClient(c.OrderRpc)),
	}
}

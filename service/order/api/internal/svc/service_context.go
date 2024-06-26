package svc

import (
	"go-zero-mall/service/order/api/internal/config"
	"go-zero-mall/service/order/rpc/ordersrv"
	"go-zero-mall/service/product/rpc/productservice"
	"go-zero-mall/service/user/rpc/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	OrderRpc ordersrv.Ordersrv
	ProductRpc productservice.ProductService
	UserRpc userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		OrderRpc: ordersrv.NewOrdersrv(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: productservice.NewProductService(zrpc.MustNewClient(c.ProductRpc)),
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
	}
}

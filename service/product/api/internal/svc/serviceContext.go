package svc

import (
	"go-zero-mall/service/product/api/internal/config"
	"go-zero-mall/service/product/rpc/productservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	ProductRpc productservice.ProductService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ProductRpc: productservice.NewProductService(zrpc.MustNewClient(c.ProductRpc)),
	}
}

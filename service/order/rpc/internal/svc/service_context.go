package svc

import (
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/order/rpc/internal/config"
	"go-zero-mall/service/product/rpc/productservice"
	"go-zero-mall/service/user/rpc/userservice"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	OrderModel model.OrderModel
	UserRpc userservice.UserService
	ProductRpc productservice.ProductService
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productservice.NewProductService(zrpc.MustNewClient(c.ProductRpc)),
	}
}

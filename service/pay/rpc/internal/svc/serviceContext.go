package svc

import (
	"go-zero-mall/service/order/rpc/ordersrv"
	"go-zero-mall/service/pay/model"
	"go-zero-mall/service/pay/rpc/internal/config"
	"go-zero-mall/service/user/rpc/userservice"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	PayModel model.PayModel
	UserRpc userservice.UserService
	OrderRpc ordersrv.Ordersrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: ordersrv.NewOrdersrv(zrpc.MustNewClient(c.OrderRpc)),
	}
}

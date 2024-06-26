package main

import (
	"flag"
	"fmt"

	"go-zero-mall/common/interceptors/rpcserver"
	"go-zero-mall/service/order/rpc/internal/config"
	"go-zero-mall/service/order/rpc/internal/server"
	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order.RegisterOrdersrvServer(grpcServer, server.NewOrdersrvServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 自定义拦截器
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	logx.DisableStat()
	s.Start()
}

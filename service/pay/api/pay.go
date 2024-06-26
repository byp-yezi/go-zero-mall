package main

import (
	"flag"
	"fmt"

	"go-zero-mall/common/result"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-mall/service/pay/api/internal/config"
	"go-zero-mall/service/pay/api/internal/handler"
	"go-zero-mall/service/pay/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/pay.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(result.JwtUnauthorizedResult))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	logx.DisableStat()
	server.Start()
}

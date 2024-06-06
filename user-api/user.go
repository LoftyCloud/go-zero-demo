package main

import (
	"flag"
	"fmt"

	"go-zero-demo/common/middleware"
	"go-zero-demo/user-api/internal/config"
	"go-zero-demo/user-api/internal/handler"
	"go-zero-demo/user-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 配置全局中间件
	server.Use(middleware.NewGlobalMiddleware().Handle)

	ctx := svc.NewServiceContext(c) // 将config传到ctx中
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

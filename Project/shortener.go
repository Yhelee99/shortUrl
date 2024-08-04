package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"Project/internal/config"
	"Project/internal/errorx"
	"Project/internal/handler"
	"Project/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/shortener-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 注册自定义错误处理方法
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, any) {
		// 判断错误类型
		switch e := err.(type) {
		case *errorx.ErrorCode: // 为自定义错误，返回响应
			return http.StatusOK, e.Response()
		default: // 不为自定义错误，默认返回
			return http.StatusInternalServerError, err
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

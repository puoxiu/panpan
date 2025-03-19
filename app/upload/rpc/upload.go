package main

import (
	"flag"
	"fmt"

	"PanPan/app/upload/rpc/internal/config"
	"PanPan/app/upload/rpc/internal/server"
	"PanPan/app/upload/rpc/internal/svc"
	"PanPan/app/upload/rpc/types/upload"
	"PanPan/common/logs/zapx"
	"PanPan/common/response/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest/httpx"

	"PanPan/common/errorx"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/upload.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		upload.RegisterUploadServer(grpcServer, server.NewUploadServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 指定rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	// 自定义错误
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

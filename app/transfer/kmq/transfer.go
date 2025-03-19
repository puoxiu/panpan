package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"

	"PanPan/common/errorx"
	"PanPan/common/logs/zapx"
	"PanPan/app/transfer/kmq/internal/config"
	"PanPan/app/transfer/kmq/internal/service"
)

var configFile = flag.String("f", "etc/transfer.yaml", "the etc file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	srv := service.NewService(c)
	queue := kq.MustNewQueue(c.KqConsumerConf, kq.WithHandle(srv.Consume))
	defer queue.Stop()

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
	fmt.Println("seckill started!!!")
	queue.Start()
}

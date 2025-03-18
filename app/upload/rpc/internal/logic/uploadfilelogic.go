package logic

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"PanPan/app/upload/model"
	"PanPan/app/upload/rpc/internal/svc"
	"PanPan/app/upload/rpc/types/upload"
	"PanPan/common/batcher"
	"PanPan/utils"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	batcher *batcher.Batcher
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	uf :=  &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}

	options := batcher.Options{
		Worker:   5,
		Buffer:   100,
		Size:     100,
		Interval: 1 * time.Second,
	}

	b := batcher.New(options)
	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % options.Worker
	}

	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*model.NewUserFile
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*model.NewUserFile))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			logx.Errorf("Batcher.Do json.Marshal msgs: %v error: %v", msgs, err)
		}
		if err = uf.svcCtx.KqPusherClient.Push(ctx, string(kd)); err != nil {
			logx.Errorf("KafkaPusher.Push kd: %s error: %v", string(kd), err)
		}
	}

	uf.batcher = b
	uf.batcher.Start()

	return uf
}

func (l *UploadFileLogic) UploadFile(in *upload.UploadFileReq) (*upload.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &upload.CommonResp{}, nil
}


// COSUpload : 上传文件到腾讯云对象存储
func (l *UploadFileLogic) COSUpload(filePath string, file *os.File) error {
	Url := l.svcCtx.Config.TencentCOS.Url
	SecretId := l.svcCtx.Config.TencentCOS.SecretId
	SecretKey := l.svcCtx.Config.TencentCOS.SecretKey

	err := utils.TencentCOSUpload(Url, SecretId, SecretKey, filePath, file)
	if err != nil {
		logc.Error(l.ctx, "上传文件失败")
		return err
	}
	return nil
}
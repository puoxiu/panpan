package logic

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"PanPan/app/upload/model"
	"PanPan/app/upload/rpc/internal/svc"
	"PanPan/app/upload/rpc/types/upload"
	"PanPan/common/errorx"
	"github.com/pkg/errors"


	"github.com/zeromicro/go-zero/core/logx"
)

type InitialMultipartUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitialMultipartUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitialMultipartUploadLogic {
	return &InitialMultipartUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InitialMultipartUpload 分块上传第一步：初始化分块上传
func (l *InitialMultipartUploadLogic) InitialMultipartUpload(in *upload.InitialMultipartUploadReq) (*upload.InitialMultipartUploadResp, error) {
	// todo: add your logic here and delete this line
	// 生成分块上传的初始化信息
	upInfo := model.MultipartUploadInfo{
		FileSha1: in.FileSha1,
		FileSize: in.FileSize,
		// 每次上传都会有一个唯一的id,根据userid+时间戳
		UploadID:   strconv.FormatInt(in.UserId, 10) + fmt.Sprintf("%x", time.Now().UnixNano()),
		ChunkSize:  100 * 1024 * 1024, // 每个分块大小：100MB
		ChunkCount: int(math.Ceil(float64(in.FileSize) / (100 * 1024 * 1024))),
	}
	// 将分块的信息写入redis
	if err := l.svcCtx.Rdb.HSet(l.ctx, "multipart_" + upInfo.UploadID, "filesha1_", upInfo.FileSha1).Err(); err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("redis写入错误"), "redis写入错误 err:%v", err)
	}
	if err := l.svcCtx.Rdb.HSet(l.ctx, "multipart_"+upInfo.UploadID, "filesize_", upInfo.FileSize).Err(); err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("redis写入错误"), "redis写入错误 err:%v", err)

	}
	if err := l.svcCtx.Rdb.HSet(l.ctx, "multipart_"+upInfo.UploadID, "chunkcount_", upInfo.ChunkCount).Err(); err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("redis写入错误"), "redis写入错误 err:%v", err)

	}

	// 返回信息给前端
	return &upload.InitialMultipartUploadResp{
		FileSha1:   upInfo.FileSha1,
		FileSize:   upInfo.FileSize,
		UploadID:   upInfo.UploadID,
		ChunkSize:  int64(upInfo.ChunkSize),
		ChunkCount: int64(upInfo.ChunkCount),
		UserId:     in.UserId,
	}, nil
}

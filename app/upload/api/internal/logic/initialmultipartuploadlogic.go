package logic

import (
	"context"

	"PanPan/app/upload/api/internal/svc"
	"PanPan/app/upload/api/internal/types"
	"PanPan/app/upload/rpc/types/upload"

	"PanPan/common/errorx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type InitialMultipartUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 初始化分块上传文件
func NewInitialMultipartUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitialMultipartUploadLogic {
	return &InitialMultipartUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// InitialMultipartUpload 初始化分块上传文件

// 1. 发送分块上传请求，传入文件的sha1值和文件大小
// 2. 生成分块上传的初始化信息，得到uploadID、分块大小、分块数量信息，服务端会将这些信息写入redis
// 3. 返回初始化信息给前端，等待前端上传分块

func (l *InitialMultipartUploadLogic) InitialMultipartUpload(req *types.InitialMultipartUploadReq) (resp *types.InitialMultipartUploadResp, err error) {
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return nil, errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}

	cnt, err := l.svcCtx.Rpc.InitialMultipartUpload(l.ctx, &upload.InitialMultipartUploadReq{
		UserId:   userId,
		FileSha1: req.FileSha1,
		FileSize: req.FileSize,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.InitialMultipartUploadResp{
		FileSha1:   cnt.FileSha1,
		FileSize:   cnt.FileSize,
		UploadID:   cnt.UploadID,
		ChunkSize:  cnt.ChunkSize,
		ChunkCount: cnt.ChunkCount,
		UserId:     cnt.UserId,
	}, nil
}

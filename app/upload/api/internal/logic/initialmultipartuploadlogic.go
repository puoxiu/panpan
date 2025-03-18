package logic

import (
	"context"

	"PanPan/app/upload/api/internal/svc"
	"PanPan/app/upload/api/internal/types"

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

func (l *InitialMultipartUploadLogic) InitialMultipartUpload(req *types.InitialMultipartUploadReq) (resp *types.InitialMultipartUploadResp, err error) {
	// todo: add your logic here and delete this line

	return
}

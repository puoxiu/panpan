package logic

import (
	"context"

	"PanPan/app/upload/rpc/internal/svc"
	"PanPan/app/upload/rpc/types/upload"

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

func (l *InitialMultipartUploadLogic) InitialMultipartUpload(in *upload.InitialMultipartUploadReq) (*upload.InitialMultipartUploadResp, error) {
	// todo: add your logic here and delete this line

	return &upload.InitialMultipartUploadResp{}, nil
}

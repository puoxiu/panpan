package logic

import (
	"context"

	"PanPan/app/upload/rpc/internal/svc"
	"PanPan/app/upload/rpc/types/upload"

	"github.com/zeromicro/go-zero/core/logx"
)

type FastUploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFastUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FastUploadFileLogic {
	return &FastUploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FastUploadFileLogic) FastUploadFile(in *upload.FastUploadFileReq) (*upload.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &upload.CommonResp{}, nil
}

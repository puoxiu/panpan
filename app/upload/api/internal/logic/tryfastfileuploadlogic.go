package logic

import (
	"context"

	"PanPan/app/upload/api/internal/svc"
	"PanPan/app/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TryFastFileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 秒传文件
func NewTryFastFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TryFastFileUploadLogic {
	return &TryFastFileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TryFastFileUploadLogic) TryFastFileUpload(req *types.TryFastUploadReq) error {
	// todo: add your logic here and delete this line

	return nil
}

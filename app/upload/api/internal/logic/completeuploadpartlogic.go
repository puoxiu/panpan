package logic

import (
	"context"

	"PanPan/app/upload/api/internal/svc"
	"PanPan/app/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteUploadPartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 合并分块文件并上传到相应存储
func NewCompleteUploadPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteUploadPartLogic {
	return &CompleteUploadPartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompleteUploadPartLogic) CompleteUploadPart(req *types.CompleteUploadPartReq) error {
	// todo: add your logic here and delete this line

	return nil
}

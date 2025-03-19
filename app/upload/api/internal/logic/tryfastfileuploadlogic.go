package logic

import (
	"context"

	"PanPan/app/upload/api/internal/svc"
	"PanPan/app/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/pkg/errors"
	"PanPan/common/errorx"
	"PanPan/app/upload/rpc/types/upload"
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

// TryFastFileUpload 秒传文件
func (l *TryFastFileUploadLogic) TryFastFileUpload(req *types.TryFastUploadReq) error {
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	
	_, err := l.svcCtx.Rpc.FastUploadFile(l.ctx, &upload.FastUploadFileReq{
		UserId:   userId,
		FileSha1: req.FileSha1,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}
	return nil
}

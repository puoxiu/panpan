package logic

import (
	"context"

	"PanPan/app/download/api/internal/svc"
	"PanPan/app/download/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadMinioLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Minio下载文件
func NewDownloadMinioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadMinioLogic {
	return &DownloadMinioLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadMinioLogic) DownloadMinio(req *types.DownloadMinioReq) error {
	// todo: add your logic here and delete this line

	return nil
}

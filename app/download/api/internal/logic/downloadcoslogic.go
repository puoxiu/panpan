package logic

import (
	"context"

	"PanPan/app/download/api/internal/svc"
	"PanPan/app/download/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadCOSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// COS下载文件
func NewDownloadCOSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadCOSLogic {
	return &DownloadCOSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadCOSLogic) DownloadCOS(req *types.DownloadCOSReq) error {
	// todo: add your logic here and delete this line

	return nil
}

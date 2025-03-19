package logic

import (
	"context"

	"PanPan/app/upload/rpc/internal/svc"
	"PanPan/app/upload/rpc/types/upload"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/pkg/errors"
	"PanPan/common/errorx"
	"strconv"
)

type UploadPartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPartLogic {
	return &UploadPartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UploadPart: 分块上传第二步：上传分块
func (l *UploadPartLogic) UploadPart(in *upload.UploadPartReq) (*upload.CommonResp, error) {
	// todo: add your logic here and delete this line
	// 将hset置为1表示已经完成该分块的上传(api模块)
	if err := l.svcCtx.Rdb.HSet(l.ctx, "multipart_"+in.UploadID, "checkindex_"+strconv.FormatInt(in.ChunkIndex, 10), 1).Err(); err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError("redis分块上传check错误"), "redis分块上传check错误 err:%v", err)
	}
	return &upload.CommonResp{}, nil
}

package logic

import (
	"context"
	"net/http"

	"PanPan/app/upload/api/internal/svc"
	"PanPan/app/upload/api/internal/types"
	"PanPan/app/upload/rpc/types/upload"

	"PanPan/common/errorx"
	"os"
	"path"
	"strconv"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传分块文件
func NewUploadPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPartLogic {
	return &UploadPartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UploadPart 上传分块文件(在initMultipartUpload得到分块信息后，上传分块文件)
func (l *UploadPartLogic) UploadPart(req *types.UploadPartReq, w http.ResponseWriter, r *http.Request) error {
	res := l.svcCtx.Rdb.HGet(l.ctx, "multipart_" + req.UploadID, "checkindex_"+strconv.FormatInt(req.ChunkIndex, 10))
	if res.Err() != nil || res.Val() == "1" {
		// 上传过了
		return errors.Wrapf(errorx.NewDefaultError("redis分块已经上传"), "redis分块已经上传,UploadID：%v, ChunkIndex:%v, err:%v", req.UploadID, req.ChunkIndex, res.Err())
	}

	filepath := l.svcCtx.Config.FileLocalPath + req.UploadID + "/" + strconv.FormatInt(req.ChunkIndex, 10)
	err := os.MkdirAll(path.Dir(filepath), 0744)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "make文件夹错误 err:%v", err)
	}

	fd, err := os.Create(filepath)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "creat文件错误 err:%v", err)

	}
	defer fd.Close()
	// 创建一个1MB大小的缓冲区
	buf := make([]byte, 1 * 1024 * 1024)
    for {
        // 从请求的Body中读取数据到缓冲区
        n, err := r.Body.Read(buf)
        if n > 0 {
            // 将缓冲区中的数据写入分块文件
            _, writeErr := fd.Write(buf[:n])
            if writeErr != nil {
                return errors.Wrapf(errorx.NewDefaultError(writeErr.Error()), "写入文件错误 err:%v", writeErr)
            }
        }
        // 读完退出循环
        if err != nil {
            break
        }
	}

	// 调用rpc更新redis关于分块文件状态
	_, err = l.svcCtx.Rpc.UploadPart(l.ctx, &upload.UploadPartReq{
		UploadID:   req.UploadID,
		ChunkIndex: req.ChunkIndex,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}

	return nil
}
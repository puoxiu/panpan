package logic

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"PanPan/app/upload/api/internal/svc"
	"PanPan/app/upload/api/internal/types"
	"PanPan/app/upload/model"
	"PanPan/app/upload/rpc/types/upload"
	"PanPan/common/errorx"
	"PanPan/utils"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传文件
func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FileUpload: 上传文件
func (l *FileUploadLogic) FileUpload(req *types.FileUploadReq, w http.ResponseWriter, r *http.Request) error {
	// 从context中获取用户id，jwt middleware
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}

	// 接收文件 暂存到本地
	file, head, err := r.FormFile("file")
	if err != nil {
		fmt.Println("文件上传失败", err)
		return errors.Wrapf(errorx.NewDefaultError("文件上传失败"), "文件上传失败 err:%v", err)
	}
	defer file.Close()

	// 计算文件hash
	file.Seek(0, 0)
	fileSha1 := utils.FileSha1(file)
	file.Seek(0, 0)

	err = os.MkdirAll(l.svcCtx.Config.FileLocalPath+fileSha1, 0755)
	if err != nil {
		log.Fatal(err)
	}

	fileMeta := model.File{
		FileName: head.Filename,
		FileSize: head.Size,
		// 存储路径，sha1 + name
		FileAddr:   l.svcCtx.Config.FileLocalPath + fileSha1 + "/" + head.Filename,
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		FileSha1:   fileSha1,
	}

	newFile, err := os.Create(fileMeta.FileAddr)
	if err != nil {
		fmt.Println("文件创建失败", err)
		return errors.Wrapf(errorx.NewDefaultError("文件创建失败"), "文件创建失败 err:%v", err)
	}
	defer newFile.Close()

	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Println("文件写入失败", err)
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "io.copy 文件失败 err:%v", err)
	}

	// 调用rpc服务，写入数据
	_, err = l.svcCtx.Rpc.UploadFile(l.ctx, &upload.UploadFileReq{
		UserId:          userId,
		FileSha1:        fileMeta.FileSha1,
		FileSize:        fileMeta.FileSize,
		FileName:        fileMeta.FileName,
		FileAddr:        fileMeta.FileAddr,
		CreateTime:      timestamppb.New(fileMeta.CreateTime),
		UpdateTime:      timestamppb.New(fileMeta.UpdateTime),
		DeleteTime:      timestamppb.New(fileMeta.DeleteTime.Time),
		Status:          fileMeta.Status,
		CurrentStoreType: req.CurrentStoreType,
	})
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "rpc.UploadFile err:%v", err)
	}

	return nil
}

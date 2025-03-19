package logic

import (
	"context"
	"fmt"
	"time"

	"PanPan/app/upload/model"
	"PanPan/app/upload/rpc/internal/svc"
	"PanPan/app/upload/rpc/types/upload"
	"PanPan/common/errorx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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

// FastUploadFile: 用于处理秒传请求
func (l *FastUploadFileLogic) FastUploadFile(in *upload.FastUploadFileReq) (*upload.CommonResp, error) {
	file := model.File{}
	if fasterr := l.svcCtx.MysqlDb.Model(&model.File{}).Where("file_sha1 = ?", in.FileSha1).First(&file).Error; fasterr == nil {
		// file查到记录,触发秒传，直接写入userfile
		fmt.Println("触发秒传")
		newUserFile := model.UserFile{
			UserId:     in.UserId,
			FileSha1:   file.FileSha1,
			FileSize:   file.FileSize,
			FileName:   file.FileName,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			Status:     file.Status,
		}
		// 判断用户是不是重新上传文件
		if err := l.svcCtx.MysqlDb.Model(&model.UserFile{}).Where("file_sha1 = ? and user_id = ?", in.FileSha1, in.UserId).First(&model.UserFile{}).Error; err != gorm.ErrRecordNotFound {
			fmt.Println("已经存在该文件")
			return nil, errors.Wrapf(errorx.NewDefaultError("已经存在该文件"), "userid:%v 已经存在该文件", in.UserId)

		} else {
			err = l.svcCtx.MysqlDb.Create(&newUserFile).Error
			if err != nil {
				return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "秒传写入userfile表失败 err:%v", err)
			}
		}
	} else if fasterr == gorm.ErrRecordNotFound {
		// filesha1值不存在，无法触发秒传
		return nil, errors.Wrapf(errorx.NewDefaultError("秒传请求失败"), "秒传请求失败 user_id:%v filesha1:%v", in.UserId, in.FileSha1)
	}
	return &upload.CommonResp{
		Code: 0,
		Message:  "秒传成功",
	}, nil
}


package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"PanPan/app/upload/model"
	"PanPan/app/upload/rpc/internal/svc"
	"PanPan/app/upload/rpc/types/upload"
	"PanPan/common/batcher"
	"PanPan/common/conf"
	"PanPan/common/errorx"
	"PanPan/utils"

	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// 用于存储文件的bucket
const bucketName = "userfile"

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	batcher *batcher.Batcher
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	uf :=  &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}

	// 下面代码为了创建一个批处理器，用于批量处理
	options := batcher.Options{
		Worker:   5,
		Buffer:   100,
		Size:     100,
		Interval: 1 * time.Second,
	}

	b := batcher.New(options)
	b.Sharding = func(key string) int {
		// 根据用户id进行分片， 保证同一个用户的数据在同一个批处理中
		// strconv.ParseInt 将string转为int64 10进制
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % options.Worker
	}

	// Do 是批处理器的处理函数，将数据发送到kafka
	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*model.NewUserFile
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*model.NewUserFile))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			logx.Errorf("Batcher.Do json.Marshal msgs: %v error: %v", msgs, err)
		}
		if err = uf.svcCtx.KqPusherClient.Push(ctx, string(kd)); err != nil {
			logx.Errorf("KafkaPusher.Push kd: %s error: %v", string(kd), err)
		}
	}

	uf.batcher = b
	uf.batcher.Start()

	return uf
}

func (l *UploadFileLogic) UploadFile(in *upload.UploadFileReq) (*upload.CommonResp, error) {
	file, err := os.Open(in.FileAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 根据类型上传文件
	switch in.CurrentStoreType {
		case int64(conf.StoreLocal):
			// 上传到本地, 当前已经上传到本地，不需要再次上传
		case int64(conf.StoreCOS):
			// 上传到腾讯云对象存储
			log.Println("开始写入cos")
			cosPath := "/cos/" + in.FileSha1 + "/" + in.FileName
			err = l.COSUpload(cosPath, file)
			if err != nil {
				return nil, errors.Wrapf(errorx.NewDefaultError("上传文件失败 err:"+err.Error()), "上传文件到COS错误 err:%v", err)
			}
			in.FileAddr = cosPath
		case int64(conf.StoreMinio):
			// 上传到Minio集群
			log.Println("开始写入minio")
			minioPath := "/minio/" + in.FileSha1 + "/" + in.FileName
			_, err := l.svcCtx.MinioDb.PutObject(context.TODO(), bucketName, minioPath, file, -1, minio.PutObjectOptions{})
			if err != nil {
				log.Println(err)
				return nil, errors.Wrapf(errorx.NewDefaultError("上传文件失败 err:"+err.Error()), "上传文件到minio错误 err : %v", err)
			}
			//改变存储路径
			in.FileAddr = minioPath
		default:
			return nil, errors.Wrapf(errorx.NewDefaultError("其他方式暂不支持"), "其他方式暂不支持")
	}

	uf := model.UserFile{
		Id:         0,
		UserId:     in.UserId,
		FileSha1:   in.FileSha1,
		FileSize:   in.FileSize,
		FileName:   in.FileName,
		CreateTime: time.Unix(in.CreateTime.GetSeconds(), 0),
		UpdateTime: time.Unix(in.UpdateTime.GetSeconds(), 0),
		DeleteTime: sql.NullTime{Time: time.Unix(in.DeleteTime.GetSeconds(), 0), Valid: true},
		Status:     in.Status,
	}
	userFile := model.NewUserFile{
		UserFile: uf,
		FileAddr: in.FileAddr,
	}

	// 将file元数据交给kafka进行异步处理
	//strconv.FormatInt 将int64转为string 10进制
	err = l.batcher.Add(strconv.FormatInt(in.UserId, 10), userFile)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(40003, errorx.ErrKafkaUserFileMeta+err.Error()), "kafka异步UserFileMeta失败 err:%v", err)
	}

	return &upload.CommonResp{
		Code:   0,
		Message:    "上传成功!",
	}, nil
}


// COSUpload : 上传文件到腾讯云对象存储
func (l *UploadFileLogic) COSUpload(filePath string, file *os.File) error {
	Url := l.svcCtx.Config.TencentCOS.Url
	SecretId := l.svcCtx.Config.TencentCOS.SecretId
	SecretKey := l.svcCtx.Config.TencentCOS.SecretKey

	err := utils.TencentCOSUpload(Url, SecretId, SecretKey, filePath, file)
	if err != nil {
		logc.Error(l.ctx, "上传文件失败")
		return err
	}
	return nil
}
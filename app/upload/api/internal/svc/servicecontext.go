package svc

import (
	"PanPan/app/upload/api/internal/config"
	"PanPan/app/upload/api/internal/middleware"
	"PanPan/app/upload/model"
	"PanPan/app/upload/rpc/uploadclient"
	"PanPan/common/init_db"

	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	JWT    rest.Middleware

	UserFileModel  model.UserFileModel
	FileModel      model.FileModel
	Rdb            *redis.ClusterClient
	MysqlDb        *gorm.DB
	MinioDb        *minio.Client
	KqPusherClient *kq.Pusher
	Rpc            uploadclient.Upload
}

func NewServiceContext(c config.Config) *ServiceContext {
	// mysql
	mysqlDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	mysqlDb.AutoMigrate(&model.UserFile{}, &model.File{})
	coon := sqlx.NewMysql(c.MysqlCluster.DataSource)

	// redis
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.RedisClusters...)
	redisDb := init_db.InitRedis(rc)

	// minio
	minioDb := init_db.InitMinio(c.MinioCluster.Endpoint, c.MinioCluster.AccessKey, c.MinioCluster.SecretKey)
	return &ServiceContext{
		Config: c,
		JWT:    middleware.NewJWTMiddleware().Handle,

		MysqlDb:        mysqlDb,
		FileModel:      model.NewFileModel(coon, c.CacheRedis),
		UserFileModel:  model.NewUserFileModel(coon, c.CacheRedis),
		Rdb:            redisDb,
		MinioDb:        minioDb,
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		Rpc:            uploadclient.NewUpload(zrpc.MustNewClient(c.Rpc)),
	}
}

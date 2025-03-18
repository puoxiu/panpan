package svc

import (
	"PanPan/app/upload/model"
	"PanPan/app/upload/rpc/internal/config"
	"PanPan/common/init_db"

	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	UserFileModel  model.UserFileModel
	FileModel      model.FileModel
	Rdb            *redis.ClusterClient
	MysqlDb        *gorm.DB
	MinioDb        *minio.Client
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	// mysql
	mysqlDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	mysqlDb.AutoMigrate(&model.UserFile{}, &model.File{})
	conn := sqlx.NewMysql(c.MysqlCluster.DataSource)


	// redis
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.RedisClusters...)
	rdb := init_db.InitRedis(rc)

	// minio
	minioDb := init_db.InitMinio(c.MinioCluster.Endpoint, c.MinioCluster.AccessKey, c.MinioCluster.SecretKey)

	return &ServiceContext{
		Config: c,

		MysqlDb:        mysqlDb,
		FileModel:      model.NewFileModel(conn, c.CacheRedis),
		UserFileModel:  model.NewUserFileModel(conn, c.CacheRedis),
		Rdb:            rdb,
		MinioDb:        minioDb,
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}

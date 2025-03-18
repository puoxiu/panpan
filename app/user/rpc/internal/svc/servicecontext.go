package svc

import (
	"PanPan/app/user/model"
	"PanPan/app/user/rpc/internal/config"
	"PanPan/common/init_db"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
	Rdb *redis.ClusterClient
	MasterDb *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.MysqlCluster.DataSource)
	masterDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	// 自动表迁移
	masterDb.AutoMigrate(&model.User{}, &model.UserAuth{})

	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.RedisClusters...)
	redisDb := init_db.InitRedis(rc)
	
	return &ServiceContext{
		Config:    c,
		MasterDb:  masterDb,
		UserModel: model.NewUserModel(coon, c.CacheRedis),
		Rdb:       redisDb,
	}
}

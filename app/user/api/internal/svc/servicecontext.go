package svc

import (
	"PanPan/app/user/api/internal/config"
	"PanPan/app/user/api/internal/middleware"
	"PanPan/app/user/model"
	"PanPan/app/user/rpc/userclient"
	"PanPan/common/init_db"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	JWT    rest.Middleware

	Rpc           userclient.User
	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
	Rdb           *redis.ClusterClient
	MysqlDb       *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// mysql
	coon := sqlx.NewMysql(c.MysqlCluster.DataSource)
	masterDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	masterDb.AutoMigrate(&model.User{}, &model.UserAuth{})

	// redis
	rc := make([]string, 1)
	rc = append(rc, c.RedisCluster.RedisClusters...)
	redisDb := init_db.InitRedis(rc)

	return &ServiceContext{
		Config: c,
		JWT:    middleware.NewJWTMiddleware().Handle,
		// 
		Rpc:           userclient.NewUser(zrpc.MustNewClient(c.Rpc)),
		MysqlDb:       masterDb,
		UserModel:     model.NewUserModel(coon, c.CacheRedis),
		UserAuthModel: model.NewUserAuthModel(coon, c.CacheRedis),
		Rdb:           redisDb,
	}
}

package config

import (
	"github.com/zeromicro/go-zero/rest"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CacheRedis cache.CacheConf

	Rpc          zrpc.RpcClientConf

	MysqlCluster struct {
		DataSource string
	}

	Github     struct {
		ClientID     string
		RedirectUrl  string
		ClientSecret string
	}

	RedisCluster struct {
		RedisClusters []string
	}
}

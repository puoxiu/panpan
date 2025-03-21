package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	CacheRedis   cache.CacheConf
	Rpc          zrpc.RpcClientConf

	MysqlCluster struct {
		DataSource string
	}

	RedisCluster struct {
		RedisClusters []string
	}

	MinioCluster struct {
		Endpoint  string
		AccessKey string
		SecretKey string
	}

	TencentCOS struct {
		Url       string
		SecretId  string
		SecretKey string
	}

	KqPusherConf struct {
		Brokers []string
		Topic   string
	}

	FileLocalPath string
}

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	MysqlCluster struct {
		DataSource string
	}
	CacheRedis   cache.CacheConf

	RedisCluster struct {
		RedisClusters []string
	}

	// 腾讯云对象存储
	TencentCOS struct {
		Url       string
		SecretId  string
		SecretKey string
	}

	// MinIO集群
	MinioCluster struct {
		Endpoint  string
		AccessKey string
		SecretKey string
	}

	// kakfka
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
}

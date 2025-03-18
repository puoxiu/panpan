package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// MysqlCluster mysql集群配置
	MysqlCluster struct {
		DataSource string
	}

	// CacheRedis 缓存配置
	CacheRedis cache.CacheConf

	// Sms 短信配置
	Credential struct {
		SecretId  string
		SecretKey string
	}

	// RedisCluster redis集群配置
	RedisCluster struct {
		RedisClusters []string 
	}
}

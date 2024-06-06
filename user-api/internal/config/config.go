package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

// 映射配置文件user.yaml中的配置项
type Config struct {
	// api和rpc公用的配置参数
	rest.RestConf // 日志信息在这个预定义结构体中被读取

	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
}

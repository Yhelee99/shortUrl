package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	ShortUrlDb struct {
		DSN string
	}

	SequenceDb struct {
		DSN string
	}

	BaseString string

	BlackList []string

	Domain string

	CacheRedis cache.CacheConf
}

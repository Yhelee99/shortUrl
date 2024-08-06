package svc

import (
	sequence "Project/Sequence"
	"Project/internal/config"
	"Project/model"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	ShortUrlDb model.ShortUrlMapModel

	Sequence  sequence.Sequence
	BlackList map[string]struct{} // 定义一个map,k为string,v为空结构体(目的：不占空间)
	Filter    bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化Mysql连接
	conn := sqlx.NewMysql(c.ShortUrlDb.DSN)

	// 把黑名单数组存到图中
	m := make(map[string]struct{}, len(c.BlackList))
	for _, v := range c.BlackList {
		m[v] = struct{}{}
	}

	// 初始化布隆过滤器
	rds := redis.New(c.CacheRedis[0].Host)
	filter := bloom.New(rds, "bloom filter", 20*(1<<20)) //byte 计算公式，20*数据量

	return &ServiceContext{
		Config:     c,
		ShortUrlDb: model.NewShortUrlMapModel(conn, c.CacheRedis),
		Sequence:   sequence.NewSeqMysql(c.SequenceDb.DSN),
		BlackList:  m,
		Filter:     *filter,
	}
}

func loadDataToBloomFliter(m model.ShortUrlMapModel) {
	sqlx.
}

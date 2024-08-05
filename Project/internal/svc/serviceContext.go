package svc

import (
	sequence "Project/Sequence"
	"Project/internal/config"
	"Project/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	ShortUrlDb model.ShortUrlMapModel

	Sequence  sequence.Sequence
	BlackList map[string]struct{} // 定义一个map,k为string,v为空结构体(目的：不占空间)
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.ShortUrlDb.DSN)
	m := make(map[string]struct{}, len(c.BlackList))
	for _, v := range c.BlackList {
		m[v] = struct{}{}
	}

	return &ServiceContext{
		Config:     c,
		ShortUrlDb: model.NewShortUrlMapModel(conn),
		Sequence:   sequence.NewSeqMysql(c.SequenceDb.DSN),
		BlackList:  m,
	}
}

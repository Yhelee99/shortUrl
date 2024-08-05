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

	Sequence sequence.Sequence
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.ShortUrlDb.DSN)

	return &ServiceContext{
		Config:     c,
		ShortUrlDb: model.NewShortUrlMapModel(conn),
		Sequence:   sequence.NewSeqMysql(c.SequenceDb.DSN),
	}
}

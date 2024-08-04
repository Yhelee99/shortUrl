package svc

import (
	"Project/internal/config"
	"Project/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	ShortUrlDb model.ShortUrlMapModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.ShortUrlDb.DSN)

	return &ServiceContext{
		Config:     c,
		ShortUrlDb: model.NewShortUrlMapModel(conn),
	}
}

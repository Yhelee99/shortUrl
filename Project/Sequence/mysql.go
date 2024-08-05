package sequence

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	sqlStr = "REPLACE INTO sequence ( `stub`, `timestamp`) VALUES ('a',NOW())"
)

type Sequence interface {
	GetNumb() (uint64, error)
}

type Mysql struct {
	conn sqlx.SqlConn
}

func NewSeqMysql(dsn string) Sequence {
	return &Mysql{sqlx.NewMysql(dsn)}
}

func (m Mysql) GetNumb() (seqid uint64, err error) {

	// StmtSession 用于准备语句（预编译）
	var s sqlx.StmtSession
	//
	s, err = m.conn.Prepare(sqlStr)
	if err != nil {
		logx.Errorw("conn.Prepare failed.", logx.Field("err", err))
		return 0, err
	}

	// 执行
	var res sql.Result
	res, err = s.Exec()
	if err != nil {
		logx.Errorw("conn.Prepare failed.", logx.Field("err", err))
		return 0, nil
	}

	// 获取最后一个id
	var lid int64
	lid, err = res.LastInsertId()
	if err != nil {
		logx.Errorw("res.LastInsertId failed.", logx.Field("err", err))
		return 0, nil
	}
	return uint64(lid), nil
}

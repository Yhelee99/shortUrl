// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sequenceFieldNames          = builder.RawFieldNames(&Sequence{})
	sequenceRows                = strings.Join(sequenceFieldNames, ",")
	sequenceRowsExpectAutoSet   = strings.Join(stringx.Remove(sequenceFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sequenceRowsWithPlaceHolder = strings.Join(stringx.Remove(sequenceFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	sequenceModel interface {
		Insert(ctx context.Context, data *Sequence) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Sequence, error)
		FindOneByStub(ctx context.Context, stub string) (*Sequence, error)
		Update(ctx context.Context, data *Sequence) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultSequenceModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Sequence struct {
		Id        uint64    `db:"id"`
		Stub      string    `db:"stub"`
		Timestamp time.Time `db:"timestamp"`
	}
)

func newSequenceModel(conn sqlx.SqlConn) *defaultSequenceModel {
	return &defaultSequenceModel{
		conn:  conn,
		table: "`sequence`",
	}
}

func (m *defaultSequenceModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSequenceModel) FindOne(ctx context.Context, id uint64) (*Sequence, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sequenceRows, m.table)
	var resp Sequence
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSequenceModel) FindOneByStub(ctx context.Context, stub string) (*Sequence, error) {
	var resp Sequence
	query := fmt.Sprintf("select %s from %s where `stub` = ? limit 1", sequenceRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, stub)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSequenceModel) Insert(ctx context.Context, data *Sequence) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sequenceRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Stub, data.Timestamp)
	return ret, err
}

func (m *defaultSequenceModel) Update(ctx context.Context, newData *Sequence) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sequenceRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Stub, newData.Timestamp, newData.Id)
	return err
}

func (m *defaultSequenceModel) tableName() string {
	return m.table
}

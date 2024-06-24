package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PayModel = (*customPayModel)(nil)

type (
	// PayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPayModel.
	PayModel interface {
		payModel
		FindOneByOid(ctx context.Context, oid int64) (*Pay, error)
	}

	customPayModel struct {
		*defaultPayModel
	}
)

// NewPayModel returns a model for the database table.
func NewPayModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PayModel {
	return &customPayModel{
		defaultPayModel: newPayModel(conn, c, opts...),
	}
}

func (m *defaultPayModel) FindOneByOid(ctx context.Context, oid int64) (*Pay, error) {
	var resp Pay
	query := fmt.Sprintf("select %s from %s where `oid` = ? limit 1", payRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, oid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
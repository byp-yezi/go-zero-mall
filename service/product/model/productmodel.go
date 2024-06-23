package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Product, error)
		SelectBuilder() squirrel.SelectBuilder
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c, opts...),
	}
}

func (m *defaultProductModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultProductModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Product, error) {
	builder = builder.Columns(productRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("stock != 0 and status != 0").Limit(uint64(pageSize)).Offset(uint64(offset)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Product
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

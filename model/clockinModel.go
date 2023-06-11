package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ClockinModel = (*customClockinModel)(nil)

type (
	// ClockinModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClockinModel.
	ClockinModel interface {
		clockinModel
	}

	customClockinModel struct {
		*defaultClockinModel
	}
)

// NewClockinModel returns a model for the database table.
func NewClockinModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ClockinModel {
	return &customClockinModel{
		defaultClockinModel: newClockinModel(conn, c, opts...),
	}
}

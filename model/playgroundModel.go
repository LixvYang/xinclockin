package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PlaygroundModel = (*customPlaygroundModel)(nil)

type (
	// PlaygroundModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPlaygroundModel.
	PlaygroundModel interface {
		playgroundModel
	}

	customPlaygroundModel struct {
		*defaultPlaygroundModel
	}
)

// NewPlaygroundModel returns a model for the database table.
func NewPlaygroundModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PlaygroundModel {
	return &customPlaygroundModel{
		defaultPlaygroundModel: newPlaygroundModel(conn, c, opts...),
	}
}

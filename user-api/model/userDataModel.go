package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserDataModel = (*customUserDataModel)(nil)

type (
	// UserDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDataModel.
	UserDataModel interface {
		userDataModel
	}

	customUserDataModel struct {
		*defaultUserDataModel
	}
)

// NewUserDataModel returns a model for the database table.
func NewUserDataModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserDataModel {
	return &customUserDataModel{
		defaultUserDataModel: newUserDataModel(conn, c, opts...),
	}
}

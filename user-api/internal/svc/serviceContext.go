package svc

import (
	"go-zero-demo/user-api/internal/config"
	"go-zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	UserDataModel model.UserDataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserDataModel: model.NewUserDataModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}

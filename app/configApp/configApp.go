package configApp

import (
	"context"
	"new_it/app/configApp/model"

	"github.com/ServiceWeaver/weaver"
)

type ConfigApp interface {
	GetDBType(ctx context.Context) (string, error)

	GetConfigMysql(ctx context.Context) (model.Mysql, error)
	GetConfigJWT(ctx context.Context) (model.JWT, error)
}

type config_app struct {
	weaver.Implements[ConfigApp]
	weaver.WithConfig[ConfigInfo]
}

func (c *config_app) GetDBType(ctx context.Context) (string, error) {
	dbtype := c.Config().DbType
	if dbtype == "" {
		dbtype = "mysql"
	}

	return dbtype, nil
}

func (c *config_app) GetConfigJWT(ctx context.Context) (model.JWT, error) {

	jwt := c.Config().JWT

	return jwt, nil
}

func (c *config_app) GetConfigMysql(ctx context.Context) (model.Mysql, error) {

	mysql := c.Config().Mysql

	return mysql, nil
}

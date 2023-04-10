package configApp

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

type ConfigApp interface {
	Get(ctx context.Context, key string) (string, error)
	GetDBType(ctx context.Context) (string, error)
	//CopyCfg(ctx context.Context) (ConfigInfo, error)

	//GetConfigMysql(ctx context.Context) (model.Mysql, error)
	//GetConfigJWT(ctx context.Context) (model.JWT, error)
}

type config_app struct {
	weaver.Implements[ConfigApp]
	weaver.WithConfig[ConfigInfo]
}

// func (c *config_app) CopyCfg(ctx context.Context) (ConfigInfo, error) {
// 	var cfg ConfigInfo

// 	cfg.DbType = c.Config().DbType
// 	cfg.Mysql.Ip = c.Config().Mysql.Ip

// 	return cfg, nil

// }

func (c *config_app) GetDBType(ctx context.Context) (string, error) {
	dbtype := c.Config().DbType
	if dbtype == "" {
		dbtype = "mysql"
	}

	return dbtype, nil
}

// func (c *config_app) GetConfigJWT(ctx context.Context) (model.JWT, error) {

// 	jwt := c.Config().JWT
// 	fmt.Println(jwt)

// 	return jwt, nil
// }

func (c *config_app) Get(_ context.Context, key string) (string, error) {

	ip := c.Config().Mysql.Ip
	port := c.Config().Mysql.Port
	if ip == "" {
		ip = "0.0.0.0"
	}
	//weaver.toml
	fmt.Println("ip=" + ip + " port=" + port + " " + c.Config().Mysql.Config)
	fmt.Println(c.Config().JWT.Issuer + " " + c.Config().Mysql.Ip)

	log := c.Logger()
	log.Debug(" ip=" + ip)

	return ip, nil
}

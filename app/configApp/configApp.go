package configApp

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

type ConfigApp interface {
	Get(ctx context.Context, key string) (string, error)
	//GetConfigMysql(ctx context.Context) (model.Mysql, error)
	//GetConfigJWT(ctx context.Context) (model.JWT, error)
}

type config_app struct {
	weaver.Implements[ConfigApp]
	weaver.WithConfig[ConfigInfo]
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

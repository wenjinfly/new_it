package configApp

import (
	"new_it/app/configApp/model"

	"github.com/ServiceWeaver/weaver"
)

type ConfigInfo struct {
	weaver.AutoMarshal
	DbType string //数据库类型
	JWT    model.JWT
	Mysql  model.Mysql
}

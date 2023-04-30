package model

import "github.com/ServiceWeaver/weaver"

type Mysql struct {
	weaver.AutoMarshal
	Ip           string // 服务器ip
	Port         string // 端口
	Config       string // 高级配置
	Dbname       string // 数据库名
	Username     string // 数据库用户名
	Password     string // 数据库密码
	MaxIdleConns int    // 空闲中的最大连接数
	MaxOpenConns int    // 打开到数据库的最大连接数
	LogMode      string // 是否开启Gorm全局日志
}

//"root:@tcp(localhost:3306)/"&parseTime=true
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Ip + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

package initialize

import (
	"new_it/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm_mysql() *gorm.DB {

	m := global.GLB_CFG_INFO.Mysql
	if m.Dbname == "" {
		panic("can not find dbname")
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		panic("can not open db error: " + err.Error())
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}

}

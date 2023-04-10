package global

import (
	"new_it/app/configApp"
	"sync"

	"github.com/ServiceWeaver/weaver"
	"gorm.io/gorm"
)

var (
	GLB_WEAVER_ROOT weaver.Instance
	GLB_DB          *gorm.DB
	GLB_DBList      map[string]*gorm.DB
	GLB_CFG_APP     configApp.ConfigInfo

	lock sync.RWMutex
)

// 通过名称获取DB
func GetGlobalDBByName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()

	db, ok := GLB_DBList[dbname]
	if !ok || db == nil {
		panic("can not find" + dbname + " dbname")
	}

	return db
}

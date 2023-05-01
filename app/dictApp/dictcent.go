package dictApp

import (
	"net/http"
	"new_it/app/dictApp/dbdata"
	"new_it/app/dictApp/model"
)

/*
Tables : 当前模块数据库表DictInfo
*/
var Tables []interface{}

func initTable() {
	Tables = append(Tables, new(model.DictType))
	Tables = append(Tables, new(model.DictInfo))

}

/*
RouterHander:路由信息
*/
type RouterHander func(w http.ResponseWriter, r *http.Request)

type RouterInfo struct {
	router  string
	handler RouterHander
}

var RouterList []RouterInfo

func initRouter() {
	//post
	//RouterList = append(RouterList, RouterInfo{"/login", api.UserApi.Login})
	//RouterList = append(RouterList, RouterInfo{"/user/register", api.UserApi.Register})

	///menu/

}

var Dbdatas []dbdata.InitDBData

// 初始化数据库表的内容
func initdbdatas() {
	Dbdatas = append(Dbdatas, dbdata.DictTypes)
	Dbdatas = append(Dbdatas, dbdata.DictInfos)
}

/*
模块初始化
*/
func init() {
	initTable()
	initdbdatas()
	initRouter()
}

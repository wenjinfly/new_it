package dictApp

import (
	"net/http"
	"new_it/app/dictApp/api"
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
	//DictType
	RouterList = append(RouterList, RouterInfo{"/dict/AddDictType", api.DirctTypeApi.AddDictType})
	RouterList = append(RouterList, RouterInfo{"/dict/DeleteDictType", api.DirctTypeApi.DeleteDictType})
	RouterList = append(RouterList, RouterInfo{"/dict/UpdateDictType", api.DirctTypeApi.UpdateDictType})
	RouterList = append(RouterList, RouterInfo{"/dict/GetDictTypeByCode", api.DirctTypeApi.GetDictTypeByCode})
	RouterList = append(RouterList, RouterInfo{"/dict/GetDictTypeList", api.DirctTypeApi.GetDictTypeList})

	//DictInfo
	RouterList = append(RouterList, RouterInfo{"/dict/AddDictInfo", api.DirctInfoApi.AddDictInfo})
	RouterList = append(RouterList, RouterInfo{"/dict/DeleteDictInfo", api.DirctInfoApi.DeleteDictInfo})
	RouterList = append(RouterList, RouterInfo{"/dict/UpdateDictInfo", api.DirctInfoApi.UpdateDictInfo})
	RouterList = append(RouterList, RouterInfo{"/dict/GetDictInfoByCode", api.DirctInfoApi.GetDictInfoByCode})
	RouterList = append(RouterList, RouterInfo{"/dict/GetDictInfoListByTypeCode", api.DirctInfoApi.GetDictInfoListByTypeCode})

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

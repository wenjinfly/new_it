package usercentApp

import (
	"net/http"
	"new_it/app/usercentApp/api"
	"new_it/app/usercentApp/model"
)

/*
Tables : 当前模块数据库表
*/
var Tables []interface{}

func initTable() {
	Tables = append(Tables, new(model.SysUsers))
	Tables = append(Tables, new(model.SysAuthorities))
	Tables = append(Tables, new(model.SysBaseMenus))
	Tables = append(Tables, new(model.SysUserAuthority))
	Tables = append(Tables, new(model.SysAuthorityMenus))
	Tables = append(Tables, new(model.SysOperationRecords))
	Tables = append(Tables, new(model.SysApis))
	Tables = append(Tables, new(model.JwtBlacklists))

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
	RouterList = append(RouterList, RouterInfo{"/login", api.UserApi.Login})
	RouterList = append(RouterList, RouterInfo{"/user/register", api.UserApi.Register})
	RouterList = append(RouterList, RouterInfo{"/user/changePassword", api.UserApi.ChangePassword})
	RouterList = append(RouterList, RouterInfo{"/user/resetPassword", api.UserApi.ResetPassword})
	///user/resetPassword
}

/*
模块初始化
*/
func init() {
	initTable()
	initRouter()
}

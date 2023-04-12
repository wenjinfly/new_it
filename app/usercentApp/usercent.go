package usercentApp

import "new_it/app/usercentApp/model"

/*
Tables : 当前模块数据库表
*/
var Tables []interface{}

func init() {
	Tables = append(Tables, new(model.SysUsers))
	Tables = append(Tables, new(model.SysAuthorities))
	Tables = append(Tables, new(model.SysBaseMenus))
	Tables = append(Tables, new(model.SysUserAuthority))
	Tables = append(Tables, new(model.SysAuthorityMenus))
	Tables = append(Tables, new(model.SysOperationRecords))
	Tables = append(Tables, new(model.SysApis))
	Tables = append(Tables, new(model.JwtBlacklists))
}

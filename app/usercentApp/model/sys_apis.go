// Package models  api
package model

import "time"

// SysApis  api。
// 说明:
// 表名:sys_apis
// group: SysApis
type SysApis struct {
	Id          uint64     `gorm:"column:id;primaryKey" json:"Id"`        //type:BIGINT UNSIGNED   comment:               version:2023-03-12 23:20
	Path        string     `gorm:"column:path" json:"Path"`               //type:string            comment:api路径        version:2023-03-12 23:20
	Description string     `gorm:"column:description" json:"Description"` //type:string            comment:api中文描述    version:2023-03-12 23:20
	ApiGroup    string     `gorm:"column:api_group" json:"ApiGroup"`      //type:string            comment:api组          version:2023-03-12 23:20
	Method      string     `gorm:"column:method" json:"Method"`           //type:string            comment:方法           version:2023-03-12 23:20
	CreatedAt   *time.Time `gorm:"column:created_at" json:"CreatedAt"`    //type:*time.Time        comment:               version:2023-03-12 23:20
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"UpdatedAt"`    //type:*time.Time        comment:               version:2023-03-12 23:20
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"DeletedAt"`    //type:*time.Time        comment:               version:2023-03-12 23:20

}

// TableName 表名:sys_apis，api。
// 说明:
func (s *SysApis) TableName() string {
	return "sys_apis"
}

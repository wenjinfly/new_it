// Package models  操作记录
package model

import "time"

// SysOperationRecords  操作记录。
// 说明:
// 表名:sys_operation_records
type SysOperationRecords struct {
	Id           uint64    `gorm:"column:id;primaryKey" json:"Id"`           //type:BIGINT UNSIGNED   comment:            version:2023-03-12 23:17
	CreatedAt    time.Time `gorm:"column:created_at" json:"CreatedAt"`       //type:*time.Time        comment:            version:2023-03-12 23:17
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"UpdatedAt"`       //type:*time.Time        comment:            version:2023-03-12 23:17
	DeletedAt    time.Time `gorm:"column:deleted_at" json:"DeletedAt"`       //type:*time.Time        comment:            version:2023-03-12 23:17
	Ip           string    `gorm:"column:ip" json:"Ip"`                      //type:string            comment:请求ip      version:2023-03-12 23:17
	Method       string    `gorm:"column:method" json:"Method"`              //type:string            comment:请求方法    version:2023-03-12 23:17
	Path         string    `gorm:"column:path" json:"Path"`                  //type:string            comment:请求路径    version:2023-03-12 23:17
	Status       bool      `gorm:"column:status" json:"Status"`              //type:BIGINT            comment:请求状态    version:2023-03-12 23:17
	Latency      bool      `gorm:"column:latency" json:"Latency"`            //type:BIGINT            comment:延迟        version:2023-03-12 23:17
	Agent        string    `gorm:"column:agent" json:"Agent"`                //type:string            comment:代理        version:2023-03-12 23:17
	ErrorMessage string    `gorm:"column:error_message" json:"ErrorMessage"` //type:string            comment:错误信息    version:2023-03-12 23:17
	Body         string    `gorm:"column:body" json:"Body"`                  //type:string            comment:请求Body    version:2023-03-12 23:17
	Resp         string    `gorm:"column:resp" json:"Resp"`                  //type:string            comment:响应Body    version:2023-03-12 23:17
	UserId       uint64    `gorm:"column:user_id" json:"UserId"`             //type:BIGINT UNSIGNED   comment:用户id      version:2023-03-12 23:17
}

// TableName 表名:sys_operation_records，操作记录。
// 说明:
func (s *SysOperationRecords) TableName() string {
	return "sys_operation_records"
}

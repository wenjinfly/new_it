package model

import (
	"new_it/common"
	"time"
)

// 收藏的任务
type FavoriteTask struct {
	Id           uint64            `gorm:"column:id;primaryKey" json:"Id"`                     //type:  comment:关系id          version:2023-04-02 16:08
	ApplyUserId  uint64            `gorm:"column:apply_user_id;primaryKey" json:"ApplyUserId"` //type:  comment:申请者用户ID    version:2023-04-02 16:08
	TaskId       uint64            `gorm:"column:task_id;primaryKey" json:"TaskId"`            //type:  comment:任务ID          version:2023-04-02 16:08
	TaskUserId   uint64            `gorm:"column:task_user_id" json:"TaskUserId"`              //type:  comment:任务ID          version:2023-04-02 16:08
	TaskName     string            `gorm:"column:task_name" json:"TaskName"`
	CommStatus   string            `gorm:"column:comm_status" json:"CommStatus"` //沟通状态
	RelationTime *common.LocalTime `gorm:"column:relation_time" json:"RelationTime"`
	CreatedAt    *time.Time        `gorm:"column:created_at" json:"CreatedAt"` //type:*time.Time   comment:创建时间              version:2023-04-02 00:05
	UpdatedAt    *time.Time        `gorm:"column:updated_at" json:"UpdatedAt"` //type:*time.Time   comment:更新时间              version:2023-04-02 00:05
	DeletedAt    *time.Time        `gorm:"column:deleted_at" json:"DeletedAt"` //type:*time.Time        comment:                version:2023-03-12 00:10
}

// TableName 表名:favorite_task，收藏的任务
// 说明:
func (u *FavoriteTask) TableName() string {
	return "favorite_task"
}

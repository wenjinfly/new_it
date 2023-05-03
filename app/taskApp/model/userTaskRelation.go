package model

import (
	"new_it/common"
	"time"
)

// UserTaskRelation  申请用户任务关系。
// 说明:
// 表名:user_task_relation
// group: UserTaskRelation
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.UserTaskRelation
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.UserTaskRelation
// version:2023-04-02 16:08
type UserTaskRelation struct {
	Id           uint64            `gorm:"column:id;primaryKey" json:"Id"`                     //type:  comment:关系id          version:2023-04-02 16:08
	ApplyUserId  uint64            `gorm:"column:apply_user_id;primaryKey" json:"ApplyUserId"` //type:  comment:申请者用户ID    version:2023-04-02 16:08
	TaskId       uint64            `gorm:"column:task_id;primaryKey" json:"TaskId"`            //type:  comment:任务ID          version:2023-04-02 16:08
	TaskUserId   uint64            `gorm:"column:task_user_id" json:"TaskUserId"`              //type:  comment:任务ID          version:2023-04-02 16:08
	CommStatus   string            `gorm:"column:comm_status" json:"CommStatus"`               //沟通状态
	RelationTime *common.LocalTime `gorm:"column:relation_time" json:"RelationTime"`
	CreatedAt    *time.Time        `gorm:"column:created_at" json:"CreatedAt"` //type:*time.Time   comment:创建时间              version:2023-04-02 00:05
	UpdatedAt    *time.Time        `gorm:"column:updated_at" json:"UpdatedAt"` //type:*time.Time   comment:更新时间              version:2023-04-02 00:05
	DeletedAt    *time.Time        `gorm:"column:deleted_at" json:"DeletedAt"` //type:*time.Time        comment:                version:2023-03-12 00:10
}

// TableName 表名:user_task_relation，申请用户任务关系(已签约)。
// 说明:
func (u *UserTaskRelation) TableName() string {
	return "user_task_relation"
}

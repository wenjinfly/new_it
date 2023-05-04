package model

import (
	"new_it/common"
	"time"
)

// Task  任务单。
// 说明:
// 表名:Task
// group: Task
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.Task
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.Task
// version:2023-04-02 00:05
type Task struct {
	TaskId               uint64            `gorm:"column:task_id;primaryKey" json:"TaskId"`                             //type:             comment:任务ID                version:2023-04-02 00:05
	UserId               uint64            `gorm:"column:user_id" json:"UserId"`                                        //type:             comment:任务拥有者的用户ID    version:2023-04-02 00:05
	TaskName             string            `gorm:"column:task_name" json:"TaskName"`                                    //type:string       comment:任务名称              version:2023-04-02 00:05
	TaskRewardMin        float64           `gorm:"column:task_reward_min" json:"TaskRewardMin"`                         //type:*float64     comment:最低酬金              version:2023-04-02 00:05
	TaskRewardMax        float64           `gorm:"column:task_reward_max" json:"TaskRewardMax"`                         //type:*float64     comment:最高酬金              version:2023-04-02 00:05
	TaskStatus           string            `gorm:"column:task_status;size:10" json:"TaskStatus"`                        //type:*int         comment:任务状态              version:2023-04-02 00:05
	TaskRequiredPerson   int               `gorm:"column:task_required_person" json:"TaskRequiredPerson"`               //type:*int         comment:任务需要人数          version:2023-04-02 00:05
	TaskDescribe         string            `gorm:"column:task_describe;type:text" json:"TaskDescribe"`                  //type:string       comment:任务描述              version:2023-04-02 00:05
	TaskPhase            string            `gorm:"column:task_phase;size:10" json:"TaskPhase"`                          //type:*int         comment:任务阶段              version:2023-04-02 00:05
	TaskDuty             string            `gorm:"column:task_duty;type:text" json:"TaskDuty"`                          //type:string       comment:任务职责              version:2023-04-02 00:05
	TaskGetQualification string            `gorm:"column:task_get_qualification;type:text" json:"TaskGetQualification"` //type:string       comment:任务任职资格          version:2023-04-02 00:05
	TaskAddress          string            `gorm:"column:task_address" json:"TaskAddress"`                              //type:string       comment:任务地址              version:2023-04-02 00:05
	TaskBegin            *common.LocalTime `gorm:"column:task_begin" json:"TaskBegin"`                                  //type:*time.Time   comment:任务开始时间          version:2023-04-02 00:05
	TaskEnd              *common.LocalTime `gorm:"column:task_end" json:"TaskEnd"`                                      //type:*time.Time   comment:任务结束时间          version:2023-04-02 00:05
	CreatedAt            *time.Time        `gorm:"column:created_at" json:"CreatedAt"`                                  //type:*time.Time   comment:创建时间              version:2023-04-02 00:05
	UpdatedAt            *time.Time        `gorm:"column:updated_at" json:"UpdatedAt"`                                  //type:*time.Time   comment:更新时间              version:2023-04-02 00:05
	DeletedAt            *time.Time        `gorm:"column:deleted_at" json:"DeletedAt"`                                  //type:*time.Time        comment:                version:2023-03-12 00:10
	Revision             int               `gorm:"column:revision" json:"Revision"`                                     //type:*int         comment:乐观锁                version:2023-04-02 00:05

}

// TableName 表名:Task，任务单。
// 说明:
func (t *Task) TableName() string {
	return "task"
}

package dbdata

import (
	"errors"
	"new_it/app/taskApp/model"
	"new_it/common"
	"new_it/global"
	"time"

	"gorm.io/gorm"
)

var TaskInfos = new(TASK_INFO)

type TASK_INFO struct{}

func (u *TASK_INFO) TableName() string {
	return "task"
}

func (u *TASK_INFO) Initialize() error {

	t1, _ := time.Parse("2006-01-02 15:04:05", "2023-05-01 10:00:05")
	t2, _ := time.Parse("2006-01-02 15:04:05", "2023-05-10 10:00:05")

	tt1 := common.LocalTime(t1)
	tt2 := common.LocalTime(t2)

	entities := []model.Task{

		{UserId: 3, TaskName: "小程序开发", TaskRewardMin: 5000.21, TaskRewardMax: 6000.66, TaskStatus: "202", TaskRequiredPerson: 2,
			TaskDescribe: "我们需要做一个小程序", TaskPhase: "112", TaskDuty: "可以保证按时开发完成商城小程序",
			TaskGetQualification: "有开发小程序经验", TaskAddress: "远程",
			TaskBegin: &tt1, TaskEnd: &tt2},
		{UserId: 3, TaskName: "第二个小程序开发", TaskRewardMin: 7000, TaskRewardMax: 8800, TaskStatus: "202", TaskRequiredPerson: 3,
			TaskDescribe: "我们需要做第二个小程序", TaskPhase: "112", TaskDuty: "可以保证按时开发完成商城小程序2",
			TaskGetQualification: "有开发小程序经验2", TaskAddress: "远程",
			TaskBegin: &tt1, TaskEnd: &tt2},
		{UserId: 3, TaskName: "第三个小程序开发", TaskRewardMin: 9000, TaskRewardMax: 10000, TaskStatus: "202", TaskRequiredPerson: 4,
			TaskDescribe: "我们需要做第三个小程序", TaskPhase: "112", TaskDuty: "可以保证按时开发完成商城小程序3",
			TaskGetQualification: "有开发小程序经验3", TaskAddress: "远程",
			TaskBegin: &tt1, TaskEnd: &tt2},
	}
	if err := global.GLB_DB.Create(&entities).Error; err != nil {
		return errors.New(u.TableName() + "表数据初始化失败!")
	}
	return nil
}

func (u *TASK_INFO) CheckDataExist() bool {
	if errors.Is(global.GLB_DB.Where("task_id = ?", "1").First(&model.Task{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

package dbdata

import (
	"errors"
	"new_it/app/taskApp/model"
	"new_it/common"
	"new_it/global"
	"time"

	"gorm.io/gorm"
)

var ChatInfo = new(Chat_INFO)

type Chat_INFO struct{}

func (u *Chat_INFO) TableName() string {
	return "task"
}

func (u *Chat_INFO) Initialize() error {

	t1, _ := time.Parse("2006-01-02 15:04:05", "2023-05-01 10:00:05")
	t2, _ := time.Parse("2006-01-02 15:04:05", "2023-05-01 10:00:06")

	tt1 := common.LocalTime(t1)
	tt2 := common.LocalTime(t2)

	entities := []model.ChatMessage{
		{ChatUuid: "0-0-1-0", FromId: 3, FromName: "测试人员", SendTime: &tt1, Content: "小程序项目经验丰富，已经按时保量完成100个了", ChatType: 1, ClassType: 2, TaskID: 1, TaskName: "小程序开发", ToId: 0, GroupId: 0}, //task下面的留言
		{ChatUuid: "2-3-0-0", FromId: 3, FromName: "测试人员", ToId: 2, ToName: "超级管理员", SendTime: &tt1, Content: "您好，当前小程序任务还可以接单么？", ChatType: 1, ClassType: 1, TaskID: 0, GroupId: 0},         //用户交流
		{ChatUuid: "2-3-0-0", FromId: 2, FromName: "超级管理员", ToId: 3, ToName: "测试人员", SendTime: &tt2, Content: "现在还可以接单，你之前做过类似的项目么？", ChatType: 1, ClassType: 1, TaskID: 0, GroupId: 0},      //用户交流

	}
	if err := global.GLB_DB.Create(&entities).Error; err != nil {
		return errors.New(u.TableName() + "表数据初始化失败!")
	}
	return nil
}

func (u *Chat_INFO) CheckDataExist() bool {
	if errors.Is(global.GLB_DB.Where("id = ?", "1").First(&model.ChatMessage{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

package dbdata

import (
	"errors"
	"new_it/app/dictApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

var DictInfos = new(DICT_INFO)

type DICT_INFO struct{}

func (u *DICT_INFO) TableName() string {
	return "dict_info"
}

func (u *DICT_INFO) Initialize() error {

	entities := []model.DictInfo{

		{Code: "101", Name: "SoftwarePlanning", CNName: "软件策划", Fixed: true, TypeCode: "100"},
		{Code: "102", Name: "SoftwareRequirements", CNName: "软件需求", Fixed: true, TypeCode: "100"},
		{Code: "103", Name: "OutlineDesign", CNName: "概要设计", Fixed: true, TypeCode: "100"},
		{Code: "104", Name: "DetailedDesign", CNName: "详细设计", Fixed: true, TypeCode: "100"},
		{Code: "105", Name: "UnitTesting", CNName: "单元测计划", Fixed: true, TypeCode: "100"},
		{Code: "106", Name: "CodingAndTesting", CNName: "编码和单测", Fixed: true, TypeCode: "100"},
		{Code: "107", Name: "SystemTestPlan", CNName: "系统测试计划", Fixed: true, TypeCode: "100"},
		{Code: "108", Name: "SystemTesting", CNName: "系统测试", Fixed: true, TypeCode: "100"},
		{Code: "109", Name: "UserTesting", CNName: "用户测试", Fixed: true, TypeCode: "100"},
		{Code: "110", Name: "SoftwarePublishing", CNName: "软件发行", Fixed: true, TypeCode: "100"},
		{Code: "111", Name: "DeploymentOperations", CNName: "部署运维", Fixed: true, TypeCode: "100"},
		{Code: "112", Name: "WholeDevelopment", CNName: "整体开发", Fixed: true, TypeCode: "100"},
		//
		{Code: "201", Name: "TaskDraft", CNName: "任务草稿", Fixed: true, TypeCode: "200"},
		{Code: "202", Name: "TaskPublishing", CNName: "任务发布", Fixed: true, TypeCode: "200"},
		{Code: "203", Name: "TaskNegotiationInProgress", CNName: "任务协商中", Fixed: true, TypeCode: "200"},
		{Code: "204", Name: "TaskInProgress", CNName: "任务进行中", Fixed: true, TypeCode: "200"},
		{Code: "205", Name: "TaskAcceptanceInProgress", CNName: "任务验收中", Fixed: true, TypeCode: "200"},
		{Code: "206", Name: "TaskCompleted", CNName: "任务已完成", Fixed: true, TypeCode: "200"},
		{Code: "207", Name: "TaskAbandoned", CNName: "任务已废弃", Fixed: true, TypeCode: "200"},
		//
		{Code: "301", Name: "Applying", CNName: "申请中", Fixed: true, TypeCode: "300"},
		{Code: "302", Name: "InNegotiation", CNName: "洽谈中", Fixed: true, TypeCode: "300"},
		{Code: "303", Name: "ConditionsMet", CNName: "条件符合", Fixed: true, TypeCode: "300"},
		{Code: "304", Name: "LackOfConditions", CNName: "条件欠缺", Fixed: true, TypeCode: "300"},
		//
		{Code: "401", Name: "Male", CNName: "男", Fixed: true, TypeCode: "400"},
		{Code: "402", Name: "Female", CNName: "女", Fixed: true, TypeCode: "400"},
		{Code: "403", Name: "NotFilled", CNName: "未填", Fixed: true, TypeCode: "400"},
	}
	if err := global.GLB_DB.Create(&entities).Error; err != nil {
		return errors.New(u.TableName() + "表数据初始化失败!")
	}
	return nil
}

func (u *DICT_INFO) CheckDataExist() bool {
	if errors.Is(global.GLB_DB.Where("code = ?", "101").First(&model.DictInfo{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

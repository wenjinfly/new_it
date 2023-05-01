package dbdata

import (
	"errors"
	"new_it/app/dictApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

var DictTypes = new(DICT_TYPE)

type DICT_TYPE struct{}

func (u *DICT_TYPE) TableName() string {
	return "dict_type"
}

func (u *DICT_TYPE) Initialize() error {

	entities := []model.DictType{
		{TypeCode: "100", TypeName: "SoftLifeCycle", TypeCNName: "软件生命周期"},
		{TypeCode: "200", TypeName: "TaskStatus", TypeCNName: "任务状态"},
		{TypeCode: "300", TypeName: "CommStatus", TypeCNName: "沟通状态"},
	}
	if err := global.GLB_DB.Create(&entities).Error; err != nil {
		return errors.New(u.TableName() + "表数据初始化失败!")
	}
	return nil
}

func (u *DICT_TYPE) CheckDataExist() bool {
	if errors.Is(global.GLB_DB.Where("type_code = ?", "100").First(&model.DictType{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

package service

import (
	"errors"
	"new_it/app/taskApp/api/request"
	"new_it/app/taskApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

type USER_TASK_REL_SERVICE struct{}

func (us *USER_TASK_REL_SERVICE) AddRelation(task model.UserTaskRelation) error {

	err := global.GLB_DB.Where("task_id = ? AND apply_user_id = ? ", task.TaskId, task.ApplyUserId).First(&task).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("当前task_id和apply_user_id已有关联")
	}

	err = global.GLB_DB.Create(&task).Error
	return err
}

func (us *USER_TASK_REL_SERVICE) DeleteRelation(param request.ParamUserTaskRelation) error {
	var task model.Task
	err := global.GLB_DB.Where("id = ? AND apply_user_id = ? ", param.Id, param.UserId).First(&task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("当前id和user_id关联的task没有找到")
	}

	return global.GLB_DB.Delete(&task).Error
}

func (us *USER_TASK_REL_SERVICE) UpdateRelationStatus(param request.ParamUserTaskStatus) error {
	var task model.UserTaskRelation

	err := global.GLB_DB.Where("task_id = ? AND apply_user_id = ? ", param.TaskId, param.UserId).First(&task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("当前task_id和user_id的task没有找到")
	}

	if task.CommStatus == param.CommStatus {
		return errors.New("相同用户相同task的相同沟通状态不需要更新")
	}

	task.CommStatus = param.CommStatus

	err = global.GLB_DB.Updates(&task).Error
	return err
}

func (us *USER_TASK_REL_SERVICE) GetRelationListByUserId(info request.ParamTaskInfoList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GLB_DB.Model(&model.UserTaskRelation{})

	var taskList []model.UserTaskRelation
	err = db.Where("apply_user_id = ?", info.UserId).Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Where("apply_user_id = ?", info.UserId).Find(&taskList).Error

	return taskList, total, err
}

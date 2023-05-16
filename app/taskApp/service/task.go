package service

import (
	"errors"
	"new_it/app/taskApp/api/request"
	"new_it/app/taskApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

type TASK_SERVICE struct{}

func (us *TASK_SERVICE) AddTask(task model.Task) error {
	var err = global.GLB_DB.Create(&task).Error
	return err
}

func (us *TASK_SERVICE) DeleteTask(param request.ParamTaskIDStatusPhash) error {

	var task model.Task
	err := global.GLB_DB.Where("task_id = ? AND user_id = ? ", param.TaskId, param.UserId).First(&task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("当前task_id和user_id的task没有找到")
	}

	return global.GLB_DB.Delete(&task).Error
}

func (us *TASK_SERVICE) UpdateTaskInfo(task model.Task) error {
	err := global.GLB_DB.Where("task_id = ? AND user_id = ? ", task.TaskId, task.UserId).First(&model.Task{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("当前task_id和user_id的task没有找到")
	}

	err = global.GLB_DB.Updates(&task).Error
	return err
}

func (us *TASK_SERVICE) UpdateTaskStatus(param request.ParamTaskIDStatusPhash) error {
	err := global.GLB_DB.Where("task_id = ? AND user_id = ? ", param.TaskId, param.UserId).First(&model.Task{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("当前task_id和user_id的task没有找到")
	}

	var task model.Task
	task.TaskId = param.TaskId
	task.TaskStatus = param.TaskStatus

	err = global.GLB_DB.Updates(&task).Error
	return err
}

func (us *TASK_SERVICE) UpdateTaskPhase(param request.ParamTaskIDStatusPhash) error {
	err := global.GLB_DB.Where("task_id = ? AND user_id = ? ", param.TaskId, param.UserId).First(&model.Task{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("当前task_id和user_id的task没有找到")
	}

	var task model.Task
	task.TaskId = param.TaskId
	task.TaskPhase = param.TaskPhase

	err = global.GLB_DB.Updates(&task).Error
	return err
}

func (us *TASK_SERVICE) GetTaskByTaskID(TaskId uint64) (task model.Task, err error) {

	err = global.GLB_DB.Where("task_id = ?", TaskId).First(&task).Error
	return
}

func (us *TASK_SERVICE) GetTaskListByUserId(info request.ParamTaskInfoList) (list interface{}, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GLB_DB.Model(&model.Task{})

	var taskList []model.Task
	err = db.Where("user_id = ?", info.UserId).Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&taskList).Error

	return taskList, total, err
}

func (us *TASK_SERVICE) GetTaskListByName(info request.ParamTaskInfoList) (list interface{}, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GLB_DB.Model(&model.Task{})

	var taskList []model.Task
	err = db.Where("task_name Like ?", "%"+info.Name+"%").Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&taskList).Error

	return taskList, total, err
}

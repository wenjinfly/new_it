package service

import (
	"new_it/app/taskApp/api/request"
	"new_it/app/taskApp/model"
	"new_it/global"
)

type TASK_SERVICE struct{}

func (us *TASK_SERVICE) AddTask(task model.Task) error {
	var err = global.GLB_DB.Create(&task).Error
	return err
}

func (us *TASK_SERVICE) DeleteTask() {
}

func (us *TASK_SERVICE) UpdateTaskInfo() {
}

func (us *TASK_SERVICE) UpdateTaskStatus() {

}

func (us *TASK_SERVICE) UpdateTaskPhase() {

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

	err = db.Limit(limit).Offset(offset).Where("user_id = ?", info.UserId).Find(&taskList).Error

	return taskList, total, err
}

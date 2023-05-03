package service

import (
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

func (us *TASK_SERVICE) GetTaskListByUserId() {
}

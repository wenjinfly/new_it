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

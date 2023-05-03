package response

import "new_it/app/taskApp/model"

type ParamTaskResponse struct {
	TaskInfo model.Task `json:"Task"`
}

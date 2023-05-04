package taskApp

import (
	"net/http"
	"new_it/app/taskApp/api"
	"new_it/app/taskApp/dbdata"
	"new_it/app/taskApp/model"
)

/*
Tables : 当前模块数据库表
*/
var Tables []interface{}

func initTable() {
	Tables = append(Tables, new(model.Task))
	Tables = append(Tables, new(model.UserTaskRelation))
	Tables = append(Tables, new(model.ChatMessage))
	Tables = append(Tables, new(model.Contract))

}

/*
RouterHander:路由信息
*/
type RouterHander func(w http.ResponseWriter, r *http.Request)

type RouterInfo struct {
	router  string
	handler RouterHander
}

var RouterList []RouterInfo

func initRouter() {
	//task
	RouterList = append(RouterList, RouterInfo{"/task/AddTask", api.TaskApi.AddTask})

	RouterList = append(RouterList, RouterInfo{"/task/DeleteTask", api.TaskApi.DeleteTask})
	RouterList = append(RouterList, RouterInfo{"/task/UpdateTaskInfo", api.TaskApi.UpdateTaskInfo})
	RouterList = append(RouterList, RouterInfo{"/task/UpdateTaskStatus", api.TaskApi.UpdateTaskStatus})
	RouterList = append(RouterList, RouterInfo{"/task/UpdateTaskPhase", api.TaskApi.UpdateTaskPhase})

	RouterList = append(RouterList, RouterInfo{"/task/GetTaskByTaskID", api.TaskApi.GetTaskByTaskID})
	RouterList = append(RouterList, RouterInfo{"/task/GetTaskListByUserId", api.TaskApi.GetTaskListByUserId})
	//todo
	RouterList = append(RouterList, RouterInfo{"/relation/AddRelation", api.UserTaskRelationApi.AddRelation})
	RouterList = append(RouterList, RouterInfo{"/relation/DeleteRelation", api.UserTaskRelationApi.DeleteRelation}) //todo
	RouterList = append(RouterList, RouterInfo{"/relation/UpdateRelationStatus", api.UserTaskRelationApi.UpdateRelationStatus})
	RouterList = append(RouterList, RouterInfo{"/relation/GetRelationListByUserId", api.UserTaskRelationApi.GetRelationListByUserId})

	//chat
	RouterList = append(RouterList, RouterInfo{"/chat/AddChat", api.ChatMessageApi.AddChat})
	RouterList = append(RouterList, RouterInfo{"/chat/GetChatListByIds", api.ChatMessageApi.GetChatListByIds})

}

var Dbdatas []dbdata.InitDBData

// 初始化数据库表的内容
func initdbdatas() {
	Dbdatas = append(Dbdatas, dbdata.TaskInfos)
	Dbdatas = append(Dbdatas, dbdata.ChatInfo)
}

/*
模块初始化
*/
func init() {
	initTable()
	initdbdatas()
	initRouter()
}

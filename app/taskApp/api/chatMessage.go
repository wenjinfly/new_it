package api

import (
	"net/http"
	"new_it/app/taskApp/api/request"
	"new_it/app/taskApp/model"
	"new_it/app/taskApp/service"
	"new_it/common"
	"new_it/errorcode"
	"new_it/utils"
	"strconv"
)

type ChatMessageAPI struct{}

// AddChat
//
//	@receiver us
//	@param w
//	@param r
func (us *ChatMessageAPI) AddChat(w http.ResponseWriter, r *http.Request) {
	//拿不到token或是uuid则说明 认证异常
	token, err := common.HttpRequestGetJWTToken(r)
	if err != nil {
		common.HttpErrorErrorResponse(w, http.StatusUnauthorized, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}
	//通过token拿到userid
	userid, errs := utils.GetUserIDFromJWT(token)
	if errs != nil {
		common.HttpErrorErrorResponse(w, http.StatusUnauthorized, errorcode.ErrUserComm.FillMsg(errs.Error()))
		return
	}

	var chat model.ChatMessage

	err = common.HttpRequest2Struct(r, &chat)

	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	if chat.FromId != userid {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("fromid 与token中的不相同,无法新增聊天记录"))
		return
	}

	setChatUUID(&chat)

	if err := service.ChatMessage.AddChat(chat); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
	} else {
		common.HttpOKResponse(w, nil)
	}
}

func setChatUUID(chat *model.ChatMessage) {
	/*
	   聊天UUID说明，id1-id2-taskid-groupid 这4个数字的字符串
	   id1 是formid 和 toid 相比较小的
	   id2 是formid 和 toid 相比较大的

	   比如：fromid 88 , toid 66 ,taskid 99 ,groupid 2
	   ClassType = 1 用户聊天
	   uuid=66-88-0-0

	   classType = 2 task下记录,任务下的留言
	   uuid=0-0-99-0

	   classType = 3 群组聊天
	   uuid=0-0-0-2
	*/
	fromid := strconv.FormatUint(chat.FromId, 10)
	toid := strconv.FormatUint(chat.ToId, 10)
	strTaskid := strconv.FormatUint(chat.TaskID, 10)
	strGroupid := strconv.FormatUint(chat.GroupId, 10)

	var id1, id2 string

	if chat.FromId <= chat.ToId {
		id1 = fromid
		id2 = toid
	} else {
		id1 = toid
		id2 = fromid
	}

	switch chat.ClassType {
	case common.CHAT_CLASS_TYPE_USER_2_USER:
		{
			chat.ChatUuid = id1 + "-" + id2 + "-0-0"
		}
	case common.CHAT_CLASS_TYPE_TASK:
		{
			chat.ChatUuid = "0-0-" + strTaskid + "-0"

		}
	case common.CHAT_CLASS_TYPE_GROUP:
		{
			chat.ChatUuid = "0-0-0" + strGroupid
		}
	}

}

func (us *ChatMessageAPI) GetChatListByIds(w http.ResponseWriter, r *http.Request) {

	var pageInfo request.ParamGetChatMessage
	err := common.HttpRequest2Struct(r, &pageInfo)

	if err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))
		return
	}

	var userid uint64
	if pageInfo.ClassType == common.CHAT_CLASS_TYPE_USER_2_USER {
		//拿不到token或是uuid则说明 认证异常
		token, err := common.HttpRequestGetJWTToken(r)
		if err != nil {
			common.HttpErrorErrorResponse(w, http.StatusUnauthorized, errorcode.ErrUserComm.FillMsg(err.Error()))
			return
		}
		//通过token拿到userid
		userid, err = utils.GetUserIDFromJWT(token)
		if err != nil {
			common.HttpErrorErrorResponse(w, http.StatusUnauthorized, errorcode.ErrUserComm.FillMsg(err.Error()))
			return
		}
	}

	if pageInfo.FromId != userid {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg("fromid 与token中的不相同,无法获取聊天记录"))
		return
	}

	chattmp := model.ChatMessage{
		FromId:    pageInfo.FromId,
		ToId:      pageInfo.ToId,
		TaskID:    pageInfo.TaskID,
		GroupId:   pageInfo.GroupId,
		ClassType: pageInfo.ClassType,
	}
	setChatUUID(&chattmp)
	pageInfo.ChatUuid = chattmp.ChatUuid

	if list, total, err := service.ChatMessage.GetChatListByIds(pageInfo); err != nil {
		common.HttpOKErrorResponse(w, errorcode.ErrUserComm.FillMsg(err.Error()))

	} else {
		common.HttpOKResponse(w, common.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		})
	}

}

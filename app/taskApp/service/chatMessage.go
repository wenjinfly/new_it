package service

import (
	"new_it/app/taskApp/api/request"
	"new_it/app/taskApp/model"
	"new_it/global"
)

type CHAT_MESSAGE struct{}

func (us *CHAT_MESSAGE) AddChat(chat model.ChatMessage) error {
	var err = global.GLB_DB.Create(&chat).Error
	return err
}

func (us *CHAT_MESSAGE) GetChatListByIds(info request.ParamGetChatMessage) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GLB_DB.Model(&model.ChatMessage{})

	var chatList []model.ChatMessage
	err = db.Where("chat_uuid = ?", info.ChatUuid).Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&chatList).Error

	return chatList, total, err
}

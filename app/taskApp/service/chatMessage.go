package service

import (
	"new_it/app/taskApp/model"
	"new_it/global"
)

type CHAT_MESSAGE struct{}

func (us *CHAT_MESSAGE) AddChat(chat model.ChatMessage) error {
	var err = global.GLB_DB.Create(&chat).Error
	return err
}

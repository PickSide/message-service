package service

import (
	"message-service/internal/database"
	"message-service/pkg/models"
)

func DeleteMessage(msgID string) error {
	return database.DeleteMessageByID(msgID)
}
func GetMessage(msgID string) (*models.Message, error) {
	return database.GetMessageByID(msgID)
}
func LoadMessages(chatroomID string) (*[]models.Message, error) {
	return database.GetChatroomMessages(chatroomID)
}
func SendMessage(req models.CreateSendMessageStruct) (*string, error) {
	return database.CreateMessage(req)
}

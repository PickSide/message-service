package service

import (
	sqlutils "message-service/internal/sql"
	"message-service/pkg/models"
)

func DeleteMessage(msgID string) error {
	return sqlutils.DeleteMessageByID(msgID)
}
func GetMessage(msgID string) (*models.Message, error) {
	return sqlutils.GetMessageByID(msgID)
}
func LoadMessages(chatroomID string) (*[]models.Message, error) {
	return sqlutils.GetChatroomMessages(chatroomID)
}
func SendMessage(req models.CreateSendMessageStruct) (*string, error) {
	return sqlutils.CreateMessage(req)
}

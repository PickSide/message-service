package database

import (
	"errors"
	"message-service/pkg/models"
	"message-service/pkg/utils"
)

func GetMessageByID(msgID string) (*models.Message, error) {
	var message models.Message

	uint64MessageID, err := utils.StringToUint64(msgID)
	if err != nil {
		return nil, errors.New("GetMessageByID - Error during conversion (Malformed ID)")
	}

	err = GetClient().QueryRow(`SELECT chatroom_id, content, delivered, sent_at, sender_id FROM messages WHERE id = ?`, uint64MessageID).Scan(
		&message.ChatroomID,
		&message.Content,
		&message.Delivered,
		&message.SentAt,
		&message.SenderID,
	)

	message.IDString = utils.Uint64ToString(message.ID)
	message.ChatroomIDString = utils.Uint64ToString(message.ChatroomID)
	message.SenderIDString = utils.Uint64ToString(message.SenderID)

	return &message, err
}

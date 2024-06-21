package database

import (
	"errors"
	"message-service/pkg/models"
	"message-service/pkg/utils"
)

func GetChatroomMessages(chatroomID string) (*[]models.Message, error) {
	var messages []models.Message

	uint64ChatroomID, err := utils.StringToUint64(chatroomID)
	if err != nil {
		return nil, errors.New("GetMessageByID - Error during conversion (Malformed ID)")
	}

	rows, err := GetClient().Query(`SELECT chatroom_id, content, delivered, sent_at, sender_id FROM messages WHERE chatroom_id = ?`, uint64ChatroomID)
	if err != nil {
		return nil, errors.New("GetMessagesForChatroom - Error fetching message for the given chatroom")
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Message

		err := rows.Scan(
			&message.ChatroomID,
			&message.Content,
			&message.Delivered,
			&message.SentAt,
			&message.SenderID,
		)
		if err != nil {
			return nil, err
		}

		message.ChatroomIDString = utils.Uint64ToString(message.ChatroomID)
		message.IDString = utils.Uint64ToString(message.ID)
		message.SenderIDString = utils.Uint64ToString(message.SenderID)

		messages = append(messages, message)
	}

	return &messages, nil
}

package sqlutils

import (
	"errors"
	"message-service/internal/database"
	"message-service/pkg/models"
	"message-service/pkg/utils"
)

func CreateMessage(req models.CreateSendMessageStruct) (*string, error) {
	uint64ChatroomID, err := utils.StringToUint64(req.ChatroomID)
	if err != nil {
		return nil, errors.New("CreateMessage - Error during conversion (Malformed ID)")
	}
	uint64SenderID, err := utils.StringToUint64(req.SenderID)
	if err != nil {
		return nil, errors.New("CreateMessage - Error during conversion (Malformed ID)")
	}

	tx, err := database.GetClient().Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	var exists bool
	err = tx.QueryRow(`SELECT EXISTS(SELECT 1 FROM chatrooms WHERE id = ?)`, uint64ChatroomID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Chatroom not found")
	}

	result, err := tx.Exec(
		`INSERT INTO messages (chatroom_id, content, sender_id) VALUES (?, ?, ?)`,
		uint64ChatroomID,
		req.Content,
		uint64SenderID,
	)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	messageID := utils.Int64ToString(lastInsertID)

	return &messageID, nil
}

func DeleteMessageByID(msgID string) error {
	uint64MessageID, err := utils.StringToUint64(msgID)
	if err != nil {
		return errors.New("DeleteMessageByID - Error during conversion (Malformed ID)")
	}

	_, err = database.GetClient().Exec(`DELETE FROM messages WHERE id = ?`, uint64MessageID)

	return err
}

func GetMessageByID(msgID string) (*models.Message, error) {
	var message models.Message

	uint64MessageID, err := utils.StringToUint64(msgID)
	if err != nil {
		return nil, errors.New("GetMessageByID - Error during conversion (Malformed ID)")
	}

	err = database.GetClient().QueryRow(`SELECT chatroom_id, content, delivered, sent_at, sender_id FROM messages WHERE id = ?`, uint64MessageID).Scan(
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

func GetChatroomMessages(chatroomID string) (*[]models.Message, error) {
	var messages []models.Message

	uint64ChatroomID, err := utils.StringToUint64(chatroomID)
	if err != nil {
		return nil, errors.New("GetMessageByID - Error during conversion (Malformed ID)")
	}

	rows, err := database.GetClient().Query(`SELECT chatroom_id, content, delivered, sent_at, sender_id FROM messages WHERE chatroom_id = ?`, uint64ChatroomID)
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

package database

import (
	"errors"
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

	tx, err := GetClient().Begin()
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

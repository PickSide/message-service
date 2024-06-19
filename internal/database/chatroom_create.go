package database

import (
	"message-service/pkg/utils"
)

func InstantiateChatroom(userIDs []string) (*string, error) {
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

	result, err := tx.Exec(`INSERT INTO chatrooms (name, created_at) VALUES (DEFAULT, DEFAULT)`)
	if err != nil {
		return nil, err
	}

	chatroomID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	for _, ID := range userIDs {
		_, err := tx.Exec(`INSERT INTO users_chatrooms (user_id, chatroom_id) VALUES (?, ?)`, ID, chatroomID)
		if err != nil {
			return nil, err
		}
	}

	chatroomIDString := utils.Int64ToString(chatroomID)

	return &chatroomIDString, nil
}

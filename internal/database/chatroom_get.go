package database

import (
	"errors"
	"message-service/pkg/models"
	"message-service/pkg/utils"
)

func LoadChatroomByID(chatroomID string) (*models.Chatroom, error) {
	var chatroom models.Chatroom

	uint64ChatroomID, err := utils.StringToUint64(chatroomID)
	if err != nil {
		return nil, errors.New("LoadChatroomByID - Error during conversion (Malformed ID)")
	}

	err = GetClient().QueryRow(`SELECT id, name, number_of_messages FROM chatrooms WHERE id = ?`, uint64ChatroomID).Scan(
		&chatroom.ID,
		&chatroom.Name,
		&chatroom.NumberOfMessages,
	)

	chatroom.IDString = utils.Uint64ToString(chatroom.ID)

	return &chatroom, err
}

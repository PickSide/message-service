package database

import (
	"errors"
	"message-service/pkg/models"
	"message-service/pkg/utils"
)

func LoadChatroomsForUser(userID string) (*[]models.Chatroom, error) {
	var chatrooms []models.Chatroom

	uint64UserID, err := utils.StringToUint64(userID)
	if err != nil {
		return nil, errors.New("GetMessageByID - Error during conversion (Malformed ID)")
	}

	rows, err := GetClient().Query(`
		SELECT cr.id, cr.name, cr.number_of_messages
		FROM chatrooms AS cr
		INNER JOIN users_chatrooms AS cru ON cr.id = cru.chatroom_id
		WHERE cru.user_id = ?;
	`, uint64UserID)
	if err != nil {
		return nil, errors.New("GetMessagesForChatroom - Error fetching message for the given chatroom")
	}
	defer rows.Close()

	for rows.Next() {
		var chatroom models.Chatroom

		err := rows.Scan(
			&chatroom.ID,
			&chatroom.Name,
			&chatroom.NumberOfMessages,
		)
		if err != nil {
			return nil, err
		}

		chatroom.IDString = utils.Uint64ToString(chatroom.ID)

		chatrooms = append(chatrooms, chatroom)
	}

	return &chatrooms, nil
}

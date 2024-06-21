package database

import (
	"errors"
	"message-service/pkg/models"
	"message-service/pkg/utils"
)

func GetChatroomParticipants(chatroomParticipantID string) (*[]models.User, error) {
	uint64ChatroomParticipantID, err := utils.StringToUint64(chatroomParticipantID)
	if err != nil {
		return nil, errors.New("GetChatroomParticipants - Error during conversion (Malformed ID)")
	}

	rows, err := GetClient().Query(`
		SELECT
			u.id,
			u.avatar,
			u.display_name,
			u.full_name
		FROM users_chatrooms
		LEFT JOIN users AS u ON users_chatrooms.user_id = u.id
		WHERE users_chatrooms.chatroom_id = ?;`,
		uint64ChatroomParticipantID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []models.User

	for rows.Next() {
		var participant models.User

		err := rows.Scan(
			&participant.ID,
			&participant.Avatar,
			&participant.DisplayName,
			&participant.FullName,
		)
		if err != nil {
			return nil, err
		}

		participant.IDString = utils.Uint64ToString(participant.ID)

		participants = append(participants, participant)
	}

	return &participants, nil
}

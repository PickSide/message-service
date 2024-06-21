package database

import (
	"errors"
	"message-service/pkg/utils"
)

func GetParticipantsID(chatroomID string) (*[]string, error) {
	var participantIDs []string

	uint64ChatroomID, err := utils.StringToUint64(chatroomID)
	if err != nil {
		return nil, errors.New("GetParticipantsID - Error during conversion (Malformed ID)")
	}

	rows, err := GetClient().Query(`
		SELECT u.id
		FROM users AS u
		INNER JOIN users_chatrooms AS uc ON u.id = uc.user_id
		WHERE uc.chatroom_id = ?;
	`, uint64ChatroomID)
	if err != nil {
		return nil, errors.New("GetParticipantsID - Error fetching message for the given chatroom")
	}
	defer rows.Close()

	for rows.Next() {
		var participantID uint64

		err := rows.Scan(&participantID)
		if err != nil {
			return nil, err
		}

		participantIDs = append(participantIDs, utils.Uint64ToString(participantID))
	}

	return &participantIDs, nil
}

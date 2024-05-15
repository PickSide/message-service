package sqlutils

import (
	"errors"
	"message-service/internal/database"
	"message-service/pkg/models"
	"message-service/pkg/utils"
	"sort"
	"strings"
)

func GetParticipantsID(chatroomID string) (*[]string, error) {
	var participantIDs []string

	uint64ChatroomID, err := utils.StringToUint64(chatroomID)
	if err != nil {
		return nil, errors.New("GetParticipantsID - Error during conversion (Malformed ID)")
	}

	rows, err := database.GetClient().Query(`
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
func LoadChatroomsForUser(userID string) (*[]models.Chatroom, error) {
	var chatrooms []models.Chatroom

	uint64UserID, err := utils.StringToUint64(userID)
	if err != nil {
		return nil, errors.New("GetMessageByID - Error during conversion (Malformed ID)")
	}

	rows, err := database.GetClient().Query(`
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
func LoadChatroomByParticipantIDs(userIDs []string) (*models.Chatroom, error) {
	var chatroom models.Chatroom

	sort.Strings(userIDs)
	userIDsString := strings.Join(userIDs, ",")

	err := database.GetClient().QueryRow(`
		SELECT cr.id, cr.name, cr.number_of_messages
		FROM chatrooms AS cr
		JOIN users_chatrooms AS cru ON cr.id = cru.chatroom_id
		GROUP BY cr.id
		HAVING COUNT(DISTINCT cru.user_id) = ?
		AND GROUP_CONCAT(DISTINCT cru.user_id ORDER BY cru.user_id) = ?
	`, len(userIDs), userIDsString).Scan(
		&chatroom.ID,
		&chatroom.Name,
		&chatroom.NumberOfMessages,
	)

	chatroom.IDString = utils.Uint64ToString(chatroom.ID)

	return &chatroom, err
}
func LoadChatroomByID(chatroomID string) (*models.Chatroom, error) {
	var chatroom models.Chatroom

	uint64ChatroomID, err := utils.StringToUint64(chatroomID)
	if err != nil {
		return nil, errors.New("LoadChatroomByID - Error during conversion (Malformed ID)")
	}

	err = database.GetClient().QueryRow(`SELECT id, name, number_of_messages FROM chatrooms WHERE id = ?`, uint64ChatroomID).Scan(
		&chatroom.ID,
		&chatroom.Name,
		&chatroom.NumberOfMessages,
	)

	chatroom.IDString = utils.Uint64ToString(chatroom.ID)

	return &chatroom, err
}
func InstantiateChatroom(userIDs []string) (*string, error) {
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

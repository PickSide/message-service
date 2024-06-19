package database

import (
	"message-service/pkg/models"
	"message-service/pkg/utils"
	"sort"
	"strings"
)

func LoadChatroomByParticipantIDs(userIDs []string) (*models.Chatroom, error) {
	var chatroom models.Chatroom

	sort.Strings(userIDs)
	userIDsString := strings.Join(userIDs, ",")

	err := GetClient().QueryRow(`
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

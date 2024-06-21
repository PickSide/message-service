package database

import (
	"errors"
	"message-service/pkg/utils"
)

func DeleteMessageByID(msgID string) error {
	uint64MessageID, err := utils.StringToUint64(msgID)
	if err != nil {
		return errors.New("DeleteMessageByID - Error during conversion (Malformed ID)")
	}

	_, err = GetClient().Exec(`DELETE FROM messages WHERE id = ?`, uint64MessageID)

	return err
}

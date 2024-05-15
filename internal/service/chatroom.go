package service

import (
	sqlutils "message-service/internal/sql"
	"message-service/pkg/models"
)

func GetParticipantsID(chatroomID string) (*[]string, error) {
	return sqlutils.GetParticipantsID(chatroomID)
}
func LoadChatrooms(participantID string) (*[]models.Chatroom, error) {
	chatrooms, err := sqlutils.LoadChatroomsForUser(participantID)
	if err != nil {
		return nil, err
	}
	for i, chatroom := range *chatrooms {
		participants, err := sqlutils.GetChatroomParticipants(chatroom.IDString)
		if err != nil {
			return nil, err
		}
		(*chatrooms)[i].Participants = *participants
	}
	return chatrooms, nil
}
func LoadChatroomFromParticipants(participantIDs []string) (*models.Chatroom, error) {
	chatroom, err := sqlutils.LoadChatroomByParticipantIDs(participantIDs)
	if err != nil {
		return nil, err
	}
	participants, err := sqlutils.GetChatroomParticipants(chatroom.IDString)
	if err != nil {
		return nil, err
	}
	chatroom.Participants = *participants

	return chatroom, nil
}
func LoadChatroomFromID(chatroomID string) (*models.Chatroom, error) {
	chatroom, err := sqlutils.LoadChatroomByID(chatroomID)
	if err != nil {
		return nil, err
	}
	participants, err := sqlutils.GetChatroomParticipants(chatroom.IDString)
	if err != nil {
		return nil, err
	}
	chatroom.Participants = *participants

	return chatroom, nil
}
func InstantiateChatroom(participantIDs []string) (*string, error) {
	return sqlutils.InstantiateChatroom(participantIDs)
}

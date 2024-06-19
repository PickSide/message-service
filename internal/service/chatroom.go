package service

import (
	"message-service/internal/database"
	"message-service/pkg/models"
)

func GetParticipantsID(chatroomID string) (*[]string, error) {
	return database.GetParticipantsID(chatroomID)
}
func LoadChatrooms(participantID string) (*[]models.Chatroom, error) {
	chatrooms, err := database.LoadChatroomsForUser(participantID)
	if err != nil {
		return nil, err
	}
	for i, chatroom := range *chatrooms {
		participants, err := database.GetChatroomParticipants(chatroom.IDString)
		if err != nil {
			return nil, err
		}
		(*chatrooms)[i].Participants = *participants
	}
	return chatrooms, nil
}
func LoadChatroomFromParticipants(participantIDs []string) (*models.Chatroom, error) {
	chatroom, err := database.LoadChatroomByParticipantIDs(participantIDs)
	if err != nil {
		return nil, err
	}
	participants, err := database.GetChatroomParticipants(chatroom.IDString)
	if err != nil {
		return nil, err
	}
	chatroom.Participants = *participants

	return chatroom, nil
}
func LoadChatroomFromID(chatroomID string) (*models.Chatroom, error) {
	chatroom, err := database.LoadChatroomByID(chatroomID)
	if err != nil {
		return nil, err
	}
	participants, err := database.GetChatroomParticipants(chatroom.IDString)
	if err != nil {
		return nil, err
	}
	chatroom.Participants = *participants

	return chatroom, nil
}
func InstantiateChatroom(participantIDs []string) (*string, error) {
	return database.InstantiateChatroom(participantIDs)
}

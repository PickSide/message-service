package websocket

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"message-service/internal/service"
	"message-service/pkg/models"

	"github.com/davecgh/go-spew/spew"
)

func ProcessIncomingMessage(msg models.SocketMessage) (*models.SocketResponse, error) {
	switch msg.EventType {
	case "connection:ready":
		return &models.SocketResponse{
			Message: "Channel ready to serve messaging",
		}, nil

	case "chatroom:open":
		var socketMsg models.LoadChatroomStruct
		if err := json.Unmarshal([]byte(msg.Content), &socketMsg); err != nil {
			return nil, err
		}
		spew.Dump("socketMsg", socketMsg)
		chatroom, err := service.LoadChatroomFromParticipants(socketMsg.ParticipantIDs)
		if err != nil && err == sql.ErrNoRows {
			chatroomID, err := service.InstantiateChatroom(socketMsg.ParticipantIDs)
			if err != nil {
				log.Println("there's an error from InstantiateChatroom")
				return nil, err
			}
			chatroom, err = service.LoadChatroomFromID(*chatroomID)
			if err != nil {
				log.Println("there's an error from LoadChatroomFromID")
				return nil, err
			}
		}
		return &models.SocketResponse{
			EventType: "chatroom:opened",
			Message:   "Loaded chatroom",
			Results:   chatroom,
		}, nil

	case "chatrooms:getall":
		var socketMsg models.LoadChatroomsStruct
		if err := json.Unmarshal([]byte(msg.Content), &socketMsg); err != nil {
			return nil, err
		}
		chatrooms, err := service.LoadChatrooms(socketMsg.ParticipantID)
		if err != nil {
			return nil, err
		}
		if len(*chatrooms) == 0 {
			return nil, errors.New("No chatrooms")
		}
		return &models.SocketResponse{
			EventType: "chatrooms:fetched",
			Message:   "Loaded chatrooms",
			Results:   chatrooms,
		}, nil

	case "chatroom:loadmessages":
		var socketMsg models.LoadMessagesStruct
		if err := json.Unmarshal([]byte(msg.Content), &socketMsg); err != nil {
			return nil, err
		}
		messages, err := service.LoadMessages(socketMsg.ChatroomID)
		if err != nil {
			return nil, err
		}
		if len(*messages) == 0 {
			return nil, errors.New("No messages")
		}
		return &models.SocketResponse{
			EventType: "chatroom:messages",
			Message:   "Loaded messages",
			Results:   messages,
		}, nil

	case "message:send":
		var messageDetails models.CreateSendMessageStruct
		if err := json.Unmarshal([]byte(msg.Content), &messageDetails); err != nil {
			return nil, err
		}
		messageID, err := service.SendMessage(messageDetails)
		if err != nil {
			return nil, err
		}
		message, err := service.GetMessage(*messageID)

		participantsInvolvedIDs, err := service.GetParticipantsID(message.ChatroomIDString)

		msg := models.SocketResponse{
			EventType: "message:received",
			Message:   "Loaded messages",
			Results:   message,
		}
		BroadcastMessageToChatroom(*participantsInvolvedIDs, msg)
		break

	case "message:delete":
		var msgID string
		if err := json.Unmarshal([]byte(msg.Content), &msgID); err != nil {
			return nil, err
		}
		return nil, service.DeleteMessage(msgID)
	}

	return nil, errors.New("Invalid format")
}

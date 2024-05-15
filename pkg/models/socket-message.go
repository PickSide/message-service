package models

import "encoding/json"

type SocketMessage struct {
	EventType string          `json:"eventType"`
	Content   json.RawMessage `json:"content"`
}
type SocketResponse struct {
	EventType string          `json:"eventType"`
	Results   any             `json:"results,omitempty"`
	Error     string          `json:"error,omitempty"`
	Message   string          `json:"message,omitempty"`
	Extra     json.RawMessage `json:"extra,omitempty"`
}
type CreateSendMessageStruct struct {
	Content    string `json:"content"`
	ChatroomID string `json:"chatroomId"`
	SenderID   string `json:"senderId"`
}
type LoadMessagesStruct struct {
	ChatroomID string `json:"chatroomId"`
}
type LoadChatroomsStruct struct {
	ParticipantID string `json:"participantId"`
}
type LoadChatroomStruct struct {
	ParticipantIDs []string `json:"participantIds"`
}

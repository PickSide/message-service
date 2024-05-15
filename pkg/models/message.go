package models

type Message struct {
	ID               uint64 `json:"-"`
	IDString         string `json:"id"`
	Content          string `json:"content"`
	Delivered        uint64 `json:"-"`
	ChatroomID       uint64 `json:"-"`
	ChatroomIDString string `json:"chatroomId"`
	SentAt           string `json:"sentAt"`
	SenderID         uint64 `json:"-"`
	SenderIDString   string `json:"senderId"`
}

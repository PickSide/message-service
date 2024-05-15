package models

type Chatroom struct {
	ID               uint64 `json:"-"`
	IDString         string `json:"id"`
	Name             string `json:"name"`
	NumberOfMessages string `json:"numberOfMessages"`
	Participants     []User `json:"participants"`
}

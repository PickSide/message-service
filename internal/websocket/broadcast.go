package websocket

import (
	"encoding/json"
	"fmt"
	"message-service/pkg/models"

	"github.com/gorilla/websocket"
)

func BroadcastMessageToClient(message models.SocketResponse, sender *websocket.Conn) {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling message:", err)
		return
	}

	for userID, conn := range Clients {
		if conn == sender {
			if err := conn.WriteMessage(websocket.TextMessage, messageJSON); err != nil {
				fmt.Println("Error broadcasting to client:", err)
				conn.Close()
				delete(Clients, userID)
			}
		}
	}
}
func BroadcastMessageToChatroom(participantsIDs []string, message models.SocketResponse) {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling message:", err)
		return
	}

	for _, ID := range participantsIDs {
		if conn, ok := Clients[ID]; ok {
			if err := conn.WriteMessage(websocket.TextMessage, messageJSON); err != nil {
				fmt.Println("Error sending message:", err)
			}
		}
	}
}

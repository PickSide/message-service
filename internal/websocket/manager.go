package websocket

import "fmt"

func HandleMessages() {
	for {
		msg := <-Broadcast

		for userID, conn := range Clients {
			err := conn.WriteJSON(msg)
			if err != nil {
				fmt.Println(err)
				conn.Close()
				delete(Clients, userID)
			}
		}
	}
}

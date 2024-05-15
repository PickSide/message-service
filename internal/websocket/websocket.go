package websocket

import (
	"fmt"
	"log"
	"message-service/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var Clients = make(map[string]*websocket.Conn)
var Broadcast = make(chan models.SocketMessage)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(g *gin.Context) {
	conn, err := upgrader.Upgrade(g.Writer, g.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	claims, exists := g.Get("claims")

	if !exists {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	userClaims := claims.(models.UserClaims)

	Clients[userClaims.UserID] = conn

	fmt.Printf("User %s connected\n", userClaims.UserID)

	for {
		var socketMsg models.SocketMessage

		err := conn.ReadJSON(&socketMsg)
		if err != nil {
			log.Println(err)
			fmt.Println(err)
			delete(Clients, userClaims.UserID)
			return
		}

		resp, err := ProcessIncomingMessage(socketMsg)
		if err != nil {
			continue
		}

		BroadcastMessageToClient(*resp, conn)
	}
}

package main

import (
	"log"
	"message-service/internal/database"
	"message-service/internal/rest"
	"message-service/internal/vault"
	"message-service/internal/websocket"
	"message-service/middlewares"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dsn := vault.Envars["DSN"].(string)
	database.Initialize(dsn)
	defer database.Close()

	go websocket.HandleMessages()

	gin.SetMode(os.Getenv("GIN_MODE"))
	g := gin.Default()
	g.Use(cors.New(buildCors()))

	as := g.Group("/message-service")

	as.GET("/ws", middlewares.ValidateUser(), websocket.HandleConnections)

	//Health
	as.GET("/health", rest.GetHealth)

	PrintServiceInformation()

	if err := g.Run(":8083"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func PrintServiceInformation() {
	log.Printf("Mode %s", os.Getenv("GIN_MODE"))
	log.Printf("Service name: %s", os.Getenv("GIN_MODE"))
	log.Printf("Version: %s", os.Getenv("SERVICE_VERSION"))
}

func buildCors() cors.Config {
	c := cors.DefaultConfig()
	c.AllowAllOrigins = false
	c.AllowCredentials = true
	c.AllowHeaders = []string{"Accept-Version", "Authorization", "Content-Type", "Origin", "X-Client-Version", "X-CSRF-Token", "X-Request-Id"}
	c.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	c.AllowWebSockets = true
	c.MaxAge = 24 * time.Hour

	c.AllowOriginFunc = func(origin string) bool {
		return true
	}
	return c
}

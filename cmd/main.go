package main

import (
	"log"
	"message-service/internal/api"
	"message-service/internal/database"
	"message-service/internal/env"
	"message-service/internal/vault"
	"message-service/internal/websocket"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize envs
	env.InitalizeEnvs()

	// Initialize database
	database.InitializeDB(vault.Envars["DSN"].(string))
	defer database.Close()

	go websocket.HandleMessages()

	gin.SetMode(env.GIN_MODE)
	g := gin.Default()
	g.Use(cors.New(buildCors()))

	g.GET("/ws", websocket.HandleConnections)

	//Health
	g.GET("/health", api.GetHealth)

	PrintServiceInformation()

	if err := g.Run(":8083"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func PrintServiceInformation() {
	log.Printf("Mode %s", env.GIN_MODE)
	log.Printf("Service name: %s", env.SERVICE_NAME)
	log.Printf("Version: %s", env.VERSION)
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

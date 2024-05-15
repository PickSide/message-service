package middlewares

import (
	"encoding/json"
	"fmt"
	"message-service/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateUser() gin.HandlerFunc {
	return func(g *gin.Context) {
		token := g.Query("token")

		req, err := http.NewRequest("GET", "http://localhost:8083/auth-service/verify-token", nil)
		if err != nil {
			g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Failed to connect to auth-service",
				"success": false,
			})
			return
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Can't call the auth-service",
				"success": false,
			})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is invalid or expired",
				"success": false,
			})
			return
		}

		var userClaims models.UserClaims

		if err := json.NewDecoder(resp.Body).Decode(&userClaims); err != nil {
			g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Failed to decode user claims",
				"success": false,
			})
			return
		}

		g.Set("claims", userClaims)

		g.Next()
	}
}

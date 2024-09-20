package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pharmacy-system/internal/utils"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Получаем токен из заголовка
		authToken := c.Request.Header.Get("Authorization")

		if authToken == "" {
			log.Printf("Authenticate - authToken is empty")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Получаем ID пользователя из JWT
		userID, err := utils.ValidateJWT(authToken)
		if err != nil {
			log.Printf("Authenticate - utils.ValidateJWT error: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Добавляем ID пользователя в контекст
		c.Set("UserID", userID)

		c.Next()
	}
}

package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"example/authzed/initializers"
	"example/authzed/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header)
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find account with token stub
		var account models.Account
		initializers.DB.First(&account, claims["stub"])
		if account.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("account", account)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}

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

func _NoAuth(c *gin.Context) {
	c.Set("auth", false)
	c.Next()
}

func CheckAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		_NoAuth(c)
		return
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
			_NoAuth(c)
			return
		}

		// Find account with token stub
		var account models.Account
		initializers.DB.First(&account, claims["stub"])
		if account.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("account", account)
		c.Next()
	} else {
		_NoAuth(c)
		return
	}
}

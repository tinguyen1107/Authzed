package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExtractBody[T interface{}](c *gin.Context) {
	var body T
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	c.Set("body", body)
}

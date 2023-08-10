package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckSpiceDbPermission(c *gin.Context) {
	fmt.Println("Check Permission from SpiceDB")

	// Create Subject

	// Create Resource

	// Call to check permission

	c.Next()
}

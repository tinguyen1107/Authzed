package controllers

import (
	"net/http"

	"example/authzed/initializers"
	"example/authzed/models"

	"github.com/gin-gonic/gin"
)

func CreateFolder(c *gin.Context) {
	var body struct {
		Name     string
		ParentId uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	doc := models.Folder{
		Name:     body.Name,
		ParentId: &body.ParentId,
	}
	result := initializers.DB.Create(&doc)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create folder",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create folder successfully"})
}

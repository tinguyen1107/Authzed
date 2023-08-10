package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"example/authzed/initializers"
	"example/authzed/models"

	"github.com/gin-gonic/gin"
)

func CreateDocument(c *gin.Context) {
	var body struct {
		Name     string
		Content  string // Raw content
		ParentId uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	doc := models.Document{
		Name:     body.Name,
		Content:  base64.StdEncoding.EncodeToString([]byte(body.Content)),
		ParentId: &body.ParentId,
	}
	result := initializers.DB.Create(&doc)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create document",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create document successfully"})
}

func EditDocument(c *gin.Context) {
	var body struct {
		Name     string
		Content  string // Raw content
		ParentId uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	doc := models.Document{
		Name:     body.Name,
		Content:  base64.StdEncoding.EncodeToString([]byte(body.Content)),
		ParentId: &body.ParentId,
	}
	result := initializers.DB.Create(&doc)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create document",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"document": fmt.Sprint(doc)})
}

func GetDocument(c *gin.Context) {
	var body struct {
		Name     string
		Content  string // Raw content
		ParentId uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	doc := models.Document{
		Name:     body.Name,
		Content:  base64.StdEncoding.EncodeToString([]byte(body.Content)),
		ParentId: &body.ParentId,
	}
	result := initializers.DB.Create(&doc)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create document",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"document": fmt.Sprint(doc)})
}

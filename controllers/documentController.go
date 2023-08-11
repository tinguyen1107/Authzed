package controllers

import (
	"encoding/base64"
	"net/http"

	"example/authzed/initializers"
	"example/authzed/models"
	"example/authzed/services"

	"github.com/gin-gonic/gin"
)

type CreateDocumentBody struct {
	Name     string
	Content  string // Raw content
	ParentId uint
}

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

	_, err := services.CreateDocument(body.Name, body.Content, body.ParentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "create document successfully"})
}

func EditDocument(c *gin.Context) {
	var body struct {
		Id      uint
		Content string // Raw content
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var doc models.Document
	initializers.DB.First(&doc, body.Id)

	if doc.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Document Id",
		})
		return
	}

	doc.Content = base64.StdEncoding.EncodeToString([]byte(body.Content))

	result := initializers.DB.Save(&doc)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to save document",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update document successfully"})
}

func DeleteDocument(c *gin.Context) {
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

	initializers.DB.Delete(&models.Document{}, "name = ? AND parent_id = ?", body.Name, body.ParentId)
	c.JSON(http.StatusOK, gin.H{"message": "delete document successfully"})
}

func GetDocument(c *gin.Context) {
	var doc models.Document
	initializers.DB.First(&doc, c.Param("id"))

	if doc.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "document is not found",
		})
		return
	}
	parsed, err := base64.StdEncoding.DecodeString(doc.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "document's content parse failed",
		})
		return
	}
	doc.Content = string(parsed)

	c.JSON(http.StatusOK, gin.H{"document": doc})
}

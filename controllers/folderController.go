package controllers

import (
	"net/http"

	"example/authzed/initializers"
	"example/authzed/models"
	"example/authzed/services"

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

	_, err := services.CreateFolder(body.Name, body.ParentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "create folder successfully"})
}

func EditFolder(c *gin.Context) {
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

	var doc models.Folder
	initializers.DB.First(&doc, body.Id)

	if doc.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Folder Id",
		})
		return
	}

	result := initializers.DB.Save(&doc)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to save folder",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "update folder successfully"})
}

func DeleteFolder(c *gin.Context) {
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

	initializers.DB.Delete(&models.Folder{}, "name = ? AND parent_id = ?", body.Name, body.ParentId)
	c.JSON(http.StatusOK, gin.H{"message": "delete folder successfully"})
}

func GetFolder(c *gin.Context) {
	var folder models.Folder
	initializers.DB.Model(&models.Folder{}).Preload("Documents").Preload("SubFolders").Find(&folder, c.Param("id"))

	if folder.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "folder is not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"folder": folder})
}

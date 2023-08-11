package services

import (
	"errors"
	"net/http"

	"example/authzed/initializers"
	"example/authzed/models"

	"github.com/gin-gonic/gin"
)

func CreateFolder(name string, parentId uint) (*models.Folder, error) {
	var parentFolder models.Folder
	initializers.DB.First(&parentFolder, parentId)

	folder := models.Folder{
		Name:     name,
		ParentId: &parentId,
	}
	err := initializers.DB.Model(&parentFolder).Association("SubFolders").Append(&folder)
	if err != nil {
		return nil, errors.New("Failed to create folder")
	}
	return &folder, nil
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

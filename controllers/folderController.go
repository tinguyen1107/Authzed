package controllers

import (
	"fmt"
	"net/http"

	"example/authzed/dtos"
	"example/authzed/initializers"
	"example/authzed/models"
	"example/authzed/services"
	"example/authzed/utils"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/gin-gonic/gin"
)

func CreateFolder(c *gin.Context) {
	var body dtos.CreateFolderDto
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	folder, err := services.CreateFolder(body.Name, body.ParentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	// Write relationship
	folderObject := utils.CreateFolderObject(folder.ID)
	parentfolderSubject := &v1.SubjectReference{Object: utils.CreateFolderObject(*folder.ParentId)}
	relationship := v1.Relationship{
		Resource: folderObject,
		Relation: "super_folder",
		Subject:  parentfolderSubject,
	}

	res, error := initializers.SpiceClient.WriteRelationships(c, &v1.WriteRelationshipsRequest{
		Updates: []*v1.RelationshipUpdate{
			{
				Operation:    v1.RelationshipUpdate_OPERATION_TOUCH,
				Relationship: &relationship,
			},
		},
	})
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to set ownership",
		})
		return
	}
	fmt.Println(res)

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

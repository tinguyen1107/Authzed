package middleware

import (
	"net/http"

	"example/authzed/dtos"
	"example/authzed/initializers"
	"example/authzed/models"
	"example/authzed/utils"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/gin-gonic/gin"
)

func RequireCreateFolderPermission(c *gin.Context) {
	account, _ := c.Get("account")
	if account == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	body, _ := c.Get("body")
	if body == nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	accountObject := &v1.SubjectReference{Object: utils.CreateAccountObject(account.(models.Account).ID)}
	folderObject := utils.CreateFolderObject(body.(dtos.CreateFolderDto).ParentId)

	resp, err := initializers.SpiceClient.CheckPermission(c, &v1.CheckPermissionRequest{
		Resource:   folderObject,
		Permission: "write",
		Subject:    accountObject,
	})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "Failed to check permission: " + err.Error()})
	}

	if resp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_NO_PERMISSION {
		c.AbortWithStatusJSON(500, gin.H{"message": "Do not have permission to do this"})
	}

	c.Next()
}

func RequireCreateDocumentPermission(c *gin.Context) {
	account, _ := c.Get("account")
	if account == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	body, _ := c.Get("body")
	if body == nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	accountObject := &v1.SubjectReference{Object: utils.CreateAccountObject(account.(models.Account).ID)}
	folderObject := utils.CreateFolderObject(body.(dtos.CreateDocumentDto).ParentId)

	resp, err := initializers.SpiceClient.CheckPermission(c, &v1.CheckPermissionRequest{
		Resource:   folderObject,
		Permission: "write",
		Subject:    accountObject,
	})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "Failed to check permission: " + err.Error()})
	}

	if resp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_NO_PERMISSION {
		c.AbortWithStatusJSON(500, gin.H{"message": "Do not have permission to do this"})
	}

	c.Next()
}

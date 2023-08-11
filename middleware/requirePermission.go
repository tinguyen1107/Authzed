package middleware

import (
	"net/http"
	"os"
	"strconv"

	"example/authzed/controllers"
	"example/authzed/initializers"
	"example/authzed/models"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/gin-gonic/gin"
)

func requirePermission(c *gin.Context) {
	// c.Request.URL.Path

	c.Next()
}

func RequireCreateFolderPermission(c *gin.Context) {
	account, _ := c.Get("account")
	if account == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	body, _ := c.Get("body")
	if body == nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	accountObject := &v1.SubjectReference{Object: &v1.ObjectReference{
		ObjectType: os.Getenv("SPICE_DB_PREFIX") + "/user",
		ObjectId:   strconv.FormatUint(uint64(account.(models.Account).ID), 10),
	}}

	folderObject := &v1.ObjectReference{
		ObjectType: os.Getenv("SPICE_DB_PREFIX") + "/document",
		ObjectId:   strconv.FormatUint(uint64(body.(controllers.CreateDocumentBody).ParentId), 10),
	}
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

package main

import (
	// "context"
	"fmt"
	// "log"
	// "net/http"

	"example/authzed/controllers"
	"example/authzed/initializers"
	"example/authzed/middleware"

	"github.com/gin-gonic/gin"
	// v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

const SPICEDB_PREFIX = "ntrongtin11702_tutorial/"

func execute(c *gin.Context) {
	fmt.Println("Execute the function")
}

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	initializers.ConnectToSpiceDb()

	r := gin.Default()

	// Connect to spice db

	// Authentication
	r.POST("/account/signup", controllers.Signup)
	r.POST("/account/login", controllers.Login)
	r.POST("/account/validate", middleware.RequireAuth, controllers.Validate)

	// Document interaction
	r.POST("/document/create", middleware.RequireAuth, controllers.CreateDocument)
	r.POST("/document/edit/:id", middleware.RequireAuth, controllers.EditDocument)
	r.POST("/document/delete/:id", middleware.RequireAuth, controllers.DeleteDocument)
	r.GET("/document/get/:id", middleware.CheckAuth, controllers.GetDocument)

	// Folder interaction
	r.POST("/folder/create", middleware.RequireAuth, controllers.CreateFolder)
	r.POST("/folder/edit/:id", middleware.RequireAuth, controllers.EditFolder)
	r.POST("/folder/delete/:id", middleware.RequireAuth, controllers.DeleteFolder)
	r.GET("/folder/get/:id", middleware.CheckAuth, controllers.GetFolder)

	// r.GET("/document/:id", middleware.CheckAuth, func(c *gin.Context) {
	// 	emilia := &v1.SubjectReference{Object: &v1.ObjectReference{
	// 		ObjectType: SPICEDB_PREFIX + "user",
	// 		ObjectId:   "",
	// 	}}
	// 	firstPost := &v1.ObjectReference{
	// 		ObjectType: SPICEDB_PREFIX + "document",
	// 		ObjectId:   c.Param("id"),
	// 	}
	// 	resp, err := initializers.SpiceClient.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
	// 		Resource:   firstPost,
	// 		Permission: "view",
	// 		Subject:    emilia,
	// 	})
	// 	if err != nil {
	// 		log.Fatalf("failed to check permission: %s", err)
	// 	}
	//
	// 	if resp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
	// 		log.Println("allowed!")
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.POST("/document/:id/comment", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.POST("/document/:id/edit", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	//
	// r.GET("/folder/:id", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.POST("/folder/:id/comment", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.POST("/folder/:id/edit", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//
	// emilia := &v1.SubjectReference{Object: &v1.ObjectReference{
	// 	ObjectType: "ntrongtin11702_tutorial/user",
	// 	ObjectId:   "emilia",
	// }}
	//
	// firstPost := &v1.ObjectReference{
	// 	ObjectType: "ntrongtin11702_tutorial/document",
	// 	ObjectId:   "1",
	// }
	//
	// if err != nil {
	// 	return
	// }
	//
	// resp, err := client.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
	// 	Resource:   firstPost,
	// 	Permission: "view",
	// 	Subject:    emilia,
	// })
	// if err != nil {
	// 	log.Fatalf("failed to check permission: %s", err)
	// }
	//
	// if resp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
	// 	log.Println("allowed!")
	// }
}

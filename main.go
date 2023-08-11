package main

import (
	"fmt"
	"log"

	"example/authzed/controllers"
	"example/authzed/dtos"
	"example/authzed/initializers"
	"example/authzed/middleware"

	"github.com/gin-gonic/gin"
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

	// Authentication
	r.POST("/account/signup", controllers.Signup)
	r.POST("/account/login", controllers.Login)
	r.POST("/account/validate", middleware.RequireAuth, controllers.Validate)

	// Document interaction
	r.POST("/document/create", middleware.RequireAuth, middleware.ExtractBody[dtos.CreateDocumentDto], middleware.RequireCreateDocumentPermission, controllers.CreateDocument)
	r.POST("/document/edit/:id", middleware.RequireAuth, controllers.EditDocument)
	r.POST("/document/delete/:id", middleware.RequireAuth, controllers.DeleteDocument)
	r.GET("/document/get/:id", middleware.CheckAuth, controllers.GetDocument)

	// Folder interaction
	r.POST("/folder/create", middleware.RequireAuth, middleware.ExtractBody[dtos.CreateFolderDto], middleware.RequireCreateFolderPermission, controllers.CreateFolder)
	r.POST("/folder/edit/:id", middleware.RequireAuth, controllers.EditFolder)
	r.POST("/folder/delete/:id", middleware.RequireAuth, controllers.DeleteFolder)
	r.GET("/folder/get/:id", middleware.CheckAuth, controllers.GetFolder)

	// Management
	// r.POST("/:relation/add/:id", middleware.RequireAuth, controllers.CreateFolder)
	// r.POST("/:relation/remove/:id", middleware.RequireAuth, controllers.EditFolder)
	// r.GET("/:relation/check/:id", middleware.RequireAuth, controllers.DeleteFolder)
	// r.GET("/:relation/list/:objectType/:id", middleware.CheckAuth, controllers.GetFolder)

	err := r.Run()
	if err != nil {
		log.Fatal("Can't start server")
	}
}

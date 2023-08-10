package main

import (
	"context"
	"example/authzed/controllers"
	"example/authzed/initializers"
	"example/authzed/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
)

const SPICEDB_PREFIX = "ntrongtin11702_tutorial/"

func execute(c *gin.Context) {
	fmt.Println("Execute the function")
}

func main() {
	initializers.LoadEnvVariables()

	r := gin.Default()

	// Connect to spice db
	systemCerts, err := grpcutil.WithSystemCerts(grpcutil.VerifyCA)
	if err != nil {
		log.Fatalf("Unable to initialize system certs: %s", err)
	}
	client, err := authzed.NewClient(
		"grpc.authzed.com:443",
		systemCerts,
		grpcutil.WithBearerToken(
			"tc_golang_server_def_59ad33598ef60d6b7d0a5c935bc33ceeb3dbb2f84d76884a097eb65f4c9684b40c4978961e91cf692140bb4764e10a2491169e8089b6a045469393b8bad756e7",
		),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}

	r.POST("/account/signup", controllers.Signup)
	r.POST("/account/login", controllers.Login)

	r.POST("/account/validate", middleware.RequrieAuth, func(c *gin.Context) {
		account, _ := c.Get("account")

		c.JSON(http.StatusOK, gin.H{"message": account})
	})

	r.GET("/document/:id", middleware.CheckAuth, func(c *gin.Context) {
		emilia := &v1.SubjectReference{Object: &v1.ObjectReference{
			ObjectType: SPICEDB_PREFIX + "user",
			ObjectId:   "",
		}}
		firstPost := &v1.ObjectReference{
			ObjectType: SPICEDB_PREFIX + "document",
			ObjectId:   c.Param("id"),
		}
		resp, err := client.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
			Resource:   firstPost,
			Permission: "view",
			Subject:    emilia,
		})
		if err != nil {
			log.Fatalf("failed to check permission: %s", err)
		}

		if resp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
			log.Println("allowed!")
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/document/:id/comment", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/document/:id/edit", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/folder/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/folder/:id/comment", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/folder/:id/edit", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	emilia := &v1.SubjectReference{Object: &v1.ObjectReference{
		ObjectType: "ntrongtin11702_tutorial/user",
		ObjectId:   "emilia",
	}}

	firstPost := &v1.ObjectReference{
		ObjectType: "ntrongtin11702_tutorial/document",
		ObjectId:   "1",
	}

	if err != nil {
		return
	}

	resp, err := client.CheckPermission(context.Background(), &v1.CheckPermissionRequest{
		Resource:   firstPost,
		Permission: "view",
		Subject:    emilia,
	})
	if err != nil {
		log.Fatalf("failed to check permission: %s", err)
	}

	if resp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
		log.Println("allowed!")
	}
}

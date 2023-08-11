package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"example/authzed/initializers"
	"example/authzed/models"
	"example/authzed/services"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create account
	folder, err := services.CreateFolder(body.Email, 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create account",
		})
	}

	account := models.Account{Email: body.Email, Password: string(hash), MyDrive: *folder}
	result := initializers.DB.Create(&account)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create account",
		})
		return
	}

	// Write relationship
	folderObject := &v1.ObjectReference{
		ObjectType: os.Getenv("SPICE_DB_PREFIX") + "/folder",
		ObjectId:   strconv.FormatUint(uint64(folder.ID), 10),
	}

	userSubject := &v1.SubjectReference{
		Object: &v1.ObjectReference{
			ObjectType: os.Getenv("SPICE_DB_PREFIX") + "/user",
			ObjectId:   strconv.FormatUint(uint64(account.ID), 10),
		},
	}

	relationship := v1.Relationship{
		Resource: folderObject,
		Relation: "owner",
		Subject:  userSubject,
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
			"error": "Failed to set ownership: " + fmt.Sprint(error),
		})
		return
	}
	// NOTE: res.WrittenAt.Token -> Zedtoken
	// 	Consistency: You can use the zedtoken to read data from SpiceDB that is consistent with a specific write. By including the zedtoken in your read request, you ensure that you're reading data as it was at the time represented by the zedtoken.
	// Concurrency Control: If you're coordinating writes across multiple parts of a distributed system, you can use the zedtoken to implement optimistic concurrency control. You can include the zedtoken in a write request to specify that the write should only succeed if the current state of the database matches the state at the time represented by the zedtoken. If another write has occurred in the meantime, the write will fail, allowing you to handle the conflict.
	// Debugging and Auditing: Storing the zedtoken with your data allows you to correlate changes in your application's data with changes in SpiceDB. This can be useful for debugging issues or for auditing changes.
	fmt.Println(res, error)
	// https://authzed.com/docs/guides/writing-relationships#two-writes--commit
	// https://authzed.com/docs/reference/zedtokens-and-zookies

	// fmt.Println("%#v", account)
	c.JSON(http.StatusOK, gin.H{"account": fmt.Sprint(account)})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var account models.Account
	initializers.DB.First(&account, "email = ?", body.Email)
	if account.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email does not exist",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   account.ID,
		"email": account.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "I'm logged in"})
}

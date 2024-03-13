package auth

import (
	"net/http"
	"shool/internal/database"
	"shool/models"
	"shool/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ResetPassword(c *gin.Context) {

	var user models.User
	userDB := database.UserDB
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	user.Email = strings.ToLower(user.Email)

	filter := bson.D{{Key: "email", Value: user.Email}}
	res := userDB.Collection.FindOne(userDB.Ctx, filter)

	if res.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "The email not registered please signup"})
		return
	}

	user.Password, err = utils.HashingPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "password",
			Value: user.Password}}}}
	updateRes, err := userDB.Collection.UpdateOne(userDB.Ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	if updateRes.ModifiedCount != 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Didn't updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated Successfully"})

}

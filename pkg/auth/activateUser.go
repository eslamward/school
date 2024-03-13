package auth

import (
	"fmt"
	"net/http"
	"shool/internal/database"
	"shool/models"
	"shool/pkg/otp"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AcivateUser(c *gin.Context) {

	var user models.User
	var fetcheduser models.User
	userDB := database.UserDB
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error binding OTP"})
		return
	}

	user.Email = strings.ToLower(user.Email)
	fmt.Println(user.Email)
	filter := bson.D{{Key: "email", Value: user.Email}}
	res := userDB.Collection.FindOne(userDB.Ctx, filter)

	if res.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "The email not registered please signup"})
		return
	}
	res.Decode(&fetcheduser)
	otpStatus := otp.CheckOTP(&user.OTP, &fetcheduser.OTP)
	if !otpStatus {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Somthing error with otp you use"})
		return
	}

	update := bson.D{
		{Key: "$set",
			Value: bson.D{{Key: "status", Value: true}}}, {Key: "$unset", Value: bson.D{{Key: "otp", Value: ""}}}}
	updateRes, err := userDB.Collection.UpdateOne(userDB.Ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	if updateRes.ModifiedCount != 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Didn't updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activated Successfully"})

}

package otp

import (
	"net/http"
	"shool/internal/database"
	"shool/models"
	"shool/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ResendOTP(c *gin.Context) {
	var fetcheduser models.User
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
		c.JSON(http.StatusBadRequest, gin.H{"Error": "This email not registered yet"})
		return
	}
	res.Decode(&fetcheduser)
	if fetcheduser.OTP.Otp == "" {
		c.JSON(http.StatusForbidden, gin.H{"Errpr": "You Acivate your account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Check Your email"})
	otp := GenerateOTP()

	ot, _ := utils.HashingPassword(otp.Otp)

	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "otp.otp",
			Value: ot}}}}
	updateRes, err := userDB.Collection.UpdateOne(userDB.Ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	if updateRes.ModifiedCount != 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error Generate new OTP"})
		return
	}

	SendOTP(user.Email, otp.Otp)
}

package auth

import (
	"net/http"
	"shool/internal/database"
	"shool/models"
	"shool/pkg/otp"
	"shool/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(c *gin.Context) {

	var user models.User
	userDB := database.UserDB

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	user.ID = primitive.NewObjectID()
	user.FirstName = strings.ToLower(user.FirstName)
	user.LastName = strings.ToLower(user.LastName)
	user.Email = strings.ToLower(user.Email)
	user.Password, _ = utils.HashingPassword(user.Password)
	user.CreatedAt = time.Now()
	otpNotHased := (*otp.GenerateOTP()).Otp
	user.OTP.Otp, err = utils.HashingPassword(otpNotHased)
	user.Status = false
	if !validateUniqueEmail(userDB, user.Email) {
		c.JSON(http.StatusOK, gin.H{"Error": "email is used before"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	_, err = userDB.Collection.InsertOne(userDB.Ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	otp.SendOTP(user.Email, otpNotHased)
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func validateUniqueEmail(userDB database.UserDatabase, email string) bool {

	cur, _ := userDB.Collection.Find(userDB.Ctx, bson.D{{Key: "email", Value: email}})

	return !cur.Next(userDB.Ctx)

}

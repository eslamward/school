package auth

import (
	"net/http"
	"shool/internal/database"
	"shool/models"
	"shool/utils"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

func SignIn(c *gin.Context) {
	var user models.User
	userDB := database.UserDB

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	user.Email = strings.ToLower(user.Email)
	cur := userDB.Collection.FindOne(userDB.Ctx, bson.D{{Key: "email", Value: user.Email}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	if cur.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "email not found"})
		return
	}

	var userFetched models.User

	err = cur.Decode(&userFetched)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	if !userFetched.Status {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Please activate your mail before signin, check your email"})
		return
	}

	if !utils.ComparePassword(user.Password, userFetched.Password) {

		c.JSON(http.StatusUnauthorized, gin.H{"Error": "invalid password"})
		return

	}

	sessionToken := xid.New().String()

	session := sessions.Default(c)

	session.Set("username", user.Email)

	session.Set("token", sessionToken)

	session.Save()
	c.JSON(http.StatusOK, gin.H{"user": userFetched})

}

func SignForm(c *gin.Context) {

	c.HTML(http.StatusOK, "sign.html", nil)
}

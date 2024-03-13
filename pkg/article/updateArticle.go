package article

import (
	"net/http"
	"shool/internal/database"
	"shool/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateArticle(c *gin.Context) {
	var article models.Article
	articleDB := database.ArticleDB
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	err = c.ShouldBindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	_, err = articleDB.Collection.UpdateOne(articleDB.Ctx,
		bson.D{primitive.E{Key: "_id", Value: objID}},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "title", Value: article.Title},
				{Key: "text", Value: article.Text},
				{Key: "category", Value: article.Category},
				{Key: "updatedAt", Value: time.Now()},
			}}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	article.ID = objID
	c.JSON(http.StatusOK, gin.H{"Message": article})

}

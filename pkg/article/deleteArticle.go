package article

import (
	"net/http"
	"shool/internal/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteArticle(c *gin.Context) {

	var articleDB = database.ArticleDB
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	res, err := articleDB.Collection.DeleteOne(articleDB.Ctx, bson.D{{Key: "_id", Value: objID}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	if res.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "the article not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Article Deleted": "deletedt successfully"})

}

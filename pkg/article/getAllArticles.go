package article

import (
	"net/http"
	"shool/internal/database"
	"shool/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetArticles(c *gin.Context) {
	var articles []models.Article
	articleDB := database.ArticleDB

	cur, err := articleDB.Collection.Find(articleDB.Ctx, bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	for cur.Next(articleDB.Ctx) {
		var article models.Article
		err := cur.Decode(&article)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		articles = append(articles, article)
	}

	defer cur.Close(articleDB.Ctx)

	c.JSON(http.StatusOK, gin.H{"articles": articles})

}

package article

import (
	"fmt"
	"net/http"
	"shool/internal/database"
	"shool/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchArticle(c *gin.Context) {

	var articles []models.Article
	articleDB := database.ArticleDB

	cat := c.Query("cat")
	fmt.Println(cat)

	// filter := bson.D{primitive.E{
	// 	Key: "category", Value: primitive.E{
	// 		Key: "$elemMatch",
	// 		Value: bson.D{
	// 			primitive.E{
	// 				Key:   "$eq",
	// 				Value: cat,
	// 			},
	// 		},
	// 	},
	// }}

	filter1 := bson.D{{Key: "category", Value: bson.D{{Key: "$all", Value: []string{cat}}}}}

	cur, err := articleDB.Collection.Find(articleDB.Ctx, filter1)
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

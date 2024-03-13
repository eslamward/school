package article

import (
	"errors"
	"log"
	"net/http"
	"shool/internal/database"
	"shool/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateArticle(c *gin.Context) {

	var article models.Article
	var articleDB = database.ArticleDB

	err := c.ShouldBindJSON(&article)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	err = validate(article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	article.ID = primitive.NewObjectID()
	article.Category = lowerString(article.Category)
	article.PublishedAt = time.Now()

	_, err = articleDB.Collection.InsertOne(articleDB.Ctx, article)

	if err != nil {
		log.Fatal(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"article": article})
}

func validate(article models.Article) error {

	if article.Author == "" ||
		article.Title == "" ||
		article.Text == "" {

		return errors.New("one or more field is empty")
	}

	return nil
}

func lowerString(list []string) []string {

	var lowerList []string

	for _, item := range list {

		lowerList = append(lowerList, strings.ToLower(item))
	}
	return lowerList
}

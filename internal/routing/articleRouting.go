package routing

import (
	"shool/pkg/article"
	"shool/pkg/auth"

	"github.com/gin-gonic/gin"
)

func ArticleRouting(router *gin.Engine) {

	gr := router.Group("/")

	gr.Use(auth.AuthMiddleWare)

	router.GET("/", article.GetArticles)
	gr.GET("article/search", article.SearchArticle)
	gr.POST("article/", article.CreateArticle)
	gr.PUT("article/:id", article.UpdateArticle)
	gr.DELETE("article/:id", article.DeleteArticle)
}

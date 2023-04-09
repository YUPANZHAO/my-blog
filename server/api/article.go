package api

import (
	"my-blog-server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticleApi(c *gin.Context) {
	// get parameters
	name := c.Query("name")
	// search article by name
	article, _ := util.GetArticle(name)
	// return json data
	c.JSON(http.StatusOK, gin.H{
		"title":     article.Title,
		"date":      article.Date,
		"category":  article.Category,
		"cover_url": article.Cover_url,
		"content":   article.Content,
	})
}

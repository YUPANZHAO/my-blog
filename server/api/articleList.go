package api

import (
	"my-blog-server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticleListApi(c *gin.Context) {
	// get parameters
	category := c.Query("category")
	// get article list pointer
	var items []gin.H
	p_articles := util.GetArticles()
	// loop article list
	for i := 0; i < len(*p_articles); i++ {
		// exclude articles that don't meet the criteria
		if category != "" && (*p_articles)[i].Category != category {
			continue
		}
		items = append(items, gin.H{
			"name":         (*p_articles)[i].Name,
			"category":     (*p_articles)[i].Category,
			"date":         (*p_articles)[i].Date,
			"cover_url":    (*p_articles)[i].Cover_url,
			"title":        (*p_articles)[i].Title,
			"introduction": (*p_articles)[i].Introduce,
		})
	}
	// return articles
	c.JSON(http.StatusOK, gin.H{
		"articles": items,
	})
}

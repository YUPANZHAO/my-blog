package api

import (
	"my-blog-server/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RecentArticle(c *gin.Context) {
	// get parameters
	max_num, err := strconv.Atoi(c.Query("max_num"))
	if err != nil { 
		max_num = 4
	}
	// get article list pointer
	var items []gin.H
	p_articles := util.GetArticles()
	// loop article list
	for i := 0; i < len(*p_articles) && i < max_num; i++ {
		items = append(items, gin.H{
			"name":         (*p_articles)[i].Name,
			"title":        (*p_articles)[i].Title,
		})
	}
	// return articles
	c.JSON(http.StatusOK, gin.H{
		"recent_articles": items,
	})
}

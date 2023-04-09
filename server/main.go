package main

import (
	"my-blog-server/api"
	"my-blog-server/util"

	"github.com/gin-gonic/gin"
)

func main() {
	go util.StartPullThread()

	server := gin.Default()
	server.GET("/article/list", api.ArticleListApi)
	server.GET("/article/paper", api.ArticleApi)
	server.GET("/article/recent", api.RecentArticle)
	server.Run(":8081")
}

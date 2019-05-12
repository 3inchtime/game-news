package routers

import (
	"game-news/routers/news"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		api.GET("news", news.GetNews)
		api.POST("news", news.CreateNews)
	}

	return r
}

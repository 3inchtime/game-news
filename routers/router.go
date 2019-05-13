package routers

import (
	"game-news/routers/news"
	"io/ioutil"
	"net/http"
	"os"

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

	r.GET("/", func(c *gin.Context) {
		file, _ := os.Open("./static/html/index.html")
		html, _ := ioutil.ReadAll(file)
		c.Data(http.StatusOK, "html", html)
	})

	return r
}

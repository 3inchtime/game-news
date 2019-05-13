package news

import (
	"game-news/dbops"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNews(c *gin.Context) {
	limit := c.Query("limit")
	num, err := strconv.Atoi(limit)

	code := 200
	if err != nil {
		code = 403
		c.JSON(http.StatusForbidden, gin.H{
			"code": code,
			"msg":  "Get News Failed",
			"data": nil,
		})
	}

	newsData, err := dbops.GetAllNews(num)

	if err != nil {
		code = 403
		c.JSON(http.StatusForbidden, gin.H{
			"code": code,
			"msg":  "Get News Failed",
			"data": nil,
		})
	}

	code = 200
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": newsData,
	})
}

func CreateNews(c *gin.Context) {
	title := c.Query("title")
	url := c.Query("url")
	media := c.Query("media")
	article := c.Query("article")

	code := 200
	if dbops.CreateNews(title, url, article, media) {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  "Create News Successful!!!",
		})
	}
}

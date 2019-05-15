package news

import (
	"game-news/dbops"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllNews(c *gin.Context) {
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
	c.JSON(http.StatusOK, newsData)
}

func GetNewsByMedia(c *gin.Context) {
	limit := c.Query("limit")
	id := c.Query("id")

	num, err := strconv.Atoi(limit)
	media_id, err := strconv.Atoi(id)

	code := 200
	if err != nil {
		code = 403
		c.JSON(http.StatusForbidden, gin.H{
			"code": code,
			"msg":  "Get News Failed",
			"data": nil,
		})
	}
	newsData, err := dbops.GetNewsByMedia(num, media_id)
	if err != nil {
		code = 403
		c.JSON(http.StatusForbidden, gin.H{
			"code": code,
			"msg":  "Get News Failed",
			"data": nil,
		})
	}

	code = 200
	c.JSON(http.StatusOK, newsData)
}

func CreateNews(c *gin.Context) {
	title := c.PostForm("title")
	url := c.PostForm("url")
	media := c.PostForm("media")
	content := c.PostForm("content")
	pub_time := c.PostForm("pub_time")

	code := 200
	if dbops.CreateNews(title, url, content, media, pub_time) {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  "Create News Successful!!!",
		})
	}
}

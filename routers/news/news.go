package news

import (
	"fmt"
	"game-news/dbops"
	"game-news/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//func GetAllNews(c *gin.Context) {
//	limit := c.Query("limit")
//	num, err := strconv.Atoi(limit)
//
//	code := 200
//	if err != nil {
//		code = 403
//		c.JSON(http.StatusForbidden, gin.H{
//			"code": code,
//			"msg":  "Get News Failed",
//			"data": nil,
//		})
//	}
//
//	newsData, err := dbops.GetAllNews(num)
//
//	if err != nil {
//		code = 403
//		c.JSON(http.StatusForbidden, gin.H{
//			"code": code,
//			"msg":  "Get News Failed",
//			"data": nil,
//		})
//	}
//
//	code = 200
//	c.JSON(http.StatusOK, newsData)
//}

func GetNewsByMedia(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	pageLimit := c.DefaultQuery("page", "1")

	mediaId, err := strconv.Atoi(id)
	page, err := strconv.Atoi(pageLimit)

	begin, end := util.GetLimit(page)
	if err != nil {
		c.String(http.StatusForbidden, fmt.Sprintf("%s请求参数错误!", err))
		return
	}

	if mediaId == 0 {
		newsData, err := dbops.GetAllNews(begin, end)
		if err != nil {
			c.String(http.StatusForbidden, fmt.Sprintf("%s获取新闻错误!", err))
			return
		}
		c.JSON(http.StatusOK, newsData)
		return
	}

	newsData, err := dbops.GetNewsByMedia(begin, end, mediaId)
	if err != nil {
		c.String(http.StatusForbidden, fmt.Sprintf("%s获取新闻错误!", err))
		return
	}

	c.JSON(http.StatusOK, newsData)
	return
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

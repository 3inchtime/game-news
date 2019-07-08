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
	mediaID := c.PostForm("media_id")
	content := c.PostForm("content")
	pubTime := c.PostForm("pub_time")

	if dbops.CreateNews(title, url, content, mediaID, pubTime) {
		c.String(http.StatusOK, fmt.Sprintf("News Created"))
	} else {
		c.String(http.StatusForbidden, fmt.Sprintf("News Created Faild"))
	}

	if dbops.CreateNewsCache(title, url, content, mediaID, pubTime) {
		fmt.Println("Create News Cache Successful")
	} else {
		fmt.Println("Create News Cache Faild")
	}
}

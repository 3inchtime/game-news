package media

import (
	"game-news/dbops"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllMedia(c *gin.Context) {

	mediaData, err := dbops.GetAllMedia()

	code := 200
	if err != nil {
		code = 403
		c.JSON(http.StatusForbidden, gin.H{
			"code": code,
			"msg":  "Get News Failed",
			"data": nil,
		})
	}

	c.JSON(http.StatusOK, mediaData)
}

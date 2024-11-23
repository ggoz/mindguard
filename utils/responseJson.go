package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	SUCCESS int = 0 // 成功
	FAILED  int = 1 // 失败
)

func Success(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  "成功",
		"data": v,
	})
}

func Failed(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": FAILED,
		"msg":  v,
	})
}

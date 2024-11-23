package controller

import (
	"github.com/gin-gonic/gin"
	"mindguard/middleware"
	"mindguard/service"
	"mindguard/utils"
	"net/http"
)

type ArticleController struct {
}

func (ac *ArticleController) Router(engine *gin.Engine) {
	// 获取所有文章
	engine.GET("/api/getarticles", middleware.JWTAuthMiddleware(), ac.GetArticles)
}

// 用户退出登录(手机/账号)
func (ac *ArticleController) GetArticles(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}
	//调用业务层
	as := service.ArticleService{}
	articles := as.GetArticles()
	//fmt.Println("articles", articles)
	if articles != nil {
		c.JSON(http.StatusOK, gin.H{
			"articles": articles,
		})
		return
	}
	utils.Failed(c, "获取所有文章失败")
}

package router

import (
	"github.com/gin-gonic/gin"
	"mindguard/controller"
)

// 注册路由
func RegisterRouter(router *gin.Engine) {
	// 注册用户路由
	new(controller.UserController).Router(router)
	// 注册文章路由
	new(controller.ArticleController).Router(router)
	// 注册心理测试路由
	new(controller.TestController).Router(router)
	// 注册预约咨询路由
	new(controller.ReservationController).Router(router)
	// 注册用户评价路由
	new(controller.EvaluationController).Router(router)
	// 注册心理咨询路由
	new(controller.ConsultController).Router(router)
}

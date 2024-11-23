package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mindguard/middleware"
	"mindguard/params"
	"mindguard/service"
	"mindguard/utils"
	"net/http"
	"strconv"
)

type EvaluationController struct {
}

func (ec *EvaluationController) Router(engine *gin.Engine) {
	// 提交学生评价
	engine.POST("/api/postevaluation", middleware.JWTAuthMiddleware(), ec.PostEvaluation)
	// 获取老师的评价
	engine.GET("/api/getcomments", middleware.JWTAuthMiddleware(), ec.GetComments)
}

// 获取老师的评价
func (ec *EvaluationController) GetComments(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
		return
	}

	// 解析参数 teacher_id
	teacher_id, ok := c.GetQuery("teacher_id")
	if !ok {
		utils.Failed(c, "参数解析失败")
		return
	}
	id, err := strconv.ParseInt(teacher_id, 10, 64)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	// 调用业务层
	es := service.EvaluateService{}
	comments := es.GetComments(id)
	fmt.Println("comments", comments)
	if comments != nil {
		c.JSON(http.StatusOK, gin.H{
			"comments": comments,
		})
		return
	}
	utils.Failed(c, "该老师暂无评论")
}

// 提交学生评价
func (ec *EvaluationController) PostEvaluation(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
		return
	}

	// 解析参数 Student_Id Teacher_Id
	var userEvaluationParams params.UserEvaluationParams
	err := c.ShouldBindJSON(&userEvaluationParams)
	fmt.Println(userEvaluationParams.Evaluator)
	fmt.Println(userEvaluationParams.Evaluated)
	fmt.Println(userEvaluationParams.Comment)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	// 调用业务层
	es := service.EvaluateService{}
	insertOk := es.PostEvaluation(userEvaluationParams.Evaluator,
		userEvaluationParams.Evaluated, userEvaluationParams.Comment)
	fmt.Println("insertOk", insertOk)
	if insertOk {
		c.JSON(http.StatusOK, gin.H{
			"msg": "插入评论成功",
		})
		return
	}
	utils.Failed(c, "插入评论失败")
}

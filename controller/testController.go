package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mindguard/middleware"
	"mindguard/params"
	"mindguard/service"
	"mindguard/utils"
	"net/http"
)

type TestController struct {
}

func (tc *TestController) Router(engine *gin.Engine) {
	// 获取所有问题和与之对应的答案
	engine.GET("/api/getqsandas", middleware.JWTAuthMiddleware(), tc.GetQuestionsAndAnswers)
	//提交测试的接口
	engine.POST("/api/submittest", middleware.JWTAuthMiddleware(), tc.SubmitTest)

}

// 获取所有问题和与之对应的答案
func (tc *TestController) SubmitTest(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 choice
	var requestData params.SubmitTestRequest
	err := c.ShouldBindJSON(&requestData)
	fmt.Println(requestData)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	//调用业务层
	ts := service.TestService{}
	insertOk, score := ts.SubmitTest(requestData)
	fmt.Println("insertOk", insertOk)
	fmt.Println("score", score)
	if insertOk {
		if score >= 0 && score <= 15 {
			c.JSON(http.StatusOK, gin.H{
				"score": score,
				"msg":   "你可能在某些方面表现较低，建议关注个人发展和提升自己的能力。",
			})
			return
		} else if score >= 16 && score <= 25 {
			c.JSON(http.StatusOK, gin.H{
				"score": score,
				"msg":   "你在某些方面表现平衡，具备一定的适应性。",
			})
			return
		} else if score >= 25 && score <= 30 {
			c.JSON(http.StatusOK, gin.H{
				"score": score,
				"msg":   "你在某些方面可能表现较高，具备强烈的某些心理特征。",
			})
			return
		}

	}
	utils.Failed(c, "提交测试失败")
}

// 获取所有问题和与之对应的答案
func (tc *TestController) GetQuestionsAndAnswers(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	//调用业务层
	ts := service.TestService{}
	questionsAndAnswers := ts.GetQuestionsAndAnswers()
	fmt.Println("questionsAndAnswers", questionsAndAnswers)
	if questionsAndAnswers != nil {
		c.JSON(http.StatusOK, gin.H{
			"questionsAndAnswers": questionsAndAnswers,
		})
		return
	}
	utils.Failed(c, "获取所有问题答案失败")
}

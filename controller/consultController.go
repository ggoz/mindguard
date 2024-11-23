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

type ConsultController struct {
}

func (cc *ConsultController) Router(engine *gin.Engine) {
	// 获取所有在线教师接口
	engine.GET("/api/getallteacher", middleware.JWTAuthMiddleware(), cc.GetAllTeachers)
	// 发送消息接口
	engine.POST("/api/sendmsg", middleware.JWTAuthMiddleware(), cc.SendMsg)
	// 根据两个发送者获取对应的消息接口
	engine.POST("/api/getmsg", middleware.JWTAuthMiddleware(), cc.GetMsg)
	// 根据两个发送者获取对应的消息接口
	engine.POST("/api/getacceptor", middleware.JWTAuthMiddleware(), cc.GetAcceptor)
	// 获取所有聊天的学生
	engine.GET("/api/getchatstudents", middleware.JWTAuthMiddleware(), cc.GetChatStudents)
}

// 获取所有聊天的学生
func (cc *ConsultController) GetChatStudents(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 acceptor
	acceptor, ok := c.GetQuery("acceptor")
	fmt.Println(acceptor)
	if !ok {
		utils.Failed(c, "参数解析失败")
		return
	}
	// 将字符串转换为 int64
	acc, err := strconv.ParseInt(acceptor, 10, 64)

	if err != nil {
		// 处理转换错误
		fmt.Println("转换失败:", err)
		return
	}
	//调用业务层
	cs := service.ConsultService{}
	students := cs.GetChatStudents(acc)
	fmt.Println("students", students)
	if students != nil {
		c.JSON(http.StatusOK, gin.H{
			"students": students,
		})
		return
	}
	utils.Failed(c, "获取所有有关聊天的学生失败")
}

// 获取接受消息的人
func (cc *ConsultController) GetAcceptor(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 appointmentId
	var getAcceptorParams params.GetAcceptorParams
	err := c.ShouldBindJSON(&getAcceptorParams)
	fmt.Println(getAcceptorParams.Acceptor)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	acc, err := strconv.Atoi(getAcceptorParams.Acceptor)

	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	//调用业务层
	cs := service.ConsultService{}
	acceptor := cs.GetAcceptor(acc)
	fmt.Println("acceptor", acceptor)
	if acceptor != nil {
		c.JSON(http.StatusOK, gin.H{
			"acceptor": acceptor,
		})
		return
	}
	utils.Failed(c, "获取所有在线教师失败")
}

// 根据两个发送者获取对应的消息
func (cc *ConsultController) GetMsg(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 appointmentId
	var getMsgParams params.GetMsgParams
	err := c.ShouldBindJSON(&getMsgParams)
	fmt.Println(getMsgParams.Sender)
	fmt.Println(getMsgParams.Acceptor)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}
	// 将字符串转为 int64
	sender, err := strconv.ParseInt(getMsgParams.Sender, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 将字符串转为 int64
	acceptor, err := strconv.ParseInt(getMsgParams.Acceptor, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//调用业务层
	cs := service.ConsultService{}
	communications := cs.GetMsg(sender, acceptor)
	fmt.Println("communications", communications)
	if communications != nil {
		c.JSON(http.StatusOK, gin.H{
			"communications": communications,
		})
		return
	}
	utils.Failed(c, "获取消息失败")
}

// 发送消息
func (cc *ConsultController) SendMsg(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 appointmentId
	var sendMsgParams params.SendMsgParams
	err := c.ShouldBindJSON(&sendMsgParams)
	fmt.Println(sendMsgParams.Sender)
	fmt.Println(sendMsgParams.Acceptor)
	fmt.Println(sendMsgParams.Message)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	// 将字符串转为 int64
	sender, err := strconv.ParseInt(sendMsgParams.Sender, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 将字符串转为 int64
	acceptor, err := strconv.ParseInt(sendMsgParams.Acceptor, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//调用业务层
	cs := service.ConsultService{}
	insertOk := cs.SendMsg(sender, acceptor, sendMsgParams.Message)
	fmt.Println("insertOk", insertOk)
	if insertOk {
		c.JSON(http.StatusOK, gin.H{
			"msg": "发送消息成功",
		})
		return
	}
	utils.Failed(c, "发送消息失败")
}

// 获取所有在线教师
func (cc *ConsultController) GetAllTeachers(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	//调用业务层
	cs := service.ConsultService{}
	teachers := cs.GetAllTeachers()
	fmt.Println("teachers", teachers)
	if teachers != nil {
		c.JSON(http.StatusOK, gin.H{
			"teachers": teachers,
		})
		return
	}
	utils.Failed(c, "获取所有在线教师失败")
}

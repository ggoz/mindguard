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

type ReservationController struct {
}

func (rc *ReservationController) Router(engine *gin.Engine) {
	// 获取所有在线教师接口
	engine.GET("/api/getallteachers", middleware.JWTAuthMiddleware(), rc.GetAllTeachers)
	// 发起预约接口
	engine.POST("/api/order", middleware.JWTAuthMiddleware(), rc.PostOrder)
	// 取消预约接口
	engine.POST("/api/cancelorder", middleware.JWTAuthMiddleware(), rc.CancelOrder)
	// 获取所有预约接口
	engine.GET("/api/getallreservations", middleware.JWTAuthMiddleware(), rc.GetAllReservations)
	// 获取所有与老师有关的预约
	engine.GET("/api/getreservationsteacher", middleware.JWTAuthMiddleware(), rc.GetTeacherReservations)
	// 老师同意预约
	engine.POST("/api/acceptorder", middleware.JWTAuthMiddleware(), rc.AcceptOrder)
	// 获取学生有关的预约
	engine.GET("/api/getstudentreservations", middleware.JWTAuthMiddleware(), rc.GetStudentReservations)
	// 获取相应的预约
	engine.GET("/api/getres", middleware.JWTAuthMiddleware(), rc.GetRes)

}

// 获取相应的预约
func (rc *ReservationController) GetRes(c *gin.Context) {
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
	tid, err := strconv.ParseInt(teacher_id, 10, 64)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	// 解析参数 student_id
	student_id, ok := c.GetQuery("student_id")
	if !ok {
		utils.Failed(c, "参数解析失败")
		return
	}
	sid, err := strconv.ParseInt(student_id, 10, 64)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	// 调用业务层
	rs := service.ReservationService{}
	reservation := rs.GetRes(tid, sid)
	fmt.Println("reservation", reservation)
	if reservation != nil {
		c.JSON(http.StatusOK, gin.H{
			"reservation": reservation,
		})
		return
	}
	utils.Failed(c, "暂时没有预约")
}

// 获取学生有关的预约
func (rc *ReservationController) GetStudentReservations(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
		return
	}

	// 解析参数 teacher_id
	student_id, ok := c.GetQuery("student_id")
	if !ok {
		utils.Failed(c, "参数解析失败")
		return
	}
	id, err := strconv.ParseInt(student_id, 10, 64)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	// 调用业务层
	rs := service.ReservationService{}
	teachers := rs.GetStudentReservations(id)
	fmt.Println("teachers", teachers)
	if teachers != nil {
		c.JSON(http.StatusOK, gin.H{
			"teachers": teachers,
		})
		return
	}
	utils.Failed(c, "暂时没有预约")
}

// 老师同意预约
func (rc *ReservationController) AcceptOrder(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 Student_Id Teacher_Id
	var acceptOrderParams params.AcceptOrderParams
	err := c.ShouldBindJSON(&acceptOrderParams)
	fmt.Println(acceptOrderParams.Student_Id)
	fmt.Println(acceptOrderParams.Teacher_Id)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	// 调用业务层
	rs := service.ReservationService{}
	modifyOk := rs.AcceptOrder(acceptOrderParams.Student_Id, acceptOrderParams.Teacher_Id)
	fmt.Println("modifyOk", modifyOk)
	if modifyOk {
		c.JSON(http.StatusOK, gin.H{
			"code": 210,
			"msg":  "接受预约成功",
		})
		return
	}
	utils.Failed(c, "该老师接受预约失败")
}

// 获取所有与老师有关的预约
func (rc *ReservationController) GetTeacherReservations(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
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
	rs := service.ReservationService{}
	students := rs.GetTeacherReservations(id)
	fmt.Println("students", students)
	if students != nil {
		c.JSON(http.StatusOK, gin.H{
			"students": students,
		})
		return
	}
	utils.Failed(c, "该老师没有预约")
}

// 获取所有预约
func (rc *ReservationController) GetAllReservations(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 调用业务层
	rs := service.ReservationService{}
	reservations := rs.GetAllReservations()
	fmt.Println("reservations", reservations)
	if reservations != nil {
		c.JSON(http.StatusOK, gin.H{
			"reservations": reservations,
		})
		return
	}
	utils.Failed(c, "获取所有预约失败")
}

// 取消预约
func (rc *ReservationController) CancelOrder(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 appointmentId
	var cancelOrderParams params.CancelOrderParams
	err := c.ShouldBindJSON(&cancelOrderParams)
	fmt.Println(cancelOrderParams.AppointmentId)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	// 调用业务层
	rs := service.ReservationService{}
	deleteOk := rs.CancelOrder(cancelOrderParams.AppointmentId)
	fmt.Println("deleteOk", deleteOk)
	if deleteOk {
		c.JSON(http.StatusOK, gin.H{
			"msg": "取消预约成功",
		})
		return
	}
	utils.Failed(c, "取消预约失败")
}

// 发起预约
func (rc *ReservationController) PostOrder(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 orderId orderedId
	var orderParams params.OrderParams
	err := c.ShouldBindJSON(&orderParams)
	fmt.Println(orderParams.OrderId)
	fmt.Println(orderParams.OrderedId)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	// 调用业务层
	rs := service.ReservationService{}
	reservation := rs.PostOrder(orderParams.OrderId, orderParams.OrderedId)
	fmt.Println("reservation", reservation)
	if reservation != nil {
		c.JSON(http.StatusOK, gin.H{
			"reservation": reservation,
		})
		return
	}
	utils.Failed(c, "添加预约失败")
}

// 获取所有教师
func (rc *ReservationController) GetAllTeachers(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 调用业务层
	rs := service.ReservationService{}
	teachers := rs.GetAllTeachers()
	fmt.Println("teachers", teachers)
	if teachers != nil {
		c.JSON(http.StatusOK, gin.H{
			"teachers": teachers,
		})
		return
	}
	utils.Failed(c, "查询所有教师失败")
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mindguard/middleware"
	"mindguard/model"
	"mindguard/params"
	"mindguard/service"
	"mindguard/utils"
	"net/http"
)

type UserController struct {
}

func (u *UserController) Router(engine *gin.Engine) {
	// 发送验证码接口
	engine.GET("/api/sendcode", u.SendSmsCode)
	// 短信登录接口
	engine.POST("/api/loginbycode", u.LoginByCode)
	// 账号密码登录接口
	engine.POST("/api/loginbypwd", u.LoginByPwd)
	// 用户退出登录接口
	engine.POST("/api/logout", middleware.JWTAuthMiddleware(), u.Logout)
	// 修改密码接口
	engine.POST("/api/modifypwd", middleware.JWTAuthMiddleware(), u.ModifyPwd)
	// 修改用户名接口
	engine.POST("/api/modifyusername", middleware.JWTAuthMiddleware(), u.ModifyUsername)
	// 修改手机号接口
	engine.POST("/api/modifyphone", middleware.JWTAuthMiddleware(), u.ModifyPhone)

}

// 修改 绑定手机号
func (u *UserController) ModifyPhone(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 oldUsername newUsername
	var modifyPhoneParams params.ModifyPhoneParams
	err := c.ShouldBindJSON(&modifyPhoneParams)
	fmt.Println(modifyPhoneParams.Id)
	fmt.Println(modifyPhoneParams.OldPhone)
	fmt.Println(modifyPhoneParams.NewPhone)
	fmt.Println(modifyPhoneParams.VerifyCode)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	// 调用业务层
	ms := service.UserService{}
	ok := ms.ModifyPhone(modifyPhoneParams.Id, modifyPhoneParams.OldPhone, modifyPhoneParams.NewPhone, modifyPhoneParams.VerifyCode)
	fmt.Println("ok", ok)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 210,
			"msg":  "修改手机号成功",
		})
		return
	}
	utils.Failed(c, "修改手机号失败")
}

// 修改用户名
func (u *UserController) ModifyUsername(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 oldUsername newUsername
	var modifyUsernameParams params.ModifyUsernameParams
	err := c.ShouldBindJSON(&modifyUsernameParams)
	fmt.Println(modifyUsernameParams.Id)
	fmt.Println(modifyUsernameParams.OldUsername)
	fmt.Println(modifyUsernameParams.NewUsername)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	// 调用业务层
	ms := service.UserService{}
	ok := ms.ModifyUsername(modifyUsernameParams.Id, modifyUsernameParams.OldUsername, modifyUsernameParams.NewUsername)
	fmt.Println("ok", ok)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 210,
			"msg":  "修改用户名成功",
		})
		return
	}
	utils.Failed(c, "修改用户名失败")
}

// 修改密码
func (u *UserController) ModifyPwd(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 newPwd oldPwd
	var modifyPwdParams params.ModifyPwdParams
	err := c.ShouldBindJSON(&modifyPwdParams)
	fmt.Println(modifyPwdParams.Id)
	fmt.Println(modifyPwdParams.OldPwd)
	fmt.Println(modifyPwdParams.NewPwd)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	// 调用业务层
	ms := service.UserService{}
	ok := ms.ModifyPwd(modifyPwdParams.Id, modifyPwdParams.OldPwd, modifyPwdParams.NewPwd)
	fmt.Println("ok", ok)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 210,
			"msg":  "修改密码成功",
		})
		return
	}
	utils.Failed(c, "修改密码失败")
}

// 用户退出登录(手机/账号)
func (u *UserController) Logout(c *gin.Context) {
	// 根据请求上下文获取信息
	_, exists := c.Get("username")
	if !exists {
		utils.Failed(c, "token错误!")
	}

	// 解析参数 id
	var logoutParams params.LogoutParams
	err := c.ShouldBindJSON(&logoutParams)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}
	fmt.Println(logoutParams.Id)

	// 调用业务层
	ms := service.UserService{}
	ok := ms.Logout(logoutParams.Id)
	fmt.Println("ok", ok)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 220,
			"msg":  "退出成功",
		})
		return
	}
	utils.Failed(c, "退出失败")
}

// 账号密码登录
func (u *UserController) LoginByPwd(c *gin.Context) {
	// 解析参数 phone code
	var pwdLoginParams params.PwdLoginParams
	err := c.ShouldBindJSON(&pwdLoginParams)
	fmt.Println(pwdLoginParams.Username)
	fmt.Println(pwdLoginParams.Password)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	// 调用业务层
	ms := service.UserService{}
	isLogin, token, user := ms.LoginByPwd(pwdLoginParams.Username, pwdLoginParams.Password)
	fmt.Println("isLogin", isLogin)
	fmt.Println("user", user)
	userRes := model.User{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Status:   user.Status,
		Avator:   user.Avator,
	}
	fmt.Println("userRes: ", userRes)
	pwdExists := false
	if user.Password != "" {
		pwdExists = true
	}
	if isLogin {
		utils.Success(c, gin.H{
			"token":     token,
			"user":      userRes,
			"pwdExists": pwdExists,
		})
		return
	}
	utils.Failed(c, "登录失败")
}

// 手机短信登录/注册 手机号+验证码
func (u *UserController) LoginByCode(c *gin.Context) {
	// 解析参数 phone code
	var phoneLoginParams params.PhoneLoginParams
	err := c.ShouldBindJSON(&phoneLoginParams)
	fmt.Println(phoneLoginParams.Phone)
	fmt.Println(phoneLoginParams.Code)
	if err != nil {
		fmt.Println(err)
		utils.Failed(c, "参数解析失败")
		return
	}

	// 调用业务层
	ms := service.UserService{}
	isLogin, token, user := ms.LoginByCode(phoneLoginParams.Phone, phoneLoginParams.Code)
	fmt.Println("isLogin", isLogin)
	userRes := model.User{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Status:   user.Status,
		Avator:   user.Avator,
	}
	pwdExists := false
	if user.Password != "" {
		pwdExists = true
	}
	fmt.Println("userRes: ", userRes)
	if isLogin {
		utils.Success(c, gin.H{
			"token":     token,
			"user":      userRes,
			"pwdExists": pwdExists,
		})
		return
	}
	utils.Failed(c, "登录失败")
}

// http://localhost:8090/api/sendcode?phone=13167582436
// 获取短信验证码
func (u *UserController) SendSmsCode(c *gin.Context) {
	// 解析参数 phone
	phone, ok := c.GetQuery("phone")
	if !ok {
		utils.Failed(c, "参数解析失败")
		return
	}

	// 调用业务层
	ms := service.UserService{}
	isSend := ms.SendCode(phone)
	fmt.Println("issend", isSend)
	if isSend {
		utils.Success(c, "发送成功")
		return
	}
	utils.Failed(c, "发送失败")
}

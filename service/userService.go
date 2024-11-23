package service

import (
	"fmt"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/goccy/go-json"
	"math/rand"
	"mindguard/dao"
	"mindguard/model"
	"mindguard/utils"
	"time"
)

type UserService struct {
}

// 修改手机号业务逻辑
func (us *UserService) ModifyPhone(id int, oldPhone, newPhone, verifyCode string) bool {
	// 从redis中查询验证码是否正确
	redisCode := utils.RediStore.Get(newPhone)
	if verifyCode != redisCode || redisCode == "" {
		fmt.Println("验证码错误")
		return false
	}

	// 从数据库中查询手机号是否存在
	user := dao.QueryUserByPhone(oldPhone)
	fmt.Println(user)
	modifyOk := false
	if user.Phone == "" && user.Id != 0 {
		// 修改手机号
		modifyOk = dao.ModifyPhoneById(id, newPhone)
		return modifyOk
	}
	if user.Phone == oldPhone {
		fmt.Println("绑定失败，手机号已存在")
		return false
	}

	// 修改手机号
	modifyOk = dao.ModifyPhoneById(id, newPhone)
	return modifyOk
}

// 修改用户名业务逻辑
func (us *UserService) ModifyUsername(id int, oldUsername string, newUsername string) bool {
	// 从数据库中查询用户名是否存在
	user := dao.QueryUserByUsername(oldUsername)
	fmt.Println(user)

	// 数据库用户名和前端用户名都为空
	if user.Username == "" && oldUsername == "" {
		// 修改用户名
		dao.ModifyUsernameById(id, newUsername)
		return true
	}

	// 用户名已存在
	if newUsername == user.Username {
		fmt.Println("用户名重复")
		return false
	}

	// 修改用户名
	modifyOk := dao.ModifyUsernameById(id, newUsername)
	return modifyOk
}

// 修改密码业务逻辑
func (us *UserService) ModifyPwd(id int, oldPwd string, newPwd string) bool {
	// 从数据库中查询密码是否正确
	queryOk := dao.QueryPwd(id, oldPwd)
	if !queryOk {
		// 密码不正确
		fmt.Println("密码查询错误")
		return false
	}
	// 数据库修改密码
	modifyOk := dao.ModifyPwdById(id, newPwd)
	if !modifyOk {
		fmt.Println("密码修改失败")
		return false
	}
	return true
}

// 退出登录业务逻辑
func (us *UserService) Logout(id int64) bool {
	// 从redis中删除token
	err := utils.RediStore.Del("token")
	if err != nil {
		fmt.Println(err)
		return false
	}

	// 修改用户在线状态
	modifyOk := dao.ModifyOnlineById(id, "0")
	return modifyOk
}

// 账号密码登录业务逻辑
func (us *UserService) LoginByPwd(username string, password string) (bool, string, *model.User) {
	if username == "" {
		return false, "", nil
	}
	// 从数据库中查询是否有账号 没有就进行注册
	user := dao.QueryUserByUsername(username)
	fmt.Println("user: ", user)
	flag := false
	u := model.User{}
	if user.Username == "" || user == nil {
		// 说明没有注册过 注册进行登录
		//密码进行加密
		flag = true
		hashPassword, err := utils.HashPassword(password)
		fmt.Println("注册hash密码", hashPassword)
		if err != nil {
			fmt.Println("密码加密错误。")
			return false, "", nil
		}

		u = model.User{
			Username: username,
			Password: hashPassword,
			Status:   "学生",
			Online:   "1",
			Avator:   "https://s11.ax1x.com/2024/01/06/pizA38I.jpg",
		}
		fmt.Println("u:", u)
		err = dao.CreateUser(&u)
		if err != nil {
			fmt.Println(err)
			return false, "", nil
		}
	} else {
		// 说明有账号 校验密码
		hashPassword, err := utils.HashPassword(password)
		if err != nil {
			fmt.Println("密码加密错误。")
			return false, "", nil
		}
		if !utils.CheckPassword(password, hashPassword) {
			fmt.Println("密码错误！")
			return false, "", nil
		}

		// 修改用户在线状态
		modifyOk := dao.ModifyOnlineById(user.Id, "1")
		if !modifyOk {
			return false, "", nil
		}
	}

	if flag {
		tokenString, _ := utils.GenToken(u.Username)
		fmt.Println(tokenString)
		// 将token存入redis
		err := utils.RediStore.Set("token", tokenString)
		if err != nil {
			fmt.Println(err)
			return false, "", nil
		}
		return true, tokenString, &u
	} else {
		tokenString, _ := utils.GenToken(user.Username)
		// 将token存入redis
		err := utils.RediStore.Set("token", tokenString)
		if err != nil {
			fmt.Println(err)
			return false, "", nil
		}
		return true, tokenString, user
	}

}

// 手机验证码登录业务逻辑
func (us *UserService) LoginByCode(phone string, code string) (bool, string, *model.User) {
	if phone == "" {
		return false, "", nil
	}
	// 从redis中查询验证码是否正确
	redisCode := utils.RediStore.Get(phone)
	if code != redisCode {
		return false, "", nil
	}

	// 从数据库中查询是否有手机号 没有就进行注册
	user := dao.QueryUserByPhone(phone)
	fmt.Println("user: ", user)
	u := model.User{}
	flag := false
	if user.Phone == "" {
		// 说明没有注册过
		// 注册进行登录
		flag = true
		fmt.Println("2")
		// 随机生成八位数字
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(100000000)
		// 构建最终字符串
		resultString := fmt.Sprintf("未命名_%08d", randomNumber)
		u = model.User{
			Username: resultString,
			Phone:    phone,
			Status:   "学生",
			Online:   "1",
			Avator:   "https://s11.ax1x.com/2024/01/06/pizAGxP.jpg",
		}
		fmt.Println("u:", u)
		err := dao.CreateUser(&u)
		if err != nil {
			fmt.Println(err)
			return false, "", nil
		}
	} else {
		// 修改用户在线状态
		modifyOk := dao.ModifyOnlineById(user.Id, "1")
		if !modifyOk {
			return false, "", nil
		}
	}

	if flag {
		tokenString, _ := utils.GenToken(u.Phone)
		fmt.Println(tokenString)
		// 将token存入redis
		err := utils.RediStore.Set("token", tokenString)
		if err != nil {
			fmt.Println(err)
			return false, "", nil
		}
		return true, tokenString, &u
	} else {
		tokenString, _ := utils.GenToken(user.Phone)
		fmt.Println(tokenString)
		// 将token存入redis
		err := utils.RediStore.Set("token", tokenString)
		if err != nil {
			fmt.Println(err)
			return false, "", nil
		}
		return true, tokenString, user
	}
}

// 获取验证码业务逻辑
func (us *UserService) SendCode(phone string) bool {
	// 1.随机生成一个6位的验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	cfg := utils.GetConfig().Sms
	//2. 调用sdk
	client, err := dysmsapi.NewClientWithAccessKey(cfg.RegionId, cfg.AppKey, cfg.AppSecret)
	if err != nil {
		fmt.Println(err)
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = cfg.SignName
	request.TemplateCode = cfg.TemplateCode

	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(response)
	fmt.Println("code: " + code)

	// 判断是否成功得到短信结果
	if response.Code == "OK" {
		// 得到短信验证码 存入redis
		err := utils.RediStore.Set(phone, code)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}

	return false
}

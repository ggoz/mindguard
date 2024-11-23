package dao

import (
	"fmt"
	"mindguard/model"
	"mindguard/utils"
)

// 根据id 修改用户在线状态
func ModifyOnlineById(id int64, online string) bool {
	tx := DB.Model(&model.User{}).Where("id = ?", id).Update("online", online)
	fmt.Println(tx)
	if tx.RowsAffected != 1 {
		return false
	}
	return true
}

// 根据id 手机号修改用户名
func ModifyPhoneById(id int, newPhone string) bool {
	tx := DB.Model(&model.User{}).Where("id = ?", id).Update("phone", newPhone)
	fmt.Println(tx)
	if tx.RowsAffected != 1 {
		return false
	}
	return true
}

// 根据id修改用户名
func ModifyUsernameById(id int, newUsername string) bool {
	tx := DB.Model(&model.User{}).Where("id = ?", id).Update("username", newUsername)
	fmt.Println(tx)
	if tx.RowsAffected != 1 {
		return false
	}
	return true
}

// 根据id修改密码
func ModifyPwdById(id int, newPwd string) bool {
	// 将密码加密后存入数据库
	password, err := utils.HashPassword(newPwd)
	if err != nil {
		fmt.Println("存入数据库失败")
		return false
	}

	tx := DB.Model(&model.User{}).Where("id = ?", id).Update("password", password)
	fmt.Println(tx)
	if tx.RowsAffected != 1 {
		return false
	}
	return true
}

// 查询密码是否正确
func QueryPwd(id int, oldPwd string) bool {
	// 根据id查找用户
	user := QueryUserById(id)
	fmt.Println(user.Password)

	// 判断密码是否正确
	// 密码为空直接比较
	if oldPwd == "" {
		if user.Password == oldPwd {
			return true
		}
	}
	// 密码不为空
	fmt.Println("old", oldPwd)
	fmt.Println("比较密码", utils.CheckPassword(oldPwd, user.Password))
	if utils.CheckPassword(oldPwd, user.Password) {
		return true
	}
	return false
}

// 根据id查询用户
func QueryUserById(id int) *model.User {
	var user model.User
	err := DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &user
}

// 创建一个用户保存到数据库
func CreateUser(user *model.User) (err error) {
	err = DB.Create(user).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return nil
}

// 根据username查询 user
func QueryUserByUsername(username string) *model.User {
	var user model.User
	err := DB.Where("username = ?", username).Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &user
}

// 根据phone查询 user
func QueryUserByPhone(phone string) *model.User {
	var user model.User
	err := DB.Where("phone = ?", phone).Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &user
}

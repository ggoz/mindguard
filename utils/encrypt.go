package utils

import "golang.org/x/crypto/bcrypt"

// bcrypt密码加密
func HashPassword(password string) (string, error) {
	// 生成哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// 将哈希值转换为字符串
	hashedPasswordString := string(hashedPassword)
	return hashedPasswordString, nil
}

func CheckPassword(inputPassword, hashedPassword string) bool {
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}

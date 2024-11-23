package dao

import (
	"fmt"
	"mindguard/model"
)

// 查询 acceptor有关对话
func GetCommunicationsByAcceptor(acceptor int64) []model.Communication {
	var communications []model.Communication
	err := DB.Where("acceptor = ?", acceptor).Find(&communications).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return communications
}

// 获取消息
func GetMsg(sender, acceptor int64) []model.Communication {
	var communications []model.Communication
	err := DB.Where("(sender = ? AND acceptor = ?) OR (sender = ? AND acceptor = ?)",
		sender, acceptor, acceptor, sender).Find(&communications).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return communications
}

// 插入消息
func InsertMsg(communication *model.Communication) bool {
	err := DB.Create(communication).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

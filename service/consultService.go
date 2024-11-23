package service

import (
	"mindguard/dao"
	"mindguard/model"
	"time"
)

type ConsultService struct {
}

// 获取所有聊天的学生
func (cs *ConsultService) GetChatStudents(acceptor int64) []model.User {
	// 从数据库 查询 acceptor有关对话
	communications := dao.GetCommunicationsByAcceptor(acceptor)

	var mySlice []int64
	uniqueAcceptors := make(map[int64]bool)
	for _, communication := range communications {
		sender := communication.Sender
		// 检查当前 Acceptor 是否已经存在于 uniqueAcceptors 中
		if _, exists := uniqueAcceptors[sender]; !exists {
			// 如果不存在，将其加入切片
			mySlice = append(mySlice, sender)
			// 标记该 Acceptor 已经存在于 uniqueAcceptors 中
			uniqueAcceptors[sender] = true
		}
	}

	var users []model.User
	for _, value := range mySlice {
		user := dao.QueryUserById(int(value))
		users = append(users, *user)
	}
	return users
}

// 获取接受消息的人
func (cs *ConsultService) GetAcceptor(acceptor int) *model.User {
	// 从数据库 获取消息
	return dao.QueryUserById(acceptor)
}

// 根据两个发送者获取对应的消息
func (cs *ConsultService) GetMsg(sender, acceptor int64) []model.Communication {
	// 从数据库 获取消息
	return dao.GetMsg(sender, acceptor)
}

// 发送消息
func (cs *ConsultService) SendMsg(sender, acceptor int64, message string) bool {
	// 从数据库 插入消息
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	communication := model.Communication{
		Sender:   sender,
		Acceptor: acceptor,
		Message:  message,
		Time:     currentTime,
	}
	return dao.InsertMsg(&communication)
}

// 获取所有教师
func (cs *ConsultService) GetAllTeachers() []model.User {
	// 从数据库 查询所有在线教师
	return dao.GetAllTeachers()
}

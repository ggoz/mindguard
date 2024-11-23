package model

// 聊天表
type Communication struct {
	Id       int64  `json:"id"`
	Sender   int64  `gorm:"varchar(255)" json:"sender"`
	Acceptor int64  `gorm:"varchar(255)" json:"acceptor"`
	Message  string `gorm:"type:text" json:"message"`
	Time     string `gorm:"varchar(255)" json:"time"`

	// 添加外键关联
	Sender1   User `gorm:"foreignKey:Sender"`
	Acceptor1 User `gorm:"foreignKey:Acceptor"`
}

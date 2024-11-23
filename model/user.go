package model

// 用户表
type User struct {
	Id       int64  `json:"id"`
	Username string `gorm:"varchar(255)" json:"username"`
	Phone    string `gorm:"varchar(11)" json:"phone"`
	Password string `gorm:"varchar(255)" json:"password"`
	Status   string `gorm:"varchar(10)" json:"status"`
	Avator   string `gorm:"varchar(255)" json:"avator"`
	// 是否在线
	Online         string          `gorm:"char(1)" json:"online"`
	UserRecords    []UserRecord    // 添加 UserRecord 切片，用于表示与 UserRecord 模型的关联
	Reservations   []Reservation   `gorm:"foreignKey:OrderId;references:Id"` // 预约的外键关联
	Evaluations    []Evaluation    `gorm:"foreignKey:Evaluator;references:Id"`
	Communications []Communication `gorm:"foreignKey:Sender;references:Id"`
}

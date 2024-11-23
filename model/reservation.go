package model

// 存储预约咨询
type Reservation struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	// 预约人
	OrderId int64 `gorm:"not null" json:"order_id"`
	// 被预约的人
	OrderedId       int64  `gorm:"not null" json:"ordered_id"`
	AppointmentTime string `gorm:"varchar(255)" json:"appointment_time"`
	Status          string `gorm:"varchar(10)" json:"status"`

	// 添加外键关联
	Order   User `gorm:"foreignKey:OrderId"`
	Ordered User `gorm:"foreignKey:OrderedId"`
}

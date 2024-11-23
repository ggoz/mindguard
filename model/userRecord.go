package model

// 存储用户的测试记录
type UserRecord struct {
	Id         int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId     int64  `gorm:"not null;foreignKey:UserId;references:Id" json:"user_id"`
	TestDate   string `gorm:"varchar(255);not null" json:"test_date"`
	TotalScore int64  `gorm:"not null" json:"total_score"`
	//UserAnswers []UserAnswer `gorm:"foreignKey:RecordId" json:"user_answers"`
}

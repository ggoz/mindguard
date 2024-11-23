package model

// 存储心理测试的问题
type Question struct {
	Id           int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	QuestionText string `gorm:"varchar(255)" json:"question_text"`
	// 添加 Answers 切片，用于表示与 Answer 模型的关联
	Answers []Answer `json:"answers"`
	// 添加 UserRecord 切片，用于表示与 UserRecord 模型的关联
	//UserAnswers []UserAnswer `json:"userAnswers"`
}

package model

// 存储问题的可能答案
type Answer struct {
	Id         int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	QuestionID int64  `gorm:"not null;foreignKey:QuestionID;references:Id" json:"question_id"`
	AnswerText string `gorm:"varchar(255);not null" json:"answer_text"`
	Score      int64  `gorm:"not null" json:"score"`
	//UserAnswers []UserAnswer // 添加 UserRecord 切片，用于表示与 UserRecord 模型的关联
}

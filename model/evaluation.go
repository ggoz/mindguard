package model

type Evaluation struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Evaluator int64  `json:"evaluator"`
	Evaluated int64  `json:"evaluated"`
	Comment   string `gorm:"type:text" json:"comment"`

	// 添加外键关联
	Evaluator1 User `gorm:"foreignKey:Evaluator"`
	Evaluated1 User `gorm:"foreignKey:Evaluated"`
}

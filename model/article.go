package model

// 文章表
type Article struct {
	Id          int64  `json:"id"`
	Title       string `gorm:"varchar(255)" json:"title"`
	Content     string `gorm:"type:text" json:"content"`
	Author_Name string `gorm:"varchar(255)" json:"author"`
	Created_At  string `gorm:"varchar(255)" json:"created_at"`
	ImageURL    string `gorm:"varchar(255)" json:"image_url"`
}

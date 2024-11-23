package dao

import (
	"fmt"
	"mindguard/model"
)

// 获取所有文章
func GetArticles() []model.Article {
	var articles []model.Article
	err := DB.Find(&articles).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return articles
}

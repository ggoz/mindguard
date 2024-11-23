package service

import (
	"mindguard/dao"
	"mindguard/model"
)

type ArticleService struct {
}

// 退出登录业务逻辑
func (as *ArticleService) GetArticles() []model.Article {
	// 从数据库中查询所有文章
	articles := dao.GetArticles()
	if articles != nil {
		return articles
	}
	return nil
}

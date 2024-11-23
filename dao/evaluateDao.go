package dao

import (
	"fmt"
	"mindguard/model"
)

// 获取老师评价
func GetTeacherComments(teacherId int64) []model.Evaluation {
	var evaluations []model.Evaluation
	err := DB.Where("evaluated = ?", teacherId).Find(&evaluations).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return evaluations
}

// 插入评价
func InsertComment(evalation *model.Evaluation) (err error) {
	err = DB.Create(evalation).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return nil
}

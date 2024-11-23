package service

import (
	"mindguard/dao"
	"mindguard/model"
)

type EvaluateService struct {
}

// 获取老师的评价
func (es *EvaluateService) GetComments(teacherId int64) []model.Evaluation {
	// 从数据库 查询老师的评价
	return dao.GetTeacherComments(teacherId)
}

// 提交学生评价
func (es *EvaluateService) PostEvaluation(evaluator, evaluated int64, comment string) bool {
	// 从数据库 插入学生评价
	evaluation := model.Evaluation{
		Evaluator: evaluator,
		Evaluated: evaluated,
		Comment:   comment,
	}
	err := dao.InsertComment(&evaluation)
	if err != nil {
		return false
	}
	return true
}

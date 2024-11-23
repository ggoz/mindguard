package dao

import (
	"fmt"
	"mindguard/model"
	"mindguard/params"
)

// 根据用户回答和题目id查询获得分数
func QueryScoreByAnswerAndUserId(choice params.Choice) *model.Answer {
	var answer model.Answer
	err := DB.Where("question_id = ? AND answer_text = ?", choice.ID, choice.SelectedOption).Find(&answer).Error
	if err != nil {
		fmt.Println("查询获得分数失败:", err)
		return nil
	}
	return &answer
}

// 插入一条用户测试记录
func InsertIntoRecord(record *model.UserRecord) (err error) {
	err = DB.Create(record).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return nil
}

// 查询所有问题和对应答案
func GetQuestionsAndAnswers() []model.Question {
	var questions []model.Question
	err := DB.Preload("Answers").Find(&questions).Error
	if err != nil {
		// 处理错误
		fmt.Println(err)
		return nil
	}
	return questions
}

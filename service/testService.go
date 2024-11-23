package service

import (
	"fmt"
	"mindguard/dao"
	"mindguard/model"
	"mindguard/params"
	"time"
)

type TestService struct {
}

// 提交测试业务逻辑
func (ts *TestService) SubmitTest(requestData params.SubmitTestRequest) (bool, int64) {
	// 根据用户选择的答案查询获得的分数
	var score int64 = 0
	for _, choice := range requestData.Choices {
		// 从数据库查询获得的分数
		answer := dao.QueryScoreByAnswerAndUserId(choice)
		score += answer.Score
	}
	fmt.Println(score)

	// 生成用户测试记录插入数据库
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(currentTime)
	record := model.UserRecord{
		UserId:     requestData.UserId,
		TestDate:   currentTime,
		TotalScore: score,
	}
	err := dao.InsertIntoRecord(&record)
	if err != nil {
		fmt.Println("插入用户测试记录失败")
		return false, 0
	}

	return true, score
}

// 获取所有问题和对应答案业务逻辑
func (ts *TestService) GetQuestionsAndAnswers() []model.Question {
	// 从数据库中查询所有问题和对应答案
	groupAnswers := dao.GetQuestionsAndAnswers()
	if groupAnswers != nil {
		return groupAnswers
	}
	return nil
}

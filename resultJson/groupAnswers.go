package resultJson

import "mindguard/model"

type GroupedAnswers struct {
	QuestionID int64          `json:"question_id"`
	Answers    []model.Answer `json:"answers"`
}

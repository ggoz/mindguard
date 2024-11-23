package params

type UserEvaluationParams struct {
	Evaluator int64  `json:"evaluator"`
	Evaluated int64  `json:"evaluated"`
	Comment   string `json:"comment"`
}

package params

type Choice struct {
	ID             int64  `json:"id"`
	SelectedOption string `json:"selectedOption"`
}

type SubmitTestRequest struct {
	UserId  int64    `json:"userId"`
	Choices []Choice `json:"choices"`
}

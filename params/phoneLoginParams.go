package params

type PhoneLoginParams struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

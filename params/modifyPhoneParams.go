package params

type ModifyPhoneParams struct {
	Id         int    `json:"id"`
	OldPhone   string `json:"oldPhone"`
	NewPhone   string `json:"newPhone"`
	VerifyCode string `json:"verifyCode"`
}

package params

type ModifyPwdParams struct {
	Id     int    `json:"id"`
	OldPwd string `json:"oldPwd"`
	NewPwd string `json:"newPwd"`
}

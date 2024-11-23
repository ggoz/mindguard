package params

type ModifyUsernameParams struct {
	Id          int    `json:"id"`
	OldUsername string `json:"oldUsername"`
	NewUsername string `json:"newUsername"`
}

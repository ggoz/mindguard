package params

type SendMsgParams struct {
	Sender   string `json:"sender"`
	Acceptor string `json:"acceptor"`
	Message  string `json:"message"`
}

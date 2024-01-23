package contract

type ReqMessageCreate struct {
	Receiver string `json:"receiver" validate:"required,email"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}

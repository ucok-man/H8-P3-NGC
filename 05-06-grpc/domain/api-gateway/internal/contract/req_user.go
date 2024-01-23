package contract

type ReqUserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ReqUserRegister struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ReqUserCreate struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

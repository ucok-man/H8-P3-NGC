package contract

type ResUserRegister struct {
	User struct {
		UserID   int    `json:"user_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
	Message string `json:"message"`
}

type ResUserLogin struct {
	AuthenticationToken struct {
		Token  string `json:"token"`
		Expiry string `json:"expiry"`
	} `json:"auhentication_token"`
}

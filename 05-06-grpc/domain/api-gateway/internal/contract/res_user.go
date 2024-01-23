package contract

type ResUserRegister struct {
	Message string `json:"message"`
}

type ResUserLogin struct {
	AuthenticationToken struct {
		Token  string `json:"token"`
		Expiry string `json:"expiry"`
	} `json:"auhentication_token"`
}

type ResUserGetAll struct {
	Users []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"users"`
}

type ResUserCreate struct {
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	Status string `json:"status"`
}

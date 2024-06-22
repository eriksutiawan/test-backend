package usecase

type AuthRegisterDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Token   string `json:"access_token"`
	Expired int64  `json:"expired_token"`
}

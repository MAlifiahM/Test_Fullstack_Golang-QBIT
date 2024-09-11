package domain

type AuthRepository interface {
	Register(user *User) error
}

type AuthService interface {
	Login(email string, password string) (map[string]string, error)
	Register(username string, email string, password string) error
}

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

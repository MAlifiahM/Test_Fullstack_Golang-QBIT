package domain

type User struct {
	ID       int    `json:"id" gorm:"primary_key, auto_increment"`
	Username string `json:"username" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
}

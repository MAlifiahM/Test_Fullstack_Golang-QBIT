package user

import (
	"case3/internal/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	return &user, r.db.Where("email = ?", email).First(&user).Error
}

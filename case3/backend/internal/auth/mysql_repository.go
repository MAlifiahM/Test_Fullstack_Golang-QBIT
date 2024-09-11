package auth

import (
	"case3/internal/domain"
	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	db *gorm.DB
}

func NewMysqlAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &mysqlAuthRepository{db: db}
}

func (r *mysqlAuthRepository) Register(user *domain.User) error {
	return r.db.Create(user).Error
}

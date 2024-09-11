package cart

import (
	"case3/internal/domain"
	"gorm.io/gorm"
	"time"
)

type mysqlCartRepository struct {
	db *gorm.DB
}

func NewMysqlCartRepository(db *gorm.DB) domain.CartRepository {
	return &mysqlCartRepository{db: db}
}

func (r *mysqlCartRepository) GetCartByUserID(userID int) ([]domain.Cart, error) {
	var carts []domain.Cart
	return carts, r.db.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&carts).Error
}

func (r *mysqlCartRepository) DeleteCartByUserID(cartID int) error {
	return r.db.Model(&domain.Cart{}).Where("id = ? AND deleted_at IS NULL", cartID).Update("deleted_at", time.Now()).Error
}

func (r *mysqlCartRepository) CreateCart(reqCart *domain.RequestCart, total int, userID int) error {
	return r.db.Create(&domain.Cart{
		ProductID: reqCart.ProductID,
		Quantity:  reqCart.Quantity,
		Total:     total,
		UserID:    userID,
	}).Error
}

func (r *mysqlCartRepository) UpdateCart(cartID int, total int, reqCart *domain.RequestCart) error {
	return r.db.Model(&domain.Cart{}).Where("id = ? AND deleted_at IS NULL", cartID).Updates(map[string]interface{}{"quantity": reqCart.Quantity, "total": total}).Error
}

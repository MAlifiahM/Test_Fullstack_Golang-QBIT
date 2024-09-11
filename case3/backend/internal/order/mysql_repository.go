package order

import (
	"case3/internal/domain"
	"gorm.io/gorm"
)

type mysqlOrderRepository struct {
	db *gorm.DB
}

func NewMysqlOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &mysqlOrderRepository{db: db}
}

func (r *mysqlOrderRepository) GetOrderByUserID(userID int) ([]domain.Order, error) {
	var orders []domain.Order
	return orders, r.db.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&orders).Error
}

func (r *mysqlOrderRepository) CreateOrder(order *domain.RequestOrder, total int, userID int) error {
	return r.db.Create(&domain.Order{
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Total:     total,
		UserID:    userID,
	}).Error
}

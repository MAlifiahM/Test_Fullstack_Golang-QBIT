package domain

type Order struct {
	ID        int `json:"id" gorm:"primary_key, auto_increment"`
	UserID    int `json:"user_id" gorm:"not null"`
	ProductID int `json:"product_id" gorm:"not null"`
	Quantity  int `json:"quantity" gorm:"not null"`
	Total     int `json:"total" gorm:"not null"`
}

type RequestOrder struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ResponseOrder struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Image     string `json:"image"`
	Name      string `json:"name"`
	Total     int    `json:"total"`
}

type OrderRepository interface {
	GetOrderByUserID(userID int) ([]Order, error)
	CreateOrder(order *RequestOrder, total int, userID int) error
}

type OrderService interface {
	GetOrderByUserID(userID int) ([]ResponseOrder, error)
	CreateOrder(order *RequestOrder, userID int) error
}

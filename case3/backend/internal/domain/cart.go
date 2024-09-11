package domain

type Cart struct {
	ID        int `json:"id" gorm:"primary_key, auto_increment"`
	UserID    int `json:"user_id" gorm:"not null"`
	ProductID int `json:"product_id" gorm:"not null"`
	Quantity  int `json:"quantity" gorm:"not null"`
	Total     int `json:"total" gorm:"not null"`
}

type CartResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Image     string `json:"image"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Total     int    `json:"total"`
}

type RequestCart struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartRepository interface {
	GetCartByUserID(userID int) ([]Cart, error)
	DeleteCartByUserID(cartID int) error
	CreateCart(reqCart *RequestCart, total int, userID int) error
	UpdateCart(cartID int, total int, reqCart *RequestCart) error
}

type CartService interface {
	GetCartByUserID(userID int) ([]CartResponse, error)
	DeleteCartByUserID(cartID int) error
	CreateCart(reqCart *RequestCart, userID int) error
	UpdateCart(cartID int, reqCart *RequestCart) error
}

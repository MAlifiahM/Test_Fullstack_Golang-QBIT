package domain

type Product struct {
	ID          int    `json:"id" gorm:"primary_key, auto_increment"`
	Name        string `json:"name" gorm:"not null"`
	Price       int    `json:"price" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Stock       int    `json:"stock" gorm:"not null"`
	Image       string `json:"image" gorm:"not null"`
	Category    string `json:"category" gorm:"not null"`
}

type RequestProduct struct {
	Name        string `json:"name" gorm:"not null"`
	Price       int    `json:"price" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Stock       int    `json:"stock" gorm:"not null"`
	Image       string `json:"image" gorm:"not null"`
	Category    string `json:"category" gorm:"not null"`
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	UpdateStock(id int, quantity int) error
	CreateProduct(req *RequestProduct) error
	UpdateProduct(id int, req *RequestProduct) error
	DeleteProduct(id int) error
}

type ProductService interface {
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	CreateProduct(req *RequestProduct) error
	UpdateProduct(id int, req *RequestProduct) error
	DeleteProduct(id int) error
}

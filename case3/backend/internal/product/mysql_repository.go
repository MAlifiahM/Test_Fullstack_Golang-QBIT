package product

import (
	"case3/internal/domain"
	"gorm.io/gorm"
	"time"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

func NewMysqlProductRepository(db *gorm.DB) domain.ProductRepository {
	return &mysqlProductRepository{db: db}
}

func (r *mysqlProductRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	return products, r.db.Where("deleted_at IS NULL").Find(&products).Error
}

func (r *mysqlProductRepository) GetByID(id int) (domain.Product, error) {
	var product domain.Product
	return product, r.db.Where("deleted_at IS NULL").First(&product, id).Error
}

func (r *mysqlProductRepository) UpdateStock(id int, quantity int) error {
	return r.db.Model(&domain.Product{}).Where("id = ? AND deleted_at IS NULL", id).Update("stock", gorm.Expr("stock - ?", quantity)).Error
}

func (r *mysqlProductRepository) CreateProduct(req *domain.RequestProduct) error {
	product := domain.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Stock:       req.Stock,
		Image:       req.Image,
		Category:    req.Category,
	}
	return r.db.Create(&product).Error
}

func (r *mysqlProductRepository) UpdateProduct(id int, req *domain.RequestProduct) error {
	return r.db.Model(&domain.Product{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&req).Error
}

func (r *mysqlProductRepository) DeleteProduct(id int) error {
	return r.db.Model(&domain.Product{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error
}

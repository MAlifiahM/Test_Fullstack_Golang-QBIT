package product

import "case3/internal/domain"

type productService struct {
	productRepo domain.ProductRepository
}

func NewProductService(productRepo domain.ProductRepository) domain.ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) GetAll() ([]domain.Product, error) {
	return s.productRepo.GetAll()
}

func (s *productService) GetByID(id int) (domain.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *productService) CreateProduct(req *domain.RequestProduct) error {
	return s.productRepo.CreateProduct(req)
}

func (s *productService) UpdateProduct(id int, req *domain.RequestProduct) error {
	return s.productRepo.UpdateProduct(id, req)
}

func (s *productService) DeleteProduct(id int) error {
	return s.productRepo.DeleteProduct(id)
}

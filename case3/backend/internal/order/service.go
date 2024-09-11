package order

import "case3/internal/domain"

type orderService struct {
	orderRepo   domain.OrderRepository
	productRepo domain.ProductRepository
}

func NewOrderService(orderRepo domain.OrderRepository, productRepo domain.ProductRepository) domain.OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *orderService) GetOrderByUserID(userID int) ([]domain.ResponseOrder, error) {
	orders, err := s.orderRepo.GetOrderByUserID(userID)

	if err != nil {
		return nil, err
	}

	var orderResponses []domain.ResponseOrder

	for _, order := range orders {
		product, _ := s.productRepo.GetByID(order.ProductID)

		orderResponses = append(orderResponses, domain.ResponseOrder{
			ID:        order.ID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
			Image:     product.Image,
			Name:      product.Name,
			Total:     order.Total,
		})
	}

	return orderResponses, nil
}

func (s *orderService) CreateOrder(order *domain.RequestOrder, userID int) error {

	product, err := s.productRepo.GetByID(order.ProductID)

	if err != nil {
		return err
	}

	if product.Stock < order.Quantity {
		return err
	}

	err = s.productRepo.UpdateStock(product.ID, order.Quantity)
	if err != nil {
		return err
	}

	total := product.Price * order.Quantity

	return s.orderRepo.CreateOrder(order, total, userID)
}

package cart

import "case3/internal/domain"

type cartService struct {
	cartRepo    domain.CartRepository
	productRepo domain.ProductRepository
}

func NewCartService(cartRepo domain.CartRepository, productRepo domain.ProductRepository) domain.CartService {
	return &cartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *cartService) GetCartByUserID(userID int) ([]domain.CartResponse, error) {
	carts, err := s.cartRepo.GetCartByUserID(userID)

	if err != nil {
		return nil, err
	}

	var cartResponses []domain.CartResponse

	for _, cart := range carts {
		product, _ := s.productRepo.GetByID(cart.ProductID)

		cartResponses = append(cartResponses, domain.CartResponse{
			ID:        cart.ID,
			UserID:    cart.UserID,
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Image:     product.Image,
			Name:      product.Name,
			Price:     product.Price,
			Total:     cart.Total,
		})
	}

	return cartResponses, nil
}

func (s *cartService) DeleteCartByUserID(cartID int) error {
	return s.cartRepo.DeleteCartByUserID(cartID)
}

func (s *cartService) CreateCart(reqCart *domain.RequestCart, userID int) error {
	product, err := s.productRepo.GetByID(reqCart.ProductID)

	if err != nil {
		return err
	}

	if product.Stock < reqCart.Quantity {
		return err
	}

	total := product.Price * reqCart.Quantity

	return s.cartRepo.CreateCart(reqCart, total, userID)
}

func (s *cartService) UpdateCart(cartID int, reqCart *domain.RequestCart) error {
	product, err := s.productRepo.GetByID(reqCart.ProductID)

	if err != nil {
		return err
	}

	if product.Stock < reqCart.Quantity {
		return err
	}

	total := product.Price * reqCart.Quantity

	return s.cartRepo.UpdateCart(cartID, total, reqCart)
}

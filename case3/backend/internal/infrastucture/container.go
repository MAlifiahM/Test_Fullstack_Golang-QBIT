package infrastucture

import (
	"case3/internal/auth"
	"case3/internal/cart"
	"case3/internal/config"
	"case3/internal/domain"
	"case3/internal/order"
	"case3/internal/product"
	"case3/internal/user"
	"case3/pkg/xlogger"
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

var (
	cfg config.Config

	userRepo domain.UserRepository

	authRepo    domain.AuthRepository
	authService domain.AuthService

	productRepo    domain.ProductRepository
	productService domain.ProductService

	orderRepo    domain.OrderRepository
	orderService domain.OrderService

	cartRepo    domain.CartRepository
	cartService domain.CartService
)

func init() {
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	xlogger.Setup(cfg)
	dbSetup()

	userRepo = user.NewMysqlUserRepository(db)

	authRepo = auth.NewMysqlAuthRepository(db)
	authService = auth.NewAuthService(authRepo, userRepo, cfg.SecretKey)

	productRepo = product.NewMysqlProductRepository(db)
	productService = product.NewProductService(productRepo)

	orderRepo = order.NewMysqlOrderRepository(db)
	orderService = order.NewOrderService(orderRepo, productRepo)

	cartRepo = cart.NewMysqlCartRepository(db)
	cartService = cart.NewCartService(cartRepo, productRepo)
}

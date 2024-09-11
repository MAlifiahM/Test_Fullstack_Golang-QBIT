package infrastucture

import (
	"case3/internal/auth"
	"case3/internal/cart"
	"case3/internal/middleware/jwt"
	"case3/internal/order"
	"case3/internal/product"
	"case3/pkg/xlogger"
	"fmt"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Run() {
	logger := xlogger.Logger

	app := fiber.New(fiber.Config{
		ProxyHeader:           cfg.ProxyHeader,
		DisableStartupMessage: true,
		ErrorHandler:          defaultErrorHandler,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
		Fields: cfg.LogFields,
	}))

	app.Use(recover2.New())
	app.Use(etag.New())
	app.Use(requestid.New())
	app.Use(cors.New())

	api := app.Group("/api")

	// route for auth
	auth.NewHttpHandler(api.Group("/auth"), authService)

	// route for product

	publicProductRoutes := api.Group("/product")
	product.NewHttpPublicProductHandler(publicProductRoutes, productService)

	protectedProductRoutes := api.Group("/product", jwt.JWTMiddleware(cfg.SecretKey))
	product.NewHttpPrivateProductHandler(protectedProductRoutes, productService)

	// route for cart
	cartRoutes := api.Group("/cart", jwt.JWTMiddleware(cfg.SecretKey))
	cart.NewHttpCartHandler(cartRoutes, cartService)

	// route for order
	orderRoutes := api.Group("/order", jwt.JWTMiddleware(cfg.SecretKey))
	order.NewHttpOrderHandler(orderRoutes, orderService)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	logger.Info().Msgf("Server is running on address: %s", addr)
	if err := app.Listen(addr); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}

package controllerBootstrapper

import (
	"database/sql"
	authController "forecasting-be/internal/controller/auth"
	forecastController "forecasting-be/internal/controller/forecast"
	orderController "forecasting-be/internal/controller/orders"
	pingController "forecasting-be/internal/controller/ping"
	productController "forecasting-be/internal/controller/products"
	authRepository "forecasting-be/internal/repository/auth"
	ordersRepository "forecasting-be/internal/repository/orders"
	productRepository "forecasting-be/internal/repository/products"
	authService "forecasting-be/internal/service/auth"
	forecastService "forecasting-be/internal/service/forecast"
	orderService "forecasting-be/internal/service/orders"
	pingService "forecasting-be/internal/service/ping"
	productService "forecasting-be/internal/service/products"
	"forecasting-be/pkg/utilities"

	"github.com/gorilla/mux"
)

func InitializeEndpoints(router *mux.Router, db *sql.DB, smtpConfig utilities.SMTPConfig) {
	pingService := pingService.NewPingService()

	pingController := pingController.NewPingController(router, pingService)
	pingController.InitEndpoints()

	orderRepository := ordersRepository.NewOrderRepository(db)
	orderService := orderService.NewOrderService(orderRepository)
	orderController := orderController.NewOrderController(router, orderService)
	orderController.InitializeEndpoints()

	authRepository := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepository, smtpConfig)
	authController := authController.NewAuthController(router, authService)
	authController.InitializeEndpoints()

	productRepository := productRepository.NewProductRepository(db)
	productService := productService.NewProductService(productRepository)
	productController := productController.NewProductController(router, productService)
	productController.InitializeEndpoints()

	forecastService := forecastService.NewForecastService(productRepository)
	forecastController := forecastController.NewForecastController(router, forecastService)
	forecastController.InitializeEndpoints()

}

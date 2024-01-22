package controllerBootstrapper

import (
	"database/sql"
	authController "forecasting-be/internal/controller/auth"
	orderController "forecasting-be/internal/controller/orders"
	pingController "forecasting-be/internal/controller/ping"
	authRepository "forecasting-be/internal/repository/auth"
	ordersRepository "forecasting-be/internal/repository/orders"
	authService "forecasting-be/internal/service/auth"
	orderService "forecasting-be/internal/service/orders"
	pingService "forecasting-be/internal/service/ping"

	"github.com/gorilla/mux"
)

func InitializeEndpoints(router *mux.Router, db *sql.DB) {
	pingService := pingService.NewPingService()

	pingController := pingController.NewPingController(router, pingService)
	pingController.InitEndpoints()

	orderRepository := ordersRepository.NewOrderRepository(db)
	orderService := orderService.NewOrderService(orderRepository)
	orderController := orderController.NewOrderController(router, orderService)
	orderController.InitializeEndpoints()

	authRepository := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepository)
	authController := authController.NewAuthController(router, authService)
	authController.InitializeEndpoints()

}

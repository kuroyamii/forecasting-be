package controllerBootstrapper

import (
	pingController "forecasting-be/internal/controller/ping"
	pingService "forecasting-be/internal/service/ping"

	"github.com/gorilla/mux"
)

func InitializeEndpoints(router *mux.Router) {
	pingService := pingService.NewPingService()

	pingController := pingController.NewPingController(router, pingService)
	pingController.InitEndpoints()
}

package forecastController

import (
	"forecasting-be/pkg/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type forecastController struct {
	r *mux.Router
}

func NewForecastController(router *mux.Router) forecastController {
	return forecastController{
		r: router,
	}
}

func (fc forecastController) InitializeEndpoints() {
	fc.r.HandleFunc("/forecast", fc.handleForecast).Methods("POST").Subrouter().Use(middlewares.ValidateAdminJWT)
}

func (fc forecastController) handleForecast(rw http.ResponseWriter, r *http.Request) {

}

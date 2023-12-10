package orderController

import (
	orderService "forecasting-be/internal/service/orders"
	baseResponse "forecasting-be/pkg/response"
	"net/http"

	"github.com/gorilla/mux"
)

type orderController struct {
	r  *mux.Router
	os orderService.OrderService
}

func NewOrderController(r *mux.Router, os orderService.OrderService) orderController {
	return orderController{
		r:  r,
		os: os,
	}
}

func (oc orderController) HandleGetOrders(rw http.ResponseWriter, r *http.Request) {
	if data, err := oc.os.GetOrders(r.Context()); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		baseResponse.NewBaseResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), baseResponse.NewErrorResponses(
			baseResponse.ErrorResponse{
				Key:   "server error",
				Value: err.Error(),
			},
		), nil)
	} else {
		rw.WriteHeader(http.StatusOK)
		baseResponse.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, data).ToJSON(rw)

	}

}

func (oc orderController) InitializeEndpoints() {
	oc.r.HandleFunc("/orders", oc.HandleGetOrders).Methods("GET")
}

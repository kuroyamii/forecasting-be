package orderController

import (
	orderService "forecasting-be/internal/service/orders"
	"forecasting-be/pkg/middlewares"
	baseResponse "forecasting-be/pkg/response"
	"forecasting-be/pkg/utilities"
	"log"
	"net/http"
	"strconv"

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
	query := r.URL.Query()
	paginateQuery := query.Get("paginate")
	page, err := strconv.ParseInt(paginateQuery, 10, 64)
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
	}
	if data, err := oc.os.GetOrders(r.Context(), page); err != nil {
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

func (oc orderController) HandleGetSalesSum(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	salesSumRequest := query.Get("month")
	month, err := strconv.ParseInt(salesSumRequest, 10, 64)
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		baseResponse.NewBaseResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			baseResponse.ErrorResponse{
				Key:   "parsing error",
				Value: err.Error(),
			},
			nil).ToJSON(rw)
		return
	}
	res, err := oc.os.GetSalesSum(r.Context(), int(month))
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		baseResponse.NewBaseResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			baseResponse.ErrorResponse{
				Key:   "parsing error",
				Value: err.Error(),
			},
			nil).ToJSON(rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		res).ToJSON(rw)
}

func (oc orderController) HandleGetTotalProduct(rw http.ResponseWriter, r *http.Request) {
	res, err := oc.os.GetTotalProduct(r.Context())
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		baseResponse.NewBaseResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			baseResponse.ErrorResponse{
				Key:   "parsing error",
				Value: err.Error(),
			},
			nil).ToJSON(rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		res).ToJSON(rw)
}
func (oc orderController) HandleGetMostBoughtCategory(rw http.ResponseWriter, r *http.Request) {
	res, err := oc.os.GetMostBoughtCategory(r.Context())
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		baseResponse.NewBaseResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			baseResponse.ErrorResponse{
				Key:   "parsing error",
				Value: err.Error(),
			},
			nil).ToJSON(rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		res).ToJSON(rw)
}

func (oc orderController) HandleGetTopTransaction(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limitString := query.Get("limit")
	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		baseResponse.NewBaseResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			baseResponse.ErrorResponse{
				Key:   "parsing error",
				Value: err.Error(),
			},
			nil).ToJSON(rw)
		return
	}
	res, err := oc.os.GetTopTransactions(r.Context(), int(limit))
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		baseResponse.NewBaseResponse(http.StatusBadRequest,
			http.StatusText(http.StatusInternalServerError),
			baseResponse.ErrorResponse{
				Key:   "error",
				Value: err.Error(),
			},
			nil).ToJSON(rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		res).ToJSON(rw)
}

func (oc orderController) InitializeEndpoints() {
	adminRouter := oc.r.PathPrefix("").Subrouter()
	adminRouter.Use(middlewares.ValidateAdminJWT)
	adminRouter.HandleFunc("/orders", oc.HandleGetOrders).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/sales-growth", oc.HandleGetSalesSum).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/total-product", oc.HandleGetTotalProduct).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/most-bought-category", oc.HandleGetMostBoughtCategory).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/top-transactions", oc.HandleGetTopTransaction).Methods("GET", "OPTIONS")
}

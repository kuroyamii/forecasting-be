package productController

import (
	productService "forecasting-be/internal/service/products"
	"forecasting-be/pkg/middlewares"
	baseResponse "forecasting-be/pkg/response"
	"forecasting-be/pkg/utilities"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type productController struct {
	r  *mux.Router
	ps productService.ProductService
}

func NewProductController(r *mux.Router, ps productService.ProductService) productController {
	return productController{
		r:  r,
		ps: ps,
	}
}

func (pc productController) InitializeEndpoints() {
	adminRouter := pc.r.NewRoute().Subrouter()
	adminRouter.Use(middlewares.ValidateAdminJWT)
	adminRouter.HandleFunc("/product-summaries", pc.handleProductSummary).Methods("GET")
}

func (pc productController) handleProductSummary(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pageQs := query.Get("page")
	page, err := strconv.ParseInt(pageQs, 10, 64)
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
	}
	limitQs := query.Get("limit")
	limit, err := strconv.ParseInt(limitQs, 10, 64)
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
	}

	res, err := pc.ps.GetProductSummary(r.Context(), int(page), int(limit))
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

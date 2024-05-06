package forecastController

import (
	"forecasting-be/internal/dto"
	forecastService "forecasting-be/internal/service/forecast"
	"forecasting-be/pkg/middlewares"
	baseResponse "forecasting-be/pkg/response"
	"forecasting-be/pkg/utilities"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type forecastController struct {
	r  *mux.Router
	fs forecastService.ForecastServie
}

func NewForecastController(router *mux.Router, fs forecastService.ForecastServie) forecastController {
	return forecastController{
		r:  router,
		fs: fs,
	}
}

func (fc forecastController) InitializeEndpoints() {
	adminRouter := fc.r.PathPrefix("").Subrouter()
	adminRouter.Use(middlewares.ValidateAdminJWT)
	adminRouter.HandleFunc("/forecast", fc.handleForecast).Methods("POST", "OPTIONS")
}

func (fc forecastController) handleForecast(rw http.ResponseWriter, r *http.Request) {
	// Initialize empty struct
	var forecastRequest dto.ForecastRequest

	// Decode request body
	err := utilities.JSONDecode(r.Body, &forecastRequest)
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

	// Validate request body
	validate := validator.New()
	err = validate.Struct(forecastRequest)
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		baseResponse.NewBaseResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			baseResponse.ErrorResponse{
				Key:   "validation error",
				Value: err.Error(),
			},
			nil).ToJSON(rw)
		return
	}

	// Forecast sales
	res, err := fc.fs.ForecastSales(r.Context(), forecastRequest.Month, forecastRequest.Year, forecastRequest.SubCategoryID)
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		baseResponse.NewBaseResponse(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			baseResponse.ErrorResponse{
				Key:   "internal server error",
				Value: err.Error(),
			},
			nil).ToJSON(rw)
		return
	}

	// Send success response
	rw.WriteHeader(http.StatusOK)
	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		res).ToJSON(rw)
}

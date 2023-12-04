package pingController

import (
	pingService "forecasting-be/internal/service/ping"
	baseResponse "forecasting-be/pkg/response"
	"net/http"

	"github.com/gorilla/mux"
)

type PingController struct {
	router *mux.Router
	ps     pingService.PingServiceImpl
}

func (pc PingController) HandleGetPing(rw http.ResponseWriter, r *http.Request) {
	data := pc.ps.GetPing()
	rw.WriteHeader(http.StatusOK)
	if err := baseResponse.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, data).ToJSON(rw); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		baseResponse.NewBaseResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), baseResponse.NewErrorResponses(
			baseResponse.ErrorResponse{
				Key:   "server error",
				Value: "error getting ping",
			},
		), nil)
	}
}

func NewPingController(router *mux.Router, ps pingService.PingServiceImpl) PingController {
	return PingController{router: router, ps: ps}
}

func (pc PingController) InitEndpoints() {
	pc.router.HandleFunc("/ping", pc.HandleGetPing).Methods("GET")
}

package authController

import (
	"forecasting-be/internal/dto"
	authService "forecasting-be/internal/service/auth"
	baseResponse "forecasting-be/pkg/response"
	"forecasting-be/pkg/utilities"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type authController struct {
	router *mux.Router
	as     authService.AuthService
}

func NewAuthController(router *mux.Router, as authService.AuthService) authController {
	return authController{
		router: router,
		as:     as,
	}
}

func (ac authController) InitializeEndpoints() {
	ac.router.HandleFunc("/signup", ac.handleSignUp).Methods("POST")
	ac.router.HandleFunc("/signin", ac.handleSignIn).Methods("POST")
	ac.router.HandleFunc("/token/refresh", ac.handleRefreshToken).Methods("POST")
}

func (ac authController) handleRefreshToken(rw http.ResponseWriter, r *http.Request) {
	var refreshTokenRequest dto.RefreshRequest
	err := utilities.JSONDecode(r.Body, &refreshTokenRequest)
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
	validate := validator.New()
	err = validate.Struct(refreshTokenRequest)
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

	res, err := ac.as.RegenerateToken(r.Context(), refreshTokenRequest.RefreshToken)
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

	rw.WriteHeader(http.StatusOK)
	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		res).ToJSON(rw)
}

func (ac authController) handleSignIn(rw http.ResponseWriter, r *http.Request) {
	var loginRequest dto.SignInRequest
	err := utilities.JSONDecode(r.Body, &loginRequest)
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

	validate := validator.New()
	err = validate.Struct(loginRequest)
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

	res, err := ac.as.SignIn(r.Context(), loginRequest.Username, loginRequest.Password)
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

	rw.WriteHeader(http.StatusOK)
	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		res).ToJSON(rw)
}

func (ac authController) handleSignUp(rw http.ResponseWriter, r *http.Request) {
	var userRequest dto.RegisterRequest
	err := utilities.JSONDecode(r.Body, &userRequest)

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

	validate := validator.New()
	err = validate.Struct(userRequest)
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

	_, err = ac.as.SignUp(r.Context(), userRequest)
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

	rw.WriteHeader(http.StatusOK)
	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		nil).ToJSON(rw)
}

package authController

import (
	"forecasting-be/internal/dto"
	authService "forecasting-be/internal/service/auth"
	"forecasting-be/pkg/middlewares"
	baseResponse "forecasting-be/pkg/response"
	"forecasting-be/pkg/utilities"
	"log"
	"net/http"
	"time"

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
	ac.router.HandleFunc("/signup", ac.handleSignUp).Methods("POST", "OPTIONS")
	ac.router.HandleFunc("/signin", ac.handleSignIn).Methods("POST", "OPTIONS")
	ac.router.HandleFunc("/token/refresh", ac.handleRefreshToken).Methods("POST", "OPTIONS")

	adminSubRouter := ac.router.PathPrefix("/admin").Subrouter()
	adminSubRouter.Use(middlewares.ValidateSuperAdminJWT)
	adminSubRouter.HandleFunc("/invite", ac.handleInviteAdmin).Methods("POST", "OPTIONS")
}

// handleInviteAdmin is for inviting admin so the admin can do signup
func (ac authController) handleInviteAdmin(rw http.ResponseWriter, r *http.Request) {
	// Initialize empty struct
	var request dto.AdminRequest

	// Decode request body
	err := utilities.JSONDecode(r.Body, &request)
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

	// Validate request
	validate := validator.New()
	err = validate.Struct(request)
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

	// invite admin
	err = ac.as.AddAdmin(r.Context(), request.Email, request.Role)
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
		nil).ToJSON(rw)
}

// handleRefreshToken handles logic to refresh the access token with refresh token
func (ac authController) handleRefreshToken(rw http.ResponseWriter, r *http.Request) {
	storedCookie, err := r.Cookie("jwt")
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		baseResponse.NewBaseResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error(), nil).ToJSON(rw)
		return
	}
	if storedCookie == nil || storedCookie.Value == "" {
		rw.WriteHeader(http.StatusBadRequest)
		baseResponse.NewBaseResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "no token", nil).ToJSON(rw)
		return
	}
	// Regenerate access token
	res, err := ac.as.RegenerateToken(r.Context(), storedCookie.Value)
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
	cookie := &http.Cookie{}
	cookie.Name = "jwt"
	cookie.Value = res.RefreshToken
	tokenEnv := utilities.GetTokenEnv()
	cookie.Expires = time.Now().Add(time.Hour * time.Duration(tokenEnv.RefreshTokenTTLHour))
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.SameSite = 4
	cookie.Secure = true
	http.SetCookie(rw, cookie)

	// Send success response
	rw.WriteHeader(http.StatusOK)
	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		dto.AccessTokenOnlyResponse{
			AccessToken: res.AccessToken,
		}).ToJSON(rw)
}

// handleSignIn handles logic for signing in user
func (ac authController) handleSignIn(rw http.ResponseWriter, r *http.Request) {
	// Initialize empty struct
	var loginRequest dto.SignInRequest

	// Decode request body
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

	// Validate request body
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

	// Signin process
	res, err := ac.as.SignIn(r.Context(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		log.Printf("%v %v\n", utilities.Red("ERROR"), err.Error())
		var statusCode int
		if err.Error() == "unauthorized" {
			statusCode = http.StatusUnauthorized
		} else {
			statusCode = http.StatusInternalServerError
		}
		rw.WriteHeader(statusCode)
		baseResponse.NewBaseResponse(statusCode,
			http.StatusText(statusCode),
			baseResponse.ErrorResponse{
				Key:   "internal server error",
				Value: err.Error(),
			},
			nil).ToJSON(rw)
		return
	}
	cookie := &http.Cookie{}
	cookie.Name = "jwt"
	cookie.Path = "/"
	cookie.Value = res.RefreshToken
	tokenEnv := utilities.GetTokenEnv()
	cookie.Expires = time.Now().Add(time.Hour * time.Duration(tokenEnv.RefreshTokenTTLHour))
	cookie.HttpOnly = true
	cookie.SameSite = 4
	cookie.Secure = true
	http.SetCookie(rw, cookie)
	// Send success response
	rw.WriteHeader(http.StatusOK)
	accessToken := dto.AccessTokenOnlyResponse{
		AccessToken: res.AccessToken,
	}

	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		accessToken).ToJSON(rw)
}

// handleSignUp handles logic for signing up invited admins
func (ac authController) handleSignUp(rw http.ResponseWriter, r *http.Request) {
	// Initialize empty struct
	var userRequest dto.RegisterRequest

	// Decode Request Body
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

	// Validate request body
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
	// Signup user
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

	// Send success response
	rw.WriteHeader(http.StatusOK)
	baseResponse.NewBaseResponse(http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		nil).ToJSON(rw)
}

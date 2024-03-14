package middlewares

import (
	"errors"
	"forecasting-be/pkg/utilities"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func ContentTypeJSON(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(rw, r)
	})
}

func LoggerMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		method := r.Method
		Origin := r.Host

		switch method {
		case "GET":
			{
				method = utilities.Green(method)
			}
		case "POST":
			{
				method = utilities.Yellow(method)
			}
		case "DELETE":
			{
				method = utilities.Red(method)
			}
		case "UPDATE":
			{
				method = utilities.Teal(method)
			}
		}
		log.Printf("%v %v", method, Origin)
		handler.ServeHTTP(rw, r)
	})
}

func ValidateAdminJWT(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			authorization := r.Header.Get("Authorization")
			auth := strings.Split(authorization, " ")[1]
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(auth, claims, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, errors.New("unauthorized")
				}
				jwtKey := os.Getenv("JWT_KEY")
				return []byte(jwtKey), nil
			})
			if err != nil {
				rw.WriteHeader(http.StatusForbidden)
				return
			}

			if token.Valid {
				data := claims["data"].(map[string]interface{})
				parsedRoleID := int(data["role_id"].(float64))
				if parsedRoleID == 1 || parsedRoleID == 2 {
					handler.ServeHTTP(rw, r)
					return
				} else {
					rw.WriteHeader(http.StatusForbidden)
					return

				}
			}
		} else {
			rw.WriteHeader(http.StatusForbidden)
			return
		}

	})
}

func ValidateSuperAdminJWT(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			authorization := r.Header.Get("Authorization")
			auth := strings.Split(authorization, " ")[1]
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(auth, claims, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, errors.New("unauthorized")
				}
				jwtKey := os.Getenv("JWT_KEY")
				return []byte(jwtKey), nil
			})
			if err != nil {
				rw.WriteHeader(http.StatusForbidden)
				return
			}

			if token.Valid {
				data := claims["data"].(map[string]interface{})
				parsedRoleID := int(data["role_id"].(float64))
				if parsedRoleID == 2 {
					handler.ServeHTTP(rw, r)
					return
				} else {
					rw.WriteHeader(http.StatusForbidden)
					return

				}
			}
		} else {
			rw.WriteHeader(http.StatusForbidden)
			return
		}

	})
}

func CORSMiddleware(whitelistedUrls map[string]bool) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, PATCH")
			rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")

			requestOriginUrl := r.Header.Get("Origin")
			log.Printf("%v CorsMiddleware: received request from %v", utilities.Info("INFO"), requestOriginUrl)
			if whitelistedUrls[requestOriginUrl] {
				rw.Header().Set("Access-Control-Allow-Origin", requestOriginUrl)

			}
			if r.Method != http.MethodOptions {
				handler.ServeHTTP(rw, r)
				return
			}
		})
	}
}

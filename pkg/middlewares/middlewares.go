package middlewares

import (
	"forecasting-be/pkg/utilities"
	"log"
	"net/http"
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

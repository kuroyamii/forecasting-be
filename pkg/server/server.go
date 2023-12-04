package server

import (
	"forecasting-be/pkg/utilities"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	Address string
	Handler http.Handler
}

func (s Server) ListenAndServe() {
	go func() {
		if err := http.ListenAndServe(s.Address, s.Handler); err != nil {
			log.Printf("%v (starting server): %v\n", utilities.Fatal("ERROR"), err.Error())
		}
	}()
	log.Printf("%v (starting server): server started, listening to %v\n", utilities.Info("INFO"), s.Address)
	relay := make(chan os.Signal, 1)
	signal.Notify(relay, os.Interrupt)
	<-relay
	log.Printf("%v (shutdown server): server shutted down", utilities.Info("INFO"))
}

func NewServer(address string, handler http.Handler) *Server {
	return &Server{
		Address: address,
		Handler: handler,
	}
}

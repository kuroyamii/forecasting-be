package main

import (
	controllerBootstrapper "forecasting-be/internal/controller_bootstrapper"
	"forecasting-be/pkg/logger"
	"forecasting-be/pkg/middlewares"
	"forecasting-be/pkg/server"
	"forecasting-be/pkg/utilities"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func initializeGlobalRouter(envVars map[string]string) *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.ContentTypeJSON)
	router.Use(middlewares.LoggerMiddleware)

	return router
}

func getEnvironmentVariables() map[string]string {
	env := make(map[string]string)
	env["SERVER_ADDRESS"] = os.Getenv("SERVER_ADDRESS")
	env["WHITELISTED_URLS"] = os.Getenv("WHITELISTED_URLS")
	env["DB_NAME"] = os.Getenv("DB_NAME")
	env["DB_ADDRESS"] = os.Getenv("DB_ADDRESS")
	env["DB_UNAME"] = os.Getenv("DB_UNAME")
	env["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	return env
}

func initLogger() {
	log.SetFlags(0)
	log.SetOutput(new(logger.LogWritter))
}

func main() {
	initLogger()
	if err := godotenv.Load(); err != nil {
		log.Printf("%v (server): %v/n", utilities.Red("ERROR"), err.Error())
	}

	environmentVariables := getEnvironmentVariables()
	db := utilities.GetDatabase(environmentVariables["DB_ADDRESS"], environmentVariables["DB_UNAME"], environmentVariables["DB_PASSWORD"], environmentVariables["DB_NAME"])
	router := initializeGlobalRouter(environmentVariables)
	controllerBootstrapper.InitializeEndpoints(router, db)
	server := server.NewServer(":8080", router)
	server.ListenAndServe()
}

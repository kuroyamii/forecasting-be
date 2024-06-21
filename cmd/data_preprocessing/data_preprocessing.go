package main

import (
	"encoding/csv"
	"fmt"
	"forecasting-be/pkg/utilities"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

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

func loadDatasetFile() []map[string]interface{} {
	file, err := os.Open("./dataset.csv")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	field_names := records[0]
	fmt.Println(field_names)
	records = records[1:]
	var dataset []map[string]interface{}
	for _, record := range records {
		data := make(map[string]interface{})
		for index, item := range record {
			data[field_names[index]] = item
		}
		dataset = append(dataset, data)
	}
	return dataset
}

func getData(string selector)

func main() {
	// dataset := loadDatasetFile()
	godotenv.Load()
	envVariables := getEnvironmentVariables()
	// fmt.Println(envVariables)
	db := utilities.GetDatabase(envVariables["DB_ADDRESS"], envVariables["DB_UNAME"], envVariables["DB_PASSWORD"], envVariables["DB_NAME"])

}

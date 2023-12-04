package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
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

	dataset := map[string]int{}
	for index, column := range records[0] {
		dataset[column] = index
	}

	for key, value := range dataset {
		fmt.Println(key, value)
	}
}

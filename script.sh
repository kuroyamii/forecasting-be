#!/bin/bash
if [[ "$1" == "data_preprocessing" ]]; then
go run ./cmd/data_preprocessing/data_preprocessing.go
else
go run ./cmd/main/main.go
fi
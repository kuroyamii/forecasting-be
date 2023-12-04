package baseResponse

import (
	"encoding/json"
	"io"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func NewBaseResponse(code int, message string, errors interface{}, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Errors:  errors,
		Data:    data,
	}
}

func (br *BaseResponse) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(&br)
}

type ErrorResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewErrorResponses(errorValue ...ErrorResponse) []ErrorResponse {
	errors := []ErrorResponse{}
	for _, value := range errorValue {
		errors = append(errors, ErrorResponse{
			Key:   value.Key,
			Value: value.Value,
		})
	}
	return errors
}

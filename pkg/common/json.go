package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var exceptionErrorResp = []byte(`{"status": 500,"message": "Internal Error"}`)

type respJSON struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// MakeGeneralResp returns the marshaled respJSON []byte with general status code and message.
// The message references from http.StatusText(code)
func MakeGeneralResp(statusCode int) []byte {
	r, err := respToString(&respJSON{
		Status:  statusCode,
		Message: http.StatusText(statusCode),
	})
	if err != nil {
		// This case should not be happened
		log.Println(err)
		return exceptionErrorResp
	}
	return r
}

// MakeGeneralResp returns the marshaled respJSON []byte with status code, message and data which could be marshal to the JSON format.
func MakeDataResp(statusCode int, message string, data interface{}) ([]byte, error) {
	return respToString(&respJSON{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

func MakeMessageResp(statusCode int, message string) []byte {
	r, err := respToString(&respJSON{
		Status:  statusCode,
		Message: message,
	})
	if err != nil {
		// This case should not be happened
		log.Println(err)
		return exceptionErrorResp
	}
	return r
}

func respToString(r *respJSON) ([]byte, error) {
	respStr, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Marshal JSON Failed: %w", err)
	}

	return respStr, nil
}

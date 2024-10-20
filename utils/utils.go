package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Errors  []SingleError `json:"errors"`
}

func (er *ErrorResponse) AddErrors(se SingleError) {
	er.Errors = append(er.Errors, se)
}

type SingleError struct {
	Detail string `json:"detail"`
	Source string `json:"source"`
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	se := SingleError{
		Detail: message,
		Source: "",
	}
	errorResponse := ErrorResponse{
		Message: message,
		Status:  statusCode,
	}
	errorResponse.AddErrors(se)
	_ = json.NewEncoder(w).Encode(errorResponse)
}

func RespondWithValidationErrors(w http.ResponseWriter, verr validator.ValidationErrors) {
	errResp := ErrorResponse{
		Message: "Validation error",
		Status:  http.StatusUnprocessableEntity,
	}
	for _, err := range verr {
		errResp.AddErrors(SingleError{
			Detail: fmt.Sprintf(
				"%s is %s and should be of type %s",
				err.Field(),
				err.Tag(),
				err.Type(),
			),
			Source: err.Field(),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errResp.Status)
	_ = json.NewEncoder(w).Encode(errResp)
}

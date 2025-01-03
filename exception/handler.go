package exception

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, err error) {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		log.Printf("Unhandled error: %v", err)
		appErr = &AppError{
			ErrCode: Unknown,
			Message: err.Error(),
			Err:     err,
		}
	}
	log.Printf("AppError occurred: Code=%s, Message=%s, Error=%v",
		appErr.ErrCode,
		appErr.Message,
		appErr.Err,
	)

	var statusCode int

	switch appErr.ErrCode {
	case NotFound:
		statusCode = http.StatusNotFound
	case Invalid:
		statusCode = http.StatusUnprocessableEntity
	case Conflict:
		statusCode = http.StatusConflict
	default:
		statusCode = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Code:    appErr.ErrCode,
		Message: appErr.Message,
	})
}

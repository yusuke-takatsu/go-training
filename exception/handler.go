package exception

import (
	"encoding/json"
	"errors"
	"net/http"
)

func Handler(w http.ResponseWriter, err error) {
	var appErr *AppError
	if !errors.As(err, &appErr) {
		appErr = &AppError{
			ErrCode: Unknown,
			Message: err.Error(),
			Err:     err,
		}
	}

	var statusCode int

	switch appErr.ErrCode {
	case NotFound:
		statusCode = http.StatusNotFound
	case InValid:
		statusCode = http.StatusUnprocessableEntity
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}

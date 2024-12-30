package exception

type AppError struct {
	ErrCode
	Message string
	Err     error
}

type ErrorResponse struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

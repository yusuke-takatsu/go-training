package exception

type AppError struct {
	ErrCode
	Message string
	Err     error
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

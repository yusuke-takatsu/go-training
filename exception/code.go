package exception

type ErrCode string

const (
	Unknown ErrCode = "U000"

	NotFound     ErrCode = "S001"
	InsertFailed ErrCode = "S002"
	Invalid      ErrCode = "S003"
)

func (code ErrCode) Wrap(err error, message string) *AppError {
	return &AppError{ErrCode: code, Message: message, Err: err}
}

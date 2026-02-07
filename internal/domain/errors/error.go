package errors

type ErrCode string

func (e ErrCode) Error() string {
	return string(e)
}

type AppError struct {
	Code    ErrCode
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return string(e.Code)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

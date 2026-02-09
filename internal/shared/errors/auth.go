package errors

const (
	ErrAuthInvalidToken           ErrCode = "AUTH_INVALID_TOKEN"
	ErrAuthExpiredToken           ErrCode = "AUTH_EXPIRED_TOKEN"
	ErrAuthTokenNotProvided       ErrCode = "AUTH_TOKEN_NOT_PROVIDED"
	ErrAuthInvalidEmailOrPassword ErrCode = "AUTH_INVALID_EMAIL_OR_PASSWORD"
)

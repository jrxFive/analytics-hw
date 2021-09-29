package types

import "errors"

var (
	ErrInvalidUnit     = errors.New("invalid time unit given")
	ErrMissingKeyAuth  = errors.New("authorization header is missing")
	ErrRequiredSetting = errors.New("required setting missing or invalid")
)

const (
	ContextBitlyAPIKey  = "api-key"
	HeaderAuthorization = "Authorization"
)

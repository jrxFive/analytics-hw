package types

import "errors"

var (
	// ErrInvalidUnit given unit is invalid and not supported by upstream API
	ErrInvalidUnit = errors.New("invalid time unit given")

	// ErrMissingKeyAuth auth key was not specified in request
	ErrMissingKeyAuth = errors.New("authorization header is missing")

	// ErrRequiredSetting required environment setting not set
	ErrRequiredSetting = errors.New("required setting missing or invalid")
)

const (
	// ContextBitlyAPIKey that will hold api key in context to communicate to upstream API with
	ContextBitlyAPIKey = "api-key"
)

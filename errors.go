package circleci

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")

	ErrRequiredEitherIDOrSlug                = errors.New("either organization ID or slug is required")
	ErrRequiredContextID                     = errors.New("context ID is required")
	ErrRequiredContextVariableName           = errors.New("environment variable name is required")
	ErrRequiredContextVariableValue          = errors.New("missing environment variable value")
	ErrRequiredProjectSlug                   = errors.New("project slug is required")
	ErrRequiredProjectCheckoutKeyType        = errors.New("project checkout key type is required")
	ErrRequiredProjectCheckoutKeyFingerprint = errors.New("project checkout key fingerprint is required")
)

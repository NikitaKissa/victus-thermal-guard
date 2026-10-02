package config

import "errors"

var (
	ErrValidation = errors.New("config validation error")
	ErrSyntax     = errors.New("syntax error during parsing config")
)

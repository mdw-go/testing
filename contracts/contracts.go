package contracts

import "errors"

type Func func(actual any, expected ...any) error

var (
	ErrAssertionFailure      = errors.New("assertion failure")
	ErrFatalAssertionFailure = errors.New("fatal")
)

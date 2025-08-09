package assert

import (
	"errors"
	"testing"
)

type Func func(actual any, expected ...any) error

var (
	ErrAssertionFailure      = errors.New("assertion failure")
	ErrFatalAssertionFailure = errors.New("fatal")
)

func So(t *testing.T, actual any, assertion Func, expected ...any) {
	t.Helper()
	_ = New(t).So(actual, assertion, expected...)
}

type T struct{ *testing.T }

func New(t *testing.T) *T {
	return &T{T: t}
}
func (this *T) So(actual any, assertion Func, expected ...any) (ok bool) {
	this.Helper()
	err := assertion(actual, expected...)
	if errors.Is(err, ErrFatalAssertionFailure) {
		this.Fatal(err)
	}
	if err != nil {
		this.Error(err)
	}
	return err == nil
}

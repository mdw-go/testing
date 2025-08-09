package should

import (
	"testing"

	"github.com/mdw-go/testing/suite"
)

type Func func(actual any, expected ...any) error

func So(t *testing.T, actual any, assertion suite.Func, expected ...any) {
	t.Helper()
	_ = suite.New(t).So(actual, assertion, expected...)
}

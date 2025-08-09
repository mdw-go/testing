package should

import (
	"testing"

	"github.com/mdw-go/testing/contracts"
)

type Func func(actual any, expected ...any) error

func So(t *testing.T, actual any, assertion contracts.Func, expected ...any) {
	t.Helper()
	_ = contracts.New(t).So(actual, assertion, expected...)
}

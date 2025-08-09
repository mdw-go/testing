package should

import (
	"testing"

	"github.com/mdw-go/testing/v2/suite"
)

func So(t *testing.T, actual any, assertion suite.Func, expected ...any) {
	t.Helper()
	_ = suite.New(t).So(actual, assertion, expected...)
}

package should

import "testing"

func So(t *testing.T, actual any, assertion Func, expected ...any) {
	t.Helper()
	_ = New(t).So(actual, assertion, expected...)
}

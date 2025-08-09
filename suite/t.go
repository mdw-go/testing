package suite

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mdw-go/testing/assert"
)

type T struct{ *testing.T }

func New(t *testing.T) *T {
	return &T{T: t}
}
func (this *T) So(actual any, assertion assert.Func, expected ...any) (ok bool) {
	this.Helper()
	err := assertion(actual, expected...)
	if errors.Is(err, assert.ErrFatalAssertionFailure) {
		this.Fatal(err)
	}
	if err != nil {
		this.Error(err)
	}
	return err == nil
}
func (this *T) Print(v ...any)            { _, _ = fmt.Fprint(this.Output(), v...) }
func (this *T) Printf(f string, v ...any) { _, _ = fmt.Fprintf(this.Output(), f, v...) }
func (this *T) Println(v ...any)          { _, _ = fmt.Fprintln(this.Output(), v...) }

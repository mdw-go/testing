package suite

import (
	"fmt"
	"testing"

	"github.com/mdw-go/testing/v2/assert"
)

type T struct{ *assert.T }

func New(t *testing.T) *T { return &T{T: assert.New(t)} }

func (this *T) Print(v ...any)            { _, _ = fmt.Fprint(this.Output(), v...) }
func (this *T) Printf(f string, v ...any) { _, _ = fmt.Fprintf(this.Output(), f, v...) }
func (this *T) Println(v ...any)          { _, _ = fmt.Fprintln(this.Output(), v...) }

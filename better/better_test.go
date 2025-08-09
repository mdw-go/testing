package better_test

import (
	"testing"

	"github.com/mdw-go/testing/better"
)

func TestSo(t *testing.T) {
	better.So(t, 1, better.Equal, 1)
}

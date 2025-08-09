package suite_test

import (
	"testing"

	"github.com/mdw-go/testing/should"
	"github.com/mdw-go/testing/suite"
)

func TestSo(t *testing.T) {
	suite.So(t, 1, should.Equal, 1)
	suite.New(t).So(1, should.Equal, 1)
}

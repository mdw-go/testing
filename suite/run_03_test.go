package suite_test

import (
	"testing"

	"github.com/mdw-go/testing/contracts"
	"github.com/mdw-go/testing/should"
	"github.com/mdw-go/testing/suite"
)

func TestSkip(t *testing.T) {
	fixture := &Suite03{T: contracts.New(t)}
	suite.Run(fixture)
	fixture.So(t.Failed(), should.BeFalse)
}

type Suite03 struct{ *contracts.T }

func (this *Suite03) SkipTestThatFails() {
	this.So(1, should.Equal, 2)
}
func (this *Suite03) SkipLongTestThatFails() {
	this.So(1, should.Equal, 2)
}

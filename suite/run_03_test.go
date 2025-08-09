package suite_test

import (
	"testing"

	"github.com/mdw-go/testing/v2/assert"
	"github.com/mdw-go/testing/v2/should"
	"github.com/mdw-go/testing/v2/suite"
)

func TestSkip(t *testing.T) {
	fixture := &Suite03{T: assert.New(t)}
	suite.Run(fixture)
	fixture.So(t.Failed(), should.BeFalse)
}

type Suite03 struct{ *assert.T }

func (this *Suite03) SkipTestThatFails() {
	this.So(1, should.Equal, 2)
}
func (this *Suite03) SkipLongTestThatFails() {
	this.So(1, should.Equal, 2)
}

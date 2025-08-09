package suite_test

import (
	"testing"

	"github.com/mdw-go/testing/v2/should"
	"github.com/mdw-go/testing/v2/suite"
)

func TestFreshFixture(t *testing.T) {
	fixture := &Suite02{T: suite.New(t)}
	suite.Run(fixture, suite.Options.UnitTests())
	fixture.So(fixture.counter, should.Equal, 0)
}

type Suite02 struct {
	*suite.T
	counter int
}

func (this *Suite02) TestSomething() {
	this.Print("*** this should appear in the test log!")
	this.counter++
}

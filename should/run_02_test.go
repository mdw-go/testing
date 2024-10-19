package should_test

import (
	"testing"

	"github.com/mdw-go/testing/should"
)

func TestFreshFixture(t *testing.T) {
	fixture := &Suite02{T: should.New(t)}
	should.Run(fixture, should.Options.UnitTests())
	fixture.So(fixture.counter, should.Equal, 0)
}

type Suite02 struct {
	*should.T
	counter int
}

func (this *Suite02) TestSomething() {
	_, _ = this.Write([]byte("*** this should appear in the test log!"))
	this.counter++
}

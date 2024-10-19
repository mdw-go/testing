package should_test

import (
	"testing"

	"github.com/mdw-go/testing/should"
)

func TestFocus(t *testing.T) {
	fixture := &Suite05{
		T:      should.New(t),
		events: make(map[string]struct{}),
	}

	should.Run(fixture, should.Options.SharedFixture())

	fixture.So(t.Failed(), should.BeFalse)
	if testing.Short() {
		fixture.So(fixture.events, should.Equal, map[string]struct{}{"1": {}})
	} else {
		fixture.So(fixture.events, should.Equal, map[string]struct{}{
			"1": {},
			"2": {},
		})
	}
}

type Suite05 struct {
	*should.T
	events map[string]struct{}
}

func (this *Suite05) FocusTest1() {
	this.events["1"] = struct{}{}
}
func (this *Suite05) FocusLongTest2() {
	this.events["2"] = struct{}{}
}
func (this *Suite05) TestThatFails() {
	this.So(1, should.Equal, 2)
}

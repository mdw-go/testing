package should_test

import (
	"testing"

	"github.com/mdw-go/testing/should"
)

func TestSkip(t *testing.T) {
	fixture := &Suite03{T: should.New(t)}
	should.Run(fixture)
	fixture.So(t.Failed(), should.BeFalse)
}

type Suite03 struct{ *should.T }

func (this *Suite03) SkipTestThatFails() {
	this.So(1, should.Equal, 2)
}
func (this *Suite03) SkipLongTestThatFails() {
	this.So(1, should.Equal, 2)
}

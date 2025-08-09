package suite_test

import (
	"testing"

	"github.com/mdw-go/testing/v2/assert"
	"github.com/mdw-go/testing/v2/should"
	"github.com/mdw-go/testing/v2/suite"
)

func TestLong(t *testing.T) {
	if !testing.Short() {
		t.Skip("This test only to be run in when -test.short flag passed.")
	}
	fixture := &Suite04{T: assert.New(t)}
	suite.Run(fixture)
	fixture.So(t.Failed(), should.BeFalse)
}

type Suite04 struct{ *assert.T }

func (this *Suite04) LongTestThatWouldFailButShouldBeSkippedInShortMode() {
	this.So(1, should.Equal, 2)
}

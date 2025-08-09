package suite_test

import (
	"testing"

	"github.com/mdw-go/testing/contracts"
	"github.com/mdw-go/testing/should"
	"github.com/mdw-go/testing/suite"
)

func TestLong(t *testing.T) {
	if !testing.Short() {
		t.Skip("This test only to be run in when -test.short flag passed.")
	}
	fixture := &Suite04{T: contracts.New(t)}
	suite.Run(fixture)
	fixture.So(t.Failed(), should.BeFalse)
}

type Suite04 struct{ *contracts.T }

func (this *Suite04) LongTestThatWouldFailButShouldBeSkippedInShortMode() {
	this.So(1, should.Equal, 2)
}

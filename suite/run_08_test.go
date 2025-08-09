package suite_test

import (
	"testing"

	"github.com/mdw-go/testing/contracts"
	"github.com/mdw-go/testing/should"
	"github.com/mdw-go/testing/suite"
)

func TestLongRunningSuite(t *testing.T) {
	fixture := &Suite08{T: contracts.New(t)}

	suite.Run(fixture, suite.Options.LongRunning())

	if testing.Short() {
		panic("should have skipped long-running test in -short mode")
	} else {
		fixture.So(fixture.events, should.Equal, []string{
			"SetupSuite",
			"Setup",
			"Test1",
			"Teardown",
			"TeardownSuite",
		})
	}
}

type Suite08 struct {
	*contracts.T
	events []string
}

func (this *Suite08) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite08) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite08) Setup()              { this.record("Setup") }
func (this *Suite08) Teardown()           { this.record("Teardown") }
func (this *Suite08) Test1()              { this.record("Test1") }
func (this *Suite08) record(event string) { this.events = append(this.events, event) }

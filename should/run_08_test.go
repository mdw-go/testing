package should_test

import (
	"testing"

	"github.com/mdw-go/testing/should"
)

func TestLongRunningSuite(t *testing.T) {
	fixture := &Suite08{T: should.New(t)}

	should.Run(fixture, should.Options.LongRunning())

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
	*should.T
	events []string
}

func (this *Suite08) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite08) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite08) Setup()              { this.record("Setup") }
func (this *Suite08) Teardown()           { this.record("Teardown") }
func (this *Suite08) Test1()              { this.record("Test1") }
func (this *Suite08) record(event string) { this.events = append(this.events, event) }

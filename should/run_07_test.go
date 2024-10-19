package should_test

import (
	"testing"

	"github.com/mdw-go/testing/should"
)

func TestSuiteWithSkippedTests(t *testing.T) {
	fixture := &Suite07{T: should.New(t)}

	should.Run(fixture, should.Options.SharedFixture())

	fixture.So(fixture.events, should.Equal, []string{
		"SetupSuite",
		"Setup",
		"Test1",
		"Teardown",
		"TeardownSuite",
	})
}

type Suite07 struct {
	*should.T
	events []string
}

func (this *Suite07) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite07) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite07) Setup()              { this.record("Setup") }
func (this *Suite07) Teardown()           { this.record("Teardown") }
func (this *Suite07) Test1()              { this.record("Test1") }
func (this *Suite07) SkipTest2()          { this.record("SkipTest2") }
func (this *Suite07) record(event string) { this.events = append(this.events, event) }

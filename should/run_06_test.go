package should_test

import (
	"testing"

	"github.com/mdw-go/testing/should"
)

func TestSuiteWithSetupsAndTeardownsSkippedEntirelyIfAllTestsSkipped(t *testing.T) {
	fixture := &Suite06{T: should.New(t)}

	should.Run(fixture, should.Options.SharedFixture())

	fixture.So(fixture.events, should.BeNil)
}

type Suite06 struct {
	*should.T
	events []string
}

func (this *Suite06) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite06) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite06) Setup()              { this.record("Setup") }
func (this *Suite06) Teardown()           { this.record("Teardown") }
func (this *Suite06) SkipTest()           { this.record("SkipTest") }
func (this *Suite06) record(event string) { this.events = append(this.events, event) }

package should_test

import (
	"testing"

	"github.com/mdw-go/testing/should"
)

func TestSuiteWithSetupsAndTeardowns(t *testing.T) {
	fixture := &Suite01{T: should.New(t)}

	should.Run(fixture, should.Options.IntegrationTests())

	fixture.So(fixture.events, should.Equal, []string{
		"SetupSuite",
		"Setup",
		"Test",
		"Teardown",
		"TeardownSuite",
	})
}

type Suite01 struct {
	*should.T
	events []string
}

func (this *Suite01) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite01) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite01) Setup()              { this.record("Setup") }
func (this *Suite01) Teardown()           { this.record("Teardown") }
func (this *Suite01) Test()               { this.record("Test") }
func (this *Suite01) record(event string) { this.events = append(this.events, event) }

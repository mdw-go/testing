package suite_test

import (
	"testing"

	"github.com/mdw-go/testing/contracts"
	"github.com/mdw-go/testing/should"
	"github.com/mdw-go/testing/suite"
)

func TestSuiteWithSetupsAndTeardownsSkippedEntirelyIfAllTestsSkipped(t *testing.T) {
	fixture := &Suite06{T: contracts.New(t)}

	suite.Run(fixture, suite.Options.SharedFixture())

	fixture.So(fixture.events, should.BeNil)
}

type Suite06 struct {
	*contracts.T
	events []string
}

func (this *Suite06) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite06) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite06) Setup()              { this.record("Setup") }
func (this *Suite06) Teardown()           { this.record("Teardown") }
func (this *Suite06) SkipTest()           { this.record("SkipTest") }
func (this *Suite06) record(event string) { this.events = append(this.events, event) }

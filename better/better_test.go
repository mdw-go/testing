package better_test

import (
	"testing"

	"github.com/mdw-go/testing/v2/assert"
	"github.com/mdw-go/testing/v2/better"
	"github.com/mdw-go/testing/v2/should"
)

func TestWrapFatalSuccess(t *testing.T) {
	err := better.Equal(1, 1)
	assert.So(t, err, should.BeNil)
}
func TestWrapFatalFailure(t *testing.T) {
	err := better.Equal(1, 2)
	assert.So(t, err, should.WrapError, assert.ErrFatalAssertionFailure)
	assert.So(t, err, should.WrapError, assert.ErrAssertionFailure)
}
func TestWrapFatalSuccess_NOT(t *testing.T) {
	err := better.NOT.Equal(1, 2)
	assert.So(t, err, should.BeNil)
}
func TestWrapFatalFailure_NOT(t *testing.T) {
	err := better.NOT.Equal(1, 1)
	assert.So(t, err, should.WrapError, assert.ErrFatalAssertionFailure)
	assert.So(t, err, should.WrapError, assert.ErrAssertionFailure)
}

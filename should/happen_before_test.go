package should_test

import (
	"testing"
	"time"

	"github.com/mdw-go/testing/v2/should"
)

func TestShouldHappenBefore(t *testing.T) {
	assert := NewAssertion(t)

	now := time.Now()
	assert.TypeMismatch(1, should.HappenBefore, now)
	assert.TypeMismatch(now, should.HappenBefore, 1)

	assert.ExpectedCountInvalid(now, should.HappenBefore)
	assert.ExpectedCountInvalid(now, should.HappenBefore, now, now)

	assert.Fail(now.Add(time.Nanosecond), should.HappenBefore, now)
	assert.Fail(now, should.HappenBefore, now)
	assert.Pass(now, should.HappenBefore, now.Add(time.Nanosecond))
}

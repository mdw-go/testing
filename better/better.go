package better

import (
	"fmt"

	"github.com/mdw-go/testing/v2/assert"
	"github.com/mdw-go/testing/v2/should"
)

func wrap(assertion assert.Func) assert.Func {
	return func(actual any, expected ...any) error {
		err := assertion(actual, expected...)
		if err != nil {
			return fmt.Errorf("%w %w", assert.ErrFatalAssertionFailure, err)
		}
		return nil
	}
}

var (
	BeChronological        = wrap(should.BeChronological)
	BeEmpty                = wrap(should.BeEmpty)
	BeFalse                = wrap(should.BeFalse)
	BeGreaterThan          = wrap(should.BeGreaterThan)
	BeGreaterThanOrEqualTo = wrap(should.BeGreaterThanOrEqualTo)
	BeIn                   = wrap(should.BeIn)
	BeLessThan             = wrap(should.BeLessThan)
	BeLessThanOrEqualTo    = wrap(should.BeLessThanOrEqualTo)
	BeNil                  = wrap(should.BeNil)
	BeTrue                 = wrap(should.BeTrue)
	Contain                = wrap(should.Contain)
	EndWith                = wrap(should.EndWith)
	Equal                  = wrap(should.Equal)
	HappenAfter            = wrap(should.HappenAfter)
	HappenBefore           = wrap(should.HappenBefore)
	HappenOn               = wrap(should.HappenOn)
	HappenWithin           = wrap(should.HappenWithin)
	HaveLength             = wrap(should.HaveLength)
	Panic                  = wrap(should.Panic)
	StartWith              = wrap(should.StartWith)
	WrapError              = wrap(should.WrapError)
)

// NOT (a singleton) constrains all negated assertions to their own namespace.
var NOT negated

type negated struct{}

func (negated) BeChronological(actual any, expected ...any) error {
	return wrap(should.NOT.BeChronological)(actual, expected...)
}
func (negated) BeEmpty(actual any, expected ...any) error {
	return wrap(should.NOT.BeEmpty)(actual, expected...)
}
func (negated) BeGreaterThan(actual any, expected ...any) error {
	return wrap(should.NOT.BeGreaterThan)(actual, expected...)
}
func (negated) BeGreaterThanOrEqualTo(actual any, expected ...any) error {
	return wrap(should.NOT.BeGreaterThanOrEqualTo)(actual, expected...)
}
func (negated) BeIn(actual any, expected ...any) error {
	return wrap(should.NOT.BeIn)(actual, expected...)
}
func (negated) BeLessThan(actual any, expected ...any) error {
	return wrap(should.NOT.BeLessThan)(actual, expected...)
}
func (negated) BeLessThanOrEqualTo(actual any, expected ...any) error {
	return wrap(should.NOT.BeLessThanOrEqualTo)(actual, expected...)
}
func (negated) BeNil(actual any, expected ...any) error {
	return wrap(should.NOT.BeNil)(actual, expected...)
}
func (negated) Contain(actual any, expected ...any) error {
	return wrap(should.NOT.Contain)(actual, expected...)
}
func (negated) Equal(actual any, expected ...any) error {
	return wrap(should.NOT.Equal)(actual, expected...)
}
func (negated) HappenOn(actual any, expected ...any) error {
	return wrap(should.NOT.HappenOn)(actual, expected...)
}
func (negated) Panic(actual any, expected ...any) error {
	return wrap(should.NOT.Panic)(actual, expected...)
}

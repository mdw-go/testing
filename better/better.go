/*
Package better

This package strives to make it easy, even fun, for software developers to produce
> a quick, sure, and repeatable proof that every element of the code works as it should.
(See [The Programmer's Oath](http://blog.cleancoder.com/uncle-bob/2015/11/18/TheProgrammersOath.html))

The simplest way is by combining the So function with the many provided assertions, such as better.Equal:

	package whatever

	import (
		"log"
		"testing"

		"github.com/mdw-go/testing/better"
	)

	func Test(t *testing.T) {
		better.So(t, 1, better.Equal, 1)
	}
*/
package better

import (
	"fmt"
	"testing"

	"github.com/mdw-go/testing/should"
	"github.com/mdw-go/testing/suite"
)

func So(t *testing.T, actual any, assertion suite.Func, expected ...any) {
	t.Helper()
	_ = suite.New(t).So(actual, assertion, expected...)
}

func wrap(assertion suite.Func) suite.Func {
	return func(actual any, expected ...any) error {
		err := assertion(actual, expected...)
		if err != nil {
			err = fmt.Errorf("%w %w", suite.ErrFatalAssertionFailure, err)
		}
		return err
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

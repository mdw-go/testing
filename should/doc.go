/*
Package should

This package strives to make it easy, even fun, for
software developers to produce
> a quick, sure, and repeatable proof that every element of the code works as it should.
(See [The Programmer's Oath](http://blog.cleancoder.com/uncle-bob/2015/11/18/TheProgrammersOath.html))

The simplest way is by combining the So function with the many provided assertions, such as should.Equal:

	package whatever

	import (
		"log"
		"testing"

		"github.com/mdw-go/testing/v2/should"
	)

	func Test(t *testing.T) {
		should.So(t, 1, should.Equal, 1)
	}
*/
package should

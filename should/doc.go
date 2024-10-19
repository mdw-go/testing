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

		"github.com/mdw-go/testing/should"
	)

	func Test(t *testing.T) {
		should.So(t, 1, should.Equal, 1)
	}

This package also implement an xUnit-style test
runner, which is based on the following packages:

  - [github.com/stretchr/testify/suite](https://pkg.go.dev/github.com/stretchr/testify/suite)
  - [github.com/smartystreets/gunit](https://pkg.go.dev/github.com/smartystreets/gunit)

For those using an IDE by JetBrains, you may
find the following "live template" helpful:

	func Test$NAME$Suite(t *testing.T) {
		should.Run(&$NAME$Suite{T: should.New(t)}, should.Options.UnitTests())
	}

	type $NAME$Suite struct {
		*should.T
	}

	func (this *$NAME$Suite) Setup() {
	}

	func (this *$NAME$Suite) Test$END$() {
	}

From a test method like the one in the template above, simply use the embedded So method:

	func (this TheSuite) TestSomething() {
		this.So(1, should.Equal, 1)
	}

Happy testing!
*/
package should

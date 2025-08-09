/*
Package suite

This package strives to make it easy, even fun, for
software developers to produce
> a quick, sure, and repeatable proof that every element of the code works as it should.
(See [The Programmer's Oath](http://blog.cleancoder.com/uncle-bob/2015/11/18/TheProgrammersOath.html))

This package implements an xUnit-style test runner, which is based on the following packages:

  - [github.com/stretchr/testify/suite](https://pkg.go.dev/github.com/stretchr/testify/suite)
  - [github.com/smartystreets/gunit](https://pkg.go.dev/github.com/smartystreets/gunit)

For those using an IDE by JetBrains, you may
find the following "live template" helpful:

	func Test$NAME$Suite(t *testing.T) {
		suite.Run(&$NAME$Suite{T: suite.New(t)}, suite.Options.UnitTests())
	}

	type $NAME$Suite struct {
		*suite.T
	}

	func (this *$NAME$Suite) Setup() {
	}

	func (this *$NAME$Suite) Test$END$() {
	}

From a test method like the one in the template above, simply use the embedded So method:

	func (this *TheSuite) TestSomething() {
		this.So(1, should.Equal, 1)
	}

Happy testing!
*/
package suite

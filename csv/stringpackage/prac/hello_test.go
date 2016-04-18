package hello_test

import (
    "testing"
   // "io"

    . "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHelloWorld(c *C) {
    c.Assert(42, Equals, "42")
    //c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
    c.Check(42, Equals, 42)
}

// import (
//   "testing"
//   "github.com/stretchr/testify/assert"
// )

// func TestSomething(t *testing.T) {

//   var a string = "Hello"
//   var b string = "Hello"

//   assert.Equal(t, a, b, "The two words should be the same.")

// }
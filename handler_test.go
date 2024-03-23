package lab2

import (
	"bytes"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestCorrectInput(c *C) {
	expected := "20" // Очікуваний результат обчислення для "5 4 *"
	output := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader("* 5 4"),
		Output: output,
	}
	err := handler.Compute()

	c.Assert(err, IsNil)
	c.Assert(output.String(), Equals, expected)
}

func (s *HandlerSuite) TestIncorrectInput(c *C) {
	output := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader("5 4 ||"),
		Output: output,
	}
	err := handler.Compute()

	c.Assert(err, NotNil)
}

func (s *HandlerSuite) TestSyntaxError(c *C) {
	output := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader("5 4 * *"), // Синтаксична помилка - зайва операція
		Output: output,
	}
	err := handler.Compute()

	c.Assert(err, NotNil)
}

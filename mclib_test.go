package mclib_test

import (
	"os"
	"testing"

	"github.com/bep/mclib"
	qt "github.com/frankban/quicktest"
)

func TestRunMainOK(t *testing.T) {
	c := qt.New(t)
	os.Args = []string{"-help"}
	c.Assert(mclib.RunMain(), qt.IsNil)

}

func TestRunMainError(t *testing.T) {
	c := qt.New(t)
	os.Args = []string{"-install", "-CAROOT", "asdfasdfasdfasdf"}
	err := mclib.RunMain()
	c.Assert(err, qt.IsNotNil)
	c.Assert(err.Error(), qt.Equals, "you can't set -[un]install and -CAROOT at the same time")

}

package mclib_test

import (
	"os"
	"testing"

	"github.com/bep/mclib"
	qt "github.com/frankban/quicktest"
)

func TestRunMain(t *testing.T) {
	c := qt.New(t)
	os.Args = []string{"-help"}
	c.Assert(mclib.RunMain(), qt.IsNil)

}

func TestGetCAROOT(t *testing.T) {
	c := qt.New(t)
	c.Assert(mclib.GetCAROOT(), qt.Not(qt.Equals), "")
}

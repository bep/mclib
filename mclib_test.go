package mclib_test

import (
	"os"
	"testing"

	"github.com/bep/mclib"
)

func TestRunMain(t *testing.T) {
	os.Args = []string{"-help"}
	mclib.RunMain()
}

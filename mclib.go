package mclib

import (
	"os"

	"github.com/bep/mclib/internal"
)

func RunMain() {
	const mkcert = "mkcert"
	if len(os.Args) == 0 || os.Args[0] != mkcert {
		// All commands needs the "mkcert" as its first argument.
		os.Args = append([]string{mkcert}, os.Args...)
	}
	internal.RunMain()
}

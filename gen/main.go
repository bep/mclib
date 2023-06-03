//go:generate go run main.go

package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bep/helpers/filehelpers"
)

func main() {
	rootDir := "../"
	if err := os.MkdirAll(filepath.Join(rootDir, "internal"), 0755); err != nil {
		log.Fatal(err)
	}

	filehelpers.CopyDir(filepath.Join(rootDir, "mkcert"), filepath.Join(rootDir, "internal"), func(path string) bool {
		return filepath.Ext(path) == ".go" && !strings.HasSuffix(path, "_test.go")
	})

	err := filepath.Walk(filepath.Join(rootDir, "internal"), func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		s := string(b)

		s = strings.Replace(s, "package main", "package internal", 1)
		if strings.HasSuffix(path, "main.go") {
			s = strings.Replace(s, "func main()", "func RunMain()", 1)
		}

		// Write to the same file.
		if err := ioutil.WriteFile(path, []byte(s), info.Mode()); err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		log.Fatal(err)
	}

}

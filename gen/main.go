//go:generate go run main.go

package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bep/helpers/filehelpers"
)

func main() {
	rootDir := "../"
	internalDir := filepath.Join(rootDir, "internal")
	// Remove old files.
	if err := os.RemoveAll(internalDir); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(internalDir, 0755); err != nil {
		log.Fatal(err)
	}

	if err := filehelpers.CopyDir(filepath.Join(rootDir, "mkcert"), internalDir, func(path string) bool {
		return filepath.Ext(path) == ".go" && !strings.HasSuffix(path, "_test.go")
	}); err != nil {
		log.Fatal(err)
	}

	if err := filehelpers.CopyFile("./truststore_other.go.txt", filepath.Join(internalDir, "truststore_other.go")); err != nil {
		log.Fatal(err)
	}

	fileReplacer := strings.NewReplacer(
		"getCAROOT()", "GetCAROOT()",
		"mkcert -install", "hugo server trust",
		"mkcert -uninstall", "hugo server trust -uninstall",
		" 👈", "",
		" 👍", "",
		"  🦊", "",
		" ℹ️", "",
		" ⚠️", "",
		" 👋", "",
	)

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

		s = fileReplacer.Replace(s)

		// We don't want os.Exit(-1) in a library.
		// E.g. log.Fatalf("ERROR: failed to execute \"%s\": %s\n\n%s\n", cmd, err, out)
		// Replace with panic.
		// NOTE: These are currently not perfect, so some edits may be needed (missing imports).
		fatalFRe := regexp.MustCompile(`log.Fatalf\((.*)\)`)
		s = fatalFRe.ReplaceAllString(s, "panic(fmt.Sprintf($1))")

		fatalLnRe := regexp.MustCompile(`log.Fatalln\((.*)\)`)
		s = fatalLnRe.ReplaceAllString(s, "panic(fmt.Sprintln($1))")

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

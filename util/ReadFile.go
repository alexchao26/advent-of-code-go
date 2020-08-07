package util

import (
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"strings"
)

/*
ReadFile is a wrapper over io/ioutil.ReadFile but also determines the
dynamic absolute path to the file.
*/
func ReadFile(pathFromCaller string) string {
	// Docs: https://golang.org/pkg/runtime/#Caller
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		// NOTE this could be updated to make ReadFile return an error, but that's overkill...
		log.Fatal("Could not find Caller of util.ReadFile")
	}

	// parse directory with pathFromCaller (which could be relative to Directory)
	absolutePath := path.Join(path.Dir(filename), pathFromCaller)

	// read the entire file & return the byte slice as a string
	content, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		log.Fatal(err)
	}
	// trim off new lines and tabs at end of input files
	strContent := string(content)
	return strings.TrimRight(strContent, "\n")
}

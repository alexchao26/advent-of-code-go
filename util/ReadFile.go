package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

/*
ReadFile takes the relative path from *the caller* and returns the contents
of the file as a string
*/
func ReadFile(pathFromCaller string) string {
	// Docs: https://golang.org/pkg/runtime/#Caller
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		// NOTE this could be updated to make ReadFile return an error, but that's overkill...
		log.Fatal("Could not find Caller of util.ReadFile")
	}

	absolutePath := path.Join(path.Dir(filename), pathFromCaller)

	fmt.Println("abs path is", absolutePath)

	content, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

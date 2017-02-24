package main

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Use `go get` to download addon and add to $GOPATH/src, useful
// for IDE auto-import and code completion, then copy entire directory
// tree to project's ./addons folder
func getAddon(args []string) error {

	// error return
	errorFunc := func(err error) error {
		return errors.New("Ponzu add failed. " + "\n" + err.Error())
	}

	var cmdOptions []string
	var addonPath = args[1]

	// Go get
	cmdOptions = append(cmdOptions, addonPath)
	get := exec.Command(gocmd, cmdOptions...)
	get.Stderr = os.Stderr
	get.Stdout = os.Stdout

	err := get.Start()
	if err != nil {
		errorFunc(err)
	}
	err = get.Wait()
	if err != nil {
		errorFunc(err)
	}

	// Copy to ./addons folder
	// GOPATH can be a list delimited by ":" on Linux or ";" on Windows
	// `go get` uses the first, this should parse out the first whatever the OS
	envGOPATH := os.Getenv("GOPATH")
	gopaths := strings.Split(envGOPATH, ":")
	gopath := gopaths[0]
	gopaths = strings.Split(envGOPATH, ";")
	gopath = gopaths[0]

	src := filepath.Join(gopath, addonPath)
	dest := filepath.Join("./addons", addonPath)

	err = copyAll(src, dest)
	if err != nil {
		errorFunc(err)
	}
	return nil
}

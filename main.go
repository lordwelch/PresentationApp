// PresentationApp project main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kardianos/osext"
	"gopkg.in/qml.v1"
)

var (
	path string
)

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()
	path, err := osext.ExecutableFolder()
	path = filepath.Clean(path + "/../src/github.com/lordwelch/PresentationApp/")
	fmt.Println(path)
	mainQml, err := engine.LoadFile(path + "/main.qml")
	if err != nil {
		return err
	}

	cellQml, err := engine.LoadFile(path + "/cell.qml")
	if err != nil {
		return err
	}

	window := mainQml.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}

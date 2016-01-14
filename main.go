// PresentationApp project main.go
package main

import (
	"fmt"
	"os"

	"github.com/kardianos/osext"
	"gopkg.in/qml.v1"
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
	controls, err := engine.LoadFile(path + "/../share/main.qml")
	if err != nil {
		return err
	}

	window := controls.CreateWindow(nil)
	window.ObjectByName("data1")

	window.Show()
	window.Wait()
	return nil
}

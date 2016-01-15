// PresentationApp project main.go
package main

import (
	"fmt"
	//"image"
	"os"
	"path/filepath"

	"github.com/kardianos/osext"
	"gopkg.in/qml.v1"
)

type cell struct {
	text string
	//img   image.Image
	qmlcell qml.Object
	index   int
}

var (
	path    string
	cellQml qml.Object
	window  *qml.Window
	slides  []cell
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

	window = mainQml.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}

func (cl *cell) addCell() {
	cl.index = len(slides)
	cl.qmlcell = cellQml.Create(nil)
	fmt.Println(cl.qmlcell.ObjectByName("celltext").Property("text"))
	cl.text = "testing 1... 2... 3..."
	fmt.Println(cl)
	slides = append(slides, *cl)
	slides[cl.index].qmlcell.Set("parent", window.ObjectByName("data1"))

}

func (cl *cell) String() string {
	return fmt.Sprint("Index: %T \nText:  %T", cl.index, cl.text)
}

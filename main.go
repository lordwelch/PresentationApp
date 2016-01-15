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
type slide []cell

var (
	path    string
	cellQml qml.Object
	window  *qml.Window
	slides  slide
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
	fmt.Println("ignore", cellQml)
	qml.RunMain(slides.addCell())

	window = mainQml.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}

func (sl *slide) addCell( /*cl *cell*/ ) {
	var cl cell
	cl.index = len(*sl)
	fmt.Println("index: ", cl.index)
	cl.qmlcell = cellQml.Create(nil)
	fmt.Println("index: ", cl.index)
	fmt.Println(cl.qmlcell.ObjectByName("celltext").Property("text"))
	cl.text = "testing 1... 2... 3..."
	fmt.Println(cl)
	//*sl = append(*sl, *cl)
	//*sl[0] //.qmlcell.Set("parent", window.ObjectByName("data1"))

}

func (cl *cell) String() string {
	return fmt.Sprint("Index: %T \nText:  %T", cl.index, cl.text)
}

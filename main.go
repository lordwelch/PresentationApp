// PresentationApp project main.go
package main

import (
	"fmt"
	//"image"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-gl/glfw/v3.1/glfw"
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
	path     string
	textEdit qml.Object
	cellQml  qml.Object
	window   *qml.Window
	slides   slide
	err      error
	monitors []*glfw.Monitor
)

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	checkMon()

	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	glfw.Terminate()
}

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func run() error {
	engine := qml.NewEngine()
	path, err = osext.ExecutableFolder()
	path = filepath.Clean(path + "/../src/github.com/lordwelch/PresentationApp/")
	//fmt.Println(path)
	mainQml, err := engine.LoadFile(path + "/main.qml")
	if err != nil {
		return err
	}

	/*projQml, err := engine.LoadFile(path + "/qml/projection.qml")
	if err != nil {
		return err
	}*/

	cellQml, err = engine.LoadFile(path + "/qml/cell.qml")
	if err != nil {
		return err
	}

	window = mainQml.CreateWindow(nil)
	//created(projQml)
	textEdit = window.ObjectByName("textEdit")
	slides.addCell()

	window.Show()
	window.Wait()
	return nil
}

func checkMon() {
	monitors = glfw.GetMonitors()
	if i := len(monitors); i < 2 {
		fmt.Println("You only have 1 monitor!!!!!!!!!!! :-P")
	} else {
		fmt.Printf("You have %d monitors\n", i)
	}
	monitorInfo()
}

func monitorInfo() {
	for _, mon := range monitors {
		fmt.Printf("monitor name: %s\n", mon.GetName())
		i, t := mon.GetPos()
		fmt.Printf("position X: %d  Y: %d\n", i, t)
		//fmt.Println(mon.)
	}
}

func (cl *cell) setSignals() {
	cl.qmlcell.ObjectByName("cellMouse").On("doubleClicked", func() {
		cellText := cl.qmlcell.ObjectByName("cellText")
		textEdit.Set("cell", cl.index)
		textEdit.Set("x", cellText.Int("x")+4)
		textEdit.Set("y", cellText.Int("y")+4)
		textEdit.Set("width", cellText.Int("width"))
		textEdit.Set("height", cellText.Int("height"))
		textEdit.Set("opacity", 100)
		textEdit.Set("visible", true)
		textEdit.ObjectByName("textEdit1").Set("focus", true)
		textEdit.Set("enabled", true) /*
			fmt.Println(textEdit.Int("x"))
			fmt.Println(textEdit.Int("y"))
			fmt.Println(textEdit.Int("width"))
			fmt.Println(textEdit.Int("height"))*/
	})
	textEdit.ObjectByName("textEdit1").On("focusChanged", func(focus bool) {
		var (
			str string
			cel cell
		)
		//fmt.Printf("focusChanged: focus: %t\n", focus)
		if !focus {
			str = textEdit.ObjectByName("textEdit1").String("text")
			cel = slides[textEdit.Int("cell")]
			cel.qmlcell.ObjectByName("cellText").Set("text", str)
		}
	})
}

func (sl *slide) addCell( /*cl *cell*/ ) {
	var cl cell

	cl.qmlcell = cellQml.Create(nil)
	cl.qmlcell.Set("objectName", fmt.Sprintf("cellRect%d", cl.index))
	cl.qmlcell.Set("parent", window.ObjectByName("data1"))

	cl.index = len(*sl)

	cl.text = "testing 1... 2... 3..."
	cl.qmlcell.ObjectByName("cellText").Set("text", cl.text)
	*sl = append(*sl, cl)

	cl.setSignals()

	//fmt.Print((*sl)[cl.index])
}

func (cl cell) String() string {
	return fmt.Sprintf("Index: %d \nText:  %s\n", cl.index, cl.text)
}

func (cl cell) GoString() string {
	return cl.String()
}

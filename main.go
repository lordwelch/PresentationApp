// PresentationApp project main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kardianos/osext"
	"github.com/lordwelch/qml"
	"gopkg.in/gographics/imagick.v2/imagick"
)

//Bool type i'm lazy wanted a toggle function
type Bool bool

type cell struct {
	text    string
	img     *imagick.MagickWand
	qmlimg  qml.Object
	qmlcell qml.Object
	index   int
}
type slide []*cell

var (
	path   string
	slides slide
	err    error
)

func main() {

	if err = qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

}

func run() error {
	imagick.Initialize()

	engine := qml.NewEngine()
	engine.AddImageProvider("images", imgProvider)
	//path for qml files TODO: change to somewhere else
	path, err = osext.ExecutableFolder()
	path = filepath.Clean(path + "/../src/github.com/lordwelch/PresentationApp/")

	mainQml, err = engine.LoadFile(path + "/main.qml")
	if err != nil {
		return err
	}

	cellQml, err = engine.LoadFile(path + "/qml/cell.qml")
	if err != nil {
		return err
	}

	qimg, err = engine.LoadFile(path + "/qml/img.qml")
	if err != nil {
		return err
	}

	window = mainQml.CreateWindow(nil)

	textEdit = window.ObjectByName("textEdit")
	//signals for whole qml
	setSignals()
	slides.add()

	//image is ready for imageprovider
	imgready = true

	window.Show()
	slides[0].clearcache()
	qml.RunMain(glInit)

	window.Wait()

	imagick.Terminate()
	return nil
}

//Adds a new cell
func (sl *slide) add( /*cl *cell*/ ) {
	var cl cell
	//gets the length so that the index is valid
	cl.index = len(*sl)

	//increase count on parent QML element
	window.ObjectByName("gridRect").Set("count", window.ObjectByName("gridRect").Int("count")+1)
	cl.qmlcell = cellQml.Create(nil)
	cl.qmlcell.Set("objectName", fmt.Sprintf("cellRect%d", len(*sl)))
	cl.qmlcell.Set("parent", window.ObjectByName("data1"))
	cl.qmlcell.Set("index", cl.index)

	//load image
	cl.img = imagick.NewMagickWand()
	cl.img.ReadImage("logo:")

	//give QML the text
	cl.qmlcell.ObjectByName("cellText").Set("text", cl.text)

	//keep the pointer/dereference (i'm not sure which it is)
	//problems occur otherwise
	*sl = append(*sl, &cl)
	cl.setSignal()

	//seperate image object in QML
	cl.qmlimg = qimg.Create(nil)
	cl.qmlimg.Set("objectName", fmt.Sprintf("cellImg%d", cl.index))
	cl.qmlimg.Set("source", fmt.Sprintf("image://images/%d"+`;`+"0", cl.index))
	cl.qmlimg.Set("parent", window.ObjectByName("data2"))
	cl.qmlimg.Set("index", cl.index)

}

//(cell) remove() should destroy everything
func (cl *cell) remove() {
	cl.text = ""
	cl.qmlimg.Destroy()
	cl.qmlcell.Destroy()
	cl.img.Destroy()
	window.ObjectByName("gridRect").Set("count", window.ObjectByName("gridRect").Int("count")-1)
	slides.remove(cl.index)
	cl.index = -1

}

//(slide) remove copied from github.com/golang/go/wiki/SliceTricks
func (sl *slide) remove(i int) {
	*sl, (*sl)[len((*sl))-1] = append((*sl)[:i], (*sl)[i+1:]...), nil
}

//Toggle lazy wanted a func for it
func (bl *Bool) Toggle() {
	if *bl == false {
		*bl = true
	} else {
		*bl = false
	}
}

//not really needed
func (cl cell) String() string {
	return fmt.Sprintf("Index: %d \nText:  %s\n", cl.index, cl.text)
}

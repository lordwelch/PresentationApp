// PresentationApp project main.go
package main

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"github.com/go-gl/glfw/v3.1/glfw"
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
	font    struct {
		name                    string
		outlineSize, size, x, y float64
		color                   color.RGBA
		outlineColor            color.RGBA
		outline                 Bool
	}
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
	win.Destroy()
	glfw.PollEvents()
	glfw.Terminate()

}

func run() error {
	imagick.Initialize()
	findfonts()

	engine = qml.NewEngine()
	engine.AddImageProvider("images", imgProvider)
	//path for qml files TODO: change to somewhere else
	path, err = osext.ExecutableFolder()
	path = filepath.Clean(path + "/../src/github.com/lordwelch/PresentationApp/")

	mainQml, err = engine.LoadFile(path + "/main.qml")
	if err != nil {
		return err
	}

	edtQml, err = engine.LoadFile(path + "/qml/songEdit.qml")
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

	window = mainQml.CreateWindow(engine.Context())
	window2 := edtQml.CreateWindow(engine.Context())

	textEdit = window.ObjectByName("textEdit")
	//signals for whole qml
	setSignals()
	slides.add()

	//image is ready for imageprovider
	imgready = true

	window.Show()
	window2.Show()
    edtQmlShow()
	slides[0].clearcache()
	qml.RunMain(glInit)

	window.Wait()

	imagick.Terminate()
	return nil
}

//Adds a new cell
func (sl *slide) add( /*cl *cell*/ ) {
	var cl cell
	cl.Init()
	//gets the length so that the index is valid
	cl.index = len(*sl)

	//increase count on parent QML element
	window.ObjectByName("gridRect").Set("count", window.ObjectByName("gridRect").Int("count")+1)
	cl.qmlcell = cellQml.Create(engine.Context())
	cl.qmlcell.Set("objectName", fmt.Sprintf("cellRect%d", len(*sl)))
	cl.qmlcell.Set("parent", window.ObjectByName("data1"))
	cl.qmlcell.Set("index", cl.index)

	//keep the pointer/dereference (i'm not sure which it is)
	//problems occur otherwise
	*sl = append(*sl, &cl)

	//seperate image object in QML
	cl.qmlimg.Set("objectName", fmt.Sprintf("cellImg%d", cl.index))
	cl.qmlimg.Set("source", fmt.Sprintf("image://images/%d"+`;`+"0", cl.index))
	cl.qmlimg.Set("parent", window.ObjectByName("data2"))
	cl.qmlimg.Set("index", cl.index)
	cl.setSignal()
	//give QML the text
	cl.qmlcell.ObjectByName("cellText").Set("text", cl.text)

}

func (cl *cell) Init() {
	cl.text = "hello this is text\nhaha\nhdsjfklfhaskjd"
	cl.index = -1
	cl.font.color, cl.font.outlineColor = color.RGBA{0, 0, 0, 1}, color.RGBA{1, 1, 1, 1}
	cl.font.name = "none"
	cl.font.outline = false
	cl.font.outlineSize = 1
	cl.font.size = 35
	cl.font.x, cl.font.y = 10, 30

	cl.qmlcell = cellQml.Create(engine.Context())
	cl.qmlimg = qimg.Create(engine.Context())

	//load image
	cl.img = imagick.NewMagickWand()
	cl.img.ReadImage("logo:")

}

//(cell) remove() should destroy everything for this cell
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

//Toggle, lazy wanted a func for it
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

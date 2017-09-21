// PresentationApp project main.go
//go:generate genqrc qml

package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/go-gl/glfw/v3.1/glfw"

	"gopkg.in/gographics/imagick.v2/imagick"
	"gopkg.in/qml.v1"
)

type Bool bool

type Cell struct {
	fnt                    Font
	image                  Image
	index, collectionIndex int
	qmlObject              qml.Object
	text                   string
	textVisible            Bool
}

type collection []*Cell

type Font struct {
	color                   color.RGBA
	name                    string
	outline                 Bool
	outlineColor            color.RGBA
	outlineSize, size, x, y float64
}

type Image struct {
	img       *imagick.MagickWand
	imgSource string
	qmlImage  qml.Object
}

type qmlVar struct {
	FontList   string
	Verses     string
	VerseOrder string
	//Img        string
}

type service []collection

var (
	currentService service
	err            error
	path           string
	slides         collection
)

func main() {

	if err = qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	glfw.Terminate()

}

func run() error {
	engine = qml.NewEngine()
	QML = &qmlVar{}
	path = "qrc:///qml"
	imagick.Initialize()
	findFonts()

	engine.Context().SetVar("go", QML)
	engine.AddImageProvider("images", imgProvider)

	err = qmlWindows()
	if err != nil {
		return err
	}

	currentService.Init(1)

	//signals for whole qml
	setSignals()

	//image is ready for imageprovider
	imgready = true

	displayImg = DisplayWindow.Root().ObjectByName("image1")
	serviceObject = serviceQml.Create(engine.Context())
	serviceObject.Set("parent", MainWindow.ObjectByName("data1"))
	serviceObject.Call("addCollection")

	//edtQmlShow()
	qml.RunMain(glInit)
	MainWindow.Wait()
	slides.destroy()

	imagick.Terminate()
	return nil
}

func (sv service) Init(uint num) {
	if num == 0 {
		num = 1
	}

	for index := 0; index < num; index++ {
		if sv == nil {
			sv.add("")
		}
	}
}

func (sv service) add(string name) {
	var (
		sl  collection
		int i = len(sv)
	)

	if len(name) <= 0 {
		name = "Song: " + fmt.Sprint(i)
	}

	sl.init()
	sv = append(sv, sl)
	//?serviceObj.Call(addCollection, name, 1)
}

func (sv service) remove(i int) {
	sv[i].destroy()

	copy(sv[i:], sv[i+1:])
	sv[len(sv)-1] = nil // or the zero value of T
	sv = sv[:len(sv)-1]

}

func (sv service) destroy() {
	for i := len(sv); i > 0; i-- {
		sv.remove(i - 1)
	}
}

func (sl collection) init(int num) {
	if num <= 0 {
		num = 1
	}

	for index := 0; index < num; index++ {
		if sl == nil {
			sl.add("")
		}
	}
}

//Adds a new cell
func (sl collection) add(string text) {
	var (
		cl  Cell
		int i = len(sl)
	)

	if len(name) <= 0 {
		name = "Slide" + fmt.Sprint(i)
	}

	cl.Init()

	//keep the pointer/dereference (i'm not sure which it is)
	//problems occur otherwise
	//*sl = append(*sl, &cl)
	sl = append(sl, &cl)

	//seperate image object in QML
	cl.image.qmlImage.Set("source", fmt.Sprintf("image://images/cell;%d", cl.index))
	cl.setSignal()
	//give QML the text

}

//(slide) remove copied from github.com/golang/go/wiki/SliceTricks
func (sl collection) remove(i int) {
	cl := sl[i]
	cl.text = ""
	cl.image.qmlImage.Destroy()
	cl.qmlObject.Destroy()
	cl.image.img.Destroy()
	MainWindow.ObjectByName("gridRect").Set("count", MainWindow.ObjectByName("gridRect").Int("count")-1)
	cl.index = -1

	copy(sl[i:], sl[i+1:])
	sl[len(sl)-1] = nil // or the zero value of T
	sl = sl[:len(sl)-1]

	//*sl, (*sl)[len((*sl))-1] = append((*sl)[:i], (*sl)[i+1:]...), nil
}

func (sl collection) destroy() {
	for i := len(sl); i > 0; i-- {
		sl.remove(i - 1)
	}
}

func (cl *Cell) Init() {
	cl.text = `hello this is text`
	cl.index = -1
	cl.fnt.color, cl.fnt.outlineColor = color.RGBA{0, 0, 0, 1}, color.RGBA{1, 1, 1, 1}
	cl.fnt.name = "none"
	cl.fnt.outline = false
	cl.fnt.outlineSize = 1
	cl.fnt.size = 35
	cl.fnt.x, cl.fnt.y = 10, 30

	cl.qmlObject = cellQml.Create(engine.Context())
	cl.image.qmlImage = imgQml.Create(engine.Context())

	//load image
	cl.image.img = imagick.NewMagickWand()
	cl.image.img.ReadImage("logo:")

}

func (cl *Cell) Select() {
	selectedCell = cl.index
	cl.qmlObject.ObjectByName("cellMouse").Call("selected")

}

//not really needed
func (cl Cell) String() string {
	return fmt.Sprintf("Index: %d \nText:  %s\n", cl.index, cl.text)
}

func (bl *Bool) Flip() {
	*bl = !*bl
}

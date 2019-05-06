// PresentationApp project main.go
//go:generate genqrc qml

package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/limetext/qml-go"
)

type Cell struct {
	Font                   Font
	image                  Image
	index, collectionIndex int
	text                   string
	textVisible            bool
}

type collection struct {
	collection []*Cell
	title      string
}

type Font struct {
	color                   color.RGBA
	name                    string
	path                    string
	outline                 bool
	outlineColor            color.RGBA
	outlineSize, size, x, y float64
}

type Image struct {
	path     string
	original image.Image
	resized  image.Image
}

type qmlVar struct {
	FontList   string
	Verses     string
	VerseOrder string
}

type service []collection

var (
	currentService = new(service)
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
	findFonts()

	engine.Context().SetVar("go", QML)
	engine.AddImageProvider("images", imgProvider)

	err = qmlWindows()
	if err != nil {
		return err
	}

	currentService.Init(1)

	//image is ready for imageprovider
	imgready = true

	displayImg = DisplayWindow.Root().ObjectByName("displayImage")
	serviceObject = serviceQml.Create(engine.Context())
	serviceObject.Set("parent", MainWindow.ObjectByName("scview"))
	serviceObject.Call("addLst", "not")

	//edtQmlShow()
	qml.RunMain(glInit)
	MainWindow.Wait()
	slides.destroy()
	fmt.Println(len(*currentService))

	return nil
}

func (sv *service) Init(num int) {
	if num <= 0 {
		num = 1
	}

	for index := 0; index < num; index++ {
		if sv == nil {
			sv.add("")
		}
	}
}

func (sv *service) add(name string) {
	var (
		sl collection
		i  = len(*sv)
	)

	if len(name) <= 0 {
		name = "Song: " + fmt.Sprint(i)
	}

	sl.init(1)
	*sv = append(*sv, sl)
	//serviceObject.Call("addLst", name)
}

func (sv *service) remove(i int) {
	(*sv)[i].destroy()

	copy((*sv)[i:], (*sv)[i+1:])
	(*sv)[len(*sv)-1] = nil // or the zero value of T
	*sv = (*sv)[:len(*sv)-1]

}

func (sv *service) destroy() {
	for i := len(*sv); i > 0; i-- {
		sv.remove(i - 1)
	}
}

func (sl *collection) init(num int) {
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
func (sl *collection) add(text string) {
	var (
		cl Cell
		i  = len(*sl)
	)

	if len(text) <= 0 {
		text = "Slide" + fmt.Sprint(i)
	}

	cl.Init()

	//keep the pointer/dereference (i'm not sure which it is)
	//problems occur otherwise
	// now Im not an idiot and I know what this does
	*sl = append(*sl, &cl)
}

// remove copied from github.com/golang/go/wiki/SliceTricks
func (sl *collection) remove(i int) {
	cl := (*sl)[i]
	cl.text = ""
	MainWindow.ObjectByName("gridRect").Set("count", MainWindow.ObjectByName("gridRect").Int("count")-1)
	cl.index = -1

	copy((*sl)[i:], (*sl)[i+1:])
	(*sl)[len(*sl)-1] = nil // or the zero value of T
	(*sl) = (*sl)[:len(*sl)-1]

	//*sl, (*sl)[len((*sl))-1] = append((*sl)[:i], (*sl)[i+1:]...), nil
}

func (sl *collection) destroy() {
	for i := len(*sl); i > 0; i-- {
		sl.remove(i - 1)
	}
}

func (cl *Cell) Init() {
	cl.text = `hello this is text`
	cl.index = -1
	cl.Font.color, cl.Font.outlineColor = color.RGBA{0, 0, 0, 1}, color.RGBA{1, 1, 1, 1}
	cl.Font.name = "none"
	cl.Font.outline = false
	cl.Font.outlineSize = 1
	cl.Font.size = 35
	cl.Font.x, cl.Font.y = 10, 30
}

//not really needed
func (cl Cell) String() string {
	return fmt.Sprintf("Index: %d \nText:  %s\n", cl.index, cl.text)
}

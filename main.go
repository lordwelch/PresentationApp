// PresentationApp project main.go
package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/gographics/imagick/imagick"
	"github.com/kardianos/osext"
	"github.com/lordwelch/qml"
)

type cell struct {
	text    string
	img     imagick.Image
	qmlimg  qml.Object
	qmlcell qml.Object
	index   int
}
type slide []*cell

var (
	x0, y0      int
	path        string
	qimg        qml.Object
	textEdit    qml.Object
	cellQml     qml.Object
	window      *qml.Window
	win         *glfw.Window
	slides      slide
	err         error
	monitors    []*glfw.Monitor
	projMonitor *glfw.Monitor
	mw1         *imagick.MagickWand
	tex1        uint32
)

func main() {

	if err = qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	defer glfw.Terminate()
	fmt.Println("test")

}

func run() error {
	var mainQml qml.Object
	imagick.Initialize()
	defer imagick.Terminate()
	mw1 = imagick.NewMagickWand()

	err = mw1.ReadImage("logo:")
	if err != nil {
		panic(err)
	}

	engine := qml.NewEngine()
	engine.AddImageProvider("images", imgProvider)

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
	slides.addCell()
	setSignals()

	window.Show()
	qml.RunMain(glInit)

	window.Wait()

	mw1.Destroy()
	imagick.Terminate()
	return nil
}

func setupScene() {

	gl.ClearColor(0.1, 0.5, 0.9, 0.0)
	mw1 = resizeImage(mw1, x0, y0, true, true)

	tex1 = newTexture(*slides[0].getImage(x0, y0))

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-1, 1, -1, 1, 1.0, 10.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

func drawSlide() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0, 0, -3.0)

	gl.Begin(gl.QUADS)

	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-1, 1, 0)

	gl.TexCoord2f(1, 0)
	gl.Vertex3f(1, 1, 0)

	gl.TexCoord2f(1, 1)
	gl.Vertex3f(1, -1, 0)

	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-1, -1, 0)

	gl.End()

}

func newTexture(rgba image.RGBA) uint32 {
	var texture1 uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture1)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture1
}

func checkMon() {
	monitors = glfw.GetMonitors()
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.AutoIconify, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)

	if i := len(monitors); i < 2 {
		fmt.Println("You only have 1 monitor!!!!!!!!!!! :-P")
		x0 = 800
		y0 = 600

		win, err = glfw.CreateWindow(x0, y0, "Cube", nil, nil)
		if err != nil {
			panic(err)
		}
		projMonitor = monitors[0]
	} else {
		fmt.Printf("You have %d monitors\n", i)
		x0 = monitors[1].GetVideoMode().Width
		y0 = monitors[1].GetVideoMode().Height
		win, err = glfw.CreateWindow(x0, y0, "Cube", nil, nil)
		fmt.Println(win.GetPos())
		win.SetPos(monitors[1].GetPos())
		fmt.Println(x0, y0)
		if err != nil {
			panic(err)
		}
		projMonitor = monitors[1]

	}
	monitorInfo()

}

func monitorInfo() {
	for _, mon := range monitors {
		fmt.Printf("monitor name: %s\n", mon.GetName())
		i, t := mon.GetPos()
		fmt.Printf("position X: %d  Y: %d\n", i, t)
	}

}

func glInit() {
	if err = glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	checkMon()

	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}
	setupScene()
	win.SetPos(projMonitor.GetPos())

	qml.Func1 = func() int {
		if !win.ShouldClose() {
			drawSlide()
			win.SwapBuffers()
			glfw.PollEvents()
			return 0
		} else {
			win.Hide()
			win.Destroy()
			return 1
		}
	}

}

func setSignals() {
	window.ObjectByName("btnAdd").On("clicked", func() {
		slides.addCell()
	})

	window.On("closing", func() {
		fmt.Println(window.Bool("cls"))
		win.Hide()
		window.Set("cls", true)

	})

	textEdit.ObjectByName("textEdit1").On("focusChanged", func(focus bool) {
		var (
			str string
			cel *cell
		)

		if !focus {

			str = textEdit.ObjectByName("textEdit1").String("text")
			cel = slides[textEdit.Int("cell")]
			if textEdit.Bool("txt") {
				cel.qmlcell.ObjectByName("cellText").Set("text", str)
				cel.text = str
			}
		}
	})

}

func (cl cell) getImage(x, y int) (img *image.RGBA) {
	mw := imagick.NewMagickWandFromImage(&cl.img)
	if (x == 0) || (y == 0) {
		x = int(mw.GetImageWidth())
		y = int(mw.GetImageHeight())
	}

	mw = resizeImage(mw, x, y, true, true)
	img = image.NewRGBA(image.Rect(0, 0, int(x), int(y)))
	if img.Stride != img.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	//	TPix, _ :=
	Tpix, _ := mw.ExportImagePixels(0, 0, uint(x), uint(y), "RGBA", imagick.PIXEL_CHAR)
	img.Pix = Tpix.([]uint8)
	defer mw.Destroy()
	return
}

func (cl *cell) setSignal() {
	cl.qmlcell.ObjectByName("cellMouse").On("doubleClicked", func() {

		textEdit.Set("cell", cl.index)
		textEdit.Set("x", cl.qmlcell.Int("x")+4)
		textEdit.Set("y", cl.qmlcell.Int("y")+4)
		textEdit.Set("width", cl.qmlcell.Int("width"))
		textEdit.Set("height", cl.qmlcell.Int("height"))
		textEdit.Set("opacity", 100)
		textEdit.Set("visible", true)
		textEdit.ObjectByName("textEdit1").Set("focus", true)
		textEdit.Set("enabled", true)

		textEdit.ObjectByName("textEdit1").Set("text", cl.text)
	})

}

func (sl *slide) addCell( /*cl *cell*/ ) {
	var cl cell

	cl.index = len(*sl)
	window.ObjectByName("gridRect").Set("count", window.ObjectByName("gridRect").Int("count")+1)
	cl.qmlcell = cellQml.Create(nil)
	cl.qmlcell.Set("objectName", fmt.Sprintf("cellRect%d", len(*sl)))
	cl.qmlcell.Set("parent", window.ObjectByName("data1"))
	cl.qmlcell.Set("index", cl.index)

	mw := imagick.NewMagickWand()
	mw.ReadImage("logo:")
	cl.img = *mw.GetImageFromMagickWand()

	cl.text = "testing 1... 2... 3..."
	cl.qmlcell.ObjectByName("cellText").Set("text", cl.text)
	*sl = append(*sl, &cl)
	cl.setSignal()

	cl.qmlimg = qimg.Create(nil)
	cl.qmlimg.Set("objectName", fmt.Sprintf("cellImg%d", cl.index))
	cl.qmlimg.Set("source", fmt.Sprintf("image://images/%d", cl.index))
	cl.qmlimg.Set("parent", window.ObjectByName("data2"))
	cl.qmlimg.Set("index", cl.index)

}

func (cl cell) String() string {
	return fmt.Sprintf("Index: %d \nText:  %s\n", cl.index, cl.text)
}

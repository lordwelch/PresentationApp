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
	img     image.RGBA
	qmlcell qml.Object
	index   int
}
type slide []*cell

var (
	x, y        int
	path        string
	textEdit    qml.Object
	cellQml     qml.Object
	window      *qml.Window
	win         *glfw.Window
	slides      slide
	err         error
	monitors    []*glfw.Monitor
	projMonitor *glfw.Monitor
	mw1, mw2    *imagick.MagickWand
	tex1        uint32
	//drawSlide func()
)

func main() {

	if err = qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	defer glfw.Terminate()
	fmt.Println("test")

}

/*func init() {
	// GLFW event handling must run on the main OS thread
	//runtime.LockOSThread()
}*/

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

	window = mainQml.CreateWindow(nil)

	textEdit = window.ObjectByName("textEdit")
	slides.addCell()
	setSignals()

	window.Show()
	qml.RunMain(glInit)

	window.Wait()
	//win.Destroy()
	mw2.Destroy()
	imagick.Terminate()
	return nil
}

func setupScene() {

	gl.ClearColor(0.1, 0.5, 0.9, 0.0)
	mw2 = resizeImage(mw1, x, y, true, true)

	tex1 = newTexture(*mw2)

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

func newTexture(mw imagick.MagickWand) uint32 {
	x1 := mw.GetImageWidth()
	y1 := mw.GetImageHeight()
	rgba := image.NewRGBA(image.Rect(0, 0, int(x1), int(y1)))
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	TPix, _ := mw.ExportImagePixels(0, 0, x1, y1, "RGBA", imagick.PIXEL_CHAR)
	rgba.Pix = TPix.([]uint8)

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
		x = 800
		y = 600
		win, err = glfw.CreateWindow(x, y, "Cube", nil, nil)
		if err != nil {
			panic(err)
		}
		projMonitor = monitors[0]
	} else {
		fmt.Printf("You have %d monitors\n", i)
		x = monitors[1].GetVideoMode().Width
		y = monitors[1].GetVideoMode().Height
		win, err = glfw.CreateWindow(x, y, "Cube", nil, nil)
		fmt.Println(win.GetPos())
		win.SetPos(monitors[1].GetPos())
		fmt.Println(x, y)
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
		//win.Destroy()
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
				//fmt.Println("haha.....:-P")
				cel.text = str
				//fmt.Println("----->"+cel.text, str, textEdit.Int("cell"))
			}
		}
	})

}

func (cl *cell) setSignal() {
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
		textEdit.Set("enabled", true)

		textEdit.ObjectByName("textEdit1").Set("text", cl.text)
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
	*sl = append(*sl, &cl)
	//fmt.Println("add Cell!!!!! :-P")
	cl.setSignal()

}

func (cl cell) String() string {
	return fmt.Sprintf("Index: %d \nText:  %s\n", cl.index, cl.text)
}
